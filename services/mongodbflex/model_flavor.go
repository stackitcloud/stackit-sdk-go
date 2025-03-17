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

// checks if the Flavor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Flavor{}

/*
	types and functions for cpu
*/

// isInteger
type FlavorGetCpuAttributeType = *int64
type FlavorGetCpuArgType = int64
type FlavorGetCpuRetType = int64

func getFlavorGetCpuAttributeTypeOk(arg FlavorGetCpuAttributeType) (ret FlavorGetCpuRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setFlavorGetCpuAttributeType(arg *FlavorGetCpuAttributeType, val FlavorGetCpuRetType) {
	*arg = &val
}

/*
	types and functions for description
*/

// isNotNullableString
type FlavorGetDescriptionAttributeType = *string

func getFlavorGetDescriptionAttributeTypeOk(arg FlavorGetDescriptionAttributeType) (ret FlavorGetDescriptionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setFlavorGetDescriptionAttributeType(arg *FlavorGetDescriptionAttributeType, val FlavorGetDescriptionRetType) {
	*arg = &val
}

type FlavorGetDescriptionArgType = string
type FlavorGetDescriptionRetType = string

/*
	types and functions for id
*/

// isNotNullableString
type FlavorGetIdAttributeType = *string

func getFlavorGetIdAttributeTypeOk(arg FlavorGetIdAttributeType) (ret FlavorGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setFlavorGetIdAttributeType(arg *FlavorGetIdAttributeType, val FlavorGetIdRetType) {
	*arg = &val
}

type FlavorGetIdArgType = string
type FlavorGetIdRetType = string

/*
	types and functions for memory
*/

// isInteger
type FlavorGetMemoryAttributeType = *int64
type FlavorGetMemoryArgType = int64
type FlavorGetMemoryRetType = int64

func getFlavorGetMemoryAttributeTypeOk(arg FlavorGetMemoryAttributeType) (ret FlavorGetMemoryRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setFlavorGetMemoryAttributeType(arg *FlavorGetMemoryAttributeType, val FlavorGetMemoryRetType) {
	*arg = &val
}

// Flavor struct for Flavor
type Flavor struct {
	Cpu         FlavorGetCpuAttributeType         `json:"cpu,omitempty"`
	Description FlavorGetDescriptionAttributeType `json:"description,omitempty"`
	Id          FlavorGetIdAttributeType          `json:"id,omitempty"`
	Memory      FlavorGetMemoryAttributeType      `json:"memory,omitempty"`
}

// NewFlavor instantiates a new Flavor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlavor() *Flavor {
	this := Flavor{}
	return &this
}

// NewFlavorWithDefaults instantiates a new Flavor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlavorWithDefaults() *Flavor {
	this := Flavor{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Flavor) GetCpu() (res FlavorGetCpuRetType) {
	res, _ = o.GetCpuOk()
	return
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flavor) GetCpuOk() (ret FlavorGetCpuRetType, ok bool) {
	return getFlavorGetCpuAttributeTypeOk(o.Cpu)
}

// HasCpu returns a boolean if a field has been set.
func (o *Flavor) HasCpu() bool {
	_, ok := o.GetCpuOk()
	return ok
}

// SetCpu gets a reference to the given int64 and assigns it to the Cpu field.
func (o *Flavor) SetCpu(v FlavorGetCpuRetType) {
	setFlavorGetCpuAttributeType(&o.Cpu, v)
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Flavor) GetDescription() (res FlavorGetDescriptionRetType) {
	res, _ = o.GetDescriptionOk()
	return
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flavor) GetDescriptionOk() (ret FlavorGetDescriptionRetType, ok bool) {
	return getFlavorGetDescriptionAttributeTypeOk(o.Description)
}

// HasDescription returns a boolean if a field has been set.
func (o *Flavor) HasDescription() bool {
	_, ok := o.GetDescriptionOk()
	return ok
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Flavor) SetDescription(v FlavorGetDescriptionRetType) {
	setFlavorGetDescriptionAttributeType(&o.Description, v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Flavor) GetId() (res FlavorGetIdRetType) {
	res, _ = o.GetIdOk()
	return
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flavor) GetIdOk() (ret FlavorGetIdRetType, ok bool) {
	return getFlavorGetIdAttributeTypeOk(o.Id)
}

// HasId returns a boolean if a field has been set.
func (o *Flavor) HasId() bool {
	_, ok := o.GetIdOk()
	return ok
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Flavor) SetId(v FlavorGetIdRetType) {
	setFlavorGetIdAttributeType(&o.Id, v)
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Flavor) GetMemory() (res FlavorGetMemoryRetType) {
	res, _ = o.GetMemoryOk()
	return
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flavor) GetMemoryOk() (ret FlavorGetMemoryRetType, ok bool) {
	return getFlavorGetMemoryAttributeTypeOk(o.Memory)
}

// HasMemory returns a boolean if a field has been set.
func (o *Flavor) HasMemory() bool {
	_, ok := o.GetMemoryOk()
	return ok
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *Flavor) SetMemory(v FlavorGetMemoryRetType) {
	setFlavorGetMemoryAttributeType(&o.Memory, v)
}

func (o Flavor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getFlavorGetCpuAttributeTypeOk(o.Cpu); ok {
		toSerialize["Cpu"] = val
	}
	if val, ok := getFlavorGetDescriptionAttributeTypeOk(o.Description); ok {
		toSerialize["Description"] = val
	}
	if val, ok := getFlavorGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getFlavorGetMemoryAttributeTypeOk(o.Memory); ok {
		toSerialize["Memory"] = val
	}
	return toSerialize, nil
}

type NullableFlavor struct {
	value *Flavor
	isSet bool
}

func (v NullableFlavor) Get() *Flavor {
	return v.value
}

func (v *NullableFlavor) Set(val *Flavor) {
	v.value = val
	v.isSet = true
}

func (v NullableFlavor) IsSet() bool {
	return v.isSet
}

func (v *NullableFlavor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlavor(val *Flavor) *NullableFlavor {
	return &NullableFlavor{value: val, isSet: true}
}

func (v NullableFlavor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlavor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
