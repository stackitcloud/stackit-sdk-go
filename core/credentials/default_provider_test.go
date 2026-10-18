package credentials

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// clearCredentialEnv removes every credential env var so the tests do not pick up
// ambient credentials from the machine running them.
func clearCredentialEnv(t *testing.T) {
	t.Helper()
	for _, env := range []string{
		"STACKIT_SERVICE_ACCOUNT_TOKEN",
		"STACKIT_SERVICE_ACCOUNT_KEY",
		"STACKIT_SERVICE_ACCOUNT_KEY_PATH",
		"STACKIT_PRIVATE_KEY",
		"STACKIT_PRIVATE_KEY_PATH",
		"STACKIT_SERVICE_ACCOUNT_EMAIL",
		"STACKIT_FEDERATED_TOKEN_FILE",
		"STACKIT_USE_CLI",
	} {
		t.Setenv(env, "")
	}
	isolateCredentialsFile(t)
	// Empty PATH so CLIProvider is deterministically skipped: whether the chain includes
	// it must not depend on the STACKIT CLI being installed on the machine running tests.
	t.Setenv("PATH", t.TempDir())
}

// isolateCredentialsFile points the credentials file at a path that does not exist, so
// tests never pick up a real ~/.stackit/credentials.json from the machine running them.
func isolateCredentialsFile(t *testing.T) {
	t.Helper()
	t.Setenv("STACKIT_CREDENTIALS_PATH", filepath.Join(t.TempDir(), "no-such-credentials.json"))
}

func TestDefaultProviderNoCredentials(t *testing.T) {
	clearCredentialEnv(t)

	_, err := NewDefaultProvider(&DefaultProviderConfig{})
	if err == nil {
		t.Fatalf("expected error when no credentials are available")
	}
	// The error must explain why each provider was skipped, not just that it failed.
	for _, want := range []string{"StaticTokenProvider", "ServiceAccountKeyProvider", "WorkloadIdentityFederationProvider", "CLIProvider"} {
		if !strings.Contains(err.Error(), want) {
			t.Fatalf("expected error to mention %s, got: %v", want, err)
		}
	}
}

func TestDefaultProviderStaticToken(t *testing.T) {
	clearCredentialEnv(t)

	//nolint:gosec // G101 False positive: This is a test JWT token, not a credential
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTl9.test"
	provider, err := NewDefaultProvider(&DefaultProviderConfig{Token: accessToken})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	token, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if token.AccessToken != accessToken {
		t.Fatalf("expected %s, got %s", accessToken, token.AccessToken)
	}
}

