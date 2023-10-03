package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/dns"
)

func main() {

	// When creating a new API client without providing any configuration, it will setup default authentication.
	// The SDK will search for a valid service account key or token in several locations.
	// It will first try to use the key flow, by looking into the variables STACKIT_SERVICE_ACCOUNT_KEY, STACKIT_SERVICE_ACCOUNT_KEY_PATH,
	// STACKIT_PRIVATE_KEY and STACKIT_PRIVATE_KEY_PATH. If the keys cannot be retrieved, it will check the credentials file located in STACKIT_CREDENTIALS_PATH, if specified, or in
	// $HOME/.stackit/credentials.json as a fallback. If the key are found and are valid, the KeyAuth flow is used.
	// If the key flow cannot be used, it will try to find a token in the STACKIT_SERVICE_ACCOUNT_TOKEN. If not present, it will
	// search in the credentials file. If the token is found, the TokenAuth flow is used.
	// In case no authentication flow can be configured, the creation of a new client fails.
	_, err := dns.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create an unauthenticated API client
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
