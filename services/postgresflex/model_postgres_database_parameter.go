/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

// PostgresDatabaseParameter struct for PostgresDatabaseParameter
type PostgresDatabaseParameter struct {
	// Context of the parameter.
	Context *string `json:"context,omitempty"`
	// Datatype describes the type of data that is used in the Value field.
	DataType *string `json:"dataType,omitempty"`
	// DefaultValue for the value field.
	DefaultValue *string `json:"defaultValue,omitempty"`
	// Description of the parameter.
	Description *string `json:"description,omitempty"`
	// Edit shows if the user can change this value.
	Edit *bool `json:"edit,omitempty"`
	// MaxValue describes the highest possible value that can be set.
	MaxValue *string `json:"maxValue,omitempty"`
	// MinValue describes the lowest possible value that can be set.
	MinValue *string `json:"minValue,omitempty"`
	// Name of the parameter.
	Name *string `json:"name,omitempty"`
	// PendingRestart describes if a parameter change requires a restart of the server.
	PendingRestart *bool `json:"pendingRestart,omitempty"`
	// ResetValue for the value field af.ter a reset.
	ResetValue *string `json:"resetValue,omitempty"`
	// Unit if the parameter has a unit if not empty.
	Unit *string `json:"unit,omitempty"`
	// Value of this parameter.
	Value *string `json:"value,omitempty"`
}
