package clients

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

var (
	testSigningKey = []byte(`Test`)
)

const testBearerToken = "eyJhbGciOiJub25lIn0.eyJleHAiOjIxNDc0ODM2NDd9." //nolint:gosec // linter false positive

func fixtureServiceAccountKey(mods ...func(*ServiceAccountKeyResponse)) *ServiceAccountKeyResponse {
	validUntil := time.Now().Add(time.Hour)
	serviceAccountKeyResponse := &ServiceAccountKeyResponse{
		Active:    true,
		CreatedAt: time.Now(),
		Credentials: &ServiceAccountKeyCredentials{
			Aud: "https://stackit-service-account-prod.apps.01.cf.eu01.stackit.cloud",
			Iss: "stackit@sa.stackit.cloud",
			Kid: uuid.New().String(),
			Sub: uuid.New(),
		},
		ID:           uuid.New(),
		KeyAlgorithm: "RSA_2048",
		KeyOrigin:    "USER_PROVIDED",
		KeyType:      "USER_MANAGED",
		PublicKey:    "...",
		ValidUntil:   &validUntil,
	}
	for _, mod := range mods {
		mod(serviceAccountKeyResponse)
	}
	return serviceAccountKeyResponse
}

func generatePrivateKey() ([]byte, error) {
	// Generate a new RSA key pair with a size of 2048 bits
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	// Encode the private key in PEM format
	privKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}

	// Print the private and public keys
	return pem.EncodeToMemory(privKeyPEM), nil
}

func TestKeyFlowInit(t *testing.T) {
	tests := []struct {
		name              string
		serviceAccountKey *ServiceAccountKeyResponse
		genPrivateKey     bool
		invalidPrivateKey bool
		wantErr           bool
	}{
		{
			name:              "ok-provided-private-key",
			serviceAccountKey: fixtureServiceAccountKey(),
			genPrivateKey:     true,
			wantErr:           false,
		},
		{
			name:              "missing_private_key",
			serviceAccountKey: fixtureServiceAccountKey(),
			genPrivateKey:     false,
			wantErr:           true,
		},
		{
			name:              "missing_service_account_key",
			serviceAccountKey: nil,
			genPrivateKey:     true,
			wantErr:           true,
		},
		{
			name:              "invalid_private_key",
			serviceAccountKey: fixtureServiceAccountKey(),
			genPrivateKey:     false,
			invalidPrivateKey: true,
			wantErr:           true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &KeyFlow{}
			cfg := &KeyFlowConfig{}
			t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY", "")
			if tt.genPrivateKey {
				privateKeyBytes, err := generatePrivateKey()
				if err != nil {
					t.Fatalf("Error generating private key: %s", err)
				}
				cfg.PrivateKey = string(privateKeyBytes)
			}
			if tt.invalidPrivateKey {
				cfg.PrivateKey = "invalid_key"
			}

			cfg.ServiceAccountKey = tt.serviceAccountKey
			if err := c.Init(cfg); (err != nil) != tt.wantErr {
				t.Errorf("KeyFlow.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			if c.config == nil {
				t.Error("config is nil")
			}
		})
	}
}

func TestSetToken(t *testing.T) {
	tests := []struct {
		name         string
		tokenInvalid bool
		refreshToken string
		wantErr      bool
	}{
		{
			name:         "ok",
			tokenInvalid: false,
			refreshToken: "refresh_token",
			wantErr:      false,
		},
		{
			name:         "invalid_token",
			tokenInvalid: true,
			refreshToken: "",
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var accessToken string
			var err error

			timestamp := time.Now().Add(24 * time.Hour)
			if tt.tokenInvalid {
				accessToken = "foo"
			} else {
				accessTokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(timestamp)})
				accessToken, err = accessTokenJWT.SignedString(testSigningKey)
				if err != nil {
					t.Fatalf("get test access token as string: %s", err)
				}
			}

			c := &KeyFlow{}
			err = c.SetToken(accessToken, tt.refreshToken)

			if (err != nil) != tt.wantErr {
				t.Errorf("KeyFlow.SetToken() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				expectedKeyFlowToken := &TokenResponseBody{
					AccessToken:  accessToken,
					ExpiresIn:    int(timestamp.Unix()),
					RefreshToken: tt.refreshToken,
					Scope:        defaultScope,
					TokenType:    defaultTokenType,
				}
				if !cmp.Equal(expectedKeyFlowToken, c.token) {
					t.Errorf("The returned result is wrong. Expected %+v, got %+v", expectedKeyFlowToken, c.token)
				}
			}
		})
	}
}

