/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"time"
)

// Network Object that represents a network.
type Network struct {
	// Date-time when resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// REQUIRED
	Name *string `json:"name"`
	// A list containing DNS Servers/Nameservers for IPv4.
	Nameservers *[]string `json:"nameservers,omitempty"`
	// A list containing DNS Servers/Nameservers for IPv6.
	NameserversV6 *[]string `json:"nameserversV6,omitempty"`
	// Universally Unique Identifier (UUID).
	// REQUIRED
	NetworkId  *string   `json:"networkId"`
	Prefixes   *[]string `json:"prefixes,omitempty"`
	PrefixesV6 *[]string `json:"prefixesV6,omitempty"`
	// Object that represents an IP address.
	PublicIp *string `json:"publicIp,omitempty"`
	// Shows if the network is routed and therefore accessible from other networks.
	Routed *bool `json:"routed,omitempty"`
	// The state of a resource object.
	// REQUIRED
	State *string `json:"state"`
	// Date-time when resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
