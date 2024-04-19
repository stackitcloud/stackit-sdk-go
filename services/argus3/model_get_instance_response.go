/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus3

type GetInstanceResponse struct {
	// REQUIRED
	DashboardUrl *string         `json:"dashboardUrl"`
	Error        *NullableString `json:"error,omitempty"`
	// REQUIRED
	Id *string `json:"id"`
	// REQUIRED
	Instance    *InstanceSensitiveData `json:"instance"`
	IsUpdatable *bool                  `json:"isUpdatable,omitempty"`
	// REQUIRED
	Message    *string            `json:"message"`
	Name       *string            `json:"name,omitempty"`
	Parameters *map[string]string `json:"parameters,omitempty"`
	// REQUIRED
	PlanId *string `json:"planId"`
	// REQUIRED
	PlanName   *string `json:"planName"`
	PlanSchema *string `json:"planSchema,omitempty"`
	// REQUIRED
	ServiceName *string `json:"serviceName"`
	// REQUIRED
	Status *string `json:"status"`
}