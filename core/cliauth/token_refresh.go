package cliauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// STACKIT OAuth2 token endpoint
	tokenEndpoint = "https://accounts.stackit.cloud/oauth2/token"
	// CLI client ID for OAuth2
	cliClientID = "stackit-cli-0000-0000-000000000001"
)

// RefreshTokenResponse represents the response from the OAuth2 token refresh endpoint
type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

// RefreshToken refreshes an expired access token using the refresh token.
// It updates the credentials in place and writes them back to storage.
//
// The function calls the STACKIT OAuth2 token endpoint with a refresh_token grant.
// If successful, it updates the access token, and optionally the refresh token and
// expiry time if provided in the response. The updated credentials are written
// back to storage (keyring or file) for persistence.
//
// Returns an error if the refresh token is empty, the HTTP request fails, or
// the token endpoint returns an error.
func RefreshToken(creds *ProviderCredentials) error {
	return RefreshTokenWithClient(creds, nil)
}

// RefreshTokenWithClient refreshes an access token using a custom HTTP client.
// If httpClient is nil, a default client with 30s timeout is used.
//
// This function is useful for testing or when custom HTTP client configuration
// (e.g., custom transport, timeouts, or proxies) is required.
func RefreshTokenWithClient(creds *ProviderCredentials, httpClient *http.Client) error {
	if creds == nil {
		return fmt.Errorf("credentials cannot be nil")
	}

	if creds.RefreshToken == "" {
		return fmt.Errorf("refresh token is empty")
	}

	// Determine which token endpoint to use
	endpoint := creds.TokenEndpoint
	if endpoint == "" {
		// Fallback to default endpoint if not set in credentials
		endpoint = tokenEndpoint
	}

	// Use default client if none provided
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 30 * time.Second}
	}

	// Build refresh request
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", creds.RefreshToken)
	data.Set("client_id", cliClientID)

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Execute request
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("token refresh failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var result RefreshTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	// Update credentials
	creds.AccessToken = result.AccessToken
	if result.RefreshToken != "" {
		creds.RefreshToken = result.RefreshToken
	}
	if result.ExpiresIn > 0 {
		creds.SessionExpiresAt = time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)
	}

	// Write back to storage
	if err := WriteCredentials(creds); err != nil {
		return fmt.Errorf("write refreshed credentials: %w", err)
	}

	return nil
}

// IsTokenExpired checks if the access token has expired or will expire soon.
// It uses a 5-minute safety margin to consider a token expired before its
// actual expiration time. This helps prevent using a token that might expire
// during a long-running operation.
//
// Returns true if:
//   - credentials are nil
//   - current time + 5 minutes is after the expiration time
//
// Returns false if:
//   - no expiration time is set (SessionExpiresAt is zero)
//   - token is still valid with safety margin
func IsTokenExpired(creds *ProviderCredentials) bool {
	if creds == nil {
		return true
	}

	if creds.SessionExpiresAt.IsZero() {
		// No expiry time, assume valid
		return false
	}

	// Consider expired if within 5 minutes of expiry (safety margin)
	return time.Now().Add(5 * time.Minute).After(creds.SessionExpiresAt)
}

// EnsureValidToken checks if the token is expired and refreshes it if needed.
// This is a convenience function that combines token expiry checking with
// automatic refresh.
//
// Returns nil if the token is still valid, or if it was successfully refreshed.
// Returns an error if the token is expired and refresh fails.
func EnsureValidToken(creds *ProviderCredentials) error {
	if !IsTokenExpired(creds) {
		return nil
	}

	return RefreshToken(creds)
}
