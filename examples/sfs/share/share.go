package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/sfs"
	"github.com/stackitcloud/stackit-sdk-go/services/sfs/wait"
)

func main() {
	// Specify the project ID and region
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	region := "eu01"

	// Create a new API client, that uses default authentication and configuration
	sfsClient, err := sfs.NewAPIClient(config.WithRegion(region))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create a resource pool in order to create a share
	createResourcePoolPayload := sfs.CreateResourcePoolPayload{
		AvailabilityZone: utils.Ptr("eu01-m"),
		Name:             utils.Ptr("example-resourcepool-share-test"),
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

	_, err = wait.CreateResourcePoolWaitHandler(context.Background(), sfsClient, projectId, region, *resourcePool.ResourcePool.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when waiting for creation of resource pool: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[sfs API] Resource pool has been successfully created.\n")

	// Create a share for the resource pool
	createSharePayload := sfs.CreateSharePayload{
		Name:                    utils.Ptr("example-share-name"),
		SpaceHardLimitGigabytes: utils.Ptr(int64(100)),
	}

	share, err := sfsClient.CreateShare(context.Background(), projectId, region, *resourcePool.ResourcePool.Id).CreateSharePayload(createSharePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when waiting calling `CreateShare`: %v\n", err)
		os.Exit(1)
	}

	if share == nil || share.Share == nil || share.Share.Id == nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error creating share. Calling API: Incomplete response")
	}

	_, err = wait.CreateShareWaitHandler(context.Background(), sfsClient, projectId, region, *resourcePool.ResourcePool.Id, *share.Share.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[sfs API] Share has been successfully created.\n")

	// Describe the share
	getShareResp, err := sfsClient.GetShare(context.Background(), projectId, region, *resourcePool.ResourcePool.Id, *share.Share.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when calling `GetShare`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[sfs API] Share %s with ID %d was received.\n", *getShareResp.Share.Name, getShareResp.Share.Id)

	// Update the share
	updateSharePayload := sfs.UpdateSharePayload{
		SpaceHardLimitGigabytes: utils.Ptr(int64(150)),
	}

	updateShare, err := sfsClient.UpdateShare(context.Background(), projectId, region, *resourcePool.ResourcePool.Id, *share.Share.Id).UpdateSharePayload(updateSharePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when calling `UpdateShare`: %v\n", err)
		os.Exit(1)
	}

	_, err = wait.UpdateShareWaitHandler(context.Background(), sfsClient, projectId, region, *resourcePool.ResourcePool.Id, *share.Share.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when waiting for update: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[sfs API] Share has been successfully updated to new limit %d.\n", *updateShare.Share.SpaceHardLimitGigabytes)

	// Delete share
	_, err = sfsClient.DeleteShare(context.Background(), projectId, region, *resourcePool.ResourcePool.Id, *share.Share.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when calling `DeleteShare`: %v\n", err)
		os.Exit(1)
	}

	_, err = wait.DeleteShareWaitHandler(context.Background(), sfsClient, projectId, region, *resourcePool.ResourcePool.Id, *share.Share.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[sfs API] Share has been successfully deleted.\n")

	// Delete resource pool
	_, err = sfsClient.DeleteResourcePool(context.Background(), projectId, region, *resourcePool.ResourcePool.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when calling `DeleteResourcePool`: %v\n", err)
		os.Exit(1)
	}

	_, err = wait.DeleteResourcePoolWaitHandler(context.Background(), sfsClient, projectId, region, *resourcePool.ResourcePool.Id).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[sfs API] Error when waiting for deletion: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[sfs API] Resource pool has been successfully deleted.\n")
}
