package clients

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
)

const (
	// API configuration options:
	Environment = "STACKIT_ENV"
)

const (
	// Known error messages
	ClientTimeoutErr           = "Client.Timeout exceeded while awaiting headers"
	ClientContextDeadlineErr   = "context deadline exceeded"
	ClientConnectionRefusedErr = "connection refused"
	ClientEOFError             = "unexpected EOF"
)

const (
	DefaultClientTimeout         = time.Minute
	DefaultRetryMaxRetries       = 3
	DefaultRetryWaitBetweenCalls = 30 * time.Second
	DefaultRetryTimeout          = 2 * time.Minute
)

type RetryConfig struct {
	MaxRetries       int           // Max retries
	WaitBetweenCalls time.Duration // Time to wait between requests
	RetryTimeout     time.Duration // Max time to re-try
	ClientTimeout    time.Duration // HTTP Client timeout
}

func NewRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxRetries:       DefaultRetryMaxRetries,
		WaitBetweenCalls: DefaultRetryWaitBetweenCalls,
		RetryTimeout:     DefaultRetryTimeout,
		ClientTimeout:    DefaultClientTimeout,
	}
}

// Do performs the request, retrying according to the configurations provided
func Do(client *http.Client, req *http.Request, cfg *RetryConfig) (resp *http.Response, err error) {
	if cfg == nil {
		cfg = NewRetryConfig()
	}
	if client == nil {
		client = &http.Client{}
	}
	client.Timeout = cfg.ClientTimeout

	maxRetries := cfg.MaxRetries
	ctx, cancel := context.WithTimeout(context.Background(), cfg.RetryTimeout)
	defer cancel()
	b := backoff.WithContext(backoff.NewConstantBackOff(cfg.WaitBetweenCalls), ctx)

	err = backoff.Retry(func() error {
		resp, err = client.Do(req) //nolint:bodyclose // body is closed by the caller functions
		if err != nil {
			if maxRetries > 0 {
				if errorIsOneOf(err, ClientTimeoutErr, ClientContextDeadlineErr, ClientConnectionRefusedErr) ||
					(req.Method == http.MethodGet && strings.Contains(err.Error(), ClientEOFError)) {
					// reduce retries counter and retry
					maxRetries--
					return err
				}
			}
			return backoff.Permanent(err)
		}
		if maxRetries > 0 && resp != nil {
			if resp.StatusCode == http.StatusBadGateway ||
				resp.StatusCode == http.StatusGatewayTimeout {
				maxRetries--
				return fmt.Errorf("request returned a gateway error")
			}
		}
		return nil
	}, b)
	if err != nil {
		return resp, fmt.Errorf("url: %s\nmethod: %s\n err: %w", req.URL.String(), req.Method, err)
	}

	return resp, err
}

// ErrorIsOneOf checks if a given error message
// has one of the provided sub strings
func errorIsOneOf(err error, msgs ...string) bool {
	if err == nil {
		return false
	}
	for _, m := range msgs {
		if strings.Contains(err.Error(), m) {
			return true
		}
	}
	return false
}
