package clients

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	// Service Account Key Flow
	// Auth flow env variables
	ServiceAccountKey     = "STACKIT_SERVICE_ACCOUNT_KEY"
	PrivateKey            = "STACKIT_PRIVATE_KEY"
	ServiceAccountKeyPath = "STACKIT_SERVICE_ACCOUNT_KEY_PATH"
	PrivateKeyPath        = "STACKIT_PRIVATE_KEY_PATH"
	tokenAPI              = "https://service-account.api.stackit.cloud/token" //nolint:gosec // linter false positive
	jwksAPI               = "https://service-account.api.stackit.cloud/.well-known/jwks.json"
	defaultTokenType      = "Bearer"
	defaultScope          = ""
)

// KeyFlow handles auth with SA key
type KeyFlow struct {
	client        *http.Client
	config        *KeyFlowConfig
	doer          func(client *http.Client, req *http.Request, cfg *RetryConfig) (resp *http.Response, err error)
	key           *ServiceAccountKeyPrivateResponse
	privateKey    *rsa.PrivateKey
	privateKeyPEM []byte
	token         *TokenResponseBody
}

// KeyFlowConfig is the flow config
type KeyFlowConfig struct {
	ServiceAccountKey string
	PrivateKey        string
	ClientRetry       *RetryConfig
	TokenUrl          string
	JWKSUrl           string
}

// TokenResponseBody is the API response
// when requesting a new token
type TokenResponseBody struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

