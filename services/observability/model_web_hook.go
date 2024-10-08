/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

// WebHook struct for WebHook
type WebHook struct {
	MsTeams      *bool `json:"msTeams,omitempty"`
	SendResolved *bool `json:"sendResolved,omitempty"`
	// REQUIRED
	Url *string `json:"url"`
}
