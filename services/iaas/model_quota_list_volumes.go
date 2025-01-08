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

// checks if the QuotaListVolumes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaListVolumes{}

// QuotaListVolumes Number of volumes.
type QuotaListVolumes struct {
	// REQUIRED
	Limit *int64 `json:"limit"`
	// REQUIRED
	Usage *int64 `json:"usage"`
}

type _QuotaListVolumes QuotaListVolumes

// NewQuotaListVolumes instantiates a new QuotaListVolumes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaListVolumes(limit *int64, usage *int64) *QuotaListVolumes {
	this := QuotaListVolumes{}
	this.Limit = limit
	this.Usage = usage
	return &this
}

// NewQuotaListVolumesWithDefaults instantiates a new QuotaListVolumes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaListVolumesWithDefaults() *QuotaListVolumes {
	this := QuotaListVolumes{}
	return &this
}

// GetLimit returns the Limit field value
func (o *QuotaListVolumes) GetLimit() *int64 {
	if o == nil || IsNil(o.Limit) {
		var ret *int64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *QuotaListVolumes) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit, true
}

// SetLimit sets field value
func (o *QuotaListVolumes) SetLimit(v *int64) {
	o.Limit = v
}

// GetUsage returns the Usage field value
func (o *QuotaListVolumes) GetUsage() *int64 {
	if o == nil || IsNil(o.Usage) {
		var ret *int64
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *QuotaListVolumes) GetUsageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usage, true
}

// SetUsage sets field value
func (o *QuotaListVolumes) SetUsage(v *int64) {
	o.Usage = v
}

func (o QuotaListVolumes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["usage"] = o.Usage
	return toSerialize, nil
}

type NullableQuotaListVolumes struct {
	value *QuotaListVolumes
	isSet bool
}

func (v NullableQuotaListVolumes) Get() *QuotaListVolumes {
	return v.value
}

func (v *NullableQuotaListVolumes) Set(val *QuotaListVolumes) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaListVolumes) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaListVolumes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaListVolumes(val *QuotaListVolumes) *NullableQuotaListVolumes {
	return &NullableQuotaListVolumes{value: val, isSet: true}
}

func (v NullableQuotaListVolumes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaListVolumes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}