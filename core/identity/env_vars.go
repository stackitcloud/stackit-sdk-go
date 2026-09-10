package identity

// Environment variable names for STACKIT authentication
const (
	// Service Account authentication
	EnvServiceAccountEmail   = "STACKIT_SERVICE_ACCOUNT_EMAIL"
	EnvServiceAccountToken   = "STACKIT_SERVICE_ACCOUNT_TOKEN"
	EnvServiceAccountKey     = "STACKIT_SERVICE_ACCOUNT_KEY"
	EnvServiceAccountKeyPath = "STACKIT_SERVICE_ACCOUNT_KEY_PATH"
	EnvPrivateKey            = "STACKIT_PRIVATE_KEY"
	EnvPrivateKeyPath        = "STACKIT_PRIVATE_KEY_PATH"
	// nolint:gosec // G101 False positive: This is a constant path, not a credential
	EnvCredentialsPath = "STACKIT_CREDENTIALS_PATH"

	// Workload Identity Federation
	// nolint:gosec // G101 False positive: This is a constant env var name, not a credential
	EnvFederatedTokenFile = "STACKIT_FEDERATED_TOKEN_FILE"
	// nolint:gosec // G101 False positive: This is a constant env var name, not a credential
	EnvIdpTokenEndpoint = "STACKIT_IDP_TOKEN_ENDPOINT"
	// nolint:gosec // G101 False positive: This is a constant env var name, not a credential
	EnvIdpTokenExpirationSec = "STACKIT_IDP_TOKEN_EXPIRATION_SECONDS"

	// Token endpoint
	// nolint:gosec // G101 False positive: This is a constant env var name, not a credential
	EnvTokenBaseUrl = "STACKIT_TOKEN_BASEURL"
)
