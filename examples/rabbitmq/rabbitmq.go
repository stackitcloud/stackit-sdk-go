package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/rabbitmq"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"
	planId := "PLAN_ID"

	// Create a new API client, that uses default authentication and configuration
	rabbitmqClient, err := rabbitmq.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the rabbitmq instances for your project
	getInstancesResp, err := rabbitmqClient.GetInstances(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
	} else {
		fmt.Printf("Number of instances: %v\n", len(*getInstancesResp.Instances))
	}

	// Get the rabbitmq offerings for your project
	getOfferingsResp, err := rabbitmqClient.GetOfferings(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetOfferings`: %v\n", err)
	} else {
		fmt.Printf("Offerings: %+v\n", getOfferingsResp.Offerings)
	}

	// Create a rabbitmq Instance
	createInstancePayload := rabbitmq.CreateInstancePayload{
		InstanceName: utils.Ptr("exampleInstance"),
		Parameters:   &rabbitmq.InstanceParameters{},
		PlanId:       utils.Ptr(planId),
	}
	createInstanceResp, err := rabbitmqClient.CreateInstance(context.Background(), projectId).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
	} else {
		fmt.Printf("Created instance with instance id \"%s\".\n", *createInstanceResp.InstanceId)
	}
}
