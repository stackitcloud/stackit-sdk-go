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

type credentialType string

type Credentials struct {
	STACKIT_SERVICE_ACCOUNT_EMAIL    string
	STACKIT_SERVICE_ACCOUNT_TOKEN    string
	STACKIT_SERVICE_ACCOUNT_KEY_PATH string
	STACKIT_PRIVATE_KEY_PATH         string
}

const (
	credentialsFilePath                                = ".stackit/credentials.json" //nolint:gosec // linter false positive
	tokenCredentialType                 credentialType = "token"
	serviceAccountKeyPathCredentialType credentialType = "service_account_key_path"
	privateKeyPathCredentialType        credentialType = "private_key_path"
)

// SetupAuth sets up authentication based on the configuration. The different options are
// custom authentication, no authentication, explicit key flow, explicit token flow or default authentication
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

// DefaultAuth will search for a valid service account key or token in several locations.
// It will first try to use the key flow, by looking into the variables STACKIT_SERVICE_ACCOUNT_KEY, STACKIT_SERVICE_ACCOUNT_KEY_PATH,
// STACKIT_PRIVATE_KEY and STACKIT_PRIVATE_KEY_PATH. If the keys cannot be retrieved, it will check the credentials file located in STACKIT_CREDENTIALS_PATH, if specified, or in
// $HOME/.stackit/credentials.json as a fallback. If the key are found and are valid, the KeyAuth flow is used.
// If the key flow cannot be used, it will try to find a token in the STACKIT_SERVICE_ACCOUNT_TOKEN. If not present, it will
// search in the credentials file. If the token is found, the TokenAuth flow is used.
// DefaultAuth returns an http.RoundTripper that can be used to make authenticated requests.
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
			return nil, fmt.Errorf("no valid credentials were found: trying key flow: %s, trying token flow: %w", keyFlowErr.Error(), err)
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
// The KeyFlow requires a service account key and a private key.
//
// If the private key is not provided explicitly, KeyAuth will check if there is one included
// in the service account key and use that one. An explicitly provided private key takes
// precedence over the one on the service account key.
func KeyAuth(cfg *config.Configuration) (http.RoundTripper, error) {
	err := getServiceAccountKey(cfg)
	if err != nil {
		return nil, fmt.Errorf("configuring key authentication: service account key could not be found: %w", err)
	}

	// Unmarshal service account key to check if private key is present
	var serviceAccountKey = &clients.ServiceAccountKeyResponse{}
	err = json.Unmarshal([]byte(cfg.ServiceAccountKey), serviceAccountKey)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling service account key: %w", err)
	}

	// Try to get private key from configuration, environment or credentials file
	err = getPrivateKey(cfg)
	if err != nil {
		// If the private key is not provided explicitly, try to extract private key from the service account key
		// and use it if present
		var extractedPrivateKey string
		if serviceAccountKey.Credentials != nil && serviceAccountKey.Credentials.PrivateKey != nil {
			extractedPrivateKey = *serviceAccountKey.Credentials.PrivateKey
		}
		if extractedPrivateKey == "" {
			return nil, fmt.Errorf("configuring key authentication: private key is not part of the service account key and could not be found: %w", err)
		}
		cfg.PrivateKey = extractedPrivateKey
	}

	if cfg.TokenCustomUrl == "" {
		tokenCustomUrl, tokenUrlSet := os.LookupEnv("STACKIT_TOKEN_BASEURL")
		if tokenUrlSet {
			cfg.TokenCustomUrl = tokenCustomUrl
		}
	}
	if cfg.JWKSCustomUrl == "" {
		jwksCustomUrl, jwksUrlSet := os.LookupEnv("STACKIT_JWKS_BASEURL")
		if jwksUrlSet {
			cfg.JWKSCustomUrl = jwksCustomUrl
		}
	}

	keyCfg := clients.KeyFlowConfig{
		ServiceAccountKey: serviceAccountKey,
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
		return nil, fmt.Errorf("unmaPrivateKeyrshalling credentials: %w", err)
	}
	return &credentials, nil
}

// readCredential reads the specified credentialType from Credentials and returns it as a string
func readCredential(cred credentialType, credentials *Credentials) (string, error) {
	var credentialValue string
	switch cred {
	case tokenCredentialType:
		credentialValue = credentials.STACKIT_SERVICE_ACCOUNT_TOKEN
		if credentialValue == "" {
			return credentialValue, fmt.Errorf("token is empty or not set")
		}
	case serviceAccountKeyPathCredentialType:
		credentialValue = credentials.STACKIT_SERVICE_ACCOUNT_KEY_PATH
		if credentialValue == "" {
			return credentialValue, fmt.Errorf("service account key path is empty or not set")
		}
	case privateKeyPathCredentialType:
		credentialValue = credentials.STACKIT_PRIVATE_KEY_PATH
		if credentialValue == "" {
			return credentialValue, fmt.Errorf("private key path is empty or not set")
		}
	default:
		return "", fmt.Errorf("invalid credential type: %s", cred)
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
	// get key from configuration
	if cfg.PrivateKey == "" {
		// get key path from configuration
		if cfg.PrivateKeyPath == "" {
			// get key path from environment
			privateKeyPath, privateKeyPathSet := os.LookupEnv("STACKIT_PRIVATE_KEY_PATH")
			if !privateKeyPathSet || privateKeyPath == "" {
				// get key path from the credentials file
				credentials, err := readCredentialsFile(cfg.CredentialsFilePath)
				if err != nil {
					return fmt.Errorf("reading from credentials file: %w", err)
				}
				privateKeyPath, err = readCredential(privateKeyPathCredentialType, credentials)
				if err != nil || privateKeyPath == "" {
					return fmt.Errorf("neither key or path are provided in the configuration, as environment variable and not present in the credentials files: %w", err)
				}
			}
			cfg.PrivateKeyPath = privateKeyPath
		}
		// Read from filepath
		privateKeyBytes, err := os.ReadFile(cfg.PrivateKeyPath)
		if err != nil {
			return err
		}
		if len(privateKeyBytes) == 0 {
			return fmt.Errorf("key path points to an empty file")
		}
		cfg.PrivateKey = string(privateKeyBytes)
	}
	return nil
}

// getServiceAccountKey searches for the service account key in the following order: client configuration, environment variable, credentials file.
func getServiceAccountKey(cfg *config.Configuration) (err error) {
	// get key from configuration
	if cfg.ServiceAccountKey == "" {
		// get key path from configuration
		if cfg.ServiceAccountKeyPath == "" {
			// get key path from environment
			serviceAccountKeyPath, serviceAccountKeyPathSet := os.LookupEnv("STACKIT_SERVICE_ACCOUNT_KEY_PATH")
			if !serviceAccountKeyPathSet || serviceAccountKeyPath == "" {
				// get key path from the credentials file
				credentials, err := readCredentialsFile(cfg.CredentialsFilePath)
				if err != nil {
					return fmt.Errorf("reading from credentials file: %w", err)
				}
				serviceAccountKeyPath, err = readCredential(serviceAccountKeyPathCredentialType, credentials)
				if err != nil || serviceAccountKeyPath == "" {
					return fmt.Errorf("neither key or path are provided in the configuration, as environment variable and not present in the credentials files: %w", err)
				}
			}
			cfg.ServiceAccountKeyPath = serviceAccountKeyPath
		}
		// Read from filepath
		serviceAccountKeyBytes, err := os.ReadFile(cfg.ServiceAccountKeyPath)
		if err != nil {
			return err
		}
		cfg.ServiceAccountKey = string(serviceAccountKeyBytes)
	}
	return nil
}
