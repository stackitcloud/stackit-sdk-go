package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/dns"
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
	wres, err := dns.CreateZoneWaitHandler(ctx, dnsClient, projectId, zoneId).SetTimeout(15 * time.Minute).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Waiting for zone update: %v\n", err)
		os.Exit(1)
	}

	// At this stage the waiter is waiting for an update to the zone
	// You can make a manual request to the DNS API updating the zone that was just created
	// The waiter will finish, and you will get the output below
	got, ok := wres.(*dns.ZoneResponse)
	if !ok {
		fmt.Fprintf(os.Stderr, "[DNS API] Returned response has unexpected type: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "[DNS API] Zone with id %s update (state: %s)\n", *got.Zone.Id, *got.Zone.State)
}
