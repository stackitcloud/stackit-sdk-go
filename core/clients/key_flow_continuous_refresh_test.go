package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	timeStartBeforeTokenExpiration := 30 * time.Millisecond
	timeBetweenContextCheck := 10 * time.Millisecond
	timeBetweenTries := 20 * time.Millisecond

	// All generated acess tokens will have this time to live
	accessTokensTimeToLive := 100 * time.Millisecond

	tests := []struct {
		desc                  string
		contextClosesIn       time.Duration
		doError               error
		expectedNumberDoCalls int
	}{
		{
			desc:                  "update access token once",
			contextClosesIn:       100 * time.Millisecond,
			expectedNumberDoCalls: 1,
		},
		{
			desc:                  "update access token twice",
			contextClosesIn:       180 * time.Millisecond,
			expectedNumberDoCalls: 2,
		},
		{
			desc:                  "context canceled at start",
			contextClosesIn:       0,
			expectedNumberDoCalls: 0,
		},
		{
			desc:                  "context canceled before start",
			contextClosesIn:       -100 * time.Millisecond,
			expectedNumberDoCalls: 0,
		},
		{
			desc:                  "context canceled before token refresh",
			contextClosesIn:       10 * time.Millisecond,
			expectedNumberDoCalls: 0,
		},
		{
			desc:                  "refresh token fails - non-API error",
			contextClosesIn:       100 * time.Millisecond,
			doError:               fmt.Errorf("something went wrong"),
			expectedNumberDoCalls: 1,
		},
		{
			desc:            "refresh token fails - API non-5xx error",
			contextClosesIn: 100 * time.Millisecond,
			doError: &oapierror.GenericOpenAPIError{
				StatusCode: http.StatusBadRequest,
			},
			expectedNumberDoCalls: 1,
		},
		{
			desc:            "refresh token fails - API 5xx error",
			contextClosesIn: 100 * time.Millisecond,
			doError: &oapierror.GenericOpenAPIError{
				StatusCode: http.StatusInternalServerError,
			},
			expectedNumberDoCalls: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			numberDoCalls := 0
			mockDo := func(client *http.Client, req *http.Request, cfg *RetryConfig) (resp *http.Response, err error) {
				numberDoCalls++

				if tt.doError != nil {
					return nil, tt.doError
				}

				accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokensTimeToLive)),
				}).SignedString([]byte("test"))
				if err != nil {
					t.Fatalf("Do call: failed to create access token: %v", err)
				}

				responseBodyStruct := TokenResponseBody{
					AccessToken: accessToken,
				}
				responseBody, err := json.Marshal(responseBodyStruct)
				if err != nil {
					t.Fatalf("Do call: failed to marshal response: %v", err)
				}
				response := &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader(responseBody)),
				}
				return response, nil
			}

			accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokensTimeToLive)),
			}).SignedString([]byte("test"))
			if err != nil {
				t.Fatalf("failed to create access token: %v", err)
			}

			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, tt.contextClosesIn)
			defer cancel()

			keyFlow := &KeyFlow{
				config: &KeyFlowConfig{
					ClientRetry:                   NewRetryConfig(),
					BackgroundTokenRefreshContext: ctx,
				},
				doer: mockDo,
				token: &TokenResponseBody{
					AccessToken: accessToken,
				},
			}

			refresher := &continuousTokenRefresher{
				keyFlow:                        keyFlow,
				timeStartBeforeTokenExpiration: timeStartBeforeTokenExpiration,
				timeBetweenContextCheck:        timeBetweenContextCheck,
				timeBetweenTries:               timeBetweenTries,
			}

			err = refresher.continuousRefreshToken()
			if err == nil {
				t.Fatalf("routine finished with non-nil error")
			}
			if numberDoCalls != tt.expectedNumberDoCalls {
				t.Fatalf("expected %d calls to API to refresh token, got %d", tt.expectedNumberDoCalls, numberDoCalls)
			}
		})
	}
}

// Tests if
//   - continuousRefreshToken() changes the token
//   - The access token can be accessed while continuousRefreshToken() is trying to update it
func TestContinuousRefreshTokenConcurrency(t *testing.T) {
	// The times here are in the order of miliseconds (so they run faster)
	// For this to work, we need to increase precision of the expiration timestamps
	jwt.TimePrecision = time.Millisecond

	// Test plan:
	// 1) continuousRefreshToken() will trigger a token update. It will be blocked in the mockDo() routine (defined below)
	// 2) After continuousRefreshToken() is blocked, a request will be made using the key flow. That request should carry the access token (shouldn't be blocked just because continuousRefreshToken() is trying to update)
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
	accessTokenFirst, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(100 * time.Millisecond)),
	}).SignedString([]byte("token-first"))
	if err != nil {
		t.Fatalf("failed to create first access token: %v", err)
	}

	// The access token that will replace accessTokenFirst
	// Has a much longer expiration timestamp
	accessTokenSecond, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}).SignedString([]byte("token-second"))
	if err != nil {
		t.Fatalf("failed to create second access token: %v", err)
	}

	if accessTokenFirst == accessTokenSecond {
		t.Fatalf("created tokens are equal")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // This cancels the refresher goroutine

	// The "do" routine, that both the keyFlow and continuousRefreshToken() use to make their requests
	// The bools are used to make sure only one request goes through on each test phase
	doTestPhase1RequestDone := false
	doTestPhase2RequestDone := false
	doTestPhase4RequestDone := false
	mockDo := func(client *http.Client, req *http.Request, cfg *RetryConfig) (resp *http.Response, err error) {
		if currentTestPhase == 1 { // Call by continuousRefreshToken()
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

			// Return response with accessTokenSecond
			responseBodyStruct := TokenResponseBody{
				AccessToken: accessTokenSecond,
			}
			responseBody, err := json.Marshal(responseBodyStruct)
			if err != nil {
				t.Fatalf("Do call: failed to marshal access token response: %v", err)
			}
			response := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(responseBody)),
			}
			return response, nil
		} else if currentTestPhase == 2 { // Call by tokenFlow, first request
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
			if authHeader != fmt.Sprintf("Bearer %s", accessTokenFirst) {
				t.Fatalf("Do call: first request didn't carry first access token")
			}

			// Return empty response
			response := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
			}
			return response, nil
		} else if currentTestPhase == 4 { // Call by tokenFlow, second request
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
		t.Fatalf("Do call: unexpected request during test phase %d", currentTestPhase)
		return nil, nil
	}

	keyFlow := &KeyFlow{
		client: &http.Client{},
		config: &KeyFlowConfig{
			ClientRetry:                   NewRetryConfig(),
			BackgroundTokenRefreshContext: ctx,
		},
		doer: mockDo,
		token: &TokenResponseBody{
			AccessToken: accessTokenFirst,
		},
	}

	// TEST START
	currentTestPhase = 1
	go continuousRefreshToken(keyFlow)

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
