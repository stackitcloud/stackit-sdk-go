package credentials

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oidcadapters"
)

const defaultProviderErrorPrefix = "default provider"

// DefaultProviderConfig contains the configuration for NewDefaultProvider.
//
// Every field is optional. Each provider in the chain resolves what it needs from this
// config first, then from environment variables, and finally from the credentials file.
// Providers whose credentials cannot be resolved are left out of the chain rather than
// failing construction.
type DefaultProviderConfig struct {
	// Token is a pre-issued access token. If empty, STACKIT_SERVICE_ACCOUNT_TOKEN is used.
	Token string
	// ServiceAccountKey is the service account key as a JSON string. If empty,
	// ServiceAccountKeyPath is used, then STACKIT_SERVICE_ACCOUNT_KEY and
	// STACKIT_SERVICE_ACCOUNT_KEY_PATH.
	ServiceAccountKey string
	// ServiceAccountKeyPath is the path of a file with the service account key. It is used
	// when ServiceAccountKey is empty.
	ServiceAccountKeyPath string
	// PrivateKey is the RSA private key as a PEM-encoded string. If empty,
	// PrivateKeyPath is used, then STACKIT_PRIVATE_KEY and STACKIT_PRIVATE_KEY_PATH.
	PrivateKey string
	// PrivateKeyPath is the path of a file with the PEM-encoded RSA private key. It is used
	// when PrivateKey is empty.
	PrivateKeyPath string
	// ServiceAccountEmail identifies the service account for workload identity
	// federation and for the instance metadata service. If empty, it is resolved from
	// STACKIT_SERVICE_ACCOUNT_EMAIL and then from the credentials file.
	ServiceAccountEmail string
	// FederatedTokenFunction provides the OIDC assertion for workload identity
	// federation. If nil, the assertion is read from STACKIT_FEDERATED_TOKEN_FILE.
	FederatedTokenFunction oidcadapters.OIDCTokenFunc
	// TokenURL overrides the token endpoint for every provider in the chain.
	TokenURL string
	// CredentialsFilePath overrides the credentials file location. If empty, the path from
	// STACKIT_CREDENTIALS_PATH is used, falling back to ~/.stackit/credentials.json.
	CredentialsFilePath string
	// TokenRefreshLeeway controls how early before expiration tokens are refreshed.
	// If zero, each provider applies its own default.
	TokenRefreshLeeway time.Duration
	// CLICommand is the STACKIT CLI executable used to reuse a local CLI session.
	// If empty, "stackit" is resolved from PATH.
	CLICommand string
	// EnableCLI adds the session of a logged-in STACKIT CLI to the end of the chain. It is
	// meant for developer machines and is off by default, since it authenticates as the
	// person logged in to the CLI rather than as the application. If false, it can also be
	// enabled with STACKIT_USE_CLI=true.
	EnableCLI bool
	// HTTPClient is used for token requests. If nil, a default client is used.
	HTTPClient *http.Client
	// Scopes are the optional OAuth2 scopes to request for the access token.
	Scopes []string
	// Resources are the optional resource identifiers to request access for.
	Resources []string
}

