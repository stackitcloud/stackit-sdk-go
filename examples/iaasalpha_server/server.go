package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha/wait"
)

func main() {
	// Specify the organization ID and project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	iaasalphaClient, err := iaasalpha.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	servers, err := iaasalphaClient.ListServers(context.Background(), projectId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `ListServers`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Number of servers: %v\n", len(*servers.Items))
	}

	// Create a server
	createServerPayload := iaasalpha.CreateServerPayload{
		Name:             utils.Ptr("example-server"),
		AvailabilityZone: utils.Ptr("eu01-1"),
		MachineType:      utils.Ptr("g1.1"),
		BootVolume: &iaasalpha.CreateServerPayloadBootVolume{
			Size: utils.Ptr(int64(64)),
			Source: &iaasalpha.BootVolumeSource{
				Id:   utils.Ptr("IMAGE_ID"),
				Type: utils.Ptr("image"),
			},
		},
	}
	server, err := iaasalphaClient.CreateServer(context.Background(), projectId).CreateServerPayload(createServerPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `CreateServer`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Triggered creation of server with ID %q.\n", *server.Id)
	}

	// Wait for creation of the server
	server, err = wait.CreateServerWaitHandler(context.Background(), iaasalphaClient, projectId, *server.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Server %q has been successfully created.\n", *server.Id)

	// Update a server
	updateServerPayload := iaasalpha.V1alpha1UpdateServerPayload{
		Name: utils.Ptr("renamed"),
	}
	server, err = iaasalphaClient.V1alpha1UpdateServer(context.Background(), projectId, *server.Id).V1alpha1UpdateServerPayload(updateServerPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `UpdateServer`: %v\n", err)
	}

	fmt.Printf("[iaasalpha API] Server %q has been successfully updated.\n", *server.Id)

	// Resize a server
	resizeServerPayload := iaasalpha.ResizeServerPayload{
		MachineType: utils.Ptr("c1.2"),
	}

	err = iaasalphaClient.ResizeServer(context.Background(), projectId, *server.Id).ResizeServerPayload(resizeServerPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `ResizeServer`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Triggered resize of server with ID %q.\n", *server.Id)
	}

	server, err = wait.ResizeServerWaitHandler(context.Background(), iaasalphaClient, projectId, *server.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for resize: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Server %q has been successfully resized.\n", *server.Id)

	// Delete a server
	err = iaasalphaClient.DeleteServer(context.Background(), projectId, *server.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `DeleteServer`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Triggered deletion of server with ID %q.\n", *server.Id)
	}

	// Wait for deletion of the server
	_, err = wait.DeleteServerWaitHandler(context.Background(), iaasalphaClient, projectId, *server.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Server %q has been successfully deleted.\n", *server.Id)
}
