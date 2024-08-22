package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/observability"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	observabilityClient, err := observability.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Observability API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the observability instances for your project
	getInstanceResp, err := observabilityClient.ListInstances(context.Background(), projectId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[Observability API] Error when calling `GetInstances`: %v\n", err)
	} else {
		fmt.Printf("[Observability API] Number of instances: %v\n", len(*getInstanceResp.Instances))
	}

	// Create an observability instance
	createInstancePayload := observability.CreateInstancePayload{
		Name:   utils.Ptr("myInstance"),
		PlanId: utils.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	}

	createInstanceResp, err := observabilityClient.CreateInstance(context.Background(), projectId).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Observability API] Error when calling `CreateInstance`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[Observability API] Created instance with id \"%s\" and dashboard url \"%s\".\n", *createInstanceResp.InstanceId, *createInstanceResp.DashboardUrl)
}
