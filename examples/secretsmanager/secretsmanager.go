package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	secretsmanager "github.com/stackitcloud/stackit-sdk-go/services/secretsmanager/v1api"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project

	// Create a new API client, that uses default authentication and configuration
	secretsmanagerClient, err := secretsmanager.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the secrets manager instances for your project
	getInstancesResp, err := secretsmanagerClient.DefaultAPI.ListInstances(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
	} else {
		fmt.Printf("Number of instances: %v\n", len(getInstancesResp.Instances))
	}

	// Create a secrets manager instance
	createInstancePayload := secretsmanager.CreateInstancePayload{
		Name: "my-secrets-manager",
	}
	createInstanceResp, err := secretsmanagerClient.DefaultAPI.CreateInstance(context.Background(), projectId).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
	} else {
		fmt.Printf("Created instance with instance id \"%s\".\n", createInstanceResp.Id)
	}
}
