/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

// EmailConfig struct for EmailConfig
type EmailConfig struct {
	AuthIdentity *string `json:"authIdentity,omitempty"`
	AuthPassword *string `json:"authPassword,omitempty"`
	AuthUsername *string `json:"authUsername,omitempty"`
	From         *string `json:"from,omitempty"`
	SendResolved *bool   `json:"sendResolved,omitempty"`
	Smarthost    *string `json:"smarthost,omitempty"`
	// REQUIRED
	To *string `json:"to"`
}
