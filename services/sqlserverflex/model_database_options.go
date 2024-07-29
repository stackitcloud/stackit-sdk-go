/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

type DatabaseOptions struct {
	// Name of the collation of the database
	CollationName *string `json:"collationName,omitempty"`
	// CompatibilityLevel of the Database.
	CompatibilityLevel *int64 `json:"compatibilityLevel,omitempty"`
	// Name of the owner of the database.
	Owner *string `json:"owner,omitempty"`
}
