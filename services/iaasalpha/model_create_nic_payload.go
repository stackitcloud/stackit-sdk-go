/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

// CreateNICPayload Object that represents a network interface.
type CreateNICPayload struct {
	// A list of IPs or CIDR notations.
	AllowedAddresses *[]AllowedAddressesInner `json:"allowedAddresses,omitempty"`
	// Universally Unique Identifier (UUID).
	// REQUIRED
	Device *string `json:"device"`
	// Universally Unique Identifier (UUID).
	// REQUIRED
	Id *string `json:"id"`
	// Object that represents an IP address.
	Ipv4 *string `json:"ipv4,omitempty"`
	// Object that represents an IPv6 address.
	Ipv6 *string `json:"ipv6,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// Object that represents an MAC address.
	// REQUIRED
	Mac *string `json:"mac"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name *string `json:"name,omitempty"`
	// Universally Unique Identifier (UUID).
	// REQUIRED
	NetworkId *string `json:"networkId"`
	// If this is set to false, then no security groups will apply to this network interface.
	NicSecurity *bool `json:"nicSecurity,omitempty"`
	// A list of UUIDs.
	SecurityGroups *[]string `json:"securityGroups,omitempty"`
	// REQUIRED
	Status *string `json:"status"`
	// REQUIRED
	Type *string `json:"type"`
}
