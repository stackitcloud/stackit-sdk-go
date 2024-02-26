package clients

import (
	"net/http"
	"time"
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
	DefaultClientTimeout = time.Minute
)

// Deprecated: retry options were removed to reduce complexity of the client. If this functionality is needed, you can provide your own custom HTTP client.
type RetryConfig struct {
	MaxRetries       int           // Max retries
	WaitBetweenCalls time.Duration // Time to wait between requests
	RetryTimeout     time.Duration // Max time to re-try
	ClientTimeout    time.Duration // HTTP Client timeout
}

// Deprecated: retry options were removed to reduce complexity of the client. If this functionality is needed, you can provide your own custom HTTP client.
func NewRetryConfig() *RetryConfig {
	return &RetryConfig{}
}

// Do performs the request
func Do(client *http.Client, req *http.Request) (resp *http.Response, err error) {
	if client == nil {
		client = http.DefaultClient
	}
	return client.Do(req)
}
