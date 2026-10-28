package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	objectstorage "github.com/stackitcloud/stackit-sdk-go/services/objectstorage/v2api"
)

/* This example shows how a middleware can be used to enable automatic
 * retries for HTTP 429 rate limits within the STACKIT Go SDK. To show
 * off that it works an amount of 100 goroutines is spawned which all try
 * to do the same SDK call in parallel.
 */

// RetryTransport wraps an underlying RoundTripper to retry on HTTP 429 rate limit errors with jitter.
type RetryTransport struct {
	Base        http.RoundTripper
	MaxRetries  int
	BaseBackoff time.Duration
	MaxJitter   time.Duration
}

func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	base := t.Base
	if base == nil {
		base = http.DefaultTransport
	}

	// Preserve request body for retries if present without mutating the original request
	getBody := req.GetBody
	if getBody == nil && req.Body != nil && req.Body != http.NoBody {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}

		err = req.Body.Close()
		if err != nil {
			return nil, err
		}

		getBody = func() (io.ReadCloser, error) {
			return io.NopCloser(bytes.NewReader(bodyBytes)), nil
		}
	}

	var resp *http.Response
	var err error

	for attempt := 0; attempt <= t.MaxRetries; attempt++ {
		reqClone := req.Clone(req.Context())
		if getBody != nil {
			var bodyErr error
			reqClone.Body, bodyErr = getBody()
			if bodyErr != nil {
				return nil, bodyErr
			}
		}

		resp, err = base.RoundTrip(reqClone)

		// If success or non-429 error, return immediately
		if err != nil || resp.StatusCode != http.StatusTooManyRequests {
			return resp, err
		}

		// Stop if max retries reached
		if attempt == t.MaxRetries {
			break
		}

		// Calculate base sleep duration (value of Retry-After Header, if header isn't present falls back to exponential backoff)
		wait := t.getWaitDuration(resp, attempt)

		// Always add random jitter regardless of Retry-After header presence. Else all
		// goroutines would try again in parallel after exactly the same interval.
		jitter := time.Duration(rand.Int64N(int64(t.MaxJitter))) //nolint:gosec // only used for jitter
		totalWait := wait + jitter

		// Drain and close response body before retrying to reuse TCP connections
		_, err = io.Copy(io.Discard, resp.Body)
		if err != nil {
			return nil, err
		}

		err = resp.Body.Close()
		if err != nil {
			return nil, err
		}

		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(totalWait):
		}
	}

	return resp, err
}

func (t *RetryTransport) getWaitDuration(resp *http.Response, attempt int) time.Duration {
	if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
		// Try parsing as integer seconds
		if seconds, err := strconv.Atoi(retryAfter); err == nil {
			return time.Duration(seconds) * time.Second
		}
		// Try parsing as HTTP-Date string
		if date, err := http.ParseTime(retryAfter); err == nil {
			if d := time.Until(date); d > 0 {
				return d
			}
		}
	}

	// Fallback to exponential backoff
	return t.BaseBackoff * (1 << attempt)
}

func NewRetryMiddleware(rt http.RoundTripper) http.RoundTripper {
	return &RetryTransport{
		Base:        rt,
		MaxRetries:  3,
		BaseBackoff: 10 * time.Second,
		MaxJitter:   500 * time.Millisecond, // Always added to wait time
	}
}

func main() {
	const (
		region = "eu01"
	)

	projectId, ok := os.LookupEnv("STACKIT_PROJECT_ID") // the uuid of your STACKIT project
	if !ok {
		fmt.Fprintf(os.Stderr, "STACKIT_PROJECT_ID env variable not set.\n")
		os.Exit(1)
	}

	ctx := context.Background()

	apiClientConfigOptions := []config.ConfigurationOption{
		config.WithMiddleware(NewRetryMiddleware),
	}

	// Create a new API client, that uses default authentication, add the retry implementation as a middleware.
	client, err := objectstorage.NewAPIClient(apiClientConfigOptions...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup

	for i := range 100 {
		wg.Add(1)

		go func(goroutineId int) {
			defer wg.Done()

			resp, err := client.DefaultAPI.ListBuckets(ctx, projectId, region).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Goroutine #%d failed: %v\n", goroutineId, err)
				return
			}

			fmt.Fprintf(os.Stderr, "Goroutine #%d completed (Amount of buckets: %d)\n", goroutineId, len(resp.Buckets))
		}(i)
	}

	wg.Wait()
	fmt.Fprintln(os.Stderr, "All goroutines finished.")
}
