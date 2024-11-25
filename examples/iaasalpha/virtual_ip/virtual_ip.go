package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha/wait"
)

func main() {
	// Specify the project ID and network ID
	projectId := "PROJECT_ID"
	networkId := "NETWORK_ID"

	// Create a new API client, that uses default authentication and configuration
	iaasalphaClient, err := iaasalpha.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	virtualIPs, err := iaasalphaClient.ListVirtualIPs(context.Background(), projectId, networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `ListVirtualIPs`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Number of virtual IPs: %v\n", len(*virtualIPs.Items))
	}

	// Create a virtual IP
	createVirtualIPPayload := iaasalpha.CreateVirtualIPPayload{
		Name: utils.Ptr("example-vip"),
		Labels: &map[string]interface{}{
			"key": "value",
		},
	}
	virtualIP, err := iaasalphaClient.CreateVirtualIP(context.Background(), projectId, networkId).CreateVirtualIPPayload(createVirtualIPPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `CreateVirtualIP`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Triggered creation of virtual IP with ID %q.\n", *virtualIP.Id)
	}

	// Wait for creation of the virtual IP
	virtualIP, err = wait.CreateVirtualIPWaitHandler(context.Background(), iaasalphaClient, projectId, networkId, *virtualIP.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Virtual IP %q has been successfully created.\n", *virtualIP.Id)

	// Delete a virtual IP
	err = iaasalphaClient.DeleteVirtualIP(context.Background(), projectId, networkId, *virtualIP.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `DeleteVirtualIP`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Triggered deletion of virtual IP with ID %q.\n", *virtualIP.Id)
	}

	// Wait for deletion of the virtual IP
	_, err = wait.DeleteVirtualIPWaitHandler(context.Background(), iaasalphaClient, projectId, networkId, *virtualIP.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Virtual IP %q has been successfully deleted.\n", *virtualIP.Id)
}
