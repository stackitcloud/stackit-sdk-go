package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	logme "github.com/stackitcloud/stackit-sdk-go/services/logme/v2api"
	"github.com/stackitcloud/stackit-sdk-go/services/logme/v2api/wait"
)

func main() {
	ctx := context.Background()

	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	planId := "PLAN_ID"
	region := "eu01"

	// Create a new API client, that uses default authentication and configuration
	logmeClient, err := logme.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the logme instances for your project
	getInstancesResp, err := logmeClient.DefaultAPI.ListInstances(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
	} else {
		fmt.Printf("Number of instances: %v\n", len(getInstancesResp.Instances))
	}

	// Get the logme offerings for your project
	getOfferingsResp, err := logmeClient.DefaultAPI.ListOfferings(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetOfferings`: %v\n", err)
	} else {
		fmt.Printf("Offerings: %+v\n", getOfferingsResp.Offerings)
	}

	// Create a logme Instance
	createInstancePayload := logme.CreateInstancePayload{
		InstanceName: "exampleInstance",
		Parameters:   &logme.InstanceParameters{},
		PlanId:       planId,
	}
	createInstanceResp, err := logmeClient.DefaultAPI.CreateInstance(ctx, projectId, region).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Triggered creation of instance with instance id \"%s\".\n", createInstanceResp.InstanceId)

	// Wait for creation of logme instance
	instance, err := wait.CreateInstanceWaitHandler(ctx, logmeClient.DefaultAPI, projectId, region, createInstanceResp.InstanceId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("LogMe instance %q has been successfully created.\n", *instance.InstanceId)

	// Update a logme Instance
	updateInstancePayload := logme.PartialUpdateInstancePayload{
		InstanceName: utils.Ptr("updatedInstance"),
		Parameters:   &logme.InstanceParameters{},
		PlanId:       &planId,
	}
	err = logmeClient.DefaultAPI.PartialUpdateInstance(ctx, projectId, region, *instance.InstanceId).PartialUpdateInstancePayload(updateInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartialUpdateInstance`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Triggered partial update of instance with instance id \"%q\".\n", *instance.InstanceId)

	// Wait for update of logme instance
	instance, err = wait.PartialUpdateInstanceWaitHandler(ctx, logmeClient.DefaultAPI, projectId, region, createInstanceResp.InstanceId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for update: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("LogMe instance %q has been successfully updated.\n", *instance.InstanceId)

	// Delete a logme instance
	err = logmeClient.DefaultAPI.DeleteInstance(ctx, projectId, region, *instance.InstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling 'DeleteInstance': %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Deleting instance with instance id %q.\n", createInstanceResp.InstanceId)

	// Wait for deletion of logme instance
	_, err = wait.DeleteInstanceWaitHandler(ctx, logmeClient.DefaultAPI, projectId, region, *instance.InstanceId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("LogMe instance %q has been successfully deleted.\n", *instance.InstanceId)
}
