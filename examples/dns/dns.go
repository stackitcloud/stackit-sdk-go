package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/dns"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	dnsClient, err := dns.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the DNS Zones for your project
	var getZoneResp *dns.ZonesResponse //nolint:golint // transparency on data model naming
	getZoneResp, err = dnsClient.GetZones(context.Background(), projectId).Execute()

	// Get only active DNS Zones for your project by adding the filter "ActiveEq(true)" to the call. More filters are available and can be chained.
	// dnsRespGetZones, err := dnsClient.ZoneApi.GetZones(context.Background(), projectId).ActiveEq(true).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Error when calling `ZoneApi.GetZones`: %v\n", err)
	} else {
		fmt.Printf("[DNS API] Number of zones: %v\n", len(*getZoneResp.Zones))
	}

	// Create a DNS Zone
	createZonePayload := dns.CreateZonePayload{
		Name:    utils.Ptr("myZone"),
		DnsName: utils.Ptr("testZone.com"),
	}
	var createZoneResp *dns.ZoneResponse //nolint:golint // transparency on data model naming
	createZoneResp, err = dnsClient.CreateZone(context.Background(), projectId).CreateZonePayload(createZonePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Error when calling `ZoneApi.CreateZone`: %v\n", err)
	} else {
		var createdZone = *createZoneResp.Zone //nolint:golint // transparency on data model naming
		fmt.Printf("[DNS API] Created zone \"%s\" with DNS name \"%s\" and zone id \"%s\".\n", *createdZone.Name, *createdZone.DnsName, *createdZone.Id)
	}

	// Get a record set of a DNS zone.
	var recordSetResp *dns.RecordSetResponse //nolint:golint // transparency on data model naming
	recordSetResp, err = dnsClient.GetRecordSet(context.Background(), projectId, "zoneId", "recordSetId").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Error when calling `GetRecordSet`: %v\n", err)
	} else {
		fmt.Printf("[DNS API] Got record set with name \"%s\".\n", *recordSetResp.Rrset.Name)
	}
}
