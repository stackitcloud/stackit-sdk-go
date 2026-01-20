package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
)

func TestContinuousRefreshToken(t *testing.T) {
	// The times here are in the order of miliseconds (so they run faster)
	// For this to work, we need to increase precision of the expiration timestamps
	jwt.TimePrecision = time.Millisecond

	// Refresher settings
	timeStartBeforeTokenExpiration := 0 * time.Second
	timeBetweenContextCheck := 50 * time.Millisecond
	timeBetweenTries := 500 * time.Millisecond

	// All generated acess tokens will have this time to live
	accessTokensTimeToLive := 1 * time.Second

	tests := []struct {
		desc                  string
		contextClosesIn       time.Duration
		doError               error
		expectedNumberDoCalls int
	}{
		{
			desc:                  "update access token never",
			contextClosesIn:       900 * time.Millisecond, // Should allow no refresh
			expectedNumberDoCalls: 0,
		},
		{
			desc:                  "update access token once",
			contextClosesIn:       1900 * time.Millisecond, // Should allow one refresh
			expectedNumberDoCalls: 1,
		},
		{
			desc:                  "update access token twice",
			contextClosesIn:       2900 * time.Millisecond, // Should allow two refreshes
			expectedNumberDoCalls: 2,
		},
		{
			desc:                  "context canceled at start",
			contextClosesIn:       0,
			expectedNumberDoCalls: 0,
		},
		{
			desc:                  "context canceled before start",
			contextClosesIn:       -150 * time.Millisecond,
			expectedNumberDoCalls: 0,
		},
		{
			desc:                  "context canceled before token refresh",
			contextClosesIn:       50 * time.Millisecond,
			expectedNumberDoCalls: 0,
		},
		{
			desc:                  "refresh token fails - error",
			contextClosesIn:       1900 * time.Millisecond,
			doError:               fmt.Errorf("something went wrong"),
			expectedNumberDoCalls: 1,
		},
		{
			desc:            "refresh token fails - API non-5xx error",
			contextClosesIn: 1900 * time.Millisecond,
			doError: &oapierror.GenericOpenAPIError{
				StatusCode: http.StatusBadRequest,
			},
			expectedNumberDoCalls: 1,
		},
		{
			desc:            "refresh token fails - API 5xx error",
			contextClosesIn: 2900 * time.Millisecond,
			doError: &oapierror.GenericOpenAPIError{
				StatusCode: http.StatusInternalServerError,
			},
			expectedNumberDoCalls: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			accessToken, err := signToken(accessTokensTimeToLive)
			if err != nil {
				t.Fatalf("failed to sign access token: %v", err)
			}
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, tt.contextClosesIn)
			defer cancel()

			authFlow := &fakeAuthFlow{
				backgroundTokenRefreshContext: ctx,
				doError:                       tt.doError,
				accessTokensTimeToLive:        accessTokensTimeToLive,
				accessToken:                   accessToken,
			}

			refresher := &continuousTokenRefresher{
				flow:                           authFlow,
				timeStartBeforeTokenExpiration: timeStartBeforeTokenExpiration,
				timeBetweenContextCheck:        timeBetweenContextCheck,
				timeBetweenTries:               timeBetweenTries,
			}

			err = refresher.continuousRefreshToken()
			if err == nil {
				t.Fatalf("routine finished with non-nil error")
			}
			numberDoCalls := authFlow.getTokenCalls()
			if numberDoCalls != tt.expectedNumberDoCalls {
				t.Fatalf("expected %d calls to API to refresh token, got %d", tt.expectedNumberDoCalls, numberDoCalls)
			}
		})
	}
}

