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
	// TokenType is the authorization scheme the token is presented with, taken from the
	// token_type of the OAuth2 response. It is "Bearer" for every flow STACKIT issues
	// today, but callers building their own Authorization header must use this rather
	// than assuming.
	// Providers whose source does not report a type default to DefaultTokenType.
	TokenType string
	ExpiresOn time.Time
	RefreshOn time.Time
}
