package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/clients"
	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/identity"
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
		if imsProvider, err := identity.NewInstanceMetadataProvider(&identity.InstanceMetadataProviderConfig{
			ServiceAccountEmail: email,
			HTTPClient:          cfg.HTTPClient,
		}); err == nil {
			providers = append(providers, imsProvider)
		}
	}

	// Try Workload Identity Federation - provider handles client cleanup internally
	if wifProvider, err := identity.NewWorkloadIdentityFederationProvider(&identity.WorkloadIdentityFederationProviderConfig{
		TokenURL:               cfg.TokenCustomUrl,
		ClientID:               cfg.ServiceAccountEmail,
		FederatedTokenFunction: cfg.ServiceAccountFederatedTokenFunc,
		HTTPClient:             cfg.HTTPClient,
		Scopes:                 cfg.Scopes,
		Resources:              cfg.Resources,
	}); err == nil {
		providers = append(providers, wifProvider)
	}

	// Try Service Account Key - provider handles client cleanup internally and credentials file resolution
	if keyProvider, err := identity.NewServiceAccountKeyProvider(&identity.ServiceAccountKeyProviderConfig{
		ServiceAccountKey:   cfg.ServiceAccountKey,
		PrivateKey:          cfg.PrivateKey,
		TokenURL:            cfg.TokenCustomUrl,
		HTTPClient:          cfg.HTTPClient,
		CredentialsFilePath: &cfg.CredentialsFilePath,
		Scopes:              cfg.Scopes,
		Resources:           cfg.Resources,
	}); err == nil {
		providers = append(providers, keyProvider)
	}

	// Try Static Token - provider handles env var resolution and credentials file resolution
	if staticProvider, err := identity.NewStaticTokenProvider(&identity.StaticTokenProviderConfig{
		Token:               cfg.Token,
		CredentialsFilePath: &cfg.CredentialsFilePath,
	}); err == nil {
		providers = append(providers, staticProvider)
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
	if cfg == nil {
		cfg = &config.Configuration{}
	}

	token, err := getTokenWithFallback(cfg)
	if err != nil {
		return nil, fmt.Errorf("resolving token: %w", err)
	}
	if token == "" {
		return nil, fmt.Errorf("STACKIT_SERVICE_ACCOUNT_TOKEN not set and no token found in credentials file")
	}

	provider, err := identity.NewStaticTokenProvider(&identity.StaticTokenProviderConfig{
		Token:               token,
		CredentialsFilePath: &cfg.CredentialsFilePath,
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

	provider, err := identity.NewServiceAccountKeyProvider(&identity.ServiceAccountKeyProviderConfig{
		ServiceAccountKey:   cfg.ServiceAccountKey,
		PrivateKey:          cfg.PrivateKey,
		TokenURL:            cfg.TokenCustomUrl,
		HTTPClient:          cfg.HTTPClient,
		CredentialsFilePath: &cfg.CredentialsFilePath,
	})
	if err != nil {
		return nil, fmt.Errorf("error initializing client: %w", err)
	}

	return newTokenProviderRoundTripper(provider, getTransportFromConfig(cfg)), nil
}

// WorkloadIdentityFederationAuth configures the wif flow and returns an http.RoundTripper
// that can be used to make authenticated requests using an access token
func WorkloadIdentityFederationAuth(cfg *config.Configuration) (http.RoundTripper, error) {
	provider, err := identity.NewWorkloadIdentityFederationProvider(&identity.WorkloadIdentityFederationProviderConfig{
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

// getServiceAccountEmail searches for an email in the following order: client configuration, environment variable, credentials file.
// is not required for authentication, so it can be empty.
func getServiceAccountEmail(cfg *config.Configuration) string {
	if cfg.ServiceAccountEmail != "" {
		return cfg.ServiceAccountEmail
	}

	email, emailSet := os.LookupEnv(identity.EnvServiceAccountEmail)
	if !emailSet || email == "" {
		credentials, err := identity.ReadCredentialsFile(cfg.CredentialsFilePath)
		if err != nil {
			// email is not required for authentication, so it shouldnt block it
			return ""
		}
		return credentials.ServiceAccountEmail
	}
	return email
}

// getKey searches for a key in the following order: client configuration, environment variable, credentials file.
func getKey(cfgKey, cfgKeyPath *string, envVarKeyPath, envVarKey string, credTypePath, credTypeKey identity.CredentialType, cfgCredFilePath string) error {
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
			credentials, err := identity.ReadCredentialsFile(cfgCredFilePath)
			if err != nil {
				return fmt.Errorf("reading from credentials file: %w", err)
			}
			// read key path from credentials file
			keyPath, err = identity.ReadCredential(credTypePath, credentials)
			if err != nil || keyPath == "" {
				// key path was not provided, read key from credentials file
				key, err = identity.ReadCredential(credTypeKey, credentials)
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
	return getKey(&cfg.ServiceAccountKey, &cfg.ServiceAccountKeyPath, identity.EnvServiceAccountKeyPath, identity.EnvServiceAccountKey, identity.CredentialTypeServiceAccountKeyPath, identity.CredentialTypeServiceAccountKey, cfg.CredentialsFilePath)
}

// getPrivateKey configures the private key in the provided configuration
func getPrivateKey(cfg *config.Configuration) error {
	return getKey(&cfg.PrivateKey, &cfg.PrivateKeyPath, identity.EnvPrivateKeyPath, identity.EnvPrivateKey, identity.CredentialTypePrivateKeyPath, identity.CredentialTypePrivateKey, cfg.CredentialsFilePath)
}

// getTokenWithFallback resolves a token from config, environment, or credentials file
func getTokenWithFallback(cfg *config.Configuration) (string, error) {
	if cfg.Token != "" {
		return cfg.Token, nil
	}

	// Try environment variable
	if token, exists := os.LookupEnv(identity.EnvServiceAccountToken); exists && token != "" {
		return token, nil
	}

	// Try credentials file
	credentials, err := identity.ReadCredentialsFile(cfg.CredentialsFilePath)
	if err != nil {
		return "", fmt.Errorf("reading from credentials file: %w", err)
	}

	token, err := identity.ReadCredential(identity.CredentialTypeToken, credentials)
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
		if tokenURL, exists := os.LookupEnv(identity.EnvTokenBaseUrl); exists && tokenURL != "" {
			cfg.TokenCustomUrl = tokenURL
		}
	}

	return nil
}
