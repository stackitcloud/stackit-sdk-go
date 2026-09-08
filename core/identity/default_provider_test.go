package identity

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/google/uuid"
)

// clearCredentialEnv removes every credential env var so the tests do not pick up
// ambient credentials from the machine running them.
func clearCredentialEnv(t *testing.T) {
	t.Helper()
	for _, env := range []string{
		EnvServiceAccountToken,
		EnvServiceAccountKey,
		EnvServiceAccountKeyPath,
		EnvPrivateKey,
		EnvPrivateKeyPath,
		EnvServiceAccountEmail,
		EnvFederatedTokenFile,
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
	t.Setenv(EnvCredentialsPath, filepath.Join(t.TempDir(), "no-such-credentials.json"))
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

// TestDefaultProviderPrefersStaticOverServiceAccountKey pins the documented order:
// a pre-issued token wins over the service account key flow, so the key flow's token
// endpoint must never be contacted.
func TestDefaultProviderPrefersStaticOverServiceAccountKey(t *testing.T) {
	clearCredentialEnv(t)

	privateKeyPEM, err := generateRSAPrivateKeyPEM()
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}

	var requests int64
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		atomic.AddInt64(&requests, 1)
		_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "from-key-flow", "expires_in": 3600})
	}))
	defer server.Close()

	saKey := ServiceAccountJson{
		Credentials: &ServiceAccountKeyCredentials{
			Aud: server.URL,
			Iss: "service-account@sa.stackit.cloud",
			Kid: "kid",
			Sub: uuid.New(),
		},
	}
	saKeyJSON, err := json.Marshal(saKey)
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

	// Both providers must be in the chain, with the static token first.
	if len(provider.chain.providers) != 2 {
		t.Fatalf("expected 2 providers in the chain, got %d", len(provider.chain.providers))
	}
	if _, ok := provider.chain.providers[0].(*StaticTokenProvider); !ok {
		t.Fatalf("expected StaticTokenProvider first, got %T", provider.chain.providers[0])
	}

	token, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if token.AccessToken != accessToken {
		t.Fatalf("expected the static token to win, got %s", token.AccessToken)
	}
	if got := atomic.LoadInt64(&requests); got != 0 {
		t.Fatalf("expected the key flow token endpoint not to be contacted, got %d requests", got)
	}
}

// TestDefaultProviderInstanceMetadataAfterConfigured pins that the ambient VM identity is
// part of the chain but never outranks explicitly configured credentials.
func TestDefaultProviderInstanceMetadataAfterConfigured(t *testing.T) {
	clearCredentialEnv(t)
	// Point the federated token file at a path that does not exist so workload identity
	// federation is deterministically skipped, whatever the machine running the test.
	t.Setenv(EnvFederatedTokenFile, filepath.Join(t.TempDir(), "no-such-token"))

	//nolint:gosec // G101 False positive: This is a test JWT token, not a credential
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTl9.test"
	provider, err := NewDefaultProvider(&DefaultProviderConfig{
		Token:               accessToken,
		ServiceAccountEmail: "service-account@sa.stackit.cloud",
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	providers := provider.chain.providers
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
	t.Setenv(EnvServiceAccountToken, accessToken)

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
	credentialsFile := filepath.Join(t.TempDir(), "credentials.json")
	credentials, err := json.Marshal(Credentials{ServiceAccountToken: accessToken})
	if err != nil {
		t.Fatalf("marshal credentials: %v", err)
	}
	if err := os.WriteFile(credentialsFile, credentials, 0o600); err != nil {
		t.Fatalf("write credentials file: %v", err)
	}
	t.Setenv(EnvCredentialsPath, credentialsFile)

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
	if err := os.WriteFile(fakeCLIPath, []byte("#!/bin/sh\nexit 1\n"), 0o700); err != nil {
		t.Fatalf("write fake cli: %v", err)
	}

	provider, err := NewDefaultProvider(&DefaultProviderConfig{
		Token:      testCLIAccessToken,
		CLICommand: fakeCLIPath,
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	providers := provider.chain.providers
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

// TestDefaultProviderCLIDisabled covers the STACKIT_USE_CLI escape hatch.
func TestDefaultProviderCLIDisabled(t *testing.T) {
	clearCredentialEnv(t)
	t.Setenv(EnvUseCLI, "false")

	fakeCLIPath := filepath.Join(t.TempDir(), "stackit")
	if err := os.WriteFile(fakeCLIPath, []byte("#!/bin/sh\nexit 1\n"), 0o700); err != nil {
		t.Fatalf("write fake cli: %v", err)
	}

	provider, err := NewDefaultProvider(&DefaultProviderConfig{
		Token:      testCLIAccessToken,
		CLICommand: fakeCLIPath,
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	for _, p := range provider.chain.providers {
		if _, ok := p.(*CLIProvider); ok {
			t.Fatalf("expected no CLIProvider when %s is false", EnvUseCLI)
		}
	}
}

// TestDefaultProviderCLIEnabledByDefault guards against the switch defaulting the wrong way.
func TestDefaultProviderCLIEnabledByDefault(t *testing.T) {
	clearCredentialEnv(t)

	fakeCLIPath := filepath.Join(t.TempDir(), "stackit")
	if err := os.WriteFile(fakeCLIPath, []byte("#!/bin/sh\nexit 1\n"), 0o700); err != nil {
		t.Fatalf("write fake cli: %v", err)
	}

	for _, value := range []string{"", "true", "not-a-bool"} {
		t.Setenv(EnvUseCLI, value)
		provider, err := NewDefaultProvider(&DefaultProviderConfig{
			Token:      testCLIAccessToken,
			CLICommand: fakeCLIPath,
		})
		if err != nil {
			t.Fatalf("expected no error: %v", err)
		}
		last := provider.chain.providers[len(provider.chain.providers)-1]
		if _, ok := last.(*CLIProvider); !ok {
			t.Fatalf("%s=%q: expected CLIProvider to be present, got %T last", EnvUseCLI, value, last)
		}
	}
}

// TestDefaultProviderCLIDisabledByConfig covers the programmatic half of the switch, and
// that an operator's STACKIT_USE_CLI=false cannot be overridden from code.
func TestDefaultProviderCLIDisabledByConfig(t *testing.T) {
	fakeCLIPath := filepath.Join(t.TempDir(), "stackit")
	if err := os.WriteFile(fakeCLIPath, []byte("#!/bin/sh\nexit 1\n"), 0o700); err != nil {
		t.Fatalf("write fake cli: %v", err)
	}

	tests := []struct {
		name       string
		env        string
		disableCLI bool
		wantCLI    bool
	}{
		{name: "enabled by default", wantCLI: true},
		{name: "disabled in code", disableCLI: true, wantCLI: false},
		{name: "disabled in the environment", env: "false", wantCLI: false},
		{name: "code cannot re-enable what the environment disabled", env: "false", disableCLI: false, wantCLI: false},
		{name: "both disabled", env: "false", disableCLI: true, wantCLI: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			clearCredentialEnv(t)
			if test.env != "" {
				t.Setenv(EnvUseCLI, test.env)
			}

			provider, err := NewDefaultProvider(&DefaultProviderConfig{
				Token:      testCLIAccessToken,
				CLICommand: fakeCLIPath,
				DisableCLI: test.disableCLI,
			})
			if err != nil {
				t.Fatalf("expected no error: %v", err)
			}

			var found bool
			for _, p := range provider.chain.providers {
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
