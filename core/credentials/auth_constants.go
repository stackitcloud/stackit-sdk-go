package credentials

// Workload Identity Federation constants
const (
	// WIF client assertion type for JWT bearer token
	wifClientAssertionType = "urn:schwarz:params:oauth:client-assertion-type:workload-jwt"
	// WIF grant type for token requests
	wifGrantType = "client_credentials"
	// WIF default token endpoint
	// nolint:gosec // G101 False positive: This is a constant URL, not a credential
	wifDefaultTokenEndpoint = "https://accounts.stackit.cloud/oauth/v2/token"
	// WIF default federated token file path in Kubernetes
	// nolint:gosec // G101 False positive: This is a constant path, not a credential
	wifDefaultFederatedTokenPath = "/var/run/secrets/stackit.cloud/serviceaccount/token"
)

// Token constants
const (
	// defaultTokenType is the authorization scheme of every token STACKIT issues today.
	defaultTokenType = "Bearer"
)
