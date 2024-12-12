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

// checks if the UserPermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPermission{}

// UserPermission struct for UserPermission
type UserPermission struct {
	// REQUIRED
	Permissions *[]ExistingPermission `json:"permissions"`
	// REQUIRED
	ResourceId *string `json:"resourceId"`
	// REQUIRED
	ResourceType *string `json:"resourceType"`
}

type _UserPermission UserPermission

// NewUserPermission instantiates a new UserPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPermission(permissions *[]ExistingPermission, resourceId *string, resourceType *string) *UserPermission {
	this := UserPermission{}
	this.Permissions = permissions
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	return &this
}

// NewUserPermissionWithDefaults instantiates a new UserPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPermissionWithDefaults() *UserPermission {
	this := UserPermission{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *UserPermission) GetPermissions() *[]ExistingPermission {
	if o == nil || IsNil(o.Permissions) {
		var ret *[]ExistingPermission
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *UserPermission) GetPermissionsOk() (*[]ExistingPermission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *UserPermission) SetPermissions(v *[]ExistingPermission) {
	o.Permissions = v
}

// GetResourceId returns the ResourceId field value
func (o *UserPermission) GetResourceId() *string {
	if o == nil || IsNil(o.ResourceId) {
		var ret *string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *UserPermission) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// SetResourceId sets field value
func (o *UserPermission) SetResourceId(v *string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value
func (o *UserPermission) GetResourceType() *string {
	if o == nil || IsNil(o.ResourceType) {
		var ret *string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *UserPermission) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// SetResourceType sets field value
func (o *UserPermission) SetResourceType(v *string) {
	o.ResourceType = v
}

func (o UserPermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissions"] = o.Permissions
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["resourceType"] = o.ResourceType
	return toSerialize, nil
}

type NullableUserPermission struct {
	value *UserPermission
	isSet bool
}

func (v NullableUserPermission) Get() *UserPermission {
	return v.value
}

func (v *NullableUserPermission) Set(val *UserPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPermission(val *UserPermission) *NullableUserPermission {
	return &NullableUserPermission{value: val, isSet: true}
}

func (v NullableUserPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
