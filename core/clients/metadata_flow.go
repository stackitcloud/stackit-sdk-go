package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
)

const (
	defaultMetadataUrl = "http://169.254.169.254"
	// The metadata service issues tokens valid for an hour
	metadataTokenExpirationLeeway = 5 * time.Minute
)

var _ AuthFlow = &MetadataFlow{}

// MetadataFlow handles auth with the service account attached to the server,
// using the tokens its metadata service issues:
// https://docs.stackit.cloud/products/iaas-api/how-tos/use-service-accounts-via-the-iaas-api/
type MetadataFlow struct {
	rt             http.RoundTripper
	metadataClient *http.Client
	config         *MetadataFlowConfig

	tokenMutex sync.RWMutex
	token      *MetadataTokenResponseBody

	// If the current access token would expire in less than TokenExpirationLeeway,
	// the client will refresh it early to prevent clock skew or other timing issues.
	tokenExpirationLeeway time.Duration
}

// MetadataFlowConfig is the flow config
type MetadataFlowConfig struct {
	ServiceAccountEmail           string
	MetadataUrl                   string
	BackgroundTokenRefreshContext context.Context // Functionality is enabled if this isn't nil
	HTTPTransport                 http.RoundTripper
	MetadataHTTPClient            *http.Client
}

// MetadataTokenResponseBody is the metadata service response
// when requesting a token
type MetadataTokenResponseBody struct {
	Token      string    `json:"token"`
	ValidUntil time.Time `json:"validUntil"`
}

// GetConfig returns the flow configuration
func (c *MetadataFlow) GetConfig() MetadataFlowConfig {
	if c.config == nil {
		return MetadataFlowConfig{}
	}
	return *c.config
}

// GetAccessToken implements AuthFlow.
func (c *MetadataFlow) GetAccessToken() (string, error) {
	if c.rt == nil {
		return "", fmt.Errorf("nil http round tripper, please run Init()")
	}

	c.tokenMutex.RLock()
	token := c.token
	c.tokenMutex.RUnlock()

	if token != nil && time.Now().Add(c.tokenExpirationLeeway).Before(token.ValidUntil) {
		return token.Token, nil
	}
	if err := c.createAccessToken(); err != nil {
		return "", fmt.Errorf("get new access token: %w", err)
	}

	c.tokenMutex.RLock()
	defer c.tokenMutex.RUnlock()
	return c.token.Token, nil
}

func (c *MetadataFlow) refreshAccessToken() error {
	return c.createAccessToken()
}

// RoundTrip implements the http.RoundTripper interface.
// It gets a token, adds it to the request's authorization header, and performs the request.
func (c *MetadataFlow) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.rt == nil {
		return nil, fmt.Errorf("please run Init()")
	}

	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	return c.rt.RoundTrip(req)
}

// getBackgroundTokenRefreshContext implements AuthFlow.
func (c *MetadataFlow) getBackgroundTokenRefreshContext() context.Context {
	return c.config.BackgroundTokenRefreshContext
}

func (c *MetadataFlow) Init(cfg *MetadataFlowConfig) error {
	// No concurrency at this point, so no mutex check needed
	c.token = nil
	c.config = cfg

	if c.config.ServiceAccountEmail == "" {
		c.config.ServiceAccountEmail = utils.GetEnvOrDefault(clientIDEnv, "")
	}

	if c.config.MetadataUrl == "" {
		c.config.MetadataUrl = defaultMetadataUrl
	}

	c.tokenExpirationLeeway = metadataTokenExpirationLeeway

	if c.rt = cfg.HTTPTransport; c.rt == nil {
		c.rt = http.DefaultTransport
	}

	if c.metadataClient = cfg.MetadataHTTPClient; c.metadataClient == nil {
		c.metadataClient = &http.Client{
			// The metadata service is link-local, so it is never reached through a proxy
			Transport: &http.Transport{},
			Timeout:   DefaultClientTimeout,
		}
	}

	err := c.validate()
	if err != nil {
		return err
	}

	if c.config.BackgroundTokenRefreshContext != nil {
		go continuousRefreshToken(c)
	}
	return nil
}

func (c *MetadataFlow) validate() error {
	if c.config.ServiceAccountEmail == "" {
		return fmt.Errorf("service account email cannot be empty")
	}
	if _, err := url.ParseRequestURI(c.config.MetadataUrl); err != nil {
		return fmt.Errorf("parse metadata URL: %w", err)
	}
	if c.tokenExpirationLeeway < 0 {
		return fmt.Errorf("token expiration leeway cannot be negative")
	}

	return nil
}

func (c *MetadataFlow) createAccessToken() (err error) {
	res, err := c.requestToken()
	if err != nil {
		return err
	}
	defer func() {
		tempErr := res.Body.Close()
		if tempErr != nil && err == nil {
			err = fmt.Errorf("close request access token response: %w", tempErr)
		}
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		apiErr := &oapierror.GenericOpenAPIError{
			StatusCode: res.StatusCode,
			Body:       body,
		}
		if res.StatusCode == http.StatusNotFound {
			return fmt.Errorf("service account %s is not attached to this server: %w", c.config.ServiceAccountEmail, apiErr)
		}
		return apiErr
	}

	token := &MetadataTokenResponseBody{}
	if err := json.Unmarshal(body, token); err != nil {
		return fmt.Errorf("unmarshal token response: %w", err)
	}
	if token.Token == "" || token.ValidUntil.IsZero() {
		return fmt.Errorf("token response lacks token or validUntil")
	}

	c.tokenMutex.Lock()
	c.token = token
	c.tokenMutex.Unlock()
	return nil
}

func (c *MetadataFlow) requestToken() (*http.Response, error) {
	tokenUrl := strings.TrimSuffix(c.config.MetadataUrl, "/") +
		"/stackit/v1/service-accounts/" + url.PathEscape(c.config.ServiceAccountEmail) + "/token"
	req, err := http.NewRequest(http.MethodGet, tokenUrl, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	return c.metadataClient.Do(req)
}
