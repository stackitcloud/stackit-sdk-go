/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

// ExportRecordSetsPayload struct for ExportRecordSetsPayload
type ExportRecordSetsPayload struct {
	ExportAsFQDN *bool   `json:"exportAsFQDN,omitempty"`
	Format       *string `json:"format,omitempty"`
}
