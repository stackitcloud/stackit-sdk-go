/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

import (
	"encoding/json"
)

// checks if the ListStoragesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListStoragesResponse{}

// ListStoragesResponse struct for ListStoragesResponse
type ListStoragesResponse struct {
	StorageClasses *[]string     `json:"storageClasses,omitempty"`
	StorageRange   *StorageRange `json:"storageRange,omitempty"`
}

// NewListStoragesResponse instantiates a new ListStoragesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStoragesResponse() *ListStoragesResponse {
	this := ListStoragesResponse{}
	return &this
}

// NewListStoragesResponseWithDefaults instantiates a new ListStoragesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStoragesResponseWithDefaults() *ListStoragesResponse {
	this := ListStoragesResponse{}
	return &this
}

// GetStorageClasses returns the StorageClasses field value if set, zero value otherwise.
func (o *ListStoragesResponse) GetStorageClasses() *[]string {
	if o == nil || IsNil(o.StorageClasses) {
		var ret *[]string
		return ret
	}
	return o.StorageClasses
}

// GetStorageClassesOk returns a tuple with the StorageClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStoragesResponse) GetStorageClassesOk() (*[]string, bool) {
	if o == nil || IsNil(o.StorageClasses) {
		return nil, false
	}
	return o.StorageClasses, true
}

// HasStorageClasses returns a boolean if a field has been set.
func (o *ListStoragesResponse) HasStorageClasses() bool {
	if o != nil && !IsNil(o.StorageClasses) {
		return true
	}

	return false
}

// SetStorageClasses gets a reference to the given []string and assigns it to the StorageClasses field.
func (o *ListStoragesResponse) SetStorageClasses(v *[]string) {
	o.StorageClasses = v
}

// GetStorageRange returns the StorageRange field value if set, zero value otherwise.
func (o *ListStoragesResponse) GetStorageRange() *StorageRange {
	if o == nil || IsNil(o.StorageRange) {
		var ret *StorageRange
		return ret
	}
	return o.StorageRange
}

// GetStorageRangeOk returns a tuple with the StorageRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStoragesResponse) GetStorageRangeOk() (*StorageRange, bool) {
	if o == nil || IsNil(o.StorageRange) {
		return nil, false
	}
	return o.StorageRange, true
}

// HasStorageRange returns a boolean if a field has been set.
func (o *ListStoragesResponse) HasStorageRange() bool {
	if o != nil && !IsNil(o.StorageRange) {
		return true
	}

	return false
}

// SetStorageRange gets a reference to the given StorageRange and assigns it to the StorageRange field.
func (o *ListStoragesResponse) SetStorageRange(v *StorageRange) {
	o.StorageRange = v
}

func (o ListStoragesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StorageClasses) {
		toSerialize["storageClasses"] = o.StorageClasses
	}
	if !IsNil(o.StorageRange) {
		toSerialize["storageRange"] = o.StorageRange
	}
	return toSerialize, nil
}

type NullableListStoragesResponse struct {
	value *ListStoragesResponse
	isSet bool
}

func (v NullableListStoragesResponse) Get() *ListStoragesResponse {
	return v.value
}

func (v *NullableListStoragesResponse) Set(val *ListStoragesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListStoragesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListStoragesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStoragesResponse(val *ListStoragesResponse) *NullableListStoragesResponse {
	return &NullableListStoragesResponse{value: val, isSet: true}
}

func (v NullableListStoragesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStoragesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
