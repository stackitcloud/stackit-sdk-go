package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas"
)

func main() {
	// Specify the organization ID, network area ID and region
	organizationId := "ORGANIZATION_ID"
	networkAreaId := "NETWORK_AREA_ID"
	region := "REGION"

	// Create a new API client, that uses default authentication and configuration
	iaasClient, err := iaas.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Add routing table in network area
	addRoutingTableToAreaPayload := iaas.AddRoutingTableToAreaPayload{
		Name:        utils.Ptr("example-routing-tables"),
		Description: utils.Ptr("example"),
	}

	routingTable, err := iaasClient.AddRoutingTableToArea(context.Background(), organizationId, networkAreaId, region).AddRoutingTableToAreaPayload(addRoutingTableToAreaPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `AddRoutingTableToArea`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[IaaS API] Added routing table %q.\n", *routingTable.Id)

	// Update routing table in network area
	updateRoutingTableOfAreaPayload := iaas.UpdateRoutingTableOfAreaPayload{
		Name:          utils.Ptr("example-routing-tables-updated"),
		DynamicRoutes: utils.Ptr(false),
	}

	updatedRoutingTable, err := iaasClient.UpdateRoutingTableOfArea(context.Background(), organizationId, networkAreaId, region, *routingTable.Id).UpdateRoutingTableOfAreaPayload(updateRoutingTableOfAreaPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `UpdateRoutingTableOfArea`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[IaaS API] Updated routing table %q.\n", *updatedRoutingTable.Id)

	// Add routes in routing table
	addRoutesToRoutingTablePayload := iaas.AddRoutesToRoutingTablePayload{
		Items: &[]iaas.Route{
			{
				Destination: &iaas.RouteDestination{
					DestinationCIDRv4: &iaas.DestinationCIDRv4{
						Type:  utils.Ptr("cidrv4"),
						Value: utils.Ptr("192.168.0.0/24"),
					},
				},
				Labels: &map[string]interface{}{
					"foo": "bar",
				},
				Nexthop: &iaas.RouteNexthop{
					NexthopIPv4: &iaas.NexthopIPv4{
						Type:  utils.Ptr("ipv4"),
						Value: utils.Ptr("10.1.2.10"),
					},
				},
			},
		},
	}

	route, err := iaasClient.AddRoutesToRoutingTable(context.Background(), organizationId, networkAreaId, region, *updatedRoutingTable.Id).AddRoutesToRoutingTablePayload(addRoutesToRoutingTablePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `AddRoutesToRoutingTable`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[IaaS API] Added routes %q to routing table %q.\n", *(*route.Items)[0].Id, *updatedRoutingTable.Id)

	// Delete route in routing table
	err = iaasClient.DeleteRouteFromRoutingTable(context.Background(), organizationId, networkAreaId, region, *updatedRoutingTable.Id, *(*route.Items)[0].Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `DeleteRouteFromRoutingTable`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[IaaS API] Deleted route %q from routing table %q.\n", *(*route.Items)[0].Id, *updatedRoutingTable.Id)

	// Delete routing table in network area
	err = iaasClient.DeleteRoutingTableFromArea(context.Background(), organizationId, networkAreaId, region, *updatedRoutingTable.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `DeleteRoutingTableFromArea`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[IaaS API] Deleted routing table %q.\n", *updatedRoutingTable.Id)
}
