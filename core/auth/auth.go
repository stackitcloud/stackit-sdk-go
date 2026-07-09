package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/clients"
	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/identity"
)

type credentialType string

type Credentials struct {
	STACKIT_SERVICE_ACCOUNT_EMAIL    string // Deprecated: ServiceAccountEmail is not required and will be removed after 12th June 2025.
	STACKIT_SERVICE_ACCOUNT_TOKEN    string
	STACKIT_SERVICE_ACCOUNT_KEY_PATH string
	STACKIT_PRIVATE_KEY_PATH         string
	STACKIT_SERVICE_ACCOUNT_KEY      string
	STACKIT_PRIVATE_KEY              string
}

const (
	credentialsFilePath                                = ".stackit/credentials.json" //nolint:gosec // linter false positive
	tokenCredentialType                 credentialType = "token"
	serviceAccountKeyCredentialType     credentialType = "service_account_key"
	serviceAccountKeyPathCredentialType credentialType = "service_account_key_path"
	privateKeyCredentialType            credentialType = "private_key"
	privateKeyPathCredentialType        credentialType = "private_key_path"
)

var userHomeDir = os.UserHomeDir

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
		noAuthRoundTripper, err := NoAuth(cfg)
		if err != nil {
			return nil, fmt.Errorf("configuring no auth client: %w", err)
		}
		return noAuthRoundTripper, nil
	} else if cfg.WorkloadIdentityFederation {
		wifRoundTripper, err := WorkloadIdentityFederationAuth(cfg)
		if err != nil {
			return nil, fmt.Errorf("configuring workload identity federation client: %w", err)
		}
		return wifRoundTripper, nil
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
	}
	authRoundTripper, err := DefaultAuth(cfg)
	if err != nil {
		return nil, fmt.Errorf("configuring default authentication: %w", err)
	}
	return authRoundTripper, nil
}

// DefaultAuth will search for a valid service account key or token in several locations.
// It will first try to use the instance metadata service (IMS), which is available on STACKIT VMs.
// If IMS is not available, it will try the workload identity federation (WIF) flow.
// If WIF is not available, it will try to use the key flow, by looking into the variables STACKIT_SERVICE_ACCOUNT_KEY, STACKIT_SERVICE_ACCOUNT_KEY_PATH,
// STACKIT_PRIVATE_KEY and STACKIT_PRIVATE_KEY_PATH. If the keys cannot be retrieved, it will check the credentials file located in STACKIT_CREDENTIALS_PATH, if specified, or in
// $HOME/.stackit/credentials.json as a fallback. If the key are found and are valid, the KeyAuth flow is used.
// If the key flow cannot be used, it will try to find a token in the STACKIT_SERVICE_ACCOUNT_TOKEN. If not present, it will
// search in the credentials file. If the token is found, the TokenAuth flow is used.
// DefaultAuth returns an http.RoundTripper that can be used to make authenticated requests.
// In case no valid credentials were found, DefaultAuth fails.
func DefaultAuth(cfg *config.Configuration) (rt http.RoundTripper, err error) {
	if cfg == nil {
		cfg = &config.Configuration{}
	}

	providers := []identity.TokenProvider{}

	// Try Instance Metadata Service (IMS) with aggressive timeout to fail fast if not in cloud
	email := getServiceAccountEmail(cfg)
	if email != "" {
		imsClient := &http.Client{Timeout: 2 * time.Second}
		if imsProvider, err := identity.NewInstanceMetadataProvider(identity.InstanceMetadataProviderConfig{
			ServiceAccountEmail: email,
			HTTPClient:          imsClient,
		}); err == nil {
			providers = append(providers, imsProvider)
		}
	}

	// Try Workload Identity Federation
	if wifProvider, err := identity.NewWorkloadIdentityFederationProvider(identity.WorkloadIdentityFederationProviderConfig{
		TokenURL:               cfg.TokenCustomUrl,
		ClientID:               cfg.ServiceAccountEmail,
		FederatedTokenFunction: cfg.ServiceAccountFederatedTokenFunc,
		HTTPClient:             cfg.HTTPClient,
	}); err == nil {
		providers = append(providers, wifProvider)
	}

	// Try Service Account Key - provider handles env var resolution, fallback to credentials file
	if keyProvider, err := identity.NewServiceAccountKeyProvider(identity.ServiceAccountKeyProviderConfig{
		ServiceAccountKey: cfg.ServiceAccountKey,
		PrivateKey:        cfg.PrivateKey,
		TokenURL:          cfg.TokenCustomUrl,
		HTTPClient:        cfg.HTTPClient,
	}); err == nil {
		providers = append(providers, keyProvider)
	} else {
		// Try resolving from credentials file if direct/env resolution fails
		keyCfg := *cfg
		if getServiceAccountKeyAndPrivateKeyWithFallback(&keyCfg) == nil {
			if keyProvider, err := identity.NewServiceAccountKeyProvider(identity.ServiceAccountKeyProviderConfig{
				ServiceAccountKey: keyCfg.ServiceAccountKey,
				PrivateKey:        keyCfg.PrivateKey,
				TokenURL:          keyCfg.TokenCustomUrl,
				HTTPClient:        keyCfg.HTTPClient,
			}); err == nil {
				providers = append(providers, keyProvider)
			}
		}
	}

	// Try Static Token - provider handles env var resolution, fallback to credentials file
	if staticProvider, err := identity.NewStaticTokenProvider(identity.StaticTokenProviderConfig{
		Token: cfg.Token,
	}); err == nil {
		providers = append(providers, staticProvider)
	} else if token, err := getTokenWithFallback(cfg); err == nil && token != "" {
		// Try resolving from credentials file if env resolution fails
		if staticProvider, err := identity.NewStaticTokenProvider(identity.StaticTokenProviderConfig{
			Token: token,
		}); err == nil {
			providers = append(providers, staticProvider)
		}
	}

	chained, err := identity.NewChainedProvider(providers...)
	if err != nil {
		return nil, fmt.Errorf("error initializing chained provider: %w", err)
	}

	return newTokenProviderRoundTripper(chained, getTransportFromConfig(cfg)), nil
}

