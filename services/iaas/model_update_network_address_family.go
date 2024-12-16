/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the UpdateNetworkAddressFamily type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkAddressFamily{}

// UpdateNetworkAddressFamily The addressFamily object for a network update request.
type UpdateNetworkAddressFamily struct {
	Ipv4 *UpdateNetworkIPv4Body `json:"ipv4,omitempty"`
	Ipv6 *UpdateNetworkIPv6Body `json:"ipv6,omitempty"`
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
func (o *UpdateNetworkAddressFamily) GetIpv4() *UpdateNetworkIPv4Body {
	if o == nil || IsNil(o.Ipv4) {
		var ret *UpdateNetworkIPv4Body
		return ret
	}
	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAddressFamily) GetIpv4Ok() (*UpdateNetworkIPv4Body, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *UpdateNetworkAddressFamily) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given UpdateNetworkIPv4Body and assigns it to the Ipv4 field.
func (o *UpdateNetworkAddressFamily) SetIpv4(v *UpdateNetworkIPv4Body) {
	o.Ipv4 = v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *UpdateNetworkAddressFamily) GetIpv6() *UpdateNetworkIPv6Body {
	if o == nil || IsNil(o.Ipv6) {
		var ret *UpdateNetworkIPv6Body
		return ret
	}
	return o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAddressFamily) GetIpv6Ok() (*UpdateNetworkIPv6Body, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *UpdateNetworkAddressFamily) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given UpdateNetworkIPv6Body and assigns it to the Ipv6 field.
func (o *UpdateNetworkAddressFamily) SetIpv6(v *UpdateNetworkIPv6Body) {
	o.Ipv6 = v
}

func (o UpdateNetworkAddressFamily) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if !IsNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
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
