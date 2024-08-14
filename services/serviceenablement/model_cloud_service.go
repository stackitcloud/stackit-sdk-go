/*
STACKIT Service Enablement API

STACKIT Service Enablement API

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceenablement

// CloudService struct for CloudService
type CloudService struct {
	Dependencies *Dependencies      `json:"dependencies,omitempty"`
	Labels       *map[string]string `json:"labels,omitempty"`
	Scope        *string            `json:"scope,omitempty"`
	// the id of the service
	ServiceId *string `json:"serviceId,omitempty"`
}
