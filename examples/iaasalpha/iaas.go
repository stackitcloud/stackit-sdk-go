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

	volumes, err := iaasalphaClient.ListVolumes(context.Background(), projectId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `ListVolumes`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Number of volumes: %v\n", len(*volumes.Items))
	}

	// Create a volume
	createVolumePayload := iaasalpha.CreateVolumePayload{
		Name:             utils.Ptr("example-volume"),
		AvailabilityZone: utils.Ptr("eu01-1"),
		Size:             utils.Ptr(int64(10)),
	}
	volume, err := iaasalphaClient.CreateVolume(context.Background(), projectId).CreateVolumePayload(createVolumePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `CreateVolume`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Triggered creation of volume with ID %q.\n", *volume.Id)
	}

	// Wait for creation of the volume
	volume, err = wait.CreateVolumeWaitHandler(context.Background(), iaasalphaClient, projectId, *volume.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Network volume %q has been successfully created.\n", *volume.Id)

	// Update a volume
	updateVolumePayload := iaasalpha.UpdateVolumePayload{
		Name: utils.Ptr("renamed"),
	}
	volume, err = iaasalphaClient.UpdateVolume(context.Background(), projectId, *volume.Id).UpdateVolumePayload(updateVolumePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `UpdateVolume`: %v\n", err)
	}

	fmt.Printf("[iaasalpha API] Network volume %q has been successfully updated.\n", *volume.Id)

	// Resize a volume
	resizeVolumePayload := iaasalpha.ResizeVolumePayload{
		Size: utils.Ptr(int64(130)),
	}
	err = iaasalphaClient.ResizeVolume(context.Background(), projectId, *volume.Id).ResizeVolumePayload(resizeVolumePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `ResizeVolume`: %v\n", err)
	}

	fmt.Printf("[iaasalpha API] Network volume %q has been successfully resized.\n", *volume.Id)

	// Delete a volume
	err = iaasalphaClient.DeleteVolume(context.Background(), projectId, *volume.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `DeleteVolume`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Triggered deletion of volume with ID %q.\n", *volume.Id)
	}

	// Wait for deletion of the volume
	_, err = wait.DeleteVolumeWaitHandler(context.Background(), iaasalphaClient, projectId, *volume.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Network volume %q has been successfully deleted.\n", *volume.Id)
}
