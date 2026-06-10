package main

import (
	"context"
	"fmt"
	"os"

	modelexperiments "github.com/stackitcloud/stackit-sdk-go/services/modelexperiments/v1api"
)

func main() {
	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	region := "eu01"
	instanceName := "instance"
	description := "description"

	modelexperimentsClient, err := modelexperiments.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create a instance
	createInstancePayload := modelexperiments.CreateInstancePayload{
		Name:        instanceName,
		Description: &description,
	}
	createInstanceResp, err := modelexperimentsClient.DefaultAPI.CreateInstance(context.Background(), projectId, region).CreateInstancePayload(createInstancePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateInstance`: %v\n", err)
	} else {
		fmt.Printf("Created instance with instance id \"%s\".\n", createInstanceResp.Instance.Id)
	}

	// Get the created instance
	getInstanceResp, err := modelexperimentsClient.DefaultAPI.GetInstance(context.Background(), projectId, region, createInstanceResp.Instance.Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetInstance`: %v\n", err)
	} else {
		fmt.Printf("Retrieved instance: %+v\n", getInstanceResp.Instance)
	}

	// Create a token for the instance
	createTokenPayload := modelexperiments.CreateInstanceTokenPayload{
		Name: "token-name",
	}
	createTokenResp, err := modelexperimentsClient.DefaultAPI.CreateInstanceToken(context.Background(), projectId, region, createInstanceResp.Instance.Id).CreateInstanceTokenPayload(createTokenPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateToken`: %v\n", err)
	} else {
		fmt.Printf("Created token: %+v\n", createTokenResp.Token)
	}
}
