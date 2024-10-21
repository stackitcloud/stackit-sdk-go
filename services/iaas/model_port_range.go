/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the PortRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortRange{}

// PortRange Object that represents a range of ports.
type PortRange struct {
	// The maximum port number. Should be greater or equal to the minimum.
	// REQUIRED
	Max *int64 `json:"max"`
	// The minimum port number. Should be less or equal to the minimum.
	// REQUIRED
	Min *int64 `json:"min"`
}

type _PortRange PortRange

// NewPortRange instantiates a new PortRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortRange(max *int64, min *int64) *PortRange {
	this := PortRange{}
	this.Max = max
	this.Min = min
	return &this
}

// NewPortRangeWithDefaults instantiates a new PortRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortRangeWithDefaults() *PortRange {
	this := PortRange{}
	return &this
}

// GetMax returns the Max field value
func (o *PortRange) GetMax() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *PortRange) GetMaxOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Max, true
}

// SetMax sets field value
func (o *PortRange) SetMax(v *int64) {
	o.Max = v
}

// GetMin returns the Min field value
func (o *PortRange) GetMin() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *PortRange) GetMinOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Min, true
}

// SetMin sets field value
func (o *PortRange) SetMin(v *int64) {
	o.Min = v
}

func (o PortRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["max"] = o.Max
	toSerialize["min"] = o.Min
	return toSerialize, nil
}

type NullablePortRange struct {
	value *PortRange
	isSet bool
}

func (v NullablePortRange) Get() *PortRange {
	return v.value
}

func (v *NullablePortRange) Set(val *PortRange) {
	v.value = val
	v.isSet = true
}

func (v NullablePortRange) IsSet() bool {
	return v.isSet
}

func (v *NullablePortRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortRange(val *PortRange) *NullablePortRange {
	return &NullablePortRange{value: val, isSet: true}
}

func (v NullablePortRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}