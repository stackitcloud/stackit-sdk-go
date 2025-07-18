package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/mongodbflex"
	"github.com/stackitcloud/stackit-sdk-go/services/mongodbflex/wait"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Specify the region
	region := "REGION"

	// Create a new API client, that uses default authentication and configuration
	mongodbflexClient, err := mongodbflex.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the MongoDB Flex instances for your project
	getInstancesResp, err := mongodbflexClient.ListInstances(context.Background(), projectId, region).Tag("tag").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListInstances`: %v\n", err)
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
	_, err = mongodbflexClient.CreateUser(context.Background(), projectId, instanceId, region).CreateUserPayload(createUserPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateUser`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created user \"%s\" associated to instance \"%s\".\n", username, *items[0].Name)

	// Restore an instance from a backup
	restoreInstancePayload := mongodbflex.RestoreInstancePayload{
		BackupId:   utils.Ptr("BACKUP_ID"),
		InstanceId: &instanceId,
	}
	_, err = mongodbflexClient.RestoreInstance(context.Background(), projectId, instanceId, region).RestoreInstancePayload(restoreInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestoreInstance`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Restoring instance \"%s\" from backup \"%s\".\n", instanceId, "BACKUP_ID")

	_, err = wait.RestoreInstanceWaitHandler(context.Background(), mongodbflexClient, projectId, instanceId, "BACKUP_ID", region).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for restore to finish: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Restored instance \"%s\" from backup \"%s\".\n", instanceId, "BACKUP_ID")
}
