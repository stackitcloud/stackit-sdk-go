package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/sfs"
	"github.com/stackitcloud/stackit-sdk-go/services/sfs/wait"
)

func main() {
	// Specify the project ID and region
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	region := "eu01"

	// Create a new API client, that uses default authentication and configuration
	sfsClient, err := sfs.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create a resource pool
	createResourcePoolPayload := sfs.CreateResourcePoolPayload{
		AvailabilityZone: utils.Ptr("eu01-m"),
		Name:             utils.Ptr("example-resourcepool-test"),
		PerformanceClass: utils.Ptr("Standard"),
		IpAcl:            &[]string{"192.168.42.1/32", "192.168.42.2/32"},
		SizeGigabytes:    utils.Ptr(int64(512)),
	}

	resourcePool, err := sfsClient.CreateResourcePool(context.Background(), projectId, region).CreateResourcePoolPayload(createResourcePoolPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when calling `CreateResourcePool`: %v\n", err)
		os.Exit(1)
	}

	if resourcePool == nil || resourcePool.ResourcePool == nil || resourcePool.ResourcePool.Id == nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error creating resource pool. Calling API: Incomplete response")
	}

	fmt.Printf("[sfs API] Triggered creation of resource pool with ID %q.\n", *resourcePool.ResourcePool.Id)
	fmt.Printf("[sfs API] Current state of the resource pool: %q\n", *resourcePool.ResourcePool.State)
	fmt.Println("[sfs API] Waiting for resource pool to be created...")

	resourcePoolResp, err := wait.CreateResourcePoolWaitHandler(context.Background(), sfsClient, projectId, region, *resourcePool.ResourcePool.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[sfs API] Resource pool has been successfully created.\n")

	// Describe the resource pool
	getResourcePoolResp, err := sfsClient.GetResourcePool(context.Background(), projectId, region, *resourcePoolResp.ResourcePool.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when calling `GetResoucePool`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[sfs API] Resource pool %s with ID %d was received.\n", *getResourcePoolResp.ResourcePool.Name, getResourcePoolResp.ResourcePool.Id)

	// Update the resource pool
	updateResourcePoolPayload := sfs.UpdateResourcePoolPayload{
		SizeGigabytes: utils.Ptr(int64(600)),
	}

	updateResourcePool, err := sfsClient.UpdateResourcePool(context.Background(), projectId, region, *resourcePoolResp.ResourcePool.Id).UpdateResourcePoolPayload(updateResourcePoolPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when calling `UpdateResourcePool`: %v\n", err)
		os.Exit(1)
	}

	_, err = wait.UpdateResourcePoolWaitHandler(context.Background(), sfsClient, projectId, region, *resourcePoolResp.ResourcePool.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when waiting for update: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[sfs API] Resource pool has been successfully updated to new size %d.\n", *updateResourcePool.ResourcePool.Space.SizeGigabytes)

	// Delete the resource pool
	_, err = sfsClient.DeleteResourcePool(context.Background(), projectId, region, *resourcePoolResp.ResourcePool.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when calling `DeleteResourcePool`: %v\n", err)
		os.Exit(1)
	}

	_, err = wait.DeleteResourcePoolWaitHandler(context.Background(), sfsClient, projectId, region, *resourcePoolResp.ResourcePool.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[sfs API] Resource pool has been successfully deleted.\n")
}
