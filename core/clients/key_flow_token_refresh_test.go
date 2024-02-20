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

func TestRefreshToken(t *testing.T) {
	// The times here are in the order of miliseconds (so they run faster)
	// For this to work, we need to increase precision of the expiration timestamps
	jwt.TimePrecision = time.Millisecond

	// Refresher settings
	timeStartBeforeTokenExpiration := 30 * time.Millisecond
	timeBetweenContextCheck := 10 * time.Millisecond
	timeBetweenTries := 10 * time.Millisecond

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
			contextClosesIn:       150 * time.Millisecond,
			expectedNumberDoCalls: 2,
		},
		{
			desc:                  "context cancelled at start",
			contextClosesIn:       0,
			expectedNumberDoCalls: 0,
		},
		{
			desc:                  "context cancelled before start",
			contextClosesIn:       -100 * time.Millisecond,
			expectedNumberDoCalls: 0,
		},
		{
			desc:                  "context cancelled before token refresh",
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
			contextClosesIn: 95 * time.Millisecond,
			doError: &oapierror.GenericOpenAPIError{
				StatusCode: http.StatusInternalServerError,
			},
			expectedNumberDoCalls: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			numberDoCalls := 0
			mockDo := func(client *http.Client, req *http.Request, cfg *RetryConfig) (resp *http.Response, err error) {
				numberDoCalls += 1

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
					ClientRetry:                     NewRetryConfig(),
					TokenRefreshInBackgroundContext: ctx,
				},
				doer: mockDo,
				token: &TokenResponseBody{
					AccessToken: accessToken,
				},
			}

			refresher := &tokenRefresher{
				keyFlow:                        keyFlow,
				timeStartBeforeTokenExpiration: timeStartBeforeTokenExpiration,
				timeBetweenContextCheck:        timeBetweenContextCheck,
				timeBetweenTries:               timeBetweenTries,
			}

			err = refresher.refreshToken()
			if err == nil {
				t.Fatalf("routine finished with non-nil error")
			}
			if numberDoCalls != tt.expectedNumberDoCalls {
				t.Fatalf("expected %d calls to API to refresh token, got %d", tt.expectedNumberDoCalls, numberDoCalls)
			}
		})
	}
}
