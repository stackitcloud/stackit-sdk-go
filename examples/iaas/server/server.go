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
		fmt.Fprintf(os.Stderr, "[iaas API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	servers, err := iaasClient.ListServers(context.Background(), projectId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `ListServers`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Number of servers: %v\n", len(*servers.Items))
	}

	// Create a server
	createServerPayload := iaas.CreateServerPayload{
		Name:             utils.Ptr("example-server"),
		AvailabilityZone: utils.Ptr("eu01-1"),
		MachineType:      utils.Ptr("g1.1"),
		BootVolume: &iaas.CreateServerPayloadBootVolume{
			Size: utils.Ptr(int64(64)),
			Source: &iaas.BootVolumeSource{
				Id:   utils.Ptr("IMAGE_ID"),
				Type: utils.Ptr("image"),
			},
		},
	}
	server, err := iaasClient.CreateServer(context.Background(), projectId).CreateServerPayload(createServerPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `CreateServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered creation of server with ID %q.\n", *server.Id)
	}

	// Wait for creation of the server
	server, err = wait.CreateServerWaitHandler(context.Background(), iaasClient, projectId, *server.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Server %q has been successfully created.\n", *server.Id)

	// Stop a server
	err = iaasClient.StopServer(context.Background(), projectId, *server.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `StopServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered stop of server with ID %q.\n", *server.Id)
	}

	// Wait for stop of the server
	server, err = wait.StopServerWaitHandler(context.Background(), iaasClient, projectId, *server.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for stop: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Server %q has been successfully stopped.\n", *server.Id)

	// Start a server
	err = iaasClient.StartServer(context.Background(), projectId, *server.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `StartServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered start of server with ID %q.\n", *server.Id)
	}

	// Wait for start of the server
	server, err = wait.StartServerWaitHandler(context.Background(), iaasClient, projectId, *server.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for start: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Server %q has been successfully started.\n", *server.Id)

	// Update a server
	updateServerPayload := iaas.UpdateServerPayload{
		Name: utils.Ptr("renamed"),
	}
	server, err = iaasClient.UpdateServer(context.Background(), projectId, *server.Id).UpdateServerPayload(updateServerPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `UpdateServer`: %v\n", err)
	}

	fmt.Printf("[iaas API] Server %q has been successfully updated.\n", *server.Id)

	// Resize a server
	resizeServerPayload := iaas.ResizeServerPayload{
		MachineType: utils.Ptr("c1.2"),
	}

	err = iaasClient.ResizeServer(context.Background(), projectId, *server.Id).ResizeServerPayload(resizeServerPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `ResizeServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered resize of server with ID %q.\n", *server.Id)
	}

	server, err = wait.ResizeServerWaitHandler(context.Background(), iaasClient, projectId, *server.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for resize: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Server %q has been successfully resized.\n", *server.Id)

	// Delete a server
	err = iaasClient.DeleteServer(context.Background(), projectId, *server.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `DeleteServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered deletion of server with ID %q.\n", *server.Id)
	}

	// Wait for deletion of the server
	_, err = wait.DeleteServerWaitHandler(context.Background(), iaasClient, projectId, *server.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Server %q has been successfully deleted.\n", *server.Id)
}
