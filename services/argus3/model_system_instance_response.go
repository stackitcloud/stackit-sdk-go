/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus3

type SystemInstanceResponse struct {
	// REQUIRED
	Instance *SystemInstance `json:"instance"`
	// REQUIRED
	Message *string `json:"message"`
}
