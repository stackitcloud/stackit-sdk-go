/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lbapplication

import (
	"encoding/json"
)

// checks if the LoadbalancerOptionMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadbalancerOptionMetrics{}

// LoadbalancerOptionMetrics struct for LoadbalancerOptionMetrics
type LoadbalancerOptionMetrics struct {
	// Credentials reference for metrics. This reference is created via the observability create endpoint and the credential needs to contain the basic auth username and password for the metrics solution the push URL points to. Then this enables monitoring via remote write for the Application Load Balancer.
	CredentialsRef *string `json:"credentialsRef,omitempty"`
	// The ARGUS/Prometheus remote write Push URL you want the metrics to be shipped to.
	PushUrl *string `json:"pushUrl,omitempty"`
}

// NewLoadbalancerOptionMetrics instantiates a new LoadbalancerOptionMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadbalancerOptionMetrics() *LoadbalancerOptionMetrics {
	this := LoadbalancerOptionMetrics{}
	return &this
}

// NewLoadbalancerOptionMetricsWithDefaults instantiates a new LoadbalancerOptionMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadbalancerOptionMetricsWithDefaults() *LoadbalancerOptionMetrics {
	this := LoadbalancerOptionMetrics{}
	return &this
}

// GetCredentialsRef returns the CredentialsRef field value if set, zero value otherwise.
func (o *LoadbalancerOptionMetrics) GetCredentialsRef() *string {
	if o == nil || IsNil(o.CredentialsRef) {
		var ret *string
		return ret
	}
	return o.CredentialsRef
}

// GetCredentialsRefOk returns a tuple with the CredentialsRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerOptionMetrics) GetCredentialsRefOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialsRef) {
		return nil, false
	}
	return o.CredentialsRef, true
}

// HasCredentialsRef returns a boolean if a field has been set.
func (o *LoadbalancerOptionMetrics) HasCredentialsRef() bool {
	if o != nil && !IsNil(o.CredentialsRef) {
		return true
	}

	return false
}

// SetCredentialsRef gets a reference to the given string and assigns it to the CredentialsRef field.
func (o *LoadbalancerOptionMetrics) SetCredentialsRef(v *string) {
	o.CredentialsRef = v
}

// GetPushUrl returns the PushUrl field value if set, zero value otherwise.
func (o *LoadbalancerOptionMetrics) GetPushUrl() *string {
	if o == nil || IsNil(o.PushUrl) {
		var ret *string
		return ret
	}
	return o.PushUrl
}

// GetPushUrlOk returns a tuple with the PushUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerOptionMetrics) GetPushUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PushUrl) {
		return nil, false
	}
	return o.PushUrl, true
}

// HasPushUrl returns a boolean if a field has been set.
func (o *LoadbalancerOptionMetrics) HasPushUrl() bool {
	if o != nil && !IsNil(o.PushUrl) {
		return true
	}

	return false
}

// SetPushUrl gets a reference to the given string and assigns it to the PushUrl field.
func (o *LoadbalancerOptionMetrics) SetPushUrl(v *string) {
	o.PushUrl = v
}

func (o LoadbalancerOptionMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CredentialsRef) {
		toSerialize["credentialsRef"] = o.CredentialsRef
	}
	if !IsNil(o.PushUrl) {
		toSerialize["pushUrl"] = o.PushUrl
	}
	return toSerialize, nil
}

type NullableLoadbalancerOptionMetrics struct {
	value *LoadbalancerOptionMetrics
	isSet bool
}

func (v NullableLoadbalancerOptionMetrics) Get() *LoadbalancerOptionMetrics {
	return v.value
}

func (v *NullableLoadbalancerOptionMetrics) Set(val *LoadbalancerOptionMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadbalancerOptionMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadbalancerOptionMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadbalancerOptionMetrics(val *LoadbalancerOptionMetrics) *NullableLoadbalancerOptionMetrics {
	return &NullableLoadbalancerOptionMetrics{value: val, isSet: true}
}

func (v NullableLoadbalancerOptionMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadbalancerOptionMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}