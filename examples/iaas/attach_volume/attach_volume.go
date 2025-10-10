package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/services/iaas"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas/wait"
)

func main() {
	// Specify the project ID, server ID, volume ID and region
	projectId := "PROJECT_ID"
	serverId := "SERVER_ID"
	volumeId := "VOLUME_ID"
	region := "REGION"

	// Create a new API client, that uses default authentication and configuration
	iaasClient, err := iaas.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	payload := iaas.AddVolumeToServerPayload{}
	_, err = iaasClient.AddVolumeToServer(context.Background(), projectId, region, serverId, volumeId).AddVolumeToServerPayload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `AddVolumeToServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered attachment of volume with ID %q.\n", volumeId)
	}

	// Wait for attachment of the volume
	_, err = wait.AddVolumeToServerWaitHandler(context.Background(), iaasClient, projectId, region, serverId, volumeId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for attachment: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Volume %q has been successfully attached to the server %s.\n", volumeId, serverId)

	err = iaasClient.RemoveVolumeFromServer(context.Background(), projectId, region, serverId, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `RemoveVolumeFromServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered removal of attachment of volume with ID %q.\n", volumeId)
	}

	// Wait for dettachment of the volume
	_, err = wait.RemoveVolumeFromServerWaitHandler(context.Background(), iaasClient, projectId, region, serverId, volumeId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for removal of attachment of volume: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Volume %q has been successfully detached from the server %s.\n", volumeId, serverId)
}
