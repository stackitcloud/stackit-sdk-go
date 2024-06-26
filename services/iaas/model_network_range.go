/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

type NetworkRange struct {
	// Universally Unique Identifier (UUID).
	NetworkRangeId *string `json:"networkRangeId,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	// REQUIRED
	Prefix *string `json:"prefix"`
}
