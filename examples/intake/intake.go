package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	intake "github.com/stackitcloud/stackit-sdk-go/services/intake/v1betaapi"
)

func main() {
	region := "eu01"          // Region where the resources will be created
	projectId := "PROJECT_ID" // Your STACKIT project ID

	dremioCatalogURI := "DREMIO_CATALOG_URI"       //nolint:gosec // E.g., "https://my-dremio-catalog.data-platform.stackit.run/iceberg"
	dremioTokenEndpoint := "DREMIO_TOKEN_ENDPOINT" //nolint:gosec // E.g., "https://my-dremio.data-platform.stackit.run/oauth/token"
	dremioPAT := "DREMIO_PERSONAL_ACCESS_TOKEN"    //nolint:gosec // Your Dremio Personal Access Token
	catalogWarehouse := "CATALOG_WAREHOUSE"        //nolint:gosec // Catalog warehouse where the data will be ingested

	intakeUserPassword := "s3cuRe_p@ssW0rd_f0r_1ntake!" //nolint:gosec // A secure password for the new intake user

	ctx := context.Background()

	intakeClient, err := intake.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Create an Intake Runner
	createRunnerPayload := intake.CreateIntakeRunnerPayload{
		DisplayName:        "my-example-runner",
		MaxMessageSizeKiB:  int32(10),
		MaxMessagesPerHour: int32(1000),
	}
	createRunnerResp, err := intakeClient.DefaultAPI.CreateIntakeRunner(ctx, projectId, region).CreateIntakeRunnerPayload(createRunnerPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating Intake Runner: %v\n", err)
		os.Exit(1)
	}
	intakeRunnerId := createRunnerResp.Id
	fmt.Printf("Triggered creation of Intake Runner with ID: %s. Waiting for it to become active...\n", intakeRunnerId)

	// Create an Intake
	dremioAuthType := intake.CatalogAuthType("dremio") // can also be set to "none" if the catalog is not authenticated
	createIntakePayload := intake.CreateIntakePayload{
		DisplayName:    "my-example-intake",
		IntakeRunnerId: intakeRunnerId,
		Catalog: intake.IntakeCatalog{
			Uri:       dremioCatalogURI,
			Warehouse: catalogWarehouse,
			Namespace: utils.Ptr("example_namespace"),
			TableName: utils.Ptr("example_table"),
			Auth: &intake.CatalogAuth{
				Type: dremioAuthType,
				Dremio: &intake.DremioAuth{
					TokenEndpoint:       dremioTokenEndpoint,
					PersonalAccessToken: dremioPAT,
				},
			},
		},
	}
	createIntakeResp, err := intakeClient.DefaultAPI.CreateIntake(ctx, projectId, region).CreateIntakePayload(createIntakePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating Intake: %v\n", err)
		os.Exit(1)
	}
	intakeId := createIntakeResp.Id
	fmt.Printf("Triggered creation of Intake with ID: %s. Waiting for it to become active...\n", intakeRunnerId)

	createIntakeUserPayload := intake.CreateIntakeUserPayload{
		DisplayName: "my-example-user",
		Password:    intakeUserPassword,
	}
	// Create an Intake user
	createIntakeUserResp, err := intakeClient.DefaultAPI.CreateIntakeUser(ctx, projectId, region, intakeId).CreateIntakeUserPayload(createIntakeUserPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating Intake User: %v\n", err)
		os.Exit(1)
	}
	intakeUserId := createIntakeUserResp.Id
	fmt.Printf("Created Intake User with ID: %s\n", intakeUserId)
}
