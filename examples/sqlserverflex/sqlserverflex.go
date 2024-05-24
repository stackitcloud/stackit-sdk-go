package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex"
	"github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex/wait"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Specify instance configuration options
	flavorId := "FLAVOR_ID"
	version := "VERSION"

	ctx := context.Background()

	// Create a new API client, that uses default authentication and configuration
	sqlserverflexClient, err := sqlserverflex.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the MongoDB Flex instances for your project
	getInstancesResp, err := sqlserverflexClient.ListInstances(ctx, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
		os.Exit(1)
	}
	items := *getInstancesResp.Items
	fmt.Printf("Number of instances: %v\n", len(items))

	// Create an instance
	createInstancePayload := sqlserverflex.CreateInstancePayload{
		Name:     utils.Ptr("my-instance"),
		FlavorId: utils.Ptr(flavorId),
		Version:  utils.Ptr(version),
	}
	instance, err := sqlserverflexClient.CreateInstance(ctx, projectId).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating SQL Server Flex instance: %v\n", err)
	}
	instanceId := *instance.Id

	_, err = wait.CreateInstanceWaitHandler(ctx, sqlserverflexClient, projectId, instanceId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for SQL Server Flex instance creation: %v\n", err)
	}

	fmt.Printf("Created SQL Server Flex instance \"%s\".\n", instanceId)

	// Delete an instance
	err = sqlserverflexClient.DeleteInstance(ctx, projectId, instanceId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting SQL Server Flex instance: %v\n", err)
	}

	_, err = wait.DeleteInstanceWaitHandler(ctx, sqlserverflexClient, projectId, instanceId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for SQL Server Flex instance deletion: %v\n", err)
	}

	fmt.Printf("Deleted SQL Server Flex instance \"%s\".\n", instanceId)
}
