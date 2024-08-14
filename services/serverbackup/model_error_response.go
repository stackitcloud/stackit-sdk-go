/*
STACKIT Server Backup Management API

API endpoints for Server Backup Operations on STACKIT Servers.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverbackup

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	// Details about the error
	// REQUIRED
	Message *string `json:"message"`
	// The string representation of the http status code (i.e. Not Found, Bad Request, etc)
	// REQUIRED
	Status *string `json:"status"`
	// The time the error occured
	// REQUIRED
	Timestamp *string `json:"timestamp"`
}
