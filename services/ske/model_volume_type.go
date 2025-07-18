/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

import (
	"encoding/json"
)

// checks if the VolumeType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeType{}

/*
	types and functions for name
*/

// isNotNullableString
type VolumeTypeGetNameAttributeType = *string

func getVolumeTypeGetNameAttributeTypeOk(arg VolumeTypeGetNameAttributeType) (ret VolumeTypeGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setVolumeTypeGetNameAttributeType(arg *VolumeTypeGetNameAttributeType, val VolumeTypeGetNameRetType) {
	*arg = &val
}

type VolumeTypeGetNameArgType = string
type VolumeTypeGetNameRetType = string

// VolumeType struct for VolumeType
type VolumeType struct {
	Name VolumeTypeGetNameAttributeType `json:"name,omitempty"`
}

// NewVolumeType instantiates a new VolumeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeType() *VolumeType {
	this := VolumeType{}
	return &this
}

// NewVolumeTypeWithDefaults instantiates a new VolumeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeTypeWithDefaults() *VolumeType {
	this := VolumeType{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VolumeType) GetName() (res VolumeTypeGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeType) GetNameOk() (ret VolumeTypeGetNameRetType, ok bool) {
	return getVolumeTypeGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *VolumeType) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VolumeType) SetName(v VolumeTypeGetNameRetType) {
	setVolumeTypeGetNameAttributeType(&o.Name, v)
}

func (o VolumeType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getVolumeTypeGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	return toSerialize, nil
}

type NullableVolumeType struct {
	value *VolumeType
	isSet bool
}

func (v NullableVolumeType) Get() *VolumeType {
	return v.value
}

func (v *NullableVolumeType) Set(val *VolumeType) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeType) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeType(val *VolumeType) *NullableVolumeType {
	return &NullableVolumeType{value: val, isSet: true}
}

func (v NullableVolumeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
