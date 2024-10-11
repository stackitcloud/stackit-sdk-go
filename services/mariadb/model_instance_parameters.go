/*
STACKIT MariaDB API

The STACKIT MariaDB API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mariadb

import (
	"encoding/json"
)

// checks if the InstanceParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceParameters{}

// InstanceParameters struct for InstanceParameters
type InstanceParameters struct {
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

// NewInstanceParameters instantiates a new InstanceParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceParameters() *InstanceParameters {
	this := InstanceParameters{}
	var enableMonitoring bool = false
	this.EnableMonitoring = &enableMonitoring
	var maxDiskThreshold int64 = 80
	this.MaxDiskThreshold = &maxDiskThreshold
	var metricsFrequency int64 = 10
	this.MetricsFrequency = &metricsFrequency
	return &this
}

// NewInstanceParametersWithDefaults instantiates a new InstanceParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceParametersWithDefaults() *InstanceParameters {
	this := InstanceParameters{}
	var enableMonitoring bool = false
	this.EnableMonitoring = &enableMonitoring
	var maxDiskThreshold int64 = 80
	this.MaxDiskThreshold = &maxDiskThreshold
	var metricsFrequency int64 = 10
	this.MetricsFrequency = &metricsFrequency
	return &this
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
	if o != nil && !IsNil(o.EnableMonitoring) {
		return true
	}

	return false
}

// SetEnableMonitoring gets a reference to the given bool and assigns it to the EnableMonitoring field.
func (o *InstanceParameters) SetEnableMonitoring(v *bool) {
	o.EnableMonitoring = v
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
	if o != nil && !IsNil(o.Graphite) {
		return true
	}

	return false
}

// SetGraphite gets a reference to the given string and assigns it to the Graphite field.
func (o *InstanceParameters) SetGraphite(v *string) {
	o.Graphite = v
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
	if o != nil && !IsNil(o.MaxDiskThreshold) {
		return true
	}

	return false
}

// SetMaxDiskThreshold gets a reference to the given int64 and assigns it to the MaxDiskThreshold field.
func (o *InstanceParameters) SetMaxDiskThreshold(v *int64) {
	o.MaxDiskThreshold = v
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
	if o != nil && !IsNil(o.MetricsFrequency) {
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
	if o != nil && !IsNil(o.MetricsPrefix) {
		return true
	}

	return false
}

// SetMetricsPrefix gets a reference to the given string and assigns it to the MetricsPrefix field.
func (o *InstanceParameters) SetMetricsPrefix(v *string) {
	o.MetricsPrefix = v
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
	if o != nil && !IsNil(o.MonitoringInstanceId) {
		return true
	}

	return false
}

// SetMonitoringInstanceId gets a reference to the given string and assigns it to the MonitoringInstanceId field.
func (o *InstanceParameters) SetMonitoringInstanceId(v *string) {
	o.MonitoringInstanceId = v
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
	if o != nil && !IsNil(o.SgwAcl) {
		return true
	}

	return false
}

// SetSgwAcl gets a reference to the given string and assigns it to the SgwAcl field.
func (o *InstanceParameters) SetSgwAcl(v *string) {
	o.SgwAcl = v
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
	if o != nil && !IsNil(o.Syslog) {
		return true
	}

	return false
}

// SetSyslog gets a reference to the given []string and assigns it to the Syslog field.
func (o *InstanceParameters) SetSyslog(v *[]string) {
	o.Syslog = v
}

func (o InstanceParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableMonitoring) {
		toSerialize["enable_monitoring"] = o.EnableMonitoring
	}
	if !IsNil(o.Graphite) {
		toSerialize["graphite"] = o.Graphite
	}
	if !IsNil(o.MaxDiskThreshold) {
		toSerialize["max_disk_threshold"] = o.MaxDiskThreshold
	}
	if !IsNil(o.MetricsFrequency) {
		toSerialize["metrics_frequency"] = o.MetricsFrequency
	}
	if !IsNil(o.MetricsPrefix) {
		toSerialize["metrics_prefix"] = o.MetricsPrefix
	}
	if !IsNil(o.MonitoringInstanceId) {
		toSerialize["monitoring_instance_id"] = o.MonitoringInstanceId
	}
	if !IsNil(o.SgwAcl) {
		toSerialize["sgw_acl"] = o.SgwAcl
	}
	if !IsNil(o.Syslog) {
		toSerialize["syslog"] = o.Syslog
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
