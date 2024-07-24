/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

type Area struct {
	// A list containing DNS Servers/Nameservers for IPv4.
	DefaultNameservers *[]string `json:"defaultNameservers,omitempty"`
	// A list of network ranges.
	NetworkRanges *[]NetworkRange `json:"networkRanges,omitempty"`
	// A list of routes.
	Routes *[]Route `json:"routes,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	TransferNetwork *string `json:"transferNetwork,omitempty"`
}
