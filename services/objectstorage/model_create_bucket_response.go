/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

type CreateBucketResponse struct {
	// Name of the bucket
	// REQUIRED
	Bucket *string `json:"bucket"`
	// Project ID
	// REQUIRED
	Project *string `json:"project"`
}
