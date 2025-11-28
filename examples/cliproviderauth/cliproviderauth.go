package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/cliauth"
	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/dns"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project

	// Example 1: Use CLI provider authentication with default profile
	// This reads credentials stored by the STACKIT CLI from the system keyring or file fallback
	// The SDK will automatically refresh tokens when needed
	flow, err := cliauth.NewCLIProviderFlow("", nil, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CLI Auth] Creating CLI provider flow: %v\n", err)
		fmt.Fprintf(os.Stderr, "Make sure you're authenticated with the STACKIT CLI first.\n")
		os.Exit(1)
	}

	// Create a DNS client using the CLI provider authentication
	dnsClient, err := dns.NewAPIClient(
		config.WithCustomAuth(flow),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Make an authenticated request
	getZoneResp, err := dnsClient.ListZones(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Error when calling `ZoneApi.GetZones`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[DNS API] Number of zones: %v\n", len(*getZoneResp.Zones))

	// Example 2: Use CLI provider authentication with a specific profile
	// This is useful when you have multiple CLI profiles configured
	profileName := "production"
	flowWithProfile, err := cliauth.NewCLIProviderFlow(profileName, nil, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CLI Auth] Creating CLI provider flow with profile '%s': %v\n", profileName, err)
		os.Exit(1)
	}

	dnsClientWithProfile, err := dns.NewAPIClient(
		config.WithCustomAuth(flowWithProfile),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client with profile: %v\n", err)
		os.Exit(1)
	}

	// Make an authenticated request with the profile
	getZoneResp2, err := dnsClientWithProfile.ListZones(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Error when calling `ZoneApi.GetZones`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[DNS API] Number of zones (profile '%s'): %v\n", profileName, len(*getZoneResp2.Zones))

	// Example 3: Direct credential access (advanced use case)
	// For cases where you need direct access to the credentials
	creds, err := cliauth.ReadCredentials("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CLI Auth] Reading credentials: %v\n", err)
		os.Exit(1)
	}

	// Check if token needs refresh
	if cliauth.IsTokenExpired(creds) {
		fmt.Println("[CLI Auth] Token is expired, refreshing...")
		err = cliauth.RefreshToken(creds)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[CLI Auth] Refreshing token: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("[CLI Auth] Token refreshed successfully")
	}

	fmt.Printf("[CLI Auth] Access token: %s...\n", creds.AccessToken[:20])
}
