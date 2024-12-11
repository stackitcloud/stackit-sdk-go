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

// checks if the Role type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Role{}

// Role struct for Role
type Role struct {
	// REQUIRED
	Description *string `json:"description"`
	Id          *string `json:"id,omitempty"`
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Permissions *[]Permission `json:"permissions"`
}

type _Role Role

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole(description *string, name *string, permissions *[]Permission) *Role {
	this := Role{}
	this.Description = description
	this.Name = name
	this.Permissions = permissions
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	return &this
}

// GetDescription returns the Description field value
func (o *Role) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Role) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description, true
}

// SetDescription sets field value
func (o *Role) SetDescription(v *string) {
	o.Description = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Role) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Role) HasId() bool {
	if o != nil && !IsNil(o.Id) && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Role) SetId(v *string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Role) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Role) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *Role) SetName(v *string) {
	o.Name = v
}

// GetPermissions returns the Permissions field value
func (o *Role) GetPermissions() *[]Permission {
	if o == nil || IsNil(o.Permissions) {
		var ret *[]Permission
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *Role) GetPermissionsOk() (*[]Permission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *Role) SetPermissions(v *[]Permission) {
	o.Permissions = v
}

func (o Role) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["permissions"] = o.Permissions
	return toSerialize, nil
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
