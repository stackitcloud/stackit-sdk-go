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

// checks if the QuotaListBackupGigabytes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaListBackupGigabytes{}

/*
	types and functions for limit
*/

// isLong
type QuotaListBackupGigabytesGetLimitAttributeType = *int64
type QuotaListBackupGigabytesGetLimitArgType = int64
type QuotaListBackupGigabytesGetLimitRetType = int64

func getQuotaListBackupGigabytesGetLimitAttributeTypeOk(arg QuotaListBackupGigabytesGetLimitAttributeType) (ret QuotaListBackupGigabytesGetLimitRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListBackupGigabytesGetLimitAttributeType(arg *QuotaListBackupGigabytesGetLimitAttributeType, val QuotaListBackupGigabytesGetLimitRetType) {
	*arg = &val
}

/*
	types and functions for usage
*/

// isLong
type QuotaListBackupGigabytesGetUsageAttributeType = *int64
type QuotaListBackupGigabytesGetUsageArgType = int64
type QuotaListBackupGigabytesGetUsageRetType = int64

func getQuotaListBackupGigabytesGetUsageAttributeTypeOk(arg QuotaListBackupGigabytesGetUsageAttributeType) (ret QuotaListBackupGigabytesGetUsageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQuotaListBackupGigabytesGetUsageAttributeType(arg *QuotaListBackupGigabytesGetUsageAttributeType, val QuotaListBackupGigabytesGetUsageRetType) {
	*arg = &val
}

// QuotaListBackupGigabytes Total size in GiB of backups.
type QuotaListBackupGigabytes struct {
	// REQUIRED
	Limit QuotaListBackupGigabytesGetLimitAttributeType `json:"limit"`
	// REQUIRED
	Usage QuotaListBackupGigabytesGetUsageAttributeType `json:"usage"`
}

type _QuotaListBackupGigabytes QuotaListBackupGigabytes

// NewQuotaListBackupGigabytes instantiates a new QuotaListBackupGigabytes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaListBackupGigabytes(limit QuotaListBackupGigabytesGetLimitArgType, usage QuotaListBackupGigabytesGetUsageArgType) *QuotaListBackupGigabytes {
	this := QuotaListBackupGigabytes{}
	setQuotaListBackupGigabytesGetLimitAttributeType(&this.Limit, limit)
	setQuotaListBackupGigabytesGetUsageAttributeType(&this.Usage, usage)
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
func (o *QuotaListBackupGigabytes) GetLimit() (ret QuotaListBackupGigabytesGetLimitRetType) {
	ret, _ = o.GetLimitOk()
	return ret
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *QuotaListBackupGigabytes) GetLimitOk() (ret QuotaListBackupGigabytesGetLimitRetType, ok bool) {
	return getQuotaListBackupGigabytesGetLimitAttributeTypeOk(o.Limit)
}

// SetLimit sets field value
func (o *QuotaListBackupGigabytes) SetLimit(v QuotaListBackupGigabytesGetLimitRetType) {
	setQuotaListBackupGigabytesGetLimitAttributeType(&o.Limit, v)
}

// GetUsage returns the Usage field value
func (o *QuotaListBackupGigabytes) GetUsage() (ret QuotaListBackupGigabytesGetUsageRetType) {
	ret, _ = o.GetUsageOk()
	return ret
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *QuotaListBackupGigabytes) GetUsageOk() (ret QuotaListBackupGigabytesGetUsageRetType, ok bool) {
	return getQuotaListBackupGigabytesGetUsageAttributeTypeOk(o.Usage)
}

// SetUsage sets field value
func (o *QuotaListBackupGigabytes) SetUsage(v QuotaListBackupGigabytesGetUsageRetType) {
	setQuotaListBackupGigabytesGetUsageAttributeType(&o.Usage, v)
}

func (o QuotaListBackupGigabytes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getQuotaListBackupGigabytesGetLimitAttributeTypeOk(o.Limit); ok {
		toSerialize["Limit"] = val
	}
	if val, ok := getQuotaListBackupGigabytesGetUsageAttributeTypeOk(o.Usage); ok {
		toSerialize["Usage"] = val
	}
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
