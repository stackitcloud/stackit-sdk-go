/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

type DetailedError struct {
	// REQUIRED
	Key *string `json:"key"`
	// REQUIRED
	Msg *string `json:"msg"`
}
