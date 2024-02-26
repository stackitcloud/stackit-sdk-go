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
	// Deprecated: retry options were removed to reduce complexity of the client. If this functionality is needed, you can provide your own custom HTTP client.
	ClientRetry *RetryConfig
}

// GetConfig returns the flow configuration
func (c *NoAuthFlow) GetConfig() NoAuthFlowConfig {
	if c.config == nil {
		return NoAuthFlowConfig{}
	}
	return *c.config
}

func (c *NoAuthFlow) Init(_ NoAuthFlowConfig) error {
	c.config = &NoAuthFlowConfig{}
	c.client = &http.Client{
		Timeout: DefaultClientTimeout,
	}
	return nil
}

// Roundtrip performs the request
func (c *NoAuthFlow) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.client == nil {
		return nil, fmt.Errorf("please run Init()")
	}
	return c.client.Do(req)
}
