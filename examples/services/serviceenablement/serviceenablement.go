package main

import (
	"context"
	"fmt"
	"os"

	serviceenablement "github.com/stackitcloud/stackit-sdk-go/services/serviceenablement/v2api"
	"github.com/stackitcloud/stackit-sdk-go/services/serviceenablement/v2api/wait"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project

	// Create a new API client, that uses default authentication and configuration
	client, err := serviceenablement.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}
	region := "eu01"

	// List Services
	listServicesResp, err := client.DefaultAPI.ListServiceStatusRegional(context.Background(), region, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListServices`: %v\n", err)
	} else {
		fmt.Printf("Number of services: %v\n", len(listServicesResp.Items))
	}

	// Get Service Id from the list of services
	serviceId := listServicesResp.Items[0].ServiceId

	getServiceStatusResp, err := client.DefaultAPI.GetServiceStatusRegional(context.Background(), region, projectId, *serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetServiceStatus`: %v\n", err)
	} else {
		fmt.Printf("Service state: %s\n", *getServiceStatusResp.State)
	}

	// Enable Service
	err = client.DefaultAPI.EnableServiceRegional(context.Background(), region, projectId, *serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnableService`: %v\n", err)
	}
	// Wait for the service to be enabled
	status, err := wait.EnableServiceWaitHandler(context.Background(), client.DefaultAPI, region, projectId, *serviceId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for service to be enabled: %v\n", err)
	} else {
		fmt.Printf("Service %q is now enabled\n", *status.ServiceId)
	}

	// Disable Service
	err = client.DefaultAPI.DisableServiceRegional(context.Background(), region, projectId, *serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisableService`: %v\n", err)
	}
	// Wait for the service to be disabled
	status, err = wait.DisableServiceWaitHandler(context.Background(), client.DefaultAPI, region, projectId, *serviceId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when waiting for service to be disabled: %v\n", err)
	} else {
		fmt.Printf("Service %q is now disabled\n", *status.ServiceId)
	}
}
