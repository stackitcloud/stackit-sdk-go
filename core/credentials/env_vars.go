package credentials

// Environment variable names for STACKIT authentication
const (
	// Service Account authentication
	envServiceAccountEmail   = "STACKIT_SERVICE_ACCOUNT_EMAIL"
	envServiceAccountToken   = "STACKIT_SERVICE_ACCOUNT_TOKEN"
	envServiceAccountKey     = "STACKIT_SERVICE_ACCOUNT_KEY"
	envServiceAccountKeyPath = "STACKIT_SERVICE_ACCOUNT_KEY_PATH"
	envPrivateKey            = "STACKIT_PRIVATE_KEY"
	envPrivateKeyPath        = "STACKIT_PRIVATE_KEY_PATH"
	// nolint:gosec // G101 False positive: This is a constant path, not a credential
	envCredentialsPath = "STACKIT_CREDENTIALS_PATH"

	// Workload Identity Federation
	// nolint:gosec // G101 False positive: This is a constant env var name, not a credential
	envFederatedTokenFile = "STACKIT_FEDERATED_TOKEN_FILE"
	// nolint:gosec // G101 False positive: This is a constant env var name, not a credential
	envIdpTokenEndpoint = "STACKIT_IDP_TOKEN_ENDPOINT"

	// Chain control
	// envUseCLI adds the CLI session step to the default chain when set to a true value.
	// The step is off by default.
	envUseCLI = "STACKIT_USE_CLI"
)
