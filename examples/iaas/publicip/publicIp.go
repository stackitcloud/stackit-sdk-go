package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas"
)

func main() {
	// Specify the project ID, network interface ID and region
	projectId := "PROJECT_ID"
	networkInterfaceId := "NETWORK_INTERFACE_ID"
	region := "REGION"

	// Create a new API client, that uses default authentication and configuration
	iaasClient, err := iaas.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	publicIps, err := iaasClient.ListPublicIPs(context.Background(), projectId, region).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `ListPublicIPs`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] Number of Public IPs: %v\n", len(*publicIps.Items))
	}

	// Create a publicIp
	createpublicIpPayload := iaas.CreatePublicIPPayload{
		NetworkInterface: iaas.NewNullableString(utils.Ptr(networkInterfaceId)),
	}
	publicIp, err := iaasClient.CreatePublicIP(context.Background(), projectId, region).CreatePublicIPPayload(createpublicIpPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `CreatePublicIP`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] public IP %q has been successfully created.\n", *publicIp.Id)
	}

	// Update a publicIp
	updatepublicIpPayload := iaas.UpdatePublicIPPayload{
		NetworkInterface: iaas.NewNullableString(nil),
	}
	publicIp, err = iaasClient.UpdatePublicIP(context.Background(), projectId, region, *publicIp.Id).UpdatePublicIPPayload(updatepublicIpPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `UpdatePublicIP`: %v\n", err)
	}

	fmt.Printf("[iaas API] public IP %q has been successfully updated.\n", *publicIp.Id)
	if publicIp.NetworkInterface == nil {
		fmt.Printf("[iaas API] Public IP network interface has been successfully removed.\n")
	} else {
		fmt.Fprintf(os.Stderr, "[iaas API] Public IP network interface has not been removed.\n")
	}

	// Delete a public IP
	err = iaasClient.DeletePublicIP(context.Background(), projectId, region, *publicIp.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[iaas API] Error when calling `DeletepublicIp`: %v\n", err)
	} else {
		fmt.Printf("[iaas API] public IP %q has been successfully deleted.\n", *publicIp.Id)
	}
}
