package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/runtime"
	"github.com/stackitcloud/stackit-sdk-go/services/postgresflex"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	postgresflexClient, err := postgresflex.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the MongoDB Flex instances for your project and capture the HTTP response using the context
	var httpResp *http.Response
	ctxWithHTTPResp := runtime.WithCaptureHTTPResponse(context.Background(), &httpResp)
	getInstancesResp, err := postgresflexClient.ListInstances(ctxWithHTTPResp, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListInstances`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Number of instances: %v\n\n", len(*getInstancesResp.Items))
	fmt.Printf("HTTP response: %+v\n\n", *httpResp)

	// Get the MongoDB Flex instances for your project and capture the HTTP request using the context
	var httpReq *http.Request
	ctxWithHTTPReq := runtime.WithCaptureHTTPRequest(context.Background(), &httpReq)
	getInstancesResp, err = postgresflexClient.ListInstances(ctxWithHTTPReq, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListInstances`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Number of instances: %v\n\n", len(*getInstancesResp.Items))
	fmt.Printf("HTTP request: %+v\n\n", *httpReq)
}
