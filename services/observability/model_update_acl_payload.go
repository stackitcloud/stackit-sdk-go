/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

// UpdateACLPayload List of cidr. Send empty string to remove acl.
type UpdateACLPayload struct {
	// list of cidr
	// REQUIRED
	Acl *[]string `json:"acl"`
}
