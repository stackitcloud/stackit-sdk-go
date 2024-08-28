/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

// SuggestedIndex struct for SuggestedIndex
type SuggestedIndex struct {
	// Unique id for this suggested index.
	Id *string `json:"id,omitempty"`
	// List of unique identifiers which correspond the query shapes in this response which pertain to this suggested index.
	Impact *[]string `json:"impact,omitempty"`
	// Array of documents that specifies a key in the index and its sort order, ascending or descending.
	Index *[]map[string]int32 `json:"index,omitempty"`
	// Namespace of the suggested index.
	Namespace *string `json:"namespace,omitempty"`
	// Estimated percentage performance improvement that the suggested index would provide.
	Weight *float64 `json:"weight,omitempty"`
}
