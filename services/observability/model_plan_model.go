/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

// PlanModel struct for PlanModel
type PlanModel struct {
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
	Id *string `json:"id"`
	// REQUIRED
	LogsAlert *int64 `json:"logsAlert"`
	// REQUIRED
	LogsStorage *int64  `json:"logsStorage"`
	Name        *string `json:"name,omitempty"`
	// REQUIRED
	PlanId *string `json:"planId"`
	// REQUIRED
	SamplesPerScrape *int64 `json:"samplesPerScrape"`
	// REQUIRED
	TargetNumber *int64 `json:"targetNumber"`
	// REQUIRED
	TotalMetricSamples *int64 `json:"totalMetricSamples"`
	// REQUIRED
	TracesStorage *int64 `json:"tracesStorage"`
}
