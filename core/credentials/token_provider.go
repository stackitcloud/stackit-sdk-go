package credentials

import "context"

// TokenProvider exposes the minimal contract for retrieving an access token.
//
// Providers cache one token per distinct combination of scopes and resources and refresh
// each independently; concurrent requests for the same combination share a single refresh.
type TokenProvider interface {
	Token(ctx context.Context, options TokenRequestOptions) (Token, error)
}

// TokenRequestOptions carries optional token request hints for providers.
//
// Providers that do not issue tokens per scope or resource ignore them.
type TokenRequestOptions struct {
	// Scopes is a list of OAuth2 scopes to request.
	// This overrides any default scopes configured in the provider.
	Scopes []string
	// Resources is a list of resource identifiers to request access for.
	// This overrides any default resources configured in the provider.
	Resources []string
}
