package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/serviceaccount"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	client, err := serviceaccount.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the service accounts for your project
	getAccountsResp, err := client.ListServiceAccounts(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetServiceAccounts`: %v\n", err)
	} else {
		fmt.Printf("Number of service accounts: %v\n", len(*getAccountsResp.Items))
	}

	// Create a service account
	createAccountPayload := serviceaccount.CreateServiceAccountPayload{
		Name: utils.Ptr("my-service-account"),
	}
	createAccountResp, err := client.CreateServiceAccount(context.Background(), projectId).CreateServiceAccountPayload(createAccountPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateServiceAccount`: %v\n", err)
	} else {
		fmt.Printf("Created service account with email \"%s\".\n", *createAccountResp.Email)
	}
}
