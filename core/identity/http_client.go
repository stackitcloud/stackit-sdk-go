package identity

import (
	"net/http"
	"time"
)

// authHTTPClient creates a new http.Client with only the transport from the given client.
// This avoids passing a client that might already have auth roundtrippers attached,
// which would cause deadlocks during token acquisition.
// If timeout > 0, it will be applied to the returned client.
func authHTTPClient(httpClient *http.Client, timeout time.Duration) *http.Client {
	var transport http.RoundTripper

	// Extract transport from the provided client if available
	if httpClient != nil && httpClient.Transport != nil {
		transport = httpClient.Transport
	}

	// Create a new clean client with just the transport
	client := &http.Client{Transport: transport}

	// Apply timeout if specified
	if timeout > 0 {
		client.Timeout = timeout
	}

	return client
}
