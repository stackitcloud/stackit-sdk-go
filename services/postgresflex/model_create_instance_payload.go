/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

// CreateInstancePayload struct for CreateInstancePayload
type CreateInstancePayload struct {
	// REQUIRED
	Acl *ACL `json:"acl"`
	// REQUIRED
	BackupSchedule *string `json:"backupSchedule"`
	// REQUIRED
	FlavorId *string `json:"flavorId"`
	// Labels field is not certain/clear
	Labels *map[string]string `json:"labels,omitempty"`
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Options *map[string]string `json:"options"`
	// REQUIRED
	Replicas *int64 `json:"replicas"`
	// REQUIRED
	Storage *Storage `json:"storage"`
	// REQUIRED
	Version *string `json:"version"`
}
