package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/kms"
	"github.com/stackitcloud/stackit-sdk-go/services/kms/wait"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"
	region := "eu01"

	// Create a new API client, that uses default authentication and configuration
	kmsClient, err := kms.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	keyRing, err := kmsClient.CreateKeyRing(ctx, projectId, region).CreateKeyRingPayload(kms.CreateKeyRingPayload{
		Description: utils.Ptr("a test keyring"),
		DisplayName: utils.Ptr("test-keyring"),
	}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot create keyring: %v\n", err)
		return
	}

	key, err := kmsClient.CreateKey(ctx, projectId, region, *keyRing.Id).CreateKeyPayload(kms.CreateKeyPayload{
		Algorithm:   kms.ALGORITHM_AES_256_GCM.Ptr(),
		Backend:     kms.BACKEND_SOFTWARE.Ptr(),
		Description: utils.Ptr("A test key"),
		DisplayName: utils.Ptr("test-key"),
		Purpose:     kms.PURPOSE_SYMMETRIC_ENCRYPT_DECRYPT.Ptr(),
	}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot create key: %v\n", err)
		return
	}
	if err := wait.CreateOrUpdateKeyWaitHandler(ctx, kmsClient, projectId, region, *key.KeyRingId, *key.Id); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create key: %v", err)
		return
	}
	fmt.Printf("created key %s\n", *key.Id)

	keyRings, err := kmsClient.ListKeyRingsExecute(ctx, projectId, region)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot list keyrings: %v\n", err)
		return
	}
	if keyrings := keyRings.KeyRings; keyrings != nil {
		if len(*keyrings) == 0 {
			fmt.Printf("no keyrings defined\n")
		} else {
			for _, keyring := range *keyrings {
				fmt.Printf("id=%s displayname=%s status=%s\n", *keyring.Id, *keyring.DisplayName, *keyring.State)
				keylist, err := kmsClient.ListKeysExecute(ctx, projectId, region, *key.KeyRingId)
				if err != nil {
					fmt.Fprintf(os.Stderr, "cannot list keys: %v", err)
					return
				}
				if keys := keylist.Keys; keys != nil {
					if len(*keys) == 0 {
						fmt.Printf("no keys\n")
					} else {
						for _, key := range *keys {
							fmt.Printf("key id=%s key name=%s key status=%s\n", *key.Id, *key.DisplayName, *key.State)
						}
					}
				}
			}
		}
	}
}
