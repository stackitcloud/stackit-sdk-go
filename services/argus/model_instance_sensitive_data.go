/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type InstanceSensitiveData struct {
	// REQUIRED
	AlertingUrl *string `json:"alertingUrl"`
	// REQUIRED
	Cluster *string `json:"cluster"`
	// REQUIRED
	DashboardUrl *string `json:"dashboardUrl"`
	// REQUIRED
	GrafanaAdminPassword *string `json:"grafanaAdminPassword"`
	// REQUIRED
	GrafanaAdminUser *string `json:"grafanaAdminUser"`
	// REQUIRED
	GrafanaPublicReadAccess *bool `json:"grafanaPublicReadAccess"`
	// REQUIRED
	GrafanaUrl *string `json:"grafanaUrl"`
	// REQUIRED
	GrafanaUseStackitSso *bool `json:"grafanaUseStackitSso"`
	// REQUIRED
	Instance *string `json:"instance"`
	// REQUIRED
	JaegerTracesUrl *string `json:"jaegerTracesUrl"`
	// REQUIRED
	JaegerUiUrl *string `json:"jaegerUiUrl"`
	// REQUIRED
	LogsPushUrl *string `json:"logsPushUrl"`
	// REQUIRED
	LogsUrl *string `json:"logsUrl"`
	// REQUIRED
	MetricsRetentionTime1h *int32 `json:"metricsRetentionTime1h"`
	// REQUIRED
	MetricsRetentionTime5m *int32 `json:"metricsRetentionTime5m"`
	// REQUIRED
	MetricsRetentionTimeRaw *int32 `json:"metricsRetentionTimeRaw"`
	// REQUIRED
	MetricsUrl *string `json:"metricsUrl"`
	Name       *string `json:"name,omitempty"`
	// REQUIRED
	OtlpTracesUrl *string `json:"otlpTracesUrl"`
	// REQUIRED
	Plan *PlanModel `json:"plan"`
	// REQUIRED
	PushMetricsUrl *string `json:"pushMetricsUrl"`
	// REQUIRED
	TargetsUrl *string `json:"targetsUrl"`
	// REQUIRED
	ZipkinSpansUrl *string `json:"zipkinSpansUrl"`
}
