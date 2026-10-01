package main

import (
	"context"
	"fmt"
	"os"

	functions "github.com/stackitcloud/stackit-sdk-go/services/functions/v1alphaapi"
	wait "github.com/stackitcloud/stackit-sdk-go/services/functions/v1alphaapi/wait"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project

	// Create a new API client, that uses default authentication and configuration
	functionsClient, err := functions.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// List functions
	functionsList, err := functionsClient.DefaultAPI.ListFunctions(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing functions: %v\n", err)
		os.Exit(1)
	}
	_, err = fmt.Fprintf(os.Stdout, "Function count: %d\n", len(functionsList.Functions))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when printing out function count: %v\n", err)
		os.Exit(1)
	}

	// List pull-credentials
	pull_credentials, err := functionsClient.DefaultAPI.ListPullCredentials(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when listing pull-credentials count: %v\n", err)
		os.Exit(1)
	}
	_, err = fmt.Fprintf(os.Stdout, "Pull-credentials count: %d\n", len(pull_credentials.PullCredentials))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when printing out pull-credentials count: %v\n", err)
		os.Exit(1)
	}

	// Create pull-credential
	payload := *functions.NewCreatePullCredentialPayload("ghcr.io", "MY_TOKEN", "MY_USERNAME")
	pull_credential, err := functionsClient.DefaultAPI.CreatePullCredential(context.Background(), projectId).CreatePullCredentialPayload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "erorr creating pull-credential: %v\n", err)
		os.Exit(1)
	}
	_, err = fmt.Fprintf(os.Stdout, "Created Pull-credentials address: %d\n", len(pull_credential.Address))
	if err != nil {
		fmt.Fprintf(os.Stderr, "can not print out function pull-credentials address: %v\n", err)
		os.Exit(1)
	}

	// List pull-credentials
	pull_credentials, err = functionsClient.DefaultAPI.ListPullCredentials(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListPullCredentials`: %v\n", err)
		os.Exit(1)
	}
	_, err = fmt.Fprintf(os.Stdout, "Pull-credentials count: %d\n", len(pull_credentials.PullCredentials))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when printing out pull-credentials count: %v\n", err)
		os.Exit(1)
	}

	// Create function deployment
	functionData := *functions.NewCreateFunctionPayload("FUNKY_FUNC")
	function, err := functionsClient.DefaultAPI.CreateFunction(context.Background(), projectId).CreateFunctionPayload(functionData).Execute()
	if err != nil || function.Id == nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateFunction`: %v\n", err)
		os.Exit(1)
	}
	_, err = fmt.Printf("Function with id %s created\n", *function.Id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error printing out function: %v\n", err)
		os.Exit(1)
	}

	// Create function revision deployment
	revisionData := *functions.NewCreateFunctionRevisionPayload(*functions.NewResourceLimits(100, "f2"), "ghcr.io/stackitcloud/example-image:0.1.0")
	functionRevision, err := functionsClient.DefaultAPI.CreateFunctionRevision(context.Background(), projectId, *function.Id).CreateFunctionRevisionPayload(revisionData).Execute()
	if err != nil || functionRevision.Id == nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateFunctionRevision`: %v\n", err)
		os.Exit(1)
	}
	_, err = fmt.Printf("Function revision created at: %s\n", functionRevision.CreatedAt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when printing out function revision: %v\n", err)
		os.Exit(1)
	}

	// Wait for creation of function revision
	revision, err := wait.CreateRevisionWaitHandler(context.Background(), functionsClient.DefaultAPI, projectId, *function.Id, *functionRevision.Id).WaitWithContext(context.Background())
	if err != nil || revision.Id == nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for revision creation: %v\n", err)
		os.Exit(1)
	}
	_, err = fmt.Printf("Function revision %s has been successfully deployed.\n", *revision.Id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when printing out function revision: %v\n", err)
		os.Exit(1)
	}
}
