package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/mongodbflex"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	mongodbflexClient, err := mongodbflex.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the MongoDB Flex instances for your project and capture the HTTP response using the context
	var httpResp *http.Response
	ctxWithHTTPResp := config.WithCaptureHTTPResponse(context.Background(), &httpResp)
	getInstancesResp, err := mongodbflexClient.ListInstances(ctxWithHTTPResp, projectId).Tag("tag").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Number of instances: %v\n\n", len(*getInstancesResp.Items))
	fmt.Printf("HTTP response: %+v\n", *httpResp)
}
