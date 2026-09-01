package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	ufw "github.com/stackitcloud/stackit-sdk-go/services/ufw/v1api"
	"github.com/stackitcloud/stackit-sdk-go/services/ufw/v1api/wait"
)

func main() {
	region := "eu01"              // Region where the resources will be created
	projectId := "PROJECT_ID"     // UUID of your STACKIT project
	instanceId := "INSTANCE_ID"   // UUID of the instance to which the firewall rule will be attached
	productType := "PRODUCT_TYPE" // Type of the instance to which the firewall rule will be attached (e.g. "edge-cloud", but you can get them from provider-options route)
	ufwRuleType := "ACL"          // Type of the rule that you want to create (ACL)

	ctx := context.Background()

	token := ""
	ufwClient, err := ufw.NewAPIClient(config.WithToken(token))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[UFW] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// List all firewall rules
	listUFWRules(ctx, ufwClient, projectId, region)

	// Create a new firewall rule
	description := "Created from SDK"
	rulePayloadToCreate := ufw.CreateRulePayload{
		InstanceId:  instanceId,
		Product:     productType,
		SourceIP:    "11.11.11.11/32",
		Type:        ufwRuleType,
		Description: &description,
	}

	createdRuleResponse, err := createFirewallRule(ctx, ufwClient, projectId, region, &rulePayloadToCreate)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[UFW] Error when creating firewall rule: %v\n", err)
		return
	}

	prettyJSON, err := json.MarshalIndent(createdRuleResponse, "", "  ")
	if err != nil {
		log.Printf("Failed to marshal rule: %v\n", err)
	}
	fmt.Printf("Created firewall rule response: %s\n", string(prettyJSON))

	// Get the firewall rule
	testGetRule, err := getFirewallRule(ctx, ufwClient, projectId, region, *createdRuleResponse.RefId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[UFW] Error when getting firewall rule: %v\n", err)
		return
	}

	prettyJSON, err = json.MarshalIndent(testGetRule, "", "  ")
	if err != nil {
		log.Printf("Failed to marshal rule: %v\n", err)
	}
	fmt.Printf("Firewall rule details: %s\n", string(prettyJSON))

	if err := verifyPayloadMatch(testGetRule, rulePayloadToCreate); err != nil {
		fmt.Fprintf(os.Stderr, "[UFW] Verification failed after creation:\n%v\n", err)
		return
	}
	fmt.Println("Created rule fields verified successfully.")

	// Update the firewall rule
	rulePayloadToUpdate := ufw.UpdateRulePayload{
		SourceIP: "22.22.22.22/32",
	}

	updatedRuleResponse, err := updateFirewallRule(ctx, ufwClient, projectId, region, *createdRuleResponse.RefId, rulePayloadToUpdate)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[UFW] Error when updating firewall rule: %v\n", err)
		return
	}

	prettyJSON, err = json.MarshalIndent(updatedRuleResponse, "", "  ")
	if err != nil {
		log.Printf("Failed to marshal rule: %v\n", err)
	}
	fmt.Printf("Updated firewall rule details: %s\n", string(prettyJSON))

	testGetRule, err = getFirewallRule(ctx, ufwClient, projectId, region, *updatedRuleResponse.RefId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[UFW] Error when getting updated firewall rule: %v\n", err)
		return
	}

	if err := verifyPayloadMatch(testGetRule, rulePayloadToUpdate); err != nil {
		fmt.Fprintf(os.Stderr, "[UFW] Verification failed after update:\n%v\n", err)
		return
	}
	fmt.Println("Updated rule fields verified successfully.")

	// Delete the firewall rule
	err = deleteFirewallRule(ctx, ufwClient, projectId, region, *updatedRuleResponse.RefId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[UFW] Error when deleting firewall rule: %v\n", err)
		return
	}

	_, err = getFirewallRule(ctx, ufwClient, projectId, region, *updatedRuleResponse.RefId)
	if !strings.Contains(err.Error(), "404") {
		fmt.Fprintf(os.Stderr, "[UFW] Error while verifying deleted rule: %v\n", err)
		return
	}

	fmt.Println("All firewall rules successfully tested and cleaned up.")
}

func listUFWRules(ctx context.Context, ufwClient *ufw.APIClient, projectId, region string) {
	listRulesResponse, err := ufwClient.DefaultAPI.ListRules(ctx, projectId, region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[UFW] Error when listing firewall rules: %v\n", err)
		return
	}

	fmt.Println("List of firewall rules:")
	for i := range listRulesResponse.Rules {
		prettyJSON, err := json.MarshalIndent(listRulesResponse.Rules[i], "", "  ")
		if err != nil {
			log.Printf("Failed to marshal rule: %v\n", err)
			continue
		}
		fmt.Println(string(prettyJSON))
	}
}

