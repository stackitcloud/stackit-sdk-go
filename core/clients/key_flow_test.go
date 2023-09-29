package clients

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const saKeyStrPattern = `{
	"active": true,
	"createdAt": "2023-03-23T18:26:20.335Z",
	"credentials": {
	  "aud": "https://stackit-service-account-prod.apps.01.cf.eu01.stackit.cloud",
	  "iss": "stackit@sa.stackit.cloud",
	  "kid": "%s",
	  "sub": "%s"
	},
	"id": "%s",
	"keyAlgorithm": "RSA_2048",
	"keyOrigin": "USER_PROVIDED",
	"keyType": "USER_MANAGED",
	"publicKey": "...",
	"validUntil": "2024-03-22T18:05:41Z"
}`

var saKey = fmt.Sprintf(saKeyStrPattern, uuid.New().String(), uuid.New().String(), uuid.New().String())

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

func TestKeyFlow_Init(t *testing.T) {
	tests := []struct {
		name              string
		serviceAccountKey string
		genPrivateKey     bool
		invalidPrivateKey bool
		wantErr           bool
	}{
		{
			name:              "ok",
			serviceAccountKey: saKey,
			genPrivateKey:     true,
			wantErr:           false,
		},
		{
			name:              "missing_private_key",
			serviceAccountKey: saKey,
			genPrivateKey:     false,
			wantErr:           true,
		},
		{
			name:              "missing_service_account_key",
			serviceAccountKey: "",
			genPrivateKey:     true,
			wantErr:           true,
		},
		{
			name:              "invalid_private_key",
			serviceAccountKey: saKey,
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
				privateKey, err := generatePrivateKey()
				if err != nil {
					t.Fatalf("Error generating private key: %s", err)
				}
				cfg.PrivateKey = string(privateKey)
			}
			if tt.invalidPrivateKey {
				cfg.PrivateKey = "invalid_key"
			}
			cfg.ServiceAccountKey = tt.serviceAccountKey
			if err := c.Init(context.Background(), cfg); (err != nil) != tt.wantErr {
				t.Errorf("KeyFlow.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			if c.config == nil {
				t.Error("config is nil")
			}
		})
	}
}

func TestKeyFlow_Do(t *testing.T) {
	type fields struct {
		client *http.Client
		config *KeyFlowConfig
	}
	type args struct{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"fail", fields{nil, nil}, args{}, 0, true},
		{"success", fields{&http.Client{}, &KeyFlowConfig{ClientRetry: &RetryConfig{}}}, args{}, http.StatusOK, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &KeyFlow{
				client: tt.fields.client,
				config: tt.fields.config,
			}
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, `{"status":"ok"}`)
			})
			server := httptest.NewServer(handler)
			defer server.Close()
			u, err := url.Parse(server.URL)
			if err != nil {
				t.Error(err)
				return
			}
			req := &http.Request{
				URL: u,
			}
			got, err := c.RoundTrip(req)
			if err == nil {
				// Defer discard and close the body
				defer func() {
					if _, discardErr := io.Copy(io.Discard, got.Body); discardErr != nil && err == nil {
						err = discardErr
					}
					if closeErr := got.Body.Close(); closeErr != nil && err == nil {
						err = closeErr
					}
				}()
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyFlow.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.StatusCode != tt.want {
				t.Errorf("KeyFlow.Do() = %v, want %v", got.StatusCode, tt.want)
			}
		})
	}
}

func TestKeyClone(t *testing.T) {
	c := &KeyFlow{
		client: &http.Client{},
		config: &KeyFlowConfig{},
		key:    &ServiceAccountKeyPrivateResponse{},
		token:  &TokenResponseBody{},
	}

	clone, ok := c.Clone().(*KeyFlow)
	if !ok {
		t.Fatalf("Type assertion failed")
	}

	if !reflect.DeepEqual(c, clone) {
		t.Errorf("Clone() = %v, want %v", clone, c)
	}
}

func TestKeyFlow_validateToken(t *testing.T) {

	// Generate a random private key
	privateKey := make([]byte, 32)
	if _, err := rand.Read(privateKey); err != nil {
		t.Fatal(err)
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		token   string
		want    bool
		wantErr bool
	}{
		{"no token", "", false, false},
		{"bad token - ask to recreate", "bad token", false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &KeyFlow{
				config: &KeyFlowConfig{
					PrivateKey: string(privateKey),
				},
				doer: do,
			}
			got, err := c.validateToken(tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyFlow.validateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("KeyFlow.validateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

type MockDoer struct {
	mock.Mock
}

func (m *MockDoer) Do(client *http.Client, req *http.Request, cfg *RetryConfig) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestGetJwksJSON(t *testing.T) {
	testCases := []struct {
		name           string
		token          string
		mockResponse   *http.Response
		mockError      error
		expectedResult []byte
		expectedError  error
	}{
		{
			name:  "Success",
			token: "test_token",
			mockResponse: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"key": "value"}`))),
			},
			mockError:      nil,
			expectedResult: []byte(`{"key": "value"}`),
			expectedError:  nil,
		},
		{
			name:           "Error",
			token:          "test_token",
			mockResponse:   nil,
			mockError:      fmt.Errorf("some error"),
			expectedResult: nil,
			expectedError:  fmt.Errorf("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if 
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, `{"status":"ok"}`)
			})
			server := httptest.NewServer(handler)
			defer server.Close()
			u, err := url.Parse(server.URL)
			if err != nil {
				t.Error(err)
				return
			}
			req := &http.Request{
				URL: u,
			}

			mockDoer := new(MockDoer)
			mockDoer.On("Do", mock.Anything).Return(tc.mockResponse, tc.mockError)

			c := &KeyFlow{
				config: &KeyFlowConfig{ClientRetry: NewRetryConfig()},
				doer:   client.Do,
			}

			result, err := c.getJwksJSON(tc.token)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expectedResult, result)
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
			mockError:     errors.New("request error"),
			expectedError: errors.New("request error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDoer := new(MockDoer)
			mockDoer.On("Do", mock.Anything).Return(tc.mockResponse, tc.mockError)

			c := &KeyFlow{
				config: &KeyFlowConfig{ClientRetry: NewRetryConfig()},
				doer:   mockDoer.Do,
			}

			res, err := c.requestToken(tc.grant, tc.assertion)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.mockResponse, res)
		})
	}
}
