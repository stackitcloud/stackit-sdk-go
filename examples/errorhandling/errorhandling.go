package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/dns"
)

func main() {
	// Any errors that occur when calling the API are converted to usual Go errors
	// We can see this by making a request with an unauthenticated client
	dnsClient, err := dns.NewAPIClient(config.WithoutAuthentication())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	_, err = dnsClient.ListZones(context.Background(), "foo-bar").Execute()
	if err == nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Unauthenticated GET request succeeded")
		os.Exit(1)
	}

	fmt.Printf("[DNS API] Unauthenticated request returned the following error: %v\n", err)
}
