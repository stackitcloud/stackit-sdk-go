package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha/wait"
)

func main() {
	// Specify the organization ID and project ID
	projectId := "PROJECT_ID"
	serverId := "SERVER_ID"
	volumeId := "VOLUME_ID"

	// Create a new API client, that uses default authentication and configuration
	iaasalphaClient, err := iaasalpha.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	_, err = iaasalphaClient.AddVolumeToServer(context.Background(), projectId, serverId, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `AddVolumeToServer`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Triggered attachment of volume with ID %q.\n", volumeId)
	}

	// Wait for attachment of the volume
	_, err = wait.AddVolumeToServerWaitHandler(context.Background(), iaasalphaClient, projectId, serverId, volumeId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for attachment: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Volume %q has been successfully attached to the server %s.\n", volumeId, serverId)

	err = iaasalphaClient.RemoveVolumeFromServer(context.Background(), projectId, serverId, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `RemoveVolumeFromServer`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Triggered removal of attachment of volume with ID %q.\n", volumeId)
	}

	// Wait for dettachment of the volume
	_, err = wait.RemoveVolumeFromServerWaitHandler(context.Background(), iaasalphaClient, projectId, serverId, volumeId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for removal of attachment of volume: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Volume %q has been successfully detacched from the server %s.\n", volumeId, serverId)
}
