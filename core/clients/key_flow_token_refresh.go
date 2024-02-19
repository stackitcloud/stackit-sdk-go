package clients

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
)

var (
	defaultTimeStartBeforeTokenExpiration = 30 * time.Minute
	defaultTimeBetweenContextCheck        = time.Second
	defaultTimeBetweenTries               = 5 * time.Minute
)

// Continuously refreshes the token of a key flow, retrying if the token API returns 5xx errrors. Writes to stderr when it terminates.
//
// To terminate this routine, close the context keyFlow.config.TokenRefreshInBackgroundContext.
func refreshToken(keyflow *KeyFlow) {
	refresher := &tokenRefresher{
		keyFlow:                        keyflow,
		timeStartBeforeTokenExpiration: defaultTimeStartBeforeTokenExpiration,
		timeBetweenCtxCheck:            defaultTimeBetweenContextCheck,
		timeBetweenTries:               defaultTimeBetweenTries,
	}
	err := refresher.refreshToken()
	fmt.Fprintf(os.Stderr, "Token refreshing terminated: %v", err)
}

type tokenRefresher struct {
	keyFlow *KeyFlow
	// Token refresh tries start at [Access token expiration timestamp] - [This duration]
	timeStartBeforeTokenExpiration time.Duration
	timeBetweenCtxCheck            time.Duration
	timeBetweenTries               time.Duration
}

// Continuously refreshes the token of a key flow, retrying if the token API returns 5xx errrors. Always returns with a non-nil error.
//
// To terminate this routine, close the context tokenRefresher.keyFlow.config.TokenRefreshInBackgroundContext.
func (refresher *tokenRefresher) refreshToken() error {
	expirationTimestamp, err := refresher.getAccessTokenExpirationTimestamp()
	if err != nil {
		return fmt.Errorf("get access token expiration timestamp: %w", err)
	}
	startRefreshTimestamp := expirationTimestamp.Add(-refresher.timeStartBeforeTokenExpiration)

	for {
		err = refresher.waitUntilTimestamp(startRefreshTimestamp)
		if err != nil {
			return err
		}

		err := refresher.keyFlow.config.TokenRefreshInBackgroundContext.Err()
		if err != nil {
			return fmt.Errorf("check context: %w", err)
		}

		ok, err := refresher.refreshTokens()
		if err != nil {
			return fmt.Errorf("refresh tokens: %w", err)
		}
		if !ok {
			startRefreshTimestamp = startRefreshTimestamp.Add(refresher.timeBetweenTries)
			continue
		}

		expirationTimestamp, err := refresher.getAccessTokenExpirationTimestamp()
		if err != nil {
			return fmt.Errorf("get access token expiration timestamp: %w", err)
		}
		startRefreshTimestamp = expirationTimestamp.Add(-refresher.timeStartBeforeTokenExpiration)
	}
}

func (refresher *tokenRefresher) getAccessTokenExpirationTimestamp() (*time.Time, error) {
	token := refresher.keyFlow.token.AccessToken

	// We can safely use ParseUnverified because we are not doing authentication of any kind
	// We're just checking the expiration time
	tokenParsed, _, err := jwt.NewParser().ParseUnverified(token, &jwt.RegisteredClaims{})
	if err != nil {
		return nil, fmt.Errorf("parse access token: %w", err)
	}
	expirationTimestampNumeric, err := tokenParsed.Claims.GetExpirationTime()
	if err != nil {
		return nil, fmt.Errorf("get expiration timestamp from access token: %w", err)
	}
	return &expirationTimestampNumeric.Time, nil
}

func (refresher *tokenRefresher) waitUntilTimestamp(timestamp time.Time) error {
	for time.Now().Before(timestamp) {
		err := refresher.keyFlow.config.TokenRefreshInBackgroundContext.Err()
		if err != nil {
			return fmt.Errorf("check context: %w", err)
		}
		time.Sleep(refresher.timeBetweenCtxCheck)
	}
	return nil
}

// Returns:
//   - (true, nil) if successful.
//   - (false, nil) if not successful but should be retried.
//   - (_, err) if not successful and shouldn't be retried.
func (refresher *tokenRefresher) refreshTokens() (bool, error) {
	err := refresher.keyFlow.createAccessTokenWithRefreshToken()
	if err == nil {
		return true, nil
	}

	// Should be retired if this is an API error with status code non-5xx
	oapiErr := oapierror.GenericOpenAPIError{}
	if !errors.As(err, &oapiErr) {
		return false, err
	}
	if oapiErr.StatusCode < 500 {
		return false, err
	}
	return false, nil
}