// NoAuth configures a flow without authentication and returns an http.RoundTripper
// that can be used to make unauthenticated requests
func NoAuth(cfgs ...*config.Configuration) (rt http.RoundTripper, err error) {
	noAuthConfig := clients.NoAuthFlowConfig{}
	noAuthRoundTripper := &clients.NoAuthFlow{}

	var cfg *config.Configuration

	if len(cfgs) > 0 {
		cfg = cfgs[0]
	} else {
		cfg = &config.Configuration{}
	}

	if cfg.HTTPClient != nil && cfg.HTTPClient.Transport != nil {
		noAuthConfig.HTTPTransport = cfg.HTTPClient.Transport
	}

	if err := noAuthRoundTripper.Init(noAuthConfig); err != nil {
		return nil, fmt.Errorf("initializing client: %w", err)
	}
	return noAuthRoundTripper, nil
}

// TokenAuth configures the token flow and returns an http.RoundTripper
// that can be used to make authenticated requests using a token
func TokenAuth(cfg *config.Configuration) (http.RoundTripper, error) {
	token, err := getTokenWithFallback(cfg)
	if err != nil {
		return nil, fmt.Errorf("resolving token: %w", err)
	}
	if token == "" {
		return nil, fmt.Errorf("STACKIT_SERVICE_ACCOUNT_TOKEN not set and no token found in credentials file")
	}

	provider, err := identity.NewStaticTokenProvider(identity.StaticTokenProviderConfig{
		Token: token,
	})
	if err != nil {
		return nil, fmt.Errorf("error initializing static token provider: %w", err)
	}

	return newTokenProviderRoundTripper(provider, getTransportFromConfig(cfg)), nil
}

// KeyAuth configures the key flow and returns an http.RoundTripper
// that can be used to make authenticated requests using an access token.
// The KeyFlow requires a service account key and a private key.
//
// Configuration is resolved in the following order:
// 1. Values from cfg parameter
// 2. Environment variables (STACKIT_SERVICE_ACCOUNT_KEY, STACKIT_PRIVATE_KEY, etc.)
// 3. Credentials file
// 4. Private key embedded in the service account key
func KeyAuth(cfg *config.Configuration) (http.RoundTripper, error) {
	if cfg == nil {
		cfg = &config.Configuration{}
	}

	if err := getServiceAccountKeyAndPrivateKeyWithFallback(cfg); err != nil {
		return nil, fmt.Errorf("configuring key authentication: %w", err)
	}

	provider, err := identity.NewServiceAccountKeyProvider(identity.ServiceAccountKeyProviderConfig{
		ServiceAccountKey: cfg.ServiceAccountKey,
		PrivateKey:        cfg.PrivateKey,
		TokenURL:          cfg.TokenCustomUrl,
		HTTPClient:        cfg.HTTPClient,
	})
	if err != nil {
		return nil, fmt.Errorf("error initializing client: %w", err)
	}

	return newTokenProviderRoundTripper(provider, getTransportFromConfig(cfg)), nil
}

// WorkloadIdentityFederationAuth configures the wif flow and returns an http.RoundTripper
// that can be used to make authenticated requests using an access token
func WorkloadIdentityFederationAuth(cfg *config.Configuration) (http.RoundTripper, error) {
	provider, err := identity.NewWorkloadIdentityFederationProvider(identity.WorkloadIdentityFederationProviderConfig{
		TokenURL:               cfg.TokenCustomUrl,
		ClientID:               cfg.ServiceAccountEmail,
		FederatedTokenFunction: cfg.ServiceAccountFederatedTokenFunc,
		HTTPClient:             cfg.HTTPClient,
	})
	if err != nil {
		return nil, fmt.Errorf("error initializing client: %w", err)
	}

	return newTokenProviderRoundTripper(provider, getTransportFromConfig(cfg)), nil
}

func getTransportFromConfig(cfg *config.Configuration) http.RoundTripper {
	if cfg != nil && cfg.HTTPClient != nil {
		return cfg.HTTPClient.Transport
	}
	return nil
}

