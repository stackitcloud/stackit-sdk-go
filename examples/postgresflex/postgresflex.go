package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/postgresflex"
	"github.com/stackitcloud/stackit-sdk-go/services/postgresflex/wait"
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

	// Get the postgresql instances for your project
	getInstancesResp, err := postgresflexClient.ListInstances(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
		os.Exit(1)
	}
	items := *getInstancesResp.Items
	fmt.Printf("Number of instances: %v\n", len(items))

	// Create a user associated to an instance
	if len(items) == 0 {
		fmt.Fprintf(os.Stderr, "Can't create a user without instances")
		os.Exit(1)
	}

	instanceId := *items[0].Id
	username := "example-user"
	createUserPayload := postgresflex.CreateUserPayload{
		Username: &username,
		Roles:    &[]string{"login"},
	}
	_, err = postgresflexClient.CreateUser(context.Background(), projectId, instanceId).CreateUserPayload(createUserPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateUser`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created user \"%s\" associated to instance \"%s\".\n", username, *items[0].Name)

	// Delete an instance
	err = postgresflexClient.DeleteInstance(context.Background(), projectId, instanceId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when delete PostgreSQL Flex instance: %v", err)
	}

	_, err = wait.DeleteInstanceWaitHandler(context.Background(), postgresflexClient, projectId, instanceId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for PostgreSQL Flex instance deletion: %v", err)
	}

	fmt.Printf("Deleted PostgreSQL Flex instance \"%s\".\n", instanceId)

	// Force delete an instance
	err = postgresflexClient.ForceDeleteInstance(context.Background(), projectId, instanceId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when force delete PostgreSQL Flex instance: %v", err)
	}

	_, err = wait.ForceDeleteInstanceWaitHandler(context.Background(), postgresflexClient, projectId, instanceId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for PostgreSQL Flex instance force deletion: %v", err)
	}

	fmt.Printf("Force deleted PostgreSQL Flex instance \"%s\".\n", instanceId)
}
