/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

// ListStoragesResponse struct for ListStoragesResponse
type ListStoragesResponse struct {
	StorageClasses *[]string     `json:"storageClasses,omitempty"`
	StorageRange   *StorageRange `json:"storageRange,omitempty"`
}
