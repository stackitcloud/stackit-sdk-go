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
	regionId := string(telemetryrouter.LISTTELEMETRYROUTERSREGIONIDPARAMETER_EU01)

	client, err := telemetryrouter.NewAPIClient()
	if err != nil {
		log.Fatalf("[TelemetryRouter API] Creating API client: %v\n", err)
	}

	// Create a Telemetry Router Instance
	log.Printf("[TelemetryRouter API] Creating TelemetryRouter Instance.\n")
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
	log.Printf("[TelemetryRouter API] Waiting for TelemetryRouter Instance to be created.\n")
	_, err = wait.CreateTelemetryRouterWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[TelemetryRouter API] Error when waiting for creation: %v\n", err)
	}
	log.Printf("[TelemetryRouter API] TelemetryRouter Instance \"%s\" has been successfully created.\n", createdInstance)

	// List Telemetry Router Instances
	log.Printf("[TelemetryRouter API] Retrieving TelemetryRouter Instances.\n")
	listResp, err := client.DefaultAPI.ListTelemetryRouters(ctx, projectId, regionId).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `ListTelemetryRouters`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Retrieved %d Telemetry Router Instances.\n", len(listResp.TelemetryRouters))

	// Get the created Telemetry Router Instance
	log.Printf("[TelemetryRouter API] Getting TelemetryRouter Instance with ID \"%s\".\n", createResp.Id)
	getResp, err := client.DefaultAPI.GetTelemetryRouter(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `GetTelemetryRouter`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Retrieved TelemetryRouter Instance with ID \"%s\" and Display Name \"%s\".\n", getResp.Id, getResp.DisplayName)

	// Update the created Telemetry Router Instance
	log.Printf("[TelemetryRouter API] Updating TelemetryRouter Instance with ID \"%s\".\n", createResp.Id)
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
	log.Printf("[TelemetryRouter API] Waiting for TelemetryRouter Instance with ID \"%s\" to be updated.\n", createResp.Id)
	_, err = wait.UpdateTelemetryRouterWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for update: %v\n", err)
	}
	log.Printf("[Telemetry Router API] TelemetryRouter Instance \"%s\" has been successfully updated.\n", createdInstance)

	// Create an Access Token
	log.Printf("[TelemetryRouter API] Creating AccessToken for TelemetryRouter Instance with ID \"%s\".\n", createResp.Id)
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
	log.Printf("[TelemetryRouter API] Waiting for AccessToken for TelemetryRouter Instance with ID \"%s\" to be created.\n", createResp.Id)
	_, err = wait.CreateAccessTokenWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, createTokenResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for creation: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Access Token \"%s\" has been successfully created.\n", createTokenResp.Id)

	// Get the created Access Token
	log.Printf("[TelemetryRouter API] Retrieving the AccessToken with ID \"%s\".\n", createTokenResp.Id)
	getTokenResp, err := client.DefaultAPI.GetAccessToken(ctx, projectId, regionId, createdInstance, createTokenResp.Id).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `GetAccessToken`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Retrieved Access Token with ID \"%s\" and Display Name \"%s\".\n", getTokenResp.Id, getTokenResp.DisplayName)

	// List Access Tokens
	log.Printf("[TelemetryRouter API] Listing AccessTokens.\n")
	listTokenResp, err := client.DefaultAPI.ListAccessTokens(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `ListAccessTokens`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Retrieved %d Access Tokens.\n", len(listTokenResp.AccessTokens))

	// Update an Access Token
	log.Printf("[Telemetry Router API] Updating the Access Token with ID \"%s\".\n", createTokenResp.Id)
	newDisplayNameNS := telemetryrouter.NullableString{}
	newDisplayNameNS.Set(utils.Ptr("my-updated-acc-token"))
	updateTokenPayload := telemetryrouter.UpdateAccessTokenPayload{
		DisplayName: newDisplayNameNS,
	}
	updateTokenResp, err := client.DefaultAPI.UpdateAccessToken(ctx, projectId, regionId, createdInstance, createTokenResp.Id).
		UpdateAccessTokenPayload(updateTokenPayload).
		Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `UpdateAccessToken`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Updated Access Token with ID \"%s\" to Display Name \"%s\".\n.\n", updateTokenResp.Id, updateTokenResp.DisplayName)

	// Wait for the Access Token to be updated
	log.Printf("[Telemetry Router API] Waiting for the Access Token with ID \"%s\" to be updated.\n", createTokenResp.Id)
	_, err = wait.UpdateAccessTokenWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, updateTokenResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for update: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Access Token \"%s\" has been successfully updated.\n", updateTokenResp.Id)

	// Delete Access Token
	log.Printf("[Telemetry Router API] Deleting for the Access Token with ID \"%s\".\n", createTokenResp.Id)
	err = client.DefaultAPI.DeleteAccessToken(ctx, projectId, regionId, createdInstance, updateTokenResp.Id).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `DeleteAccessToken`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Deleted Access Tokens.\n")

	// Wait for the Access Token to be deleted
	log.Printf("[Telemetry Router API] Waiting for the Access Token with ID \"%s\" to be deleted.\n", createTokenResp.Id)
	_, err = wait.DeleteAccessTokenWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, updateTokenResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for deletion: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Access Token \"%s\" has been successfully deleted.\n", updateTokenResp.Id)

	// Create S3 Destination
	log.Printf("[Telemetry Router API] Creating Destination.\n")
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
				Endpoint: "https://portal.stackit.cloud/",
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
	log.Printf("[Telemetry Router API] Waiting for a Destination to be created.\n")
	_, err = wait.CreateDestinationWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, createS3DestinationResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for creation: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Destination \"%s\" has been successfully created.\n", createS3DestinationResp.Id)

	// Get the created S3 Destination
	log.Printf("[Telemetry Router API] Retrieving Destination with ID \"%s\".\n", createS3DestinationResp.Id)
	getS3DestinationResp, err := client.DefaultAPI.GetDestination(ctx, projectId, regionId, createdInstance, createS3DestinationResp.Id).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `GetDestination`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Retrieved Destination with ID \"%s\" and Display Name \"%s\".\n", getS3DestinationResp.Id, getS3DestinationResp.DisplayName)

	// List Destinations
	log.Printf("[Telemetry Router API] Listing Destinations.\n")
	listDestinationResp, err := client.DefaultAPI.ListDestinations(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `ListDestinations`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Retrieved %d Destinations.\n", len(listDestinationResp.Destinations))

	// Update S3 Destination
	log.Printf("[Telemetry Router API] Updating Destination with ID \"%s\".\n", createS3DestinationResp.Id)
	updateS3DestinationPayload := telemetryrouter.UpdateDestinationPayload{
		DisplayName: utils.Ptr("updated-s3-destination"),
		Config: &telemetryrouter.DestinationConfig{
			ConfigType: "S3",
			S3: &telemetryrouter.DestinationConfigS3{
				AccessKey: &telemetryrouter.DestinationConfigS3AccessKey{
					Id:     "id",
					Secret: "secret",
				},
				Bucket:   "bucket",
				Endpoint: "https://portal.stackit.cloud/",
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
	log.Printf("[Telemetry Router API] Waiting for the Destination with ID \"%s\" to be updated.\n", createS3DestinationResp.Id)
	_, err = wait.UpdateDestinationWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, updateS3DestinationResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for update: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Destination \"%s\" has been successfully updated.\n", updateS3DestinationResp.Id)

	// Delete S3 Destination
	log.Printf("[Telemetry Router API] Deleting Destination with ID \"%s\".\n", updateS3DestinationResp.Id)
	err = client.DefaultAPI.DeleteDestination(ctx, projectId, regionId, createdInstance, updateS3DestinationResp.Id).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `DeleteDestination`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Deleted Destination with ID \"%s\".\n", updateS3DestinationResp.Id)

	// Wait for the S3 Destination to be deleted
	log.Printf("[Telemetry Router API] Waiting for Destination with ID \"%s\" to be deleted.\n", updateS3DestinationResp.Id)
	_, err = wait.DeleteDestinationWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, updateS3DestinationResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for deletion: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Destination \"%s\" has been successfully deleted.\n", updateS3DestinationResp.Id)

	// Create OpenTelemetry Destination
	log.Printf("[Telemetry Router API] Creating Destination.\n")
	createOpenTelemetryDestinationPayload := telemetryrouter.CreateDestinationPayload{
		DisplayName: "otlp-destination",
		Config: telemetryrouter.DestinationConfig{
			ConfigType: "OpenTelemetry",
			OpenTelemetry: &telemetryrouter.DestinationConfigOpenTelemetry{
				BasicAuth: &telemetryrouter.DestinationConfigOpenTelemetryBasicAuth{
					Password: "password",
					Username: "user",
				},
				Uri: "https://portal.stackit.cloud/v1/logs",
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
	log.Printf("[Telemetry Router API] Waiting for Destination to be created.\n")
	_, err = wait.CreateDestinationWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, createOpenTelemetryDestinationResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for creation: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Destination \"%s\" has been successfully created.\n", createOpenTelemetryDestinationResp.Id)

	// Get the created OpenTelemetry Destination
	log.Printf("[Telemetry Router API] Receiving Destination with ID \"%s\".\n", createOpenTelemetryDestinationResp.Id)
	getOpenTelemetryDestinationResp, err := client.DefaultAPI.GetDestination(ctx, projectId, regionId, createdInstance, createOpenTelemetryDestinationResp.Id).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `GetDestination`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Received Destination with ID \"%s\".\n", getOpenTelemetryDestinationResp.Id)

	// List Destinations
	log.Printf("[Telemetry Router API] Listing Destinations.\n")
	listDestinationsResp, err := client.DefaultAPI.ListDestinations(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `ListDestinations`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Received %d Destinations.\n", len(listDestinationsResp.Destinations))

	// Update S3 Destination
	log.Printf("[Telemetry Router API] Updating Destination with ID \"%s\".\n", createOpenTelemetryDestinationResp.Id)
	updateOTLPDestinationPayload := telemetryrouter.UpdateDestinationPayload{
		DisplayName: utils.Ptr("updated-otlp-destination"),
		Config: &telemetryrouter.DestinationConfig{
			ConfigType: "OpenTelemetry",
			OpenTelemetry: &telemetryrouter.DestinationConfigOpenTelemetry{
				BasicAuth: &telemetryrouter.DestinationConfigOpenTelemetryBasicAuth{
					Username: "user1",
					Password: "pass",
				},
				Uri: "https://portal.stackit.cloud/v1/logs",
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
	updateOTLPDestinationResp, err := client.DefaultAPI.UpdateDestination(ctx, projectId, regionId, createdInstance, createOpenTelemetryDestinationResp.Id).
		UpdateDestinationPayload(updateOTLPDestinationPayload).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `UpdateDestination`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Updated Destination with ID \"%s\".\n", updateOTLPDestinationResp.Id)

	// Wait for the OpenTelemetry Destination to be updated
	log.Printf("[Telemetry Router API] Waiting for Destination \"%s\" to be updated.\n", createOpenTelemetryDestinationResp.Id)
	_, err = wait.UpdateDestinationWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, updateOTLPDestinationResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for update: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Destination \"%s\" has been successfully updated.\n", updateOTLPDestinationResp.Id)

	// Delete OpenTelemetry Destination
	log.Printf("[Telemetry Router API] Deleting Destination with ID \"%s\".\n", updateOTLPDestinationResp.Id)
	err = client.DefaultAPI.DeleteDestination(ctx, projectId, regionId, createdInstance, updateOTLPDestinationResp.Id).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `DeleteDestination`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Deleted Destination with ID \"%s\".\n", updateOTLPDestinationResp.Id)

	// Wait for the OpenTelemetry Destination to be deleted
	log.Printf("[Telemetry Router API] Waiting for Destination \"%s\" to be deleted.\n", updateOTLPDestinationResp.Id)
	_, err = wait.DeleteDestinationWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance, updateOTLPDestinationResp.Id).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for deletion: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Destination \"%s\" has been successfully deleted.\n", updateOTLPDestinationResp.Id)

	// Delete the created TelemetryRouter Instance
	log.Printf("[Telemetry Router API] Deleting Telemetry Router Instance with ID \"%s\".\n", createdInstance)
	err = client.DefaultAPI.DeleteTelemetryRouter(ctx, projectId, regionId, createdInstance).Execute()
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when calling `DeleteTelemetryRouter`: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Deleted Telemetry Router Instance with ID \"%s\".\n", createdInstance)

	// Wait for the TelemetryRouter Instance to be deleted
	log.Printf("[Telemetry Router API] Waiting for Telemetry Router Instance with ID \"%s\" to be deleted.\n", createdInstance)
	_, err = wait.DeleteTelemetryRouterWaitHandler(ctx, client.DefaultAPI, projectId, regionId, createdInstance).WaitWithContext(context.Background())
	if err != nil {
		log.Fatalf("[Telemetry Router API] Error when waiting for deletion: %v\n", err)
	}
	log.Printf("[Telemetry Router API] Telemetry Router Instance \"%s\" has been successfully deleted.\n", createdInstance)
}
