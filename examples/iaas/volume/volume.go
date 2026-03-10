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
	// Specify the project ID and region
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	region := "eu01"

	// Create a new API client, that uses default authentication and configuration
	iaasClient, err := iaas.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	volumes, err := iaasClient.DefaultAPI.ListVolumes(context.Background(), projectId, region).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `ListVolumes`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Number of volumes: %v\n", len(volumes.Items))
	}

	// Create a volume
	createVolumePayload := iaas.CreateVolumePayload{
		Name:             utils.Ptr("example-volume"),
		AvailabilityZone: "eu01-1",
		Size:             utils.Ptr(int64(10)),
	}
	volume, err := iaasClient.DefaultAPI.CreateVolume(context.Background(), projectId, region).CreateVolumePayload(createVolumePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `CreateVolume`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered creation of volume with ID %q.\n", *volume.Id)
	}

	// Wait for creation of the volume
	volume, err = wait.CreateVolumeWaitHandler(context.Background(), iaasClient.DefaultAPI, projectId, region, *volume.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Volume %q has been successfully created.\n", *volume.Id)

	// Update a volume
	updateVolumePayload := iaas.UpdateVolumePayload{
		Name: utils.Ptr("renamed"),
	}
	volume, err = iaasClient.DefaultAPI.UpdateVolume(context.Background(), projectId, region, *volume.Id).UpdateVolumePayload(updateVolumePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `UpdateVolume`: %v\n", err)
	}

	fmt.Printf("[iaas API] Volume %q has been successfully updated.\n", *volume.Id)

	// Resize a volume
	resizeVolumePayload := iaas.ResizeVolumePayload{
		Size: int64(130),
	}
	err = iaasClient.DefaultAPI.ResizeVolume(context.Background(), projectId, region, *volume.Id).ResizeVolumePayload(resizeVolumePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `ResizeVolume`: %v\n", err)
	}

	fmt.Printf("[iaas API] Volume %q has been successfully resized.\n", *volume.Id)

	// Delete a volume
	err = iaasClient.DefaultAPI.DeleteVolume(context.Background(), projectId, region, *volume.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `DeleteVolume`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered deletion of volume with ID %q.\n", *volume.Id)
	}

	// Wait for deletion of the volume
	_, err = wait.DeleteVolumeWaitHandler(context.Background(), iaasClient.DefaultAPI, projectId, region, *volume.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Volume %q has been successfully deleted.\n", *volume.Id)
}
