package main

import (
	"context"
	"fmt"
	"os"

	sdkConfig "github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/intake"
	"github.com/stackitcloud/stackit-sdk-go/services/intake/wait"
)

func main() {
	region := "REGION"        // Region where the resources will be created
	projectId := "PROJECT_ID" // Your STACKIT project ID

	dremioCatalogURI := "DREMIO_CATALOG_URI"       // E.g., "https://my-dremio-catalog.data-platform.stackit.run/iceberg"
	dremioTokenEndpoint := "DREMIO_TOKEN_ENDPOINT" // E.g., "https://my-dremio.data-platform.stackit.run/oauth/token"
	dremioPAT := "DREMIO_PERSONAL_ACCESS_TOKEN"    // Your Dremio Personal Access Token
	catalogWarehouse := "CATALOG_WAREHOUSE"        // Catalog warehouse where the data will be ingested

	intakeUserPassword := "s3cuRe_p@ssW0rd_f0r_1ntake!" // A secure password for the new intake user

	ctx := context.Background()

	intakeClient, err := intake.NewAPIClient(
		sdkConfig.WithRegion(region),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create an Intake Runner
	createRunnerPayload := intake.CreateIntakeRunnerPayload{
		DisplayName:        utils.Ptr("my-example-runner"),
		MaxMessageSizeKiB:  utils.Ptr(int64(10)),
		MaxMessagesPerHour: utils.Ptr(int64(1000)),
	}
	createRunnerResp, err := intakeClient.CreateIntakeRunner(ctx, projectId, region).CreateIntakeRunnerPayload(createRunnerPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating Intake Runner: %v\n", err)
		os.Exit(1)
	}
	intakeRunnerId := *createRunnerResp.Id
	fmt.Printf("Triggered creation of Intake Runner with ID: %s. Waiting for it to become active...\n", intakeRunnerId)

	// Wait for the Intake Runner to become active
	_, err = wait.CreateOrUpdateIntakeRunnerWaitHandler(ctx, intakeClient, projectId, region, intakeRunnerId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for Intake Runner: %v\n", err)
		os.Exit(1)
	}

	// Create an Intake
	dremioAuthType := intake.CatalogAuthType("dremio") // can also be set to "none" if the catalog is not authenticated
	createIntakePayload := intake.CreateIntakePayload{
		DisplayName:    utils.Ptr("my-example-intake"),
		IntakeRunnerId: utils.Ptr(intakeRunnerId),
		Catalog: &intake.IntakeCatalog{
			Uri:       utils.Ptr(dremioCatalogURI),
			Warehouse: utils.Ptr(catalogWarehouse),
			Namespace: utils.Ptr("example_namespace"),
			TableName: utils.Ptr("example_table"),
			Auth: &intake.CatalogAuth{
				Type: &dremioAuthType,
				Dremio: &intake.DremioAuth{
					TokenEndpoint:       utils.Ptr(dremioTokenEndpoint),
					PersonalAccessToken: utils.Ptr(dremioPAT),
				},
			},
		},
	}
	createIntakeResp, err := intakeClient.CreateIntake(ctx, projectId, region).CreateIntakePayload(createIntakePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating Intake: %v\n", err)
		os.Exit(1)
	}
	intakeId := *createIntakeResp.Id
	fmt.Printf("Triggered creation of Intake with ID: %s. Waiting for it to become active...\n", intakeRunnerId)

	// Wait for the Intake to become active
	_, err = wait.CreateOrUpdateIntakeWaitHandler(ctx, intakeClient, projectId, region, intakeId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for Intake: %v\n", err)
		os.Exit(1)
	}

	createIntakeUserPayload := intake.CreateIntakeUserPayload{
		DisplayName: utils.Ptr("my-example-user"),
		Password:    utils.Ptr(intakeUserPassword),
	}
	createIntakeUserResp, err := intakeClient.CreateIntakeUser(ctx, projectId, region, intakeId).CreateIntakeUserPayload(createIntakeUserPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating Intake User: %v\n", err)
		os.Exit(1)
	}
	intakeUserId := *createIntakeUserResp.Id
	fmt.Printf("Triggered creation of Intake User with ID: %s. Waiting for it to become active...\n", intakeRunnerId)

	// Wait for the Intake User to become active
	_, err = wait.CreateOrUpdateIntakeUserWaitHandler(ctx, intakeClient, projectId, region, intakeId, intakeUserId).WaitWithContext(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for Intake User: %v\n", err)
		os.Exit(1)
	}
}
