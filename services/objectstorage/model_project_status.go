/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

// ProjectStatus struct for ProjectStatus
type ProjectStatus struct {
	// Project ID
	// REQUIRED
	Project *string `json:"project"`
	// REQUIRED
	Scope *ProjectScope `json:"scope"`
}
