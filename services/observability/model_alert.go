/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

// Alert struct for Alert
type Alert struct {
	Global       *Global         `json:"global,omitempty"`
	InhibitRules *[]InhibitRules `json:"inhibitRules,omitempty"`
	// REQUIRED
	Receivers *[]Receivers `json:"receivers"`
	// REQUIRED
	Route *Route `json:"route"`
}