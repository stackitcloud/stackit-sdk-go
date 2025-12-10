package clients

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestWorkloadIdentityFlowInit(t *testing.T) {
	tests := []struct {
		name                 string
		clientID             string
		clientIDAsEnv        bool
		customTokenUrl       string
		customTokenUrlEnv    bool
		tokenExpiration      string
		validAssertion       bool
		tokenFilePathAsEnv   bool
		missingTokenFilePath bool
		wantErr              bool
	}{
		{
			name:           "ok setting all",
			clientID:       "test@stackit.cloud",
			validAssertion: true,
			wantErr:        false,
		},
		{
			name:               "ok using defaults from envs",
			clientID:           "test@stackit.cloud",
			clientIDAsEnv:      true,
			tokenFilePathAsEnv: true,
			customTokenUrlEnv:  true,
			validAssertion:     true,
			wantErr:            false,
		},
		{
			name:               "ok using defaults from envs",
			clientID:           "test@stackit.cloud",
			clientIDAsEnv:      true,
			tokenFilePathAsEnv: true,
			customTokenUrlEnv:  true,
			validAssertion:     true,
			wantErr:            false,
		},
		{
			name:           "missing client id",
			validAssertion: true,
			wantErr:        true,
		},
		{
			name:                 "missing assertion",
			clientID:             "test@stackit.cloud",
			missingTokenFilePath: true,
			wantErr:              true,
		},
		{
			name:     "invalid assertion",
			clientID: "test@stackit.cloud",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flow := &WorkloadIdentityFederationFlow{}
			flowConfig := &WorkloadIdentityFederationFlowConfig{}
			if tt.customTokenUrl != "" {
				if tt.customTokenUrlEnv {
					t.Setenv("STACKIT_IDP_ENDPOINT", tt.customTokenUrl)
				} else {
					flowConfig.TokenUrl = tt.customTokenUrl
				}
			}

			if tt.clientID != "" {
				if tt.clientIDAsEnv {
					t.Setenv("STACKIT_SERVICE_ACCOUNT_EMAIL", tt.clientID)
				} else {
					flowConfig.ClientID = tt.clientID
				}
			}
			if tt.tokenExpiration != "" {
				flowConfig.TokenExpiration = tt.tokenExpiration
			}

			if !tt.missingTokenFilePath {
				file, err := os.CreateTemp("", "*.token")
				if err != nil {
					log.Fatal(err)
				}
				defer os.Remove(file.Name())
				if tt.validAssertion {
					token, err := signTokenWithSubject("subject", time.Minute)
					if err != nil {
						t.Fatalf("failed to create token: %v", err)
					}
					os.WriteFile(file.Name(), []byte(token), os.ModeAppend)
				}
				if tt.tokenFilePathAsEnv {
					t.Setenv("STACKIT_FEDERATED_TOKEN_FILE", file.Name())
				} else {
					flowConfig.FederatedTokenFilePath = file.Name()
				}
			}

			if err := flow.Init(flowConfig); (err != nil) != tt.wantErr {
				t.Errorf("KeyFlow.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			if flow.config == nil {
				t.Error("config is nil")
			}

			if flow.config.ClientID != tt.clientID {
				t.Errorf("clientID mismatch, want %s, got %s", tt.clientID, flow.config.ClientID)
			}

			if tt.customTokenUrl != "" && flow.config.TokenUrl != tt.customTokenUrl {
				t.Errorf("tokenUrl mismatch, want %s, got %s", tt.customTokenUrl, flow.config.TokenUrl)
			}

			if tt.customTokenUrl == "" && flow.config.TokenUrl != "https://accounts.stackit.cloud/oauth/v2/token" {
				t.Errorf("tokenUrl mismatch, want %s, got %s", "https://accounts.stackit.cloud/oauth/v2/token", flow.config.TokenUrl)
			}

			if tt.missingTokenFilePath && flow.config.FederatedTokenFilePath != "/var/run/secrets/stackit.cloud/serviceaccount/token" {
				t.Errorf("clientID mismatch, want %s, got %s", "/var/run/secrets/stackit.cloud/serviceaccount/token", flow.config.FederatedTokenFilePath)
			}

			if !tt.missingTokenFilePath && flow.config.FederatedTokenFilePath == "/var/run/secrets/stackit.cloud/serviceaccount/token" {
				t.Errorf("clientID mismatch, want different from %s", flow.config.FederatedTokenFilePath)
			}

			if tt.tokenExpiration != "" && flow.config.TokenExpiration != tt.tokenExpiration {
				t.Errorf("tokenExpiration mismatch, want %s, got %s", tt.tokenExpiration, flow.config.TokenExpiration)
			}
		})
	}
}

