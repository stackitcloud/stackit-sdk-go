package identity

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	instanceMetadataEndpointTemplate = "http://169.254.169.254/stackit/v1/service-accounts/<SERVICEACCOUNTMAIL>/token"
	instanceMetadataPathPlaceholder  = "<SERVICEACCOUNTMAIL>"
	instanceMetadataDefaultLeeway    = time.Minute
	instanceMetadataErrorPrefix      = "instance metadata provider"
)

// InstanceMetadataProviderConfig contains the configuration for InstanceMetadataProvider.
type InstanceMetadataProviderConfig struct {
	ServiceAccountEmail string
	// TokenRefreshLeeway controls how early before expiration the token is refreshed.
	// If zero, defaults to 1 minute.
	TokenRefreshLeeway time.Duration
	// HTTPClient is used for token requests. If nil, a default client with a 1-minute timeout is used.
	HTTPClient *http.Client
}

var _ TokenProvider = (*InstanceMetadataProvider)(nil)

// InstanceMetadata retrieves access tokens from STACKIT VM instance metadata endpoint.
type InstanceMetadataProvider struct {
	name             string
	httpClient       *http.Client
	serviceAccount   string
	endpointTemplate string
	tokenLeeway      time.Duration

	tokenMutex sync.RWMutex
	token      Token
}

type instanceMetadataTokenResponse struct {
	Token      string `json:"token"`
	ValidUntil string `json:"validUntil"`
}

// NewInstanceMetadataProvider creates an InstanceMetadata provider.
func NewInstanceMetadataProvider(cfg *InstanceMetadataProviderConfig) (*InstanceMetadataProvider, error) {
	if cfg == nil {
		return nil, fmt.Errorf("%s: config cannot be nil", instanceMetadataErrorPrefix)
	}

	if cfg.ServiceAccountEmail == "" {
		return nil, fmt.Errorf("%s: service account email cannot be empty", instanceMetadataErrorPrefix)
	}

	// Create a clean client to avoid deadlocks if the provided client already has auth roundtrippers
	httpClient := authHTTPClient(cfg.HTTPClient, time.Minute)

	leeway := cfg.TokenRefreshLeeway
	if leeway == 0 {
		leeway = instanceMetadataDefaultLeeway
	}

	return &InstanceMetadataProvider{
		name:             "InstanceMetadataProvider",
		httpClient:       httpClient,
		serviceAccount:   cfg.ServiceAccountEmail,
		endpointTemplate: instanceMetadataEndpointTemplate,
		tokenLeeway:      leeway,
	}, nil
}

// Token returns a valid access token.
func (p *InstanceMetadataProvider) Token(ctx context.Context, _ TokenRequestOptions) (Token, error) {
	if p == nil {
		return Token{}, fmt.Errorf("%s: provider is not initialized", instanceMetadataErrorPrefix)
	}
	if ctx == nil {
		ctx = context.Background()
	}

	now := time.Now().Add(p.tokenLeeway)
	p.tokenMutex.RLock()
	cached := p.token
	p.tokenMutex.RUnlock()
	if cached.AccessToken != "" && cached.RefreshOn.After(now) {
		return cached, nil
	}

	fresh, err := p.requestToken(ctx)
	if err != nil {
		return Token{}, err
	}
	DebugContext(ctx, "identity: authenticated", "provider", p.name)

	p.tokenMutex.Lock()
	p.token = fresh
	p.tokenMutex.Unlock()

	return fresh, nil
}

func (p *InstanceMetadataProvider) requestToken(ctx context.Context) (Token, error) {
	url := strings.ReplaceAll(p.endpointTemplate, instanceMetadataPathPlaceholder, p.serviceAccount)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return Token{}, fmt.Errorf("%s: create instance metadata request: %w", instanceMetadataErrorPrefix, err)
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return Token{}, fmt.Errorf("%s: request instance metadata token: %w", instanceMetadataErrorPrefix, err)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return Token{}, fmt.Errorf("%s: instance metadata request failed with status %d", instanceMetadataErrorPrefix, res.StatusCode)
	}

	var tokenResponse instanceMetadataTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&tokenResponse); err != nil {
		return Token{}, fmt.Errorf("%s: decode instance metadata token response: %w", instanceMetadataErrorPrefix, err)
	}
	if tokenResponse.Token == "" {
		return Token{}, fmt.Errorf("%s: instance metadata token response did not include token", instanceMetadataErrorPrefix)
	}
	if tokenResponse.ValidUntil == "" {
		return Token{}, fmt.Errorf("%s: instance metadata token response did not include validUntil", instanceMetadataErrorPrefix)
	}

	expiresAt, err := time.Parse(time.RFC3339Nano, tokenResponse.ValidUntil)
	if err != nil {
		expiresAt, err = time.Parse(time.RFC3339, tokenResponse.ValidUntil)
		if err != nil {
			return Token{}, fmt.Errorf("%s: parse validUntil timestamp: %w", instanceMetadataErrorPrefix, err)
		}
	}

	return Token{
		AccessToken: tokenResponse.Token,
		ExpiresOn:   expiresAt,
		RefreshOn:   expiresAt,
	}, nil
}
