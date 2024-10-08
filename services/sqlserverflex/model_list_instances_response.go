/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

// ListInstancesResponse struct for ListInstancesResponse
type ListInstancesResponse struct {
	Count *int64                  `json:"count,omitempty"`
	Items *[]InstanceListInstance `json:"items,omitempty"`
}
