package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/logme"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"
	planId := "PLAN_ID"

	// Create a new API client, that uses default authentication and configuration
	logmeClient, err := logme.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the logme instances for your project
	getInstancesResp, err := logmeClient.GetInstances(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
	} else {
		fmt.Printf("Number of instances: %v\n", len(*getInstancesResp.Instances))
	}

	// Get the logme offerings for your project
	getOfferingsResp, err := logmeClient.GetOfferings(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetOfferings`: %v\n", err)
	} else {
		fmt.Printf("Offerings: %+v\n", getOfferingsResp.Offerings)
	}

	// Create a logme Instance
	createInstancePayload := logme.CreateInstancePayload{
		InstanceName: utils.Ptr("exampleInstance"),
		Parameters:   &logme.InstanceParameters{},
		PlanId:       utils.Ptr(planId),
	}
	createInstanceResp, err := logmeClient.CreateInstance(context.Background(), projectId).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
	} else {
		fmt.Printf("Created instance with instance id \"%s\".\n", *createInstanceResp.InstanceId)
	}
}
