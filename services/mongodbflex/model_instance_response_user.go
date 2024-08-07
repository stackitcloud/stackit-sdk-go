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

// checks if the InstanceResponseUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceResponseUser{}

// InstanceResponseUser struct for InstanceResponseUser
type InstanceResponseUser struct {
	Database *string   `json:"database,omitempty"`
	Host     *string   `json:"host,omitempty"`
	Id       *string   `json:"id,omitempty"`
	Port     *int64    `json:"port,omitempty"`
	Roles    *[]string `json:"roles,omitempty"`
	Username *string   `json:"username,omitempty"`
}

// NewInstanceResponseUser instantiates a new InstanceResponseUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceResponseUser() *InstanceResponseUser {
	this := InstanceResponseUser{}
	return &this
}

// NewInstanceResponseUserWithDefaults instantiates a new InstanceResponseUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceResponseUserWithDefaults() *InstanceResponseUser {
	this := InstanceResponseUser{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *InstanceResponseUser) GetDatabase() *string {
	if o == nil || IsNil(o.Database) {
		var ret *string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResponseUser) GetDatabaseOk() (*string, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *InstanceResponseUser) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *InstanceResponseUser) SetDatabase(v *string) {
	o.Database = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *InstanceResponseUser) GetHost() *string {
	if o == nil || IsNil(o.Host) {
		var ret *string
		return ret
	}
	return o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResponseUser) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *InstanceResponseUser) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *InstanceResponseUser) SetHost(v *string) {
	o.Host = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstanceResponseUser) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResponseUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InstanceResponseUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InstanceResponseUser) SetId(v *string) {
	o.Id = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InstanceResponseUser) GetPort() *int64 {
	if o == nil || IsNil(o.Port) {
		var ret *int64
		return ret
	}
	return o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResponseUser) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InstanceResponseUser) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *InstanceResponseUser) SetPort(v *int64) {
	o.Port = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *InstanceResponseUser) GetRoles() *[]string {
	if o == nil || IsNil(o.Roles) {
		var ret *[]string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResponseUser) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *InstanceResponseUser) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *InstanceResponseUser) SetRoles(v *[]string) {
	o.Roles = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *InstanceResponseUser) GetUsername() *string {
	if o == nil || IsNil(o.Username) {
		var ret *string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResponseUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *InstanceResponseUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *InstanceResponseUser) SetUsername(v *string) {
	o.Username = v
}

func (o InstanceResponseUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
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

type NullableInstanceResponseUser struct {
	value *InstanceResponseUser
	isSet bool
}

func (v NullableInstanceResponseUser) Get() *InstanceResponseUser {
	return v.value
}

func (v *NullableInstanceResponseUser) Set(val *InstanceResponseUser) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceResponseUser) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceResponseUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceResponseUser(val *InstanceResponseUser) *NullableInstanceResponseUser {
	return &NullableInstanceResponseUser{value: val, isSet: true}
}

func (v NullableInstanceResponseUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceResponseUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
