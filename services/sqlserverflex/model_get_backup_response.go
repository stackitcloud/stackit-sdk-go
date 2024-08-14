/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

// GetBackupResponse struct for GetBackupResponse
type GetBackupResponse struct {
	// Backup end time in UTC
	EndTime *string `json:"endTime,omitempty"`
	// Backup error
	Error *string `json:"error,omitempty"`
	// Backup id
	Id *string `json:"id,omitempty"`
	// Backup labels
	Labels *[]string `json:"labels,omitempty"`
	// Backup name
	Name *string `json:"name,omitempty"`
	// Backup specific options
	Options *map[string]string `json:"options,omitempty"`
	// Backup size in byte
	Size *int64 `json:"size,omitempty"`
	// Backup start time in UTC
	StartTime *string `json:"startTime,omitempty"`
}
