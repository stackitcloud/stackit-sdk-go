/*
STACKIT Redis API

The STACKIT Redis API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redis

// InstanceParameters struct for InstanceParameters
type InstanceParameters struct {
	// The unit is milliseconds.
	DownAfterMilliseconds *int64 `json:"down-after-milliseconds,omitempty"`
	EnableMonitoring      *bool  `json:"enable_monitoring,omitempty"`
	// The unit is milliseconds.
	FailoverTimeout *int64 `json:"failover-timeout,omitempty"`
	// If you want to monitor your service with Graphite, you can set the custom parameter graphite. It expects the host and port where the Graphite metrics should be sent to.
	Graphite             *string `json:"graphite,omitempty"`
	LazyfreeLazyEviction *string `json:"lazyfree-lazy-eviction,omitempty"`
	LazyfreeLazyExpire   *string `json:"lazyfree-lazy-expire,omitempty"`
	LuaTimeLimit         *int64  `json:"lua-time-limit,omitempty"`
	// This component monitors ephemeral and persistent disk usage. If one of these disk usages reaches the default configured threshold of 80%, the a9s Parachute stops all processes on that node.
	MaxDiskThreshold *int64  `json:"max_disk_threshold,omitempty"`
	Maxclients       *int64  `json:"maxclients,omitempty"`
	MaxmemoryPolicy  *string `json:"maxmemory-policy,omitempty"`
	MaxmemorySamples *int64  `json:"maxmemory-samples,omitempty"`
	// Frequency of metrics being emitted in seconds
	MetricsFrequency *int64 `json:"metrics_frequency,omitempty"`
	// Depending on your graphite provider, you might need to prefix the metrics with a certain value, like an API key for example.
	MetricsPrefix *string `json:"metrics_prefix,omitempty"`
	// The unit is seconds.
	MinReplicasMaxLag    *int64  `json:"min_replicas_max_lag,omitempty"`
	MonitoringInstanceId *string `json:"monitoring_instance_id,omitempty"`
	// The allowed value must include the following characters only: [K,E,g,$,l,s,h,z,x,e,A,t]
	NotifyKeyspaceEvents *string `json:"notify-keyspace-events,omitempty"`
	// Comma separated list of IP networks in CIDR notation which are allowed to access this instance.
	SgwAcl *string `json:"sgw_acl,omitempty"`
	// This setting must follow the original Redis configuration for RDB.
	Snapshot        *string   `json:"snapshot,omitempty"`
	Syslog          *[]string `json:"syslog,omitempty"`
	TlsCiphers      *[]string `json:"tls-ciphers,omitempty"`
	TlsCiphersuites *string   `json:"tls-ciphersuites,omitempty"`
	TlsProtocols    *string   `json:"tls-protocols,omitempty"`
}
