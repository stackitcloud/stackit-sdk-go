/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type Plan struct {
	// REQUIRED
	AlertMatchers *int64 `json:"alertMatchers"`
	// REQUIRED
	AlertReceivers *int64 `json:"alertReceivers"`
	// REQUIRED
	AlertRules *int64   `json:"alertRules"`
	Amount     *float64 `json:"amount,omitempty"`
	// REQUIRED
	BucketSize  *int64  `json:"bucketSize"`
	Description *string `json:"description,omitempty"`
	// REQUIRED
	GrafanaGlobalDashboards *int64 `json:"grafanaGlobalDashboards"`
	// REQUIRED
	GrafanaGlobalOrgs *int64 `json:"grafanaGlobalOrgs"`
	// REQUIRED
	GrafanaGlobalSessions *int64 `json:"grafanaGlobalSessions"`
	// REQUIRED
	GrafanaGlobalUsers *int64 `json:"grafanaGlobalUsers"`
	// REQUIRED
	Id       *string `json:"id"`
	IsFree   *bool   `json:"isFree,omitempty"`
	IsPublic *bool   `json:"isPublic,omitempty"`
	// REQUIRED
	LogsAlert *int64 `json:"logsAlert"`
	// REQUIRED
	LogsStorage *int64  `json:"logsStorage"`
	Name        *string `json:"name,omitempty"`
	// REQUIRED
	PlanId *string `json:"planId"`
	// REQUIRED
	SamplesPerScrape *int64  `json:"samplesPerScrape"`
	Schema           *string `json:"schema,omitempty"`
	// REQUIRED
	TargetNumber *int64 `json:"targetNumber"`
	// REQUIRED
	TotalMetricSamples *int64 `json:"totalMetricSamples"`
	// REQUIRED
	TracesStorage *int64 `json:"tracesStorage"`
}
