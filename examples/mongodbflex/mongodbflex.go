package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/mongodbflex"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	mongodbflexClient, err := mongodbflex.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the MongoDB Flex instances for your project
	getInstancesResp, err := mongodbflexClient.GetInstances(context.Background(), projectId).Tag("tag").Execute()
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
	createUserPayload := mongodbflex.CreateUserPayload{
		Username: &username,
		Database: utils.Ptr("default"),
		Roles:    &[]string{"read"},
	}
	_, err = mongodbflexClient.CreateUser(context.Background(), projectId, instanceId).CreateUserPayload(createUserPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateUser`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created user \"%s\" associated to instance \"%s\".\n", username, *items[0].Name)
}
