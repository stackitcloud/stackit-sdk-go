/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

// CreateRecordSetPayload RRSetPost for rr set info.
type CreateRecordSetPayload struct {
	// user comment
	Comment *string `json:"comment,omitempty"`
	// name of the record which should be a valid domain according to rfc1035 Section 2.3.4
	// REQUIRED
	Name *string `json:"name"`
	// records
	// REQUIRED
	Records *[]RecordPayload `json:"records"`
	// time to live. If nothing provided we will set the zone ttl.
	Ttl *int64 `json:"ttl,omitempty"`
	// record set type
	// REQUIRED
	Type *string `json:"type"`
}
