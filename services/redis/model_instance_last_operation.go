/*
STACKIT Redis API

The STACKIT Redis API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redis

import (
	"encoding/json"
)

// checks if the InstanceLastOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceLastOperation{}

// InstanceLastOperation struct for InstanceLastOperation
type InstanceLastOperation struct {
	// REQUIRED
	Description *string `json:"description"`
	// REQUIRED
	State *string `json:"state"`
	// REQUIRED
	Type *string `json:"type"`
}

type _InstanceLastOperation InstanceLastOperation

// NewInstanceLastOperation instantiates a new InstanceLastOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceLastOperation(description *string, state *string, type_ *string) *InstanceLastOperation {
	this := InstanceLastOperation{}
	this.Description = description
	this.State = state
	this.Type = type_
	return &this
}

// NewInstanceLastOperationWithDefaults instantiates a new InstanceLastOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceLastOperationWithDefaults() *InstanceLastOperation {
	this := InstanceLastOperation{}
	return &this
}

// GetDescription returns the Description field value
func (o *InstanceLastOperation) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InstanceLastOperation) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description, true
}

// SetDescription sets field value
func (o *InstanceLastOperation) SetDescription(v *string) {
	o.Description = v
}

// GetState returns the State field value
func (o *InstanceLastOperation) GetState() *string {
	if o == nil || IsNil(o.State) {
		var ret *string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *InstanceLastOperation) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State, true
}

// SetState sets field value
func (o *InstanceLastOperation) SetState(v *string) {
	o.State = v
}

// GetType returns the Type field value
func (o *InstanceLastOperation) GetType() *string {
	if o == nil || IsNil(o.Type) {
		var ret *string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InstanceLastOperation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *InstanceLastOperation) SetType(v *string) {
	o.Type = v
}

func (o InstanceLastOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["state"] = o.State
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableInstanceLastOperation struct {
	value *InstanceLastOperation
	isSet bool
}

func (v NullableInstanceLastOperation) Get() *InstanceLastOperation {
	return v.value
}

func (v *NullableInstanceLastOperation) Set(val *InstanceLastOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceLastOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceLastOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceLastOperation(val *InstanceLastOperation) *NullableInstanceLastOperation {
	return &NullableInstanceLastOperation{value: val, isSet: true}
}

func (v NullableInstanceLastOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceLastOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