func signTokenWithSubject(sub string, expiration time.Duration) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
		Subject:   sub,
	}).SignedString([]byte("test"))
}

func TestWorkloadIdentityFlowRoundTrip(t *testing.T) {
	validSub := "valid-sub"
	serviceAccountSub := "sa-sub"
	tests := []struct {
		name           string
		clientID       string
		validAssertion bool
		wantErr        bool
	}{
		{
			name:           "ok setting all",
			clientID:       "test@stackit.cloud",
			validAssertion: true,
			wantErr:        false,
		},
		{
			name:           "invalid assertion",
			clientID:       "test@stackit.cloud",
			validAssertion: false,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r.ParseForm()
				assertionType := r.PostForm.Get("client_assertion_type")
				if assertionType != "urn:schwarz:params:oauth:client-assertion-type:workload-jwt" {
					t.Fatalf("invalid assertion type: %s", assertionType)
				}
				grantType := r.PostForm.Get("grant_type")
				if grantType != "client_credentials" {
					t.Fatalf("invalid grant type: %s", assertionType)
				}
				context, _, err := jwt.NewParser().ParseUnverified(r.PostForm.Get("client_assertion"), jwt.MapClaims{})
				if err != nil {
					t.Fatalf("failed to validate token: %v", err)
				}
				sub, err := context.Claims.GetSubject()
				if err != nil {
					t.Fatalf("failed to validate token sub: %v", err)
				}
				if sub != validSub {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				token, err := signTokenWithSubject(serviceAccountSub, time.Minute)
				if err != nil {
					t.Fatalf("failed to create token: %v", err)
				}

				tokenResponse := &TokenResponseBody{
					AccessToken: token,
					ExpiresIn:   60,
					TokenType:   "Bearer",
				}

				payload, err := json.Marshal(tokenResponse)
				if err != nil {
					t.Fatalf("failed to create token payload: %v", err)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(payload)
			}))
			t.Cleanup(authServer.Close)

			protectedResource := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				context, _, err := jwt.NewParser().ParseUnverified(strings.Fields(r.Header.Get("Authorization"))[1], jwt.MapClaims{})
				if err != nil {
					t.Fatalf("failed to validate token: %v", err)
				}

				sub, err := context.Claims.GetSubject()
				if err != nil {
					t.Fatalf("failed to validate token sub: %v", err)
				}
				if sub != serviceAccountSub {
					t.Fatalf("invalid token on protected resource: %v", err)
				}

				w.WriteHeader(http.StatusOK)
			}))
			t.Cleanup(protectedResource.Close)

			flow := &WorkloadIdentityFederationFlow{}
			flowConfig := &WorkloadIdentityFederationFlowConfig{}
			flowConfig.TokenUrl = authServer.URL

			flowConfig.ClientID = tt.clientID
			file, err := os.CreateTemp("", "*.token")
			if err != nil {
				log.Fatal(err)
			}
			defer os.Remove(file.Name())
			flowConfig.FederatedTokenFilePath = file.Name()

			subject := "wrong"
			if tt.validAssertion {
				subject = validSub
			}
			token, err := signTokenWithSubject(subject, time.Minute)
			if err != nil {
				t.Fatalf("failed to create token: %v", err)
			}
			os.WriteFile(file.Name(), []byte(token), os.ModeAppend)

			if err := flow.Init(flowConfig); err != nil {
				t.Errorf("KeyFlow.Init() error = %v", err)
			}
			if flow.config == nil {
				t.Error("config is nil")
			}

			client := http.Client{
				Transport: flow,
			}
			resp, err := client.Get(protectedResource.URL)
			if (err != nil || resp.StatusCode != http.StatusOK) && !tt.wantErr {
				t.Fatalf("failed request to protected resource: %v", err)
			}
		})
	}
}

