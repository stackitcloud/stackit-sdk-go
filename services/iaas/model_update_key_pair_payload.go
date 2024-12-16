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

// checks if the UpdateKeyPairPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateKeyPairPayload{}

// UpdateKeyPairPayload Object that represents an update request body of a public key of an SSH keypair.
type UpdateKeyPairPayload struct {
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
}

// NewUpdateKeyPairPayload instantiates a new UpdateKeyPairPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateKeyPairPayload() *UpdateKeyPairPayload {
	this := UpdateKeyPairPayload{}
	return &this
}

// NewUpdateKeyPairPayloadWithDefaults instantiates a new UpdateKeyPairPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateKeyPairPayloadWithDefaults() *UpdateKeyPairPayload {
	this := UpdateKeyPairPayload{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateKeyPairPayload) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateKeyPairPayload) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *UpdateKeyPairPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *UpdateKeyPairPayload) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

func (o UpdateKeyPairPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableUpdateKeyPairPayload struct {
	value *UpdateKeyPairPayload
	isSet bool
}

func (v NullableUpdateKeyPairPayload) Get() *UpdateKeyPairPayload {
	return v.value
}

func (v *NullableUpdateKeyPairPayload) Set(val *UpdateKeyPairPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateKeyPairPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateKeyPairPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateKeyPairPayload(val *UpdateKeyPairPayload) *NullableUpdateKeyPairPayload {
	return &NullableUpdateKeyPairPayload{value: val, isSet: true}
}

func (v NullableUpdateKeyPairPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateKeyPairPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
