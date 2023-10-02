package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/dns"
)

func main() {
	// Create a new API client, that uses default authentication.
	// The SDK will first try to find a token in the STACKIT_SERVICE_ACCOUNT_TOKEN env var. If not present, it will
	// check the credentials file located in the path defined by the STACKIT_CREDENTIALS_PATH env var, if specified,
	// or in $HOME/.stackit/credentials.json as a fallback. If the token is found, all of the requests are authenticated using that token.
	_, err := dns.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// // Create an unauthenticated API client
	_, err = dns.NewAPIClient(config.WithoutAuthentication())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create a new API client, that will authenticate using the provided bearer token
	token := "TOKEN"
	_, err = dns.NewAPIClient(config.WithToken(token))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create a new API client, that will authenticate using the key flow
	saKeyPath := "/path/to/service_account_key.json"
	privateKeyPath := "/path/to/private.key"
	dnsClient, err := dns.NewAPIClient(
		config.WithServiceAccountKeyPath(saKeyPath),
		config.WithPrivateKeyPath(privateKeyPath),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Check that you can make an authenticated request
	getZoneResp, err := dnsClient.GetZones(context.Background(), "PROJECT_ID").Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Error when calling `ZoneApi.GetZones`: %v\n", err)
	} else {
		fmt.Printf("[DNS API] Number of zones: %v\n", len(*getZoneResp.Zones))
	}
}
