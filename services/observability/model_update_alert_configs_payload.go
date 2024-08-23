/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

// UpdateAlertConfigsPayload Alert config
type UpdateAlertConfigsPayload struct {
	Global       *UpdateAlertConfigsPayloadGlobal       `json:"global,omitempty"`
	InhibitRules *UpdateAlertConfigsPayloadInhibitRules `json:"inhibitRules,omitempty"`
	// A list of notification receivers.
	// REQUIRED
	Receivers *[]UpdateAlertConfigsPayloadReceiversInner `json:"receivers"`
	// REQUIRED
	Route *UpdateAlertConfigsPayloadRoute `json:"route"`
}
