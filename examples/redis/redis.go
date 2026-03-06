package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	redis "github.com/stackitcloud/stackit-sdk-go/services/redis/v1api"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	planId := "PLAN_ID"

	// Create a new API client, that uses default authentication and configuration
	redisClient, err := redis.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the redis instances for your project
	getInstancesResp, err := redisClient.DefaultAPI.ListInstances(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
	} else {
		fmt.Printf("Number of instances: %v\n", len(getInstancesResp.Instances))
	}

	// Get the redis offerings for your project
	getOfferingsResp, err := redisClient.DefaultAPI.ListOfferings(context.Background(), projectId).Execute()
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
	createInstanceResp, err := redisClient.DefaultAPI.CreateInstance(context.Background(), projectId).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
	} else {
		fmt.Printf("Created instance with instance id \"%s\".\n", createInstanceResp.InstanceId)
	}
}
