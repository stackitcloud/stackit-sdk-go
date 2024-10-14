/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
)

// checks if the UserResponseUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResponseUser{}

// UserResponseUser struct for UserResponseUser
type UserResponseUser struct {
	DefaultDatabase *string   `json:"default_database,omitempty"`
	Host            *string   `json:"host,omitempty"`
	Id              *string   `json:"id,omitempty"`
	Port            *int64    `json:"port,omitempty"`
	Roles           *[]string `json:"roles,omitempty"`
	Username        *string   `json:"username,omitempty"`
}

// NewUserResponseUser instantiates a new UserResponseUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponseUser() *UserResponseUser {
	this := UserResponseUser{}
	return &this
}

// NewUserResponseUserWithDefaults instantiates a new UserResponseUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseUserWithDefaults() *UserResponseUser {
	this := UserResponseUser{}
	return &this
}

// GetDefaultDatabase returns the DefaultDatabase field value if set, zero value otherwise.
func (o *UserResponseUser) GetDefaultDatabase() *string {
	if o == nil || IsNil(o.DefaultDatabase) {
		var ret *string
		return ret
	}
	return o.DefaultDatabase
}

// GetDefaultDatabaseOk returns a tuple with the DefaultDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseUser) GetDefaultDatabaseOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultDatabase) {
		return nil, false
	}
	return o.DefaultDatabase, true
}

// HasDefaultDatabase returns a boolean if a field has been set.
func (o *UserResponseUser) HasDefaultDatabase() bool {
	if o != nil && !IsNil(o.DefaultDatabase) {
		return true
	}

	return false
}

// SetDefaultDatabase gets a reference to the given string and assigns it to the DefaultDatabase field.
func (o *UserResponseUser) SetDefaultDatabase(v *string) {
	o.DefaultDatabase = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *UserResponseUser) GetHost() *string {
	if o == nil || IsNil(o.Host) {
		var ret *string
		return ret
	}
	return o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseUser) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *UserResponseUser) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *UserResponseUser) SetHost(v *string) {
	o.Host = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserResponseUser) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserResponseUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserResponseUser) SetId(v *string) {
	o.Id = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UserResponseUser) GetPort() *int64 {
	if o == nil || IsNil(o.Port) {
		var ret *int64
		return ret
	}
	return o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseUser) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UserResponseUser) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *UserResponseUser) SetPort(v *int64) {
	o.Port = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserResponseUser) GetRoles() *[]string {
	if o == nil || IsNil(o.Roles) {
		var ret *[]string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseUser) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserResponseUser) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UserResponseUser) SetRoles(v *[]string) {
	o.Roles = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserResponseUser) GetUsername() *string {
	if o == nil || IsNil(o.Username) {
		var ret *string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserResponseUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserResponseUser) SetUsername(v *string) {
	o.Username = v
}

func (o UserResponseUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultDatabase) {
		toSerialize["default_database"] = o.DefaultDatabase
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableUserResponseUser struct {
	value *UserResponseUser
	isSet bool
}

func (v NullableUserResponseUser) Get() *UserResponseUser {
	return v.value
}

func (v *NullableUserResponseUser) Set(val *UserResponseUser) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponseUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponseUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponseUser(val *UserResponseUser) *NullableUserResponseUser {
	return &NullableUserResponseUser{value: val, isSet: true}
}

func (v NullableUserResponseUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponseUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
