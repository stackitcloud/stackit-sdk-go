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

/*
	types and functions for description
*/

// isNotNullableString
type InstanceLastOperationGetDescriptionAttributeType = *string

func getInstanceLastOperationGetDescriptionAttributeTypeOk(arg InstanceLastOperationGetDescriptionAttributeType) (ret InstanceLastOperationGetDescriptionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceLastOperationGetDescriptionAttributeType(arg *InstanceLastOperationGetDescriptionAttributeType, val InstanceLastOperationGetDescriptionRetType) {
	*arg = &val
}

type InstanceLastOperationGetDescriptionArgType = string
type InstanceLastOperationGetDescriptionRetType = string

/*
	types and functions for state
*/

// isEnumRef
type InstanceLastOperationGetStateAttributeType = *string
type InstanceLastOperationGetStateArgType = string
type InstanceLastOperationGetStateRetType = string

func getInstanceLastOperationGetStateAttributeTypeOk(arg InstanceLastOperationGetStateAttributeType) (ret InstanceLastOperationGetStateRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceLastOperationGetStateAttributeType(arg *InstanceLastOperationGetStateAttributeType, val InstanceLastOperationGetStateRetType) {
	*arg = &val
}

/*
	types and functions for type
*/

// isEnumRef
type InstanceLastOperationGetTypeAttributeType = *string
type InstanceLastOperationGetTypeArgType = string
type InstanceLastOperationGetTypeRetType = string

func getInstanceLastOperationGetTypeAttributeTypeOk(arg InstanceLastOperationGetTypeAttributeType) (ret InstanceLastOperationGetTypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceLastOperationGetTypeAttributeType(arg *InstanceLastOperationGetTypeAttributeType, val InstanceLastOperationGetTypeRetType) {
	*arg = &val
}

// InstanceLastOperation struct for InstanceLastOperation
type InstanceLastOperation struct {
	// REQUIRED
	Description InstanceLastOperationGetDescriptionAttributeType `json:"description"`
	// REQUIRED
	State InstanceLastOperationGetStateAttributeType `json:"state"`
	// REQUIRED
	Type InstanceLastOperationGetTypeAttributeType `json:"type"`
}

type _InstanceLastOperation InstanceLastOperation

// NewInstanceLastOperation instantiates a new InstanceLastOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceLastOperation(description InstanceLastOperationGetDescriptionArgType, state InstanceLastOperationGetStateArgType, type_ InstanceLastOperationGetTypeArgType) *InstanceLastOperation {
	this := InstanceLastOperation{}
	setInstanceLastOperationGetDescriptionAttributeType(&this.Description, description)
	setInstanceLastOperationGetStateAttributeType(&this.State, state)
	setInstanceLastOperationGetTypeAttributeType(&this.Type, type_)
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
func (o *InstanceLastOperation) GetDescription() (ret InstanceLastOperationGetDescriptionRetType) {
	ret, _ = o.GetDescriptionOk()
	return ret
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InstanceLastOperation) GetDescriptionOk() (ret InstanceLastOperationGetDescriptionRetType, ok bool) {
	return getInstanceLastOperationGetDescriptionAttributeTypeOk(o.Description)
}

// SetDescription sets field value
func (o *InstanceLastOperation) SetDescription(v InstanceLastOperationGetDescriptionRetType) {
	setInstanceLastOperationGetDescriptionAttributeType(&o.Description, v)
}

// GetState returns the State field value
func (o *InstanceLastOperation) GetState() (ret InstanceLastOperationGetStateRetType) {
	ret, _ = o.GetStateOk()
	return ret
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *InstanceLastOperation) GetStateOk() (ret InstanceLastOperationGetStateRetType, ok bool) {
	return getInstanceLastOperationGetStateAttributeTypeOk(o.State)
}

// SetState sets field value
func (o *InstanceLastOperation) SetState(v InstanceLastOperationGetStateRetType) {
	setInstanceLastOperationGetStateAttributeType(&o.State, v)
}

// GetType returns the Type field value
func (o *InstanceLastOperation) GetType() (ret InstanceLastOperationGetTypeRetType) {
	ret, _ = o.GetTypeOk()
	return ret
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InstanceLastOperation) GetTypeOk() (ret InstanceLastOperationGetTypeRetType, ok bool) {
	return getInstanceLastOperationGetTypeAttributeTypeOk(o.Type)
}

// SetType sets field value
func (o *InstanceLastOperation) SetType(v InstanceLastOperationGetTypeRetType) {
	setInstanceLastOperationGetTypeAttributeType(&o.Type, v)
}

func (o InstanceLastOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getInstanceLastOperationGetDescriptionAttributeTypeOk(o.Description); ok {
		toSerialize["Description"] = val
	}
	if val, ok := getInstanceLastOperationGetStateAttributeTypeOk(o.State); ok {
		toSerialize["State"] = val
	}
	if val, ok := getInstanceLastOperationGetTypeAttributeTypeOk(o.Type); ok {
		toSerialize["Type"] = val
	}
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
