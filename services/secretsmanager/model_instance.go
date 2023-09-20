/*
STACKIT Secrets Manager API

This API provides endpoints for managing the Secrets-Manager.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretsmanager

type Instance struct {
	// The API endpoint for connecting to the secrets engine.
	// REQUIRED
	ApiUrl *string `json:"apiUrl"`
	// The date and time the creation of the Secrets Manager instance was finished.
	CreationFinishedDate *string `json:"creationFinishedDate,omitempty"`
	// The date and time the creation of the Secrets Manager instance was triggered.
	// REQUIRED
	CreationStartDate *string `json:"creationStartDate"`
	// A auto generated unique id which identifies the secrets manager instances.
	// REQUIRED
	Id *string `json:"id"`
	// A user chosen name to distinguish multiple secrets manager instances.
	// REQUIRED
	Name *string `json:"name"`
	// The number of secrets currently stored inside of the instance. This value will be updated once per hour.
	// REQUIRED
	SecretCount *int32 `json:"secretCount"`
	// The name of the secrets engine.
	// REQUIRED
	SecretsEngine *string `json:"secretsEngine"`
	// The current state of the Secrets Manager instance.
	// REQUIRED
	State              *string `json:"state"`
	UpdateFinishedDate *string `json:"updateFinishedDate,omitempty"`
	UpdateStartDate    *string `json:"updateStartDate,omitempty"`
}
