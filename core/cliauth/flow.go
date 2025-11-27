package cliauth

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

// CLIProviderFlow implements http.RoundTripper for CLI provider authentication.
// It handles automatic token refresh when access tokens expire and provides
// thread-safe concurrent access to credentials.
//
// The flow reads credentials from STACKIT CLI storage (keyring or file),
// adds authentication headers to HTTP requests, and automatically refreshes
// tokens when they expire.
//
// Optional background token refresh can be enabled by providing a context
// via WithBackgroundTokenRefresh. When enabled, a goroutine will monitor
// token expiration and refresh proactively.
type CLIProviderFlow struct {
	rt             http.RoundTripper
	profile        string
	creds          *ProviderCredentials
	tokenMutex     sync.RWMutex
	httpClient     *http.Client
	refreshContext context.Context // If set, enables background token refresh
	initialized    bool
}

// NewCLIProviderFlow creates a new CLI provider flow with the given profile.
// The profile parameter follows the same resolution order as ReadCredentials.
//
// If baseTransport is nil, http.DefaultTransport is used.
// If httpClient is nil, a default client is created for token refresh operations.
func NewCLIProviderFlow(profile string, baseTransport http.RoundTripper, httpClient *http.Client) (*CLIProviderFlow, error) {
	return NewCLIProviderFlowWithContext(profile, baseTransport, httpClient, nil)
}

// NewCLIProviderFlowWithContext creates a new CLI provider flow with optional background refresh.
// The profile parameter follows the same resolution order as ReadCredentials.
//
// If baseTransport is nil, http.DefaultTransport is used.
// If httpClient is nil, a default client is created for token refresh operations.
// If refreshContext is non-nil, background token refresh is enabled.
func NewCLIProviderFlowWithContext(profile string, baseTransport http.RoundTripper, httpClient *http.Client, refreshContext context.Context) (*CLIProviderFlow, error) {
	if baseTransport == nil {
		baseTransport = http.DefaultTransport
	}

	flow := &CLIProviderFlow{
		rt:             baseTransport,
		profile:        profile,
		httpClient:     httpClient,
		refreshContext: refreshContext,
	}

	// Initialize credentials
	if err := flow.init(); err != nil {
		return nil, err
	}

	return flow, nil
}

// init initializes the flow by reading credentials from storage
func (f *CLIProviderFlow) init() error {
	creds, err := ReadCredentials(f.profile)
	if err != nil {
		return fmt.Errorf("read CLI credentials: %w", err)
	}

	// Ensure token is valid before proceeding
	if IsTokenExpired(creds) {
		if err := RefreshTokenWithClient(creds, f.httpClient); err != nil {
			return fmt.Errorf("refresh expired token: %w", err)
		}
	}

	f.creds = creds
	f.initialized = true

	// Start background token refresh if context is provided
	if f.refreshContext != nil {
		go continuousRefreshToken(f)
	}

	return nil
}

// WithBackgroundTokenRefresh enables background token refresh for this flow.
// When enabled, a goroutine will monitor token expiration and automatically
// refresh the token before it expires.
//
// The goroutine terminates when the provided context is canceled.
//
// This method must be called before the flow is initialized (i.e., before
// NewCLIProviderFlow returns or before init() is called).
//
// Example:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	flow := &CLIProviderFlow{}
//	flow.WithBackgroundTokenRefresh(ctx)
//	flow, err := NewCLIProviderFlow("", nil, nil)
//
// Or more commonly, via the SDK configuration:
//
//	client, err := dns.NewAPIClient(
//	    config.WithCLIProviderAuth(""),
//	    config.WithCLIBackgroundTokenRefresh(ctx),
//	)
func (f *CLIProviderFlow) WithBackgroundTokenRefresh(ctx context.Context) *CLIProviderFlow {
	f.refreshContext = ctx
	return f
}

// RoundTrip implements the http.RoundTripper interface.
// It adds the Authorization header with the current access token and
// handles automatic token refresh if needed.
//
// This method is thread-safe and can be called concurrently from multiple goroutines.
func (f *CLIProviderFlow) RoundTrip(req *http.Request) (*http.Response, error) {
	if !f.initialized {
		return nil, fmt.Errorf("CLI provider flow not initialized")
	}

	// Get current token with read lock
	f.tokenMutex.RLock()
	token := f.creds.AccessToken
	needsRefresh := IsTokenExpired(f.creds)
	f.tokenMutex.RUnlock()

	// Refresh token if needed
	if needsRefresh {
		if err := f.refreshToken(); err != nil {
			return nil, fmt.Errorf("refresh token: %w", err)
		}

		// Get refreshed token
		f.tokenMutex.RLock()
		token = f.creds.AccessToken
		f.tokenMutex.RUnlock()
	}

	// Clone request to avoid modifying the original
	clonedReq := req.Clone(req.Context())
	if clonedReq.Header == nil {
		clonedReq.Header = make(http.Header)
	}

	// Add Authorization header
	clonedReq.Header.Set("Authorization", "Bearer "+token)

	// Execute request
	return f.rt.RoundTrip(clonedReq)
}

// refreshToken refreshes the access token with write lock to prevent concurrent refreshes
func (f *CLIProviderFlow) refreshToken() error {
	f.tokenMutex.Lock()
	defer f.tokenMutex.Unlock()

	// Double-check if refresh is still needed (another goroutine might have refreshed)
	if !IsTokenExpired(f.creds) {
		return nil
	}

	// Refresh the token
	if err := RefreshTokenWithClient(f.creds, f.httpClient); err != nil {
		return err
	}

	return nil
}

// GetCredentials returns a copy of the current credentials.
// This method is thread-safe.
//
// Note: The returned credentials are a snapshot and may become outdated
// if the flow refreshes tokens in the background.
func (f *CLIProviderFlow) GetCredentials() *ProviderCredentials {
	f.tokenMutex.RLock()
	defer f.tokenMutex.RUnlock()

	// Return a copy to prevent external modification
	credsCopy := *f.creds
	return &credsCopy
}
