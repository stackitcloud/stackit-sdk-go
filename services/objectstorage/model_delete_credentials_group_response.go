/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

type DeleteCredentialsGroupResponse struct {
	// The ID of the credentials group
	// REQUIRED
	CredentialsGroupId *string `json:"credentialsGroupId"`
	// Project ID
	// REQUIRED
	Project *string `json:"project"`
}
