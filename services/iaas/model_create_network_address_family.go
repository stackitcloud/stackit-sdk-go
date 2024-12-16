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

// checks if the CreateNetworkAddressFamily type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkAddressFamily{}

// CreateNetworkAddressFamily The addressFamily object for a network create request.
type CreateNetworkAddressFamily struct {
	Ipv4 *CreateNetworkIPv4Body `json:"ipv4,omitempty"`
	Ipv6 *CreateNetworkIPv6Body `json:"ipv6,omitempty"`
}

// NewCreateNetworkAddressFamily instantiates a new CreateNetworkAddressFamily object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkAddressFamily() *CreateNetworkAddressFamily {
	this := CreateNetworkAddressFamily{}
	return &this
}

// NewCreateNetworkAddressFamilyWithDefaults instantiates a new CreateNetworkAddressFamily object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkAddressFamilyWithDefaults() *CreateNetworkAddressFamily {
	this := CreateNetworkAddressFamily{}
	return &this
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *CreateNetworkAddressFamily) GetIpv4() *CreateNetworkIPv4Body {
	if o == nil || IsNil(o.Ipv4) {
		var ret *CreateNetworkIPv4Body
		return ret
	}
	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkAddressFamily) GetIpv4Ok() (*CreateNetworkIPv4Body, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *CreateNetworkAddressFamily) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given CreateNetworkIPv4Body and assigns it to the Ipv4 field.
func (o *CreateNetworkAddressFamily) SetIpv4(v *CreateNetworkIPv4Body) {
	o.Ipv4 = v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *CreateNetworkAddressFamily) GetIpv6() *CreateNetworkIPv6Body {
	if o == nil || IsNil(o.Ipv6) {
		var ret *CreateNetworkIPv6Body
		return ret
	}
	return o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkAddressFamily) GetIpv6Ok() (*CreateNetworkIPv6Body, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *CreateNetworkAddressFamily) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given CreateNetworkIPv6Body and assigns it to the Ipv6 field.
func (o *CreateNetworkAddressFamily) SetIpv6(v *CreateNetworkIPv6Body) {
	o.Ipv6 = v
}

func (o CreateNetworkAddressFamily) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if !IsNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	return toSerialize, nil
}

type NullableCreateNetworkAddressFamily struct {
	value *CreateNetworkAddressFamily
	isSet bool
}

func (v NullableCreateNetworkAddressFamily) Get() *CreateNetworkAddressFamily {
	return v.value
}

func (v *NullableCreateNetworkAddressFamily) Set(val *CreateNetworkAddressFamily) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkAddressFamily) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkAddressFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkAddressFamily(val *CreateNetworkAddressFamily) *NullableCreateNetworkAddressFamily {
	return &NullableCreateNetworkAddressFamily{value: val, isSet: true}
}

func (v NullableCreateNetworkAddressFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkAddressFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
