/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

// Project Object that represents a STACKIT project.
type Project struct {
	// REQUIRED
	AreaId         *AreaId `json:"areaId"`
	InternetAccess *bool   `json:"internetAccess,omitempty"`
	// The identifier (ID) of an OpenStack project.
	OpenstackProjectId *string `json:"openstackProjectId,omitempty"`
	// Universally Unique Identifier (UUID).
	// REQUIRED
	ProjectId *string `json:"projectId"`
	// The state of a resource object.
	// REQUIRED
	State *string `json:"state"`
}
