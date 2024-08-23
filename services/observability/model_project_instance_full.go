/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

// ProjectInstanceFull struct for ProjectInstanceFull
type ProjectInstanceFull struct {
	// REQUIRED
	Id *string `json:"id"`
	// REQUIRED
	Instance *string `json:"instance"`
	Name     *string `json:"name,omitempty"`
	// REQUIRED
	PlanName *string `json:"planName"`
	// REQUIRED
	ServiceName *string `json:"serviceName"`
	// REQUIRED
	Status *string `json:"status"`
}
