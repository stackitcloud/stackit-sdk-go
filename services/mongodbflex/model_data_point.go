/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

// DataPoint struct for DataPoint
type DataPoint struct {
	Timestamp *string  `json:"timestamp,omitempty"`
	Value     *float64 `json:"value,omitempty"`
}
