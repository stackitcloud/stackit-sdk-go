/*
STACKIT Marketplace API

API to manage STACKIT Marketplace.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackitmarketplace

import (
	"encoding/json"
)

// checks if the VendorId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VendorId{}

// VendorId The vendor ID.
type VendorId struct {
}

// NewVendorId instantiates a new VendorId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorId() *VendorId {
	this := VendorId{}
	return &this
}

// NewVendorIdWithDefaults instantiates a new VendorId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorIdWithDefaults() *VendorId {
	this := VendorId{}
	return &this
}

func (o VendorId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableVendorId struct {
	value *VendorId
	isSet bool
}

func (v NullableVendorId) Get() *VendorId {
	return v.value
}

func (v *NullableVendorId) Set(val *VendorId) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorId) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorId(val *VendorId) *NullableVendorId {
	return &NullableVendorId{value: val, isSet: true}
}

func (v NullableVendorId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
