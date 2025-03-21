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

// checks if the ListStoragesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListStoragesResponse{}

/*
	types and functions for storageClasses
*/

// isArray
type ListStoragesResponseGetStorageClassesAttributeType = *[]string
type ListStoragesResponseGetStorageClassesArgType = []string
type ListStoragesResponseGetStorageClassesRetType = []string

func getListStoragesResponseGetStorageClassesAttributeTypeOk(arg ListStoragesResponseGetStorageClassesAttributeType) (ret ListStoragesResponseGetStorageClassesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListStoragesResponseGetStorageClassesAttributeType(arg *ListStoragesResponseGetStorageClassesAttributeType, val ListStoragesResponseGetStorageClassesRetType) {
	*arg = &val
}

/*
	types and functions for storageRange
*/

// isModel
type ListStoragesResponseGetStorageRangeAttributeType = *StorageRange
type ListStoragesResponseGetStorageRangeArgType = StorageRange
type ListStoragesResponseGetStorageRangeRetType = StorageRange

func getListStoragesResponseGetStorageRangeAttributeTypeOk(arg ListStoragesResponseGetStorageRangeAttributeType) (ret ListStoragesResponseGetStorageRangeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListStoragesResponseGetStorageRangeAttributeType(arg *ListStoragesResponseGetStorageRangeAttributeType, val ListStoragesResponseGetStorageRangeRetType) {
	*arg = &val
}

// ListStoragesResponse struct for ListStoragesResponse
type ListStoragesResponse struct {
	StorageClasses ListStoragesResponseGetStorageClassesAttributeType `json:"storageClasses,omitempty"`
	StorageRange   ListStoragesResponseGetStorageRangeAttributeType   `json:"storageRange,omitempty"`
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
func (o *ListStoragesResponse) GetStorageClasses() (res ListStoragesResponseGetStorageClassesRetType) {
	res, _ = o.GetStorageClassesOk()
	return
}

// GetStorageClassesOk returns a tuple with the StorageClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStoragesResponse) GetStorageClassesOk() (ret ListStoragesResponseGetStorageClassesRetType, ok bool) {
	return getListStoragesResponseGetStorageClassesAttributeTypeOk(o.StorageClasses)
}

// HasStorageClasses returns a boolean if a field has been set.
func (o *ListStoragesResponse) HasStorageClasses() bool {
	_, ok := o.GetStorageClassesOk()
	return ok
}

// SetStorageClasses gets a reference to the given []string and assigns it to the StorageClasses field.
func (o *ListStoragesResponse) SetStorageClasses(v ListStoragesResponseGetStorageClassesRetType) {
	setListStoragesResponseGetStorageClassesAttributeType(&o.StorageClasses, v)
}

// GetStorageRange returns the StorageRange field value if set, zero value otherwise.
func (o *ListStoragesResponse) GetStorageRange() (res ListStoragesResponseGetStorageRangeRetType) {
	res, _ = o.GetStorageRangeOk()
	return
}

// GetStorageRangeOk returns a tuple with the StorageRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStoragesResponse) GetStorageRangeOk() (ret ListStoragesResponseGetStorageRangeRetType, ok bool) {
	return getListStoragesResponseGetStorageRangeAttributeTypeOk(o.StorageRange)
}

// HasStorageRange returns a boolean if a field has been set.
func (o *ListStoragesResponse) HasStorageRange() bool {
	_, ok := o.GetStorageRangeOk()
	return ok
}

// SetStorageRange gets a reference to the given StorageRange and assigns it to the StorageRange field.
func (o *ListStoragesResponse) SetStorageRange(v ListStoragesResponseGetStorageRangeRetType) {
	setListStoragesResponseGetStorageRangeAttributeType(&o.StorageRange, v)
}

func (o ListStoragesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListStoragesResponseGetStorageClassesAttributeTypeOk(o.StorageClasses); ok {
		toSerialize["StorageClasses"] = val
	}
	if val, ok := getListStoragesResponseGetStorageRangeAttributeTypeOk(o.StorageRange); ok {
		toSerialize["StorageRange"] = val
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
