/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

import (
	"encoding/json"
)

// checks if the UpdateUserPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserPayload{}

/*
	types and functions for database
*/

// isNotNullableString
type UpdateUserPayloadGetDatabaseAttributeType = *string

func getUpdateUserPayloadGetDatabaseAttributeTypeOk(arg UpdateUserPayloadGetDatabaseAttributeType) (ret UpdateUserPayloadGetDatabaseRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateUserPayloadGetDatabaseAttributeType(arg *UpdateUserPayloadGetDatabaseAttributeType, val UpdateUserPayloadGetDatabaseRetType) {
	*arg = &val
}

type UpdateUserPayloadGetDatabaseArgType = string
type UpdateUserPayloadGetDatabaseRetType = string

/*
	types and functions for roles
*/

// isArray
type UpdateUserPayloadGetRolesAttributeType = *[]string
type UpdateUserPayloadGetRolesArgType = []string
type UpdateUserPayloadGetRolesRetType = []string

func getUpdateUserPayloadGetRolesAttributeTypeOk(arg UpdateUserPayloadGetRolesAttributeType) (ret UpdateUserPayloadGetRolesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateUserPayloadGetRolesAttributeType(arg *UpdateUserPayloadGetRolesAttributeType, val UpdateUserPayloadGetRolesRetType) {
	*arg = &val
}

// UpdateUserPayload struct for UpdateUserPayload
type UpdateUserPayload struct {
	// REQUIRED
	Database UpdateUserPayloadGetDatabaseAttributeType `json:"database" required:"true"`
	// REQUIRED
	Roles UpdateUserPayloadGetRolesAttributeType `json:"roles" required:"true"`
}

type _UpdateUserPayload UpdateUserPayload

// NewUpdateUserPayload instantiates a new UpdateUserPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserPayload(database UpdateUserPayloadGetDatabaseArgType, roles UpdateUserPayloadGetRolesArgType) *UpdateUserPayload {
	this := UpdateUserPayload{}
	setUpdateUserPayloadGetDatabaseAttributeType(&this.Database, database)
	setUpdateUserPayloadGetRolesAttributeType(&this.Roles, roles)
	return &this
}

// NewUpdateUserPayloadWithDefaults instantiates a new UpdateUserPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserPayloadWithDefaults() *UpdateUserPayload {
	this := UpdateUserPayload{}
	return &this
}

// GetDatabase returns the Database field value
func (o *UpdateUserPayload) GetDatabase() (ret UpdateUserPayloadGetDatabaseRetType) {
	ret, _ = o.GetDatabaseOk()
	return ret
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *UpdateUserPayload) GetDatabaseOk() (ret UpdateUserPayloadGetDatabaseRetType, ok bool) {
	return getUpdateUserPayloadGetDatabaseAttributeTypeOk(o.Database)
}

// SetDatabase sets field value
func (o *UpdateUserPayload) SetDatabase(v UpdateUserPayloadGetDatabaseRetType) {
	setUpdateUserPayloadGetDatabaseAttributeType(&o.Database, v)
}

// GetRoles returns the Roles field value
func (o *UpdateUserPayload) GetRoles() (ret UpdateUserPayloadGetRolesRetType) {
	ret, _ = o.GetRolesOk()
	return ret
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *UpdateUserPayload) GetRolesOk() (ret UpdateUserPayloadGetRolesRetType, ok bool) {
	return getUpdateUserPayloadGetRolesAttributeTypeOk(o.Roles)
}

// SetRoles sets field value
func (o *UpdateUserPayload) SetRoles(v UpdateUserPayloadGetRolesRetType) {
	setUpdateUserPayloadGetRolesAttributeType(&o.Roles, v)
}

func (o UpdateUserPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getUpdateUserPayloadGetDatabaseAttributeTypeOk(o.Database); ok {
		toSerialize["Database"] = val
	}
	if val, ok := getUpdateUserPayloadGetRolesAttributeTypeOk(o.Roles); ok {
		toSerialize["Roles"] = val
	}
	return toSerialize, nil
}

type NullableUpdateUserPayload struct {
	value *UpdateUserPayload
	isSet bool
}

func (v NullableUpdateUserPayload) Get() *UpdateUserPayload {
	return v.value
}

func (v *NullableUpdateUserPayload) Set(val *UpdateUserPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserPayload(val *UpdateUserPayload) *NullableUpdateUserPayload {
	return &NullableUpdateUserPayload{value: val, isSet: true}
}

func (v NullableUpdateUserPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
