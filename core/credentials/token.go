package credentials

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Token is a normalized access token response.
type Token struct {
	AccessToken string
	// TokenType is the authorization scheme the token is presented with, taken from the
	// token_type of the OAuth2 response. It is "Bearer" for every flow STACKIT issues
	// today, but callers building their own Authorization header must use this rather
	// than assuming.
	// Providers whose source does not report a type default to "Bearer".
	TokenType string
	ExpiresOn time.Time
	RefreshOn time.Time
}

// getTokenExpiration extracts the expiration time from a JWT token.
// If expiresIn is provided (> 0), it's used to calculate expiration as Now + expiresIn seconds.
// Otherwise, the exp claim is parsed from the token.
func getTokenExpiration(accessToken string, expiresIn int) (time.Time, error) {
	if expiresIn > 0 {
		return time.Now().Add(time.Duration(expiresIn) * time.Second), nil
	}

	parsed, _, err := jwt.NewParser().ParseUnverified(accessToken, &jwt.RegisteredClaims{})
	if err != nil {
		return time.Time{}, fmt.Errorf("parse access token: %w", err)
	}
	exp, err := parsed.Claims.GetExpirationTime()
	if err != nil {
		return time.Time{}, fmt.Errorf("get token expiration: %w", err)
	}
	return exp.Time, nil
}

// tokenType normalizes the token_type of an OAuth2 response. Sources that do not report
// one are treated as defaultTokenType, which is what every STACKIT flow issues today.
func tokenType(reported string) string {
	if reported == "" {
		return defaultTokenType
	}
	return reported
}
