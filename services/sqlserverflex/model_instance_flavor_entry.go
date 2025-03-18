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

/*
	types and functions for categories
*/

// isNotNullableString
type InstanceFlavorEntryGetCategoriesAttributeType = *string

func getInstanceFlavorEntryGetCategoriesAttributeTypeOk(arg InstanceFlavorEntryGetCategoriesAttributeType) (ret InstanceFlavorEntryGetCategoriesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceFlavorEntryGetCategoriesAttributeType(arg *InstanceFlavorEntryGetCategoriesAttributeType, val InstanceFlavorEntryGetCategoriesRetType) {
	*arg = &val
}

type InstanceFlavorEntryGetCategoriesArgType = string
type InstanceFlavorEntryGetCategoriesRetType = string

/*
	types and functions for cpu
*/

// isInteger
type InstanceFlavorEntryGetCpuAttributeType = *int64
type InstanceFlavorEntryGetCpuArgType = int64
type InstanceFlavorEntryGetCpuRetType = int64

func getInstanceFlavorEntryGetCpuAttributeTypeOk(arg InstanceFlavorEntryGetCpuAttributeType) (ret InstanceFlavorEntryGetCpuRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceFlavorEntryGetCpuAttributeType(arg *InstanceFlavorEntryGetCpuAttributeType, val InstanceFlavorEntryGetCpuRetType) {
	*arg = &val
}

/*
	types and functions for description
*/

// isNotNullableString
type InstanceFlavorEntryGetDescriptionAttributeType = *string

func getInstanceFlavorEntryGetDescriptionAttributeTypeOk(arg InstanceFlavorEntryGetDescriptionAttributeType) (ret InstanceFlavorEntryGetDescriptionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceFlavorEntryGetDescriptionAttributeType(arg *InstanceFlavorEntryGetDescriptionAttributeType, val InstanceFlavorEntryGetDescriptionRetType) {
	*arg = &val
}

type InstanceFlavorEntryGetDescriptionArgType = string
type InstanceFlavorEntryGetDescriptionRetType = string

/*
	types and functions for id
*/

// isNotNullableString
type InstanceFlavorEntryGetIdAttributeType = *string

func getInstanceFlavorEntryGetIdAttributeTypeOk(arg InstanceFlavorEntryGetIdAttributeType) (ret InstanceFlavorEntryGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceFlavorEntryGetIdAttributeType(arg *InstanceFlavorEntryGetIdAttributeType, val InstanceFlavorEntryGetIdRetType) {
	*arg = &val
}

type InstanceFlavorEntryGetIdArgType = string
type InstanceFlavorEntryGetIdRetType = string

/*
	types and functions for memory
*/

// isInteger
type InstanceFlavorEntryGetMemoryAttributeType = *int64
type InstanceFlavorEntryGetMemoryArgType = int64
type InstanceFlavorEntryGetMemoryRetType = int64

func getInstanceFlavorEntryGetMemoryAttributeTypeOk(arg InstanceFlavorEntryGetMemoryAttributeType) (ret InstanceFlavorEntryGetMemoryRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setInstanceFlavorEntryGetMemoryAttributeType(arg *InstanceFlavorEntryGetMemoryAttributeType, val InstanceFlavorEntryGetMemoryRetType) {
	*arg = &val
}

// InstanceFlavorEntry struct for InstanceFlavorEntry
type InstanceFlavorEntry struct {
	Categories InstanceFlavorEntryGetCategoriesAttributeType `json:"categories,omitempty"`
	// Can be cast to int32 without loss of precision.
	Cpu         InstanceFlavorEntryGetCpuAttributeType         `json:"cpu,omitempty"`
	Description InstanceFlavorEntryGetDescriptionAttributeType `json:"description,omitempty"`
	Id          InstanceFlavorEntryGetIdAttributeType          `json:"id,omitempty"`
	// Can be cast to int32 without loss of precision.
	Memory InstanceFlavorEntryGetMemoryAttributeType `json:"memory,omitempty"`
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
func (o *InstanceFlavorEntry) GetCategories() (res InstanceFlavorEntryGetCategoriesRetType) {
	res, _ = o.GetCategoriesOk()
	return
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceFlavorEntry) GetCategoriesOk() (ret InstanceFlavorEntryGetCategoriesRetType, ok bool) {
	return getInstanceFlavorEntryGetCategoriesAttributeTypeOk(o.Categories)
}

// HasCategories returns a boolean if a field has been set.
func (o *InstanceFlavorEntry) HasCategories() bool {
	_, ok := o.GetCategoriesOk()
	return ok
}

// SetCategories gets a reference to the given string and assigns it to the Categories field.
func (o *InstanceFlavorEntry) SetCategories(v InstanceFlavorEntryGetCategoriesRetType) {
	setInstanceFlavorEntryGetCategoriesAttributeType(&o.Categories, v)
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *InstanceFlavorEntry) GetCpu() (res InstanceFlavorEntryGetCpuRetType) {
	res, _ = o.GetCpuOk()
	return
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceFlavorEntry) GetCpuOk() (ret InstanceFlavorEntryGetCpuRetType, ok bool) {
	return getInstanceFlavorEntryGetCpuAttributeTypeOk(o.Cpu)
}

// HasCpu returns a boolean if a field has been set.
func (o *InstanceFlavorEntry) HasCpu() bool {
	_, ok := o.GetCpuOk()
	return ok
}

// SetCpu gets a reference to the given int64 and assigns it to the Cpu field.
func (o *InstanceFlavorEntry) SetCpu(v InstanceFlavorEntryGetCpuRetType) {
	setInstanceFlavorEntryGetCpuAttributeType(&o.Cpu, v)
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InstanceFlavorEntry) GetDescription() (res InstanceFlavorEntryGetDescriptionRetType) {
	res, _ = o.GetDescriptionOk()
	return
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceFlavorEntry) GetDescriptionOk() (ret InstanceFlavorEntryGetDescriptionRetType, ok bool) {
	return getInstanceFlavorEntryGetDescriptionAttributeTypeOk(o.Description)
}

// HasDescription returns a boolean if a field has been set.
func (o *InstanceFlavorEntry) HasDescription() bool {
	_, ok := o.GetDescriptionOk()
	return ok
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InstanceFlavorEntry) SetDescription(v InstanceFlavorEntryGetDescriptionRetType) {
	setInstanceFlavorEntryGetDescriptionAttributeType(&o.Description, v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstanceFlavorEntry) GetId() (res InstanceFlavorEntryGetIdRetType) {
	res, _ = o.GetIdOk()
	return
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceFlavorEntry) GetIdOk() (ret InstanceFlavorEntryGetIdRetType, ok bool) {
	return getInstanceFlavorEntryGetIdAttributeTypeOk(o.Id)
}

// HasId returns a boolean if a field has been set.
func (o *InstanceFlavorEntry) HasId() bool {
	_, ok := o.GetIdOk()
	return ok
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InstanceFlavorEntry) SetId(v InstanceFlavorEntryGetIdRetType) {
	setInstanceFlavorEntryGetIdAttributeType(&o.Id, v)
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *InstanceFlavorEntry) GetMemory() (res InstanceFlavorEntryGetMemoryRetType) {
	res, _ = o.GetMemoryOk()
	return
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceFlavorEntry) GetMemoryOk() (ret InstanceFlavorEntryGetMemoryRetType, ok bool) {
	return getInstanceFlavorEntryGetMemoryAttributeTypeOk(o.Memory)
}

// HasMemory returns a boolean if a field has been set.
func (o *InstanceFlavorEntry) HasMemory() bool {
	_, ok := o.GetMemoryOk()
	return ok
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *InstanceFlavorEntry) SetMemory(v InstanceFlavorEntryGetMemoryRetType) {
	setInstanceFlavorEntryGetMemoryAttributeType(&o.Memory, v)
}

func (o InstanceFlavorEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getInstanceFlavorEntryGetCategoriesAttributeTypeOk(o.Categories); ok {
		toSerialize["Categories"] = val
	}
	if val, ok := getInstanceFlavorEntryGetCpuAttributeTypeOk(o.Cpu); ok {
		toSerialize["Cpu"] = val
	}
	if val, ok := getInstanceFlavorEntryGetDescriptionAttributeTypeOk(o.Description); ok {
		toSerialize["Description"] = val
	}
	if val, ok := getInstanceFlavorEntryGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getInstanceFlavorEntryGetMemoryAttributeTypeOk(o.Memory); ok {
		toSerialize["Memory"] = val
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
