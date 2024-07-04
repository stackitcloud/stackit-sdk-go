/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

type UpdateInstancePayload struct {
	Acl            *ACL    `json:"acl,omitempty"`
	BackupSchedule *string `json:"backupSchedule,omitempty"`
	FlavorId       *string `json:"flavorId,omitempty"`
	// Labels field is not certain/clear
	Labels   *map[string]string `json:"labels,omitempty"`
	Name     *string            `json:"name,omitempty"`
	Options  *map[string]string `json:"options,omitempty"`
	Replicas *int64             `json:"replicas,omitempty"`
	Storage  *Storage           `json:"storage,omitempty"`
	Version  *string            `json:"version,omitempty"`
}