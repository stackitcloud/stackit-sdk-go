package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/runtime"
	iaas "github.com/stackitcloud/stackit-sdk-go/services/iaas/v2api"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas/v2api/wait"
)

func main() {
	// Specify the project ID, server ID, volume ID and region
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	serverId := "SERVER_ID"
	volumeId := "VOLUME_ID"
	region := "eu01"

	// Create a new API client, that uses default authentication and configuration
	iaasClient, err := iaas.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	var httpResp *http.Response
	ctx := runtime.WithCaptureHTTPResponse(context.Background(), &httpResp)

	payload := iaas.AddVolumeToServerPayload{}
	_, err = iaasClient.DefaultAPI.AddVolumeToServer(ctx, projectId, region, serverId, volumeId).AddVolumeToServerPayload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `AddVolumeToServer`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[iaas API] Triggered attachment of volume with ID %q.\n", volumeId)

	xRequestId := httpResp.Header.Get(wait.XRequestIDHeader)
	// Wait for attachment of the volume
	_, err = wait.ProjectRequestWaitHandler(ctx, iaasClient.DefaultAPI, projectId, region, xRequestId).SetSleepBeforeWait(500 * time.Millisecond).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for attachment: %v\n", err)
		os.Exit(1)
	}

	err = iaasClient.DefaultAPI.RemoveVolumeFromServer(ctx, projectId, region, serverId, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `RemoveVolumeFromServer`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[iaas API] Triggered removal of attachment of volume with ID %q.\n", volumeId)

	xRequestId = httpResp.Header.Get(wait.XRequestIDHeader)
	// Wait for detachment of the volume
	_, err = wait.ProjectRequestWaitHandler(ctx, iaasClient.DefaultAPI, projectId, region, xRequestId).SetSleepBeforeWait(500 * time.Millisecond).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for removal of attachment of volume: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Volume %q has been successfully detached from the server %s.\n", volumeId, serverId)
}
