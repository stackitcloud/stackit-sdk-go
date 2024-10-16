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

// checks if the UpdateAreaAddressFamily type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAreaAddressFamily{}

// UpdateAreaAddressFamily The addressFamily object for a area update request.
type UpdateAreaAddressFamily struct {
	Ipv4 *UpdateAreaIPv4 `json:"ipv4,omitempty"`
}

// NewUpdateAreaAddressFamily instantiates a new UpdateAreaAddressFamily object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAreaAddressFamily() *UpdateAreaAddressFamily {
	this := UpdateAreaAddressFamily{}
	return &this
}

// NewUpdateAreaAddressFamilyWithDefaults instantiates a new UpdateAreaAddressFamily object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAreaAddressFamilyWithDefaults() *UpdateAreaAddressFamily {
	this := UpdateAreaAddressFamily{}
	return &this
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *UpdateAreaAddressFamily) GetIpv4() *UpdateAreaIPv4 {
	if o == nil || IsNil(o.Ipv4) {
		var ret *UpdateAreaIPv4
		return ret
	}
	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAreaAddressFamily) GetIpv4Ok() (*UpdateAreaIPv4, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *UpdateAreaAddressFamily) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given UpdateAreaIPv4 and assigns it to the Ipv4 field.
func (o *UpdateAreaAddressFamily) SetIpv4(v *UpdateAreaIPv4) {
	o.Ipv4 = v
}

func (o UpdateAreaAddressFamily) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	return toSerialize, nil
}

type NullableUpdateAreaAddressFamily struct {
	value *UpdateAreaAddressFamily
	isSet bool
}

func (v NullableUpdateAreaAddressFamily) Get() *UpdateAreaAddressFamily {
	return v.value
}

func (v *NullableUpdateAreaAddressFamily) Set(val *UpdateAreaAddressFamily) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAreaAddressFamily) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAreaAddressFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAreaAddressFamily(val *UpdateAreaAddressFamily) *NullableUpdateAreaAddressFamily {
	return &NullableUpdateAreaAddressFamily{value: val, isSet: true}
}

func (v NullableUpdateAreaAddressFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAreaAddressFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
