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

// checks if the UpdateKeyPairPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateKeyPairPayload{}

/*
	types and functions for labels
*/

// isFreeform
type UpdateKeyPairPayloadGetLabelsAttributeType = *map[string]interface{}
type UpdateKeyPairPayloadGetLabelsArgType = map[string]interface{}
type UpdateKeyPairPayloadGetLabelsRetType = map[string]interface{}

func getUpdateKeyPairPayloadGetLabelsAttributeTypeOk(arg UpdateKeyPairPayloadGetLabelsAttributeType) (ret UpdateKeyPairPayloadGetLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateKeyPairPayloadGetLabelsAttributeType(arg *UpdateKeyPairPayloadGetLabelsAttributeType, val UpdateKeyPairPayloadGetLabelsRetType) {
	*arg = &val
}

// UpdateKeyPairPayload Object that represents an update request body of a public key of an SSH keypair.
type UpdateKeyPairPayload struct {
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels UpdateKeyPairPayloadGetLabelsAttributeType `json:"labels,omitempty"`
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
func (o *UpdateKeyPairPayload) GetLabels() (res UpdateKeyPairPayloadGetLabelsRetType) {
	res, _ = o.GetLabelsOk()
	return
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateKeyPairPayload) GetLabelsOk() (ret UpdateKeyPairPayloadGetLabelsRetType, ok bool) {
	return getUpdateKeyPairPayloadGetLabelsAttributeTypeOk(o.Labels)
}

// HasLabels returns a boolean if a field has been set.
func (o *UpdateKeyPairPayload) HasLabels() bool {
	_, ok := o.GetLabelsOk()
	return ok
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *UpdateKeyPairPayload) SetLabels(v UpdateKeyPairPayloadGetLabelsRetType) {
	setUpdateKeyPairPayloadGetLabelsAttributeType(&o.Labels, v)
}

func (o UpdateKeyPairPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getUpdateKeyPairPayloadGetLabelsAttributeTypeOk(o.Labels); ok {
		toSerialize["Labels"] = val
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
