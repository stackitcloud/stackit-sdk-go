package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas"
)

func main() {
	// Specify the organization ID and project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	iaasClient, err := iaas.NewAPIClient(
		config.WithRegion("eu01"),
		config.WithCliAuth(),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	servers, err := iaasClient.ListServers(context.Background(), projectId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `ListServers`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Number of servers: %v\n", len(*servers.Items))
	}

}
