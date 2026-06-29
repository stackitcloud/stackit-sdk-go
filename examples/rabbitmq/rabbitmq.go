package main

import (
	"context"
	"fmt"
	"os"

	rabbitmq "github.com/stackitcloud/stackit-sdk-go/services/rabbitmq/v2api"
	wait "github.com/stackitcloud/stackit-sdk-go/services/rabbitmq/v2api/wait"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	region := "eu01"
	planId := "PLAN_ID"

	// Create a new API client, that uses default authentication and configuration
	rabbitmqClient, err := rabbitmq.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the rabbitmq instances for your project
	getInstancesResp, err := rabbitmqClient.DefaultAPI.ListInstances(context.Background(), projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Number of instances: %v\n", len(getInstancesResp.Instances))

	// Get the rabbitmq offerings for your project
	getOfferingsResp, err := rabbitmqClient.DefaultAPI.ListOfferings(context.Background(), projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetOfferings`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Offerings: %+v\n", getOfferingsResp.Offerings)

	// Create a rabbitmq Instance
	createInstancePayload := rabbitmq.CreateInstancePayload{
		InstanceName: "exampleInstance",
		Parameters:   &rabbitmq.InstanceParameters{},
		PlanId:       planId,
	}
	createInstanceResp, err := rabbitmqClient.DefaultAPI.CreateInstance(context.Background(), projectId, region).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created instance with instance id %q.\n", createInstanceResp.InstanceId)

	// Wait for creation of rabbitmq instance
	instance, err := wait.CreateInstanceWaitHandler(context.Background(), rabbitmqClient.DefaultAPI, projectId, region, createInstanceResp.InstanceId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Rabbitmq instance %q has been successfully created.\n", *instance.InstanceId)

	// Delete a rabbitmq instance
	err = rabbitmqClient.DefaultAPI.DeleteInstance(context.Background(), projectId, region, *instance.InstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling 'DeleteInstance': %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Deleting instance with instance id %q.\n", createInstanceResp.InstanceId)

	// Wait for deletion of rabbitmq instance
	_, err = wait.DeleteInstanceWaitHandler(context.Background(), rabbitmqClient.DefaultAPI, projectId, region, *instance.InstanceId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Rabbitmq instance %q has been successfully deleted.\n", *instance.InstanceId)
}
