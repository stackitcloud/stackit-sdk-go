/*
STACKIT LogMe API

The STACKIT LogMe API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logme

type LogMe2InstanceParameters struct {
	EnableMonitoring     *bool   `json:"enable_monitoring,omitempty"`
	FluentdTcp           *int64  `json:"fluentd-tcp,omitempty"`
	FluentdTls           *int64  `json:"fluentd-tls,omitempty"`
	FluentdTlsCiphers    *string `json:"fluentd-tls-ciphers,omitempty"`
	FluentdTlsMaxVersion *string `json:"fluentd-tls-max-version,omitempty"`
	FluentdTlsMinVersion *string `json:"fluentd-tls-min-version,omitempty"`
	FluentdTlsVersion    *string `json:"fluentd-tls-version,omitempty"`
	FluentdUdp           *int64  `json:"fluentd-udp,omitempty"`
	// If you want to monitor your service with Graphite, you can set the custom parameter graphite. It expects the host and port where the Graphite metrics should be sent to.
	Graphite *string                   `json:"graphite,omitempty"`
	Groks    *[]map[string]interface{} `json:"groks,omitempty"`
	// Combination of an integer and a timerange when an index will be considered 'd' and can be deleted from OpenSearch. Possible values for the timerange are s, m, h and d.
	IsmDeletionAfter *string  `json:"ism_deletion_after,omitempty"`
	IsmJitter        *float64 `json:"ism_jitter,omitempty"`
	IsmJobInterval   *int64   `json:"ism_job_interval,omitempty"`
	// Default: not set, 46% of available memory will be used. The amount of memory (in MB) allocated as heap by the JVM for OpenSearch.
	JavaHeapspace *int64 `json:"java_heapspace,omitempty"`
	// The amount of memory (in MB) used by the JVM to store metadata for OpenSearch.
	JavaMaxmetaspace *int64 `json:"java_maxmetaspace,omitempty"`
	// This component monitors ephemeral and persistent disk usage. If one of these disk usages reaches the default configured threshold of 80%, the a9s Parachute stops all processes on that node.
	MaxDiskThreshold *int64 `json:"max_disk_threshold,omitempty"`
	// Frequency of metrics being emitted in seconds
	MetricsFrequency *int64 `json:"metrics_frequency,omitempty"`
	// Depending on your graphite provider, you might need to prefix the metrics with a certain value, like an API key for example.
	MetricsPrefix          *string   `json:"metrics_prefix,omitempty"`
	MonitoringInstanceId   *string   `json:"monitoring_instance_id,omitempty"`
	OpensearchTlsCiphers   *[]string `json:"opensearch-tls-ciphers,omitempty"`
	OpensearchTlsProtocols *[]string `json:"opensearch-tls-protocols,omitempty"`
	// Comma separated list of IP networks in CIDR notation which are allowed to access this instance.
	SgwAcl       *string   `json:"sgw_acl,omitempty"`
	Syslog       *[]string `json:"syslog,omitempty"`
	SyslogUseUdp *string   `json:"syslog-use-udp,omitempty"`
}
