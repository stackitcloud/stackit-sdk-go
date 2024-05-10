/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

type V1UpdateNetworkIPv4 struct {
	// List of DNS Servers/Nameservers.
	Nameservers *[]string `json:"nameservers,omitempty"`
}
