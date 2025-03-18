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

// checks if the RescueServerPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RescueServerPayload{}

/*
	types and functions for image
*/

// isNotNullableString
type RescueServerPayloadGetImageAttributeType = *string

func getRescueServerPayloadGetImageAttributeTypeOk(arg RescueServerPayloadGetImageAttributeType) (ret RescueServerPayloadGetImageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRescueServerPayloadGetImageAttributeType(arg *RescueServerPayloadGetImageAttributeType, val RescueServerPayloadGetImageRetType) {
	*arg = &val
}

type RescueServerPayloadGetImageArgType = string
type RescueServerPayloadGetImageRetType = string

// RescueServerPayload struct for RescueServerPayload
type RescueServerPayload struct {
	// Universally Unique Identifier (UUID).
	// REQUIRED
	Image RescueServerPayloadGetImageAttributeType `json:"image"`
}

type _RescueServerPayload RescueServerPayload

// NewRescueServerPayload instantiates a new RescueServerPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRescueServerPayload(image RescueServerPayloadGetImageArgType) *RescueServerPayload {
	this := RescueServerPayload{}
	setRescueServerPayloadGetImageAttributeType(&this.Image, image)
	return &this
}

// NewRescueServerPayloadWithDefaults instantiates a new RescueServerPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRescueServerPayloadWithDefaults() *RescueServerPayload {
	this := RescueServerPayload{}
	return &this
}

// GetImage returns the Image field value
func (o *RescueServerPayload) GetImage() (ret RescueServerPayloadGetImageRetType) {
	ret, _ = o.GetImageOk()
	return ret
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *RescueServerPayload) GetImageOk() (ret RescueServerPayloadGetImageRetType, ok bool) {
	return getRescueServerPayloadGetImageAttributeTypeOk(o.Image)
}

// SetImage sets field value
func (o *RescueServerPayload) SetImage(v RescueServerPayloadGetImageRetType) {
	setRescueServerPayloadGetImageAttributeType(&o.Image, v)
}

func (o RescueServerPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getRescueServerPayloadGetImageAttributeTypeOk(o.Image); ok {
		toSerialize["Image"] = val
	}
	return toSerialize, nil
}

type NullableRescueServerPayload struct {
	value *RescueServerPayload
	isSet bool
}

func (v NullableRescueServerPayload) Get() *RescueServerPayload {
	return v.value
}

func (v *NullableRescueServerPayload) Set(val *RescueServerPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableRescueServerPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableRescueServerPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRescueServerPayload(val *RescueServerPayload) *NullableRescueServerPayload {
	return &NullableRescueServerPayload{value: val, isSet: true}
}

func (v NullableRescueServerPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRescueServerPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
