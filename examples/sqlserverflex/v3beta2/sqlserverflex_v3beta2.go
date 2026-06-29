package main

import (
	"context"
	"fmt"
	"os"

	sqlserverflex "github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex/v3beta2api"
	"github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex/v3beta2api/wait"
)

func main() {
	const (
		projectId = "PROJECT_ID" // the uuid of your STACKIT project
		region    = "eu01"       // specify the region

		// Specify instance configuration options
		version = sqlserverflex.INSTANCEVERSION__2022
		// You can find a valid flavorId, by calling this API https://docs.api.stackit.cloud/documentation/mssql-flex-service/version/v3beta2#tag/Flavors
		// or using sqlserverflexClient.DefaultAPI.GetFlavors(ctx, projectId, region).Execute()
		flavorId = "FLAVOR_ID"
	)

	ctx := context.Background()

	// Create a new API client, that uses default authentication and configuration
	sqlserverflexClient, err := sqlserverflex.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// List the SQLServer Flex instances for your project
	listInstancesResp, err := sqlserverflexClient.DefaultAPI.ListInstances(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListInstances`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Number of instances: %v\n", len(listInstancesResp.Instances))

	// Create an instance
	createInstancePayload := sqlserverflex.CreateInstancePayload{
		Name:           "my-instance",
		FlavorId:       flavorId,
		Version:        version,
		BackupSchedule: "0 2 * * *",
		RetentionDays:  30,
		Network: sqlserverflex.CreateInstancePayloadNetwork{
			Acl: []string{"1.2.3.4/32"},
		},
		Storage: sqlserverflex.StorageCreate{
			Class: "premium-perf2-stackit",
			Size:  5,
		},
	}
	instance, err := sqlserverflexClient.DefaultAPI.CreateInstance(ctx, projectId, region).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating SQL Server Flex instance: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Triggered SQL Server Flex instance creation %q.\n", instance.Id)

	// Wait for the instance to become active
	_, err = wait.CreateInstanceWaitHandler(ctx, sqlserverflexClient.DefaultAPI, projectId, region, instance.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for SQL Server Flex instance creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created SQL Server Flex instance %q.\n", instance.Id)

	// Get an instance
	instanceResp, err := sqlserverflexClient.DefaultAPI.GetInstance(ctx, projectId, region, instance.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstance`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Get instance %+v.\n", instanceResp)

	// Delete an instance
	err = sqlserverflexClient.DefaultAPI.DeleteInstance(ctx, projectId, region, instance.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting SQL Server Flex instance: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Triggered SQL Server Flex instance deletion %q.\n", instance.Id)

	// Wait for the instance to be deleted, using the wait handler
	_, err = wait.DeleteInstanceWaitHandler(ctx, sqlserverflexClient.DefaultAPI, projectId, region, instance.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for SQL Server Flex instance deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Deleted SQL Server Flex instance %q.\n", instance.Id)
}
