/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

// SlowQuery struct for SlowQuery
type SlowQuery struct {
	// The raw log line pertaining to the slow query.
	Line *string `json:"line,omitempty"`
	// The namespace in which the slow query ran.
	Namespace *string `json:"namespace,omitempty"`
}