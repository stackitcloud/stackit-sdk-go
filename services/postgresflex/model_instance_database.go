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

// checks if the InstanceDatabase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceDatabase{}

// InstanceDatabase struct for InstanceDatabase
type InstanceDatabase struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Database specific options
	Options *map[string]interface{} `json:"options,omitempty"`
}

// NewInstanceDatabase instantiates a new InstanceDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceDatabase() *InstanceDatabase {
	this := InstanceDatabase{}
	return &this
}

// NewInstanceDatabaseWithDefaults instantiates a new InstanceDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceDatabaseWithDefaults() *InstanceDatabase {
	this := InstanceDatabase{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstanceDatabase) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceDatabase) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InstanceDatabase) HasId() bool {
	if o != nil && !IsNil(o.Id) && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InstanceDatabase) SetId(v *string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceDatabase) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceDatabase) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstanceDatabase) HasName() bool {
	if o != nil && !IsNil(o.Name) && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceDatabase) SetName(v *string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *InstanceDatabase) GetOptions() *map[string]interface{} {
	if o == nil || IsNil(o.Options) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceDatabase) GetOptionsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return &map[string]interface{}{}, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *InstanceDatabase) HasOptions() bool {
	if o != nil && !IsNil(o.Options) && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *InstanceDatabase) SetOptions(v *map[string]interface{}) {
	o.Options = v
}

func (o InstanceDatabase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableInstanceDatabase struct {
	value *InstanceDatabase
	isSet bool
}

func (v NullableInstanceDatabase) Get() *InstanceDatabase {
	return v.value
}

func (v *NullableInstanceDatabase) Set(val *InstanceDatabase) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceDatabase(val *InstanceDatabase) *NullableInstanceDatabase {
	return &NullableInstanceDatabase{value: val, isSet: true}
}

func (v NullableInstanceDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
