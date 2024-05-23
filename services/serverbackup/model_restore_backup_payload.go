/*
STACKIT Server Backup Management API

API endpoints for Server Backup Operations on STACKIT Servers.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverbackup

type RestoreBackupPayload struct {
	// REQUIRED
	StartServerAfterRestore *bool     `json:"startServerAfterRestore"`
	VolumeIds               *[]string `json:"volumeIds,omitempty"`
}
