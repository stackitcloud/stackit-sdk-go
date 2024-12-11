/*
STACKIT Redis API

The STACKIT Redis API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redis

import (
	"encoding/json"
)

// checks if the InstanceParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceParameters{}

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

// NewInstanceParameters instantiates a new InstanceParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceParameters() *InstanceParameters {
	this := InstanceParameters{}
	var downAfterMilliseconds int64 = 10000
	this.DownAfterMilliseconds = &downAfterMilliseconds
	var enableMonitoring bool = false
	this.EnableMonitoring = &enableMonitoring
	var failoverTimeout int64 = 30000
	this.FailoverTimeout = &failoverTimeout
	var lazyfreeLazyEviction string = "no"
	this.LazyfreeLazyEviction = &lazyfreeLazyEviction
	var lazyfreeLazyExpire string = "no"
	this.LazyfreeLazyExpire = &lazyfreeLazyExpire
	var luaTimeLimit int64 = 5000
	this.LuaTimeLimit = &luaTimeLimit
	var maxDiskThreshold int64 = 80
	this.MaxDiskThreshold = &maxDiskThreshold
	var maxclients int64 = 10000
	this.Maxclients = &maxclients
	var maxmemoryPolicy string = "volatile-lru"
	this.MaxmemoryPolicy = &maxmemoryPolicy
	var maxmemorySamples int64 = 5
	this.MaxmemorySamples = &maxmemorySamples
	var metricsFrequency int64 = 10
	this.MetricsFrequency = &metricsFrequency
	var minReplicasMaxLag int64 = 10
	this.MinReplicasMaxLag = &minReplicasMaxLag
	var notifyKeyspaceEvents string = ""
	this.NotifyKeyspaceEvents = &notifyKeyspaceEvents
	return &this
}

// NewInstanceParametersWithDefaults instantiates a new InstanceParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceParametersWithDefaults() *InstanceParameters {
	this := InstanceParameters{}
	var downAfterMilliseconds int64 = 10000
	this.DownAfterMilliseconds = &downAfterMilliseconds
	var enableMonitoring bool = false
	this.EnableMonitoring = &enableMonitoring
	var failoverTimeout int64 = 30000
	this.FailoverTimeout = &failoverTimeout
	var lazyfreeLazyEviction string = "no"
	this.LazyfreeLazyEviction = &lazyfreeLazyEviction
	var lazyfreeLazyExpire string = "no"
	this.LazyfreeLazyExpire = &lazyfreeLazyExpire
	var luaTimeLimit int64 = 5000
	this.LuaTimeLimit = &luaTimeLimit
	var maxDiskThreshold int64 = 80
	this.MaxDiskThreshold = &maxDiskThreshold
	var maxclients int64 = 10000
	this.Maxclients = &maxclients
	var maxmemoryPolicy string = "volatile-lru"
	this.MaxmemoryPolicy = &maxmemoryPolicy
	var maxmemorySamples int64 = 5
	this.MaxmemorySamples = &maxmemorySamples
	var metricsFrequency int64 = 10
	this.MetricsFrequency = &metricsFrequency
	var minReplicasMaxLag int64 = 10
	this.MinReplicasMaxLag = &minReplicasMaxLag
	var notifyKeyspaceEvents string = ""
	this.NotifyKeyspaceEvents = &notifyKeyspaceEvents
	return &this
}

// GetDownAfterMilliseconds returns the DownAfterMilliseconds field value if set, zero value otherwise.
func (o *InstanceParameters) GetDownAfterMilliseconds() *int64 {
	if o == nil || IsNil(o.DownAfterMilliseconds) {
		var ret *int64
		return ret
	}
	return o.DownAfterMilliseconds
}

// GetDownAfterMillisecondsOk returns a tuple with the DownAfterMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetDownAfterMillisecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.DownAfterMilliseconds) {
		return nil, false
	}
	return o.DownAfterMilliseconds, true
}

// HasDownAfterMilliseconds returns a boolean if a field has been set.
func (o *InstanceParameters) HasDownAfterMilliseconds() bool {
	if o != nil && !IsNil(o.DownAfterMilliseconds) && !IsNil(o.DownAfterMilliseconds) {
		return true
	}

	return false
}

// SetDownAfterMilliseconds gets a reference to the given int64 and assigns it to the DownAfterMilliseconds field.
func (o *InstanceParameters) SetDownAfterMilliseconds(v *int64) {
	o.DownAfterMilliseconds = v
}

// GetEnableMonitoring returns the EnableMonitoring field value if set, zero value otherwise.
func (o *InstanceParameters) GetEnableMonitoring() *bool {
	if o == nil || IsNil(o.EnableMonitoring) {
		var ret *bool
		return ret
	}
	return o.EnableMonitoring
}

// GetEnableMonitoringOk returns a tuple with the EnableMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetEnableMonitoringOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableMonitoring) {
		return nil, false
	}
	return o.EnableMonitoring, true
}

// HasEnableMonitoring returns a boolean if a field has been set.
func (o *InstanceParameters) HasEnableMonitoring() bool {
	if o != nil && !IsNil(o.EnableMonitoring) && !IsNil(o.EnableMonitoring) {
		return true
	}

	return false
}

// SetEnableMonitoring gets a reference to the given bool and assigns it to the EnableMonitoring field.
func (o *InstanceParameters) SetEnableMonitoring(v *bool) {
	o.EnableMonitoring = v
}

// GetFailoverTimeout returns the FailoverTimeout field value if set, zero value otherwise.
func (o *InstanceParameters) GetFailoverTimeout() *int64 {
	if o == nil || IsNil(o.FailoverTimeout) {
		var ret *int64
		return ret
	}
	return o.FailoverTimeout
}

// GetFailoverTimeoutOk returns a tuple with the FailoverTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetFailoverTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.FailoverTimeout) {
		return nil, false
	}
	return o.FailoverTimeout, true
}

// HasFailoverTimeout returns a boolean if a field has been set.
func (o *InstanceParameters) HasFailoverTimeout() bool {
	if o != nil && !IsNil(o.FailoverTimeout) && !IsNil(o.FailoverTimeout) {
		return true
	}

	return false
}

// SetFailoverTimeout gets a reference to the given int64 and assigns it to the FailoverTimeout field.
func (o *InstanceParameters) SetFailoverTimeout(v *int64) {
	o.FailoverTimeout = v
}

// GetGraphite returns the Graphite field value if set, zero value otherwise.
func (o *InstanceParameters) GetGraphite() *string {
	if o == nil || IsNil(o.Graphite) {
		var ret *string
		return ret
	}
	return o.Graphite
}

// GetGraphiteOk returns a tuple with the Graphite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetGraphiteOk() (*string, bool) {
	if o == nil || IsNil(o.Graphite) {
		return nil, false
	}
	return o.Graphite, true
}

// HasGraphite returns a boolean if a field has been set.
func (o *InstanceParameters) HasGraphite() bool {
	if o != nil && !IsNil(o.Graphite) && !IsNil(o.Graphite) {
		return true
	}

	return false
}

// SetGraphite gets a reference to the given string and assigns it to the Graphite field.
func (o *InstanceParameters) SetGraphite(v *string) {
	o.Graphite = v
}

// GetLazyfreeLazyEviction returns the LazyfreeLazyEviction field value if set, zero value otherwise.
func (o *InstanceParameters) GetLazyfreeLazyEviction() *string {
	if o == nil || IsNil(o.LazyfreeLazyEviction) {
		var ret *string
		return ret
	}
	return o.LazyfreeLazyEviction
}

// GetLazyfreeLazyEvictionOk returns a tuple with the LazyfreeLazyEviction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetLazyfreeLazyEvictionOk() (*string, bool) {
	if o == nil || IsNil(o.LazyfreeLazyEviction) {
		return nil, false
	}
	return o.LazyfreeLazyEviction, true
}

// HasLazyfreeLazyEviction returns a boolean if a field has been set.
func (o *InstanceParameters) HasLazyfreeLazyEviction() bool {
	if o != nil && !IsNil(o.LazyfreeLazyEviction) && !IsNil(o.LazyfreeLazyEviction) {
		return true
	}

	return false
}

// SetLazyfreeLazyEviction gets a reference to the given string and assigns it to the LazyfreeLazyEviction field.
func (o *InstanceParameters) SetLazyfreeLazyEviction(v *string) {
	o.LazyfreeLazyEviction = v
}

// GetLazyfreeLazyExpire returns the LazyfreeLazyExpire field value if set, zero value otherwise.
func (o *InstanceParameters) GetLazyfreeLazyExpire() *string {
	if o == nil || IsNil(o.LazyfreeLazyExpire) {
		var ret *string
		return ret
	}
	return o.LazyfreeLazyExpire
}

// GetLazyfreeLazyExpireOk returns a tuple with the LazyfreeLazyExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetLazyfreeLazyExpireOk() (*string, bool) {
	if o == nil || IsNil(o.LazyfreeLazyExpire) {
		return nil, false
	}
	return o.LazyfreeLazyExpire, true
}

// HasLazyfreeLazyExpire returns a boolean if a field has been set.
func (o *InstanceParameters) HasLazyfreeLazyExpire() bool {
	if o != nil && !IsNil(o.LazyfreeLazyExpire) && !IsNil(o.LazyfreeLazyExpire) {
		return true
	}

	return false
}

// SetLazyfreeLazyExpire gets a reference to the given string and assigns it to the LazyfreeLazyExpire field.
func (o *InstanceParameters) SetLazyfreeLazyExpire(v *string) {
	o.LazyfreeLazyExpire = v
}

// GetLuaTimeLimit returns the LuaTimeLimit field value if set, zero value otherwise.
func (o *InstanceParameters) GetLuaTimeLimit() *int64 {
	if o == nil || IsNil(o.LuaTimeLimit) {
		var ret *int64
		return ret
	}
	return o.LuaTimeLimit
}

// GetLuaTimeLimitOk returns a tuple with the LuaTimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetLuaTimeLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.LuaTimeLimit) {
		return nil, false
	}
	return o.LuaTimeLimit, true
}

// HasLuaTimeLimit returns a boolean if a field has been set.
func (o *InstanceParameters) HasLuaTimeLimit() bool {
	if o != nil && !IsNil(o.LuaTimeLimit) && !IsNil(o.LuaTimeLimit) {
		return true
	}

	return false
}

// SetLuaTimeLimit gets a reference to the given int64 and assigns it to the LuaTimeLimit field.
func (o *InstanceParameters) SetLuaTimeLimit(v *int64) {
	o.LuaTimeLimit = v
}

// GetMaxDiskThreshold returns the MaxDiskThreshold field value if set, zero value otherwise.
func (o *InstanceParameters) GetMaxDiskThreshold() *int64 {
	if o == nil || IsNil(o.MaxDiskThreshold) {
		var ret *int64
		return ret
	}
	return o.MaxDiskThreshold
}

// GetMaxDiskThresholdOk returns a tuple with the MaxDiskThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetMaxDiskThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxDiskThreshold) {
		return nil, false
	}
	return o.MaxDiskThreshold, true
}

// HasMaxDiskThreshold returns a boolean if a field has been set.
func (o *InstanceParameters) HasMaxDiskThreshold() bool {
	if o != nil && !IsNil(o.MaxDiskThreshold) && !IsNil(o.MaxDiskThreshold) {
		return true
	}

	return false
}

// SetMaxDiskThreshold gets a reference to the given int64 and assigns it to the MaxDiskThreshold field.
func (o *InstanceParameters) SetMaxDiskThreshold(v *int64) {
	o.MaxDiskThreshold = v
}

// GetMaxclients returns the Maxclients field value if set, zero value otherwise.
func (o *InstanceParameters) GetMaxclients() *int64 {
	if o == nil || IsNil(o.Maxclients) {
		var ret *int64
		return ret
	}
	return o.Maxclients
}

// GetMaxclientsOk returns a tuple with the Maxclients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetMaxclientsOk() (*int64, bool) {
	if o == nil || IsNil(o.Maxclients) {
		return nil, false
	}
	return o.Maxclients, true
}

// HasMaxclients returns a boolean if a field has been set.
func (o *InstanceParameters) HasMaxclients() bool {
	if o != nil && !IsNil(o.Maxclients) && !IsNil(o.Maxclients) {
		return true
	}

	return false
}

// SetMaxclients gets a reference to the given int64 and assigns it to the Maxclients field.
func (o *InstanceParameters) SetMaxclients(v *int64) {
	o.Maxclients = v
}

// GetMaxmemoryPolicy returns the MaxmemoryPolicy field value if set, zero value otherwise.
func (o *InstanceParameters) GetMaxmemoryPolicy() *string {
	if o == nil || IsNil(o.MaxmemoryPolicy) {
		var ret *string
		return ret
	}
	return o.MaxmemoryPolicy
}

// GetMaxmemoryPolicyOk returns a tuple with the MaxmemoryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetMaxmemoryPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.MaxmemoryPolicy) {
		return nil, false
	}
	return o.MaxmemoryPolicy, true
}

// HasMaxmemoryPolicy returns a boolean if a field has been set.
func (o *InstanceParameters) HasMaxmemoryPolicy() bool {
	if o != nil && !IsNil(o.MaxmemoryPolicy) && !IsNil(o.MaxmemoryPolicy) {
		return true
	}

	return false
}

// SetMaxmemoryPolicy gets a reference to the given string and assigns it to the MaxmemoryPolicy field.
func (o *InstanceParameters) SetMaxmemoryPolicy(v *string) {
	o.MaxmemoryPolicy = v
}

// GetMaxmemorySamples returns the MaxmemorySamples field value if set, zero value otherwise.
func (o *InstanceParameters) GetMaxmemorySamples() *int64 {
	if o == nil || IsNil(o.MaxmemorySamples) {
		var ret *int64
		return ret
	}
	return o.MaxmemorySamples
}

// GetMaxmemorySamplesOk returns a tuple with the MaxmemorySamples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetMaxmemorySamplesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxmemorySamples) {
		return nil, false
	}
	return o.MaxmemorySamples, true
}

// HasMaxmemorySamples returns a boolean if a field has been set.
func (o *InstanceParameters) HasMaxmemorySamples() bool {
	if o != nil && !IsNil(o.MaxmemorySamples) && !IsNil(o.MaxmemorySamples) {
		return true
	}

	return false
}

// SetMaxmemorySamples gets a reference to the given int64 and assigns it to the MaxmemorySamples field.
func (o *InstanceParameters) SetMaxmemorySamples(v *int64) {
	o.MaxmemorySamples = v
}

// GetMetricsFrequency returns the MetricsFrequency field value if set, zero value otherwise.
func (o *InstanceParameters) GetMetricsFrequency() *int64 {
	if o == nil || IsNil(o.MetricsFrequency) {
		var ret *int64
		return ret
	}
	return o.MetricsFrequency
}

// GetMetricsFrequencyOk returns a tuple with the MetricsFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetMetricsFrequencyOk() (*int64, bool) {
	if o == nil || IsNil(o.MetricsFrequency) {
		return nil, false
	}
	return o.MetricsFrequency, true
}

// HasMetricsFrequency returns a boolean if a field has been set.
func (o *InstanceParameters) HasMetricsFrequency() bool {
	if o != nil && !IsNil(o.MetricsFrequency) && !IsNil(o.MetricsFrequency) {
		return true
	}

	return false
}

// SetMetricsFrequency gets a reference to the given int64 and assigns it to the MetricsFrequency field.
func (o *InstanceParameters) SetMetricsFrequency(v *int64) {
	o.MetricsFrequency = v
}

// GetMetricsPrefix returns the MetricsPrefix field value if set, zero value otherwise.
func (o *InstanceParameters) GetMetricsPrefix() *string {
	if o == nil || IsNil(o.MetricsPrefix) {
		var ret *string
		return ret
	}
	return o.MetricsPrefix
}

// GetMetricsPrefixOk returns a tuple with the MetricsPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetMetricsPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.MetricsPrefix) {
		return nil, false
	}
	return o.MetricsPrefix, true
}

// HasMetricsPrefix returns a boolean if a field has been set.
func (o *InstanceParameters) HasMetricsPrefix() bool {
	if o != nil && !IsNil(o.MetricsPrefix) && !IsNil(o.MetricsPrefix) {
		return true
	}

	return false
}

// SetMetricsPrefix gets a reference to the given string and assigns it to the MetricsPrefix field.
func (o *InstanceParameters) SetMetricsPrefix(v *string) {
	o.MetricsPrefix = v
}

// GetMinReplicasMaxLag returns the MinReplicasMaxLag field value if set, zero value otherwise.
func (o *InstanceParameters) GetMinReplicasMaxLag() *int64 {
	if o == nil || IsNil(o.MinReplicasMaxLag) {
		var ret *int64
		return ret
	}
	return o.MinReplicasMaxLag
}

// GetMinReplicasMaxLagOk returns a tuple with the MinReplicasMaxLag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetMinReplicasMaxLagOk() (*int64, bool) {
	if o == nil || IsNil(o.MinReplicasMaxLag) {
		return nil, false
	}
	return o.MinReplicasMaxLag, true
}

// HasMinReplicasMaxLag returns a boolean if a field has been set.
func (o *InstanceParameters) HasMinReplicasMaxLag() bool {
	if o != nil && !IsNil(o.MinReplicasMaxLag) && !IsNil(o.MinReplicasMaxLag) {
		return true
	}

	return false
}

// SetMinReplicasMaxLag gets a reference to the given int64 and assigns it to the MinReplicasMaxLag field.
func (o *InstanceParameters) SetMinReplicasMaxLag(v *int64) {
	o.MinReplicasMaxLag = v
}

// GetMonitoringInstanceId returns the MonitoringInstanceId field value if set, zero value otherwise.
func (o *InstanceParameters) GetMonitoringInstanceId() *string {
	if o == nil || IsNil(o.MonitoringInstanceId) {
		var ret *string
		return ret
	}
	return o.MonitoringInstanceId
}

// GetMonitoringInstanceIdOk returns a tuple with the MonitoringInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetMonitoringInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MonitoringInstanceId) {
		return nil, false
	}
	return o.MonitoringInstanceId, true
}

// HasMonitoringInstanceId returns a boolean if a field has been set.
func (o *InstanceParameters) HasMonitoringInstanceId() bool {
	if o != nil && !IsNil(o.MonitoringInstanceId) && !IsNil(o.MonitoringInstanceId) {
		return true
	}

	return false
}

// SetMonitoringInstanceId gets a reference to the given string and assigns it to the MonitoringInstanceId field.
func (o *InstanceParameters) SetMonitoringInstanceId(v *string) {
	o.MonitoringInstanceId = v
}

// GetNotifyKeyspaceEvents returns the NotifyKeyspaceEvents field value if set, zero value otherwise.
func (o *InstanceParameters) GetNotifyKeyspaceEvents() *string {
	if o == nil || IsNil(o.NotifyKeyspaceEvents) {
		var ret *string
		return ret
	}
	return o.NotifyKeyspaceEvents
}

// GetNotifyKeyspaceEventsOk returns a tuple with the NotifyKeyspaceEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetNotifyKeyspaceEventsOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyKeyspaceEvents) {
		return nil, false
	}
	return o.NotifyKeyspaceEvents, true
}

// HasNotifyKeyspaceEvents returns a boolean if a field has been set.
func (o *InstanceParameters) HasNotifyKeyspaceEvents() bool {
	if o != nil && !IsNil(o.NotifyKeyspaceEvents) && !IsNil(o.NotifyKeyspaceEvents) {
		return true
	}

	return false
}

// SetNotifyKeyspaceEvents gets a reference to the given string and assigns it to the NotifyKeyspaceEvents field.
func (o *InstanceParameters) SetNotifyKeyspaceEvents(v *string) {
	o.NotifyKeyspaceEvents = v
}

// GetSgwAcl returns the SgwAcl field value if set, zero value otherwise.
func (o *InstanceParameters) GetSgwAcl() *string {
	if o == nil || IsNil(o.SgwAcl) {
		var ret *string
		return ret
	}
	return o.SgwAcl
}

// GetSgwAclOk returns a tuple with the SgwAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetSgwAclOk() (*string, bool) {
	if o == nil || IsNil(o.SgwAcl) {
		return nil, false
	}
	return o.SgwAcl, true
}

// HasSgwAcl returns a boolean if a field has been set.
func (o *InstanceParameters) HasSgwAcl() bool {
	if o != nil && !IsNil(o.SgwAcl) && !IsNil(o.SgwAcl) {
		return true
	}

	return false
}

// SetSgwAcl gets a reference to the given string and assigns it to the SgwAcl field.
func (o *InstanceParameters) SetSgwAcl(v *string) {
	o.SgwAcl = v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *InstanceParameters) GetSnapshot() *string {
	if o == nil || IsNil(o.Snapshot) {
		var ret *string
		return ret
	}
	return o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetSnapshotOk() (*string, bool) {
	if o == nil || IsNil(o.Snapshot) {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *InstanceParameters) HasSnapshot() bool {
	if o != nil && !IsNil(o.Snapshot) && !IsNil(o.Snapshot) {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given string and assigns it to the Snapshot field.
func (o *InstanceParameters) SetSnapshot(v *string) {
	o.Snapshot = v
}

// GetSyslog returns the Syslog field value if set, zero value otherwise.
func (o *InstanceParameters) GetSyslog() *[]string {
	if o == nil || IsNil(o.Syslog) {
		var ret *[]string
		return ret
	}
	return o.Syslog
}

// GetSyslogOk returns a tuple with the Syslog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetSyslogOk() (*[]string, bool) {
	if o == nil || IsNil(o.Syslog) {
		return nil, false
	}
	return o.Syslog, true
}

// HasSyslog returns a boolean if a field has been set.
func (o *InstanceParameters) HasSyslog() bool {
	if o != nil && !IsNil(o.Syslog) && !IsNil(o.Syslog) {
		return true
	}

	return false
}

// SetSyslog gets a reference to the given []string and assigns it to the Syslog field.
func (o *InstanceParameters) SetSyslog(v *[]string) {
	o.Syslog = v
}

// GetTlsCiphers returns the TlsCiphers field value if set, zero value otherwise.
func (o *InstanceParameters) GetTlsCiphers() *[]string {
	if o == nil || IsNil(o.TlsCiphers) {
		var ret *[]string
		return ret
	}
	return o.TlsCiphers
}

// GetTlsCiphersOk returns a tuple with the TlsCiphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetTlsCiphersOk() (*[]string, bool) {
	if o == nil || IsNil(o.TlsCiphers) {
		return nil, false
	}
	return o.TlsCiphers, true
}

// HasTlsCiphers returns a boolean if a field has been set.
func (o *InstanceParameters) HasTlsCiphers() bool {
	if o != nil && !IsNil(o.TlsCiphers) && !IsNil(o.TlsCiphers) {
		return true
	}

	return false
}

// SetTlsCiphers gets a reference to the given []string and assigns it to the TlsCiphers field.
func (o *InstanceParameters) SetTlsCiphers(v *[]string) {
	o.TlsCiphers = v
}

// GetTlsCiphersuites returns the TlsCiphersuites field value if set, zero value otherwise.
func (o *InstanceParameters) GetTlsCiphersuites() *string {
	if o == nil || IsNil(o.TlsCiphersuites) {
		var ret *string
		return ret
	}
	return o.TlsCiphersuites
}

// GetTlsCiphersuitesOk returns a tuple with the TlsCiphersuites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetTlsCiphersuitesOk() (*string, bool) {
	if o == nil || IsNil(o.TlsCiphersuites) {
		return nil, false
	}
	return o.TlsCiphersuites, true
}

// HasTlsCiphersuites returns a boolean if a field has been set.
func (o *InstanceParameters) HasTlsCiphersuites() bool {
	if o != nil && !IsNil(o.TlsCiphersuites) && !IsNil(o.TlsCiphersuites) {
		return true
	}

	return false
}

// SetTlsCiphersuites gets a reference to the given string and assigns it to the TlsCiphersuites field.
func (o *InstanceParameters) SetTlsCiphersuites(v *string) {
	o.TlsCiphersuites = v
}

// GetTlsProtocols returns the TlsProtocols field value if set, zero value otherwise.
func (o *InstanceParameters) GetTlsProtocols() *string {
	if o == nil || IsNil(o.TlsProtocols) {
		var ret *string
		return ret
	}
	return o.TlsProtocols
}

// GetTlsProtocolsOk returns a tuple with the TlsProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceParameters) GetTlsProtocolsOk() (*string, bool) {
	if o == nil || IsNil(o.TlsProtocols) {
		return nil, false
	}
	return o.TlsProtocols, true
}

// HasTlsProtocols returns a boolean if a field has been set.
func (o *InstanceParameters) HasTlsProtocols() bool {
	if o != nil && !IsNil(o.TlsProtocols) && !IsNil(o.TlsProtocols) {
		return true
	}

	return false
}

// SetTlsProtocols gets a reference to the given string and assigns it to the TlsProtocols field.
func (o *InstanceParameters) SetTlsProtocols(v *string) {
	o.TlsProtocols = v
}

func (o InstanceParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DownAfterMilliseconds) {
		toSerialize["down-after-milliseconds"] = o.DownAfterMilliseconds
	}
	if !IsNil(o.EnableMonitoring) {
		toSerialize["enable_monitoring"] = o.EnableMonitoring
	}
	if !IsNil(o.FailoverTimeout) {
		toSerialize["failover-timeout"] = o.FailoverTimeout
	}
	if !IsNil(o.Graphite) {
		toSerialize["graphite"] = o.Graphite
	}
	if !IsNil(o.LazyfreeLazyEviction) {
		toSerialize["lazyfree-lazy-eviction"] = o.LazyfreeLazyEviction
	}
	if !IsNil(o.LazyfreeLazyExpire) {
		toSerialize["lazyfree-lazy-expire"] = o.LazyfreeLazyExpire
	}
	if !IsNil(o.LuaTimeLimit) {
		toSerialize["lua-time-limit"] = o.LuaTimeLimit
	}
	if !IsNil(o.MaxDiskThreshold) {
		toSerialize["max_disk_threshold"] = o.MaxDiskThreshold
	}
	if !IsNil(o.Maxclients) {
		toSerialize["maxclients"] = o.Maxclients
	}
	if !IsNil(o.MaxmemoryPolicy) {
		toSerialize["maxmemory-policy"] = o.MaxmemoryPolicy
	}
	if !IsNil(o.MaxmemorySamples) {
		toSerialize["maxmemory-samples"] = o.MaxmemorySamples
	}
	if !IsNil(o.MetricsFrequency) {
		toSerialize["metrics_frequency"] = o.MetricsFrequency
	}
	if !IsNil(o.MetricsPrefix) {
		toSerialize["metrics_prefix"] = o.MetricsPrefix
	}
	if !IsNil(o.MinReplicasMaxLag) {
		toSerialize["min_replicas_max_lag"] = o.MinReplicasMaxLag
	}
	if !IsNil(o.MonitoringInstanceId) {
		toSerialize["monitoring_instance_id"] = o.MonitoringInstanceId
	}
	if !IsNil(o.NotifyKeyspaceEvents) {
		toSerialize["notify-keyspace-events"] = o.NotifyKeyspaceEvents
	}
	if !IsNil(o.SgwAcl) {
		toSerialize["sgw_acl"] = o.SgwAcl
	}
	if !IsNil(o.Snapshot) {
		toSerialize["snapshot"] = o.Snapshot
	}
	if !IsNil(o.Syslog) {
		toSerialize["syslog"] = o.Syslog
	}
	if !IsNil(o.TlsCiphers) {
		toSerialize["tls-ciphers"] = o.TlsCiphers
	}
	if !IsNil(o.TlsCiphersuites) {
		toSerialize["tls-ciphersuites"] = o.TlsCiphersuites
	}
	if !IsNil(o.TlsProtocols) {
		toSerialize["tls-protocols"] = o.TlsProtocols
	}
	return toSerialize, nil
}

type NullableInstanceParameters struct {
	value *InstanceParameters
	isSet bool
}

func (v NullableInstanceParameters) Get() *InstanceParameters {
	return v.value
}

func (v *NullableInstanceParameters) Set(val *InstanceParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceParameters(val *InstanceParameters) *NullableInstanceParameters {
	return &NullableInstanceParameters{value: val, isSet: true}
}

func (v NullableInstanceParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
