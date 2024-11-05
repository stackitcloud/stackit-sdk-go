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

	volumes, err := iaasClient.ListVolumes(context.Background(), projectId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `ListVolumes`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Number of volumes: %v\n", len(*volumes.Items))
	}

	// Create a volume
	createVolumePayload := iaas.CreateVolumePayload{
		Name:             utils.Ptr("example-volume"),
		AvailabilityZone: utils.Ptr("eu01-1"),
		Size:             utils.Ptr(int64(10)),
	}
	volume, err := iaasClient.CreateVolume(context.Background(), projectId).CreateVolumePayload(createVolumePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `CreateVolume`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered creation of volume with ID %q.\n", *volume.Id)
	}

	// Wait for creation of the volume
	volume, err = wait.CreateVolumeWaitHandler(context.Background(), iaasClient, projectId, *volume.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Volume %q has been successfully created.\n", *volume.Id)

	// Update a volume
	updateVolumePayload := iaas.UpdateVolumePayload{
		Name: utils.Ptr("renamed"),
	}
	volume, err = iaasClient.UpdateVolume(context.Background(), projectId, *volume.Id).UpdateVolumePayload(updateVolumePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `UpdateVolume`: %v\n", err)
	}

	fmt.Printf("[iaas API] Volume %q has been successfully updated.\n", *volume.Id)

	// Resize a volume
	resizeVolumePayload := iaas.ResizeVolumePayload{
		Size: utils.Ptr(int64(130)),
	}
	err = iaasClient.ResizeVolume(context.Background(), projectId, *volume.Id).ResizeVolumePayload(resizeVolumePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `ResizeVolume`: %v\n", err)
	}

	fmt.Printf("[iaas API] Volume %q has been successfully resized.\n", *volume.Id)

	// Delete a volume
	err = iaasClient.DeleteVolume(context.Background(), projectId, *volume.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `DeleteVolume`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered deletion of volume with ID %q.\n", *volume.Id)
	}

	// Wait for deletion of the volume
	_, err = wait.DeleteVolumeWaitHandler(context.Background(), iaasClient, projectId, *volume.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Volume %q has been successfully deleted.\n", *volume.Id)
}
