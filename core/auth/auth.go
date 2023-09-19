package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/stackitcloud/stackit-sdk-go/core/clients"
	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

const credentialsFilePath = ".stackit/credentials.json" //nolint:gosec // linter false positive

type Credentials struct {
	STACKIT_SERVICE_ACCOUNT_EMAIL string
	STACKIT_SERVICE_ACCOUNT_TOKEN string
}

// SetupAuth sets up authentication based on the configuration. The different options are
// custom authentication, no authentication, explicit token flow or default authentication
func SetupAuth(cfg *config.Configuration) (rt http.RoundTripper, err error) {
	if cfg == nil {
		cfg = &config.Configuration{}
	}

	if cfg.CustomAuth != nil {
		return cfg.CustomAuth, nil
	} else if cfg.NoAuth {
		noAuthRoundTripper, err := NoAuth(cfg.RetryOptions)
		if err != nil {
			return nil, fmt.Errorf("configuring no auth client: %w", err)
		}
		return noAuthRoundTripper, nil
	} else if cfg.Token != "" {
		email := getServiceAccountEmail(cfg)
		cfg.ServiceAccountEmail = email
		tokenRoundTripper, err := TokenAuth(cfg.Token, cfg.RetryOptions)
		if err != nil {
			return nil, fmt.Errorf("configuring token authentication: %w", err)
		}
		return tokenRoundTripper, nil
	} else {
		authRoundTripper, err := DefaultAuth(cfg)
		if err != nil {
			return nil, fmt.Errorf("configuring default authentication: %w", err)
		}
		return authRoundTripper, nil
	}
}

// DefaultAuth will search for a valid token in several locations.
// It will first try to find the token in the STACKIT_SERVICE_ACCOUNT_TOKEN. If not present, it will
// check the credentials file located in STACKIT_CREDENTIALS_PATH, if specified, or in
// $HOME/.stackit/credentials.json as a fallback. If the token is found, the TokenAuth flow is used and DefaultAuth returns
// an http.RoundTripper that can be used to make authenticated requests.
// In case the token is not found, DefaultAuth fails.
func DefaultAuth(cfg *config.Configuration) (rt http.RoundTripper, err error) {
	if cfg == nil {
		cfg = &config.Configuration{}
	}

	var token string
	// Check token
	token, tokenSet := os.LookupEnv("STACKIT_SERVICE_ACCOUNT_TOKEN")
	if !tokenSet || token == "" {
		token, err = readTokenFromCredentialsFile(cfg.CredentialsFilePath)
		if err != nil {
			return nil, fmt.Errorf("STACKIT_SERVICE_ACCOUNT_TOKEN not set. Trying to read from credentials file: %w", err)
		}
	}

	// Get service account  email
	email := getServiceAccountEmail(cfg)
	cfg.ServiceAccountEmail = email

	rt, err = TokenAuth(token, cfg.RetryOptions)
	if err != nil {
		return nil, fmt.Errorf("using token flow: %w", err)
	}

	return rt, nil
}

// NoAuth configures a flow without authentication and returns an http.RoundTripper
// that can be used to make unauthenticated requests
func NoAuth(retryOptions *clients.RetryConfig) (rt http.RoundTripper, err error) {
	noAuthConfig := clients.NoAuthFlowConfig{
		ClientRetry: retryOptions,
	}
	noAuthRoundTripper := &clients.NoAuthFlow{}
	if err := noAuthRoundTripper.Init(noAuthConfig); err != nil {
		return nil, fmt.Errorf("initializing client: %w", err)
	}
	return noAuthRoundTripper, nil
}

func readTokenFromCredentialsFile(path string) (string, error) {
	credentials, err := readCredentialsFile(path)
	if err != nil {
		return "", fmt.Errorf("reading from file: %w", err)
	}

	token := credentials.STACKIT_SERVICE_ACCOUNT_TOKEN
	if token == "" {
		return token, fmt.Errorf("token is empty or not set")
	}

	return token, nil
}

func readCredentialsFile(path string) (*Credentials, error) {
	if path == "" {
		customPath, customPathSet := os.LookupEnv("STACKIT_CREDENTIALS_PATH")
		if !customPathSet || customPath == "" {
			path = credentialsFilePath
			home, err := os.UserHomeDir()
			if err != nil {
				return nil, fmt.Errorf("getting home directory: %w", err)
			}
			path = filepath.Join(home, path)
		} else {
			path = customPath
		}
	}

	credentialsRaw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}

	var credentials Credentials
	err = json.Unmarshal(credentialsRaw, &credentials)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling credentials: %w", err)
	}
	return &credentials, nil
}

// getServiceAccountEmail searches for an email in the following order: client configuration, environment variable, credentials file.
// is not required for authentication, so it can be empty.
// It may required for some operations in the terraform provider, though
func getServiceAccountEmail(cfg *config.Configuration) string {
	if cfg.ServiceAccountEmail != "" {
		return cfg.ServiceAccountEmail
	}

	email, emailSet := os.LookupEnv("STACKIT_SERVICE_ACCOUNT_EMAIL")
	if !emailSet || email == "" {
		credentials, err := readCredentialsFile(cfg.CredentialsFilePath)
		if err != nil {
			// email is not required for authentication, so it shouldnt block it
			return ""
		}
		return credentials.STACKIT_SERVICE_ACCOUNT_EMAIL
	}
	return email
}

// TokenAuth configures the token flow and returns an http.RoundTripper
// that can be used to make authenticated requests using a token
func TokenAuth(token string, retryOptions *clients.RetryConfig) (http.RoundTripper, error) {
	tokenCfg := clients.TokenFlowConfig{
		ServiceAccountToken: token,
		ClientRetry:         retryOptions,
	}

	client := &clients.TokenFlow{}
	if err := client.Init(context.Background(), &tokenCfg); err != nil {
		return nil, fmt.Errorf("error initializing client: %w", err)
	}

	return client, nil
}
