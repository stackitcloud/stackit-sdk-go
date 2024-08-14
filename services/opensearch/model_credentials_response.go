/*
STACKIT Opensearch API

The STACKIT Opensearch API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opensearch

// CredentialsResponse struct for CredentialsResponse
type CredentialsResponse struct {
	// REQUIRED
	Id  *string         `json:"id"`
	Raw *RawCredentials `json:"raw,omitempty"`
	// REQUIRED
	Uri *string `json:"uri"`
}
