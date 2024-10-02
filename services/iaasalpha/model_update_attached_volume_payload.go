/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

// UpdateAttachedVolumePayload Object that represents a Volume attachment to a server.
type UpdateAttachedVolumePayload struct {
	// Delete the volume during the termination of the server. Defaults to false.
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty"`
	// Universally Unique Identifier (UUID).
	ServerId *string `json:"serverId,omitempty"`
	// Universally Unique Identifier (UUID).
	VolumeId *string `json:"volumeId,omitempty"`
}