/*
STACKIT RabbitMQ API

The STACKIT RabbitMQ API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rabbitmq

type MariaDBInstanceParameters struct {
	EnableMonitoring *bool `json:"enable_monitoring,omitempty"`
	// If you want to monitor your service with Graphite, you can set the custom parameter graphite. It expects the host and port where the Graphite metrics should be sent to.
	Graphite *string `json:"graphite,omitempty"`
	// This component monitors ephemeral and persistent disk usage. If one of these disk usages reaches the default configured threshold of 80%, the a9s Parachute stops all processes on that node.
	MaxDiskThreshold *int64 `json:"max_disk_threshold,omitempty"`
	// Frequency of metrics being emitted in seconds
	MetricsFrequency *int64 `json:"metrics_frequency,omitempty"`
	// Depending on your graphite provider, you might need to prefix the metrics with a certain value, like an API key for example.
	MetricsPrefix        *string `json:"metrics_prefix,omitempty"`
	MonitoringInstanceId *string `json:"monitoring_instance_id,omitempty"`
	// Comma separated list of IP networks in CIDR notation which are allowed to access this instance.
	SgwAcl *string   `json:"sgw_acl,omitempty"`
	Syslog *[]string `json:"syslog,omitempty"`
}
