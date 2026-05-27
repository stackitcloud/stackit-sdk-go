package main

import (
	"context"
	"fmt"
	"os"

	vpn "github.com/stackitcloud/stackit-sdk-go/services/vpn/v1api"
)

func main() {
	region := "eu01"          // Region where the resources will be created
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	planId := "PLAN_ID"       // the id of the plan you want to use for the VPN Gateway. You can get the available plans with `ListPlans`.

	// STACKIT VPN enforces the following requirements for a secure PSK:
	// - must be at least 20 characters long
	// - must be at least 16 different characters
	// - must have at least one upper case letter
	// - must have at least one lower case letter
	// - must have at least one number
	psk := "Super.$ecret_Shared3Key12345"

	// Create a new API client, that uses default authentication and configuration
	vpnClient, err := vpn.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create a VPN Gateway
	createVpnGatewayPayload := vpn.CreateGatewayPayload{
		DisplayName: "exampleVpnGateway",
		PlanId:      planId,
		RoutingType: vpn.ROUTINGTYPE_ROUTE_BASED,
		AvailabilityZones: vpn.CreateGatewayPayloadAvailabilityZones{
			Tunnel1: "eu01-1",
			Tunnel2: "eu01-2",
		},
	}

	gatewayResp, err := vpnClient.DefaultAPI.CreateGateway(context.Background(), projectId, region).CreateGatewayPayload(createVpnGatewayPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateVpnGateway`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created VPN Gateway with id \"%s\".\n", *gatewayResp.Id)

	// Create a VPN Connection
	phase1 := vpn.TunnelConfigurationPhase1{
		DhGroups:             []vpn.PhaseDhGroupsInner{"ecp384"},
		EncryptionAlgorithms: []vpn.PhaseEncryptionAlgorithmsInner{"aes256"},
		IntegrityAlgorithms:  []vpn.PhaseIntegrityAlgorithmsInner{"sha2_384"},
	}
	phase2 := vpn.TunnelConfigurationPhase2{
		DhGroups:             []vpn.PhaseDhGroupsInner{"ecp384"},
		EncryptionAlgorithms: []vpn.PhaseEncryptionAlgorithmsInner{"aes256"},
		IntegrityAlgorithms:  []vpn.PhaseIntegrityAlgorithmsInner{"sha2_384"},
	}
	tunnel := vpn.TunnelConfiguration{
		Phase1:        phase1,
		Phase2:        phase2,
		PreSharedKey:  &psk,
		RemoteAddress: "0.0.0.0",
	}

	createGatewayConnectionPayload := vpn.CreateGatewayConnectionPayload{
		DisplayName: "exampleVpnConnection",
		Tunnel1:     tunnel,
		Tunnel2:     tunnel,
	}
	connectionResp, err := vpnClient.DefaultAPI.CreateGatewayConnection(context.Background(), projectId, region, *gatewayResp.Id).CreateGatewayConnectionPayload(createGatewayConnectionPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateVpnConnection`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created VPN Connection with id \"%s\".\n", *connectionResp.Id)

	// Delete the VPN Connection
	err = vpnClient.DefaultAPI.DeleteGatewayConnection(context.Background(), projectId, region, *gatewayResp.Id, *connectionResp.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteVpnConnection`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Deleted VPN Connection with id \"%s\".\n", *connectionResp.Id)

	// Delete the VPN Gateway
	err = vpnClient.DefaultAPI.DeleteGateway(context.Background(), projectId, region, *gatewayResp.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteVpnGateway`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Deleted VPN Gateway with id \"%s\".\n", *gatewayResp.Id)
}
