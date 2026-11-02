package main

import (
	"context"
	"log"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	cost "github.com/stackitcloud/stackit-sdk-go/services/cost/v3api"
)

func main() {
	ctx := context.Background()

	projectID := "PROJECT_ID"                   // the uuid of your STACKIT project
	customerAccountId := "CUSTOMER_ACCOUNT_ID"  // the uuid of your customer account
	resellerAccountId := "RESELLER_ACCOUNT_ID"  // the uuid of your reseller account
	serviceAccKeyPath := "/path/to/sa-key.json" // path to your service account key

	client, err := cost.NewAPIClient(
		config.WithServiceAccountKeyPath(serviceAccKeyPath),
	)
	if err != nil {
		log.Fatalf("[Cost API] Creating API client: %v\n", err)
	}

	// list costs for customer
	customerResp, err := client.DefaultAPI.
		ListCostsForCustomer(ctx, customerAccountId).
		From("2025-10-10").
		To("2026-03-17").
		Execute()
	if err != nil {
		log.Fatalf("[Cost API] Error when calling `ListCostsForCustomer`: %v\n", err)
	}
	for _, c := range customerResp {
		log.Printf(
			"[Cost API] Costs for customer account \"%s\", project \"%s\": %+v\n",
			customerAccountId,
			c.ProjectCostWithDetailedServices.ProjectName,
			c.ProjectCostWithDetailedServices.TotalCharge,
		)
	}

	// get costs for project
	projectResp, err := client.DefaultAPI.
		GetCostsForProject(ctx, customerAccountId, projectID).
		From("2025-10-10").
		To("2026-03-17").
		Execute()
	if err != nil {
		log.Fatalf("[Cost API] Error when calling `ListCostsForProject`: %v\n", err)
	}
	log.Printf(
		"[Cost API] Costs for project \"%s\": %+v\n",
		projectResp.ProjectCostWithDetailedServices.ProjectName,
		projectResp.ProjectCostWithDetailedServices.TotalCharge,
	)

	// list costs for reseller
	resellerResp, err := client.DefaultAPI.
		ListCostsForReseller(ctx, resellerAccountId).
		From("2025-10-10").
		To("2026-03-17").
		Execute()
	if err != nil {
		log.Fatalf("[Cost API] Error when calling `ListCostsForReseller`: %v\n", err)
	}
	for _, c := range resellerResp {
		log.Printf(
			"[Cost API] Costs for reseller with customer account \"%s\", project \"%s\": %+v\n",
			resellerAccountId,
			c.ProjectCostWithDetailedServices.ProjectName,
			c.ProjectCostWithDetailedServices.TotalCharge,
		)
	}
}
