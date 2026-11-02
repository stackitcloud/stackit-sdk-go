package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	iaas "github.com/stackitcloud/stackit-sdk-go/services/iaas/v2api"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas/v2api/wait"
)

func main() {
	// Specify the organization ID and region
	organizationId := "ORGANIZATION_ID"
	region := "eu01"

	// Create a new API client, that uses default authentication and configuration
	iaasClient, err := iaas.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// List the network areas for your organization
	var areas *iaas.NetworkAreaListResponse //nolint:golint // transparency on data model naming
	areas, err = iaasClient.DefaultAPI.ListNetworkAreas(context.Background(), organizationId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `ListNetworkAreas`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[IaaS API] Number of network areas: %v\n", len(areas.Items))

	// Create a network area
	createNetworkAreaPayload := iaas.CreateNetworkAreaPayload{
		Name: "example-network-area",
		Labels: map[string]interface{}{
			"key": "value",
		},
	}
	area, err := iaasClient.DefaultAPI.CreateNetworkArea(context.Background(), organizationId).CreateNetworkAreaPayload(createNetworkAreaPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `CreateNetworkAreas`: %v\n", err)
	} else {
		fmt.Printf("[IaaS API] Created network area with ID %q.\n", *area.Id)
	}

	fmt.Printf("[IaaS API] Network area %q has been successfully created.\n", *area.Id)

	// Update a network area
	updateNetworkAreaPayload := iaas.PartialUpdateNetworkAreaPayload{
		Name: utils.Ptr(area.Name + "-renamed"),
	}
	updatedArea, err := iaasClient.DefaultAPI.PartialUpdateNetworkArea(context.Background(), organizationId, *area.Id).PartialUpdateNetworkAreaPayload(updateNetworkAreaPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `PartialUpdateNetworkArea`: %v\n", err)
	} else {
		fmt.Printf("[IaaS API] Updated network area with ID %q.\n", *updatedArea.Id)
	}

	fmt.Printf("[IaaS API] Network area %q has been successfully updated.\n", *updatedArea.Id)

	// Create a network area region
	createNetworkAreaRegionPayload := iaas.CreateNetworkAreaRegionPayload{
		Ipv4: &iaas.RegionalAreaIPv4{
			DefaultNameservers: []string{
				"1.2.3.4",
			},
			DefaultPrefixLen: int64(25),
			MaxPrefixLen:     int64(29),
			MinPrefixLen:     int64(24),
			NetworkRanges: []iaas.NetworkRange{
				{
					Prefix: "1.2.3.0/24",
				},
			},
			TransferNetwork: "1.2.4.0/24",
		},
	}

	resp, err := iaasClient.DefaultAPI.CreateNetworkAreaRegion(context.Background(), organizationId, *updatedArea.Id, region).CreateNetworkAreaRegionPayload(createNetworkAreaRegionPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `ConfigureNetworkAreaRegion`: %v\n", err)
	} else {
		fmt.Printf("[IaaS API] Triggered configuration of network area with ID %q.\n", *updatedArea.Id)
	}
	fmt.Printf("[IaaS API] Status of updated network area config %q.\n", *resp.Status)

	_, err = wait.CreateNetworkAreaRegionWaitHandler(context.Background(), iaasClient.DefaultAPI, organizationId, *updatedArea.Id, region).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when waiting for network configuration creation: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[IaaS API] Network area %q configuration for region %q has been successfully created.\n", *updatedArea.Id, region)

	// Update a network area configuration
	updateConfigNetworkAreaPayload := iaas.UpdateNetworkAreaRegionPayload{
		Ipv4: &iaas.UpdateRegionalAreaIPv4{
			DefaultNameservers: []string{
				"2.2.3.4",
			},
			DefaultPrefixLen: utils.Ptr(int64(26)),
		},
	}

	_, err = iaasClient.DefaultAPI.UpdateNetworkAreaRegion(context.Background(), organizationId, *updatedArea.Id, region).UpdateNetworkAreaRegionPayload(updateConfigNetworkAreaPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `UpdateNetworkAreaRegion`: %v\n", err)
	} else {
		fmt.Printf("[IaaS API] Network area %q configuration for region %q has been successfully updated.\n", *updatedArea.Id, region)
	}

	// Delete regional network area
	err = iaasClient.DefaultAPI.DeleteNetworkAreaRegion(context.Background(), organizationId, *updatedArea.Id, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `DeleteNetworkAreaRegion`: %v\n", err)
	}

	_, err = wait.DeleteNetworkAreaRegionWaitHandler(context.Background(), iaasClient.DefaultAPI, organizationId, *updatedArea.Id, region).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when waiting for network configuration deletion: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[IaaS API] Network area %q configuration for region %q has been successfully deleted.\n", *updatedArea.Id, region)

	// Delete a network area
	err = iaasClient.DefaultAPI.DeleteNetworkArea(context.Background(), organizationId, *updatedArea.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `DeleteNetworkArea`: %v\n", err)
	} else {
		fmt.Printf("[IaaS API] Network area %q has been successfully deleted.\n", *updatedArea.Id)
	}
}
