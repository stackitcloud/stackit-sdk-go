/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

import (
	"encoding/json"
)

// checks if the InstanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceResponse{}

// InstanceResponse struct for InstanceResponse
type InstanceResponse struct {
	Item *Instance `json:"item,omitempty"`
}

// NewInstanceResponse instantiates a new InstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceResponse() *InstanceResponse {
	this := InstanceResponse{}
	return &this
}

// NewInstanceResponseWithDefaults instantiates a new InstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceResponseWithDefaults() *InstanceResponse {
	this := InstanceResponse{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *InstanceResponse) GetItem() *Instance {
	if o == nil || IsNil(o.Item) {
		var ret *Instance
		return ret
	}
	return o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResponse) GetItemOk() (*Instance, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *InstanceResponse) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given Instance and assigns it to the Item field.
func (o *InstanceResponse) SetItem(v *Instance) {
	o.Item = v
}

func (o InstanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	return toSerialize, nil
}

type NullableInstanceResponse struct {
	value *InstanceResponse
	isSet bool
}

func (v NullableInstanceResponse) Get() *InstanceResponse {
	return v.value
}

func (v *NullableInstanceResponse) Set(val *InstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceResponse(val *InstanceResponse) *NullableInstanceResponse {
	return &NullableInstanceResponse{value: val, isSet: true}
}

func (v NullableInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
