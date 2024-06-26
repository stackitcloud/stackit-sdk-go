/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

type CreateNetworkPayload struct {
	AddressFamily *CreateNetworkAddressFamily `json:"addressFamily,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	// REQUIRED
	Name *string `json:"name"`
}
