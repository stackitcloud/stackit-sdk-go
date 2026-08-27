package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	automation "github.com/stackitcloud/stackit-sdk-go/services/automation/v1betaapi"
	"github.com/stackitcloud/stackit-sdk-go/services/automation/v1betaapi/wait"
)

func main() {
	ctx := context.Background()

	region := "eu01"                                    // Region where the resources will be created
	projectId := "66b03760-125b-4c63-9624-741e8a418c02" // the uuid of your STACKIT project

	// Create a new API client, that uses default authentication and configuration
	automationClient, err := automation.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Automation API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// List the available volume templates and pick the first one
	templatesResp, err := automationClient.DefaultAPI.ListVolumeTemplates(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Automation API] Error when calling `ListVolumeTemplates`: %v\n", err)
		os.Exit(1)
	}
	if len(templatesResp.Items) == 0 {
		fmt.Fprintln(os.Stderr, "[Automation API] No volume templates available")
		os.Exit(1)
	}
	templateId := templatesResp.Items[0].Id
	fmt.Printf("[Automation API] Using volume template \"%s\" with id \"%s\".\n", templatesResp.Items[0].Name, templateId)

	// Create a volume automation with a schedule trigger (daily at 02:00)
	input := automation.VolumeRecoveryPointManagementInputAsVolumeAutomationInput(&automation.VolumeRecoveryPointManagementInput{
		Kind: "VolumeRecoveryPointManagement",
		SnapshotRetentionPolicy: automation.SnapshotRetentionPolicyCountAsSnapshotRetentionPolicy(
			automation.NewSnapshotRetentionPolicyCount(automation.SNAPSHOTRETENTIONPOLICYCOUNTKIND_COUNT, 6),
		),
	})
	createVolumeAutomationPayload := automation.CreateVolumeAutomationPayload{
		Name:        utils.Ptr("exampleVolumeAutomation"),
		Description: utils.Ptr("Example volume automation created via the STACKIT SDK for Go"),
		TemplateId:  templateId,
		Input:       &input,
		Triggers: &automation.AutomationTriggers{
			Schedule: &automation.AutomationScheduleTrigger{
				Rrule: "FREQ=DAILY;BYHOUR=2;BYMINUTE=0;BYSECOND=0",
			},
		},
	}
	automationResp, err := automationClient.DefaultAPI.CreateVolumeAutomation(ctx, projectId, region).CreateVolumeAutomationPayload(createVolumeAutomationPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Automation API] Error when calling `CreateVolumeAutomation`: %v\n", err)
		os.Exit(1)
	}
	automationId := automationResp.Id
	fmt.Printf("[Automation API] Created volume automation with id \"%s\".\n", automationId)

	// List volume automations
	listResp, err := automationClient.DefaultAPI.ListVolumeAutomations(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Automation API] Error when calling `ListVolumeAutomations`: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("[Automation API] Volume automations:")
	for _, item := range listResp.Items {
		fmt.Printf("- %s (%s)\n", item.Id, *item.Name)
	}

	// Trigger an execution of the volume automation
	executionResp, err := automationClient.DefaultAPI.CreateVolumeExecution(ctx, projectId, region, automationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Automation API] Error when calling `CreateVolumeExecution`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[Automation API] Triggered execution with id \"%s\" (status: %s).\n", executionResp.Id, executionResp.Status)

	// Wait for the execution to finish, otherwise the deletion below fails
	waitResp, err := wait.CreateVolumeExecutionWaitHandler(ctx, automationClient.DefaultAPI, projectId, region, automationId, executionResp.Id).
		SetTimeout(10 * time.Minute).
		WaitWithContext(ctx)
	if err != nil {
		// Continue with the deletion: failed or terminated executions no longer block it
		fmt.Fprintf(os.Stderr, "[Automation API] Error waiting for execution to finish: %v\n", err)
	}
	fmt.Printf("[Automation API] Finished execution with id \"%s\" (status: %s).\n", waitResp.Id, waitResp.Status)

	// Delete the volume automation
	err = automationClient.DefaultAPI.DeleteVolumeAutomation(ctx, projectId, region, automationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Automation API] Error when calling `DeleteVolumeAutomation`: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[Automation API] Deleted volume automation with id \"%s\".\n", automationId)
}
