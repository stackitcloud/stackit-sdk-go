package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/ske"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	skeClient, err := ske.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// The following operations assume you have a ske project already created. If you dont, run:
	// createProjectResponse, err := skeClient.CreateProject(context.Background(), projectId).Execute()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error when calling `CreateProject`: %v\n", err)
	// } else {
	// 	fmt.Printf("Project created with projectId: %v\n", len(*createProjectResponse.ProjectId))
	// }

	// Get the ske clusters for your project
	getClustersResp, err := skeClient.ListClusters(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetClusters`: %v\n", err)
	} else {
		fmt.Printf("Number of clusters: %v\n", len(*getClustersResp.Items))
	}

	var availableVersion string
	// Get the ske provider options
	getOptionsResp, err := skeClient.ListProviderOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetOptions`: %v\n", err)
	} else {
		availableVersions := *getOptionsResp.KubernetesVersions
		availableVersion = *availableVersions[0].Version
		fmt.Printf("First available version: %v\n", availableVersion)
	}

	// Create an ske cluster
	createInstancePayload := ske.CreateOrUpdateClusterPayload{
		Kubernetes: &ske.Kubernetes{
			Version: utils.Ptr(availableVersion),
		},
		Nodepools: &[]ske.Nodepool{
			{
				AvailabilityZones: utils.Ptr([]string{"eu01-3"}),
				Machine: &ske.Machine{
					Image: &ske.Image{
						Name:    utils.Ptr("name"),
						Version: utils.Ptr("3510.2.5"),
					},
					Type: utils.Ptr("b1.2"),
				},
				Maximum: utils.Ptr(int64(3)),
				Minimum: utils.Ptr(int64(2)),
				Name:    utils.Ptr("my-nodepool"),
				Volume: &ske.Volume{
					Size: utils.Ptr(int64(20)),
					Type: utils.Ptr("storage_premium_perf0"),
				},
			},
		},
	}
	clusterName := "cl-name"
	createClusterResp, err := skeClient.CreateOrUpdateCluster(context.Background(), projectId, clusterName).CreateOrUpdateClusterPayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateCluster`: %v\n", err)
	} else {
		fmt.Printf("Created cluster with cluster name \"%s\".\n", *createClusterResp.Name)
	}
}
