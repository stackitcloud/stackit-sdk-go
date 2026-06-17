package main

import (
	"context"
	"fmt"
	"os"

	modelexperiments "github.com/stackitcloud/stackit-sdk-go/services/modelexperiments/v1api"
	"github.com/stackitcloud/stackit-sdk-go/services/modelexperiments/v1api/wait"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	region := "eu01"
	instanceName := "instance"
	newInstanceName := "newInstance"
	newTokenName := "newTokenName"
	description := "description"
	ctx := context.Background()

	modelexperimentsClient, err := modelexperiments.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create an Model Experiments Instance
	createInstancePayload := modelexperiments.CreateInstancePayload{
		Name:        instanceName,
		Description: &description,
	}
	createInstanceResp, err := modelexperimentsClient.DefaultAPI.CreateInstance(ctx, projectId, region).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when calling `CreateInstance`: %v\n", err)
	} else {
		fmt.Printf("[Model Experiments] Created Model Experiments Instance with ID: \"%s\".\n", createInstanceResp.Instance.Id)
	}

	// Wait for the Model Experiments Instance to be ready
	fmt.Printf("[Model Experiments] Waiting for Model Experiments Instance to be created.\n")
	_, err = wait.CreateModelExperimentsInstanceWaitHandler(ctx, modelexperimentsClient.DefaultAPI, region, projectId, createInstanceResp.Instance.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when waiting for creation: %v\n", err)
	}
	fmt.Printf("[Model Experiments] Model Experiments Instance \"%s\" has been successfully created.\n", createInstanceResp.Instance.Id)

	// Get the created Model Experiments Instance
	getInstanceResp, err := modelexperimentsClient.DefaultAPI.GetInstance(ctx, projectId, region, createInstanceResp.Instance.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when calling `GetInstance`: %v\n", err)
	} else {
		fmt.Printf("[Model Experiments] Retrieved Model Experiments Instance with ID: \"%s\" and display name: \"%s\"\n", getInstanceResp.Instance.Id, getInstanceResp.Instance.Name)
	}

	// List Model Experiments Instances
	listInstaceResp, err := modelexperimentsClient.DefaultAPI.ListInstances(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when calling `ListInstances`: %v\n", err)
	} else {
		fmt.Printf("[Model Experiments] Retrieved %d instances\n", len(listInstaceResp.Instances))
	}

	// Update an Model Experiments Instance
	partialUpdateInstancePayload := modelexperiments.PartialUpdateInstancePayload{
		Name: &newInstanceName,
	}
	partialUpdateInstanceResp, err := modelexperimentsClient.DefaultAPI.PartialUpdateInstance(ctx, projectId, region, getInstanceResp.Instance.Id).PartialUpdateInstancePayload(partialUpdateInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when calling `PartialUpdateInstance`: %v\n", err)
	} else {
		fmt.Printf("[Model Experiments] Updated Model Experiments Instance with ID: \"%s\" and display name: \"%s\"\n", partialUpdateInstanceResp.Instance.Id, partialUpdateInstanceResp.Instance.Name)
	}

	// Delete an Model Experiments Instance
	deleteInstanceResp, err := modelexperimentsClient.DefaultAPI.DeleteInstance(ctx, projectId, region, getInstanceResp.Instance.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when calling `DeleteInstance`: %v\n", err)
	} else {
		fmt.Printf("[Model Experiments] Deleted Model Experiments Instance with ID: \"%s\" and display name: \"%s\"\n", deleteInstanceResp.Instance.Id, deleteInstanceResp.Instance.Name)
	}

	// Wait for the Model Experiments Instance to be deleted
	fmt.Printf("[Model Experiments] Waiting for Model Experiments Instance to be deleted.\n")
	_, err = wait.DeleteModelExperimentsInstanceWaitHandler(ctx, modelexperimentsClient.DefaultAPI, region, projectId, getInstanceResp.Instance.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when waiting for deletion: %v\n", err)
	}
	fmt.Printf("[Model Experiments] Model Experiments Instance \"%s\" has been successfully deleted.\n", deleteInstanceResp.Instance.Id)

	// Create an Model Experiments Instance Token
	createTokenPayload := modelexperiments.CreateInstanceTokenPayload{
		Name: "token-name",
	}
	createTokenResp, err := modelexperimentsClient.DefaultAPI.CreateInstanceToken(ctx, projectId, region, createInstanceResp.Instance.Id).CreateInstanceTokenPayload(createTokenPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when calling `CreateToken`: %v\n", err)
	} else {
		fmt.Printf("[Model Experiments] Created Model Experiments Instance Token: %+v\n", createTokenResp.Token)
	}

	// Wait for the Model Experiments Instance Token to be ready
	fmt.Printf("[Model Experiments] Waiting for Model Experiments Instance Token to be created.\n")
	_, err = wait.CreateModelExperimentsInstanceTokenWaitHandler(ctx, modelexperimentsClient.DefaultAPI, region, projectId, getInstanceResp.Instance.Id, createTokenResp.Token.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when waiting for creation: %v\n", err)
	}
	fmt.Printf("[Model Experiments] Model Experiments Instance Token \"%s\" has been successfully created.\n", createTokenResp.Token.Id)

	// Get the created Model Experiments Instance Token
	getTokenResp, err := modelexperimentsClient.DefaultAPI.GetInstanceToken(ctx, projectId, region, createTokenResp.Token.Id, getInstanceResp.Instance.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when calling `GetInstanceToken`: %v\n", err)
	} else {
		fmt.Printf("[Model Experiments] Retrieved Model Experiments Instance Token with ID: \"%s\" and display name: \"%s\"\n", getTokenResp.Token.Id, getTokenResp.Token.Name)
	}

	// List Model Experiments Instance Tokens
	listTokenResp, err := modelexperimentsClient.DefaultAPI.ListInstanceTokens(ctx, projectId, region, getInstanceResp.Instance.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when calling `ListInstanceTokens`: %v\n", err)
	} else {
		fmt.Printf("[Model Experiments] Retrieved %d instance tokens\n", len(listTokenResp.Tokens))
	}

	// Update an Model Experiments Instance Token
	partialUpdateInstanceTokenPayload := modelexperiments.PartialUpdateInstanceTokenPayload{
		Name: &newTokenName,
	}
	partialUpdateTokenResp, err := modelexperimentsClient.DefaultAPI.PartialUpdateInstanceToken(ctx, projectId, region, getTokenResp.Token.Id, getInstanceResp.Instance.Id).PartialUpdateInstanceTokenPayload(partialUpdateInstanceTokenPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when calling `PartialUpdateInstanceToken`: %v\n", err)
	} else {
		fmt.Printf("[Model Experiments] Updated Model Experiments Instance Token with ID: \"%s\" and display name: \"%s\"\n", partialUpdateTokenResp.Token.Id, partialUpdateTokenResp.Token.Name)
	}

	// Delete an Model Experiments Instance Token
	deleteTokenResp, err := modelexperimentsClient.DefaultAPI.DeleteInstanceToken(ctx, projectId, region, getTokenResp.Token.Id, getInstanceResp.Instance.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when calling `DeleteInstanceToken`: %v\n", err)
	} else {
		fmt.Printf("[Model Experiments] Deleted Model Experiments Instance Token with ID: \"%s\" and display name: \"%s\"\n", deleteTokenResp.Token.Id, deleteTokenResp.Token.Name)
	}

	// Wait for the Model Experiments Instance Token to be deleted
	fmt.Printf("[Model Experiments] Waiting for Model Experiments Instance Token to be deleted.\n")
	_, err = wait.DeleteModelExperimentsInstanceTokenWaitHandler(ctx, modelexperimentsClient.DefaultAPI, region, projectId, getInstanceResp.Instance.Id, deleteTokenResp.Token.Id).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Model Experiments] Error when waiting for deletion: %v\n", err)
	}
	fmt.Printf("[Model Experiments] Model Experiments Instance Token \"%s\" has been successfully deleted.\n", deleteTokenResp.Token.Id)
}
