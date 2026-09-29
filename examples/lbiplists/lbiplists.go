package main

import (
	"context"
	"fmt"
	"os"
	"time"

	lbiplists "github.com/stackitcloud/stackit-sdk-go/services/lbiplists/v1alphaapi"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project

	// Specify the region
	region := "eu01"

	// Name of the IP list used in this example
	name := "sdk-example-ip-list"

	ctx := context.Background()

	// Create a new API client, that uses default authentication and configuration
	lbiplistsClient, err := lbiplists.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create an IP list
	createIPListPayload := lbiplists.UploadIPListPayload{
		Name:        name,
		FileContent: "10.0.0.0/8\n192.168.0.0/16",
	}
	createIPListResp, err := lbiplistsClient.DefaultAPI.UploadIPList(ctx, projectId, region, name).UploadIPListPayload(createIPListPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadIPList`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created IP list %q with %d CIDR entries.\n", createIPListResp.Name, *createIPListResp.NumberOfIps)

	// List the IP lists for your project
	listIPListsResp, err := lbiplistsClient.DefaultAPI.ListIPLists(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListIPLists`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Number of IP lists: %v\n", len(listIPListsResp.Items))

	// Get the IP list
	getIPListResp, err := lbiplistsClient.DefaultAPI.GetIPList(ctx, projectId, region, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetIPList`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Fetched IP list %q with content hash %q.\n", getIPListResp.Name, *getIPListResp.ContentHash)

	// UploadIPList is limited to 1 request per minute; wait before updating
	// the list we just created to avoid a 429 Too Many Requests error.
	fmt.Println("Waiting 60s before updating the IP list (UploadIPList is rate-limited to 1 request/minute)...")
	time.Sleep(60 * time.Second)

	// Update the IP list by uploading new content
	updateIPListPayload := lbiplists.UploadIPListPayload{
		Name:        name,
		FileContent: "203.0.113.0/24",
		Labels:      &map[string]string{"purpose": "sdk-example"},
	}
	updateIPListResp, err := lbiplistsClient.DefaultAPI.UploadIPList(ctx, projectId, region, name).UploadIPListPayload(updateIPListPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadIPList`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Updated IP list %q, now containing %d CIDR entries.\n", updateIPListResp.Name, *updateIPListResp.NumberOfIps)

	// Get the IP list again to show the update
	getUpdatedIPListResp, err := lbiplistsClient.DefaultAPI.GetIPList(ctx, projectId, region, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetIPList`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Fetched updated IP list %q with new content hash %q and labels %v.\n", getUpdatedIPListResp.Name, *getUpdatedIPListResp.ContentHash, *getUpdatedIPListResp.Labels)

	// Delete the IP list
	_, err = lbiplistsClient.DefaultAPI.DeleteIPList(ctx, projectId, region, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteIPList`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Deleted IP list %q.\n", name)

	// List the IP lists again to show the clean slate
	listIPListsResp, err = lbiplistsClient.DefaultAPI.ListIPLists(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListIPLists`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Number of IP lists after cleanup: %v\n", len(listIPListsResp.Items))
}