func TestTokenExpired(t *testing.T) {
	tests := []struct {
		desc              string
		tokenInvalid      bool
		tokenExpiresAt    time.Time
		expectedErr       bool
		expectedIsExpired bool
	}{
		{
			desc:              "token valid",
			tokenExpiresAt:    time.Now().Add(time.Hour),
			expectedErr:       false,
			expectedIsExpired: false,
		},
		{
			desc:              "token expired",
			tokenExpiresAt:    time.Now().Add(-time.Hour),
			expectedErr:       false,
			expectedIsExpired: true,
		},
		{
			desc:         "token invalid",
			tokenInvalid: true,
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			var err error
			token := "foo"
			if !tt.tokenInvalid {
				token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(tt.tokenExpiresAt),
				}).SignedString([]byte("test"))
				if err != nil {
					t.Fatalf("failed to create token: %v", err)
				}
			}

			isExpired, err := tokenExpired(token)
			if err != nil && !tt.expectedErr {
				t.Fatalf("failed on valid input: %v", err)
			}
			if err != nil {
				return
			}
			if isExpired != tt.expectedIsExpired {
				t.Fatalf("expected isValid to be %t, got %t", tt.expectedIsExpired, isExpired)
			}
		})
	}
}

func TestRequestToken(t *testing.T) {
	testCases := []struct {
		name          string
		grant         string
		assertion     string
		mockResponse  *http.Response
		mockError     error
		expectedError error
	}{
		{
			name:      "Success",
			grant:     "test_grant",
			assertion: "test_assertion",
			mockResponse: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"access_token": "test_token"}`)),
			},
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "Error",
			grant:         "test_grant",
			assertion:     "test_assertion",
			mockResponse:  nil,
			mockError:     fmt.Errorf("request error"),
			expectedError: fmt.Errorf("request error"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			c := &KeyFlow{
				authClient: &http.Client{
					Transport: mockTransportFn{func(_ *http.Request) (*http.Response, error) {
						return tt.mockResponse, tt.mockError
					}},
				},
				config: &KeyFlowConfig{},
				rt:     http.DefaultTransport,
			}

			res, err := c.requestToken(tt.grant, tt.assertion)
			defer func() {
				if res != nil {
					tempErr := res.Body.Close()
					if tempErr != nil {
						t.Errorf("closing request token response: %s", tempErr.Error())
					}
				}
			}()
			if tt.expectedError != nil {
				if err == nil {
					t.Errorf("Expected error '%v' but no error was returned", tt.expectedError)
				} else if errors.Is(err, tt.expectedError) {
					t.Errorf("Error is not correct. Expected %v, got %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but error was returned: %v", err)
				}
				if !cmp.Equal(tt.mockResponse, res, cmp.AllowUnexported(strings.Reader{})) {
					t.Errorf("The returned result is wrong. Expected %v, got %v", tt.mockResponse, res)
				}
			}
		})
	}
}

func TestKeyFlow_Do(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		keyFlow   *KeyFlow
		handlerFn func(tb testing.TB) http.HandlerFunc
		want      int
		wantErr   bool
	}{
		{
			name:    "success",
			keyFlow: &KeyFlow{rt: http.DefaultTransport, config: &KeyFlowConfig{}},
			handlerFn: func(tb testing.TB) http.HandlerFunc {
				tb.Helper()

				return func(w http.ResponseWriter, r *http.Request) {
					if r.Header.Get("Authorization") != "Bearer "+testBearerToken {
						tb.Errorf("expected Authorization header to be 'Bearer %s', but got %s", testBearerToken, r.Header.Get("Authorization"))
					}

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					_, _ = fmt.Fprintln(w, `{"status":"ok"}`)
				}
			},
			want:    http.StatusOK,
			wantErr: false,
		},
		{
			name:    "success with code 500",
			keyFlow: &KeyFlow{rt: http.DefaultTransport, config: &KeyFlowConfig{}},
			handlerFn: func(_ testing.TB) http.HandlerFunc {
				return func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "text/html")
					w.WriteHeader(http.StatusInternalServerError)
					_, _ = fmt.Fprintln(w, `<html>Internal Server Error</html>`)
				}
			},
			want:    http.StatusInternalServerError,
			wantErr: false,
		},
		{
			name: "success with custom transport",
			keyFlow: &KeyFlow{
				rt: mockTransportFn{
					fn: func(req *http.Request) (*http.Response, error) {
						req.Header.Set("User-Agent", "custom_transport")

						return http.DefaultTransport.RoundTrip(req)
					},
				},
				config: &KeyFlowConfig{},
			},
			handlerFn: func(tb testing.TB) http.HandlerFunc {
				tb.Helper()

				return func(w http.ResponseWriter, r *http.Request) {
					if r.Header.Get("User-Agent") != "custom_transport" {
						tb.Errorf("expected User-Agent header to be 'custom_transport', but got %s", r.Header.Get("User-Agent"))
					}

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					_, _ = fmt.Fprintln(w, `{"status":"ok"}`)
				}
			},
			want:    http.StatusOK,
			wantErr: false,
		},
		{
			name: "fail with custom proxy",
			keyFlow: &KeyFlow{
				rt: &http.Transport{
					Proxy: func(_ *http.Request) (*url.URL, error) {
						return nil, fmt.Errorf("proxy error")
					},
				},
				config: &KeyFlowConfig{},
			},
			handlerFn: func(testing.TB) http.HandlerFunc {
				return func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					_, _ = fmt.Fprintln(w, `{"status":"ok"}`)
				}
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			t.Cleanup(cancel) // This cancels the refresher goroutine

			privateKeyBytes, err := generatePrivateKey()
			if err != nil {
				t.Errorf("no error is expected, but got %v", err)
			}

			tt.keyFlow.config.ServiceAccountKey = fixtureServiceAccountKey()
			tt.keyFlow.config.PrivateKey = string(privateKeyBytes)
			tt.keyFlow.config.BackgroundTokenRefreshContext = ctx
			tt.keyFlow.authClient = &http.Client{
				Transport: mockTransportFn{
					fn: func(_ *http.Request) (*http.Response, error) {
						res := httptest.NewRecorder()
						res.WriteHeader(http.StatusOK)
						res.Header().Set("Content-Type", "application/json")

						token := &TokenResponseBody{
							AccessToken:  testBearerToken,
							ExpiresIn:    2147483647,
							RefreshToken: testBearerToken,
							TokenType:    "Bearer",
						}

						if err := json.NewEncoder(res.Body).Encode(token); err != nil {
							t.Logf("no error is expected, but got %v", err)
						}

						return res.Result(), nil
					},
				},
			}

			if err := tt.keyFlow.validate(); err != nil {
				t.Errorf("no error is expected, but got %v", err)
			}

			go continuousRefreshToken(tt.keyFlow)

			tokenCtx, tokenCancel := context.WithTimeout(context.Background(), 1*time.Second)

		token:
			for {
				select {
				case <-tokenCtx.Done():
					t.Error(tokenCtx.Err())
				case <-time.After(50 * time.Millisecond):
					tt.keyFlow.tokenMutex.RLock()
					if tt.keyFlow.token != nil {
						tt.keyFlow.tokenMutex.RUnlock()
						tokenCancel()
						break token
					}

					tt.keyFlow.tokenMutex.RUnlock()
				}
			}

			server := httptest.NewServer(tt.handlerFn(t))
			t.Cleanup(server.Close)

			u, err := url.Parse(server.URL)
			if err != nil {
				t.Errorf("no error is expected, but got %v", err)
			}

			req, err := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
			if err != nil {
				t.Errorf("no error is expected, but got %v", err)
			}

			httpClient := &http.Client{
				Transport: tt.keyFlow,
			}

			res, err := httpClient.Do(req)

			if tt.wantErr {
				if err == nil {
					t.Errorf("error is expected, but got %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("no error is expected, but got %v", err)
				}

				if res.StatusCode != tt.want {
					t.Errorf("expected status code %d, but got %d", tt.want, res.StatusCode)
				}

				// Defer discard and close the body
				t.Cleanup(func() {
					if _, err := io.Copy(io.Discard, res.Body); err != nil {
						t.Errorf("no error is expected, but got %v", err)
					}

					if err := res.Body.Close(); err != nil {
						t.Errorf("no error is expected, but got %v", err)
					}
				})
			}
		})
	}
}

type mockTransportFn struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (m mockTransportFn) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.fn(req)
}
