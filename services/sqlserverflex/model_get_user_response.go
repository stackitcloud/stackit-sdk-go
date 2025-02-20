/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
)

// checks if the GetUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserResponse{}

// GetUserResponse struct for GetUserResponse
type GetUserResponse struct {
	Item *UserResponseUser `json:"item,omitempty"`
}

// NewGetUserResponse instantiates a new GetUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserResponse() *GetUserResponse {
	this := GetUserResponse{}
	return &this
}

// NewGetUserResponseWithDefaults instantiates a new GetUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserResponseWithDefaults() *GetUserResponse {
	this := GetUserResponse{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *GetUserResponse) GetItem() *UserResponseUser {
	if o == nil || IsNil(o.Item) {
		var ret *UserResponseUser
		return ret
	}
	return o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserResponse) GetItemOk() (*UserResponseUser, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *GetUserResponse) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given UserResponseUser and assigns it to the Item field.
func (o *GetUserResponse) SetItem(v *UserResponseUser) {
	o.Item = v
}

func (o GetUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	return toSerialize, nil
}

type NullableGetUserResponse struct {
	value *GetUserResponse
	isSet bool
}

func (v NullableGetUserResponse) Get() *GetUserResponse {
	return v.value
}

func (v *NullableGetUserResponse) Set(val *GetUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserResponse(val *GetUserResponse) *NullableGetUserResponse {
	return &NullableGetUserResponse{value: val, isSet: true}
}

func (v NullableGetUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
