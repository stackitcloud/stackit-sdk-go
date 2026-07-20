package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	postgresflex "github.com/stackitcloud/stackit-sdk-go/services/postgresflex/v3beta1api"
	"github.com/stackitcloud/stackit-sdk-go/services/postgresflex/v3beta1api/wait"
)

func main() {
	const (
		projectId = "PROJECT_ID" // the uuid of your STACKIT project
		region    = "eu01"       // Specify the region
		// Specify instance configuration options
		version = "VERSION"
		// You can find a valid flavorId, by calling this API https://docs.api.stackit.cloud/documentation/postgres-flex-service/version/v3alpha1#tag/Flavor/operation/getFlavorsRequest
		// or using postgresflexClient.DefaultAPI.ListFlavors(ctx, projectId, region).Execute()
		flavorId = "FLAVOR_ID"
	)

	ctx := context.Background()

	// Create a new API client, that uses default authentication and configuration
	postgresflexClient, err := postgresflex.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// List the PostgreSQL Flex instances for your project
	getInstancesResp, err := postgresflexClient.DefaultAPI.ListInstances(ctx, projectId, region).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListInstances`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Number of instances: %v\n", len(getInstancesResp.Instances))

	// Create an instance
	createInstancePayload := postgresflex.CreateInstancePayload{
		Name:           "my-instance",
		FlavorId:       flavorId,
		Version:        version,
		BackupSchedule: "0 2 * * *",
		RetentionDays:  *postgresflex.NewNullableInt32(utils.Ptr(int32(35))),
		Network: postgresflex.InstanceNetworkCreate{
			Acl: []string{"1.2.3.4/32"},
		},
		Storage: postgresflex.StorageCreate{
			Class: utils.Ptr("premium-perf2-stackit"),
			Size:  5,
		},
	}
	instance, err := postgresflexClient.DefaultAPI.CreateInstance(ctx, projectId, region).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating PostgreSQL Flex instance: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Triggered PostgreSQL Flex instance creation %q.\n", instance.Id)

	// Wait for the instance to become active
	_, err = wait.CreateInstanceWaitHandler(ctx, postgresflexClient.DefaultAPI, projectId, region, instance.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for PostgreSQL Flex instance creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created PostgreSQL Flex instance %q.\n", instance.Id)

	// Get an instance
	instanceResp, err := postgresflexClient.DefaultAPI.GetInstance(ctx, projectId, region, instance.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstance`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Get instance %+v.\n", instanceResp)

	// Create a user associated to an instance
	instanceId := instanceResp.Id
	username := "myUser"
	createUserPayload := postgresflex.CreateUserPayload{
		Name:  username,
		Roles: []string{"login"},
	}
	user, err := postgresflexClient.DefaultAPI.CreateUser(ctx, projectId, region, instanceId).CreateUserPayload(createUserPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateUser`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Triggered creation of user %q associated to instance %q.\n", username, instanceResp.Name)

	// Wait for the user to become active
	_, err = wait.CreateUserWaitHandler(ctx, postgresflexClient.DefaultAPI, projectId, region, instance.Id, user.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for PostgreSQL Flex user creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created PostgreSQL Flex user %q.\n", user.Id)

	// Get a user
	userResp, err := postgresflexClient.DefaultAPI.GetUser(ctx, projectId, region, instance.Id, user.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetUser`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Get user %+v.\n", userResp)

	// Delete a user
	err = postgresflexClient.DefaultAPI.DeleteUser(ctx, projectId, region, instanceId, user.Id).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when deleting PostgreSQL Flex user: %v", err)
		os.Exit(1)
	}

	_, err = wait.DeleteUserWaitHandler(ctx, postgresflexClient.DefaultAPI, projectId, region, instanceId, user.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for PostgreSQL Flex user deletion: %v", err)
	}

	fmt.Printf("Deleted PostgreSQL Flex user %q.\n", instanceId)

	// Update an instance
	updateInstancePayload := postgresflex.PartialUpdateInstancePayload{
		Name: utils.Ptr("updated-instance"),
	}
	err = postgresflexClient.DefaultAPI.PartialUpdateInstance(ctx, projectId, region, instanceId).PartialUpdateInstancePayload(updateInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating PostgreSQL Flex instance: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Triggered PostgreSQL Flex instance update %q.\n", instance.Id)

	// Wait for the instance to become active
	_, err = wait.PartialUpdateInstanceWaitHandler(ctx, postgresflexClient.DefaultAPI, projectId, region, instanceId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for PostgreSQL Flex instance update: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Updated PostgreSQL Flex instance %q.\n", instance.Id)

	// Delete an instance
	err = postgresflexClient.DefaultAPI.DeleteInstance(ctx, projectId, region, instanceId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when deleting PostgreSQL Flex instance: %v", err)
		os.Exit(1)
	}

	_, err = wait.DeleteInstanceWaitHandler(ctx, postgresflexClient.DefaultAPI, projectId, region, instanceId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for PostgreSQL Flex instance deletion: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Deleted PostgreSQL Flex instance %q.\n", instanceId)
}
