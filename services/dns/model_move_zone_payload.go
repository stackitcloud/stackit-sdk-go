/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

// MoveZonePayload body to move zone from one project to another.
type MoveZonePayload struct {
	// Code to move the zone. It must be valid, not expired and belong
	// REQUIRED
	Code *string `json:"code"`
	// ZoneDnsName is the dns name of the zone to move
	// REQUIRED
	ZoneDnsName *string `json:"zoneDnsName"`
}
