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

// checks if the CreateDatabasePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDatabasePayload{}

/*
	types and functions for name
*/

// isNotNullableString
type CreateDatabasePayloadGetNameAttributeType = *string

func getCreateDatabasePayloadGetNameAttributeTypeOk(arg CreateDatabasePayloadGetNameAttributeType) (ret CreateDatabasePayloadGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateDatabasePayloadGetNameAttributeType(arg *CreateDatabasePayloadGetNameAttributeType, val CreateDatabasePayloadGetNameRetType) {
	*arg = &val
}

type CreateDatabasePayloadGetNameArgType = string
type CreateDatabasePayloadGetNameRetType = string

/*
	types and functions for options
*/

// isModel
type CreateDatabasePayloadGetOptionsAttributeType = *DatabaseDocumentationCreateDatabaseRequestOptions
type CreateDatabasePayloadGetOptionsArgType = DatabaseDocumentationCreateDatabaseRequestOptions
type CreateDatabasePayloadGetOptionsRetType = DatabaseDocumentationCreateDatabaseRequestOptions

func getCreateDatabasePayloadGetOptionsAttributeTypeOk(arg CreateDatabasePayloadGetOptionsAttributeType) (ret CreateDatabasePayloadGetOptionsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateDatabasePayloadGetOptionsAttributeType(arg *CreateDatabasePayloadGetOptionsAttributeType, val CreateDatabasePayloadGetOptionsRetType) {
	*arg = &val
}

// CreateDatabasePayload struct for CreateDatabasePayload
type CreateDatabasePayload struct {
	// REQUIRED
	Name CreateDatabasePayloadGetNameAttributeType `json:"name"`
	// REQUIRED
	Options CreateDatabasePayloadGetOptionsAttributeType `json:"options"`
}

type _CreateDatabasePayload CreateDatabasePayload

// NewCreateDatabasePayload instantiates a new CreateDatabasePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDatabasePayload(name CreateDatabasePayloadGetNameArgType, options CreateDatabasePayloadGetOptionsArgType) *CreateDatabasePayload {
	this := CreateDatabasePayload{}
	setCreateDatabasePayloadGetNameAttributeType(&this.Name, name)
	setCreateDatabasePayloadGetOptionsAttributeType(&this.Options, options)
	return &this
}

// NewCreateDatabasePayloadWithDefaults instantiates a new CreateDatabasePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDatabasePayloadWithDefaults() *CreateDatabasePayload {
	this := CreateDatabasePayload{}
	return &this
}

// GetName returns the Name field value
func (o *CreateDatabasePayload) GetName() (ret CreateDatabasePayloadGetNameRetType) {
	ret, _ = o.GetNameOk()
	return ret
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDatabasePayload) GetNameOk() (ret CreateDatabasePayloadGetNameRetType, ok bool) {
	return getCreateDatabasePayloadGetNameAttributeTypeOk(o.Name)
}

// SetName sets field value
func (o *CreateDatabasePayload) SetName(v CreateDatabasePayloadGetNameRetType) {
	setCreateDatabasePayloadGetNameAttributeType(&o.Name, v)
}

// GetOptions returns the Options field value
func (o *CreateDatabasePayload) GetOptions() (ret CreateDatabasePayloadGetOptionsRetType) {
	ret, _ = o.GetOptionsOk()
	return ret
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *CreateDatabasePayload) GetOptionsOk() (ret CreateDatabasePayloadGetOptionsRetType, ok bool) {
	return getCreateDatabasePayloadGetOptionsAttributeTypeOk(o.Options)
}

// SetOptions sets field value
func (o *CreateDatabasePayload) SetOptions(v CreateDatabasePayloadGetOptionsRetType) {
	setCreateDatabasePayloadGetOptionsAttributeType(&o.Options, v)
}

func (o CreateDatabasePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateDatabasePayloadGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getCreateDatabasePayloadGetOptionsAttributeTypeOk(o.Options); ok {
		toSerialize["Options"] = val
	}
	return toSerialize, nil
}

type NullableCreateDatabasePayload struct {
	value *CreateDatabasePayload
	isSet bool
}

func (v NullableCreateDatabasePayload) Get() *CreateDatabasePayload {
	return v.value
}

func (v *NullableCreateDatabasePayload) Set(val *CreateDatabasePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDatabasePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDatabasePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDatabasePayload(val *CreateDatabasePayload) *NullableCreateDatabasePayload {
	return &NullableCreateDatabasePayload{value: val, isSet: true}
}

func (v NullableCreateDatabasePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDatabasePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
