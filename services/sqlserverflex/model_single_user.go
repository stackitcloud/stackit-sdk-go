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

// checks if the SingleUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleUser{}

/*
	types and functions for default_database
*/

// isNotNullableString
type SingleUserGetDefaultDatabaseAttributeType = *string

func getSingleUserGetDefaultDatabaseAttributeTypeOk(arg SingleUserGetDefaultDatabaseAttributeType) (ret SingleUserGetDefaultDatabaseRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSingleUserGetDefaultDatabaseAttributeType(arg *SingleUserGetDefaultDatabaseAttributeType, val SingleUserGetDefaultDatabaseRetType) {
	*arg = &val
}

type SingleUserGetDefaultDatabaseArgType = string
type SingleUserGetDefaultDatabaseRetType = string

/*
	types and functions for host
*/

// isNotNullableString
type SingleUserGetHostAttributeType = *string

func getSingleUserGetHostAttributeTypeOk(arg SingleUserGetHostAttributeType) (ret SingleUserGetHostRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSingleUserGetHostAttributeType(arg *SingleUserGetHostAttributeType, val SingleUserGetHostRetType) {
	*arg = &val
}

type SingleUserGetHostArgType = string
type SingleUserGetHostRetType = string

/*
	types and functions for id
*/

// isNotNullableString
type SingleUserGetIdAttributeType = *string

func getSingleUserGetIdAttributeTypeOk(arg SingleUserGetIdAttributeType) (ret SingleUserGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSingleUserGetIdAttributeType(arg *SingleUserGetIdAttributeType, val SingleUserGetIdRetType) {
	*arg = &val
}

type SingleUserGetIdArgType = string
type SingleUserGetIdRetType = string

/*
	types and functions for password
*/

// isNotNullableString
type SingleUserGetPasswordAttributeType = *string

func getSingleUserGetPasswordAttributeTypeOk(arg SingleUserGetPasswordAttributeType) (ret SingleUserGetPasswordRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSingleUserGetPasswordAttributeType(arg *SingleUserGetPasswordAttributeType, val SingleUserGetPasswordRetType) {
	*arg = &val
}

type SingleUserGetPasswordArgType = string
type SingleUserGetPasswordRetType = string

/*
	types and functions for port
*/

// isLong
type SingleUserGetPortAttributeType = *int64
type SingleUserGetPortArgType = int64
type SingleUserGetPortRetType = int64

func getSingleUserGetPortAttributeTypeOk(arg SingleUserGetPortAttributeType) (ret SingleUserGetPortRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSingleUserGetPortAttributeType(arg *SingleUserGetPortAttributeType, val SingleUserGetPortRetType) {
	*arg = &val
}

/*
	types and functions for roles
*/

// isArray
type SingleUserGetRolesAttributeType = *[]string
type SingleUserGetRolesArgType = []string
type SingleUserGetRolesRetType = []string

func getSingleUserGetRolesAttributeTypeOk(arg SingleUserGetRolesAttributeType) (ret SingleUserGetRolesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSingleUserGetRolesAttributeType(arg *SingleUserGetRolesAttributeType, val SingleUserGetRolesRetType) {
	*arg = &val
}

/*
	types and functions for uri
*/

// isNotNullableString
type SingleUserGetUriAttributeType = *string

func getSingleUserGetUriAttributeTypeOk(arg SingleUserGetUriAttributeType) (ret SingleUserGetUriRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSingleUserGetUriAttributeType(arg *SingleUserGetUriAttributeType, val SingleUserGetUriRetType) {
	*arg = &val
}

type SingleUserGetUriArgType = string
type SingleUserGetUriRetType = string

/*
	types and functions for username
*/

// isNotNullableString
type SingleUserGetUsernameAttributeType = *string

func getSingleUserGetUsernameAttributeTypeOk(arg SingleUserGetUsernameAttributeType) (ret SingleUserGetUsernameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSingleUserGetUsernameAttributeType(arg *SingleUserGetUsernameAttributeType, val SingleUserGetUsernameRetType) {
	*arg = &val
}

type SingleUserGetUsernameArgType = string
type SingleUserGetUsernameRetType = string

// SingleUser struct for SingleUser
type SingleUser struct {
	DefaultDatabase SingleUserGetDefaultDatabaseAttributeType `json:"default_database,omitempty"`
	Host            SingleUserGetHostAttributeType            `json:"host,omitempty"`
	Id              SingleUserGetIdAttributeType              `json:"id,omitempty"`
	Password        SingleUserGetPasswordAttributeType        `json:"password,omitempty"`
	Port            SingleUserGetPortAttributeType            `json:"port,omitempty"`
	Roles           SingleUserGetRolesAttributeType           `json:"roles,omitempty"`
	Uri             SingleUserGetUriAttributeType             `json:"uri,omitempty"`
	Username        SingleUserGetUsernameAttributeType        `json:"username,omitempty"`
}

// NewSingleUser instantiates a new SingleUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleUser() *SingleUser {
	this := SingleUser{}
	return &this
}

// NewSingleUserWithDefaults instantiates a new SingleUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleUserWithDefaults() *SingleUser {
	this := SingleUser{}
	return &this
}

// GetDefaultDatabase returns the DefaultDatabase field value if set, zero value otherwise.
func (o *SingleUser) GetDefaultDatabase() (res SingleUserGetDefaultDatabaseRetType) {
	res, _ = o.GetDefaultDatabaseOk()
	return
}

// GetDefaultDatabaseOk returns a tuple with the DefaultDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUser) GetDefaultDatabaseOk() (ret SingleUserGetDefaultDatabaseRetType, ok bool) {
	return getSingleUserGetDefaultDatabaseAttributeTypeOk(o.DefaultDatabase)
}

// HasDefaultDatabase returns a boolean if a field has been set.
func (o *SingleUser) HasDefaultDatabase() bool {
	_, ok := o.GetDefaultDatabaseOk()
	return ok
}

// SetDefaultDatabase gets a reference to the given string and assigns it to the DefaultDatabase field.
func (o *SingleUser) SetDefaultDatabase(v SingleUserGetDefaultDatabaseRetType) {
	setSingleUserGetDefaultDatabaseAttributeType(&o.DefaultDatabase, v)
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *SingleUser) GetHost() (res SingleUserGetHostRetType) {
	res, _ = o.GetHostOk()
	return
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUser) GetHostOk() (ret SingleUserGetHostRetType, ok bool) {
	return getSingleUserGetHostAttributeTypeOk(o.Host)
}

// HasHost returns a boolean if a field has been set.
func (o *SingleUser) HasHost() bool {
	_, ok := o.GetHostOk()
	return ok
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *SingleUser) SetHost(v SingleUserGetHostRetType) {
	setSingleUserGetHostAttributeType(&o.Host, v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SingleUser) GetId() (res SingleUserGetIdRetType) {
	res, _ = o.GetIdOk()
	return
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUser) GetIdOk() (ret SingleUserGetIdRetType, ok bool) {
	return getSingleUserGetIdAttributeTypeOk(o.Id)
}

// HasId returns a boolean if a field has been set.
func (o *SingleUser) HasId() bool {
	_, ok := o.GetIdOk()
	return ok
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SingleUser) SetId(v SingleUserGetIdRetType) {
	setSingleUserGetIdAttributeType(&o.Id, v)
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SingleUser) GetPassword() (res SingleUserGetPasswordRetType) {
	res, _ = o.GetPasswordOk()
	return
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUser) GetPasswordOk() (ret SingleUserGetPasswordRetType, ok bool) {
	return getSingleUserGetPasswordAttributeTypeOk(o.Password)
}

// HasPassword returns a boolean if a field has been set.
func (o *SingleUser) HasPassword() bool {
	_, ok := o.GetPasswordOk()
	return ok
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SingleUser) SetPassword(v SingleUserGetPasswordRetType) {
	setSingleUserGetPasswordAttributeType(&o.Password, v)
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SingleUser) GetPort() (res SingleUserGetPortRetType) {
	res, _ = o.GetPortOk()
	return
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUser) GetPortOk() (ret SingleUserGetPortRetType, ok bool) {
	return getSingleUserGetPortAttributeTypeOk(o.Port)
}

// HasPort returns a boolean if a field has been set.
func (o *SingleUser) HasPort() bool {
	_, ok := o.GetPortOk()
	return ok
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *SingleUser) SetPort(v SingleUserGetPortRetType) {
	setSingleUserGetPortAttributeType(&o.Port, v)
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *SingleUser) GetRoles() (res SingleUserGetRolesRetType) {
	res, _ = o.GetRolesOk()
	return
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUser) GetRolesOk() (ret SingleUserGetRolesRetType, ok bool) {
	return getSingleUserGetRolesAttributeTypeOk(o.Roles)
}

// HasRoles returns a boolean if a field has been set.
func (o *SingleUser) HasRoles() bool {
	_, ok := o.GetRolesOk()
	return ok
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *SingleUser) SetRoles(v SingleUserGetRolesRetType) {
	setSingleUserGetRolesAttributeType(&o.Roles, v)
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *SingleUser) GetUri() (res SingleUserGetUriRetType) {
	res, _ = o.GetUriOk()
	return
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUser) GetUriOk() (ret SingleUserGetUriRetType, ok bool) {
	return getSingleUserGetUriAttributeTypeOk(o.Uri)
}

// HasUri returns a boolean if a field has been set.
func (o *SingleUser) HasUri() bool {
	_, ok := o.GetUriOk()
	return ok
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *SingleUser) SetUri(v SingleUserGetUriRetType) {
	setSingleUserGetUriAttributeType(&o.Uri, v)
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SingleUser) GetUsername() (res SingleUserGetUsernameRetType) {
	res, _ = o.GetUsernameOk()
	return
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUser) GetUsernameOk() (ret SingleUserGetUsernameRetType, ok bool) {
	return getSingleUserGetUsernameAttributeTypeOk(o.Username)
}

// HasUsername returns a boolean if a field has been set.
func (o *SingleUser) HasUsername() bool {
	_, ok := o.GetUsernameOk()
	return ok
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SingleUser) SetUsername(v SingleUserGetUsernameRetType) {
	setSingleUserGetUsernameAttributeType(&o.Username, v)
}

func (o SingleUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getSingleUserGetDefaultDatabaseAttributeTypeOk(o.DefaultDatabase); ok {
		toSerialize["DefaultDatabase"] = val
	}
	if val, ok := getSingleUserGetHostAttributeTypeOk(o.Host); ok {
		toSerialize["Host"] = val
	}
	if val, ok := getSingleUserGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getSingleUserGetPasswordAttributeTypeOk(o.Password); ok {
		toSerialize["Password"] = val
	}
	if val, ok := getSingleUserGetPortAttributeTypeOk(o.Port); ok {
		toSerialize["Port"] = val
	}
	if val, ok := getSingleUserGetRolesAttributeTypeOk(o.Roles); ok {
		toSerialize["Roles"] = val
	}
	if val, ok := getSingleUserGetUriAttributeTypeOk(o.Uri); ok {
		toSerialize["Uri"] = val
	}
	if val, ok := getSingleUserGetUsernameAttributeTypeOk(o.Username); ok {
		toSerialize["Username"] = val
	}
	return toSerialize, nil
}

type NullableSingleUser struct {
	value *SingleUser
	isSet bool
}

func (v NullableSingleUser) Get() *SingleUser {
	return v.value
}

func (v *NullableSingleUser) Set(val *SingleUser) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleUser) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleUser(val *SingleUser) *NullableSingleUser {
	return &NullableSingleUser{value: val, isSet: true}
}

func (v NullableSingleUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
