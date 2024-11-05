package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas/wait"
)

func main() {
	// Specify the organization ID and project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	iaasClient, err := iaas.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create a network
	createNetworkPayload := iaas.CreateNetworkPayload{
		Name: utils.Ptr("example-network"),
		AddressFamily: &iaas.CreateNetworkAddressFamily{
			Ipv4: &iaas.CreateNetworkIPv4Body{
				PrefixLength: utils.Ptr(int64(24)),
				Nameservers:  &[]string{"1.2.3.4"},
			},
		},
	}

	network, err := iaasClient.CreateNetwork(context.Background(), projectId).CreateNetworkPayload(createNetworkPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `CreateNetwork`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[IaaS API] Triggered creation of network with ID %q.\n", *network.NetworkId)
	fmt.Printf("[Iaas API] Current state of the network: %q\n", *network.State)
	fmt.Println("[Iaas API] Waiting for network to be created...")

	network, err = wait.CreateNetworkWaitHandler(context.Background(), iaasClient, projectId, *network.NetworkId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[IaaS API] Network has been successfully created.\n")

	// Update a network
	updateNetworkPayload := iaas.PartialUpdateNetworkPayload{
		Name: utils.Ptr("example-network-test-renamed"),
	}

	err = iaasClient.PartialUpdateNetwork(context.Background(), projectId, *network.NetworkId).PartialUpdateNetworkPayload(updateNetworkPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `PartialUpdateNetwork`: %v\n", err)
		os.Exit(1)
	}

	_, err = wait.UpdateNetworkWaitHandler(context.Background(), iaasClient, projectId, *network.NetworkId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when waiting for update: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[IaaS API] Network has been successfully updated.\n")

	// Delete a network
	err = iaasClient.DeleteNetwork(context.Background(), projectId, *network.NetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `DeleteNetwork`: %v\n", err)
		os.Exit(1)
	}

	_, err = wait.DeleteNetworkWaitHandler(context.Background(), iaasClient, projectId, *network.NetworkId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[IaaS API] Network has been successfully deleted.\n")
}
