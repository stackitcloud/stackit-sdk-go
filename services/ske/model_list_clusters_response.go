/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

// ListClustersResponse struct for ListClustersResponse
type ListClustersResponse struct {
	Items *[]Cluster `json:"items,omitempty"`
}
