/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorization

type DomainFilter struct {
	Domains *[]string `json:"domains,omitempty"`
	// REQUIRED
	Effective *bool `json:"effective"`
	// REQUIRED
	ResourceId *string `json:"resourceId"`
	// REQUIRED
	ResourceType *string `json:"resourceType"`
	UpdatedAt    *string `json:"updatedAt,omitempty"`
}
