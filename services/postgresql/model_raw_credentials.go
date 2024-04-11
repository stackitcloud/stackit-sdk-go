/*
STACKIT PostgreSQL API

The STACKIT PostgreSQL API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresql

type RawCredentials struct {
	// REQUIRED
	Credentials *PostgresQLCredentials `json:"credentials"`
}
