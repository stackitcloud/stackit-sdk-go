/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus2

type Receivers struct {
	EmailConfigs *[]EmailConfig `json:"emailConfigs,omitempty"`
	// REQUIRED
	Name            *string           `json:"name"`
	OpsgenieConfigs *[]OpsgenieConfig `json:"opsgenieConfigs,omitempty"`
	WebHookConfigs  *[]WebHook        `json:"webHookConfigs,omitempty"`
}
