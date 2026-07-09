package identity

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	serviceAccountKeyDefaultTokenURL = "https://accounts.stackit.cloud/oauth/v2/token"
	serviceAccountJWTBearerGrantType = "urn:ietf:params:oauth:grant-type:jwt-bearer"
	serviceAccountKeyDefaultLeeway   = time.Minute
	serviceAccountKeyErrorPrefix     = "service account key provider"
)

// ServiceAccountKeyProviderConfig contains the configuration for ServiceAccountKeyProvider.
type ServiceAccountKeyProviderConfig struct {
	// ServiceAccountKey is the service account key as a JSON string.
	// If empty, attempts to read from STACKIT_SERVICE_ACCOUNT_KEY environment variable,
	// then from STACKIT_SERVICE_ACCOUNT_KEY_PATH, then optionally from credentials file if CredentialsFilePath is set.
	ServiceAccountKey string
	// PrivateKey is the RSA private key as a PEM-encoded string.
	// If empty, attempts to read from STACKIT_PRIVATE_KEY environment variable,
	// then from STACKIT_PRIVATE_KEY_PATH, then optionally from credentials file if CredentialsFilePath is set.
	PrivateKey string
	// TokenURL overrides the token endpoint. If empty, the value from the service account key is used,
	// falling back to the default STACKIT token endpoint.
	TokenURL string
	// TokenRefreshLeeway controls how early before expiration the token is refreshed.
	// If zero, defaults to 1 minute.
	TokenRefreshLeeway time.Duration
	// HTTPClient is used for token requests. If nil, a default client with a 1-minute timeout is used.
	HTTPClient *http.Client
	// CredentialsFilePath is a pointer to the path to the credentials file. When set (not nil), it enables
	// automatic resolution of key/private key from the credentials file if they are not found in config or
	// environment variables. If the pointer is nil, credentials file resolution is skipped.
	// When set to a pointer to an empty string, uses the default credentials file path (~/.stackit/credentials.json).
	CredentialsFilePath *string
	// Optional scopes to request for the access token
	Scopes []string
	// Optional resources to request for the access token
	Resources []string
}

var _ TokenProvider = (*ServiceAccountKeyProvider)(nil)

// ServiceAccountKey retrieves tokens using STACKIT service account key flow.
type ServiceAccountKeyProvider struct {
	name       string
	httpClient *http.Client
	tokenURL   string
	json       *ServiceAccountJson
	privateKey *rsa.PrivateKey

	tokenLeeway time.Duration
	tokenMutex  sync.RWMutex
	token       Token
	scopes      string
	resources   []string
}

type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// NewServiceAccountKeyProvider creates a ServiceAccountKey provider.
// It resolves configuration from environment variables if not provided in the config.
func NewServiceAccountKeyProvider(cfg ServiceAccountKeyProviderConfig) (*ServiceAccountKeyProvider, error) {
	// Resolve service account key from config, env vars, or credentials file
	serviceAccountKeyJSON := cfg.ServiceAccountKey
	if serviceAccountKeyJSON == "" {
		// Try STACKIT_SERVICE_ACCOUNT_KEY env var
		if key, exists := os.LookupEnv(EnvServiceAccountKey); exists && key != "" {
			serviceAccountKeyJSON = key
		} else if keyPath, exists := os.LookupEnv(EnvServiceAccountKeyPath); exists && keyPath != "" {
			// Try STACKIT_SERVICE_ACCOUNT_KEY_PATH env var
			data, err := os.ReadFile(keyPath)
			if err != nil {
				return nil, fmt.Errorf("%s: read service account key from path: %w", serviceAccountKeyErrorPrefix, err)
			}
			serviceAccountKeyJSON = string(data)
		} else if cfg.CredentialsFilePath != nil {
			// Try credentials file if pointer is set (not nil)
			credentials, err := ReadCredentialsFile(*cfg.CredentialsFilePath)
			if err == nil {
				if key, err := ReadCredential(CredentialTypeServiceAccountKey, credentials); err == nil {
					serviceAccountKeyJSON = key
				}
			}
		}
	}

	if serviceAccountKeyJSON == "" {
		return nil, fmt.Errorf("%s: service account key not provided and not found in environment variables", serviceAccountKeyErrorPrefix)
	}

	// Parse service account key JSON
	var serviceAccountKey = &ServiceAccountJson{}
	if err := json.Unmarshal([]byte(serviceAccountKeyJSON), serviceAccountKey); err != nil {
		return nil, fmt.Errorf("%s: parse service account key JSON: %w", serviceAccountKeyErrorPrefix, err)
	}

	if serviceAccountKey.Credentials == nil {
		return nil, fmt.Errorf("%s: service account key credentials cannot be empty", serviceAccountKeyErrorPrefix)
	}

	// Resolve private key from config, env vars, embedded in service account key, or credentials file
	privateKeyPEM := cfg.PrivateKey
	if privateKeyPEM == "" {
		// Try STACKIT_PRIVATE_KEY env var
		if key, exists := os.LookupEnv(EnvPrivateKey); exists && key != "" {
			privateKeyPEM = key
		} else if keyPath, exists := os.LookupEnv(EnvPrivateKeyPath); exists && keyPath != "" {
			// Try STACKIT_PRIVATE_KEY_PATH env var
			data, err := os.ReadFile(keyPath)
			if err != nil {
				return nil, fmt.Errorf("%s: read private key from path: %w", serviceAccountKeyErrorPrefix, err)
			}
			privateKeyPEM = string(data)
		} else if serviceAccountKey.Credentials.PrivateKey != nil {
			// Try private key embedded in service account key
			privateKeyPEM = *serviceAccountKey.Credentials.PrivateKey
		} else if cfg.CredentialsFilePath != nil {
			// Try credentials file if pointer is set and not found elsewhere
			credentials, err := ReadCredentialsFile(*cfg.CredentialsFilePath)
			if err == nil {
				if key, err := ReadCredential(CredentialTypePrivateKey, credentials); err == nil {
					privateKeyPEM = key
				}
			}
		}
	}

	if privateKeyPEM == "" {
		return nil, fmt.Errorf("%s: private key cannot be empty", serviceAccountKeyErrorPrefix)
	}

	// Parse RSA private key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyPEM))
	if err != nil {
		return nil, fmt.Errorf("%s: parse private key from PEM: %w", serviceAccountKeyErrorPrefix, err)
	}

	// Resolve token URL
	tokenURL := cfg.TokenURL
	if tokenURL == "" {
		tokenURL = serviceAccountKey.Credentials.TokenEndpoint
	}
	if tokenURL == "" {
		tokenURL = serviceAccountKeyDefaultTokenURL
	}

	// Resolve HTTP client - use clean client to avoid deadlocks if the provided client already has auth roundtrippers
	httpClient := authHTTPClient(cfg.HTTPClient, time.Minute)

	// Resolve token refresh leeway
	leeway := cfg.TokenRefreshLeeway
	if leeway == 0 {
		leeway = serviceAccountKeyDefaultLeeway
	}

	return &ServiceAccountKeyProvider{
		name:        "ServiceAccountKeyProvider",
		httpClient:  httpClient,
		tokenURL:    tokenURL,
		json:        serviceAccountKey,
		privateKey:  privateKey,
		tokenLeeway: leeway,
		scopes:      strings.Join(cfg.Scopes, " "),
		resources:   cfg.Resources,
	}, nil
}