// func TestRequestToken(t *testing.T) {
// 	testCases := []struct {
// 		name          string
// 		grant         string
// 		assertion     string
// 		mockResponse  *http.Response
// 		mockError     error
// 		expectedError error
// 	}{
// 		{
// 			name:      "Success",
// 			grant:     "test_grant",
// 			assertion: "test_assertion",
// 			mockResponse: &http.Response{
// 				StatusCode: 200,
// 				Body:       io.NopCloser(strings.NewReader(`{"access_token": "test_token"}`)),
// 			},
// 			mockError:     nil,
// 			expectedError: nil,
// 		},
// 		{
// 			name:          "Error",
// 			grant:         "test_grant",
// 			assertion:     "test_assertion",
// 			mockResponse:  nil,
// 			mockError:     fmt.Errorf("request error"),
// 			expectedError: fmt.Errorf("request error"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			keyFlow := &KeyFlow{}
// 			privateKeyBytes, err := generatePrivateKey()
// 			if err != nil {
// 				t.Fatalf("Error generating private key: %s", err)
// 			}
// 			keyFlowConfig := &KeyFlowConfig{
// 				AuthHTTPClient: &http.Client{
// 					Transport: mockTransportFn{func(_ *http.Request) (*http.Response, error) {
// 						return tt.mockResponse, tt.mockError
// 					}},
// 				},
// 				ServiceAccountKey: fixtureServiceAccountKey(),
// 				PrivateKey:        string(privateKeyBytes),
// 				HTTPTransport:     http.DefaultTransport,
// 			}
// 			err = keyFlow.Init(keyFlowConfig)
// 			if err != nil {
// 				t.Fatalf("failed to initialize key flow: %v", err)
// 			}

// 			res, err := keyFlow.requestToken(tt.grant, tt.assertion)
// 			defer func() {
// 				if res != nil {
// 					tempErr := res.Body.Close()
// 					if tempErr != nil {
// 						t.Errorf("closing request token response: %s", tempErr.Error())
// 					}
// 				}
// 			}()
// 			if tt.expectedError != nil {
// 				if err == nil {
// 					t.Errorf("Expected error '%v' but no error was returned", tt.expectedError)
// 				} else if errors.Is(err, tt.expectedError) {
// 					t.Errorf("Error is not correct. Expected %v, got %v", tt.expectedError, err)
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("Expected no error but error was returned: %v", err)
// 				}
// 				if !cmp.Equal(tt.mockResponse, res, cmp.AllowUnexported(strings.Reader{})) {
// 					t.Errorf("The returned result is wrong. Expected %v, got %v", tt.mockResponse, res)
// 				}
// 			}
// 		})
// 	}
// }

// func TestKeyFlow_Do(t *testing.T) {
// 	t.Parallel()

// 	tests := []struct {
// 		name      string
// 		handlerFn func(tb testing.TB) http.HandlerFunc
// 		want      int
// 		wantErr   bool
// 	}{
// 		{
// 			name: "success",
// 			handlerFn: func(tb testing.TB) http.HandlerFunc {
// 				tb.Helper()

// 				return func(w http.ResponseWriter, r *http.Request) {
// 					if r.Header.Get("Authorization") != "Bearer "+testBearerToken {
// 						tb.Errorf("expected Authorization header to be 'Bearer %s', but got %s", testBearerToken, r.Header.Get("Authorization"))
// 					}

// 					w.Header().Set("Content-Type", "application/json")
// 					w.WriteHeader(http.StatusOK)
// 					_, _ = fmt.Fprintln(w, `{"status":"ok"}`)
// 				}
// 			},
// 			want:    http.StatusOK,
// 			wantErr: false,
// 		},
// 		{
// 			name: "success with code 500",
// 			handlerFn: func(_ testing.TB) http.HandlerFunc {
// 				return func(w http.ResponseWriter, _ *http.Request) {
// 					w.Header().Set("Content-Type", "text/html")
// 					w.WriteHeader(http.StatusInternalServerError)
// 					_, _ = fmt.Fprintln(w, `<html>Internal Server Error</html>`)
// 				}
// 			},
// 			want:    http.StatusInternalServerError,
// 			wantErr: false,
// 		},
// 		{
// 			name: "success with custom transport",
// 			handlerFn: func(tb testing.TB) http.HandlerFunc {
// 				tb.Helper()

// 				return func(w http.ResponseWriter, r *http.Request) {
// 					if r.Header.Get("User-Agent") != "custom_transport" {
// 						tb.Errorf("expected User-Agent header to be 'custom_transport', but got %s", r.Header.Get("User-Agent"))
// 					}

// 					w.Header().Set("Content-Type", "application/json")
// 					w.WriteHeader(http.StatusOK)
// 					_, _ = fmt.Fprintln(w, `{"status":"ok"}`)
// 				}
// 			},
// 			want:    http.StatusOK,
// 			wantErr: false,
// 		},
// 		{
// 			name: "fail with custom proxy",
// 			handlerFn: func(testing.TB) http.HandlerFunc {
// 				return func(w http.ResponseWriter, _ *http.Request) {
// 					w.Header().Set("Content-Type", "application/json")
// 					w.WriteHeader(http.StatusOK)
// 					_, _ = fmt.Fprintln(w, `{"status":"ok"}`)
// 				}
// 			},
// 			want:    0,
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctx := context.Background()
// 			ctx, cancel := context.WithCancel(ctx)
// 			t.Cleanup(cancel) // This cancels the refresher goroutine

