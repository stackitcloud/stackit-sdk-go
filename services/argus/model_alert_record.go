/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type AlertRecord struct {
	// REQUIRED
	Expr   *string            `json:"expr"`
	Labels *map[string]string `json:"labels,omitempty"`
	// REQUIRED
	Record *string `json:"record"`
}