// TestDefaultProviderPrefersServiceAccountKeyOverStatic pins the documented order:
// the service account key flow wins over a pre-issued token.
func TestDefaultProviderPrefersServiceAccountKeyOverStatic(t *testing.T) {
	clearCredentialEnv(t)
	t.Setenv("STACKIT_FEDERATED_TOKEN_FILE", filepath.Join(t.TempDir(), "no-such-token"))

	privateKeyPEM, err := generateRSAPrivateKeyPEM()
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "from-key-flow", "expires_in": 3600})
	}))
	defer server.Close()

	saKeyJSON, err := json.Marshal(serviceAccountJSON{
		Credentials: &serviceAccountKeyCredentials{
			Aud: server.URL,
			Iss: "service-account@sa.stackit.cloud",
			Kid: "kid",
			Sub: uuid.New(),
		},
	})
	if err != nil {
		t.Fatalf("marshal service account key: %v", err)
	}

	//nolint:gosec // G101 False positive: This is a test JWT token, not a credential
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTl9.test"
	provider, err := NewDefaultProvider(&DefaultProviderConfig{
		Token:             accessToken,
		ServiceAccountKey: string(saKeyJSON),
		PrivateKey:        string(privateKeyPEM),
		TokenURL:          server.URL,
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	if len(provider.providers) != 2 {
		t.Fatalf("expected 2 providers in the chain, got %d", len(provider.providers))
	}
	if _, ok := provider.providers[0].(*ServiceAccountKeyProvider); !ok {
		t.Fatalf("expected ServiceAccountKeyProvider first, got %T", provider.providers[0])
	}
	if _, ok := provider.providers[1].(*StaticTokenProvider); !ok {
		t.Fatalf("expected StaticTokenProvider second, got %T", provider.providers[1])
	}

	token, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if token.AccessToken != "from-key-flow" {
		t.Fatalf("expected the key flow token to win, got %s", token.AccessToken)
	}
}

// TestDefaultProviderPrefersWorkloadIdentityFederationOverServiceAccountKey pins that the
// secretless flow outranks the one based on a long-lived key.
func TestDefaultProviderPrefersWorkloadIdentityFederationOverServiceAccountKey(t *testing.T) {
	clearCredentialEnv(t)

	privateKeyPEM, err := generateRSAPrivateKeyPEM()
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}
	saKeyJSON, err := json.Marshal(serviceAccountJSON{
		Credentials: &serviceAccountKeyCredentials{Aud: "aud", Iss: "iss", Kid: "kid", Sub: uuid.New()},
	})
	if err != nil {
		t.Fatalf("marshal service account key: %v", err)
	}

	provider, err := NewDefaultProvider(&DefaultProviderConfig{
		ServiceAccountEmail:    "service-account@sa.stackit.cloud",
		FederatedTokenFunction: func(context.Context) (string, error) { return "federated-token", nil },
		ServiceAccountKey:      string(saKeyJSON),
		PrivateKey:             string(privateKeyPEM),
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	providers := provider.providers
	if _, ok := providers[0].(*WorkloadIdentityFederationProvider); !ok {
		t.Fatalf("expected WorkloadIdentityFederationProvider first, got %T", providers[0])
	}
	if _, ok := providers[1].(*ServiceAccountKeyProvider); !ok {
		t.Fatalf("expected ServiceAccountKeyProvider second, got %T", providers[1])
	}
}

// TestDefaultProviderInstanceMetadataAfterConfigured pins that the ambient VM identity is
// part of the chain but never outranks explicitly configured credentials.
func TestDefaultProviderInstanceMetadataAfterConfigured(t *testing.T) {
	clearCredentialEnv(t)
	// Point the federated token file at a path that does not exist so workload identity
	// federation is deterministically skipped, whatever the machine running the test.
	t.Setenv("STACKIT_FEDERATED_TOKEN_FILE", filepath.Join(t.TempDir(), "no-such-token"))

	//nolint:gosec // G101 False positive: This is a test JWT token, not a credential
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTl9.test"
	provider, err := NewDefaultProvider(&DefaultProviderConfig{
		Token:               accessToken,
		ServiceAccountEmail: "service-account@sa.stackit.cloud",
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	providers := provider.providers
	if len(providers) != 2 {
		t.Fatalf("expected 2 providers in the chain, got %d", len(providers))
	}
	if _, ok := providers[0].(*StaticTokenProvider); !ok {
		t.Fatalf("expected StaticTokenProvider first, got %T", providers[0])
	}
	if _, ok := providers[1].(*InstanceMetadataProvider); !ok {
		t.Fatalf("expected InstanceMetadataProvider after the configured credentials, got %T", providers[1])
	}

	// The metadata endpoint is link-local and unreachable here; the static token must win
	// without it ever being contacted.
	token, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if token.AccessToken != accessToken {
		t.Fatalf("expected the static token to win, got %s", token.AccessToken)
	}
}

func TestDefaultProviderResolvesTokenFromEnvironment(t *testing.T) {
	clearCredentialEnv(t)

	//nolint:gosec // G101 False positive: This is a test JWT token, not a credential
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTl9.test"
	t.Setenv("STACKIT_SERVICE_ACCOUNT_TOKEN", accessToken)

	provider, err := NewDefaultProvider(&DefaultProviderConfig{})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	token, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if token.AccessToken != accessToken {
		t.Fatalf("expected %s, got %s", accessToken, token.AccessToken)
	}
}

// TestDefaultProviderResolvesTokenFromCredentialsFile covers the third step of the
// documented resolution order: config, then environment, then the credentials file. The
// caller does not have to opt into the file by setting CredentialsFilePath.
func TestDefaultProviderResolvesTokenFromCredentialsFile(t *testing.T) {
	clearCredentialEnv(t)

	//nolint:gosec // G101 False positive: This is a test JWT token, not a credential
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTl9.test"
	credentialsPath := filepath.Join(t.TempDir(), "credentials.json")
	credentials, err := json.Marshal(credentialsFile{ServiceAccountToken: accessToken})
	if err != nil {
		t.Fatalf("marshal credentials: %v", err)
	}
	if err := os.WriteFile(credentialsPath, credentials, 0o600); err != nil {
		t.Fatalf("write credentials file: %v", err)
	}
	t.Setenv("STACKIT_CREDENTIALS_PATH", credentialsPath)

	provider, err := NewDefaultProvider(&DefaultProviderConfig{})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	token, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if token.AccessToken != accessToken {
		t.Fatalf("expected the token from the credentials file, got %s", token.AccessToken)
	}
}

// TestDefaultProviderCLIIsLast pins the last step of the documented chain: the CLI session
// is available but never outranks anything else.
func TestDefaultProviderCLIIsLast(t *testing.T) {
	clearCredentialEnv(t)

	// A stand-in for the STACKIT CLI, so the test does not depend on it being installed.
	fakeCLIPath := filepath.Join(t.TempDir(), "stackit")
	if err := os.WriteFile(fakeCLIPath, []byte("#!/bin/sh\nexit 1\n"), 0o700); err != nil { //nolint:gosec // G306: the fake CLI has to be executable
		t.Fatalf("write fake cli: %v", err)
	}

	provider, err := NewDefaultProvider(&DefaultProviderConfig{
		Token:      testCLIAccessToken,
		CLICommand: fakeCLIPath,
		EnableCLI:  true,
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	providers := provider.providers
	if len(providers) != 2 {
		t.Fatalf("expected 2 providers in the chain, got %d", len(providers))
	}
	if _, ok := providers[0].(*StaticTokenProvider); !ok {
		t.Fatalf("expected StaticTokenProvider first, got %T", providers[0])
	}
	if _, ok := providers[len(providers)-1].(*CLIProvider); !ok {
		t.Fatalf("expected CLIProvider last, got %T", providers[len(providers)-1])
	}

	// The fake CLI always fails, so the static token must win without it being invoked.
	token, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if token.AccessToken != testCLIAccessToken {
		t.Fatalf("expected the static token to win, got %s", token.AccessToken)
	}
}

// TestDefaultProviderCLIOptIn covers the two ways to enable the CLI session step, and that
// it is off by default.
func TestDefaultProviderCLIOptIn(t *testing.T) {
	fakeCLIPath := filepath.Join(t.TempDir(), "stackit")
	if err := os.WriteFile(fakeCLIPath, []byte("#!/bin/sh\nexit 1\n"), 0o700); err != nil { //nolint:gosec // G306: the fake CLI has to be executable
		t.Fatalf("write fake cli: %v", err)
	}

	tests := []struct {
		name      string
		env       string
		enableCLI bool
		wantCLI   bool
	}{
		{name: "off by default", wantCLI: false},
		{name: "enabled in code", enableCLI: true, wantCLI: true},
		{name: "enabled in the environment", env: "true", wantCLI: true},
		{name: "environment false leaves it off", env: "false", wantCLI: false},
		{name: "a value that is not a boolean leaves it off", env: "not-a-bool", wantCLI: false},
		{name: "code wins over an environment false", env: "false", enableCLI: true, wantCLI: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			clearCredentialEnv(t)
			t.Setenv("STACKIT_USE_CLI", test.env)

			provider, err := NewDefaultProvider(&DefaultProviderConfig{
				Token:      testCLIAccessToken,
				CLICommand: fakeCLIPath,
				EnableCLI:  test.enableCLI,
			})
			if err != nil {
				t.Fatalf("expected no error: %v", err)
			}

			var found bool
			for _, p := range provider.providers {
				if _, ok := p.(*CLIProvider); ok {
					found = true
				}
			}
			if found != test.wantCLI {
				t.Fatalf("expected CLIProvider present=%v, got %v", test.wantCLI, found)
			}
		})
	}
}

func TestDefaultProviderServiceAccountKeyFromPaths(t *testing.T) {
	clearCredentialEnv(t)
	privateKeyPEM, _ := newRSAKey(t)
	tokenURL, assertion := keyResolutionServer(t)

	provider, err := NewDefaultProvider(&DefaultProviderConfig{
		ServiceAccountKeyPath: writeTestFile(t, serviceAccountKeyJSONFor(t, "from-path")),
		PrivateKeyPath:        writeTestFile(t, privateKeyPEM),
		TokenURL:              tokenURL,
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if _, ok := provider.providers[0].(*ServiceAccountKeyProvider); !ok {
		t.Fatalf("expected ServiceAccountKeyProvider first, got %T", provider.providers[0])
	}
	if _, err := provider.Token(context.Background(), TokenRequestOptions{}); err != nil {
		t.Fatalf("token: %v", err)
	}
	claims := jwt.MapClaims{}
	if _, _, err := jwt.NewParser().ParseUnverified(assertion(), claims); err != nil {
		t.Fatalf("parse assertion: %v", err)
	}
	if claims["iss"] != "from-path" {
		t.Fatalf("expected the key read from the path, got iss %v", claims["iss"])
	}
}

// TestDefaultProviderReadsCredentialsFileOnce pins that the providers of the default chain
// share one read of the credentials file. Every read of a file that others can read logs a
// warning, so the number of warnings is the number of reads.
func TestDefaultProviderReadsCredentialsFileOnce(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("permission bits are synthesized on Windows")
	}
	clearCredentialEnv(t)
	logs := captureLogs(t)

	_, err := NewDefaultProvider(&DefaultProviderConfig{
		CredentialsFilePath: writeCredentialsFile(t, 0o644),
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if got := strings.Count(logs.String(), "readable by other users"); got != 1 {
		t.Fatalf("expected the credentials file to be read once, it was read %d times", got)
	}
}

// TestDefaultProviderDoesNotReadCredentialsFileWhenNotNeeded pins that the file is only read
// when some value is missing from the config and the environment.
func TestDefaultProviderDoesNotReadCredentialsFileWhenNotNeeded(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("permission bits are synthesized on Windows")
	}
	clearCredentialEnv(t)
	logs := captureLogs(t)

	//nolint:gosec // G101 False positive: This is a test JWT token, not a credential
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTl9.test"
	privateKeyPEM, _ := newRSAKey(t)
	_, err := NewDefaultProvider(&DefaultProviderConfig{
		Token:               accessToken,
		ServiceAccountKey:   serviceAccountKeyJSONFor(t, "iss"),
		PrivateKey:          privateKeyPEM,
		ServiceAccountEmail: "sa@sa.stackit.cloud",
		FederatedTokenFunction: func(context.Context) (string, error) {
			return "federated", nil
		},
		CredentialsFilePath: writeCredentialsFile(t, 0o644),
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if logs.Len() != 0 {
		t.Fatalf("expected the credentials file not to be read, got: %q", logs.String())
	}
}
