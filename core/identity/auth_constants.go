package identity

// Workload Identity Federation constants
const (
	// WIF client assertion type for JWT bearer token
	WifClientAssertionType = "urn:schwarz:params:oauth:client-assertion-type:workload-jwt"
	// WIF grant type for token requests
	WifGrantType = "client_credentials"
	// WIF default token endpoint
	// nolint:gosec // G101 False positive: This is a constant URL, not a credential
	WifDefaultTokenEndpoint = "https://accounts.stackit.cloud/oauth/v2/token"
	// WIF default federated token file path in Kubernetes
	// nolint:gosec // G101 False positive: This is a constant path, not a credential
	WifDefaultFederatedTokenPath = "/var/run/secrets/stackit.cloud/serviceaccount/token"
	// WIF default token expiration leeway
	WifDefaultTokenExpiration = "1h"
)

// Service Account Key Flow constants
const (
	// Service Account Key API endpoint
	// nolint:gosec // G101 False positive: This is a constant URL, not a credential
	KeyFlowTokenAPI = "https://accounts.stackit.cloud/oauth/v2/token"
	// Default token type for Bearer token
	DefaultTokenType = "Bearer"
	// Default scope for token requests
	DefaultScope = ""
)
