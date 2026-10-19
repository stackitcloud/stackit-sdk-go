package credentials

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// keyResolutionServer answers token requests and remembers the assertion it received.
func keyResolutionServer(t *testing.T) (url string, assertion func() string) {
	t.Helper()
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Subject:   "subject",
	}).SignedString([]byte("test"))
	if err != nil {
		t.Fatalf("generate access token: %v", err)
	}
	var received string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Errorf("parse form: %v", err)
		}
		received = r.PostForm.Get("assertion")
		_ = json.NewEncoder(w).Encode(map[string]any{"access_token": accessToken, "expires_in": 3600})
	}))
	t.Cleanup(server.Close)
	return server.URL, func() string { return received }
}

func serviceAccountKeyJSONFor(t *testing.T, issuer string) string {
	t.Helper()
	content, err := json.Marshal(serviceAccountJSON{
		Credentials: &serviceAccountKeyCredentials{Aud: "aud", Iss: issuer, Kid: "kid", Sub: uuid.New()},
	})
	if err != nil {
		t.Fatalf("marshal service account key: %v", err)
	}
	return string(content)
}

func writeTestFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "file")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write file: %v", err)
	}
	return path
}

func newRSAKey(t *testing.T) (pem string, key *rsa.PrivateKey) {
	t.Helper()
	raw, err := generateRSAPrivateKeyPEM()
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}
	key, err = jwt.ParseRSAPrivateKeyFromPEM(raw)
	if err != nil {
		t.Fatalf("parse private key: %v", err)
	}
	return string(raw), key
}

func requestToken(t *testing.T, cfg *ServiceAccountKeyProviderConfig, tokenURL string) {
	t.Helper()
	cfg.TokenURL = tokenURL
	provider, err := NewServiceAccountKeyProvider(cfg)
	if err != nil {
		t.Fatalf("new provider: %v", err)
	}
	if _, err := provider.Token(context.Background(), TokenRequestOptions{}); err != nil {
		t.Fatalf("token: %v", err)
	}
}

// TestServiceAccountKeyResolutionOrder pins where the service account key is taken from:
// the config, then the path in the config, then the environment variable, then the path
// in the environment, so that anything explicit in code wins over the environment.
func TestServiceAccountKeyResolutionOrder(t *testing.T) {
	privateKeyPEM, _ := newRSAKey(t)
	tokenURL, assertion := keyResolutionServer(t)

	tests := []struct {
		name                                     string
		config, configPath, environment, envPath bool
		wantIssuer                               string
	}{
		{name: "config content wins over everything", config: true, configPath: true, environment: true, envPath: true, wantIssuer: "config"},
		{name: "config path wins over the environment", configPath: true, environment: true, envPath: true, wantIssuer: "config-path"},
		{name: "environment content wins over the environment path", environment: true, envPath: true, wantIssuer: "environment"},
		{name: "environment path alone", envPath: true, wantIssuer: "environment-path"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			clearCredentialEnv(t)
			cfg := &ServiceAccountKeyProviderConfig{PrivateKey: privateKeyPEM}
			if test.config {
				cfg.ServiceAccountKey = serviceAccountKeyJSONFor(t, "config")
			}
			if test.configPath {
				cfg.ServiceAccountKeyPath = writeTestFile(t, serviceAccountKeyJSONFor(t, "config-path"))
			}
			if test.environment {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY", serviceAccountKeyJSONFor(t, "environment"))
			}
			if test.envPath {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY_PATH", writeTestFile(t, serviceAccountKeyJSONFor(t, "environment-path")))
			}

			requestToken(t, cfg, tokenURL)

			claims := jwt.MapClaims{}
			if _, _, err := jwt.NewParser().ParseUnverified(assertion(), claims); err != nil {
				t.Fatalf("parse assertion: %v", err)
			}
			if claims["iss"] != test.wantIssuer {
				t.Fatalf("expected the key from %q, got %v", test.wantIssuer, claims["iss"])
			}
		})
	}
}

