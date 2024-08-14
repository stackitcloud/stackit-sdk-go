/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package membership

// CreateResourcePayload struct for CreateResourcePayload
type CreateResourcePayload struct {
	Members *[]Member `json:"members,omitempty"`
	// REQUIRED
	ParentId *string `json:"parentId"`
	// REQUIRED
	ParentType    *string `json:"parentType"`
	ResourceAlias *string `json:"resourceAlias,omitempty"`
	// REQUIRED
	ResourceId *string `json:"resourceId"`
}