// ServiceAccountKeyPrivateResponse is the API response
// when creating a new SA key
type ServiceAccountKeyPrivateResponse struct {
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"createdAt"`
	Credentials struct {
		Aud        string    `json:"aud"`
		Iss        string    `json:"iss"`
		Kid        string    `json:"kid"`
		PrivateKey *string   `json:"privateKey,omitempty"`
		Sub        uuid.UUID `json:"sub"`
	} `json:"credentials"`
	ID           uuid.UUID  `json:"id"`
	KeyAlgorithm string     `json:"keyAlgorithm"`
	KeyOrigin    string     `json:"keyOrigin"`
	KeyType      string     `json:"keyType"`
	PublicKey    string     `json:"publicKey"`
	ValidUntil   *time.Time `json:"validUntil,omitempty"`
}

// GetConfig returns the flow configuration
func (c *KeyFlow) GetConfig() KeyFlowConfig {
	if c.config == nil {
		return KeyFlowConfig{}
	}
	return *c.config
}

// GetServiceAccountEmail returns the service account email
func (c *KeyFlow) GetServiceAccountEmail() string {
	if c.key == nil {
		return ""
	}
	return c.key.Credentials.Iss
}

// GetToken returns the token field
func (c *KeyFlow) GetToken() TokenResponseBody {
	if c.token == nil {
		return TokenResponseBody{}
	}
	return *c.token
}

func (c *KeyFlow) Init(cfg *KeyFlowConfig) error {
	c.token = &TokenResponseBody{}
	c.config = cfg
	c.doer = Do

	// set defaults if no custom token and jwks url are provided
	if c.config.TokenUrl == "" {
		c.config.TokenUrl = tokenAPI
	}
	if c.config.JWKSUrl == "" {
		c.config.JWKSUrl = jwksAPI
	}
	c.configureHTTPClient()
	if c.config.ClientRetry == nil {
		c.config.ClientRetry = NewRetryConfig()
	}
	return c.validate()
}

// SetToken can be used to set an access and refresh token manually in the client
// the other fields in the token field are determined by inspecting the token or setting default values
func (c *KeyFlow) SetToken(accessToken, refreshToken string) error {
	// We can safely use ParseUnverified because we are not authenticating the user,
	// We are parsing the token just to get the expiration time claim
	parsedAccessToken, _, err := jwt.NewParser().ParseUnverified(accessToken, &jwt.RegisteredClaims{})
	if err != nil {
		return fmt.Errorf("parse access token to read expiration time: %w", err)
	}
	exp, err := parsedAccessToken.Claims.GetExpirationTime()
	if err != nil {
		return fmt.Errorf("get expiration time from access token: %w", err)
	}

	c.token = &TokenResponseBody{
		AccessToken:  accessToken,
		ExpiresIn:    int(exp.Time.Unix()),
		Scope:        defaultScope,
		RefreshToken: refreshToken,
		TokenType:    defaultTokenType,
	}
	return nil
}

// Clone creates a clone of the client
func (c *KeyFlow) Clone() interface{} {
	sc := *c
	nc := &sc
	cl := *nc.client
	cf := *nc.config
	ke := *nc.key
	to := *nc.token
	nc.client = &cl
	nc.config = &cf
	nc.key = &ke
	nc.token = &to
	return c
}

// Roundtrip performs the request
func (c *KeyFlow) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.client == nil {
		return nil, fmt.Errorf("please run Init()")
	}

	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	return c.doer(c.client, req, c.config.ClientRetry)
}

// GetAccessToken returns a short-lived access token and saves the access and refresh tokens in the token field
func (c *KeyFlow) GetAccessToken() (string, error) {
	accessTokenIsValid, err := c.validateToken(c.token.AccessToken)
	if err != nil {
		return "", fmt.Errorf("failed initial validation: %w", err)
	}
	if accessTokenIsValid {
		return c.token.AccessToken, nil
	}
	if err := c.recreateAccessToken(); err != nil {
		return "", fmt.Errorf("failed during token recreation: %w", err)
	}
	return c.token.AccessToken, nil
}

// configureHTTPClient configures the HTTP client
func (c *KeyFlow) configureHTTPClient() {
	client := &http.Client{}
	client.Timeout = DefaultClientTimeout
	c.client = client
}

// validate the client is configured well
func (c *KeyFlow) validate() error {
	if c.config.ServiceAccountKey == "" {
		return fmt.Errorf("service account access key cannot be empty")
	}
	if c.config.PrivateKey == "" {
		return fmt.Errorf("private key cannot be empty")
	}

	c.key = &ServiceAccountKeyPrivateResponse{}
	err := json.Unmarshal([]byte(c.config.ServiceAccountKey), c.key)
	if err != nil {
		return fmt.Errorf("unmarshalling service account key: %w", err)
	}

	c.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(c.config.PrivateKey))
	if err != nil {
		return fmt.Errorf("parsing private key from PEM file: %w", err)
	}

	// Encode the private key in PEM format
	privKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(c.privateKey),
	}
	c.privateKeyPEM = pem.EncodeToMemory(privKeyPEM)

	return nil
}

// Flow auth functions

// recreateAccessToken is used to create a new access token
// when the existing one isn't valid anymore
func (c *KeyFlow) recreateAccessToken() error {
	refreshTokenIsValid, err := c.validateToken(c.token.RefreshToken)
	if err != nil {
		return err
	}
	if refreshTokenIsValid {
		return c.createAccessTokenWithRefreshToken()
	}
	return c.createAccessToken()
}

// createAccessToken creates an access token using self signed JWT
func (c *KeyFlow) createAccessToken() (err error) {
	grant := "urn:ietf:params:oauth:grant-type:jwt-bearer"
	assertion, err := c.generateSelfSignedJWT()
	if err != nil {
		return err
	}
	res, err := c.requestToken(grant, assertion)
	if err != nil {
		return err
	}
	defer func() {
		tempErr := res.Body.Close()
		if tempErr != nil {
			err = fmt.Errorf("closing request access token response: %w", tempErr)
		}
	}()
	return c.parseTokenResponse(res)
}

// createAccessTokenWithRefreshToken creates an access token using
// an existing pre-validated refresh token
func (c *KeyFlow) createAccessTokenWithRefreshToken() (err error) {
	res, err := c.requestToken("refresh_token", c.token.RefreshToken)
	if err != nil {
		return err
	}
	defer func() {
		tempErr := res.Body.Close()
		if tempErr != nil {
			err = fmt.Errorf("closing request access token with refresh token response: %w", tempErr)
		}
	}()
	return c.parseTokenResponse(res)
}

// generateSelfSignedJWT generates JWT token
func (c *KeyFlow) generateSelfSignedJWT() (string, error) {
	claims := jwt.MapClaims{
		"iss": c.key.Credentials.Iss,
		"sub": c.key.Credentials.Sub,
		"jti": uuid.New(),
		"aud": c.key.Credentials.Aud,
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	token.Header["kid"] = c.key.Credentials.Kid
	tokenString, err := token.SignedString(c.privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// requestToken makes a request to the SA token API
func (c *KeyFlow) requestToken(grant, assertion string) (*http.Response, error) {
	body := url.Values{}
	body.Set("grant_type", grant)
	body.Set("assertion", assertion)
	payload := strings.NewReader(body.Encode())
	req, err := http.NewRequest(http.MethodPost, c.config.TokenUrl, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return c.doer(&http.Client{}, req, c.config.ClientRetry)
}

// parseTokenResponse parses the response from the server
func (c *KeyFlow) parseTokenResponse(res *http.Response) error {
	if res == nil {
		return fmt.Errorf("received bad response from API")
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("received: %+v", res)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	c.token = &TokenResponseBody{}
	return json.Unmarshal(body, c.token)
}

// validateToken returns true if token is valid
func (c *KeyFlow) validateToken(token string) (bool, error) {
	if token == "" {
		return false, nil
	}
	if _, err := c.parseToken(token); err != nil {
		if strings.Contains(err.Error(), "401") {
			c.token = &TokenResponseBody{}
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// parseToken parses and validates a JWT token
func (c *KeyFlow) parseToken(token string) (*jwt.Token, error) {
	b, err := c.getJwksJSON(token)
	if err != nil {
		return nil, err
	}
	var jwksBytes = json.RawMessage(b)
	jwks, err := keyfunc.NewJSON(jwksBytes)
	if err != nil {
		return nil, err
	}

	return jwt.Parse(token, jwks.Keyfunc)
}

func (c *KeyFlow) getJwksJSON(token string) (jwks []byte, err error) {
	req, err := http.NewRequest("GET", c.config.JWKSUrl, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := c.doer(&http.Client{}, req, c.config.ClientRetry)
	if err != nil {
		return nil, err
	}
	defer func() {
		tempErr := res.Body.Close()
		if tempErr != nil {
			jwks = nil
			err = fmt.Errorf("closing get jwks response: %w", tempErr)
		}
	}()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("getting jwks return error status %s", res.Status)
	}
	return io.ReadAll(res.Body)
}
