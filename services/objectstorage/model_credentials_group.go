/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage.  **Disclaimer**: The API is still under development.

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

type CredentialsGroup struct {
	// The ID of the credentials group
	// REQUIRED
	CredentialsGroupId *string `json:"credentialsGroupId"`
	// Name of the group holding credentials
	// REQUIRED
	DisplayName *string `json:"displayName"`
	// Credentials group URN
	// REQUIRED
	Urn *string `json:"urn"`
}