// 			privateKeyBytes, err := generatePrivateKey()
// 			if err != nil {
// 				t.Errorf("no error is expected, but got %v", err)
// 			}

// 			keyFlow := &KeyFlow{}
// 			keyFlowConfig := &KeyFlowConfig{
// 				ServiceAccountKey:             fixtureServiceAccountKey(),
// 				PrivateKey:                    string(privateKeyBytes),
// 				BackgroundTokenRefreshContext: ctx,
// 				HTTPTransport: func() http.RoundTripper {
// 					switch tt.name {
// 					case "success with custom transport":
// 						return mockTransportFn{
// 							fn: func(req *http.Request) (*http.Response, error) {
// 								req.Header.Set("User-Agent", "custom_transport")
// 								return http.DefaultTransport.RoundTrip(req)
// 							},
// 						}
// 					case "fail with custom proxy":
// 						return &http.Transport{
// 							Proxy: func(_ *http.Request) (*url.URL, error) {
// 								return nil, fmt.Errorf("proxy error")
// 							},
// 						}
// 					default:
// 						return http.DefaultTransport
// 					}
// 				}(),
// 				AuthHTTPClient: &http.Client{
// 					Transport: mockTransportFn{
// 						fn: func(_ *http.Request) (*http.Response, error) {
// 							res := httptest.NewRecorder()
// 							res.WriteHeader(http.StatusOK)
// 							res.Header().Set("Content-Type", "application/json")

// 							token := &TokenResponseBody{
// 								AccessToken: testBearerToken,
// 								ExpiresIn:   2147483647,
// 								TokenType:   "Bearer",
// 							}

// 							if err := json.NewEncoder(res.Body).Encode(token); err != nil {
// 								t.Logf("no error is expected, but got %v", err)
// 							}

// 							return res.Result(), nil
// 						},
// 					},
// 				},
// 			}
// 			err = keyFlow.Init(keyFlowConfig)
// 			if err != nil {
// 				t.Fatalf("failed to initialize key flow: %v", err)
// 			}

// 			go continuousRefreshToken(keyFlow)

// 			tokenCtx, tokenCancel := context.WithTimeout(context.Background(), 1*time.Second)

// 		token:
// 			for {
// 				select {
// 				case <-tokenCtx.Done():
// 					t.Error(tokenCtx.Err())
// 				case <-time.After(50 * time.Millisecond):
// 					keyFlow.tokenMutex.RLock()
// 					if keyFlow.token != nil {
// 						keyFlow.tokenMutex.RUnlock()
// 						tokenCancel()
// 						break token
// 					}

// 					keyFlow.tokenMutex.RUnlock()
// 				}
// 			}

// 			server := httptest.NewServer(tt.handlerFn(t))
// 			t.Cleanup(server.Close)

// 			u, err := url.Parse(server.URL)
// 			if err != nil {
// 				t.Errorf("no error is expected, but got %v", err)
// 			}

// 			req, err := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
// 			if err != nil {
// 				t.Errorf("no error is expected, but got %v", err)
// 			}

// 			httpClient := &http.Client{
// 				Transport: keyFlow,
// 			}

// 			res, err := httpClient.Do(req)

// 			if tt.wantErr {
// 				if err == nil {
// 					t.Errorf("error is expected, but got %v", err)
// 				}
// 			} else {
// 				if err != nil {
// 					t.Errorf("no error is expected, but got %v", err)
// 				}

// 				if res.StatusCode != tt.want {
// 					t.Errorf("expected status code %d, but got %d", tt.want, res.StatusCode)
// 				}

// 				// Defer discard and close the body
// 				t.Cleanup(func() {
// 					if _, err := io.Copy(io.Discard, res.Body); err != nil {
// 						t.Errorf("no error is expected, but got %v", err)
// 					}

// 					if err := res.Body.Close(); err != nil {
// 						t.Errorf("no error is expected, but got %v", err)
// 					}
// 				})
// 			}
// 		})
// 	}
// }

// type mockTransportFn struct {
// 	fn func(req *http.Request) (*http.Response, error)
// }

// func (m mockTransportFn) RoundTrip(req *http.Request) (*http.Response, error) {
// 	return m.fn(req)
// }
