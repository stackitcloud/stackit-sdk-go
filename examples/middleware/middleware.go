package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/argus"
)

// RequestCapturer is a middleware that captures the request and appends it to a variable.
func RequestCapturer(status string) config.Middleware {
	return func(rt http.RoundTripper) http.RoundTripper {
		return &roundTripperWithCapture{rt, status}
	}
}

// roundTripperWithCapture wraps an existing RoundTripper and captures requests.
type roundTripperWithCapture struct {
	transport http.RoundTripper
	status    string
}

func (rt roundTripperWithCapture) RoundTrip(req *http.Request) (*http.Response, error) {
	returnStr := fmt.Sprintf("%s %s", req.Method, req.URL.String())

	fmt.Println("Captured request:", returnStr)
	fmt.Println("Status:", rt.status)

	// Proceed with the original round trip
	return rt.transport.RoundTrip(req)
}

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication
	// Use the `WithMiddleware` option to add the middleware to the client
	// The middleware will be executed in the reverse order they are added
	// In this case, the middleware with status "1" will be executed first
	argusClient, err := argus.NewAPIClient(
		config.WithRegion("eu01"),
		config.WithMiddleware(RequestCapturer("2")),
		config.WithMiddleware(RequestCapturer("1")),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Argus API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the argus instances for your project
	getInstanceResp, err := argusClient.ListInstances(context.Background(), projectId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[Argus API] Error when calling `GetInstances`: %v\n", err)
	} else {
		fmt.Printf("[Argus API] Number of instances: %v\n", len(*getInstanceResp.Instances))
	}
}
