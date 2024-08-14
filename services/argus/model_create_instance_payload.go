/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

// CreateInstancePayload Create update instance body.
type CreateInstancePayload struct {
	// Name of the service
	Name *string `json:"name,omitempty"`
	// additional parameters
	Parameter *map[string]interface{} `json:"parameter,omitempty"`
	// uuid of the plan to create/update
	// REQUIRED
	PlanId *string `json:"planId"`
}
