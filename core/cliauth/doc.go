// Package cliauth provides authentication support for STACKIT CLI provider credentials.
//
// This package enables applications (like the Terraform Provider) to use credentials
// stored by the STACKIT CLI without direct dependency on CLI code. It supports:
//
//   - Reading credentials from system keyring or file fallback
//   - Automatic OAuth2 token refresh
//   - Multiple CLI profiles
//   - Bidirectional credential sync (writeback after refresh)
//
// # Storage Locations
//
// Credentials are stored in two locations with automatic fallback:
//
//  1. System Keyring (preferred):
//     - macOS: Keychain
//     - Linux: Secret Service API / libsecret
//     - Windows: Credential Manager
//     - Service name: "stackit-cli-provider" or "stackit-cli-provider/{profile}"
//
//  2. File Fallback:
//     - Default profile: ~/.stackit/cli-provider-auth-storage.txt
//     - Custom profiles: ~/.stackit/profiles/{profile}/cli-provider-auth-storage.txt
//     - Format: Base64-encoded JSON
//
// # Profile Resolution
//
// Profiles are resolved in the following order:
//  1. Explicit profile parameter
//  2. STACKIT_CLI_PROFILE environment variable
//  3. ~/.config/stackit/cli-profile.txt
//  4. "default"
//
// # Usage Example - RoundTripper (Recommended)
//
// The recommended way to use this package is through the CLIProviderFlow,
// which implements http.RoundTripper and handles automatic token refresh:
//
//	flow, err := cliauth.NewCLIProviderFlow("", nil, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	client := &http.Client{Transport: flow}
//	// Token refresh happens automatically on requests
//
// # Usage Example - Direct Credential Access
//
// For advanced use cases, you can directly access credentials:
//
//	// Read credentials
//	creds, err := cliauth.ReadCredentials("")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Check if token is expired
//	if cliauth.IsTokenExpired(creds) {
//	    err = cliauth.RefreshToken(creds)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	}
//
//	// Use the access token
//	req.Header.Set("Authorization", "Bearer "+creds.AccessToken)
//
// # Usage Example - With Custom Profile
//
//	// Use a specific profile (e.g., "production")
//	flow, err := cliauth.NewCLIProviderFlow("production", nil, nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	client := &http.Client{Transport: flow}
//
// # Backward Compatibility
//
// This package maintains 100% backward compatibility with credentials created by
// existing STACKIT CLI versions. All file paths, formats, and keyring service names
// match the CLI exactly. Users can seamlessly switch between CLI and SDK-based tools
// without re-authenticating.
//
// # Thread Safety
//
// The CLIProviderFlow type is thread-safe and can be used concurrently from
// multiple goroutines. All other functions are safe to call concurrently, but
// they operate on independent credential instances.
//
// # Error Handling
//
// All functions return descriptive errors that can be inspected using standard
// Go error handling patterns. Common error scenarios include:
//
//   - No credentials found (user not authenticated)
//   - Expired refresh token (re-authentication required)
//   - Network errors during token refresh
//   - File system errors (permissions, missing directories)
//   - Keyring access errors (platform-specific)
package cliauth
