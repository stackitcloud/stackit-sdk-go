package clients

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
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
			mockDo := func(client *http.Client, req *http.Request) (resp *http.Response, err error) {
				return tt.mockResponse, tt.mockError
			}

			c := &KeyFlow{
				config: &KeyFlowConfig{},
				doer:   mockDo,
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
				} else if tt.expectedError.Error() != err.Error() {
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
