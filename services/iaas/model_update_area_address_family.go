/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

// UpdateAreaAddressFamily The addressFamily object for a area update request.
type UpdateAreaAddressFamily struct {
	Ipv4 *UpdateAreaIPv4 `json:"ipv4,omitempty"`
}
