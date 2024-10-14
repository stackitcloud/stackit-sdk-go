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

// checks if the InstanceListInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceListInstance{}

// InstanceListInstance struct for InstanceListInstance
type InstanceListInstance struct {
	Id     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewInstanceListInstance instantiates a new InstanceListInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceListInstance() *InstanceListInstance {
	this := InstanceListInstance{}
	return &this
}

// NewInstanceListInstanceWithDefaults instantiates a new InstanceListInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceListInstanceWithDefaults() *InstanceListInstance {
	this := InstanceListInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstanceListInstance) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListInstance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InstanceListInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InstanceListInstance) SetId(v *string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceListInstance) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstanceListInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceListInstance) SetName(v *string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InstanceListInstance) GetStatus() *string {
	if o == nil || IsNil(o.Status) {
		var ret *string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceListInstance) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InstanceListInstance) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InstanceListInstance) SetStatus(v *string) {
	o.Status = v
}

func (o InstanceListInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableInstanceListInstance struct {
	value *InstanceListInstance
	isSet bool
}

func (v NullableInstanceListInstance) Get() *InstanceListInstance {
	return v.value
}

func (v *NullableInstanceListInstance) Set(val *InstanceListInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceListInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceListInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceListInstance(val *InstanceListInstance) *NullableInstanceListInstance {
	return &NullableInstanceListInstance{value: val, isSet: true}
}

func (v NullableInstanceListInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceListInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
