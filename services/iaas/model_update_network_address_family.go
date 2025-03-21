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

// checks if the UpdateNetworkAddressFamily type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkAddressFamily{}

/*
	types and functions for ipv4
*/

// isModel
type UpdateNetworkAddressFamilyGetIpv4AttributeType = *UpdateNetworkIPv4Body
type UpdateNetworkAddressFamilyGetIpv4ArgType = UpdateNetworkIPv4Body
type UpdateNetworkAddressFamilyGetIpv4RetType = UpdateNetworkIPv4Body

func getUpdateNetworkAddressFamilyGetIpv4AttributeTypeOk(arg UpdateNetworkAddressFamilyGetIpv4AttributeType) (ret UpdateNetworkAddressFamilyGetIpv4RetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateNetworkAddressFamilyGetIpv4AttributeType(arg *UpdateNetworkAddressFamilyGetIpv4AttributeType, val UpdateNetworkAddressFamilyGetIpv4RetType) {
	*arg = &val
}

/*
	types and functions for ipv6
*/

// isModel
type UpdateNetworkAddressFamilyGetIpv6AttributeType = *UpdateNetworkIPv6Body
type UpdateNetworkAddressFamilyGetIpv6ArgType = UpdateNetworkIPv6Body
type UpdateNetworkAddressFamilyGetIpv6RetType = UpdateNetworkIPv6Body

func getUpdateNetworkAddressFamilyGetIpv6AttributeTypeOk(arg UpdateNetworkAddressFamilyGetIpv6AttributeType) (ret UpdateNetworkAddressFamilyGetIpv6RetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateNetworkAddressFamilyGetIpv6AttributeType(arg *UpdateNetworkAddressFamilyGetIpv6AttributeType, val UpdateNetworkAddressFamilyGetIpv6RetType) {
	*arg = &val
}

// UpdateNetworkAddressFamily The addressFamily object for a network update request.
type UpdateNetworkAddressFamily struct {
	Ipv4 UpdateNetworkAddressFamilyGetIpv4AttributeType `json:"ipv4,omitempty"`
	Ipv6 UpdateNetworkAddressFamilyGetIpv6AttributeType `json:"ipv6,omitempty"`
}

// NewUpdateNetworkAddressFamily instantiates a new UpdateNetworkAddressFamily object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkAddressFamily() *UpdateNetworkAddressFamily {
	this := UpdateNetworkAddressFamily{}
	return &this
}

// NewUpdateNetworkAddressFamilyWithDefaults instantiates a new UpdateNetworkAddressFamily object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkAddressFamilyWithDefaults() *UpdateNetworkAddressFamily {
	this := UpdateNetworkAddressFamily{}
	return &this
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *UpdateNetworkAddressFamily) GetIpv4() (res UpdateNetworkAddressFamilyGetIpv4RetType) {
	res, _ = o.GetIpv4Ok()
	return
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAddressFamily) GetIpv4Ok() (ret UpdateNetworkAddressFamilyGetIpv4RetType, ok bool) {
	return getUpdateNetworkAddressFamilyGetIpv4AttributeTypeOk(o.Ipv4)
}

// HasIpv4 returns a boolean if a field has been set.
func (o *UpdateNetworkAddressFamily) HasIpv4() bool {
	_, ok := o.GetIpv4Ok()
	return ok
}

// SetIpv4 gets a reference to the given UpdateNetworkIPv4Body and assigns it to the Ipv4 field.
func (o *UpdateNetworkAddressFamily) SetIpv4(v UpdateNetworkAddressFamilyGetIpv4RetType) {
	setUpdateNetworkAddressFamilyGetIpv4AttributeType(&o.Ipv4, v)
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *UpdateNetworkAddressFamily) GetIpv6() (res UpdateNetworkAddressFamilyGetIpv6RetType) {
	res, _ = o.GetIpv6Ok()
	return
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAddressFamily) GetIpv6Ok() (ret UpdateNetworkAddressFamilyGetIpv6RetType, ok bool) {
	return getUpdateNetworkAddressFamilyGetIpv6AttributeTypeOk(o.Ipv6)
}

// HasIpv6 returns a boolean if a field has been set.
func (o *UpdateNetworkAddressFamily) HasIpv6() bool {
	_, ok := o.GetIpv6Ok()
	return ok
}

// SetIpv6 gets a reference to the given UpdateNetworkIPv6Body and assigns it to the Ipv6 field.
func (o *UpdateNetworkAddressFamily) SetIpv6(v UpdateNetworkAddressFamilyGetIpv6RetType) {
	setUpdateNetworkAddressFamilyGetIpv6AttributeType(&o.Ipv6, v)
}

func (o UpdateNetworkAddressFamily) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getUpdateNetworkAddressFamilyGetIpv4AttributeTypeOk(o.Ipv4); ok {
		toSerialize["Ipv4"] = val
	}
	if val, ok := getUpdateNetworkAddressFamilyGetIpv6AttributeTypeOk(o.Ipv6); ok {
		toSerialize["Ipv6"] = val
	}
	return toSerialize, nil
}

type NullableUpdateNetworkAddressFamily struct {
	value *UpdateNetworkAddressFamily
	isSet bool
}

func (v NullableUpdateNetworkAddressFamily) Get() *UpdateNetworkAddressFamily {
	return v.value
}

func (v *NullableUpdateNetworkAddressFamily) Set(val *UpdateNetworkAddressFamily) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkAddressFamily) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkAddressFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkAddressFamily(val *UpdateNetworkAddressFamily) *NullableUpdateNetworkAddressFamily {
	return &NullableUpdateNetworkAddressFamily{value: val, isSet: true}
}

func (v NullableUpdateNetworkAddressFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkAddressFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
