/*
STACKIT Run Commands Service API

API endpoints for the STACKIT Run Commands Service API

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package runcommand

// Commands struct for Commands
type Commands struct {
	CommandTemplateName  *string `json:"commandTemplateName,omitempty"`
	CommandTemplateTitle *string `json:"commandTemplateTitle,omitempty"`
	FinishedAt           *string `json:"finishedAt,omitempty"`
	// Can be cast to int32 without loss of precision.
	Id        *int64  `json:"id,omitempty"`
	StartedAt *string `json:"startedAt,omitempty"`
	Status    *string `json:"status,omitempty"`
}
