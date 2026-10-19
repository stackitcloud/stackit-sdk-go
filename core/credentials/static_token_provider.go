package credentials

import (
	"context"
	"fmt"
)

const staticTokenErrorPrefix = "static token provider"

var _ TokenProvider = (*StaticTokenProvider)(nil)

// StaticTokenProviderConfig contains configuration for StaticTokenProvider.
type StaticTokenProviderConfig struct {
	// Token is the static access token. If empty, it is resolved from the
	// STACKIT_SERVICE_ACCOUNT_TOKEN env var and then from the credentials file.
	Token string
	// CredentialsFilePath overrides the credentials file location. If empty, the path from
	// STACKIT_CREDENTIALS_PATH is used, falling back to ~/.stackit/credentials.json.
	CredentialsFilePath string

	// credentialsFile is shared by NewDefaultProvider so the file is read once per chain.
	credentialsFile *credentialsFileSource
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
	file := fileSourceFor(cfg.credentialsFile, cfg.CredentialsFilePath)
	token, _ := resolveValue(cfg.Token, "", envServiceAccountToken, "",
		fromFile(file, func(c *credentialsFile) string { return c.ServiceAccountToken }))
	if token == "" {
		return nil, fmt.Errorf("static token provider: %s not set and no token provided in config", envServiceAccountToken)
	}

	expiresOn, err := getTokenExpiration(token, 0)
	if err != nil {
		return nil, fmt.Errorf("static token expiration invalid: %w", err)
	}

	accessToken := Token{
		AccessToken: token,
		// A pre-issued token carries no type of its own.
		TokenType: defaultTokenType,
		ExpiresOn: expiresOn,
		RefreshOn: expiresOn,
	}

	return &StaticTokenProvider{
		name:  "StaticTokenProvider",
		token: accessToken,
	}, nil
}

// Token returns the static access token, parsing its expiration from the JWT claims.
func (p *StaticTokenProvider) Token(_ context.Context, _ TokenRequestOptions) (Token, error) {
	return p.token, nil
}
