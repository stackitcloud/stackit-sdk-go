/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"encoding/json"
)

// checks if the TLSConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TLSConfig{}

// TLSConfig struct for TLSConfig
type TLSConfig struct {
	InsecureSkipVerify *bool `json:"insecureSkipVerify,omitempty"`
}

// NewTLSConfig instantiates a new TLSConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSConfig() *TLSConfig {
	this := TLSConfig{}
	var insecureSkipVerify bool = false
	this.InsecureSkipVerify = &insecureSkipVerify
	return &this
}

// NewTLSConfigWithDefaults instantiates a new TLSConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSConfigWithDefaults() *TLSConfig {
	this := TLSConfig{}
	var insecureSkipVerify bool = false
	this.InsecureSkipVerify = &insecureSkipVerify
	return &this
}

// GetInsecureSkipVerify returns the InsecureSkipVerify field value if set, zero value otherwise.
func (o *TLSConfig) GetInsecureSkipVerify() *bool {
	if o == nil || IsNil(o.InsecureSkipVerify) {
		var ret *bool
		return ret
	}
	return o.InsecureSkipVerify
}

// GetInsecureSkipVerifyOk returns a tuple with the InsecureSkipVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfig) GetInsecureSkipVerifyOk() (*bool, bool) {
	if o == nil || IsNil(o.InsecureSkipVerify) {
		return nil, false
	}
	return o.InsecureSkipVerify, true
}

// HasInsecureSkipVerify returns a boolean if a field has been set.
func (o *TLSConfig) HasInsecureSkipVerify() bool {
	if o != nil && !IsNil(o.InsecureSkipVerify) {
		return true
	}

	return false
}

// SetInsecureSkipVerify gets a reference to the given bool and assigns it to the InsecureSkipVerify field.
func (o *TLSConfig) SetInsecureSkipVerify(v *bool) {
	o.InsecureSkipVerify = v
}

func (o TLSConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InsecureSkipVerify) {
		toSerialize["insecureSkipVerify"] = o.InsecureSkipVerify
	}
	return toSerialize, nil
}

type NullableTLSConfig struct {
	value *TLSConfig
	isSet bool
}

func (v NullableTLSConfig) Get() *TLSConfig {
	return v.value
}

func (v *NullableTLSConfig) Set(val *TLSConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTLSConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTLSConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTLSConfig(val *TLSConfig) *NullableTLSConfig {
	return &NullableTLSConfig{value: val, isSet: true}
}

func (v NullableTLSConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTLSConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
