package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha"
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

	publicIps, err := iaasalphaClient.ListPublicIPs(context.Background(), projectId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `ListPublicIPs`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] Number of Public IPs: %v\n", len(*publicIps.Items))
	}

	// Create a publicIp
	createpublicIpPayload := iaasalpha.CreatePublicIPPayload{
		NetworkInterface: iaasalpha.NewNullableString(utils.Ptr("NIC_ID")),
	}
	publicIp, err := iaasalphaClient.CreatePublicIP(context.Background(), projectId).CreatePublicIPPayload(createpublicIpPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `CreatePublicIP`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] public IP %q has been successfully created.\n", *publicIp.Id)
	}

	// Update a publicIp
	updatepublicIpPayload := iaasalpha.UpdatePublicIPPayload{
		NetworkInterface: iaasalpha.NewNullableString(nil),
	}
	publicIp, err = iaasalphaClient.UpdatePublicIP(context.Background(), projectId, *publicIp.Id).UpdatePublicIPPayload(updatepublicIpPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `UpdatePublicIP`: %v\n", err)
	}

	fmt.Printf("[iaasalpha API] public IP %q has been successfully updated.\n", *publicIp.Id)
	if publicIp.NetworkInterface == nil {
		fmt.Printf("[iaasalpha API] Public IP network interface has been successfully removed.\n")
	} else {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Public IP network interface has not been removed.\n")
	}

	// Delete a public IP
	err = iaasalphaClient.DeletePublicIP(context.Background(), projectId, *publicIp.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaasalpha API] Error when calling `DeletepublicIp`: %v\n", err)
	} else {
		fmt.Printf("[iaasalpha API] public IP %q has been successfully deleted.\n", *publicIp.Id)
	}
}
