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
	securityGroupId := "SECURITY_GROUP_ID"

	// Create a new API client, that uses default authentication and configuration
	iaasalphaClient, err := iaasalpha.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Attach an existing network interface to an existing server
	var httpResp *http.Response
	ctxWithHTTPResp := runtime.WithCaptureHTTPResponse(context.Background(), &httpResp)
	err = iaasalphaClient.AddSecurityGroupToServer(ctxWithHTTPResp, projectId, serverId, securityGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `AddSecurityGroupToServer`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Triggered attachment of security group with ID %q.\n", securityGroupId)
	}
	requestId := httpResp.Header["wait.XRequestIDHeader"][0]

	// Wait for attachment of the security group
	_, err = wait.ProjectRequestWaitHandler(context.Background(), iaasalphaClient, projectId, requestId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for attachment: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] Security group %q has been successfully attached to the server %s.\n", securityGroupId, serverId)

	err = iaasalphaClient.RemoveSecurityGroupFromServer(ctxWithHTTPResp, projectId, serverId, securityGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `RemoveSecurityGroupFromServer`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Triggered removal of attachment of security group with ID %q.\n", securityGroupId)
	}

	requestId = httpResp.Header["wait.XRequestIDHeader"][0]

	// Wait for dettachment of the security group
	_, err = wait.ProjectRequestWaitHandler(context.Background(), iaasalphaClient, projectId, requestId).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when waiting for removal of attachment of SecurityGroup: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[iaasalpha API] SecurityGroup %q has been successfully detacched from the server %s.\n", securityGroupId, serverId)
}
