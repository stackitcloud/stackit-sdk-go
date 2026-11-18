package main

import (
	"context"
	"fmt"
	"os"

	valkey "github.com/stackitcloud/stackit-sdk-go/services/valkey/v2api"
	"github.com/stackitcloud/stackit-sdk-go/services/valkey/v2api/wait"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	planId := "PLAN_ID"       // the uuid of the valkey plan
	region := "eu01"

	// Create a new API client, that uses default authentication and configuration
	valkeyClient, err := valkey.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the valkey instances for your project
	getInstancesResp, err := valkeyClient.DefaultAPI.ListInstances(context.Background(), projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
	} else {
		fmt.Printf("Number of instances: %v\n", len(getInstancesResp.Instances))
	}

	// Get the valkey offerings for your project
	getOfferingsResp, err := valkeyClient.DefaultAPI.ListOfferings(context.Background(), projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetOfferings`: %v\n", err)
	} else {
		fmt.Printf("Offerings: %+v\n", getOfferingsResp.Offerings)
	}

	// Create a valkey Instance
	createInstancePayload := valkey.CreateInstancePayload{
		InstanceName: "exampleInstance",
		Parameters:   &valkey.InstanceParameters{},
		PlanId:       planId,
	}
	createInstanceResp, err := valkeyClient.DefaultAPI.CreateInstance(context.Background(), projectId, region).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Triggered creation of instance with instance id \"%s\".\n", createInstanceResp.InstanceId)

	// Wait for creation of valkey instance
	instance, err := wait.CreateInstanceWaitHandler(context.Background(), valkeyClient.DefaultAPI, projectId, region, createInstanceResp.InstanceId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("valkey instance %q has been successfully created.\n", *instance.InstanceId)

	// Delete a valkey instance
	err = valkeyClient.DefaultAPI.DeleteInstance(context.Background(), projectId, region, *instance.InstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling 'DeleteInstance': %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Deleting instance with instance id %q.\n", createInstanceResp.InstanceId)

	// Wait for deletion of valkey instance
	_, err = wait.DeleteInstanceWaitHandler(context.Background(), valkeyClient.DefaultAPI, projectId, region, *instance.InstanceId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("valkey instance %q has been successfully deleted.\n", *instance.InstanceId)
}
