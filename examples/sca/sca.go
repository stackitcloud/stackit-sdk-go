package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	sca "github.com/stackitcloud/stackit-sdk-go/services/sca/v1alphaapi"
	"github.com/stackitcloud/stackit-sdk-go/services/sca/v1alphaapi/wait"
)

func main() {
	region := "eu01"          // Region where the resources will be created
	projectID := "PROJECT_ID" // the uuid of your STACKIT project

	// Create a new API client, that uses default authentication and configuration
	token := "TOKEN"
	scaClient, err := sca.NewAPIClient(config.WithRegion(region), config.WithToken(token))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create environment
	createEnvironmentPayload := sca.CreateEnvironmentPayload{
		DisplayName: "environment-name",
	}
	env, err := scaClient.DefaultAPI.CreateEnvironment(context.Background(), projectID).
		CreateEnvironmentPayload(createEnvironmentPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateEnvironment`: %v\n", err)
	} else {
		fmt.Printf("Created environment with id %q\n", env.GetId())
	}

	// Get environment
	getEnvResp, err := scaClient.DefaultAPI.GetEnvironment(context.Background(), projectID, env.GetId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetEnvironment`: %v\n", err)
	} else {
		fmt.Printf("Got environment with id %q\n", getEnvResp.GetId())
	}

	// List environments
	listEnvsResp, err := scaClient.DefaultAPI.ListEnvironments(context.Background(), projectID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListEnvironments`: %v\n", err)
	} else {
		fmt.Printf("Number of environments in project: %d\n", len(listEnvsResp.Items))
	}

	// List all applications of a given project
	listAppsResp, err := scaClient.DefaultAPI.ListProjectApplications(context.Background(), projectID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListProjectApplications`: %v\n", err)
	} else {
		fmt.Printf("Number of applications in project: %d\n", len(listAppsResp.Items))
	}

	// Create an application within an environment
	environmentID := "ENVIRONMENT_ID"
	createApplicationPayload := sca.CreateApplicationPayload{
		DisplayName: "application-name",
		Containers: []sca.Container{{
			Name:   "nginx-container",
			Image:  "nginxinc/nginx-unprivileged",
			Memory: sca.PtrInt32(100),
			Cpu:    sca.PtrInt32(100),
		}},
		Network: sca.Network{
			PublicIngress: true,
			Port:          sca.PtrInt32(8080),
		},
		Scaling: sca.Scaling{
			Type: sca.SCALINGTYPE_SCALING_TYPE_MANUAL,
			ManualScaling: &sca.ManualScaling{
				Instances: 1,
			},
		},
	}

	app, err := scaClient.DefaultAPI.CreateApplication(context.Background(), projectID, environmentID).
		CreateApplicationPayload(createApplicationPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateApplication`: %v\n", err)
	} else {
		fmt.Printf("Triggered application creation whit id %q\n", app.GetId())
	}

	_, err = wait.CreateApplicationWaitHandler(context.Background(), scaClient.DefaultAPI, projectID, environmentID, app.GetId()).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateOrUpdateCluster`: %v\n", err)
	} else {
		fmt.Printf("Application created witht id %q\n", app.GetId())
	}

	// Get application's details
	getAppResp, err := scaClient.DefaultAPI.GetApplication(context.Background(), projectID, environmentID, app.GetId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetApplication`: %v\n", err)
	} else {
		fmt.Printf("Got application with id %q\n", getAppResp.GetId())
	}

	// List applications in an environment
	listEnvAppsResp, err := scaClient.DefaultAPI.ListApplications(context.Background(), projectID, environmentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListApplications`: %v\n", err)
	} else {
		fmt.Printf("Number of applications in environment: %d\n", len(listEnvAppsResp.Items))
	}

	// Update application
	updateApplicationPayload := sca.UpdateApplicationPayload{
		Scaling: &sca.Scaling{
			Type: sca.SCALINGTYPE_SCALING_TYPE_MANUAL,
			AutoScaling: &sca.AutoScaling{
				AllowScaleToZero: sca.PtrBool(true),
				MinInstances:     1,
				MaxInstances:     2,
				Rules: []sca.ScaleRule{{
					Name: "http-scaling-rule",
					Type: sca.RULETYPE_RULE_TYPE_HTTP,
					HttpRule: &sca.HttpScaleRule{
						Concurrency: sca.PtrInt32(10),
						Rps:         sca.PtrInt32(10),
					},
				}},
			},
		},
	}
	updateAppResp, err := scaClient.DefaultAPI.UpdateApplication(context.Background(), projectID, environmentID, app.GetId()).
		UpdateApplicationPayload(updateApplicationPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateApplication`: %v\n", err)
	} else {
		fmt.Printf("Updated application with id %q\n", updateAppResp.GetId())
	}

	deleteApplicationResp, err := scaClient.DefaultAPI.DeleteApplication(context.Background(), projectID, environmentID, app.GetId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteApplication`: %v\n", err)
	} else {
		fmt.Printf("Deleted application with id %q\n", *deleteApplicationResp.Id)
	}
}
