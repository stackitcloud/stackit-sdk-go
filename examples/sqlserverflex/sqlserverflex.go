package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex"
	"github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex/wait"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Specify the region
	region := "REGION"

	// Specify instance configuration options
	flavorId := "FLAVOR_ID"
	version := "VERSION"

	ctx := context.Background()

	// Create a new API client, that uses default authentication and configuration
	sqlserverflexClient, err := sqlserverflex.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the SQLServer Flex instances for your project
	getInstancesResp, err := sqlserverflexClient.ListInstances(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListInstances`: %v\n", err)
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
	instance, err := sqlserverflexClient.CreateInstance(ctx, projectId, region).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating SQL Server Flex instance: %v\n", err)
	}
	instanceId := *instance.Id

	_, err = wait.CreateInstanceWaitHandler(ctx, sqlserverflexClient, projectId, instanceId, region).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for SQL Server Flex instance creation: %v\n", err)
	}

	fmt.Printf("Created SQL Server Flex instance \"%s\".\n", instanceId)

	// Delete an instance
	err = sqlserverflexClient.DeleteInstance(ctx, projectId, instanceId, region).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting SQL Server Flex instance: %v\n", err)
	}

	_, err = wait.DeleteInstanceWaitHandler(ctx, sqlserverflexClient, projectId, instanceId, region).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for SQL Server Flex instance deletion: %v\n", err)
	}

	fmt.Printf("Deleted SQL Server Flex instance \"%s\".\n", instanceId)
}
