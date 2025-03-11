/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

import (
	"encoding/json"
)

// checks if the LoadbalancerOptionObservability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadbalancerOptionObservability{}

// LoadbalancerOptionObservability We offer Load Balancer metrics observability via ARGUS or external solutions. Not changeable after creation.
type LoadbalancerOptionObservability struct {
	Logs    *LoadbalancerOptionLogs    `json:"logs,omitempty"`
	Metrics *LoadbalancerOptionMetrics `json:"metrics,omitempty"`
}

// NewLoadbalancerOptionObservability instantiates a new LoadbalancerOptionObservability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadbalancerOptionObservability() *LoadbalancerOptionObservability {
	this := LoadbalancerOptionObservability{}
	return &this
}

// NewLoadbalancerOptionObservabilityWithDefaults instantiates a new LoadbalancerOptionObservability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadbalancerOptionObservabilityWithDefaults() *LoadbalancerOptionObservability {
	this := LoadbalancerOptionObservability{}
	return &this
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *LoadbalancerOptionObservability) GetLogs() *LoadbalancerOptionLogs {
	if o == nil || IsNil(o.Logs) {
		var ret *LoadbalancerOptionLogs
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerOptionObservability) GetLogsOk() (*LoadbalancerOptionLogs, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *LoadbalancerOptionObservability) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given LoadbalancerOptionLogs and assigns it to the Logs field.
func (o *LoadbalancerOptionObservability) SetLogs(v *LoadbalancerOptionLogs) {
	o.Logs = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *LoadbalancerOptionObservability) GetMetrics() *LoadbalancerOptionMetrics {
	if o == nil || IsNil(o.Metrics) {
		var ret *LoadbalancerOptionMetrics
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerOptionObservability) GetMetricsOk() (*LoadbalancerOptionMetrics, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *LoadbalancerOptionObservability) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given LoadbalancerOptionMetrics and assigns it to the Metrics field.
func (o *LoadbalancerOptionObservability) SetMetrics(v *LoadbalancerOptionMetrics) {
	o.Metrics = v
}

func (o LoadbalancerOptionObservability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	return toSerialize, nil
}

type NullableLoadbalancerOptionObservability struct {
	value *LoadbalancerOptionObservability
	isSet bool
}

func (v NullableLoadbalancerOptionObservability) Get() *LoadbalancerOptionObservability {
	return v.value
}

func (v *NullableLoadbalancerOptionObservability) Set(val *LoadbalancerOptionObservability) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadbalancerOptionObservability) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadbalancerOptionObservability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadbalancerOptionObservability(val *LoadbalancerOptionObservability) *NullableLoadbalancerOptionObservability {
	return &NullableLoadbalancerOptionObservability{value: val, isSet: true}
}

func (v NullableLoadbalancerOptionObservability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadbalancerOptionObservability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
