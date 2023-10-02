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

const (
	credentialsFilePath             = ".stackit/credentials.json" //nolint:gosec // linter false positive
	tokenCredentialType             = "token"
	serviceAccountKeyCredentialType = "serviceAccountKey"
	privateKeyCredentialType        = "privateKey"
)

type Credentials struct {
	STACKIT_SERVICE_ACCOUNT_EMAIL string
	STACKIT_SERVICE_ACCOUNT_TOKEN string
	STACKIT_SERVICE_ACCOUNT_KEY   string
	STACKIT_PRIVATE_KEY           string
}

// SetupAuth sets up authentication based on the configuration. The different options are
// custom authentication, no authentication, explicit token flow or default authentication
func SetupAuth(cfg *config.Configuration) (rt http.RoundTripper, err error) {
	if cfg == nil {
		cfg = &config.Configuration{}
		email := getServiceAccountEmail(cfg)
		cfg.ServiceAccountEmail = email
	}

	if cfg.CustomAuth != nil {
		return cfg.CustomAuth, nil
	} else if cfg.NoAuth {
		noAuthRoundTripper, err := NoAuth(cfg.RetryOptions)
		if err != nil {
			return nil, fmt.Errorf("configuring no auth client: %w", err)
		}
		return noAuthRoundTripper, nil
	} else if cfg.ServiceAccountKey != "" || cfg.ServiceAccountKeyPath != "" {
		keyRoundTripper, err := KeyAuth(cfg)
		if err != nil {
			return nil, fmt.Errorf("configuring key authentication: %w", err)
		}
		return keyRoundTripper, nil
	} else if cfg.Token != "" {
		tokenRoundTripper, err := TokenAuth(cfg)
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

	// Key flow
	rt, err = KeyAuth(cfg)
	if err != nil {
		keyFlowErr := err
		// Token flow
		rt, err = TokenAuth(cfg)
		if err != nil {
			return nil, fmt.Errorf("no valid credentials were found: trying key flow: %s, trying token flow: %s", keyFlowErr.Error(), err)
		}
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

// TokenAuth configures the token flow and returns an http.RoundTripper
// that can be used to make authenticated requests using a token
func TokenAuth(cfg *config.Configuration) (http.RoundTripper, error) {
	// Check token
	if cfg.Token == "" {
		token, tokenSet := os.LookupEnv("STACKIT_SERVICE_ACCOUNT_TOKEN")
		if !tokenSet || token == "" {
			credentials, err := readCredentialsFile(cfg.CredentialsFilePath)
			if err != nil {
				return nil, fmt.Errorf("reading from credentials file: %w", err)
			}
			token, err = readCredential(tokenCredentialType, credentials)
			if err != nil {
				return nil, fmt.Errorf("STACKIT_SERVICE_ACCOUNT_TOKEN not set. Trying to read from credentials file: %w", err)
			}
		}
		cfg.Token = token
	}

	tokenCfg := clients.TokenFlowConfig{
		ServiceAccountToken: cfg.Token,
		ClientRetry:         cfg.RetryOptions,
	}

	client := &clients.TokenFlow{}
	if err := client.Init(context.Background(), &tokenCfg); err != nil {
		return nil, fmt.Errorf("error initializing client: %w", err)
	}

	return client, nil
}

// KeyAuth configures the key flow and returns an http.RoundTripper
// that can be used to make authenticated requests using an access token
func KeyAuth(cfg *config.Configuration) (http.RoundTripper, error) {
	err := getServiceAccountKey(cfg)
	if err != nil {
		return nil, fmt.Errorf("configuring key authentication: service account key could not be found: %w", err)
	}

	err = getPrivateKey(cfg)
	if err != nil {
		return nil, fmt.Errorf("configuring key authentication: private key could not be found: %w", err)
	}

	keyCfg := clients.KeyFlowConfig{
		ServiceAccountKey: cfg.ServiceAccountKey,
		PrivateKey:        cfg.PrivateKey,
		ClientRetry:       cfg.RetryOptions,
		TokenUrl:          cfg.TokenCustomUrl,
		JWKSUrl:           cfg.JWKSCustomUrl,
	}

	client := &clients.KeyFlow{}
	if err := client.Init(&keyCfg); err != nil {
		return nil, fmt.Errorf("error initializing client: %w", err)
	}

	return client, nil
}

// readCredentialsFile reads the credentials file from the specified path and returns Credentials
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

// readCredential reads the specified credentialType from Credentials and returns it as a string
func readCredential(credentialType string, credentials *Credentials) (string, error) {
	var credentialValue string
	switch credentialType {
	case tokenCredentialType:
		credentialValue = credentials.STACKIT_SERVICE_ACCOUNT_TOKEN
		if credentialValue == "" {
			return credentialValue, fmt.Errorf("token is empty or not set")
		}
	case serviceAccountKeyCredentialType:
		credentialValue = credentials.STACKIT_SERVICE_ACCOUNT_KEY
		if credentialValue == "" {
			return credentialValue, fmt.Errorf("service account key is empty or not set")
		}
	case privateKeyCredentialType:
		credentialValue = credentials.STACKIT_PRIVATE_KEY
		if credentialValue == "" {
			return credentialValue, fmt.Errorf("private key is empty or not set")
		}
	default:
		return "", fmt.Errorf("invalid credential type: %s", credentialType)
	}

	return credentialValue, nil
}

// getServiceAccountEmail searches for an email in the following order: client configuration, environment variable, credentials file.
// is not required for authentication, so it can be empty.
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

// getPrivateKey searches for the private key in the following order: client configuration, environment variable, credentials file.
func getPrivateKey(cfg *config.Configuration) (err error) {
	var privateKeyFilePathFound bool
	// get key from configuration
	if cfg.PrivateKey == "" {
		// get key path from configuration
		if cfg.PrivateKeyPath == "" {
			// get key from environment
			privateKey, privateKeySet := os.LookupEnv("STACKIT_PRIVATE_KEY")
			if !privateKeySet || privateKey == "" {
				// get key path from environment
				privateKeyPath, privateKeyPathSet := os.LookupEnv("STACKIT_PRIVATE_KEY_PATH")
				if !privateKeyPathSet || privateKeyPath == "" {
					// get key from the credentials file
					credentials, err := readCredentialsFile(cfg.CredentialsFilePath)
					if err != nil {
						return fmt.Errorf("reading from credentials file: %w", err)
					}
					privateKey, err = readCredential(privateKeyCredentialType, credentials)
					if err != nil || privateKey == "" {
						return fmt.Errorf("neither key or path are provided in the configuration, as environment variable and not present in the credentials files: %w", err)
					}
					cfg.PrivateKey = privateKey
				}
				privateKeyFilePathFound = true
				cfg.PrivateKeyPath = privateKeyPath
			}
			cfg.PrivateKey = privateKey
		} else {
			privateKeyFilePathFound = true
		}
		if privateKeyFilePathFound && cfg.PrivateKey == "" {
			// Read from filepath
			privateKeyBytes, err := os.ReadFile(cfg.PrivateKeyPath)
			if err != nil {
				return err
			}
			cfg.PrivateKey = string(privateKeyBytes)
		}
	}
	return nil
}

// getServiceAccountKey searches for the service account key in the following order: client configuration, environment variable, credentials file.
func getServiceAccountKey(cfg *config.Configuration) (err error) {
	var serviceAccountKeyFilePathFound bool
	// get key from configuration
	if cfg.ServiceAccountKey == "" {
		// get key path from configuration
		if cfg.ServiceAccountKeyPath == "" {
			// get key from environment
			serviceAccountKey, serviceAccountKeySet := os.LookupEnv("STACKIT_SERVICE_ACCOUNT_KEY")
			if !serviceAccountKeySet || serviceAccountKey == "" {
				// get key path from environment
				serviceAccountKeyPath, serviceAccountKeyPathSet := os.LookupEnv("STACKIT_SERVICE_ACCOUNT_KEY_PATH")
				if !serviceAccountKeyPathSet || serviceAccountKeyPath == "" {
					// get key from the credentials file
					credentials, err := readCredentialsFile(cfg.CredentialsFilePath)
					if err != nil {
						return fmt.Errorf("reading from credentials file: %w", err)
					}
					serviceAccountKey, err = readCredential(serviceAccountKeyCredentialType, credentials)
					if err != nil || serviceAccountKey == "" {
						return fmt.Errorf("neither key or path are provided in the configuration, as environment variable and not present in the credentials files: %w", err)
					}
					cfg.ServiceAccountKey = serviceAccountKey
				}
				serviceAccountKeyFilePathFound = true
				cfg.ServiceAccountKeyPath = serviceAccountKeyPath
			}
			cfg.ServiceAccountKey = serviceAccountKey
		} else {
			serviceAccountKeyFilePathFound = true
		}
		if serviceAccountKeyFilePathFound && cfg.ServiceAccountKey == "" {
			// Read from filepath
			serviceAccountKeyBytes, err := os.ReadFile(cfg.ServiceAccountKeyPath)
			if err != nil {
				return err
			}
			cfg.ServiceAccountKey = string(serviceAccountKeyBytes)
		}
	}
	return nil
}
