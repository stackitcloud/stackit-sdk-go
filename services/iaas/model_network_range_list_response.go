/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

// NetworkRangeListResponse Network Range list response.
type NetworkRangeListResponse struct {
	// A list of network ranges.
	// REQUIRED
	Items *[]NetworkRange `json:"items"`
}
