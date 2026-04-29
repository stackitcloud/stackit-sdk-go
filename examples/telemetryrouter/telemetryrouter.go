package main

import (
	"context"
	"log"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	telemetryrouter "github.com/stackitcloud/stackit-sdk-go/services/telemetryrouter/v1betaapi"
	"github.com/stackitcloud/stackit-sdk-go/services/telemetryrouter/v1betaapi/wait"
)

func main() {
	ctx := context.Background()

	projectId := "PROJECT_ID" // the uuid of your STACKIT project
	regionId := "eu01"

	client, err := telemetryrouter.NewAPIClient()
	if err != nil {
		log.Fatalf("[TelemetryRouter API] Creating API client: %v\n", err)
	}

	// Create a Telemetry Router Instance
	var createdInstance string
	createInstancePayload := telemetryrouter.CreateTelemetryRouterPayload{
		DisplayName: "my-tlmr-instance",
	}
	createResp, err := client.DefaultAPI.CreateTelemetryRouter(ctx, projectId, regionId).
		CreateTelemetryRouterPayload(createInstancePayload).
		Execute()
	if err != nil {
		log.Fatalf("[TelemetryRouter API] Error when calling `CreateTelemetryRouter`: %v\n", err)
	}
	log.Printf("[TelemetryRouter API] Created TelemetryRouter Instance with ID \"%s\".\n", createResp.Id)
	createdInstance = createResp.Id

	// Wait for the Telemetry Router Instance to be ready
	_, err = wait.CreateTelemetryRouterWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryRouter API] Error when waiting for creation: %v\n", err)
	}
	log.Printf("[TelemetryRouter API] TelemetryRouter Instance %q has been successfully created.\n", createdInstance)

	// List Telemetry Router Instances
	listResp, err := client.DefaultAPI.ListTelemetryRouters(ctx, projectId, regionId).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `ListTelemetryRouters`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Retrieved %d Logs Instances.\n", len(listResp.TelemetryRouters))

	// Get the created Telemetry Router Instance
	getResp, err := client.DefaultAPI.GetTelemetryRouter(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `GetTelemetryRouter`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Retrieved TelemetryRouter Instance with ID \"%s\" and Display Name \"%s\".\n", getResp.Id, getResp.DisplayName)

	// Update the created Telemetry Router Instance
	updatePayload := telemetryrouter.UpdateTelemetryRouterPayload{
		DisplayName: utils.Ptr("my-updated-tlmr-instance"),
		Filter: &telemetryrouter.ConfigFilter{
			Attributes: []telemetryrouter.ConfigFilterAttributes{
				{
					Key:     "service.name",
					Level:   "logRecord",
					Matcher: "=",
					Values:  []string{"service1", "service2"},
				},
			},
		},
	}
	updateResp, err := client.DefaultAPI.UpdateTelemetryRouter(ctx, projectId, regionId, createdInstance).
		UpdateTelemetryRouterPayload(updatePayload).
		Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `UpdateTelemetryRouter`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Updated Telemetry Router Instance with ID \"%s\" to Display Name \"%s\".\n", updateResp.Id, updateResp.DisplayName)

	// Wait for the Telemetry Router Instance to be updated
	_, err = wait.UpdateTelemetryRouterWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for update: %v\n", err)
	}
	log.Printf("[Telemetry Router API] TelemetryRouter Instance %q has been successfully updated.\n", createdInstance)

	// Create an Access Token
	createTokenPayload := telemetryrouter.CreateAccessTokenPayload{
		DisplayName: "my-access-token",
	}
	createTokenResp, err := client.DefaultAPI.CreateAccessToken(ctx, projectId, regionId, createdInstance).
		CreateAccessTokenPayload(createTokenPayload).
		Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `CreateAccessToken`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Created Access Token with ID \"%s\".\n", createTokenResp.Id)

	// Wait for the Access Token to be created
	_, err = wait.CreateAccessTokenWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, createTokenResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for creation: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Access Token %q has been successfully created.\n", createTokenResp.Id)

	// Get the created Access Token
	getTokenResp, err := client.DefaultAPI.GetAccessToken(ctx, projectId, regionId, createdInstance, createTokenResp.Id).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `GetAccessToken`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Retrieved Access Token with ID \"%s\" and Display Name \"%s\".\n", getTokenResp.Id, getTokenResp.DisplayName)

	// List Access Tokens
	listTokenResp, err := client.DefaultAPI.ListAccessTokens(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `ListAccessTokens`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Retrieved %d Access Tokens.\n", len(listTokenResp.AccessTokens))

	newDisplayName := "my-updated-acc-token"
	newDisplayNameNS := telemetryrouter.NullableString{}
	newDisplayNameNS.Set(&newDisplayName)
	updateTokenPayload := telemetryrouter.UpdateAccessTokenPayload{
		DisplayName: newDisplayNameNS,
	}
	// Update an Access Token
	updateTokenResp, err := client.DefaultAPI.UpdateAccessToken(ctx, projectId, regionId, createdInstance, createTokenResp.Id).
		UpdateAccessTokenPayload(updateTokenPayload).
		Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `UpdateAccessToken`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Updated Access Token with ID \"%s\" to Display Name \"%s\".\n.\n", updateTokenResp.Id, updateTokenResp.DisplayName)

	// Wait for the Access Token to be updated
	_, err = wait.UpdateAccessTokenWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, updateTokenResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for update: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Access Token %q has been successfully updated.\n", updateTokenResp.Id)

	// Delete Access Token
	err = client.DefaultAPI.DeleteAccessToken(ctx, projectId, regionId, createdInstance, updateTokenResp.Id).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `DeleteAccessToken`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Deleted Access Tokens.\n")

	// Wait for the Access Token to be deleted
	_, err = wait.DeleteAccessTokenWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, updateTokenResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for deletion: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Access Token %q has been successfully deleted.\n", updateTokenResp.Id)

	// Create S3 Destination
	createS3DestinationPayload := telemetryrouter.CreateDestinationPayload{
		DisplayName: "s3-destination",
		Config: telemetryrouter.DestinationConfig{
			ConfigType: "S3",
			S3: &telemetryrouter.DestinationConfigS3{
				AccessKey: &telemetryrouter.DestinationConfigS3AccessKey{
					Id:     "id",
					Secret: "secret",
				},
				Bucket:   "bucket",
				Endpoint: "https://bucket.endpoint.com",
			},
		},
	}
	createS3DestinationResp, err := client.DefaultAPI.CreateDestination(ctx, projectId, regionId, createdInstance).
		CreateDestinationPayload(createS3DestinationPayload).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `CreateDestination`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Created Destination with ID \"%s\".\n", createS3DestinationResp.Id)

	// Wait for the S3 Destination to be created
	_, err = wait.CreateDestinationWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, createS3DestinationResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for creation: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Destination %q has been successfully created.\n", createS3DestinationResp.Id)

	// Get the created S3 Destination
	getS3DestinationResp, err := client.DefaultAPI.GetDestination(ctx, projectId, regionId, createdInstance, createS3DestinationResp.Id).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `GetDestination`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Retrieved Destination with ID \"%s\" and Display Name \"%s\".\n", getS3DestinationResp.Id, getS3DestinationResp.DisplayName)

	// List Destinations
	listDestinationResp, err := client.DefaultAPI.ListDestinations(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `ListDestinations`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Retrieved %d Destinations.\n", len(listDestinationResp.Destinations))

	// Update S3 Destination
	newS3DestinationDisplayName := "updated-s3-destination"
	updateS3DestinationPayload := telemetryrouter.UpdateDestinationPayload{
		DisplayName: &newS3DestinationDisplayName,
		Config: &telemetryrouter.DestinationConfig{
			ConfigType: "S3",
			S3: &telemetryrouter.DestinationConfigS3{
				AccessKey: &telemetryrouter.DestinationConfigS3AccessKey{
					Id:     "id",
					Secret: "secret",
				},
				Bucket:   "bucket",
				Endpoint: "https://bucket.endpoint.com",
			},
			Filter: &telemetryrouter.ConfigFilter{
				Attributes: []telemetryrouter.ConfigFilterAttributes{
					{
						Key:     "service.name",
						Level:   "logRecord",
						Matcher: "=",
						Values:  []string{"service1", "service2"},
					},
				},
			},
		},
	}
	updateS3DestinationResp, err := client.DefaultAPI.UpdateDestination(ctx, projectId, regionId, createdInstance, createS3DestinationResp.Id).
		UpdateDestinationPayload(updateS3DestinationPayload).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `UpdateDestination`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Updated Destination with ID \"%s\".\n", updateS3DestinationResp.Id)

	// Wait for the S3 Destination to be updated
	_, err = wait.UpdateDestinationWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, updateS3DestinationResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for update: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Destination %q has been successfully updated.\n", updateS3DestinationResp.Id)

	// Delete S3 Destination
	err = client.DefaultAPI.DeleteDestination(ctx, projectId, regionId, createdInstance, updateS3DestinationResp.Id).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `DeleteDestination`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Deleted Destination with ID \"%s\".\n", updateS3DestinationResp.Id)

	// Wait for the S3 Destination to be deleted
	_, err = wait.DeleteDestinationWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, updateS3DestinationResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for deletion: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Destination %q has been successfully deleted.\n", updateS3DestinationResp.Id)

	// Create OpenTelemetry Destination
	createOpenTelemetryDestinationPayload := telemetryrouter.CreateDestinationPayload{
		DisplayName: "otlp-destination",
		Config: telemetryrouter.DestinationConfig{
			ConfigType: "OpenTelemetry",
			OpenTelemetry: &telemetryrouter.DestinationConfigOpenTelemetry{
				BasicAuth: &telemetryrouter.DestinationConfigOpenTelemetryBasicAuth{
					Password: "password",
					Username: "user",
				},
				Uri: "https://otlp.endpoint.com/v1/logs",
			},
		},
	}
	createOpenTelemetryDestinationResp, err := client.DefaultAPI.CreateDestination(ctx, projectId, regionId, createdInstance).
		CreateDestinationPayload(createOpenTelemetryDestinationPayload).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `CreateDestination`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Created Destination with ID \"%s\".\n", createOpenTelemetryDestinationResp.Id)

	// Wait for the OpenTelemetry Destination to be created
	_, err = wait.CreateDestinationWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, createOpenTelemetryDestinationResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for creation: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Destination %q has been successfully created.\n", createOpenTelemetryDestinationResp.Id)

	// Get the created OpenTelemetry Destination
	getOpenTelemetryDestinationResp, err := client.DefaultAPI.GetDestination(ctx, projectId, regionId, createdInstance, createOpenTelemetryDestinationResp.Id).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `GetDestination`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Received Destination with ID \"%s\".\n", getOpenTelemetryDestinationResp.Id)

	// List Destinations
	listDestinationsResp, err := client.DefaultAPI.ListDestinations(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `ListDestinations`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Received %d Destinations.\n", len(listDestinationsResp.Destinations))

	// Update S3 Destination
	newOTLPDestinationDisplayName := "updated-otlp-destination"
	updateOTLPDestinationPayload := telemetryrouter.UpdateDestinationPayload{
		DisplayName: &newOTLPDestinationDisplayName,
		Config: &telemetryrouter.DestinationConfig{
			ConfigType: "OpenTelemetry",
			OpenTelemetry: &telemetryrouter.DestinationConfigOpenTelemetry{
				BasicAuth: &telemetryrouter.DestinationConfigOpenTelemetryBasicAuth{
					Username: "user1",
					Password: "pass",
				},
				Uri: "https://otlp.endpoint.com/v1/logs",
			},
			Filter: &telemetryrouter.ConfigFilter{
				Attributes: []telemetryrouter.ConfigFilterAttributes{
					{
						Key:     "service.name",
						Level:   "logRecord",
						Matcher: "=",
						Values:  []string{"service1", "service2"},
					},
				},
			},
		},
	}
	updateOTLPDestinationResp, err := client.DefaultAPI.UpdateDestination(ctx, projectId, regionId, createdInstance, createS3DestinationResp.Id).
		UpdateDestinationPayload(updateOTLPDestinationPayload).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `UpdateDestination`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Updated Destination with ID \"%s\".\n", updateOTLPDestinationResp.Id)

	// Wait for the OpenTelemetry Destination to be updated
	_, err = wait.UpdateDestinationWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, updateOTLPDestinationResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for update: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Destination %q has been successfully updated.\n", updateOTLPDestinationResp.Id)

	// Delete OpenTelemetry Destination
	err = client.DefaultAPI.DeleteDestination(ctx, projectId, regionId, createdInstance, updateOTLPDestinationResp.Id).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `DeleteDestination`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Deleted Destination with ID \"%s\".\n", updateOTLPDestinationResp.Id)

	// Wait for the OpenTelemetry Destination to be deleted
	_, err = wait.DeleteDestinationWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, updateOTLPDestinationResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for deletion: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Destination %q has been successfully deleted.\n", updateOTLPDestinationResp.Id)

	// Delete the created TelemetryRouter Instance
	err = client.DefaultAPI.DeleteTelemetryRouter(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `DeleteTelemetryRouter`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Deleted Telemetry Router Instance with ID \"%s\".\n", createdInstance)

	// Wait for the TelemetryRouter Instance to be deleted
	_, err = wait.DeleteTelemetryRouterWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for deletion: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Telemetry Router Instance %q has been successfully deleted.\n", createdInstance)
}
