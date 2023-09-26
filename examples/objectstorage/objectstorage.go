package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/services/objectstorage"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	objectStorageClient, err := objectstorage.NewAPIClient(
		config.WithRegion("eu01"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the object storage buckets for your project
	getBucketsResp, err := objectStorageClient.GetBuckets(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetBuckets`: %v\n", err)
	} else {
		fmt.Printf("Number of buckets: %v\n", len(*getBucketsResp.Buckets))
	}

	// Create an object storage bucket
	createBucketResp, err := objectStorageClient.CreateBucket(context.Background(), projectId, "example-bucket").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateBucket`: %v\n", err)
	} else {
		fmt.Printf("Created bucket with name \"%s\".\n", *createBucketResp.Bucket)
	}
}
