package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	kms "github.com/stackitcloud/stackit-sdk-go/services/kms/v1api"
	"github.com/stackitcloud/stackit-sdk-go/services/kms/v1api/wait"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	region := "eu01"

	// Create a new API client, that uses default authentication and configuration
	kmsClient, err := kms.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[kms API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()

	keyRing, err := kmsClient.DefaultAPI.CreateKeyRing(ctx, projectId, region).
		CreateKeyRingPayload(
			kms.CreateKeyRingPayload{
				Description: utils.Ptr("a test keyring"),
				DisplayName: "test-keyring",
			},
		).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[kms API] Cannot create keyring: %v\n", err)
		os.Exit(1)
	}

	// Create a key
	key, err := kmsClient.DefaultAPI.CreateKey(ctx, projectId, region, keyRing.Id).
		CreateKeyPayload(
			kms.CreateKeyPayload{
				Algorithm:   kms.ALGORITHM_AES_256_GCM,
				Description: utils.Ptr("A test key"),
				DisplayName: "test-key",
				Purpose:     kms.PURPOSE_SYMMETRIC_ENCRYPT_DECRYPT,
			}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[kms API] Cannot create key: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[kms API] Triggered creation of key: %v\n", key.Id)

	// Wait for creation of key
	key, err = wait.CreateOrUpdateKeyWaitHandler(ctx, kmsClient.DefaultAPI, projectId, region, key.KeyRingId, key.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[kms API] Error when waiting for creation: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[kms API] Created key %s\n", key.Id)

	// List key rings
	keyRings, err := kmsClient.DefaultAPI.ListKeyRings(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[kms API] Cannot list keyrings: %v\n", err)
		os.Exit(1)
	}

	if keyrings := keyRings.KeyRings; keyrings != nil {
		fmt.Printf("[kms API] Number of keyrings: %v\n", len(keyrings))

		for _, keyring := range keyrings {
			keylist, err := kmsClient.DefaultAPI.ListKeys(ctx, projectId, region, keyring.Id).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "[kms API] Cannot list keys: %v\n", err)
				os.Exit(1)
			}

			if keys := keylist.Keys; keys != nil {
				fmt.Printf("[kms API] Keys in Keyring %s: %v\n", keyring.Id, len(keys))
			}
		}
	}
}
