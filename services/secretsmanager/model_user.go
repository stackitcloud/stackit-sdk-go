/*
STACKIT Secrets Manager API

This API provides endpoints for managing the Secrets-Manager.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretsmanager

type User struct {
	// A user chosen description to differentiate between multiple users.
	// REQUIRED
	Description *string `json:"description"`
	// A auto generated unique id which identifies the users.
	// REQUIRED
	Id *string `json:"id"`
	// A auto generated password for logging in with the user.
	// REQUIRED
	Password *string `json:"password"`
	// A auto generated username for logging in with the user.
	// REQUIRED
	Username *string `json:"username"`
	// Is true if the user has write access to the secrets engine. Is false for a read-only user.
	// REQUIRED
	Write *bool `json:"write"`
}
