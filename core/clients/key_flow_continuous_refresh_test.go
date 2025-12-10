package clients

import (
	"context"
	"fmt"
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

func (f *fakeAuthFlow) RoundTrip(req *http.Request) (*http.Response, error) {
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
func (f *fakeAuthFlow) GetBackgroundTokenRefreshContext() context.Context {
	return f.backgroundTokenRefreshContext
}

func (f *fakeAuthFlow) getTokenCalls() int {
	return f.tokenCounter
}
