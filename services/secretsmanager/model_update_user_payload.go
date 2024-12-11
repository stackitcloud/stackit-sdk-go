/*
STACKIT Secrets Manager API

This API provides endpoints for managing the Secrets-Manager.

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretsmanager

import (
	"encoding/json"
)

// checks if the UpdateUserPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserPayload{}

// UpdateUserPayload struct for UpdateUserPayload
type UpdateUserPayload struct {
	// Is true if the user has write access to the secrets engine. Is false for a read-only user.
	Write *bool `json:"write,omitempty"`
}

// NewUpdateUserPayload instantiates a new UpdateUserPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserPayload() *UpdateUserPayload {
	this := UpdateUserPayload{}
	return &this
}

// NewUpdateUserPayloadWithDefaults instantiates a new UpdateUserPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserPayloadWithDefaults() *UpdateUserPayload {
	this := UpdateUserPayload{}
	return &this
}

// GetWrite returns the Write field value if set, zero value otherwise.
func (o *UpdateUserPayload) GetWrite() *bool {
	if o == nil || IsNil(o.Write) {
		var ret *bool
		return ret
	}
	return o.Write
}

// GetWriteOk returns a tuple with the Write field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserPayload) GetWriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Write) {
		return nil, false
	}
	return o.Write, true
}

// HasWrite returns a boolean if a field has been set.
func (o *UpdateUserPayload) HasWrite() bool {
	if o != nil && !IsNil(o.Write) && !IsNil(o.Write) {
		return true
	}

	return false
}

// SetWrite gets a reference to the given bool and assigns it to the Write field.
func (o *UpdateUserPayload) SetWrite(v *bool) {
	o.Write = v
}

func (o UpdateUserPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Write) {
		toSerialize["write"] = o.Write
	}
	return toSerialize, nil
}

type NullableUpdateUserPayload struct {
	value *UpdateUserPayload
	isSet bool
}

func (v NullableUpdateUserPayload) Get() *UpdateUserPayload {
	return v.value
}

func (v *NullableUpdateUserPayload) Set(val *UpdateUserPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserPayload(val *UpdateUserPayload) *NullableUpdateUserPayload {
	return &NullableUpdateUserPayload{value: val, isSet: true}
}

func (v NullableUpdateUserPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
