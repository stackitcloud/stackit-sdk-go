package identity

import (
	"context"
	"time"
)

// TokenProvider exposes the minimal contract for retrieving an access token.
type TokenProvider interface {
	Token(ctx context.Context, options TokenRequestOptions) (Token, error)
}

// TokenRequestOptions carries optional token request hints for providers.
//
// The initial implementation keeps these fields as forward-compatible extension
// points. Providers may ignore them when not applicable.
type TokenRequestOptions struct {
	// Scopes is a list of OAuth2 scopes to request.
	// This overrides any default scopes configured in the provider.
	Scopes []string
	// Resources is a list of resource identifiers to request access for.
	// This overrides any default scopes configured in the provider.
	Resources []string
}

// Token is a normalized access token response.
type Token struct {
	AccessToken string
	ExpiresOn   time.Time
	RefreshOn   time.Time
}
