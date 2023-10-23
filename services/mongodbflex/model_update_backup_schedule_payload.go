/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

type UpdateBackupSchedulePayload struct {
	BackupSchedule                 *string `json:"backupSchedule,omitempty"`
	DailySnapshotRetentionDays     *int64  `json:"dailySnapshotRetentionDays,omitempty"`
	MonthlySnapshotRetentionMonths *int64  `json:"monthlySnapshotRetentionMonths,omitempty"`
	PointInTimeWindowHours         *int64  `json:"pointInTimeWindowHours,omitempty"`
	SnapshotRetentionDays          *int64  `json:"snapshotRetentionDays,omitempty"`
	WeeklySnapshotRetentionWeeks   *int64  `json:"weeklySnapshotRetentionWeeks,omitempty"`
}
