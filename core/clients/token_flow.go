package clients

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	// Service Account Token Flow
	// Auth flow env variables
	ServiceAccountToken = "STACKIT_SERVICE_ACCOUNT_TOKEN"
)

// TokenFlow handles auth with SA static token
type TokenFlow struct {
	client *http.Client
	config *TokenFlowConfig
}

// TokenFlowConfig is the flow config
type TokenFlowConfig struct {
	ServiceAccountEmail string
	ServiceAccountToken string
	ClientRetry         *RetryConfig
}

// GetConfig returns the flow configuration
func (c *TokenFlow) GetConfig() TokenFlowConfig {
	if c.config == nil {
		return TokenFlowConfig{}
	}
	return *c.config
}

func (c *TokenFlow) Init(ctx context.Context, cfg *TokenFlowConfig) error {
	c.config = cfg
	c.configureHTTPClient(ctx)
	if c.config.ClientRetry == nil {
		c.config.ClientRetry = NewRetryConfig()
	}
	return c.validate()
}

// Clone creates a clone of the client
func (c *TokenFlow) Clone() interface{} {
	sc := *c
	nc := &sc
	cl := *nc.client
	cf := *nc.config
	nc.client = &cl
	nc.config = &cf
	return c
}

// configureHTTPClient configures the HTTP client
func (c *TokenFlow) configureHTTPClient(ctx context.Context) {
	sts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.config.ServiceAccountToken},
	)
	o2nc := oauth2.NewClient(ctx, sts)
	o2nc.Timeout = DefaultClientTimeout
	c.client = o2nc
}

// validate the client is configured well
func (c *TokenFlow) validate() error {
	if c.config.ServiceAccountToken == "" {
		return fmt.Errorf("service account access token cannot be empty")
	}
	return nil
}

// Roundtrip performs the request
func (c *TokenFlow) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.client == nil {
		return nil, fmt.Errorf("please run Init()")
	}
	return do(c.client, req, c.config.ClientRetry)
}
