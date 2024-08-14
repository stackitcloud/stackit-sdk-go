/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

// Bucket struct for Bucket
type Bucket struct {
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Region *string `json:"region"`
	// URL in path style
	// REQUIRED
	UrlPathStyle *string `json:"urlPathStyle"`
	// URL in virtual hosted style
	// REQUIRED
	UrlVirtualHostedStyle *string `json:"urlVirtualHostedStyle"`
}
