/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

// Receiver struct for Receiver
type Receiver struct {
	// REQUIRED
	Data *Receivers `json:"data"`
	// REQUIRED
	Message *string `json:"message"`
}
