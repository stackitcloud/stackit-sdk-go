package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	git "github.com/stackitcloud/stackit-sdk-go/services/git/v1betaapi"
	"github.com/stackitcloud/stackit-sdk-go/services/git/v1betaapi/wait"
)

func main() {
	ctx := context.Background()

	projectId := "PROJECT_ID" // the uuid of your STACKIT project

	// Create a new API client, that uses default authentication and configuration
	gitClient, err := git.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the Git instances for your project
	listInstancesResp, err := gitClient.DefaultAPI.ListInstances(ctx, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListInstances`: %v\n", err)
	} else {
		fmt.Printf("Number of instances: %v\n", len(listInstancesResp.Instances))
	}

	// Get the Git offerings for your project
	getFlavorsResp, err := gitClient.DefaultAPI.ListFlavors(ctx, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListFlavors`: %v\n", err)
	} else {
		fmt.Printf("Flavors: %+v\n", getFlavorsResp.Flavors)
	}

	// Create a Git instance
	createInstancePayload := git.CreateInstancePayload{
		Name:   "example",
		Acl:    []string{"1.2.3.4/32"},
		Flavor: utils.Ptr(git.CreateInstancePayloadFlavor(getFlavorsResp.Flavors[0].Id)),
	}
	createInstanceResp, err := gitClient.DefaultAPI.CreateInstance(ctx, projectId).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Triggered creation of instance with instance id \"%s\".\n", createInstanceResp.Id)

	// Wait for creation of Git instance
	instance, err := wait.CreateGitInstanceWaitHandler(ctx, gitClient.DefaultAPI, projectId, createInstanceResp.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Git instance %q has been successfully created.\n", instance.Id)

	// Delete a Git instance
	err = gitClient.DefaultAPI.DeleteInstance(ctx, projectId, instance.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling 'DeleteInstance': %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Deleting instance with instance id %q.\n", createInstanceResp.Id)

	// Wait for deletion of Git instance
	_, err = wait.DeleteGitInstanceWaitHandler(ctx, gitClient.DefaultAPI, projectId, instance.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Git instance %q has been successfully deleted.\n", instance.Id)
}
