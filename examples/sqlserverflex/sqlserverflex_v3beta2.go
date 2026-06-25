package main

import (
	"context"
	"fmt"
	"os"

	sqlserverflex "github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex/v3beta2api"
	"github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex/v3beta2api/wait"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project

	// Specify the region
	region := "eu01"

	// Specify instance configuration options
	flavorId := "FLAVOR_ID"
	version := sqlserverflex.INSTANCEVERSION__2022

	ctx := context.Background()

	// Create a new API client, that uses default authentication and configuration
	sqlserverflexClient, err := sqlserverflex.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the SQLServer Flex instances for your project
	getInstancesResp, err := sqlserverflexClient.DefaultAPI.ListInstances(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListInstances`: %v\n", err)
		os.Exit(1)
	}
	items := getInstancesResp.Instances
	fmt.Printf("Number of instances: %v\n", len(items))

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
	instanceId := instance.Id
	fmt.Printf("Triggered SQL Server Flex instance creation %q.\n", instanceId)

	// Wait for the instance to become active
	_, err = wait.CreateInstanceWaitHandler(ctx, sqlserverflexClient.DefaultAPI, projectId, region, instanceId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for SQL Server Flex instance creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created SQL Server Flex instance %q.\n", instanceId)

	// Get an instance
	instanceResp, err := sqlserverflexClient.DefaultAPI.GetInstance(ctx, projectId, region, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstance`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Get instance %+v.\n", instanceResp)

	// Delete an instance
	err = sqlserverflexClient.DefaultAPI.DeleteInstance(ctx, projectId, region, instanceId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting SQL Server Flex instance: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Triggered SQL Server Flex instance deletion %q.\n", instanceId)

	// Wait for the instance to be deleted, using the wait handler
	_, err = wait.DeleteInstanceWaitHandler(ctx, sqlserverflexClient.DefaultAPI, projectId, region, instanceId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for SQL Server Flex instance deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Deleted SQL Server Flex instance %q.\n", instanceId)
}
