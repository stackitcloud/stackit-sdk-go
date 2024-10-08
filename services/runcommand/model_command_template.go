/*
STACKIT Run Commands Service API

API endpoints for the STACKIT Run Commands Service API

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package runcommand

// CommandTemplate struct for CommandTemplate
type CommandTemplate struct {
	Name   *string   `json:"name,omitempty"`
	OsType *[]string `json:"osType,omitempty"`
	Title  *string   `json:"title,omitempty"`
}
