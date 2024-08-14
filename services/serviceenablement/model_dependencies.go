/*
STACKIT Service Enablement API

STACKIT Service Enablement API

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceenablement

// Dependencies struct for Dependencies
type Dependencies struct {
	// a list of service IDs which this service depend on. If the service is enabled, those service are enabled as well automatically.
	Hard *[]string `json:"hard,omitempty"`
	// a list of service IDs which this service depend on. When they are disabled a notification is sent.
	Soft *[]string `json:"soft,omitempty"`
}
