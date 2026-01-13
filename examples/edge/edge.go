package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/edge"
	"github.com/stackitcloud/stackit-sdk-go/services/edge/wait"
	"gopkg.in/yaml.v3"
)

func main() {
	// Mandatory parameters
	projectId := "PROJECT_ID"
	region := "eu01"

	// Create a new API client, that uses default authentication and configuration
	client, err := edge.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Edge API] Failed to create API client: %v\n", err)
		os.Exit(1)
	}

	// Create an Edge Instance with default values
	var (
		payload  = edge.NewPostInstancesPayloadWithDefaults()
		instance *edge.Instance
		ctx      = context.Background()
	)
	payload.DisplayName = utils.Ptr("example")
	instance, err = client.PostInstances(ctx, projectId, region).PostInstancesPayload(*payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Edge API] Failed to create Instance: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[Edge API] Instance creation started, Id: %s\n", instance.GetId())

	// Wait for Edge Instance to become active
	waitResult, err := wait.CreateOrUpdateInstanceWaitHandler(
		ctx,
		client,
		projectId,
		region,
		instance.GetId(),
	).SetTimeout(10 * time.Minute).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Edge API] Failed wait for Instance creation: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[Edge API] Instance created, Id: %s, Status: %s\n, URL: %s\n", instance.GetId(), waitResult.GetStatus(), instance.GetFrontendUrl())

	// Create a service token to login to the instance UI and wait for the instance to become ready
	token, err := wait.TokenWaitHandler(
		ctx,
		client,
		projectId,
		region,
		instance.GetId(),
		utils.Ptr(int64(600)),
	).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Edge API] Failed wait for token creation: %v\n", err)
		os.Exit(1)
	}
	if token != nil && token.Token != nil {
		fmt.Printf("Token: %s\n", *token.Token)
	}

	// Create a kubeconfig to interact with the instances kubernetes API and wait for the instance to become ready
	kubeconfig, err := wait.KubeconfigWaitHandler(
		ctx,
		client,
		projectId,
		region,
		instance.GetId(),
		utils.Ptr(int64(600)),
	).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Edge API] Failed wait for kubeconfig creation: %v\n", err)
		os.Exit(1)
	}
	if kubeconfig != nil && kubeconfig.Kubeconfig != nil {
		yamlData, err := yaml.Marshal(kubeconfig.Kubeconfig)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[Edge API] Failed to marshal kubeconfig: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Kubeconfig:")
		fmt.Println(string(yamlData))
	}

	// Delete Edge Instance
	err = client.DeleteInstance(ctx, projectId, region, instance.GetId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Edge API] Failed to delete instance: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[Edge API] Instance deletion started, Id: %s\n", instance.GetId())

	// Wait for Edge instance deletion
	_, err = wait.DeleteInstanceWaitHandler(ctx, client, projectId, region, instance.GetId()).SetTimeout(10 * time.Minute).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Edge API] Failed wait for Instance deletion: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("[Edge API] Instance deleted")
}
