package identity

import (
	"context"
	"fmt"
	"os"
)

var _ TokenProvider = (*StaticTokenProvider)(nil)

// StaticTokenProviderConfig contains configuration for StaticTokenProvider.
type StaticTokenProviderConfig struct {
	// Token is the static access token. If empty, will attempt to resolve from STACKIT_SERVICE_ACCOUNT_TOKEN env var.
	Token string
}

// StaticTokenProvider provides a static access token.
type StaticTokenProvider struct {
	name  string
	token Token
}

// NewStaticTokenProvider creates a StaticTokenProvider, resolving the token from config or environment.
func NewStaticTokenProvider(cfg StaticTokenProviderConfig) (*StaticTokenProvider, error) {
	token := cfg.Token
	if token == "" {
		token, _ = os.LookupEnv("STACKIT_SERVICE_ACCOUNT_TOKEN")
	}
	if token == "" {
		return nil, fmt.Errorf("static token provider: STACKIT_SERVICE_ACCOUNT_TOKEN not set and no token provided in config")
	}

	expiresOn, err := getTokenExpiration(token, 0)
	if err != nil {
		return nil, fmt.Errorf("static token expiration invalid: %w", err)
	}

	accessToken := Token{
		AccessToken: token,
		ExpiresOn:   expiresOn,
		RefreshOn:   expiresOn,
	}

	return &StaticTokenProvider{
		name:  "StaticTokenProvider",
		token: accessToken,
	}, nil
}

// Token returns the static access token, parsing its expiration from the JWT claims.
func (p *StaticTokenProvider) Token(_ context.Context, _ TokenRequestOptions) (Token, error) {
	if p.token.AccessToken == "" {
		return Token{}, fmt.Errorf("static token is empty")
	}
	return p.token, nil
}
