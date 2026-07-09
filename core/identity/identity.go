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
	Scopes []string
}

// Token is a normalized access token response.
type Token struct {
	AccessToken string
	ExpiresOn   time.Time
	RefreshOn   time.Time
}