// Tests if
//   - continuousRefreshToken() updates access token using the refresh token
//   - The access token can be accessed while continuousRefreshToken() is trying to update it
func TestContinuousRefreshTokenConcurrency(t *testing.T) {
	// The times here are in the order of miliseconds (so they run faster)
	// For this to work, we need to increase precision of the expiration timestamps
	jwt.TimePrecision = time.Millisecond

	// Test plan:
	// 1) continuousRefreshToken() will trigger a token update. It will be blocked in the mockDo() routine (defined below)
	// 2) After continuousRefreshToken() is blocked, a request will be made using the key flow. That request should carry the access token (shouldn't be blocked just because continuousRefreshToken() is trying to refresh the token)
	// 3) After the request is successful, continuousRefreshToken() will be unblocked
	// 4) After waiting a bit, a new request will be made using the key flow. That request should carry the new access token

	// Where we're at in the test plan:
	// - Starts at 0
	// - Is set to 1 before continuousRefreshToken() is called
	// - Is set to 2 once the continuousRefreshToken() is blocked
	// - Is set to 3 once the first request goes through and is checked
	// - Is set to 4 after a small wait after continuousRefreshToken() is unblocked
	currentTestPhase := 0

	// Used to signal continuousRefreshToken() has been blocked
	chanBlockContinuousRefreshToken := make(chan bool)

	// Used to signal continuousRefreshToken() should be unblocked
	chanUnblockContinuousRefreshToken := make(chan bool)

	// The access token at the start
	accessTokenFirst, err := signToken(10 * time.Second)
	if err != nil {
		t.Fatalf("failed to create first access token: %v", err)
	}

	// The access token that will replace accessTokenFirst
	// Has a much longer expiration timestamp
	accessTokenSecond, err := signToken(time.Hour)
	if err != nil {
		t.Fatalf("failed to create second access token: %v", err)
	}

	if accessTokenFirst == accessTokenSecond {
		t.Fatalf("created tokens are equal")
	}

	// The refresh token used to update the access token
	refreshToken, err := signToken(time.Hour)
	if err != nil {
		t.Fatalf("failed to create refresh token: %v", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // This cancels the refresher goroutine

	// Extract host from tokenAPI constant for consistency
	tokenURL, _ := url.Parse(tokenAPI)
	tokenHost := tokenURL.Host

	// The Do() routine, that both the keyFlow and continuousRefreshToken() use to make their requests
	// The bools are used to make sure only one request goes through on each test phase
	doTestPhase1RequestDone := false
	doTestPhase2RequestDone := false
	doTestPhase4RequestDone := false
	mockDo := func(req *http.Request) (resp *http.Response, err error) {
		// Handle auth requests (token refresh)
		if req.URL.Host == tokenHost {
			switch currentTestPhase {
			default:
				// After phase 1, allow additional auth requests but don't fail the test
				// This handles the continuous nature of the refresh routine
				if currentTestPhase > 1 {
					// Return a valid response for any additional auth requests
					newAccessToken, err := signToken(time.Hour)
					if err != nil {
						t.Fatalf("Do call: failed to create additional access token: %v", err)
					}
					responseBodyStruct := TokenResponseBody{
						AccessToken:  newAccessToken,
						RefreshToken: refreshToken,
					}
					responseBody, err := json.Marshal(responseBodyStruct)
					if err != nil {
						t.Fatalf("Do call: failed to marshal additional response: %v", err)
					}
					response := &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(bytes.NewReader(responseBody)),
					}
					return response, nil
				}
				t.Fatalf("Do call: unexpected request during test phase %d", currentTestPhase)
				return nil, nil
			case 1: // Call by continuousRefreshToken()
				if doTestPhase1RequestDone {
					t.Fatalf("Do call: multiple requests during test phase 1")
				}
				doTestPhase1RequestDone = true

				currentTestPhase = 2
				chanBlockContinuousRefreshToken <- true

				// Wait until continuousRefreshToken() is to be unblocked
				<-chanUnblockContinuousRefreshToken

				if currentTestPhase != 3 {
					t.Fatalf("Do call: after unlocking refreshToken(), expected test phase to be 3, got %d", currentTestPhase)
				}

				// Check required fields are passed
				err = req.ParseForm()
				if err != nil {
					t.Fatalf("Do call: failed to parse body form: %v", err)
				}
				reqGrantType := req.Form.Get("grant_type")
				if reqGrantType != "refresh_token" {
					t.Fatalf("Do call: failed request to refresh token: call to refresh access expected to have grant type as %q, found %q instead", "refresh_token", reqGrantType)
				}
				reqRefreshToken := req.Form.Get("refresh_token")
				if reqRefreshToken != refreshToken {
					t.Fatalf("Do call: failed request to refresh token: call to refresh access token did not have the expected refresh token set")
				}

				// Return response with accessTokenSecond
				responseBodyStruct := TokenResponseBody{
					AccessToken:  accessTokenSecond,
					RefreshToken: refreshToken,
				}
				responseBody, err := json.Marshal(responseBodyStruct)
				if err != nil {
					t.Fatalf("Do call: failed request to refresh token: marshal access token response: %v", err)
				}
				response := &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader(responseBody)),
				}
				return response, nil
			}
		}

		// Handle regular HTTP requests
		switch currentTestPhase {
		default:
			t.Fatalf("Do call: unexpected request during test phase %d", currentTestPhase)
			return nil, nil
		case 2: // Call by tokenFlow, first request
			if doTestPhase2RequestDone {
				t.Fatalf("Do call: multiple requests during test phase 2")
			}
			doTestPhase2RequestDone = true

			// Check host and access token
			host := req.Host
			expectedHost := "first-request-url.com"
			if host != expectedHost {
				t.Fatalf("Do call: first request expected to have host %q, found %q", expectedHost, host)
			}
			authHeader := req.Header.Get("Authorization")
			expectedAuthHeader := fmt.Sprintf("Bearer %s", accessTokenFirst)
			if authHeader != expectedAuthHeader {
				t.Fatalf("Do call: first request didn't carry first access token. Expected: %s, Got: %s", expectedAuthHeader, authHeader)
			}

			// Return empty response
			response := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
			}
			return response, nil
		case 4: // Call by tokenFlow, second request
			if doTestPhase4RequestDone {
				t.Fatalf("Do call: multiple requests during test phase 4")
			}
			doTestPhase4RequestDone = true

			// Check host and access token
			host := req.Host
			expectedHost := "second-request-url.com"
			if host != expectedHost {
				t.Fatalf("Do call: second request expected to have host %q, found %q", expectedHost, host)
			}
			authHeader := req.Header.Get("Authorization")
			if authHeader != fmt.Sprintf("Bearer %s", accessTokenSecond) {
				t.Fatalf("Do call: second request didn't carry second access token")
			}

			// Return empty response
			response := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
			}
			return response, nil
		}
	}

	keyFlow := &KeyFlow{}
	privateKeyBytes, err := generatePrivateKey()
	if err != nil {
		t.Fatalf("Error generating private key: %s", err)
	}
	keyFlowConfig := &KeyFlowConfig{
		ServiceAccountKey: fixtureServiceAccountKey(),
		PrivateKey:        string(privateKeyBytes),
		AuthHTTPClient: &http.Client{
			Transport: mockTransportFn{mockDo},
		},
		HTTPTransport: mockTransportFn{mockDo}, // Use same mock for regular requests
		// Don't start continuous refresh automatically
		BackgroundTokenRefreshContext: nil,
	}
	err = keyFlow.Init(keyFlowConfig)
	if err != nil {
		t.Fatalf("failed to initialize key flow: %v", err)
	}

	// Set the token after initialization
	err = keyFlow.SetToken(accessTokenFirst, refreshToken)
	if err != nil {
		t.Fatalf("failed to set token: %v", err)
	}

	// Set the context for continuous refresh
	keyFlow.config.BackgroundTokenRefreshContext = ctx

	// Create a custom refresher with shorter timing for the test
	refresher := &continuousTokenRefresher{
		flow:                           keyFlow,
		timeStartBeforeTokenExpiration: 9 * time.Second, // Start 9 seconds before expiration
		timeBetweenContextCheck:        5 * time.Millisecond,
		timeBetweenTries:               40 * time.Millisecond,
	}

	// TEST START
	currentTestPhase = 1
	// Ignore returned error as expected in test
	go func() {
		_ = refresher.continuousRefreshToken()
	}()

	// Wait until continuousRefreshToken() is blocked
	<-chanBlockContinuousRefreshToken

	if currentTestPhase != 2 {
		t.Fatalf("Unexpected test phase %d after continuousRefreshToken() was blocked", currentTestPhase)
	}

	// Perform first request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://first-request-url.com", http.NoBody)
	if err != nil {
		t.Fatalf("Create first request failed: %v", err)
	}
	resp, err := keyFlow.RoundTrip(req)
	if err != nil {
		t.Fatalf("Perform first request failed: %v", err)
	}
	err = resp.Body.Close()
	if err != nil {
		t.Fatalf("First request body failed to close: %v", err)
	}

	// Unblock continuousRefreshToken()
	currentTestPhase = 3
	chanUnblockContinuousRefreshToken <- true

	// Wait for a bit
	time.Sleep(10 * time.Millisecond)
	currentTestPhase = 4

	// Perform second request
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, "http://second-request-url.com", http.NoBody)
	if err != nil {
		t.Fatalf("Create second request failed: %v", err)
	}
	resp, err = keyFlow.RoundTrip(req)
	if err != nil {
		t.Fatalf("Second request failed: %v", err)
	}
	err = resp.Body.Close()
	if err != nil {
		t.Fatalf("Second request body failed to close: %v", err)
	}
}