// readCredentialsFile reads the credentials file from the specified path and returns Credentials
func readCredentialsFile(path string) (*Credentials, error) {
	if path == "" {
		customPath, customPathSet := os.LookupEnv("STACKIT_CREDENTIALS_PATH")
		if !customPathSet || customPath == "" {
			path = credentialsFilePath
			home, err := userHomeDir()
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

// getKey searches for a key in the following order: client configuration, environment variable, credentials file.
func getKey(cfgKey, cfgKeyPath *string, envVarKeyPath, envVarKey string, credTypePath, credTypeKey credentialType, cfgCredFilePath string) error {
	if *cfgKey != "" {
		return nil
	}
	if *cfgKeyPath == "" {
		// check environment variable: path
		keyPath, keyPathSet := os.LookupEnv(envVarKeyPath)
		// check environment variable: key
		key, keySet := os.LookupEnv(envVarKey)
		// if both are not set -> read from credentials file
		if (!keyPathSet || keyPath == "") && (!keySet || key == "") {
			credentials, err := readCredentialsFile(cfgCredFilePath)
			if err != nil {
				return fmt.Errorf("reading from credentials file: %w", err)
			}
			// read key path from credentials file
			keyPath, err = readCredential(credTypePath, credentials)
			if err != nil || keyPath == "" {
				// key path was not provided, read key from credentials file
				key, err = readCredential(credTypeKey, credentials)
				if err != nil || key == "" {
					return fmt.Errorf("neither key nor path is provided in the configuration, environment variable, or credentials file: %w", err)
				}
				*cfgKey = key
				return nil
			}
		} else if !keyPathSet || keyPath == "" {
			// key path was not provided, use key
			*cfgKey = key
			return nil
		}
		*cfgKeyPath = keyPath
	}
	keyBytes, err := os.ReadFile(*cfgKeyPath)
	if err != nil {
		return fmt.Errorf("reading key from file path: %w", err)
	}
	if len(keyBytes) == 0 {
		return fmt.Errorf("key path points to an empty file")
	}
	*cfgKey = string(keyBytes)
	return nil
}

// getServiceAccountKey configures the service account key in the provided configuration
func getServiceAccountKey(cfg *config.Configuration) error {
	return getKey(&cfg.ServiceAccountKey, &cfg.ServiceAccountKeyPath, "STACKIT_SERVICE_ACCOUNT_KEY_PATH", "STACKIT_SERVICE_ACCOUNT_KEY", serviceAccountKeyPathCredentialType, serviceAccountKeyCredentialType, cfg.CredentialsFilePath)
}

// getPrivateKey configures the private key in the provided configuration
func getPrivateKey(cfg *config.Configuration) error {
	return getKey(&cfg.PrivateKey, &cfg.PrivateKeyPath, "STACKIT_PRIVATE_KEY_PATH", "STACKIT_PRIVATE_KEY", privateKeyPathCredentialType, privateKeyCredentialType, cfg.CredentialsFilePath)
}

// getTokenWithFallback resolves a token from config, environment, or credentials file
func getTokenWithFallback(cfg *config.Configuration) (string, error) {
	if cfg.Token != "" {
		return cfg.Token, nil
	}

	// Try environment variable
	if token, exists := os.LookupEnv("STACKIT_SERVICE_ACCOUNT_TOKEN"); exists && token != "" {
		return token, nil
	}

	// Try credentials file
	credentials, err := readCredentialsFile(cfg.CredentialsFilePath)
	if err != nil {
		return "", fmt.Errorf("reading from credentials file: %w", err)
	}

	token, err := readCredential(tokenCredentialType, credentials)
	if err != nil {
		return "", fmt.Errorf("token not found in credentials file: %w", err)
	}

	return token, nil
}

// getServiceAccountKeyAndPrivateKeyWithFallback resolves service account key and private key with fallback to credentials file
// It also attempts to extract private key from the service account key JSON if not provided separately
func getServiceAccountKeyAndPrivateKeyWithFallback(cfg *config.Configuration) error {
	// Resolve service account key
	err := getServiceAccountKey(cfg)
	if err != nil {
		return fmt.Errorf("service account key could not be found: %w", err)
	}

	// Try to get private key
	err = getPrivateKey(cfg)
	if err != nil {
		// If the private key is not provided explicitly, try to extract it from the service account key JSON
		cfg.PrivateKey = ""
	}

	// If private key is still empty, try to extract from service account key JSON
	if cfg.PrivateKey == "" && cfg.ServiceAccountKey != "" {
		var serviceAccountKey = &identity.ServiceAccountJson{}
		if err := json.Unmarshal([]byte(cfg.ServiceAccountKey), serviceAccountKey); err == nil {
			if serviceAccountKey.Credentials != nil && serviceAccountKey.Credentials.PrivateKey != nil {
				cfg.PrivateKey = *serviceAccountKey.Credentials.PrivateKey
			}
		}
	}

	// Resolve token URL from env if not set
	if cfg.TokenCustomUrl == "" {
		if tokenURL, exists := os.LookupEnv("STACKIT_TOKEN_BASEURL"); exists && tokenURL != "" {
			cfg.TokenCustomUrl = tokenURL
		}
	}

	return nil
}
