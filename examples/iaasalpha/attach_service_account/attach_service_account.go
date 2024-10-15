package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/runtime"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha/wait"
)

func main() {
	// Specify the organization ID and project ID
	projectId := "PROJECT_ID"
	serverId := "SERVER_ID"
	serviceAccountMail := "SERVICE_ACCOUNT_MAIL"

	// Create a new API client, that uses default authentication and configuration
	iaasalphaClient, err := iaasalpha.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Attach an existing service account to an existing server
	var httpResp *http.Response
	ctxWithHTTPResp := runtime.WithCaptureHTTPResponse(context.Background(), &httpResp)
	_, err = iaasalphaClient.AddServiceAccountToServer(ctxWithHTTPResp, projectId, serverId, serviceAccountMail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `AddServiceAccountToServer`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Triggered attachment of service account with mail %q.\n", serviceAccountMail)
	}
	requestId := httpResp.Header[wait.XRequestIDHeader][0]

	// Wait for attachment of the service account
	_, err = wait.ProjectRequestWaitHandler(context.Background(), iaasalphaClient, projectId, requestId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for attachment: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Service account %q has been successfully attached to the server %s.\n", serviceAccountMail, serverId)

	_, err = iaasalphaClient.RemoveServiceAccountFromServer(ctxWithHTTPResp, projectId, serverId, serviceAccountMail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `RemoveServiceAccountFromServer`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Triggered removal of attachment of service account with mail %q.\n", serviceAccountMail)
	}

	requestId = httpResp.Header[wait.XRequestIDHeader][0]

	// Wait for dettachment of the service account
	_, err = wait.ProjectRequestWaitHandler(context.Background(), iaasalphaClient, projectId, requestId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for removal of attachment of service account: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Service account %q has been successfully detached from the server %s.\n", serviceAccountMail, serverId)
}