func signToken(expiration time.Duration) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
	}).SignedString([]byte("test"))
}

var _ AuthFlow = &fakeAuthFlow{}

type fakeAuthFlow struct {
	backgroundTokenRefreshContext context.Context
	tokenCounter                  int
	doError                       error
	accessTokensTimeToLive        time.Duration
	accessToken                   string
}

func (f *fakeAuthFlow) RoundTrip(_ *http.Request) (*http.Response, error) {
	return nil, nil
}
func (f *fakeAuthFlow) GetAccessToken() (string, error) {
	expired, err := tokenExpired(f.accessToken, 0)
	if err != nil {
		return "", err
	}
	if !expired {
		return f.accessToken, nil
	}
	f.tokenCounter++
	if f.doError != nil {
		return "", f.doError
	}
	accessToken, err := signToken(f.accessTokensTimeToLive)
	if err != nil {
		return "", f.doError
	}
	f.accessToken = accessToken
	return accessToken, nil
}

func (f *fakeAuthFlow) RefreshAccessToken() error {
	_, err := f.GetAccessToken()
	return err
}

func (f *fakeAuthFlow) GetBackgroundTokenRefreshContext() context.Context {
	return f.backgroundTokenRefreshContext
}

func (f *fakeAuthFlow) getTokenCalls() int {
	return f.tokenCounter
}
