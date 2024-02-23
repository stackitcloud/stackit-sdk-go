package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/dns"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a cancellable context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // This is needed to abort the background routine, to prevent resource leaks

	// Create a new API client, that uses key flow and has background token refresh enabled
	dnsClient, err := dns.NewAPIClient(
		config.WithServiceAccountKeyPath("path/to/service/account/key"),
		config.WithBackgroundTokenRefresh(ctx),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when creating API client: %v\n", err)
	}

	// The client can then be used as normal, a new token will be fetched whenever the current one is close to expiring
	// Use this for long-running operations (hours)
	for i := 0; i < 10; i++ {
		// Get the DNS zones for your project
		getZoneResp, err := dnsClient.ListZones(context.Background(), projectId).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ListZones`: %v\n", err)
		}
		fmt.Printf("Number of DNS zones: %v\n", len(*getZoneResp.Zones))

		time.Sleep(time.Hour)
	}
}
