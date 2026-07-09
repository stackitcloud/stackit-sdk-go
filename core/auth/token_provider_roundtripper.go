package auth

import (
	"fmt"
	"net/http"

	"github.com/stackitcloud/stackit-sdk-go/core/identity"
)

const authorizationSchemeBearer = "Bearer"

type tokenProviderRoundTripper struct {
	provider identity.TokenProvider
	rt       http.RoundTripper
}

func newTokenProviderRoundTripper(provider identity.TokenProvider, rt http.RoundTripper) http.RoundTripper {
	if rt == nil {
		rt = http.DefaultTransport
	}
	return &tokenProviderRoundTripper{
		provider: provider,
		rt:       rt,
	}
}

func (r *tokenProviderRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if r.provider == nil {
		return nil, fmt.Errorf("token provider cannot be nil")
	}

	token, err := r.provider.Token(req.Context(), identity.TokenRequestOptions{})
	if err != nil {
		return nil, fmt.Errorf("get access token from provider: %w", err)
	}
	if token.AccessToken == "" {
		return nil, fmt.Errorf("token provider returned empty access token")
	}

	requestCopy := req.Clone(req.Context())
	requestCopy.Header.Set("Authorization", fmt.Sprintf("%s %s", authorizationSchemeBearer, token.AccessToken))
	return r.rt.RoundTrip(requestCopy)
}
