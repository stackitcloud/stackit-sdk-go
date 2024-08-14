/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorization

// ListPermissionsResponse struct for ListPermissionsResponse
type ListPermissionsResponse struct {
	// REQUIRED
	Permissions *[]Permission `json:"permissions"`
}
