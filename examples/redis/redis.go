package main

import (
	"context"
	"fmt"
	"os"

	redis "github.com/stackitcloud/stackit-sdk-go/services/redis/v2api"
	wait "github.com/stackitcloud/stackit-sdk-go/services/redis/v2api/wait"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	planId := "PLAN_ID"
	region := "eu01"

	// Create a new API client, that uses default authentication and configuration
	redisClient, err := redis.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the redis instances for your project
	getInstancesResp, err := redisClient.DefaultAPI.ListInstances(context.Background(), projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
	} else {
		fmt.Printf("Number of instances: %v\n", len(getInstancesResp.Instances))
	}

	// Get the redis offerings for your project
	getOfferingsResp, err := redisClient.DefaultAPI.ListOfferings(context.Background(), projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetOfferings`: %v\n", err)
	} else {
		fmt.Printf("Offerings: %+v\n", getOfferingsResp.Offerings)
	}

	// Create a redis Instance
	createInstancePayload := redis.CreateInstancePayload{
		InstanceName: "exampleInstance",
		Parameters:   &redis.InstanceParameters{},
		PlanId:       planId,
	}
	createInstanceResp, err := redisClient.DefaultAPI.CreateInstance(context.Background(), projectId, region).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created instance with instance id \"%s\".\n", createInstanceResp.InstanceId)

	// Wait for creation of redis instance
	instance, err := wait.CreateInstanceWaitHandler(context.Background(), redisClient.DefaultAPI, projectId, region, createInstanceResp.InstanceId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Redis instance %q has been successfully created.\n", *instance.InstanceId)

	// Delete a redis instance
	err = redisClient.DefaultAPI.DeleteInstance(context.Background(), projectId, region, *instance.InstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling 'DeleteInstance': %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Deleting instance with instance id %q.\n", createInstanceResp.InstanceId)

	// Wait for deletion of redis instance
	_, err = wait.DeleteInstanceWaitHandler(context.Background(), redisClient.DefaultAPI, projectId, region, *instance.InstanceId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Redis instance %q has been successfully deleted.\n", *instance.InstanceId)
}
