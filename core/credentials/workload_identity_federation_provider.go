package credentials

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oidcadapters"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
)

const (
	workloadIdentityClientIDEnv   = envServiceAccountEmail
	workloadIdentityDefaultLeeway = time.Minute
	workloadIdentityErrorPrefix   = "workload identity federation provider"
)

// WorkloadIdentityFederationProviderConfig contains the configuration for WorkloadIdentityFederationProvider.
type WorkloadIdentityFederationProviderConfig struct {
	// TokenURL overrides the token endpoint. If empty, STACKIT_IDP_TOKEN_ENDPOINT env var is used,
	// falling back to the default STACKIT WIF token endpoint.
	TokenURL string
	// ClientID is the service account email. If empty, STACKIT_SERVICE_ACCOUNT_EMAIL env var is used.
	ClientID string
	// FederatedTokenFunction provides the OIDC assertion token. If nil, the default filesystem reader is used.
	FederatedTokenFunction oidcadapters.OIDCTokenFunc
	// TokenRefreshLeeway controls how early before expiration the token is refreshed.
	// If zero, defaults to 1 minute.
	TokenRefreshLeeway time.Duration
	// HTTPClient is used for token requests. If nil, a default client with a 1-minute timeout is used.
	HTTPClient *http.Client
	// Optional scopes to request for the access token
	Scopes []string
	// Optional resources to request for the access token
	Resources []string
}

var _ TokenProvider = (*WorkloadIdentityFederationProvider)(nil)

// WorkloadIdentityFederation retrieves tokens using WIF flow.
type WorkloadIdentityFederationProvider struct {
	name                   string
	httpClient             *http.Client
	tokenURL               string
	clientID               string
	federatedTokenFunction oidcadapters.OIDCTokenFunc
	tokenLeeway            time.Duration
	cache                  tokenCache
	scopes                 []string
	resources              []string
}

type workloadIdentityTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// NewWorkloadIdentityFederationProvider creates a WorkloadIdentityFederation provider.
func NewWorkloadIdentityFederationProvider(cfg *WorkloadIdentityFederationProviderConfig) (*WorkloadIdentityFederationProvider, error) {
	if cfg == nil {
		return nil, fmt.Errorf("%s: config cannot be nil", workloadIdentityErrorPrefix)
	}
	tokenURL := cfg.TokenURL
	if tokenURL == "" {
		tokenURL = utils.GetEnvOrDefault(envIdpTokenEndpoint, wifDefaultTokenEndpoint)
	}

	clientID := cfg.ClientID
	if clientID == "" {
		clientID = utils.GetEnvOrDefault(workloadIdentityClientIDEnv, "")
	}
	if clientID == "" {
		return nil, fmt.Errorf("%s: client ID cannot be empty", workloadIdentityErrorPrefix)
	}

	federatedTokenFunction := cfg.FederatedTokenFunction
	if federatedTokenFunction == nil {
		federatedTokenFunction = oidcadapters.ReadJWTFromFileSystem(utils.GetEnvOrDefault(envFederatedTokenFile, wifDefaultFederatedTokenPath))
	}
	if _, err := federatedTokenFunction(context.Background()); err != nil {
		return nil, fmt.Errorf("%s: error reading federated token file: %w", workloadIdentityErrorPrefix, err)
	}

	httpClient := authHTTPClient(cfg.HTTPClient, time.Minute)

	leeway := cfg.TokenRefreshLeeway
	if leeway == 0 {
		leeway = workloadIdentityDefaultLeeway
	}

	return &WorkloadIdentityFederationProvider{
		name:                   "WorkloadIdentityFederationProvider",
		httpClient:             httpClient,
		tokenURL:               tokenURL,
		clientID:               clientID,
		federatedTokenFunction: federatedTokenFunction,
		tokenLeeway:            leeway,
		scopes:                 cfg.Scopes,
		resources:              cfg.Resources,
	}, nil
}

// Token returns a valid access token.
func (p *WorkloadIdentityFederationProvider) Token(ctx context.Context, opt TokenRequestOptions) (Token, error) {
	scopes, resources := effectiveRequest(opt, p.scopes, p.resources)
	return p.cache.get(ctx, requestKey(scopes, resources), p.tokenLeeway, func(ctx context.Context) (Token, error) {
		fresh, err := p.requestToken(ctx, scopes, resources)
		if err != nil {
			return Token{}, err
		}
		debugContext(ctx, "credentials: authenticated", "provider", p.name)
		return fresh, nil
	})
}

func (p *WorkloadIdentityFederationProvider) requestToken(ctx context.Context, scopes, resources []string) (Token, error) {
	clientAssertion, err := p.federatedTokenFunction(ctx)
	if err != nil {
		return Token{}, fmt.Errorf("%s: get federated token: %w", workloadIdentityErrorPrefix, err)
	}

	body := url.Values{}
	body.Set("grant_type", wifGrantType)
	body.Set("client_assertion_type", wifClientAssertionType)
	body.Set("client_assertion", clientAssertion)
	body.Set("client_id", p.clientID)
	setScopesAndResources(body, scopes, resources)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.tokenURL, strings.NewReader(body.Encode()))
	if err != nil {
		return Token{}, fmt.Errorf("%s: create token request: %w", workloadIdentityErrorPrefix, err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := p.httpClient.Do(req)
	if err != nil {
		return Token{}, fmt.Errorf("%s: request access token: %w", workloadIdentityErrorPrefix, err)
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		bodyRaw, _ := io.ReadAll(res.Body)
		return Token{}, fmt.Errorf("%s: token request failed with status %d: %s", workloadIdentityErrorPrefix, res.StatusCode, string(bodyRaw))
	}

	var tokenResponse workloadIdentityTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&tokenResponse); err != nil {
		return Token{}, fmt.Errorf("%s: decode token response: %w", workloadIdentityErrorPrefix, err)
	}
	if tokenResponse.AccessToken == "" {
		return Token{}, fmt.Errorf("%s: token response did not include access token", workloadIdentityErrorPrefix)
	}

	expiresOn, err := getTokenExpiration(tokenResponse.AccessToken, tokenResponse.ExpiresIn)
	if err != nil {
		return Token{}, err
	}

	refreshOn := expiresOn.Add(-p.tokenLeeway)
	return Token{
		AccessToken: tokenResponse.AccessToken,
		TokenType:   tokenType(tokenResponse.TokenType),
		ExpiresOn:   expiresOn,
		RefreshOn:   refreshOn,
	}, nil
}
