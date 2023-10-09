package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/argus"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	argusClient, err := argus.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Argus API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the argus instances for your project
	getInstanceResp, err := argusClient.GetInstances(context.Background(), projectId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[Argus API] Error when calling `GetInstances`: %v\n", err)
	} else {
		fmt.Printf("[Argus API] Number of instances: %v\n", len(*getInstanceResp.Instances))
	}

	// Create a argus Instance
	createInstancePayload := argus.CreateInstancePayload{
		Name:   utils.Ptr("myInstance"),
		PlanId: utils.Ptr("plan-uuid"),
	}

	createInstanceResp, err := argusClient.CreateInstance(context.Background(), projectId).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Argus API] Error when calling `CreateInstance`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[argus API] Created instance with id \"%s\" and dashboard url \"%s\".\n", *createInstanceResp.InstanceId, *createInstanceResp.DashboardUrl)
}
