package cliauth

import (
	"fmt"
	"os"
	"time"
)

var (
	// Start refresh attempts this duration before token expiration
	defaultTimeStartBeforeTokenExpiration = 5 * time.Minute
	// Check context cancellation this frequently while waiting
	defaultTimeBetweenContextCheck = time.Second
	// Retry interval on refresh failures
	defaultTimeBetweenTries = 2 * time.Minute
)

// continuousRefreshToken continuously refreshes the CLI provider token in the background.
// It monitors token expiration and automatically refreshes before the token expires.
//
// The goroutine terminates when:
//   - The context is canceled
//   - A non-retryable error occurs
//
// To terminate this routine, cancel the context in flow.refreshContext.
func continuousRefreshToken(flow *CLIProviderFlow) {
	refresher := &continuousTokenRefresher{
		flow:                           flow,
		timeStartBeforeTokenExpiration: defaultTimeStartBeforeTokenExpiration,
		timeBetweenContextCheck:        defaultTimeBetweenContextCheck,
		timeBetweenTries:               defaultTimeBetweenTries,
	}
	err := refresher.continuousRefreshToken()
	fmt.Fprintf(os.Stderr, "CLI provider token refreshing terminated: %v\n", err)
}

type continuousTokenRefresher struct {
	flow                           *CLIProviderFlow
	timeStartBeforeTokenExpiration time.Duration
	timeBetweenContextCheck        time.Duration
	timeBetweenTries               time.Duration
}

// continuousRefreshToken runs the main background refresh loop.
// It waits until the token is close to expiring, then refreshes it.
// Always returns with a non-nil error (indicating why it terminated).
func (r *continuousTokenRefresher) continuousRefreshToken() error {
	// Compute initial refresh timestamp
	startRefreshTimestamp := r.getNextRefreshTimestamp()

	for {
		// Wait until it's time to refresh (or context is canceled)
		err := r.waitUntilTimestamp(startRefreshTimestamp)
		if err != nil {
			return err
		}

		// Check if context was canceled
		err = r.flow.refreshContext.Err()
		if err != nil {
			return fmt.Errorf("context canceled: %w", err)
		}

		// Attempt to refresh the token
		ok, err := r.refreshToken()
		if err != nil {
			return fmt.Errorf("refresh token: %w", err)
		}

		if !ok {
			// Refresh failed (but is retryable), try again later
			startRefreshTimestamp = time.Now().Add(r.timeBetweenTries)
			continue
		}

		// Refresh succeeded, compute next refresh time
		startRefreshTimestamp = r.getNextRefreshTimestamp()
	}
}

// getNextRefreshTimestamp calculates when the next token refresh should start.
// Returns now if token is already expired, otherwise returns expiry time minus safety margin.
func (r *continuousTokenRefresher) getNextRefreshTimestamp() time.Time {
	r.flow.tokenMutex.RLock()
	expiresAt := r.flow.creds.SessionExpiresAt
	r.flow.tokenMutex.RUnlock()

	// If no expiry time set, check again in 5 minutes
	if expiresAt.IsZero() {
		return time.Now().Add(5 * time.Minute)
	}

	// If already expired, refresh immediately
	if time.Now().After(expiresAt) {
		return time.Now()
	}

	// Schedule refresh before expiration (with safety margin)
	return expiresAt.Add(-r.timeStartBeforeTokenExpiration)
}

// waitUntilTimestamp blocks until the target timestamp is reached or context is canceled.
// Periodically checks if the context has been canceled.
func (r *continuousTokenRefresher) waitUntilTimestamp(timestamp time.Time) error {
	for time.Now().Before(timestamp) {
		// Check if context was canceled
		err := r.flow.refreshContext.Err()
		if err != nil {
			return fmt.Errorf("context canceled during wait: %w", err)
		}

		// Sleep briefly before checking again
		time.Sleep(r.timeBetweenContextCheck)
	}
	return nil
}

// refreshToken attempts to refresh the access token.
// Returns:
//   - (true, nil) if refresh succeeded
//   - (false, nil) if refresh failed but should be retried (e.g., network error)
//   - (false, err) if refresh failed and should not be retried (e.g., invalid refresh token)
func (r *continuousTokenRefresher) refreshToken() (bool, error) {
	// Acquire write lock for refresh
	r.flow.tokenMutex.Lock()
	defer r.flow.tokenMutex.Unlock()

	// Double-check if refresh is still needed (another goroutine might have refreshed)
	if !IsTokenExpired(r.flow.creds) {
		return true, nil
	}

	// Attempt refresh
	err := RefreshTokenWithClient(r.flow.creds, r.flow.httpClient)
	if err == nil {
		return true, nil
	}

	// Check if error is retryable
	// Network errors, 5xx errors are retryable
	// 4xx errors (invalid refresh token) are not retryable
	errStr := err.Error()

	// Non-retryable errors (invalid refresh token, auth errors)
	if contains(errStr, "status 400") || contains(errStr, "status 401") ||
	   contains(errStr, "status 403") || contains(errStr, "refresh token is empty") {
		return false, fmt.Errorf("token refresh failed (non-retryable): %w", err)
	}

	// Retryable errors (network issues, 5xx errors)
	return false, nil
}

// contains checks if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) &&
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
		containsMiddle(s, substr)))
}

func containsMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
