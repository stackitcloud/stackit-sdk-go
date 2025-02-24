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

// checks if the InstanceFlavorEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceFlavorEntry{}

// InstanceFlavorEntry struct for InstanceFlavorEntry
type InstanceFlavorEntry struct {
	Categories *string `json:"categories,omitempty"`
	// Can be cast to int32 without loss of precision.
	Cpu         *int64  `json:"cpu,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	// Can be cast to int32 without loss of precision.
	Memory *int64 `json:"memory,omitempty"`
}

// NewInstanceFlavorEntry instantiates a new InstanceFlavorEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceFlavorEntry() *InstanceFlavorEntry {
	this := InstanceFlavorEntry{}
	return &this
}

// NewInstanceFlavorEntryWithDefaults instantiates a new InstanceFlavorEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceFlavorEntryWithDefaults() *InstanceFlavorEntry {
	this := InstanceFlavorEntry{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *InstanceFlavorEntry) GetCategories() *string {
	if o == nil || IsNil(o.Categories) {
		var ret *string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceFlavorEntry) GetCategoriesOk() (*string, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *InstanceFlavorEntry) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given string and assigns it to the Categories field.
func (o *InstanceFlavorEntry) SetCategories(v *string) {
	o.Categories = v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *InstanceFlavorEntry) GetCpu() *int64 {
	if o == nil || IsNil(o.Cpu) {
		var ret *int64
		return ret
	}
	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceFlavorEntry) GetCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *InstanceFlavorEntry) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int64 and assigns it to the Cpu field.
func (o *InstanceFlavorEntry) SetCpu(v *int64) {
	o.Cpu = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InstanceFlavorEntry) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceFlavorEntry) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InstanceFlavorEntry) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InstanceFlavorEntry) SetDescription(v *string) {
	o.Description = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstanceFlavorEntry) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceFlavorEntry) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InstanceFlavorEntry) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InstanceFlavorEntry) SetId(v *string) {
	o.Id = v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *InstanceFlavorEntry) GetMemory() *int64 {
	if o == nil || IsNil(o.Memory) {
		var ret *int64
		return ret
	}
	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceFlavorEntry) GetMemoryOk() (*int64, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *InstanceFlavorEntry) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *InstanceFlavorEntry) SetMemory(v *int64) {
	o.Memory = v
}

func (o InstanceFlavorEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return toSerialize, nil
}

type NullableInstanceFlavorEntry struct {
	value *InstanceFlavorEntry
	isSet bool
}

func (v NullableInstanceFlavorEntry) Get() *InstanceFlavorEntry {
	return v.value
}

func (v *NullableInstanceFlavorEntry) Set(val *InstanceFlavorEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceFlavorEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceFlavorEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceFlavorEntry(val *InstanceFlavorEntry) *NullableInstanceFlavorEntry {
	return &NullableInstanceFlavorEntry{value: val, isSet: true}
}

func (v NullableInstanceFlavorEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceFlavorEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
