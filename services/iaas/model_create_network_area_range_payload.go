/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

type CreateNetworkAreaRangePayload struct {
	// A list of Network ranges.
	Ipv4 *[]NetworkRange `json:"ipv4,omitempty"`
}
