/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
)

// checks if the InstanceDocumentationStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceDocumentationStorage{}

// InstanceDocumentationStorage struct for InstanceDocumentationStorage
type InstanceDocumentationStorage struct {
	// Class of the instance.
	Class *string `json:"class,omitempty"`
	// Size of the instance storage in GB
	Size *int64 `json:"size,omitempty"`
}

// NewInstanceDocumentationStorage instantiates a new InstanceDocumentationStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceDocumentationStorage() *InstanceDocumentationStorage {
	this := InstanceDocumentationStorage{}
	var class string = "premium-perf12-stackit"
	this.Class = &class
	return &this
}

// NewInstanceDocumentationStorageWithDefaults instantiates a new InstanceDocumentationStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceDocumentationStorageWithDefaults() *InstanceDocumentationStorage {
	this := InstanceDocumentationStorage{}
	var class string = "premium-perf12-stackit"
	this.Class = &class
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *InstanceDocumentationStorage) GetClass() *string {
	if o == nil || IsNil(o.Class) {
		var ret *string
		return ret
	}
	return o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceDocumentationStorage) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *InstanceDocumentationStorage) HasClass() bool {
	if o != nil && !IsNil(o.Class) && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *InstanceDocumentationStorage) SetClass(v *string) {
	o.Class = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *InstanceDocumentationStorage) GetSize() *int64 {
	if o == nil || IsNil(o.Size) {
		var ret *int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceDocumentationStorage) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *InstanceDocumentationStorage) HasSize() bool {
	if o != nil && !IsNil(o.Size) && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *InstanceDocumentationStorage) SetSize(v *int64) {
	o.Size = v
}

func (o InstanceDocumentationStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Class) {
		toSerialize["class"] = o.Class
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableInstanceDocumentationStorage struct {
	value *InstanceDocumentationStorage
	isSet bool
}

func (v NullableInstanceDocumentationStorage) Get() *InstanceDocumentationStorage {
	return v.value
}

func (v *NullableInstanceDocumentationStorage) Set(val *InstanceDocumentationStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceDocumentationStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceDocumentationStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceDocumentationStorage(val *InstanceDocumentationStorage) *NullableInstanceDocumentationStorage {
	return &NullableInstanceDocumentationStorage{value: val, isSet: true}
}

func (v NullableInstanceDocumentationStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceDocumentationStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
