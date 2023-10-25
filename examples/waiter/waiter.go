package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/dns"
	"github.com/stackitcloud/stackit-sdk-go/services/dns/wait"
)

func main() {
	ctx := context.Background()

	// Specify the project ID and DNS name (must be a valid and unique FQDN)
	projectId := "PROJECT_ID"
	dnsName := "zoneTest.com"

	// Create a new API client, that uses default authentication.
	dnsClient, err := dns.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create a DNS Zone
	createZonePayload := dns.CreateZonePayload{
		Name:    utils.Ptr("myZone"),
		DnsName: utils.Ptr(dnsName),
	}

	createZoneResp, err := dnsClient.CreateZone(ctx, projectId).CreateZonePayload(createZonePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Error when calling `ZoneApi.CreateZone`: %v\n", err)
		os.Exit(1)
	}

	zoneId := *createZoneResp.Zone.Id

	// The following will wait until the DNS zone is finshed being created
	wres, err := wait.CreateZoneWaitHandler(ctx, dnsClient, projectId, zoneId).SetTimeout(15 * time.Minute).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Waiting for zone creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "[DNS API] Zone with id %s created (state: %s)\n", *wres.Zone.Id, *wres.Zone.State) // The state is always successful
}
