/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

// GetCredentialsResponse struct for GetCredentialsResponse
type GetCredentialsResponse struct {
	CredentialsInfo *map[string]string `json:"credentialsInfo,omitempty"`
	// REQUIRED
	Id *string `json:"id"`
	// REQUIRED
	Message *string `json:"message"`
	// REQUIRED
	Name *string `json:"name"`
}
