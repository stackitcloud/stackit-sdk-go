/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

// Machine struct for Machine
type Machine struct {
	// REQUIRED
	Image *Image `json:"image"`
	// For valid types please take a look at [provider-options](#tag/ProviderOptions/operation/SkeService_GetProviderOptions) `machineTypes`.
	// REQUIRED
	Type *string `json:"type"`
}
