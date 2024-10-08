/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

// DeleteAccessKeyResponse struct for DeleteAccessKeyResponse
type DeleteAccessKeyResponse struct {
	// Identifies the pair of access key and secret access key for deletion
	// REQUIRED
	KeyId *string `json:"keyId"`
	// Project ID
	// REQUIRED
	Project *string `json:"project"`
}
