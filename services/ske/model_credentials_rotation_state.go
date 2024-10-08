/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

// CredentialsRotationState struct for CredentialsRotationState
type CredentialsRotationState struct {
	// Format: `2024-02-15T11:06:29Z`
	LastCompletionTime *string `json:"lastCompletionTime,omitempty"`
	// Format: `2024-02-15T11:06:29Z`
	LastInitiationTime *string `json:"lastInitiationTime,omitempty"`
	// Phase of the credentials rotation. `NEVER` indicates that no credentials rotation has been performed using the new credentials rotation endpoints yet. Using the deprecated [rotate-credentials](#tag/Credentials/operation/SkeService_GetClusterCredentials) endpoint will not update this status field.
	Phase *string `json:"phase,omitempty"`
}