func getFirewallRule(ctx context.Context, ufwClient *ufw.APIClient, projectId, region, ruleId string) (*ufw.RuleResponse, error) {
	return ufwClient.DefaultAPI.GetRule(ctx, projectId, region, ruleId).Execute()
}

func createFirewallRule(ctx context.Context, ufwClient *ufw.APIClient, projectId, region string, payload *ufw.CreateRulePayload) (*ufw.SecurityRuleSuccessfullyCreatedResponse, error) {
	createdFirewallRule, err := ufwClient.DefaultAPI.CreateRule(ctx, projectId, region).CreateRulePayload(*payload).Execute()
	if err != nil {
		return nil, err
	}

	createdFirewallRuleId := createdFirewallRule.RefId
	fmt.Printf("[UFW] Triggered creation of firewall rule with ID: %s\n", *createdFirewallRuleId)

	_, err = wait.CreateRuleWaitHandler(ctx, ufwClient.DefaultAPI, projectId, region, *createdFirewallRuleId).WaitWithContext(ctx)
	if err != nil {
		return nil, err
	}

	return createdFirewallRule, nil
}

func updateFirewallRule(ctx context.Context, ufwClient *ufw.APIClient, projectId, region, ruleId string, payload ufw.UpdateRulePayload) (*ufw.SecurityRuleSuccessfullyCreatedResponse, error) {
	updatedFirewallRule, err := ufwClient.DefaultAPI.UpdateRule(ctx, projectId, region, ruleId).UpdateRulePayload(payload).Execute()
	if err != nil {
		return nil, err
	}

	fmt.Printf("[UFW] Triggered update of firewall rule with ID: %s\n", *updatedFirewallRule.RefId)

	_, err = wait.UpdateRuleWaitHandler(ctx, ufwClient.DefaultAPI, projectId, region, *updatedFirewallRule.RefId).WaitWithContext(ctx)
	if err != nil {
		return nil, err
	}

	return updatedFirewallRule, nil
}

func deleteFirewallRule(ctx context.Context, ufwClient *ufw.APIClient, projectId, region, ruleId string) error {
	_, err := ufwClient.DefaultAPI.DeleteRule(ctx, projectId, region, ruleId).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("[UFW] Triggered deletion of firewall rule with ID: %s\n", ruleId)

	_, err = wait.DeleteRuleWaitHandler(ctx, ufwClient.DefaultAPI, projectId, region, ruleId).WaitWithContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func verifyPayloadMatch(actual, expected any) error {
	var mismatches []string

	actualVal := reflect.ValueOf(actual)
	if actualVal.Kind() == reflect.Pointer {
		actualVal = actualVal.Elem()
	}

	expectedVal := reflect.ValueOf(expected)
	if expectedVal.Kind() == reflect.Pointer {
		expectedVal = expectedVal.Elem()
	}

	expectedType := expectedVal.Type()

	for i := 0; i < expectedVal.NumField(); i++ {
		fieldName := expectedType.Field(i).Name

		if fieldName == "AdditionalProperties" {
			continue // AdditionalProperties field is not part of the API response struct
		}

		if fieldName == "Description" {
			continue // Description field is overridden by the API, will be fixed in the future
		}

		expectedField := expectedVal.Field(i)

		// Look for a matching field in the API response struct
		actualField := actualVal.FieldByName(fieldName)
		if !actualField.IsValid() {
			continue // Field exists in payload but not in response struct, safe to skip
		}

		// Dereference pointers cleanly to get string representations
		expStr := formatReflectValue(expectedField)
		actStr := formatReflectValue(actualField)

		// Skip uninitialized/nil fields in the expected payload
		// (e.g. fields omitted from an UpdateRulePayload)
		if expStr == "" || expStr == "<nil>" {
			continue
		}

		if expStr != actStr {
			mismatches = append(mismatches, fmt.Sprintf("%s: expected %q, got %q", fieldName, expStr, actStr))
		}
	}

	if len(mismatches) > 0 {
		return fmt.Errorf("field mismatches found:\n- %s", strings.Join(mismatches, "\n- "))
	}
	return nil
}

func formatReflectValue(v reflect.Value) string {
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return "<nil>"
		}
		v = v.Elem()
	}
	return fmt.Sprintf("%v", v.Interface())
}
