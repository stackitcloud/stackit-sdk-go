package identity

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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
