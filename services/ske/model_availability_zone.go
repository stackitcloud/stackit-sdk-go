/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

import (
	"encoding/json"
)

// checks if the AvailabilityZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailabilityZone{}

/*
	types and functions for name
*/

// isNotNullableString
type AvailabilityZoneGetNameAttributeType = *string

func getAvailabilityZoneGetNameAttributeTypeOk(arg AvailabilityZoneGetNameAttributeType) (ret AvailabilityZoneGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAvailabilityZoneGetNameAttributeType(arg *AvailabilityZoneGetNameAttributeType, val AvailabilityZoneGetNameRetType) {
	*arg = &val
}

type AvailabilityZoneGetNameArgType = string
type AvailabilityZoneGetNameRetType = string

// AvailabilityZone struct for AvailabilityZone
type AvailabilityZone struct {
	Name AvailabilityZoneGetNameAttributeType `json:"name,omitempty"`
}

// NewAvailabilityZone instantiates a new AvailabilityZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityZone() *AvailabilityZone {
	this := AvailabilityZone{}
	return &this
}

// NewAvailabilityZoneWithDefaults instantiates a new AvailabilityZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityZoneWithDefaults() *AvailabilityZone {
	this := AvailabilityZone{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AvailabilityZone) GetName() (res AvailabilityZoneGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityZone) GetNameOk() (ret AvailabilityZoneGetNameRetType, ok bool) {
	return getAvailabilityZoneGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *AvailabilityZone) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AvailabilityZone) SetName(v AvailabilityZoneGetNameRetType) {
	setAvailabilityZoneGetNameAttributeType(&o.Name, v)
}

func (o AvailabilityZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getAvailabilityZoneGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	return toSerialize, nil
}

type NullableAvailabilityZone struct {
	value *AvailabilityZone
	isSet bool
}

func (v NullableAvailabilityZone) Get() *AvailabilityZone {
	return v.value
}

func (v *NullableAvailabilityZone) Set(val *AvailabilityZone) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityZone) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityZone(val *AvailabilityZone) *NullableAvailabilityZone {
	return &NullableAvailabilityZone{value: val, isSet: true}
}

func (v NullableAvailabilityZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
