package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/loadbalancer"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	loadbalancerClient, err := loadbalancer.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Enable load balancer functionality for project
	_, err = loadbalancerClient.EnableLoadBalancing(context.Background(), projectId).XRequestID("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnableLoadBalancing`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Enabled load balancer functionality for project %s.\n", projectId)

	// List the load balancer instances for your project
	listInstancesResp, err := loadbalancerClient.ListLoadBalancers(context.Background(), projectId).Execute()
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
				Role:      utils.Ptr(int32(1)),
			},
		},
		Listeners: &[]loadbalancer.Listener{
			{
				DisplayName: utils.Ptr("example-listener"),
				Port:        utils.Ptr(int32(1)),
				Protocol:    utils.Ptr(int32(1)),
				TargetPool:  utils.Ptr("example-target-pool"),
			},
		},
		TargetPools: &[]loadbalancer.TargetPool{
			{
				Name:       utils.Ptr("example-target-pool"),
				TargetPort: utils.Ptr(int32(1)),
				Targets: &[]loadbalancer.Target{
					{
						DisplayName: utils.Ptr("example-target"),
						Ip:          utils.Ptr("x.x.x.x"),
					},
				},
			},
		},
	}
	createLoadBalancerRes, err := loadbalancerClient.CreateLoadBalancer(context.Background(), projectId).CreateLoadBalancerPayload(createLoadBalancerPayload).XRequestID("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateLoadBalancer`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created load balancer \"%s\".\n", *createLoadBalancerRes.Name)
}
