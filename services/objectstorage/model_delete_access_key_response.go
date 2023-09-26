/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage.  **Disclaimer**: The API is still under development.

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

type DeleteAccessKeyResponse struct {
	// Identifies the pair of access key and secret access key for deletion
	// REQUIRED
	KeyId *string `json:"keyId"`
	// Project ID
	// REQUIRED
	Project *string `json:"project"`
}
