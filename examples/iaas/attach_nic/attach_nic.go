package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/runtime"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas/wait"
)

func main() {
	// Specify the organization ID and project ID
	projectId := "PROJECT_ID"
	serverId := "SERVER_ID"
	nicId := "NIC_ID"

	// Create a new API client, that uses default authentication and configuration
	iaasClient, err := iaas.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Attach an existing network interface to an existing server
	var httpResp *http.Response
	ctxWithHTTPResp := runtime.WithCaptureHTTPResponse(context.Background(), &httpResp)
	err = iaasClient.AddNICToServer(ctxWithHTTPResp, projectId, serverId, nicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `AddNICToServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered attachment of nic with ID %q.\n", nicId)
	}
	requestId := httpResp.Header[wait.XRequestIDHeader][0]

	// Wait for attachment of the nic
	_, err = wait.ProjectRequestWaitHandler(context.Background(), iaasClient, projectId, requestId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for attachment: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Nic %q has been successfully attached to the server %s.\n", nicId, serverId)

	err = iaasClient.RemoveNICFromServer(ctxWithHTTPResp, projectId, serverId, nicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `RemoveNICFromServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered removal of attachment of nic with ID %q.\n", nicId)
	}

	requestId = httpResp.Header[wait.XRequestIDHeader][0]

	// Wait for dettachment of the nic
	_, err = wait.ProjectRequestWaitHandler(context.Background(), iaasClient, projectId, requestId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for removal of attachment of NIC: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] NIC %q has been successfully detached from the server %s.\n", nicId, serverId)
}
