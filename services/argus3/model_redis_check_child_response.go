/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus3

type RedisCheckChildResponse struct {
	// REQUIRED
	Id       *string `json:"id"`
	Password *string `json:"password,omitempty"`
	// REQUIRED
	Server   *string `json:"server"`
	Username *string `json:"username,omitempty"`
}