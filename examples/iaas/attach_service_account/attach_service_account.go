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
	serviceAccountMail := "SERVICE_ACCOUNT_MAIL"

	// Create a new API client, that uses default authentication and configuration
	iaasClient, err := iaas.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Attach an existing service account to an existing server
	var httpResp *http.Response
	ctxWithHTTPResp := runtime.WithCaptureHTTPResponse(context.Background(), &httpResp)
	_, err = iaasClient.AddServiceAccountToServer(ctxWithHTTPResp, projectId, serverId, serviceAccountMail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `AddServiceAccountToServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered attachment of service account with mail %q.\n", serviceAccountMail)
	}
	requestId := httpResp.Header[wait.XRequestIDHeader][0]

	// Wait for attachment of the service account
	_, err = wait.ProjectRequestWaitHandler(context.Background(), iaasClient, projectId, requestId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for attachment: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Service account %q has been successfully attached to the server %s.\n", serviceAccountMail, serverId)

	_, err = iaasClient.RemoveServiceAccountFromServer(ctxWithHTTPResp, projectId, serverId, serviceAccountMail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `RemoveServiceAccountFromServer`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Triggered removal of attachment of service account with mail %q.\n", serviceAccountMail)
	}

	requestId = httpResp.Header[wait.XRequestIDHeader][0]

	// Wait for dettachment of the service account
	_, err = wait.ProjectRequestWaitHandler(context.Background(), iaasClient, projectId, requestId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when waiting for removal of attachment of service account: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaas API] Service account %q has been successfully detached from the server %s.\n", serviceAccountMail, serverId)
}
