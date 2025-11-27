package config

import (
	"context"
	"fmt"

	"github.com/stackitcloud/stackit-sdk-go/core/cliauth"
)

// WithCLIProviderAuth returns a ConfigurationOption that configures authentication
// using STACKIT CLI API credentials.
//
// This option enables the SDK to use credentials stored by the STACKIT CLI
// (via 'stackit auth api login') directly, without requiring external adapters.
//
// Profile resolution order:
//  1. Explicit profile parameter (if non-empty)
//  2. STACKIT_CLI_PROFILE environment variable
//  3. ~/.config/stackit/cli-profile.txt
//  4. "default"
//
// The authentication flow:
//   - Reads credentials from system keyring or file fallback
//   - Automatically refreshes expired tokens
//   - Writes refreshed tokens back to storage (bidirectional sync)
//
// Returns an AuthenticationError if no CLI credentials are found or cannot be initialized.
//
// Example usage:
//
//	// Use default profile
//	client, err := dns.NewAPIClient(
//	    config.WithCLIProviderAuth(""),
//	)
//
//	// Use custom profile
//	client, err := dns.NewAPIClient(
//	    config.WithCLIProviderAuth("production"),
//	)
func WithCLIProviderAuth(profile string) ConfigurationOption {
	return func(c *Configuration) error {
		// Create CLI provider flow with optional background refresh context
		flow, err := cliauth.NewCLIProviderFlowWithContext(
			profile,
			nil,
			nil,
			c.BackgroundTokenRefreshContext,
		)
		if err != nil {
			// Return the error directly - it already has a good message
			return err
		}

		// Configure the SDK to use CLI authentication
		c.CustomAuth = flow
		return nil
	}
}

// AuthenticationError indicates that CLI provider authentication failed.
// This error is returned when credentials are not found or cannot be initialized.
type AuthenticationError struct {
	msg   string
	cause error
}

// Error implements the error interface.
func (e *AuthenticationError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.cause)
	}
	return e.msg
}

// Unwrap returns the underlying cause of the authentication error, if any.
// This allows errors.Is and errors.As to work with wrapped errors.
func (e *AuthenticationError) Unwrap() error {
	return e.cause
}

// WithCLIBackgroundTokenRefresh returns a ConfigurationOption that enables
// background token refresh for CLI API authentication.
//
// When enabled, a goroutine will monitor CLI token expiration and automatically
// refresh the token before it expires. The goroutine is terminated when the
// provided context is canceled.
//
// This option only has effect when used together with WithCLIProviderAuth.
// It must be applied BEFORE WithCLIProviderAuth in the configuration chain.
//
// Example usage:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	client, err := dns.NewAPIClient(
//	    config.WithCLIBackgroundTokenRefresh(ctx),
//	    config.WithCLIProviderAuth(""),
//	)
//
// Note: The background refresh goroutine will write status messages to stderr
// when it terminates.
func WithCLIBackgroundTokenRefresh(ctx context.Context) ConfigurationOption {
	return func(c *Configuration) error {
		if ctx == nil {
			return fmt.Errorf("context for CLI background token refresh cannot be nil")
		}

		// Store context for CLI auth flow to use
		// Note: This assumes CLIProviderAuth flow will check for this
		c.BackgroundTokenRefreshContext = ctx
		return nil
	}
}
