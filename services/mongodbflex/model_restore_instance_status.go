/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

// RestoreInstanceStatus struct for RestoreInstanceStatus
type RestoreInstanceStatus struct {
	BackupID   *string `json:"backupID,omitempty"`
	Date       *string `json:"date,omitempty"`
	Id         *string `json:"id,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	Status     *string `json:"status,omitempty"`
}
