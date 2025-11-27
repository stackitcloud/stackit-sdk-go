package config

import (
	"errors"
	"net/http"
	"testing"
)

// mockCLIAuthProvider is a mock implementation for testing
type mockCLIAuthProvider struct {
	isAuthenticated bool
	authFlowError   error
	roundTripper    http.RoundTripper
}

func (m *mockCLIAuthProvider) IsAuthenticated() bool {
	return m.isAuthenticated
}

func (m *mockCLIAuthProvider) GetAuthFlow() (http.RoundTripper, error) {
	if m.authFlowError != nil {
		return nil, m.authFlowError
	}
	return m.roundTripper, nil
}

// mockRoundTripper is a simple mock RoundTripper for testing
type mockRoundTripper struct{}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 200}, nil
}

func TestWithCLIProviderAuth_Success(t *testing.T) {
	mockRT := &mockRoundTripper{}
	provider := &mockCLIAuthProvider{
		isAuthenticated: true,
		roundTripper:    mockRT,
	}

	cfg := &Configuration{}
	opt := WithCLIProviderAuth(provider)
	err := opt(cfg)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if cfg.CustomAuth == nil {
		t.Error("Expected CustomAuth to be set")
	}

	if cfg.CustomAuth != mockRT {
		t.Error("Expected CustomAuth to be the mock RoundTripper")
	}
}

func TestWithCLIProviderAuth_NilProvider(t *testing.T) {
	cfg := &Configuration{}
	opt := WithCLIProviderAuth(nil)
	err := opt(cfg)

	if err == nil {
		t.Error("Expected error for nil provider")
	}

	var authErr *AuthenticationError
	if !errors.As(err, &authErr) {
		t.Errorf("Expected AuthenticationError, got: %T", err)
	}

	if authErr.msg != "CLI auth provider cannot be nil" {
		t.Errorf("Unexpected error message: %s", authErr.msg)
	}
}

func TestWithCLIProviderAuth_NotAuthenticated(t *testing.T) {
	provider := &mockCLIAuthProvider{
		isAuthenticated: false,
	}

	cfg := &Configuration{}
	opt := WithCLIProviderAuth(provider)
	err := opt(cfg)

	if err == nil {
		t.Error("Expected error when not authenticated")
	}

	var authErr *AuthenticationError
	if !errors.As(err, &authErr) {
		t.Errorf("Expected AuthenticationError, got: %T", err)
	}

	expectedMsg := "not authenticated with CLI provider credentials: please run authentication command (e.g., 'stackit auth provider login')"
	if authErr.msg != expectedMsg {
		t.Errorf("Expected message '%s', got: %s", expectedMsg, authErr.msg)
	}
}

func TestWithCLIProviderAuth_GetAuthFlowError(t *testing.T) {
	testErr := errors.New("failed to get auth flow")
	provider := &mockCLIAuthProvider{
		isAuthenticated: true,
		authFlowError:   testErr,
	}

	cfg := &Configuration{}
	opt := WithCLIProviderAuth(provider)
	err := opt(cfg)

	if err == nil {
		t.Error("Expected error when GetAuthFlow fails")
	}

	var authErr *AuthenticationError
	if !errors.As(err, &authErr) {
		t.Errorf("Expected AuthenticationError, got: %T", err)
	}

	if authErr.msg != "failed to initialize CLI provider authentication" {
		t.Errorf("Unexpected error message: %s", authErr.msg)
	}

	if !errors.Is(err, testErr) {
		t.Error("Expected wrapped error to be accessible")
	}
}

func TestAuthenticationError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *AuthenticationError
		expected string
	}{
		{
			name:     "simple error",
			err:      &AuthenticationError{msg: "test error"},
			expected: "test error",
		},
		{
			name: "error with cause",
			err: &AuthenticationError{
				msg:   "wrapper",
				cause: errors.New("underlying"),
			},
			expected: "wrapper: underlying",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Error() != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, tt.err.Error())
			}
		})
	}
}

func TestAuthenticationError_Unwrap(t *testing.T) {
	underlyingErr := errors.New("underlying error")
	authErr := &AuthenticationError{
		msg:   "wrapper",
		cause: underlyingErr,
	}

	unwrapped := authErr.Unwrap()
	if unwrapped != underlyingErr {
		t.Errorf("Expected unwrapped error to be %v, got %v", underlyingErr, unwrapped)
	}

	// Test with no cause
	authErrNoCause := &AuthenticationError{msg: "no cause"}
	if authErrNoCause.Unwrap() != nil {
		t.Error("Expected Unwrap to return nil when no cause")
	}
}

func TestCLIAuthProvider_IntegrationPattern(t *testing.T) {
	// This test demonstrates the expected integration pattern
	// (without actually importing the CLI)

	// Simulate authenticated scenario
	provider := &mockCLIAuthProvider{
		isAuthenticated: true,
		roundTripper:    &mockRoundTripper{},
	}

	// Create configuration
	cfg := &Configuration{
		HTTPClient: &http.Client{},
	}

	// Apply CLI auth configuration
	err := WithCLIProviderAuth(provider)(cfg)
	if err != nil {
		t.Fatalf("Failed to configure CLI auth: %v", err)
	}

	// Verify CustomAuth was set
	if cfg.CustomAuth == nil {
		t.Error("Expected CustomAuth to be configured")
	}

	// Verify it's the expected RoundTripper
	if cfg.CustomAuth != provider.roundTripper {
		t.Error("CustomAuth should be set to the provider's RoundTripper")
	}
}

func TestCLIAuthProvider_ChainedConfiguration(t *testing.T) {
	// Test that CLI auth can be chained with other configuration options
	provider := &mockCLIAuthProvider{
		isAuthenticated: true,
		roundTripper:    &mockRoundTripper{},
	}

	cfg := &Configuration{}

	// Chain multiple configuration options
	opts := []ConfigurationOption{
		WithRegion("eu01"),
		WithCLIProviderAuth(provider),
		WithUserAgent("test-agent"),
	}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			t.Fatalf("Configuration option failed: %v", err)
		}
	}

	// Verify all options were applied
	if cfg.Region != "eu01" {
		t.Error("Expected Region to be set")
	}
	if cfg.CustomAuth == nil {
		t.Error("Expected CustomAuth to be set")
	}
	if cfg.UserAgent != "test-agent" {
		t.Error("Expected UserAgent to be set")
	}
}
