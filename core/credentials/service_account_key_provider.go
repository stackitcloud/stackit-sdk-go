package credentials

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	// nolint:gosec // G101 False positive: This is a constant URL, not a credential
	serviceAccountKeyDefaultTokenURL = "https://accounts.stackit.cloud/oauth/v2/token"
	// nolint:gosec // G101 False positive: This is a constant protocol string, not a credential
	serviceAccountJWTBearerGrantType = "urn:ietf:params:oauth:grant-type:jwt-bearer"
	serviceAccountKeyDefaultLeeway   = time.Minute
	serviceAccountKeyErrorPrefix     = "service account key provider"
)

// ServiceAccountKeyProviderConfig contains the configuration for ServiceAccountKeyProvider.
type ServiceAccountKeyProviderConfig struct {
	// ServiceAccountKey is the service account key as a JSON string.
	// If empty, ServiceAccountKeyPath is used, then the STACKIT_SERVICE_ACCOUNT_KEY
	// environment variable, then STACKIT_SERVICE_ACCOUNT_KEY_PATH, then the credentials file.
	ServiceAccountKey string
	// ServiceAccountKeyPath is the path of a file with the service account key.
	// It is used when ServiceAccountKey is empty.
	ServiceAccountKeyPath string
	// PrivateKey is the RSA private key as a PEM-encoded string.
	// If empty, PrivateKeyPath is used, then the STACKIT_PRIVATE_KEY environment variable,
	// then STACKIT_PRIVATE_KEY_PATH, then the private key embedded in the service account
	// key, then the credentials file.
	PrivateKey string
	// PrivateKeyPath is the path of a file with the PEM-encoded RSA private key.
	// It is used when PrivateKey is empty.
	PrivateKeyPath string
	// TokenURL overrides the token endpoint. If empty, the value from the service account key is used,
	// falling back to the default STACKIT token endpoint.
	TokenURL string
	// TokenRefreshLeeway controls how early before expiration the token is refreshed.
	// If zero, defaults to 1 minute.
	TokenRefreshLeeway time.Duration
	// HTTPClient is used for token requests. If nil, a default client with a 1-minute timeout is used.
	HTTPClient *http.Client
	// CredentialsFilePath overrides the credentials file location. If empty, the path from
	// STACKIT_CREDENTIALS_PATH is used, falling back to ~/.stackit/credentials.json.
	CredentialsFilePath string
	// Optional scopes to request for the access token
	Scopes []string
	// Optional resources to request for the access token
	Resources []string

	// credentialsFile is shared by NewDefaultProvider so the file is read once per chain.
	credentialsFile *credentialsFileSource
}

var _ TokenProvider = (*ServiceAccountKeyProvider)(nil)

// ServiceAccountKey retrieves tokens using STACKIT service account key flow.
type ServiceAccountKeyProvider struct {
	name       string
	httpClient *http.Client
	tokenURL   string
	json       *serviceAccountJSON
	privateKey *rsa.PrivateKey

	tokenLeeway time.Duration
	cache       tokenCache
	scopes      []string
	resources   []string
}

type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// NewServiceAccountKeyProvider creates a ServiceAccountKey provider.
// It resolves configuration from environment variables if not provided in the config.
func NewServiceAccountKeyProvider(cfg *ServiceAccountKeyProviderConfig) (*ServiceAccountKeyProvider, error) {
	if cfg == nil {
		return nil, fmt.Errorf("%s: config cannot be nil", serviceAccountKeyErrorPrefix)
	}
	file := fileSourceFor(cfg.credentialsFile, cfg.CredentialsFilePath)

	// Resolve service account key from config, env vars, or credentials file
	serviceAccountKeyJSON, err := resolveValue(cfg.ServiceAccountKey, cfg.ServiceAccountKeyPath, envServiceAccountKey, envServiceAccountKeyPath,
		fromFile(file, func(c *credentialsFile) string { return c.ServiceAccountKey }))
	if err != nil {
		return nil, fmt.Errorf("%s: read service account key from path: %w", serviceAccountKeyErrorPrefix, err)
	}
	if serviceAccountKeyJSON == "" {
		return nil, fmt.Errorf("%s: service account key not provided and not found in environment variables", serviceAccountKeyErrorPrefix)
	}

	// Parse service account key JSON
	var serviceAccountKey = &serviceAccountJSON{}
	if err := json.Unmarshal([]byte(serviceAccountKeyJSON), serviceAccountKey); err != nil {
		return nil, fmt.Errorf("%s: parse service account key JSON: %w", serviceAccountKeyErrorPrefix, err)
	}

	if serviceAccountKey.Credentials == nil {
		return nil, fmt.Errorf("%s: service account key credentials cannot be empty", serviceAccountKeyErrorPrefix)
	}

	// Resolve private key from config, env vars, embedded in service account key, or credentials file
	privateKeyPEM, err := resolveValue(cfg.PrivateKey, cfg.PrivateKeyPath, envPrivateKey, envPrivateKeyPath,
		func() string {
			if serviceAccountKey.Credentials.PrivateKey != nil {
				return *serviceAccountKey.Credentials.PrivateKey
			}
			return ""
		},
		fromFile(file, func(c *credentialsFile) string { return c.PrivateKey }))
	if err != nil {
		return nil, fmt.Errorf("%s: read private key from path: %w", serviceAccountKeyErrorPrefix, err)
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
		scopes:      cfg.Scopes,
		resources:   cfg.Resources,
	}, nil
}

// Token returns a valid access token.
func (p *ServiceAccountKeyProvider) Token(ctx context.Context, opt TokenRequestOptions) (Token, error) {
	scopes, resources := effectiveRequest(opt, p.scopes, p.resources)
	return p.cache.get(ctx, requestKey(scopes, resources), p.tokenLeeway, func(ctx context.Context) (Token, error) {
		fresh, err := p.requestToken(ctx, scopes, resources)
		if err != nil {
			return Token{}, err
		}
		debugContext(ctx, "credentials: authenticated", "provider", p.name)
		return fresh, nil
	})
}

func (p *ServiceAccountKeyProvider) requestToken(ctx context.Context, scopes, resources []string) (Token, error) {
	assertion, err := p.generateSelfSignedJWT()
	if err != nil {
		return Token{}, err
	}

	body := url.Values{}
	body.Set("grant_type", serviceAccountJWTBearerGrantType)
	body.Set("assertion", assertion)
	setScopesAndResources(body, scopes, resources)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.tokenURL, strings.NewReader(body.Encode()))
	if err != nil {
		return Token{}, fmt.Errorf("%s: create token request: %w", serviceAccountKeyErrorPrefix, err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := p.httpClient.Do(req)
	if err != nil {
		return Token{}, fmt.Errorf("%s: request access token: %w", serviceAccountKeyErrorPrefix, err)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		bodyRaw, _ := io.ReadAll(res.Body)
		errorContext(ctx, "credentials: token request failed", "status", res.StatusCode, "body", string(bodyRaw))
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
		TokenType:   tokenType(tokenResponse.TokenType),
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
