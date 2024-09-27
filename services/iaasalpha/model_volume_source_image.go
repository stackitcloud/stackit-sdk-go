/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

// VolumeSourceImage The volume source type to generate a new volume based on an image.
type VolumeSourceImage struct {
	// Universally Unique Identifier (UUID).
	// REQUIRED
	ImageId *string `json:"imageId"`
}