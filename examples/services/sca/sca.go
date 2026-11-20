package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	sca "github.com/stackitcloud/stackit-sdk-go/services/sca/v1betaapi"
	"github.com/stackitcloud/stackit-sdk-go/services/sca/v1betaapi/wait"
)

func main() {
	region := "eu01"          // Region where the resources will be created
	projectID := "PROJECT_ID" // the uuid of your STACKIT project

	// Create a new API client, that uses default authentication and configuration
	scaClient, err := sca.NewAPIClient(config.WithRegion(region))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create environment
	// WARNING: Keep in mind that there is no endpoint to delete environments right now.
	createEnvironmentPayload := sca.CreateEnvironmentPayload{
		DisplayName: "environment-name",
	}
	env, err := scaClient.DefaultAPI.CreateEnvironment(context.Background(), projectID, region).
		CreateEnvironmentPayload(createEnvironmentPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateEnvironment`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Triggered environment creation with id %q\n", env.GetId())

	_, err = wait.CreateEnvironmentWaitHandler(context.Background(), scaClient.DefaultAPI, projectID, region, env.GetId()).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateEnvironmentWaitHandler`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created environment with id %q\n", env.GetId())

	// Get environment
	getEnvResp, err := scaClient.DefaultAPI.GetEnvironment(context.Background(), projectID, region, env.GetId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetEnvironment`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Got environment with id %q\n", getEnvResp.GetId())

	// List environments
	listEnvsResp, err := scaClient.DefaultAPI.ListEnvironments(context.Background(), projectID, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListEnvironments`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Number of environments in project: %d\n", len(listEnvsResp.Items))

	// List all applications of a given project
	listAppsResp, err := scaClient.DefaultAPI.ListProjectApplications(context.Background(), projectID, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListProjectApplications`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Number of applications in project: %d\n", len(listAppsResp.Items))

	// Create an application within an environment
	createApplicationPayload := sca.CreateApplicationPayload{
		DisplayName: "application-name",
		Containers: []sca.Container{{
			Name:   "nginx-container",
			Image:  "nginxinc/nginx-unprivileged",
			Memory: sca.PtrInt32(1000),
			Cpu:    sca.PtrInt32(1000),
		}},
		Network: &sca.ApplicationNetwork{
			Public: &sca.ApplicationPublicNetwork{
				Enabled: true,
				Port:    sca.PtrInt32(8080),
			},
		},
		Scaling: sca.ApplicationScaling{
			Type: sca.APPLICATIONSCALINGTYPE_SCALING_TYPE_MANUAL,
			ManualScaling: &sca.ApplicationManualScaling{
				Instances: 1,
			},
		},
	}

	app, err := scaClient.DefaultAPI.CreateApplication(context.Background(), projectID, region, env.GetId()).
		CreateApplicationPayload(createApplicationPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateApplication`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Triggered application creation with id %q\n", app.GetId())

	_, err = wait.CreateApplicationWaitHandler(context.Background(), scaClient.DefaultAPI, projectID, region, env.GetId(), app.GetId()).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateApplicationWaitHandler`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Application created with id %q\n", app.GetId())

	// Get application's details
	getAppResp, err := scaClient.DefaultAPI.GetApplication(context.Background(), projectID, region, env.GetId(), app.GetId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetApplication`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Got application with id %q\n", getAppResp.GetId())

	// List applications in an environment
	listEnvAppsResp, err := scaClient.DefaultAPI.ListApplications(context.Background(), projectID, region, env.GetId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListApplications`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Number of applications in environment: %d\n", len(listEnvAppsResp.Items))

	logs, err := scaClient.DefaultAPI.GetApplicationLogs(context.Background(), projectID, region, env.GetId(), app.GetId()).
		Instance(getAppResp.Status.Instances[0].GetName()).
		Container(getAppResp.Containers[0].GetName()).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetApplicationLogs`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d logs for application %q\n", len(logs.GetLogs()), app.GetId())

	events, err := scaClient.DefaultAPI.GetApplicationEvents(context.Background(), projectID, region, env.GetId(), app.GetId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetApplicationEvents`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d events for application %q\n", len(events.GetEvents()), app.GetId())

	// Update application
	updateApplicationPayload := sca.UpdateApplicationPayload{
		Scaling: &sca.ApplicationScaling{
			Type: sca.APPLICATIONSCALINGTYPE_SCALING_TYPE_AUTO,
			AutoScaling: &sca.ApplicationAutoScaling{
				AllowScaleToZero: sca.PtrBool(true),
				MinInstances:     1,
				MaxInstances:     2,
				Rules: []sca.ApplicationScaleRule{{
					Name: "http-scaling-rule",
					Type: sca.APPLICATIONRULETYPE_RULE_TYPE_HTTP,
					HttpRule: &sca.ApplicationHttpScaleRule{
						Concurrency: sca.PtrInt32(10),
						Rps:         sca.PtrInt32(10),
					},
				}},
			},
		},
	}
	updateAppResp, err := scaClient.DefaultAPI.UpdateApplication(context.Background(), projectID, region, env.GetId(), app.GetId()).
		UpdateApplicationPayload(updateApplicationPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateApplication`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Triggered application update with id %q\n", app.GetId())

	_, err = wait.UpdateApplicationWaitHandler(context.Background(), scaClient.DefaultAPI, projectID, region, env.GetId(), app.GetId()).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatedApplicationWaitHandler`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Updated application with id %q\n", updateAppResp.GetId())

	deleteApplicationResp, err := scaClient.DefaultAPI.DeleteApplication(context.Background(), projectID, region, env.GetId(), app.GetId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteApplication`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Triggered application deletion with id %q\n", *deleteApplicationResp.Id)

	_, err = wait.DeleteApplicationWaitHandler(context.Background(), scaClient.DefaultAPI, projectID, region, env.GetId(), app.GetId()).WaitWithContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteApplicationWaitHandler`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Application deleted with id %q\n", app.GetId())

	overview, err := scaClient.DefaultAPI.GetOverview(context.Background(), projectID, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetOverview`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Got project %q overview\n", projectID)
	fmt.Printf("Number of applications in project %q: %d\n", projectID, *overview.GetSummary().ApplicationsTotal)
	fmt.Printf("Number of running applications in project %q: %d\n", projectID, *overview.GetSummary().ApplicationsRunning)

	_, err = scaClient.DefaultAPI.UpdateEnvironmentRegistries(context.Background(), projectID, region, env.GetId()).UpdateEnvironmentRegistriesPayload(sca.UpdateEnvironmentRegistriesPayload{
		Registries: []sca.EnvironmentRegistry{{
			Server:   "index.docker.io",
			Username: sca.PtrString("test-user"),
			Password: sca.PtrString("test-password"),
		}},
	}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateEnvironmentRegistries`: %v\n", err)
		os.Exit(1)
	}

	registries, err := scaClient.DefaultAPI.ListEnvironmentRegistries(context.Background(), projectID, region, env.GetId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListEnvironmentRegistries`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Got %d registries\n", len(registries.GetRegistries()))

	_, err = scaClient.DefaultAPI.UpdateEnvironmentRegistries(context.Background(), projectID, region, env.GetId()).UpdateEnvironmentRegistriesPayload(sca.UpdateEnvironmentRegistriesPayload{
		Registries: nil,
	}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateEnvironmentRegistries`: %v\n", err)
		os.Exit(1)
	}

	_, err = scaClient.DefaultAPI.UpdateEnvironmentSecrets(context.Background(), projectID, region, env.GetId()).UpdateEnvironmentSecretsPayload(sca.UpdateEnvironmentSecretsPayload{
		Secrets: []sca.EnvironmentSecret{{
			Key:   "secret-key",
			Value: sca.PtrString("secret-val"),
		}},
	}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateEnvironmentSecrets`: %v\n", err)
		os.Exit(1)
	}

	secrets, err := scaClient.DefaultAPI.ListEnvironmentSecrets(context.Background(), projectID, region, env.GetId()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListEnvironmentSecrets`: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Got %d secrets\n", len(secrets.GetSecrets()))

	_, err = scaClient.DefaultAPI.UpdateEnvironmentSecrets(context.Background(), projectID, region, env.GetId()).UpdateEnvironmentSecretsPayload(sca.UpdateEnvironmentSecretsPayload{
		Secrets: nil,
	}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateEnvironmentSecrets`: %v\n", err)
		os.Exit(1)
	}
}
