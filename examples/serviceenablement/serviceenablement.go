package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/serviceenablement"
	"github.com/stackitcloud/stackit-sdk-go/services/serviceenablement/wait"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	client, err := serviceenablement.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// List Services
	listServicesResp, err := client.ListServiceStatusExecute(context.Background(), projectId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListServices`: %v\n", err)
	} else {
		fmt.Printf("Number of services: %v\n", len(*listServicesResp.Items))
	}

	// Get Service Id from the list of services
	serviceId := (*listServicesResp.Items)[0].ServiceId

	getServiceStatusResp, err := client.GetServiceStatus(context.Background(), projectId, *serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetServiceStatus`: %v\n", err)
	} else {
		fmt.Printf("Service state: %s\n", *getServiceStatusResp.State)
	}

	// Enable Service
	err = client.EnableService(context.Background(), projectId, *serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnableService`: %v\n", err)
	}
	// Wait for the service to be enabled
	status, err := wait.EnableServiceWaitHandler(context.Background(), client, projectId, *serviceId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for service to be enabled: %v\n", err)
	} else {
		fmt.Printf("Service %q is now enabled\n", *status.ServiceId)
	}

	// Disable Service
	err = client.DisableService(context.Background(), projectId, *serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisableService`: %v\n", err)
	}
	// Wait for the service to be disabled
	status, err = wait.DisableServiceWaitHandler(context.Background(), client, projectId, *serviceId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for service to be disabled: %v\n", err)
	} else {
		fmt.Printf("Service %q is now disabled\n", *status.ServiceId)
	}
}
