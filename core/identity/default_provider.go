package identity

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oidcadapters"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
)

const defaultProviderErrorPrefix = "default provider"

var _ TokenProvider = (*DefaultProvider)(nil)

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
	// STACKIT_SERVICE_ACCOUNT_KEY and STACKIT_SERVICE_ACCOUNT_KEY_PATH are used.
	ServiceAccountKey string
	// PrivateKey is the RSA private key as a PEM-encoded string. If empty,
	// STACKIT_PRIVATE_KEY and STACKIT_PRIVATE_KEY_PATH are used.
	PrivateKey string
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
	// HTTPClient is used for token requests. If nil, a default client is used.
	HTTPClient *http.Client
	// Scopes are the optional OAuth2 scopes to request for the access token.
	Scopes []string
	// Resources are the optional resource identifiers to request access for.
	Resources []string
}

// DefaultProvider is the opinionated credential chain STACKIT ships out of the box.
// It tries, in order:
//
//  1. StaticTokenProvider — a pre-issued token from config, environment or credentials file.
//  2. ServiceAccountKeyProvider — the service account key flow.
//  3. WorkloadIdentityFederationProvider — workload identity federation.
//  4. InstanceMetadataProvider — the token of the service account attached to a STACKIT VM.
//
// Explicitly configured credentials therefore always take precedence over the ambient
// identity of the machine, which is only consulted once everything else has failed. This
// mirrors the credential chains of the other major cloud SDKs, where the instance metadata
// service is present but last (AWS EC2 IMDS, GCP metadata server, Azure managed identity).
// Placing it late also means its network probe is only paid as a last resort, so callers
// who are not running on a STACKIT VM do not wait on it.
//
// DefaultProvider is a convenience for getting started. In production, constructing
// the single provider you actually use makes authentication more predictable and
// easier to debug.
type DefaultProvider struct {
	chain *ChainedProvider
}

// NewDefaultProvider creates the default STACKIT credential chain. Providers whose
// credentials cannot be resolved are skipped; an error is returned only when no
// provider at all could be configured, listing why each one was skipped.
func NewDefaultProvider(cfg *DefaultProviderConfig) (*DefaultProvider, error) {
	if cfg == nil {
		cfg = &DefaultProviderConfig{}
	}
	var (
		providers []TokenProvider
		skipped   []string
	)
	add := func(name string, provider TokenProvider, err error) {
		if err != nil {
			skipped = append(skipped, fmt.Sprintf("%s: %v", name, err))
			return
		}
		providers = append(providers, provider)
	}

	staticProvider, err := NewStaticTokenProvider(&StaticTokenProviderConfig{
		Token:               cfg.Token,
		CredentialsFilePath: cfg.CredentialsFilePath,
	})
	add("StaticTokenProvider", staticProvider, err)

	keyProvider, err := NewServiceAccountKeyProvider(&ServiceAccountKeyProviderConfig{
		ServiceAccountKey:   cfg.ServiceAccountKey,
		PrivateKey:          cfg.PrivateKey,
		TokenURL:            cfg.TokenURL,
		TokenRefreshLeeway:  cfg.TokenRefreshLeeway,
		HTTPClient:          cfg.HTTPClient,
		CredentialsFilePath: cfg.CredentialsFilePath,
		Scopes:              cfg.Scopes,
		Resources:           cfg.Resources,
	})
	add("ServiceAccountKeyProvider", keyProvider, err)

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

	// Ambient VM identity goes last: it must never outrank credentials the caller
	// configured, and its probe is only worth paying once nothing else has worked.
	metadataProvider, err := NewInstanceMetadataProvider(&InstanceMetadataProviderConfig{
		ServiceAccountEmail: serviceAccountEmail(cfg),
		TokenRefreshLeeway:  cfg.TokenRefreshLeeway,
		HTTPClient:          cfg.HTTPClient,
	})
	add("InstanceMetadataProvider", metadataProvider, err)

	if len(providers) == 0 {
		return nil, fmt.Errorf("%s: no valid credentials were found: %s", defaultProviderErrorPrefix, strings.Join(skipped, "; "))
	}

	chain, err := NewChainedProvider(providers...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", defaultProviderErrorPrefix, err)
	}
	chain.name = "DefaultProvider"

	return &DefaultProvider{chain: chain}, nil
}

// Token returns the first token retrieved successfully from the default chain.
func (p *DefaultProvider) Token(ctx context.Context, options TokenRequestOptions) (Token, error) {
	return p.chain.Token(ctx, options)
}

// serviceAccountEmail resolves the service account email from config, then the
// environment, then the credentials file, matching how the other providers in the chain
// resolve their own credentials. InstanceMetadataProvider, unlike
// WorkloadIdentityFederationProvider, does not do this resolution itself.
func serviceAccountEmail(cfg *DefaultProviderConfig) string {
	if cfg.ServiceAccountEmail != "" {
		return cfg.ServiceAccountEmail
	}
	if email := utils.GetEnvOrDefault(EnvServiceAccountEmail, ""); email != "" {
		return email
	}
	credentials, err := ReadCredentialsFile(cfg.CredentialsFilePath)
	if err != nil {
		return ""
	}
	return credentials.ServiceAccountEmail
}
