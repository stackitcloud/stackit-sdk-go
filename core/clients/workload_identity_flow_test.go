package clients

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oidcadapters"

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
					t.Setenv("STACKIT_IDP_TOKEN_ENDPOINT", tt.customTokenUrl)
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
				defer func() {
					err := os.Remove(file.Name())
					if err != nil {
						t.Fatalf("Removing temporary file: %s", err)
					}
				}()
				if tt.validAssertion {
					token, err := signTokenWithSubject("subject", time.Minute)
					if err != nil {
						t.Fatalf("failed to create token: %v", err)
					}
					err = os.WriteFile(file.Name(), []byte(token), os.ModeAppend)
					if err != nil {
						t.Fatalf("writing temporary file: %s", err)
					}
				}
				if tt.tokenFilePathAsEnv {
					t.Setenv("STACKIT_FEDERATED_TOKEN_FILE", file.Name())
				} else {
					flowConfig.FederatedTokenFunction = oidcadapters.ReadJWTFromFileSystem(file.Name())
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
		injectToken    bool
		wantErr        bool
	}{
		{
			name:           "ok setting all",
			clientID:       "test@stackit.cloud",
			validAssertion: true,
			wantErr:        false,
		},
		{
			name:           "injected token ok",
			clientID:       "test@stackit.cloud",
			validAssertion: true,
			injectToken:    true,
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
				err := r.ParseForm()
				if err != nil {
					t.Fatalf("failed to parse form: %v", err)
				}
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
				_, err = w.Write(payload)
				if err != nil {
					t.Fatalf("writing response: %s", err)
				}
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

			subject := "wrong"
			if tt.validAssertion {
				subject = validSub
			}
			token, err := signTokenWithSubject(subject, time.Minute)
			if err != nil {
				t.Fatalf("failed to create token: %v", err)
			}

			if tt.injectToken {
				flowConfig.FederatedTokenFunction = func(context.Context) (string, error) {
					return token, nil
				}
			} else {
				file, err := os.CreateTemp("", "*.token")
				if err != nil {
					log.Fatal(err)
				}
				defer func() {
					err := os.Remove(file.Name())
					if err != nil {
						t.Fatalf("Removing temporary file: %s", err)
					}
				}()
				flowConfig.FederatedTokenFunction = oidcadapters.ReadJWTFromFileSystem(file.Name())
				err = os.WriteFile(file.Name(), []byte(token), os.ModeAppend)
				if err != nil {
					t.Fatalf("writing temporary file: %s", err)
				}
			}

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
			if resp != nil && resp.Body != nil {
				if err := resp.Body.Close(); err != nil {
					t.Errorf("resp.Body.Close() error = %v", err)
				}
			}
		})
	}
}
