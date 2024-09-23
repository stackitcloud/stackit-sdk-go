/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

// HandlersInstancesSuggestedIndexesResponse struct for HandlersInstancesSuggestedIndexesResponse
type HandlersInstancesSuggestedIndexesResponse struct {
	// Documents with information about the query shapes that are served by the suggested indexes.
	Shapes *[]Shape `json:"shapes,omitempty"`
	// Documents with information about the indexes suggested by the Performance Advisor.
	SuggestedIndexes *[]SuggestedIndex `json:"suggestedIndexes,omitempty"`
}