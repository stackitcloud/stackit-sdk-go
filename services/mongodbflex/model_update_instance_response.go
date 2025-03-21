/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

import (
	"encoding/json"
)

// checks if the UpdateInstanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInstanceResponse{}

/*
	types and functions for item
*/

// isModel
type UpdateInstanceResponseGetItemAttributeType = *Instance
type UpdateInstanceResponseGetItemArgType = Instance
type UpdateInstanceResponseGetItemRetType = Instance

func getUpdateInstanceResponseGetItemAttributeTypeOk(arg UpdateInstanceResponseGetItemAttributeType) (ret UpdateInstanceResponseGetItemRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateInstanceResponseGetItemAttributeType(arg *UpdateInstanceResponseGetItemAttributeType, val UpdateInstanceResponseGetItemRetType) {
	*arg = &val
}

// UpdateInstanceResponse struct for UpdateInstanceResponse
type UpdateInstanceResponse struct {
	Item UpdateInstanceResponseGetItemAttributeType `json:"item,omitempty"`
}

// NewUpdateInstanceResponse instantiates a new UpdateInstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInstanceResponse() *UpdateInstanceResponse {
	this := UpdateInstanceResponse{}
	return &this
}

// NewUpdateInstanceResponseWithDefaults instantiates a new UpdateInstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInstanceResponseWithDefaults() *UpdateInstanceResponse {
	this := UpdateInstanceResponse{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *UpdateInstanceResponse) GetItem() (res UpdateInstanceResponseGetItemRetType) {
	res, _ = o.GetItemOk()
	return
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInstanceResponse) GetItemOk() (ret UpdateInstanceResponseGetItemRetType, ok bool) {
	return getUpdateInstanceResponseGetItemAttributeTypeOk(o.Item)
}

// HasItem returns a boolean if a field has been set.
func (o *UpdateInstanceResponse) HasItem() bool {
	_, ok := o.GetItemOk()
	return ok
}

// SetItem gets a reference to the given Instance and assigns it to the Item field.
func (o *UpdateInstanceResponse) SetItem(v UpdateInstanceResponseGetItemRetType) {
	setUpdateInstanceResponseGetItemAttributeType(&o.Item, v)
}

func (o UpdateInstanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getUpdateInstanceResponseGetItemAttributeTypeOk(o.Item); ok {
		toSerialize["Item"] = val
	}
	return toSerialize, nil
}

type NullableUpdateInstanceResponse struct {
	value *UpdateInstanceResponse
	isSet bool
}

func (v NullableUpdateInstanceResponse) Get() *UpdateInstanceResponse {
	return v.value
}

func (v *NullableUpdateInstanceResponse) Set(val *UpdateInstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInstanceResponse(val *UpdateInstanceResponse) *NullableUpdateInstanceResponse {
	return &NullableUpdateInstanceResponse{value: val, isSet: true}
}

func (v NullableUpdateInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
