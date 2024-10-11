/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

// CreateServerGroupPayload Definition of a server group.
type CreateServerGroupPayload struct {
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// The servers that are part of the server group.
	Members *[]string `json:"members,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	// REQUIRED
	Name *string `json:"name"`
	// The server group policy.
	// REQUIRED
	Policy *string `json:"policy"`
}
