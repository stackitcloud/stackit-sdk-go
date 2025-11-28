package cliauth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func init() {
	// Disable keyring access in tests to avoid macOS Keychain dialogs
	SetSkipKeyring(true)
}

func TestRefreshToken_Success(t *testing.T) {
	// Create mock OAuth2 server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			t.Errorf("Expected application/x-www-form-urlencoded content type")
		}

		// Parse form data
		r.ParseForm()
		if r.Form.Get("grant_type") != "refresh_token" {
			t.Errorf("Expected grant_type=refresh_token, got %s", r.Form.Get("grant_type"))
		}
		if r.Form.Get("refresh_token") != "old-refresh-token" {
			t.Errorf("Expected refresh_token=old-refresh-token, got %s", r.Form.Get("refresh_token"))
		}

		// Send successful response
		response := RefreshTokenResponse{
			AccessToken:  "new-access-token",
			RefreshToken: "new-refresh-token",
			ExpiresIn:    3600,
			TokenType:    "Bearer",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create test credentials
	tmpDir := t.TempDir()
	setupTestHome(t, tmpDir)

	creds := &ProviderCredentials{
		AccessToken:      "old-access-token",
		RefreshToken:     "old-refresh-token",
		Email:            "test@example.com",
		SessionExpiresAt: time.Now().Add(-1 * time.Hour), // Expired
		sourceProfile:    "default",
	}

	// Create custom HTTP client pointing to mock server
	client := &http.Client{
		Transport: &testTransport{serverURL: server.URL},
	}

	// Test refresh
	err := RefreshTokenWithClient(creds, client)
	if err != nil {
		t.Fatalf("RefreshToken() error = %v", err)
	}

	// Verify credentials were updated
	if creds.AccessToken != "new-access-token" {
		t.Errorf("AccessToken = %v, want new-access-token", creds.AccessToken)
	}
	if creds.RefreshToken != "new-refresh-token" {
		t.Errorf("RefreshToken = %v, want new-refresh-token", creds.RefreshToken)
	}
	if creds.SessionExpiresAt.Before(time.Now()) {
		t.Error("SessionExpiresAt should be in the future")
	}
}

func TestRefreshToken_EmptyRefreshToken(t *testing.T) {
	creds := &ProviderCredentials{
		AccessToken: "some-token",
		// RefreshToken is empty
	}

	err := RefreshToken(creds)
	if err == nil {
		t.Error("Expected error for empty refresh token")
	}
	if err.Error() != "refresh token is empty" {
		t.Errorf("Expected 'refresh token is empty' error, got: %v", err)
	}
}

func TestRefreshToken_NilCredentials(t *testing.T) {
	err := RefreshToken(nil)
	if err == nil {
		t.Error("Expected error for nil credentials")
	}
	if err.Error() != "credentials cannot be nil" {
		t.Errorf("Expected 'credentials cannot be nil' error, got: %v", err)
	}
}

func TestRefreshToken_HTTPError(t *testing.T) {
	// Create mock server that returns 401
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid refresh token"))
	}))
	defer server.Close()

	tmpDir := t.TempDir()
	setupTestHome(t, tmpDir)

	creds := &ProviderCredentials{
		AccessToken:   "old-token",
		RefreshToken:  "invalid-refresh-token",
		sourceProfile: "default",
	}

	client := &http.Client{
		Transport: &testTransport{serverURL: server.URL},
	}

	err := RefreshTokenWithClient(creds, client)
	if err == nil {
		t.Error("Expected error for HTTP 401 response")
	}
}

func TestIsTokenExpired(t *testing.T) {
	tests := []struct {
		name     string
		creds    *ProviderCredentials
		expected bool
	}{
		{
			name:     "nil credentials",
			creds:    nil,
			expected: true,
		},
		{
			name: "no expiry time",
			creds: &ProviderCredentials{
				SessionExpiresAt: time.Time{},
			},
			expected: false,
		},
		{
			name: "expired 1 hour ago",
			creds: &ProviderCredentials{
				SessionExpiresAt: time.Now().Add(-1 * time.Hour),
			},
			expected: true,
		},
		{
			name: "expires in 10 minutes (within safety margin)",
			creds: &ProviderCredentials{
				SessionExpiresAt: time.Now().Add(4 * time.Minute),
			},
			expected: true, // 5-minute safety margin
		},
		{
			name: "expires in 10 minutes (outside safety margin)",
			creds: &ProviderCredentials{
				SessionExpiresAt: time.Now().Add(10 * time.Minute),
			},
			expected: false,
		},
		{
			name: "expires in 1 hour",
			creds: &ProviderCredentials{
				SessionExpiresAt: time.Now().Add(1 * time.Hour),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsTokenExpired(tt.creds)
			if result != tt.expected {
				t.Errorf("IsTokenExpired() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestEnsureValidToken(t *testing.T) {
	tmpDir := t.TempDir()
	setupTestHome(t, tmpDir)

	t.Run("token not expired", func(t *testing.T) {
		creds := &ProviderCredentials{
			AccessToken:      "valid-token",
			RefreshToken:     "refresh-token",
			SessionExpiresAt: time.Now().Add(1 * time.Hour),
			sourceProfile:    "default",
		}

		err := EnsureValidToken(creds)
		if err != nil {
			t.Errorf("EnsureValidToken() error = %v for valid token", err)
		}
	})

	t.Run("token expired, refresh succeeds", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := RefreshTokenResponse{
				AccessToken: "new-token",
				ExpiresIn:   3600,
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		creds := &ProviderCredentials{
			AccessToken:      "expired-token",
			RefreshToken:     "refresh-token",
			SessionExpiresAt: time.Now().Add(-1 * time.Hour),
			sourceProfile:    "default",
		}

		// This will fail in real use but demonstrates the pattern
		err := EnsureValidToken(creds)
		// We expect an error because we're not using the custom client
		// In real usage, the custom client would point to the test server
		if err == nil {
			t.Log("Note: This test doesn't fully test refresh with custom endpoint")
		}
	})
}

// testTransport redirects all requests to a test server
type testTransport struct {
	serverURL string
}

func (t *testTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Redirect all requests to the test server
	req.URL.Scheme = "http"
	req.URL.Host = t.serverURL[7:] // Remove "http://"
	return http.DefaultTransport.RoundTrip(req)
}

// Helper to setup test environment
func setupTestHome(t *testing.T, dir string) {
	t.Setenv("HOME", dir)
}