// Token returns a valid access token.
func (p *ServiceAccountKeyProvider) Token(ctx context.Context, _ TokenRequestOptions) (Token, error) {
	if p == nil || p.httpClient == nil {
		return Token{}, fmt.Errorf("%s: provider is not initialized", serviceAccountKeyErrorPrefix)
	}
	if ctx == nil {
		ctx = context.Background()
	}

	now := time.Now().Add(p.tokenLeeway)
	p.tokenMutex.RLock()
	cached := p.token
	p.tokenMutex.RUnlock()
	if cached.AccessToken != "" && cached.RefreshOn.After(now) {
		return cached, nil
	}

	fresh, err := p.requestToken(ctx)
	if err != nil {
		return Token{}, err
	}
	DebugContext(ctx, "identity: authenticated", "provider", p.name)
	p.tokenMutex.Lock()
	p.token = fresh
	p.tokenMutex.Unlock()
	return fresh, nil
}

func (p *ServiceAccountKeyProvider) requestToken(ctx context.Context) (Token, error) {
	assertion, err := p.generateSelfSignedJWT()
	if err != nil {
		return Token{}, err
	}

	body := url.Values{}
	body.Set("grant_type", serviceAccountJWTBearerGrantType)
	body.Set("assertion", assertion)
	if p.scopes != "" {
		body.Set("scope", p.scopes)
	}
	for _, resource := range p.resources {
		body.Add("resource", resource)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.tokenURL, strings.NewReader(body.Encode()))
	if err != nil {
		return Token{}, fmt.Errorf("%s: create token request: %w", serviceAccountKeyErrorPrefix, err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := p.httpClient.Do(req)
	if err != nil {
		return Token{}, fmt.Errorf("%s: request access token: %w", serviceAccountKeyErrorPrefix, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bodyRaw, _ := io.ReadAll(res.Body)
		ErrorContext(ctx, "identity: token request failed", "status", res.StatusCode, "body", string(bodyRaw))
		return Token{}, fmt.Errorf("%s: token request failed with status %d: %s", serviceAccountKeyErrorPrefix, res.StatusCode, string(bodyRaw))
	}

	var tokenResponse oauthTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&tokenResponse); err != nil {
		return Token{}, fmt.Errorf("%s: decode token response: %w", serviceAccountKeyErrorPrefix, err)
	}
	if tokenResponse.AccessToken == "" {
		return Token{}, fmt.Errorf("%s: token response did not include access token", serviceAccountKeyErrorPrefix)
	}

	expiresOn, err := getTokenExpiration(tokenResponse.AccessToken, tokenResponse.ExpiresIn)
	if err != nil {
		return Token{}, fmt.Errorf("%s: %w", serviceAccountKeyErrorPrefix, err)
	}

	refreshOn := expiresOn.Add(-p.tokenLeeway)
	return Token{
		AccessToken: tokenResponse.AccessToken,
		ExpiresOn:   expiresOn,
		RefreshOn:   refreshOn,
	}, nil
}

func (p *ServiceAccountKeyProvider) generateSelfSignedJWT() (string, error) {
	claims := jwt.MapClaims{
		"iss": p.json.Credentials.Iss,
		"sub": p.json.Credentials.Sub,
		"jti": uuid.New(),
		"aud": p.json.Credentials.Aud,
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	token.Header["kid"] = p.json.Credentials.Kid

	tokenString, err := token.SignedString(p.privateKey)
	if err != nil {
		return "", fmt.Errorf("%s: sign self-signed jwt: %w", serviceAccountKeyErrorPrefix, err)
	}
	return tokenString, nil
}
