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

// checks if the VolumeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeSource{}

/*
	types and functions for id
*/

// isNotNullableString
type VolumeSourceGetIdAttributeType = *string

func getVolumeSourceGetIdAttributeTypeOk(arg VolumeSourceGetIdAttributeType) (ret VolumeSourceGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setVolumeSourceGetIdAttributeType(arg *VolumeSourceGetIdAttributeType, val VolumeSourceGetIdRetType) {
	*arg = &val
}

type VolumeSourceGetIdArgType = string
type VolumeSourceGetIdRetType = string

/*
	types and functions for type
*/

// isNotNullableString
type VolumeSourceGetTypeAttributeType = *string

func getVolumeSourceGetTypeAttributeTypeOk(arg VolumeSourceGetTypeAttributeType) (ret VolumeSourceGetTypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setVolumeSourceGetTypeAttributeType(arg *VolumeSourceGetTypeAttributeType, val VolumeSourceGetTypeRetType) {
	*arg = &val
}

type VolumeSourceGetTypeArgType = string
type VolumeSourceGetTypeRetType = string

// VolumeSource The source object of a volume.
type VolumeSource struct {
	// Universally Unique Identifier (UUID).
	// REQUIRED
	Id VolumeSourceGetIdAttributeType `json:"id"`
	// The source types of a volume. Possible values: `image`, `volume`, `snapshot`, `backup`.
	// REQUIRED
	Type VolumeSourceGetTypeAttributeType `json:"type"`
}

type _VolumeSource VolumeSource

// NewVolumeSource instantiates a new VolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeSource(id VolumeSourceGetIdArgType, type_ VolumeSourceGetTypeArgType) *VolumeSource {
	this := VolumeSource{}
	setVolumeSourceGetIdAttributeType(&this.Id, id)
	setVolumeSourceGetTypeAttributeType(&this.Type, type_)
	return &this
}

// NewVolumeSourceWithDefaults instantiates a new VolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeSourceWithDefaults() *VolumeSource {
	this := VolumeSource{}
	return &this
}

// GetId returns the Id field value
func (o *VolumeSource) GetId() (ret VolumeSourceGetIdRetType) {
	ret, _ = o.GetIdOk()
	return ret
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VolumeSource) GetIdOk() (ret VolumeSourceGetIdRetType, ok bool) {
	return getVolumeSourceGetIdAttributeTypeOk(o.Id)
}

// SetId sets field value
func (o *VolumeSource) SetId(v VolumeSourceGetIdRetType) {
	setVolumeSourceGetIdAttributeType(&o.Id, v)
}

// GetType returns the Type field value
func (o *VolumeSource) GetType() (ret VolumeSourceGetTypeRetType) {
	ret, _ = o.GetTypeOk()
	return ret
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VolumeSource) GetTypeOk() (ret VolumeSourceGetTypeRetType, ok bool) {
	return getVolumeSourceGetTypeAttributeTypeOk(o.Type)
}

// SetType sets field value
func (o *VolumeSource) SetType(v VolumeSourceGetTypeRetType) {
	setVolumeSourceGetTypeAttributeType(&o.Type, v)
}

func (o VolumeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getVolumeSourceGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getVolumeSourceGetTypeAttributeTypeOk(o.Type); ok {
		toSerialize["Type"] = val
	}
	return toSerialize, nil
}

type NullableVolumeSource struct {
	value *VolumeSource
	isSet bool
}

func (v NullableVolumeSource) Get() *VolumeSource {
	return v.value
}

func (v *NullableVolumeSource) Set(val *VolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeSource(val *VolumeSource) *NullableVolumeSource {
	return &NullableVolumeSource{value: val, isSet: true}
}

func (v NullableVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
