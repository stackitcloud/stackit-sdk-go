/*
STACKIT MariaDB API

The STACKIT MariaDB API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mariadb

type LogMeCredentials struct {
	// REQUIRED
	Host *string `json:"host"`
	// REQUIRED
	Password       *string `json:"password"`
	Port           *int64  `json:"port,omitempty"`
	SyslogDrainUrl *string `json:"syslog_drain_url,omitempty"`
	Uri            *string `json:"uri,omitempty"`
	// REQUIRED
	Username *string `json:"username"`
}