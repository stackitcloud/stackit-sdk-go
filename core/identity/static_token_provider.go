package identity

import (
	"context"
	"fmt"
	"os"
)

const staticTokenErrorPrefix = "static token provider"

var _ TokenProvider = (*StaticTokenProvider)(nil)

// StaticTokenProviderConfig contains configuration for StaticTokenProvider.
type StaticTokenProviderConfig struct {
	// Token is the static access token. If empty, will attempt to resolve from STACKIT_SERVICE_ACCOUNT_TOKEN env var
	// or from credentials file if CredentialsFilePath is set.
	Token string
	// CredentialsFilePath is a pointer to the path to the credentials file. When set (not nil), it enables
	// automatic resolution of token from the credentials file if it is not found in config or environment variables.
	// If the pointer is nil, credentials file resolution is skipped.
	// When set to a pointer to an empty string, uses the default credentials file path (~/.stackit/credentials.json).
	CredentialsFilePath *string
}

// StaticTokenProvider provides a static access token.
type StaticTokenProvider struct {
	name  string
	token Token
}

// NewStaticTokenProvider creates a StaticTokenProvider, resolving the token from config or environment.
func NewStaticTokenProvider(cfg *StaticTokenProviderConfig) (*StaticTokenProvider, error) {
	if cfg == nil {
		return nil, fmt.Errorf("%s: config cannot be nil", staticTokenErrorPrefix)
	}
	token := cfg.Token
	if token == "" {
		if val, found := os.LookupEnv(EnvServiceAccountToken); found && val != "" {
			token = val
		} else if cfg.CredentialsFilePath != nil {
			// Try credentials file if pointer is set (not nil)
			credentials, err := ReadCredentialsFile(*cfg.CredentialsFilePath)
			if err == nil {
				if credToken, err := ReadCredential(CredentialTypeToken, credentials); err == nil {
					token = credToken
				}
			}
		}
	}
	if token == "" {
		return nil, fmt.Errorf("static token provider: %s not set and no token provided in config", EnvServiceAccountToken)
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
		return Token{}, fmt.Errorf("static token provider: token is empty")
	}
	return p.token, nil
}
