package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/redis"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"
	planId := "PLAN_ID"

	// Create a new API client, that uses default authentication and configuration
	redisClient, err := redis.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the redis instances for your project
	getInstancesResp, err := redisClient.GetInstances(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
	} else {
		fmt.Printf("Number of instances: %v\n", len(*getInstancesResp.Instances))
	}

	// Get the redis offerings for your project
	getOfferingsResp, err := redisClient.GetOfferings(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetOfferings`: %v\n", err)
	} else {
		fmt.Printf("Offerings: %+v\n", getOfferingsResp.Offerings)
	}

	// Create a redis Instance
	createInstancePayload := redis.CreateInstancePayload{
		InstanceName: utils.Ptr("exampleInstance"),
		Parameters:   &redis.InstanceParameters{},
		PlanId:       utils.Ptr(planId),
	}
	createInstanceResp, err := redisClient.CreateInstance(context.Background(), projectId).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
	} else {
		fmt.Printf("Created instance with instance id \"%s\".\n", *createInstanceResp.InstanceId)
	}
}
