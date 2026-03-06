package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	logme "github.com/stackitcloud/stackit-sdk-go/services/logme/v1api"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	planId := "PLAN_ID"

	// Create a new API client, that uses default authentication and configuration
	logmeClient, err := logme.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the logme instances for your project
	getInstancesResp, err := logmeClient.DefaultAPI.ListInstances(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
	} else {
		fmt.Printf("Number of instances: %v\n", len(getInstancesResp.Instances))
	}

	// Get the logme offerings for your project
	getOfferingsResp, err := logmeClient.DefaultAPI.ListOfferings(context.Background(), projectId).Execute()
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
	createInstanceResp, err := logmeClient.DefaultAPI.CreateInstance(context.Background(), projectId).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
	} else {
		fmt.Printf("Created instance with instance id \"%s\".\n", createInstanceResp.InstanceId)
	}
}
