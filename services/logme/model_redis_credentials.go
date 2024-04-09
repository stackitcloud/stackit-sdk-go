/*
STACKIT LogMe API

The STACKIT LogMe API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logme

type RedisCredentials struct {
	// REQUIRED
	Host             *string   `json:"host"`
	Hosts            *[]string `json:"hosts,omitempty"`
	LoadBalancedHost *string   `json:"load_balanced_host,omitempty"`
	// REQUIRED
	Password *string `json:"password"`
	Port     *int64  `json:"port,omitempty"`
	Uri      *string `json:"uri,omitempty"`
	// REQUIRED
	Username *string `json:"username"`
}
