package main

import (
	"context"
	"fmt"
	"os"

	registry "github.com/stackitcloud/stackit-sdk-go/services/registry/v1api"
	"github.com/stackitcloud/stackit-sdk-go/services/registry/v1api/wait"
)

func main() {
	projectId := "PROJECT_ID" // the UUID of your STACKIT project
	region := "eu01"          // STACKIT region, e.g. "eu01"

	// Create a new API client using default authentication and configuration
	client, err := registry.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()

	// 1. List existing artifactories (registries) in the project
	listResp, err := client.DefaultAPI.ListArtifactories(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing artifactories: %v\n", err)
	} else {
		fmt.Printf("Found %d artifactory(ies)\n", len(listResp.Artifactories))
		for _, artifactory := range listResp.Artifactories {
			fmt.Printf("- %s (ID: %s, State: %s)\n", artifactory.Name, artifactory.Id, artifactory.State)
		}
	}

	// 2. Create a new artifactory
	artifactoryName := "my-registry"
	createPayload := registry.CreateArtifactoryPayload{
		Name: artifactoryName,
	}
	createResp, err := client.DefaultAPI.CreateArtifactory(ctx, projectId, region).
		CreateArtifactoryPayload(createPayload).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating artifactory: %v\n", err)
		os.Exit(1)
	}

	artifactoryId := createResp.Id
	fmt.Printf("Created artifactory with ID %q. Waiting for it to become ready...\n", artifactoryId)

	// 3. Wait for the artifactory to be ready
	_, err = wait.CreateArtifactoryWaitHandler(ctx, client.DefaultAPI, projectId, region, artifactoryId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for artifactory to be ready: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Artifactory is ready.")

	// 4. Retrieve artifactory details
	artifactory, err := client.DefaultAPI.GetArtifactory(ctx, projectId, region, artifactoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting artifactory details: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Artifactory details: Name=%s, State=%s, Url=%s\n", artifactory.Name, artifactory.State, artifactory.Url)

	// 5. Delete the artifactory
	err = client.DefaultAPI.DeleteArtifactory(ctx, projectId, region, artifactoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting artifactory: %v\n", err)
		os.Exit(1)
	}

	// Wait for the artifactory to be deleted
	_, err = wait.DeleteArtifactoryWaitHandler(ctx, client.DefaultAPI, projectId, region, artifactoryId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for artifactory to be deleted: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Artifactory deleted successfully.")
}
