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

// checks if the Storage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Storage{}

// Storage struct for Storage
type Storage struct {
	Class *string `json:"class,omitempty"`
	Size  *int64  `json:"size,omitempty"`
}

// NewStorage instantiates a new Storage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorage() *Storage {
	this := Storage{}
	return &this
}

// NewStorageWithDefaults instantiates a new Storage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageWithDefaults() *Storage {
	this := Storage{}
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *Storage) GetClass() *string {
	if o == nil || IsNil(o.Class) {
		var ret *string
		return ret
	}
	return o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *Storage) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *Storage) SetClass(v *string) {
	o.Class = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Storage) GetSize() *int64 {
	if o == nil || IsNil(o.Size) {
		var ret *int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Storage) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *Storage) SetSize(v *int64) {
	o.Size = v
}

func (o Storage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Class) {
		toSerialize["class"] = o.Class
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableStorage struct {
	value *Storage
	isSet bool
}

func (v NullableStorage) Get() *Storage {
	return v.value
}

func (v *NullableStorage) Set(val *Storage) {
	v.value = val
	v.isSet = true
}

func (v NullableStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorage(val *Storage) *NullableStorage {
	return &NullableStorage{value: val, isSet: true}
}

func (v NullableStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
