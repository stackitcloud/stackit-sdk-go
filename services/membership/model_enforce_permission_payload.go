/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package membership

// EnforcePermissionPayload struct for EnforcePermissionPayload
type EnforcePermissionPayload struct {
	// REQUIRED
	Actions     *[]string    `json:"actions"`
	Consistency *Consistency `json:"consistency,omitempty"`
	// REQUIRED
	Resource *string `json:"resource"`
	// REQUIRED
	ResourceType *string `json:"resourceType"`
	// REQUIRED
	Subject *string `json:"subject"`
}
