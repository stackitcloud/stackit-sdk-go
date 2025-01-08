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

// checks if the QuotaListBackupGigabytes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaListBackupGigabytes{}

// QuotaListBackupGigabytes Total size in GiB of backups.
type QuotaListBackupGigabytes struct {
	// REQUIRED
	Limit *int64 `json:"limit"`
	// REQUIRED
	Usage *int64 `json:"usage"`
}

type _QuotaListBackupGigabytes QuotaListBackupGigabytes

// NewQuotaListBackupGigabytes instantiates a new QuotaListBackupGigabytes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaListBackupGigabytes(limit *int64, usage *int64) *QuotaListBackupGigabytes {
	this := QuotaListBackupGigabytes{}
	this.Limit = limit
	this.Usage = usage
	return &this
}

// NewQuotaListBackupGigabytesWithDefaults instantiates a new QuotaListBackupGigabytes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaListBackupGigabytesWithDefaults() *QuotaListBackupGigabytes {
	this := QuotaListBackupGigabytes{}
	return &this
}

// GetLimit returns the Limit field value
func (o *QuotaListBackupGigabytes) GetLimit() *int64 {
	if o == nil || IsNil(o.Limit) {
		var ret *int64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *QuotaListBackupGigabytes) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit, true
}

// SetLimit sets field value
func (o *QuotaListBackupGigabytes) SetLimit(v *int64) {
	o.Limit = v
}

// GetUsage returns the Usage field value
func (o *QuotaListBackupGigabytes) GetUsage() *int64 {
	if o == nil || IsNil(o.Usage) {
		var ret *int64
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *QuotaListBackupGigabytes) GetUsageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usage, true
}

// SetUsage sets field value
func (o *QuotaListBackupGigabytes) SetUsage(v *int64) {
	o.Usage = v
}

func (o QuotaListBackupGigabytes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["usage"] = o.Usage
	return toSerialize, nil
}

type NullableQuotaListBackupGigabytes struct {
	value *QuotaListBackupGigabytes
	isSet bool
}

func (v NullableQuotaListBackupGigabytes) Get() *QuotaListBackupGigabytes {
	return v.value
}

func (v *NullableQuotaListBackupGigabytes) Set(val *QuotaListBackupGigabytes) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaListBackupGigabytes) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaListBackupGigabytes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaListBackupGigabytes(val *QuotaListBackupGigabytes) *NullableQuotaListBackupGigabytes {
	return &NullableQuotaListBackupGigabytes{value: val, isSet: true}
}

func (v NullableQuotaListBackupGigabytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaListBackupGigabytes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}