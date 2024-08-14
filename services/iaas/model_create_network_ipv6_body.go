/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

// CreateNetworkIPv6Body The config object for a IPv6 network.
type CreateNetworkIPv6Body struct {
	// A list containing DNS Servers/Nameservers for IPv6.
	Nameservers  *[]string `json:"nameservers,omitempty"`
	PrefixLength *int64    `json:"prefixLength,omitempty"`
}
