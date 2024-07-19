/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

type CreateAreaIPv4 struct {
	DefaultNameservers *[]string `json:"defaultNameservers,omitempty"`
	// A list of network ranges.
	// REQUIRED
	NetworkRanges *[]NetworkRange `json:"networkRanges"`
	// A list of routes.
	Routes *[]Route `json:"routes,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	// REQUIRED
	TransferNetwork *string `json:"transferNetwork"`
	// The default prefix length for networks in the network area.
	DefaultPrefixLen *int64 `json:"defaultPrefixLen,omitempty"`
	// The maximal prefix length for networks in the network area.
	MaxPrefixLen *int64 `json:"maxPrefixLen,omitempty"`
	// The minimal prefix length for networks in the network area.
	MinPrefixLen *int64 `json:"minPrefixLen,omitempty"`
}
