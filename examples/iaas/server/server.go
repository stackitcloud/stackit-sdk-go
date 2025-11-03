package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas/wait"
)

func main() {
	// Specify the project ID, image ID and region
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	imageId := "IMAGE_ID"
	region := "eu01"

	// Create a new API client, that uses default authentication and configuration
	iaasClient, err := iaas.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	servers, err := iaasClient.ListServers(context.Background(), projectId, region).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `ListServers`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Number of servers: %v\n", len(*servers.Items))
	}

	// Create a network
	createNetworkPayload := iaas.CreateNetworkPayload{
		Name: utils.Ptr("example-network"),
		Ipv4: &iaas.CreateNetworkIPv4{
			CreateNetworkIPv4WithPrefixLength: &iaas.CreateNetworkIPv4WithPrefixLength{
				PrefixLength: utils.Ptr(int64(24)),
			},
		},
	}

	network, err := iaasClient.CreateNetwork(context.Background(), projectId, region).CreateNetworkPayload(createNetworkPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `CreateNetwork`: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("[Iaas API] Waiting for network to be created...")

	network, err = wait.CreateNetworkWaitHandler(context.Background(), iaasClient, projectId, region, *network.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}

	// Create a server
	createServerPayload := iaas.CreateServerPayload{
		Name:             utils.Ptr("example-server"),
		AvailabilityZone: utils.Ptr("eu01-1"),
		MachineType:      utils.Ptr("g1.1"),
		BootVolume: &iaas.ServerBootVolume{
			Size: utils.Ptr(int64(64)),
			Source: &iaas.BootVolumeSource{
				Id:   utils.Ptr(imageId),
				Type: utils.Ptr("image"),
			},
		},
		Networking: &iaas.CreateServerPayloadAllOfNetworking{
			CreateServerNetworking: &iaas.CreateServerNetworking{
				NetworkId: network.Id,
			},
		},
	}
	server, err := iaasClient.CreateServer(context.Background(), projectId, region).CreateServerPayload(createServerPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `CreateServer`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[iaas API] Triggered creation of server with ID %q.\n", *server.Id)

	// Wait for creation of the server
	server, err = wait.CreateServerWaitHandler(context.Background(), iaasClient, projectId, region, *server.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Server %q has been successfully created.\n", *server.Id)

	// Stop a server
	err = iaasClient.StopServer(context.Background(), projectId, region, *server.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `StopServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered stop of server with ID %q.\n", *server.Id)
	}

	// Wait for stop of the server
	server, err = wait.StopServerWaitHandler(context.Background(), iaasClient, projectId, region, *server.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for stop: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Server %q has been successfully stopped.\n", *server.Id)

	// Start a server
	err = iaasClient.StartServer(context.Background(), projectId, region, *server.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `StartServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered start of server with ID %q.\n", *server.Id)
	}

	// Wait for start of the server
	server, err = wait.StartServerWaitHandler(context.Background(), iaasClient, projectId, region, *server.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for start: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Server %q has been successfully started.\n", *server.Id)

	// Update a server
	updateServerPayload := iaas.UpdateServerPayload{
		Name: utils.Ptr("renamed"),
	}
	server, err = iaasClient.UpdateServer(context.Background(), projectId, region, *server.Id).UpdateServerPayload(updateServerPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `UpdateServer`: %v\n", err)
	}

	fmt.Printf("[iaas API] Server %q has been successfully updated.\n", *server.Id)

	// Resize a server
	resizeServerPayload := iaas.ResizeServerPayload{
		MachineType: utils.Ptr("c1.2"),
	}

	err = iaasClient.ResizeServer(context.Background(), projectId, region, *server.Id).ResizeServerPayload(resizeServerPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `ResizeServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered resize of server with ID %q.\n", *server.Id)
	}

	server, err = wait.ResizeServerWaitHandler(context.Background(), iaasClient, projectId, region, *server.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for resize: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Server %q has been successfully resized.\n", *server.Id)

	// Delete a server
	err = iaasClient.DeleteServer(context.Background(), projectId, region, *server.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `DeleteServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered deletion of server with ID %q.\n", *server.Id)
	}

	// Wait for deletion of the server
	_, err = wait.DeleteServerWaitHandler(context.Background(), iaasClient, projectId, region, *server.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Server %q has been successfully deleted.\n", *server.Id)

	// Delete a network
	err = iaasClient.DeleteNetwork(context.Background(), projectId, region, *network.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when calling `DeleteNetwork`: %v\n", err)
		os.Exit(1)
	}

	_, err = wait.DeleteNetworkWaitHandler(context.Background(), iaasClient, projectId, region, *network.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[IaaS API] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[IaaS API] Network has been successfully deleted.\n")
}
