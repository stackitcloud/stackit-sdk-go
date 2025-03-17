/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorization

import (
	"encoding/json"
)

// checks if the Permission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Permission{}

/*
	types and functions for description
*/

// isNotNullableString
type PermissionGetDescriptionAttributeType = *string

func getPermissionGetDescriptionAttributeTypeOk(arg PermissionGetDescriptionAttributeType) (ret PermissionGetDescriptionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPermissionGetDescriptionAttributeType(arg *PermissionGetDescriptionAttributeType, val PermissionGetDescriptionRetType) {
	*arg = &val
}

type PermissionGetDescriptionArgType = string
type PermissionGetDescriptionRetType = string

/*
	types and functions for name
*/

// isNotNullableString
type PermissionGetNameAttributeType = *string

func getPermissionGetNameAttributeTypeOk(arg PermissionGetNameAttributeType) (ret PermissionGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPermissionGetNameAttributeType(arg *PermissionGetNameAttributeType, val PermissionGetNameRetType) {
	*arg = &val
}

type PermissionGetNameArgType = string
type PermissionGetNameRetType = string

// Permission struct for Permission
type Permission struct {
	// REQUIRED
	Description PermissionGetDescriptionAttributeType `json:"description"`
	// REQUIRED
	Name PermissionGetNameAttributeType `json:"name"`
}

type _Permission Permission

// NewPermission instantiates a new Permission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermission(description PermissionGetDescriptionArgType, name PermissionGetNameArgType) *Permission {
	this := Permission{}
	setPermissionGetDescriptionAttributeType(&this.Description, description)
	setPermissionGetNameAttributeType(&this.Name, name)
	return &this
}

// NewPermissionWithDefaults instantiates a new Permission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionWithDefaults() *Permission {
	this := Permission{}
	return &this
}

// GetDescription returns the Description field value
func (o *Permission) GetDescription() (ret PermissionGetDescriptionRetType) {
	ret, _ = o.GetDescriptionOk()
	return ret
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Permission) GetDescriptionOk() (ret PermissionGetDescriptionRetType, ok bool) {
	return getPermissionGetDescriptionAttributeTypeOk(o.Description)
}

// SetDescription sets field value
func (o *Permission) SetDescription(v PermissionGetDescriptionRetType) {
	setPermissionGetDescriptionAttributeType(&o.Description, v)
}

// GetName returns the Name field value
func (o *Permission) GetName() (ret PermissionGetNameRetType) {
	ret, _ = o.GetNameOk()
	return ret
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Permission) GetNameOk() (ret PermissionGetNameRetType, ok bool) {
	return getPermissionGetNameAttributeTypeOk(o.Name)
}

// SetName sets field value
func (o *Permission) SetName(v PermissionGetNameRetType) {
	setPermissionGetNameAttributeType(&o.Name, v)
}

func (o Permission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getPermissionGetDescriptionAttributeTypeOk(o.Description); ok {
		toSerialize["Description"] = val
	}
	if val, ok := getPermissionGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	return toSerialize, nil
}

type NullablePermission struct {
	value *Permission
	isSet bool
}

func (v NullablePermission) Get() *Permission {
	return v.value
}

func (v *NullablePermission) Set(val *Permission) {
	v.value = val
	v.isSet = true
}

func (v NullablePermission) IsSet() bool {
	return v.isSet
}

func (v *NullablePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermission(val *Permission) *NullablePermission {
	return &NullablePermission{value: val, isSet: true}
}

func (v NullablePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
