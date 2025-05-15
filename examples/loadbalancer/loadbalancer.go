package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/loadbalancer"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Specify the region
	region := "REGION"

	// Create a new API client, that uses default authentication and configuration
	loadbalancerClient, err := loadbalancer.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// List the load balancer instances for your project
	listInstancesResp, err := loadbalancerClient.ListLoadBalancers(context.Background(), projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListLoadBalancers`: %v\n", err)
		os.Exit(1)
	}
	if listInstancesResp.LoadBalancers == nil {
		fmt.Printf("Current project does not have any load balancer instances.\n")
	} else {
		lbs := *listInstancesResp.LoadBalancers
		fmt.Printf("Number of instances: %v\n", len(lbs))
	}

	// Create a load balancer instance
	createLoadBalancerPayload := loadbalancer.CreateLoadBalancerPayload{
		Name: utils.Ptr("example-instance"),
		Options: &loadbalancer.LoadBalancerOptions{
			PrivateNetworkOnly: utils.Ptr(true),
		},
		Networks: &[]loadbalancer.Network{
			{
				NetworkId: utils.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
				Role:      utils.Ptr(loadbalancer.NETWORKROLE_LISTENERS_AND_TARGETS),
			},
		},
		Listeners: &[]loadbalancer.Listener{
			{
				DisplayName: utils.Ptr("example-listener"),
				Port:        utils.Ptr(int64(1)),
				Protocol:    utils.Ptr(loadbalancer.LISTENERPROTOCOL_TCP),
				TargetPool:  utils.Ptr("example-target-pool"),
			},
		},
		TargetPools: &[]loadbalancer.TargetPool{
			{
				Name:       utils.Ptr("example-target-pool"),
				TargetPort: utils.Ptr(int64(1)),
				Targets: &[]loadbalancer.Target{
					{
						DisplayName: utils.Ptr("example-target"),
						Ip:          utils.Ptr("x.x.x.x"),
					},
				},
			},
		},
	}
	createLoadBalancerRes, err := loadbalancerClient.CreateLoadBalancer(context.Background(), projectId, region).CreateLoadBalancerPayload(createLoadBalancerPayload).XRequestID("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateLoadBalancer`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created load balancer \"%s\".\n", *createLoadBalancerRes.Name)
}
