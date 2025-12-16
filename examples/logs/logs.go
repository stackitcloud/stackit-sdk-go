package main

import (
	"context"
	"log"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/logs"
)

func main() {
	ctx := context.Background()

	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	regionId := "eu01"

	client, err := logs.NewAPIClient(
		config.WithRegion(regionId),
	)
	if err != nil {
		log.Fatalf("[Logs API] Creating API client: %v\n", err)
	}

	// Create a Logs Instance
	var createdInstance string
	createInstancePayload := logs.CreateLogsInstancePayload{
		DisplayName:   utils.Ptr("my-logs-instance"),
		RetentionDays: utils.Ptr(int64(1)),
	}
	createResp, err := client.CreateLogsInstance(ctx, projectId, regionId).
		CreateLogsInstancePayload(createInstancePayload).
		Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `CreateLogsInstance`: %v\n", err)
	}
	createdInstance = *createResp.Id
	log.Printf("[Logs API] Created Logs Instance with ID \"%s\".\n", createdInstance)

	// List Logs Instances
	listResp, err := client.ListLogsInstances(ctx, projectId, regionId).Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `ListLogsInstances`: %v\n", err)
	}
	log.Printf("[Logs API] Retrieved %d Logs Instances.\n", len(*listResp.Instances))

	// Get the created Logs Instance
	getResp, err := client.GetLogsInstance(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `GetLogsInstance`: %v\n", err)
	}
	log.Printf("[Logs API] Retrieved Logs Instance with ID \"%s\" and Display Name \"%s\".\n", *getResp.Id, *getResp.DisplayName)

	// Update the created Logs Instance
	updatePayload := logs.UpdateLogsInstancePayload{
		DisplayName:   utils.Ptr("my-updated-logs-instance"),
		RetentionDays: utils.Ptr(int64(7)),
	}
	updateResp, err := client.UpdateLogsInstance(ctx, projectId, regionId, createdInstance).
		UpdateLogsInstancePayload(updatePayload).
		Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `UpdateLogsInstance`: %v\n", err)
	}
	log.Printf("[Logs API] Updated Logs Instance with ID \"%s\" to Display Name \"%s\".\n", *updateResp.Id, *updateResp.DisplayName)

	// Create an Access Token
	createTokenPayload := logs.CreateAccessTokenPayload{
		DisplayName: utils.Ptr("my-access-token"),
		Permissions: &[]string{"read"},
	}
	createTokenResp, err := client.CreateAccessToken(ctx, projectId, regionId, createdInstance).
		CreateAccessTokenPayload(createTokenPayload).
		Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `CreateAccessToken`: %v\n", err)
	}
	log.Printf("[Logs API] Created Access Token with ID \"%s\".\n", *createTokenResp.Id)

	// Add Access Token to Logs Instance
	err = client.UpdateAccessToken(ctx, projectId, regionId, createdInstance, *createTokenResp.Id).
		// needs at least an empty payload
		UpdateAccessTokenPayload(logs.UpdateAccessTokenPayload{}).
		Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `UpdateAccessToken`: %v\n", err)
	}

	// Delete all Access Tokens from Logs Instance
	tokenList, err := client.DeleteAllAccessTokens(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `DeleteAllAccessTokens`: %v\n", err)
	}
	log.Printf("[Logs API] Deleted %d Access Tokens from Logs Instance with ID \"%s\".\n", len(*tokenList.Tokens), createdInstance)

	// Delete the created Logs Instance
	err = client.DeleteLogsInstance(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `DeleteLogsInstance`: %v\n", err)
	}
	log.Printf("[Logs API] Deleted Logs Instance with ID \"%s\".\n", createdInstance)
}