// TestServiceAccountKeyPrivateKeyResolutionOrder does the same for the private key, which
// has one more source: the one embedded in the service account key.
func TestServiceAccountKeyPrivateKeyResolutionOrder(t *testing.T) {
	tokenURL, assertion := keyResolutionServer(t)

	// Case i creates a distinct key for source i and for every source below it, so the key
	// that must win is the one of source i.
	order := []string{"config", "configPath", "environment", "envPath", "embedded", "file"}
	for i, want := range order {
		t.Run("wins: "+want, func(t *testing.T) {
			clearCredentialEnv(t)
			pems := map[string]string{}
			rsaKeys := map[string]*rsa.PrivateKey{}
			for _, source := range order[i:] { // only this source and the ones below it exist
				pems[source], rsaKeys[source] = newRSAKey(t)
			}

			cfg := &ServiceAccountKeyProviderConfig{}
			keyJSON := serviceAccountKeyJSONFor(t, "iss")
			if pem, ok := pems["embedded"]; ok {
				var key serviceAccountJSON
				if err := json.Unmarshal([]byte(keyJSON), &key); err != nil {
					t.Fatal(err)
				}
				key.Credentials.PrivateKey = &pem
				content, _ := json.Marshal(key)
				keyJSON = string(content)
			}
			cfg.ServiceAccountKey = keyJSON
			if pem, ok := pems["config"]; ok {
				cfg.PrivateKey = pem
			}
			if pem, ok := pems["configPath"]; ok {
				cfg.PrivateKeyPath = writeTestFile(t, pem)
			}
			if pem, ok := pems["environment"]; ok {
				t.Setenv("STACKIT_PRIVATE_KEY", pem)
			}
			if pem, ok := pems["envPath"]; ok {
				t.Setenv("STACKIT_PRIVATE_KEY_PATH", writeTestFile(t, pem))
			}
			if pem, ok := pems["file"]; ok {
				credentialsPath := writeTestFile(t, `{"STACKIT_PRIVATE_KEY": `+strings.TrimSpace(mustJSONString(t, pem))+`}`)
				cfg.CredentialsFilePath = credentialsPath
			}

			requestToken(t, cfg, tokenURL)

			_, err := jwt.Parse(assertion(), func(*jwt.Token) (any, error) {
				return &rsaKeys[want].PublicKey, nil
			}, jwt.WithValidMethods([]string{"RS512"}))
			if err != nil {
				t.Fatalf("the assertion was not signed with the private key from %q: %v", want, err)
			}
		})
	}
}

func mustJSONString(t *testing.T, s string) string {
	t.Helper()
	content, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	return string(content)
}

func TestServiceAccountKeyUnreadablePaths(t *testing.T) {
	privateKeyPEM, _ := newRSAKey(t)
	missing := filepath.Join(t.TempDir(), "missing")

	t.Run("service account key path", func(t *testing.T) {
		clearCredentialEnv(t)
		_, err := NewServiceAccountKeyProvider(&ServiceAccountKeyProviderConfig{ServiceAccountKeyPath: missing, PrivateKey: privateKeyPEM})
		if err == nil || !strings.Contains(err.Error(), "read service account key from path") {
			t.Fatalf("expected an error reading the service account key path, got %v", err)
		}
	})
	t.Run("private key path", func(t *testing.T) {
		clearCredentialEnv(t)
		_, err := NewServiceAccountKeyProvider(&ServiceAccountKeyProviderConfig{ServiceAccountKey: serviceAccountKeyJSONFor(t, "iss"), PrivateKeyPath: missing})
		if err == nil || !strings.Contains(err.Error(), "read private key from path") {
			t.Fatalf("expected an error reading the private key path, got %v", err)
		}
	})
	t.Run("content in the config makes the path irrelevant", func(t *testing.T) {
		clearCredentialEnv(t)
		_, err := NewServiceAccountKeyProvider(&ServiceAccountKeyProviderConfig{
			ServiceAccountKey: serviceAccountKeyJSONFor(t, "iss"), ServiceAccountKeyPath: missing,
			PrivateKey: privateKeyPEM, PrivateKeyPath: missing,
		})
		if err != nil {
			t.Fatalf("expected the unused paths to be ignored, got %v", err)
		}
	})
}
