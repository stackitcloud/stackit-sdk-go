/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

// InstanceDatabase struct for InstanceDatabase
type InstanceDatabase struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Database specific options
	Options *map[string]interface{} `json:"options,omitempty"`
}
