package main

import (
	"context"
	"log"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	logs "github.com/stackitcloud/stackit-sdk-go/services/logs/v1api"
)

func main() {
	ctx := context.Background()

	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	regionId := "eu01"

	client, err := logs.NewAPIClient()
	if err != nil {
		log.Fatalf("[Logs API] Creating API client: %v\n", err)
	}

	// Create a Logs Instance
	var createdInstance string
	createInstancePayload := logs.CreateLogsInstancePayload{
		DisplayName:   "my-logs-instance",
		RetentionDays: int32(1),
	}
	createResp, err := client.DefaultAPI.CreateLogsInstance(ctx, projectId, regionId).
		CreateLogsInstancePayload(createInstancePayload).
		Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `CreateLogsInstance`: %v\n", err)
	}
	log.Printf("[Logs API] Created Logs Instance with ID \"%s\".\n", createResp.Id)

	// List Logs Instances
	listResp, err := client.DefaultAPI.ListLogsInstances(ctx, projectId, regionId).Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `ListLogsInstances`: %v\n", err)
	}
	log.Printf("[Logs API] Retrieved %d Logs Instances.\n", len(listResp.Instances))

	// Get the created Logs Instance
	getResp, err := client.DefaultAPI.GetLogsInstance(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `GetLogsInstance`: %v\n", err)
	}
	log.Printf("[Logs API] Retrieved Logs Instance with ID \"%s\" and Display Name \"%s\".\n", getResp.Id, getResp.DisplayName)

	// Update the created Logs Instance
	updatePayload := logs.UpdateLogsInstancePayload{
		DisplayName:   utils.Ptr("my-updated-logs-instance"),
		RetentionDays: utils.Ptr(int32(7)),
	}
	updateResp, err := client.DefaultAPI.UpdateLogsInstance(ctx, projectId, regionId, createdInstance).
		UpdateLogsInstancePayload(updatePayload).
		Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `UpdateLogsInstance`: %v\n", err)
	}
	log.Printf("[Logs API] Updated Logs Instance with ID \"%s\" to Display Name \"%s\".\n", updateResp.Id, updateResp.DisplayName)

	// Create an Access Token
	createTokenPayload := logs.CreateAccessTokenPayload{
		DisplayName: "my-access-token",
		Permissions: []string{"read"},
	}
	createTokenResp, err := client.DefaultAPI.CreateAccessToken(ctx, projectId, regionId, createdInstance).
		CreateAccessTokenPayload(createTokenPayload).
		Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `CreateAccessToken`: %v\n", err)
	}
	log.Printf("[Logs API] Created Access Token with ID \"%s\".\n", createTokenResp.Id)

	// Add Access Token to Logs Instance
	err = client.DefaultAPI.UpdateAccessToken(ctx, projectId, regionId, createdInstance, createTokenResp.Id).
		// needs at least an empty payload
		UpdateAccessTokenPayload(logs.UpdateAccessTokenPayload{}).
		Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `UpdateAccessToken`: %v\n", err)
	}

	// Delete all Access Tokens from Logs Instance
	tokenList, err := client.DefaultAPI.DeleteAllAccessTokens(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `DeleteAllAccessTokens`: %v\n", err)
	}
	log.Printf("[Logs API] Deleted %d Access Tokens from Logs Instance with ID \"%s\".\n", len(tokenList.Tokens), createdInstance)

	// Delete the created Logs Instance
	err = client.DefaultAPI.DeleteLogsInstance(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Logs API] Error when calling `DeleteLogsInstance`: %v\n", err)
	}
	log.Printf("[Logs API] Deleted Logs Instance with ID \"%s\".\n", createdInstance)
}
