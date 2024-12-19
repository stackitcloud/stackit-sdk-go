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

// checks if the QuotaListNics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaListNics{}

// QuotaListNics Number of network interfaces.
type QuotaListNics struct {
	// REQUIRED
	Limit *int64 `json:"limit"`
	// REQUIRED
	Usage *int64 `json:"usage"`
}

type _QuotaListNics QuotaListNics

// NewQuotaListNics instantiates a new QuotaListNics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaListNics(limit *int64, usage *int64) *QuotaListNics {
	this := QuotaListNics{}
	this.Limit = limit
	this.Usage = usage
	return &this
}

// NewQuotaListNicsWithDefaults instantiates a new QuotaListNics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaListNicsWithDefaults() *QuotaListNics {
	this := QuotaListNics{}
	return &this
}

// GetLimit returns the Limit field value
func (o *QuotaListNics) GetLimit() *int64 {
	if o == nil || IsNil(o.Limit) {
		var ret *int64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *QuotaListNics) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit, true
}

// SetLimit sets field value
func (o *QuotaListNics) SetLimit(v *int64) {
	o.Limit = v
}

// GetUsage returns the Usage field value
func (o *QuotaListNics) GetUsage() *int64 {
	if o == nil || IsNil(o.Usage) {
		var ret *int64
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *QuotaListNics) GetUsageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usage, true
}

// SetUsage sets field value
func (o *QuotaListNics) SetUsage(v *int64) {
	o.Usage = v
}

func (o QuotaListNics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["usage"] = o.Usage
	return toSerialize, nil
}

type NullableQuotaListNics struct {
	value *QuotaListNics
	isSet bool
}

func (v NullableQuotaListNics) Get() *QuotaListNics {
	return v.value
}

func (v *NullableQuotaListNics) Set(val *QuotaListNics) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaListNics) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaListNics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaListNics(val *QuotaListNics) *NullableQuotaListNics {
	return &NullableQuotaListNics{value: val, isSet: true}
}

func (v NullableQuotaListNics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaListNics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
