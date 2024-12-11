/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
)

// checks if the AreaPrefixConfigIPv4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AreaPrefixConfigIPv4{}

// AreaPrefixConfigIPv4 The IPv4 prefix config for a network area.
type AreaPrefixConfigIPv4 struct {
	// The default prefix length for networks in the network area.
	DefaultPrefixLen *int64 `json:"defaultPrefixLen,omitempty"`
	// The maximal prefix length for networks in the network area.
	MaxPrefixLen *int64 `json:"maxPrefixLen,omitempty"`
	// The minimal prefix length for networks in the network area.
	MinPrefixLen *int64 `json:"minPrefixLen,omitempty"`
}

// NewAreaPrefixConfigIPv4 instantiates a new AreaPrefixConfigIPv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAreaPrefixConfigIPv4() *AreaPrefixConfigIPv4 {
	this := AreaPrefixConfigIPv4{}
	var defaultPrefixLen int64 = 25
	this.DefaultPrefixLen = &defaultPrefixLen
	var maxPrefixLen int64 = 29
	this.MaxPrefixLen = &maxPrefixLen
	var minPrefixLen int64 = 24
	this.MinPrefixLen = &minPrefixLen
	return &this
}

// NewAreaPrefixConfigIPv4WithDefaults instantiates a new AreaPrefixConfigIPv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAreaPrefixConfigIPv4WithDefaults() *AreaPrefixConfigIPv4 {
	this := AreaPrefixConfigIPv4{}
	var defaultPrefixLen int64 = 25
	this.DefaultPrefixLen = &defaultPrefixLen
	var maxPrefixLen int64 = 29
	this.MaxPrefixLen = &maxPrefixLen
	var minPrefixLen int64 = 24
	this.MinPrefixLen = &minPrefixLen
	return &this
}

// GetDefaultPrefixLen returns the DefaultPrefixLen field value if set, zero value otherwise.
func (o *AreaPrefixConfigIPv4) GetDefaultPrefixLen() *int64 {
	if o == nil || IsNil(o.DefaultPrefixLen) {
		var ret *int64
		return ret
	}
	return o.DefaultPrefixLen
}

// GetDefaultPrefixLenOk returns a tuple with the DefaultPrefixLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaPrefixConfigIPv4) GetDefaultPrefixLenOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultPrefixLen) {
		return nil, false
	}
	return o.DefaultPrefixLen, true
}

// HasDefaultPrefixLen returns a boolean if a field has been set.
func (o *AreaPrefixConfigIPv4) HasDefaultPrefixLen() bool {
	if o != nil && !IsNil(o.DefaultPrefixLen) && !IsNil(o.DefaultPrefixLen) {
		return true
	}

	return false
}

// SetDefaultPrefixLen gets a reference to the given int64 and assigns it to the DefaultPrefixLen field.
func (o *AreaPrefixConfigIPv4) SetDefaultPrefixLen(v *int64) {
	o.DefaultPrefixLen = v
}

// GetMaxPrefixLen returns the MaxPrefixLen field value if set, zero value otherwise.
func (o *AreaPrefixConfigIPv4) GetMaxPrefixLen() *int64 {
	if o == nil || IsNil(o.MaxPrefixLen) {
		var ret *int64
		return ret
	}
	return o.MaxPrefixLen
}

// GetMaxPrefixLenOk returns a tuple with the MaxPrefixLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaPrefixConfigIPv4) GetMaxPrefixLenOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxPrefixLen) {
		return nil, false
	}
	return o.MaxPrefixLen, true
}

// HasMaxPrefixLen returns a boolean if a field has been set.
func (o *AreaPrefixConfigIPv4) HasMaxPrefixLen() bool {
	if o != nil && !IsNil(o.MaxPrefixLen) && !IsNil(o.MaxPrefixLen) {
		return true
	}

	return false
}

// SetMaxPrefixLen gets a reference to the given int64 and assigns it to the MaxPrefixLen field.
func (o *AreaPrefixConfigIPv4) SetMaxPrefixLen(v *int64) {
	o.MaxPrefixLen = v
}

// GetMinPrefixLen returns the MinPrefixLen field value if set, zero value otherwise.
func (o *AreaPrefixConfigIPv4) GetMinPrefixLen() *int64 {
	if o == nil || IsNil(o.MinPrefixLen) {
		var ret *int64
		return ret
	}
	return o.MinPrefixLen
}

// GetMinPrefixLenOk returns a tuple with the MinPrefixLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaPrefixConfigIPv4) GetMinPrefixLenOk() (*int64, bool) {
	if o == nil || IsNil(o.MinPrefixLen) {
		return nil, false
	}
	return o.MinPrefixLen, true
}

// HasMinPrefixLen returns a boolean if a field has been set.
func (o *AreaPrefixConfigIPv4) HasMinPrefixLen() bool {
	if o != nil && !IsNil(o.MinPrefixLen) && !IsNil(o.MinPrefixLen) {
		return true
	}

	return false
}

// SetMinPrefixLen gets a reference to the given int64 and assigns it to the MinPrefixLen field.
func (o *AreaPrefixConfigIPv4) SetMinPrefixLen(v *int64) {
	o.MinPrefixLen = v
}

func (o AreaPrefixConfigIPv4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultPrefixLen) {
		toSerialize["defaultPrefixLen"] = o.DefaultPrefixLen
	}
	if !IsNil(o.MaxPrefixLen) {
		toSerialize["maxPrefixLen"] = o.MaxPrefixLen
	}
	if !IsNil(o.MinPrefixLen) {
		toSerialize["minPrefixLen"] = o.MinPrefixLen
	}
	return toSerialize, nil
}

type NullableAreaPrefixConfigIPv4 struct {
	value *AreaPrefixConfigIPv4
	isSet bool
}

func (v NullableAreaPrefixConfigIPv4) Get() *AreaPrefixConfigIPv4 {
	return v.value
}

func (v *NullableAreaPrefixConfigIPv4) Set(val *AreaPrefixConfigIPv4) {
	v.value = val
	v.isSet = true
}

func (v NullableAreaPrefixConfigIPv4) IsSet() bool {
	return v.isSet
}

func (v *NullableAreaPrefixConfigIPv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAreaPrefixConfigIPv4(val *AreaPrefixConfigIPv4) *NullableAreaPrefixConfigIPv4 {
	return &NullableAreaPrefixConfigIPv4{value: val, isSet: true}
}

func (v NullableAreaPrefixConfigIPv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAreaPrefixConfigIPv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
