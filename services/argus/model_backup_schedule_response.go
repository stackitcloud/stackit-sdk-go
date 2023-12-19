/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type BackupScheduleResponse struct {
	AlertConfigBackupSchedules *[]BackupSchedule `json:"alertConfigBackupSchedules,omitempty"`
	AlertRulesBackupSchedules  *[]BackupSchedule `json:"alertRulesBackupSchedules,omitempty"`
	GrafanaBackupSchedules     *[]BackupSchedule `json:"grafanaBackupSchedules,omitempty"`
	// REQUIRED
	Message                     *string           `json:"message"`
	ScrapeConfigBackupSchedules *[]BackupSchedule `json:"scrapeConfigBackupSchedules,omitempty"`
}
