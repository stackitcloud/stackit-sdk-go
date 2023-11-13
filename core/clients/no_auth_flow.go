package clients

import (
	"fmt"
	"net/http"
)

type NoAuthFlow struct {
	client *http.Client
	config *NoAuthFlowConfig
}

// NoAuthFlowConfig holds the configuration for the unauthenticated flow
type NoAuthFlowConfig struct {
	ClientRetry *RetryConfig
}

// GetConfig returns the flow configuration
func (c *NoAuthFlow) GetConfig() NoAuthFlowConfig {
	if c.config == nil {
		return NoAuthFlowConfig{}
	}
	return *c.config
}

func (c *NoAuthFlow) Init(cfg NoAuthFlowConfig) error {
	c.config = &NoAuthFlowConfig{
		ClientRetry: cfg.ClientRetry,
	}
	c.client = &http.Client{
		Timeout: DefaultClientTimeout,
	}
	if c.config.ClientRetry == nil {
		c.config.ClientRetry = NewRetryConfig()
	}
	return nil
}

// Clone creates a clone of the client
func (c *NoAuthFlow) Clone() interface{} {
	sc := *c
	nc := &sc
	cl := *nc.client
	cf := *nc.config
	nc.client = &cl
	nc.config = &cf
	return c
}

// Roundtrip performs the request
func (c *NoAuthFlow) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.client == nil {
		return nil, fmt.Errorf("please run Init()")
	}
	return Do(c.client, req, c.config.ClientRetry)
}
