package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	dremio "github.com/stackitcloud/stackit-sdk-go/services/dremio/v1alphaapi"
)

func main() {
	region := "eu01"          // Region where the resources will be created
	projectId := "PROJECT_ID" // Your STACKIT project ID

	ctx := context.Background()

	dremioClient, err := dremio.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Creating a Dremio instance
	createDremioInstancePayload := dremio.CreateDremioInstancePayload{
		DisplayName: "myExampleDremioInstance",
		Description: utils.Ptr("This is a Dremio instance."),
	}
	createDremioInstanceResponse, err := dremioClient.DefaultAPI.CreateDremioInstance(ctx, projectId, region).CreateDremioInstancePayload(createDremioInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when creating Dremio instance: %v\n", err)
		os.Exit(1)
	}
	dremioId := createDremioInstanceResponse.Id
	fmt.Printf("[Dremio] Triggered creation of Dremio with ID: %s. Waiting for it to become active... \n", dremioId)

	// Creating a Dremio user
	createDremioUserPayload := dremio.CreateDremioUserPayload{
		Description: utils.Ptr("This is a new user."),
		Email:       "newuser@some.domain",
		FirstName:   "New",
		LastName:    "User",
		Name:        "newUser",
		Password:    "aV3ryS4feP4ssw0rd6",
	}
	createDremioUserResponse, err := dremioClient.DefaultAPI.CreateDremioUser(ctx, projectId, region, dremioId).CreateDremioUserPayload(createDremioUserPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Dremio] Error when creating Dremio user: %v\n", err)
		os.Exit(1)
	}
	dremioUserId := createDremioUserResponse.Id
	fmt.Printf("[Dremio] Created Dremio User with ID: %s\n", dremioUserId)
}
