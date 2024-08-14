/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

// RestoreRunningRestore struct for RestoreRunningRestore
type RestoreRunningRestore struct {
	Command                 *string `json:"command,omitempty"`
	DatabaseName            *string `json:"database_name,omitempty"`
	EstimatedCompletionTime *string `json:"estimated_completion_time,omitempty"`
	PercentComplete         *int64  `json:"percent_complete,omitempty"`
	StartTime               *string `json:"start_time,omitempty"`
}
