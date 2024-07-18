/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

type SingleDatabase struct {
	Id      *string                `json:"id,omitempty"`
	Name    *string                `json:"name,omitempty"`
	Options *SingleDatabaseOptions `json:"options,omitempty"`
}
