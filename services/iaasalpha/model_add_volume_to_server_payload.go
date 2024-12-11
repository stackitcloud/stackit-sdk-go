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

// checks if the AddVolumeToServerPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddVolumeToServerPayload{}

// AddVolumeToServerPayload Object that represents a Volume attachment to a server.
type AddVolumeToServerPayload struct {
	// Delete the volume during the termination of the server. Defaults to false.
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty"`
	// Universally Unique Identifier (UUID).
	ServerId *string `json:"serverId,omitempty"`
	// Universally Unique Identifier (UUID).
	VolumeId *string `json:"volumeId,omitempty"`
}

// NewAddVolumeToServerPayload instantiates a new AddVolumeToServerPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVolumeToServerPayload() *AddVolumeToServerPayload {
	this := AddVolumeToServerPayload{}
	var deleteOnTermination bool = false
	this.DeleteOnTermination = &deleteOnTermination
	return &this
}

// NewAddVolumeToServerPayloadWithDefaults instantiates a new AddVolumeToServerPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVolumeToServerPayloadWithDefaults() *AddVolumeToServerPayload {
	this := AddVolumeToServerPayload{}
	var deleteOnTermination bool = false
	this.DeleteOnTermination = &deleteOnTermination
	return &this
}

// GetDeleteOnTermination returns the DeleteOnTermination field value if set, zero value otherwise.
func (o *AddVolumeToServerPayload) GetDeleteOnTermination() *bool {
	if o == nil || IsNil(o.DeleteOnTermination) {
		var ret *bool
		return ret
	}
	return o.DeleteOnTermination
}

// GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVolumeToServerPayload) GetDeleteOnTerminationOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteOnTermination) {
		return nil, false
	}
	return o.DeleteOnTermination, true
}

// HasDeleteOnTermination returns a boolean if a field has been set.
func (o *AddVolumeToServerPayload) HasDeleteOnTermination() bool {
	if o != nil && !IsNil(o.DeleteOnTermination) && !IsNil(o.DeleteOnTermination) {
		return true
	}

	return false
}

// SetDeleteOnTermination gets a reference to the given bool and assigns it to the DeleteOnTermination field.
func (o *AddVolumeToServerPayload) SetDeleteOnTermination(v *bool) {
	o.DeleteOnTermination = v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *AddVolumeToServerPayload) GetServerId() *string {
	if o == nil || IsNil(o.ServerId) {
		var ret *string
		return ret
	}
	return o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVolumeToServerPayload) GetServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *AddVolumeToServerPayload) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *AddVolumeToServerPayload) SetServerId(v *string) {
	o.ServerId = v
}

// GetVolumeId returns the VolumeId field value if set, zero value otherwise.
func (o *AddVolumeToServerPayload) GetVolumeId() *string {
	if o == nil || IsNil(o.VolumeId) {
		var ret *string
		return ret
	}
	return o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddVolumeToServerPayload) GetVolumeIdOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeId) {
		return nil, false
	}
	return o.VolumeId, true
}

// HasVolumeId returns a boolean if a field has been set.
func (o *AddVolumeToServerPayload) HasVolumeId() bool {
	if o != nil && !IsNil(o.VolumeId) && !IsNil(o.VolumeId) {
		return true
	}

	return false
}

// SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.
func (o *AddVolumeToServerPayload) SetVolumeId(v *string) {
	o.VolumeId = v
}

func (o AddVolumeToServerPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeleteOnTermination) {
		toSerialize["deleteOnTermination"] = o.DeleteOnTermination
	}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if !IsNil(o.VolumeId) {
		toSerialize["volumeId"] = o.VolumeId
	}
	return toSerialize, nil
}

type NullableAddVolumeToServerPayload struct {
	value *AddVolumeToServerPayload
	isSet bool
}

func (v NullableAddVolumeToServerPayload) Get() *AddVolumeToServerPayload {
	return v.value
}

func (v *NullableAddVolumeToServerPayload) Set(val *AddVolumeToServerPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVolumeToServerPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVolumeToServerPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVolumeToServerPayload(val *AddVolumeToServerPayload) *NullableAddVolumeToServerPayload {
	return &NullableAddVolumeToServerPayload{value: val, isSet: true}
}

func (v NullableAddVolumeToServerPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVolumeToServerPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
