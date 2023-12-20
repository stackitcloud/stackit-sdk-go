package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/mariadb"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"
	planId := "PLAN_ID"

	// Create a new API client, that uses default authentication and configuration
	mariadbClient, err := mariadb.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the mariadb instances for your project
	getInstancesResp, err := mariadbClient.ListInstances(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstances`: %v\n", err)
	} else {
		fmt.Printf("Number of instances: %v\n", len(*getInstancesResp.Instances))
	}

	// Get the mariadb offerings for your project
	getOfferingsResp, err := mariadbClient.ListOfferings(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetOfferings`: %v\n", err)
	} else {
		fmt.Printf("Offerings: %+v\n", getOfferingsResp.Offerings)
	}

	// Create a mariadb Instance
	createInstancePayload := mariadb.CreateInstancePayload{
		InstanceName: utils.Ptr("exampleInstance"),
		Parameters:   &mariadb.InstanceParameters{},
		PlanId:       utils.Ptr(planId),
	}
	createInstanceResp, err := mariadbClient.CreateInstance(context.Background(), projectId).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
	} else {
		fmt.Printf("Created instance with instance id \"%s\".\n", *createInstanceResp.InstanceId)
	}
}
