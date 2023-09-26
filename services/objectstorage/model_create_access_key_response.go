/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage.  **Disclaimer**: The API is still under development.

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

type CreateAccessKeyResponse struct {
	// Access key
	// REQUIRED
	AccessKey *string `json:"accessKey"`
	// Obfuscated access key
	// REQUIRED
	DisplayName *string `json:"displayName"`
	// Expiration date. Null means never expires.
	// REQUIRED
	Expires *string `json:"expires"`
	// Identifies the pair of access key and secret access key for deletion
	// REQUIRED
	KeyId *string `json:"keyId"`
	// Project ID
	// REQUIRED
	Project *string `json:"project"`
	// Secret access key
	// REQUIRED
	SecretAccessKey *string `json:"secretAccessKey"`
}