// NewDefaultProvider creates the opinionated credential chain STACKIT ships out of the box.
// It tries, in order:
//
//  1. WorkloadIdentityFederationProvider — workload identity federation.
//  2. ServiceAccountKeyProvider — the service account key flow.
//  3. StaticTokenProvider — a pre-issued token from config, environment or credentials file.
//  4. InstanceMetadataProvider — the token of the service account attached to a STACKIT VM.
//  5. CLIProvider — the session of a logged-in STACKIT CLI, for developer machines.
//
// The most secure flows come first: workload identity federation needs no secret and issues
// short-lived tokens, the service account key flow issues short-lived tokens from a
// long-lived key, and a pre-issued token is the least secure, so it only applies when
// neither is available. The chain cannot tell whether a flow was configured in code or in
// the environment, so a flow set up in the environment can take precedence over a less
// secure one passed in DefaultProviderConfig.
//
// Credentials of the application therefore always take precedence over the ambient
// identity of the machine, which is only consulted once everything else has failed. This
// mirrors the credential chains of the other major cloud SDKs, where the instance metadata
// service is present but last (AWS EC2 IMDS, GCP metadata server, Azure managed identity).
// Placing it late also means its network probe is only paid as a last resort, so callers
// who are not running on a STACKIT VM do not wait on it.
//
// The CLI session is opt-in (DefaultProviderConfig.EnableCLI or STACKIT_USE_CLI=true) and
// goes last, as in the other SDKs, whose chains end with the local developer tooling
// (Azure's AzureCLICredential, GCP's gcloud credentials). It is not on by default because
// it authenticates as the person logged in rather than as the application, so a
// misconfigured service credential would otherwise fall back to a person's permissions
// without anyone noticing. Note that no interactive login is part of this chain: a default
// chain that can open a browser would be a poor surprise in CI or in a script, so obtaining
// a session interactively stays an explicit, opt-in operation.
//
// Providers whose credentials cannot be resolved are skipped; an error is returned only when
// no provider at all could be configured, listing why each one was skipped.
//
// The default chain is a convenience for getting started. In production, constructing the
// single provider you actually use makes authentication more predictable and easier to debug.
func NewDefaultProvider(cfg *DefaultProviderConfig) (*ChainedProvider, error) {
	if cfg == nil {
		cfg = &DefaultProviderConfig{}
	}
	var (
		providers []TokenProvider
		skipped   []string
	)
	// The credentials file is read once for the whole chain, not once per provider.
	file := newCredentialsFileSource(cfg.CredentialsFilePath)
	add := func(name string, provider TokenProvider, err error) {
		if err != nil {
			skipped = append(skipped, fmt.Sprintf("%s: %v", name, err))
			debugContext(context.Background(), "credentials: provider left out of the default chain", "provider", name, "reason", err.Error())
			return
		}
		providers = append(providers, provider)
	}

	wifProvider, err := NewWorkloadIdentityFederationProvider(&WorkloadIdentityFederationProviderConfig{
		TokenURL:               cfg.TokenURL,
		ClientID:               cfg.ServiceAccountEmail,
		FederatedTokenFunction: cfg.FederatedTokenFunction,
		TokenRefreshLeeway:     cfg.TokenRefreshLeeway,
		HTTPClient:             cfg.HTTPClient,
		Scopes:                 cfg.Scopes,
		Resources:              cfg.Resources,
	})
	add("WorkloadIdentityFederationProvider", wifProvider, err)

	keyProvider, err := NewServiceAccountKeyProvider(&ServiceAccountKeyProviderConfig{
		ServiceAccountKey:     cfg.ServiceAccountKey,
		ServiceAccountKeyPath: cfg.ServiceAccountKeyPath,
		PrivateKey:            cfg.PrivateKey,
		PrivateKeyPath:        cfg.PrivateKeyPath,
		TokenURL:              cfg.TokenURL,
		TokenRefreshLeeway:    cfg.TokenRefreshLeeway,
		HTTPClient:            cfg.HTTPClient,
		CredentialsFilePath:   cfg.CredentialsFilePath,
		Scopes:                cfg.Scopes,
		Resources:             cfg.Resources,
		credentialsFile:       file,
	})
	add("ServiceAccountKeyProvider", keyProvider, err)

	staticProvider, err := NewStaticTokenProvider(&StaticTokenProviderConfig{
		Token:               cfg.Token,
		CredentialsFilePath: cfg.CredentialsFilePath,
		credentialsFile:     file,
	})
	add("StaticTokenProvider", staticProvider, err)

	// Ambient VM identity goes last: it must never outrank credentials the caller
	// configured, and its probe is only worth paying once nothing else has worked.
	metadataProvider, err := NewInstanceMetadataProvider(&InstanceMetadataProviderConfig{
		ServiceAccountEmail: serviceAccountEmail(cfg, file),
		TokenRefreshLeeway:  cfg.TokenRefreshLeeway,
		HTTPClient:          cfg.HTTPClient,
	})
	add("InstanceMetadataProvider", metadataProvider, err)

	// The CLI session is the developer-machine fallback, so it goes last. It is opt-in, and
	// is skipped wherever the CLI is not installed, which is every CI runner, container and
	// server.
	if reason := cliStepDisabledReason(cfg); reason != "" {
		skipped = append(skipped, "CLIProvider: "+reason)
		debugContext(context.Background(), "credentials: provider left out of the default chain", "provider", "CLIProvider", "reason", reason)
	} else {
		cliProvider, err := NewCLIProvider(&CLIProviderConfig{
			Command:            cfg.CLICommand,
			TokenRefreshLeeway: cfg.TokenRefreshLeeway,
		})
		add("CLIProvider", cliProvider, err)
	}

	if len(providers) == 0 {
		return nil, fmt.Errorf("%s: no valid credentials were found: %s", defaultProviderErrorPrefix, strings.Join(skipped, "; "))
	}

	chain, err := NewChainedProvider(providers...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", defaultProviderErrorPrefix, err)
	}
	chain.name = "DefaultProvider"

	return chain, nil
}

// serviceAccountEmail resolves the service account email from config, then the
// environment, then the credentials file, matching how the other providers in the chain
// resolve their own credentials. InstanceMetadataProvider, unlike
// WorkloadIdentityFederationProvider, does not do this resolution itself.
func serviceAccountEmail(cfg *DefaultProviderConfig, file *credentialsFileSource) string {
	email, _ := resolveValue(cfg.ServiceAccountEmail, "", envServiceAccountEmail, "",
		fromFile(file, func(c *credentialsFile) string { return c.ServiceAccountEmail }))
	return email
}

// cliStepDisabledReason reports why the default chain leaves out the CLI session step, or
// an empty string when the step is included. The step is opt-in: it is included when
// DefaultProviderConfig.EnableCLI is set or, failing that, when STACKIT_USE_CLI is true.
// Nothing turns it off once it has been enabled, so the code never overrides the environment
// nor the other way around. A value of STACKIT_USE_CLI that is not a boolean leaves it off.
//
// Constructing a CLIProvider directly is unaffected — an explicit choice is not overridden.
func cliStepDisabledReason(cfg *DefaultProviderConfig) string {
	if cfg.EnableCLI {
		return ""
	}
	if enabled, err := strconv.ParseBool(os.Getenv(envUseCLI)); err == nil && enabled {
		return ""
	}
	return fmt.Sprintf("not enabled (set DefaultProviderConfig.EnableCLI or %s=true)", envUseCLI)
}
