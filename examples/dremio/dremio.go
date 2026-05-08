package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/dremio/v1alphaapi/wait/wait"

	dremio "github.com/stackitcloud/stackit-sdk-go/services/dremio/v1alphaapi"
)

func main() {
	region := "eu01"                            // Region where the resources will be created
	projectId := "PROJECT_ID"                   // Your STACKIT project ID
	serviceAccountKeyPath := "sa-key-path.json" // Path to your STACKIT service account json

	ctx := context.Background()

	dremioClient, err := dremio.NewAPIClient(
		config.WithRegion(region),
		config.WithServiceAccountKeyPath(serviceAccountKeyPath),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Creating a Dremio instance
	createDremioInstancePayload := dremio.CreateDremioInstancePayload{
		DisplayName: "myExampleDremioInstance",
		Description: utils.Ptr("This is a Dremio instance."),
	}
	createDremioInstanceResp, err := dremioClient.DefaultAPI.CreateDremioInstance(ctx, projectId, region).CreateDremioInstancePayload(createDremioInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when creating Dremio instance: %v\n", err)
		os.Exit(1)
	}
	dremioId := createDremioInstanceResp.Id
	fmt.Printf("[Dremio] Triggered creation of Dremio with ID: %s. Waiting for it to become active... \n", dremioId)

	_, err = wait.CreateDremioWaitHandler(ctx, dremioClient.DefaultAPI, projectId, region, dremioId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}

	// Update a Dremio instance
	_, err = dremioClient.DefaultAPI.UpdateDremioInstance(ctx, projectId, region, dremioId).UpdateDremioInstancePayload(dremio.UpdateDremioInstancePayload{
		DisplayName: utils.Ptr("myExampleDremioInstanceUpdated"),
	}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when updating Dremio instance: %v\n", err)
		os.Exit(1)
	}
	_, err = wait.UpdateDremioWaitHandler(ctx, dremioClient.DefaultAPI, projectId, region, dremioId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when waiting for update: %v\n", err)
		os.Exit(1)
	}

	// Listing Dremio instances
	listResp, err := dremioClient.DefaultAPI.ListDremioInstances(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when listing Dremio instances: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Dremio instances:")
	for i := range listResp.Dremios {
		instance := listResp.Dremios[i]
		fmt.Printf("%s - %s\n", instance.Id, instance.DisplayName)
	}
	fmt.Println()

	// Creating a Dremio user
	createDremioUserPayload := dremio.CreateDremioUserPayload{
		Description: utils.Ptr("This is a new user."),
		Email:       "newuser@some.domain",
		FirstName:   "New",
		LastName:    "User",
		Name:        "newUser",
		Password:    "aV3ryS4feP4ssw0rd6",
	}
	createDremioUserResp, err := dremioClient.DefaultAPI.CreateDremioUser(ctx, projectId, region, dremioId).CreateDremioUserPayload(createDremioUserPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when creating Dremio user: %v\n", err)
		os.Exit(1)
	}
	dremioUserId := createDremioUserResp.Id
	fmt.Printf("[Dremio] Created Dremio user with ID: %s\n", dremioUserId)

	_, err = wait.CreateDremioUserWaitHandler(ctx, dremioClient.DefaultAPI, projectId, region, dremioId, dremioUserId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when waiting for user creation: %v\n", err)
		os.Exit(1)
	}

	// List Dremio users
	listUsersResp, err := dremioClient.DefaultAPI.ListDremioUsers(ctx, projectId, region, dremioId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when listing Dremio users: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Dremio users:")
	for i := range listUsersResp.DremioUsers {
		user := &listUsersResp.DremioUsers[i]
		fmt.Printf("%s - %s\n", user.Id, user.Name)
	}
	fmt.Println()

	// Delete a Dremio user
	err = dremioClient.DefaultAPI.DeleteDremioUser(ctx, projectId, region, dremioId, dremioUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when deleting Dremio user: %v\n", err)
		os.Exit(1)
	}
	_, err = wait.DeleteDremioUserWaitHandler(ctx, dremioClient.DefaultAPI, projectId, region, dremioId, dremioUserId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when waiting for user deletion: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[Dremio] Deleted Dremio user with ID: %s\n", dremioUserId)

	// Delete a Dremio instance
	err = dremioClient.DefaultAPI.DeleteDremioInstance(ctx, projectId, region, dremioId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when deleting Dremio instance: %v\n", err)
		os.Exit(1)
	}
	_, err = wait.DeleteDremioWaitHandler(ctx, dremioClient.DefaultAPI, projectId, region, dremioId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[Dremio] Deleted Dremio instance with ID: %s\n", dremioId)
}
