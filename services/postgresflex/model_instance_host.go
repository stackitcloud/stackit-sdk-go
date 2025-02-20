/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

import (
	"encoding/json"
)

// checks if the InstanceHost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceHost{}

// InstanceHost struct for InstanceHost
type InstanceHost struct {
	HostMetrics *[]InstanceHostMetric `json:"hostMetrics,omitempty"`
	Id          *string               `json:"id,omitempty"`
}

// NewInstanceHost instantiates a new InstanceHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceHost() *InstanceHost {
	this := InstanceHost{}
	return &this
}

// NewInstanceHostWithDefaults instantiates a new InstanceHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceHostWithDefaults() *InstanceHost {
	this := InstanceHost{}
	return &this
}

// GetHostMetrics returns the HostMetrics field value if set, zero value otherwise.
func (o *InstanceHost) GetHostMetrics() *[]InstanceHostMetric {
	if o == nil || IsNil(o.HostMetrics) {
		var ret *[]InstanceHostMetric
		return ret
	}
	return o.HostMetrics
}

// GetHostMetricsOk returns a tuple with the HostMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceHost) GetHostMetricsOk() (*[]InstanceHostMetric, bool) {
	if o == nil || IsNil(o.HostMetrics) {
		return nil, false
	}
	return o.HostMetrics, true
}

// HasHostMetrics returns a boolean if a field has been set.
func (o *InstanceHost) HasHostMetrics() bool {
	if o != nil && !IsNil(o.HostMetrics) {
		return true
	}

	return false
}

// SetHostMetrics gets a reference to the given []InstanceHostMetric and assigns it to the HostMetrics field.
func (o *InstanceHost) SetHostMetrics(v *[]InstanceHostMetric) {
	o.HostMetrics = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstanceHost) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceHost) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InstanceHost) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InstanceHost) SetId(v *string) {
	o.Id = v
}

func (o InstanceHost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HostMetrics) {
		toSerialize["hostMetrics"] = o.HostMetrics
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableInstanceHost struct {
	value *InstanceHost
	isSet bool
}

func (v NullableInstanceHost) Get() *InstanceHost {
	return v.value
}

func (v *NullableInstanceHost) Set(val *InstanceHost) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceHost) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceHost(val *InstanceHost) *NullableInstanceHost {
	return &NullableInstanceHost{value: val, isSet: true}
}

func (v NullableInstanceHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
