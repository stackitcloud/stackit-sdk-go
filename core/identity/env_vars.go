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
	EnvCredentialsPath       = "STACKIT_CREDENTIALS_PATH"

	// Workload Identity Federation
	EnvFederatedTokenFile    = "STACKIT_FEDERATED_TOKEN_FILE"
	EnvIdpTokenEndpoint      = "STACKIT_IDP_TOKEN_ENDPOINT"
	EnvIdpTokenExpirationSec = "STACKIT_IDP_TOKEN_EXPIRATION_SECONDS"

	// Token endpoint
	EnvTokenBaseUrl = "STACKIT_TOKEN_BASEURL"
)
