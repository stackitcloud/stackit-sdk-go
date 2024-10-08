/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

// NetworkAreaListResponse Network area list response.
type NetworkAreaListResponse struct {
	// A list of network areas.
	// REQUIRED
	Items *[]NetworkArea `json:"items"`
}
