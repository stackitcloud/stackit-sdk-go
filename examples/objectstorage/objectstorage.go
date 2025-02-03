package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/services/objectstorage"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"
	region := "eu01"

	// Create a new API client, that uses default authentication and configuration
	objectStorageClient, err := objectstorage.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the object storage buckets for your project
	getBucketsResp, err := objectStorageClient.ListBuckets(context.Background(), projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetBuckets`: %v\n", err)
	} else {
		fmt.Printf("Number of buckets: %v\n", len(*getBucketsResp.Buckets))
	}

	// Create an object storage bucket
	createBucketResp, err := objectStorageClient.CreateBucket(context.Background(), projectId, region, "example-bucket").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateBucket`: %v\n", err)
	} else {
		fmt.Printf("Created bucket with name \"%s\".\n", *createBucketResp.Bucket)
	}
}
