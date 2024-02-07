/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorization

import (
	"time"
)

type ErrorResponse struct {
	// REQUIRED
	Error *string `json:"error"`
	// REQUIRED
	Message *string `json:"message"`
	// REQUIRED
	Path *string `json:"path"`
	// REQUIRED
	Status *int64 `json:"status"`
	// REQUIRED
	TimeStamp *time.Time `json:"timeStamp"`
}
