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

// checks if the PartialUpdateInstanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartialUpdateInstanceResponse{}

/*
	types and functions for item
*/

// isModel
type PartialUpdateInstanceResponseGetItemAttributeType = *Instance
type PartialUpdateInstanceResponseGetItemArgType = Instance
type PartialUpdateInstanceResponseGetItemRetType = Instance

func getPartialUpdateInstanceResponseGetItemAttributeTypeOk(arg PartialUpdateInstanceResponseGetItemAttributeType) (ret PartialUpdateInstanceResponseGetItemRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPartialUpdateInstanceResponseGetItemAttributeType(arg *PartialUpdateInstanceResponseGetItemAttributeType, val PartialUpdateInstanceResponseGetItemRetType) {
	*arg = &val
}

// PartialUpdateInstanceResponse struct for PartialUpdateInstanceResponse
type PartialUpdateInstanceResponse struct {
	Item PartialUpdateInstanceResponseGetItemAttributeType `json:"item,omitempty"`
}

// NewPartialUpdateInstanceResponse instantiates a new PartialUpdateInstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialUpdateInstanceResponse() *PartialUpdateInstanceResponse {
	this := PartialUpdateInstanceResponse{}
	return &this
}

// NewPartialUpdateInstanceResponseWithDefaults instantiates a new PartialUpdateInstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialUpdateInstanceResponseWithDefaults() *PartialUpdateInstanceResponse {
	this := PartialUpdateInstanceResponse{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *PartialUpdateInstanceResponse) GetItem() (res PartialUpdateInstanceResponseGetItemRetType) {
	res, _ = o.GetItemOk()
	return
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateInstanceResponse) GetItemOk() (ret PartialUpdateInstanceResponseGetItemRetType, ok bool) {
	return getPartialUpdateInstanceResponseGetItemAttributeTypeOk(o.Item)
}

// HasItem returns a boolean if a field has been set.
func (o *PartialUpdateInstanceResponse) HasItem() bool {
	_, ok := o.GetItemOk()
	return ok
}

// SetItem gets a reference to the given Instance and assigns it to the Item field.
func (o *PartialUpdateInstanceResponse) SetItem(v PartialUpdateInstanceResponseGetItemRetType) {
	setPartialUpdateInstanceResponseGetItemAttributeType(&o.Item, v)
}

func (o PartialUpdateInstanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getPartialUpdateInstanceResponseGetItemAttributeTypeOk(o.Item); ok {
		toSerialize["Item"] = val
	}
	return toSerialize, nil
}

type NullablePartialUpdateInstanceResponse struct {
	value *PartialUpdateInstanceResponse
	isSet bool
}

func (v NullablePartialUpdateInstanceResponse) Get() *PartialUpdateInstanceResponse {
	return v.value
}

func (v *NullablePartialUpdateInstanceResponse) Set(val *PartialUpdateInstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialUpdateInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialUpdateInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialUpdateInstanceResponse(val *PartialUpdateInstanceResponse) *NullablePartialUpdateInstanceResponse {
	return &NullablePartialUpdateInstanceResponse{value: val, isSet: true}
}

func (v NullablePartialUpdateInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialUpdateInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
