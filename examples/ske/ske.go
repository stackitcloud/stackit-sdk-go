package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/ske"
	"github.com/stackitcloud/stackit-sdk-go/services/ske/wait"
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
		fmt.Printf("Triggered creation of cluster with name \"%s\".\n", *createClusterResp.Name)
	}

	// Wait for cluster creation to complete
	_, err = wait.CreateOrUpdateClusterWaitHandler(context.Background(), skeClient, projectId, clusterName).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateOrUpdateCluster`: %v\n", err)
	} else {
		fmt.Printf("Cluster created.\n")
	}

	// Start cluster credential rotation
	_, err = skeClient.StartCredentialsRotationExecute(context.Background(), projectId, clusterName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StartCredentialsRotation`: %v\n", err)
	} else {
		fmt.Printf("Triggered start of cluster credentials rotation.\n")
	}

	// Wait for cluster credential rotation to be prepared
	_, err = wait.StartCredentialsRotationWaitHandler(context.Background(), skeClient, projectId, clusterName).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StartRotateCredentials`: %v\n", err)
	} else {
		fmt.Printf("Cluster credentials ready to rotate.\n")
	}

	// Complete cluster credential rotation
	_, err = skeClient.CompleteCredentialsRotationExecute(context.Background(), projectId, clusterName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompleteCredentialsRotation`: %v\n", err)
	} else {
		fmt.Printf("Triggered completion of cluster credentials rotation.\n")
	}

	// Wait for cluster credential rotation to be completed
	_, err = wait.CompleteCredentialsRotationWaitHandler(context.Background(), skeClient, projectId, clusterName).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompleteRotateCredentials`: %v\n", err)
	} else {
		fmt.Printf("Completed cluster credentials rotation.\n")
	}

}
