/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 2beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alb

import (
	"encoding/json"
)

// checks if the CreateCredentialsPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCredentialsPayload{}

/*
	types and functions for displayName
*/

// isNotNullableString
type CreateCredentialsPayloadGetDisplayNameAttributeType = *string

func getCreateCredentialsPayloadGetDisplayNameAttributeTypeOk(arg CreateCredentialsPayloadGetDisplayNameAttributeType) (ret CreateCredentialsPayloadGetDisplayNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateCredentialsPayloadGetDisplayNameAttributeType(arg *CreateCredentialsPayloadGetDisplayNameAttributeType, val CreateCredentialsPayloadGetDisplayNameRetType) {
	*arg = &val
}

type CreateCredentialsPayloadGetDisplayNameArgType = string
type CreateCredentialsPayloadGetDisplayNameRetType = string

/*
	types and functions for password
*/

// isNotNullableString
type CreateCredentialsPayloadGetPasswordAttributeType = *string

func getCreateCredentialsPayloadGetPasswordAttributeTypeOk(arg CreateCredentialsPayloadGetPasswordAttributeType) (ret CreateCredentialsPayloadGetPasswordRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateCredentialsPayloadGetPasswordAttributeType(arg *CreateCredentialsPayloadGetPasswordAttributeType, val CreateCredentialsPayloadGetPasswordRetType) {
	*arg = &val
}

type CreateCredentialsPayloadGetPasswordArgType = string
type CreateCredentialsPayloadGetPasswordRetType = string

/*
	types and functions for username
*/

// isNotNullableString
type CreateCredentialsPayloadGetUsernameAttributeType = *string

func getCreateCredentialsPayloadGetUsernameAttributeTypeOk(arg CreateCredentialsPayloadGetUsernameAttributeType) (ret CreateCredentialsPayloadGetUsernameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateCredentialsPayloadGetUsernameAttributeType(arg *CreateCredentialsPayloadGetUsernameAttributeType, val CreateCredentialsPayloadGetUsernameRetType) {
	*arg = &val
}

type CreateCredentialsPayloadGetUsernameArgType = string
type CreateCredentialsPayloadGetUsernameRetType = string

// CreateCredentialsPayload struct for CreateCredentialsPayload
type CreateCredentialsPayload struct {
	// Credential name
	DisplayName CreateCredentialsPayloadGetDisplayNameAttributeType `json:"displayName,omitempty"`
	// A valid password used for an existing ARGUS instance, which is used during basic auth.
	Password CreateCredentialsPayloadGetPasswordAttributeType `json:"password,omitempty"`
	// A valid username used for an existing ARGUS instance, which is used during basic auth.
	Username CreateCredentialsPayloadGetUsernameAttributeType `json:"username,omitempty"`
}

// NewCreateCredentialsPayload instantiates a new CreateCredentialsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCredentialsPayload() *CreateCredentialsPayload {
	this := CreateCredentialsPayload{}
	return &this
}

// NewCreateCredentialsPayloadWithDefaults instantiates a new CreateCredentialsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCredentialsPayloadWithDefaults() *CreateCredentialsPayload {
	this := CreateCredentialsPayload{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CreateCredentialsPayload) GetDisplayName() (res CreateCredentialsPayloadGetDisplayNameRetType) {
	res, _ = o.GetDisplayNameOk()
	return
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCredentialsPayload) GetDisplayNameOk() (ret CreateCredentialsPayloadGetDisplayNameRetType, ok bool) {
	return getCreateCredentialsPayloadGetDisplayNameAttributeTypeOk(o.DisplayName)
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateCredentialsPayload) HasDisplayName() bool {
	_, ok := o.GetDisplayNameOk()
	return ok
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CreateCredentialsPayload) SetDisplayName(v CreateCredentialsPayloadGetDisplayNameRetType) {
	setCreateCredentialsPayloadGetDisplayNameAttributeType(&o.DisplayName, v)
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateCredentialsPayload) GetPassword() (res CreateCredentialsPayloadGetPasswordRetType) {
	res, _ = o.GetPasswordOk()
	return
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCredentialsPayload) GetPasswordOk() (ret CreateCredentialsPayloadGetPasswordRetType, ok bool) {
	return getCreateCredentialsPayloadGetPasswordAttributeTypeOk(o.Password)
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateCredentialsPayload) HasPassword() bool {
	_, ok := o.GetPasswordOk()
	return ok
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateCredentialsPayload) SetPassword(v CreateCredentialsPayloadGetPasswordRetType) {
	setCreateCredentialsPayloadGetPasswordAttributeType(&o.Password, v)
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateCredentialsPayload) GetUsername() (res CreateCredentialsPayloadGetUsernameRetType) {
	res, _ = o.GetUsernameOk()
	return
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCredentialsPayload) GetUsernameOk() (ret CreateCredentialsPayloadGetUsernameRetType, ok bool) {
	return getCreateCredentialsPayloadGetUsernameAttributeTypeOk(o.Username)
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateCredentialsPayload) HasUsername() bool {
	_, ok := o.GetUsernameOk()
	return ok
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateCredentialsPayload) SetUsername(v CreateCredentialsPayloadGetUsernameRetType) {
	setCreateCredentialsPayloadGetUsernameAttributeType(&o.Username, v)
}

func (o CreateCredentialsPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateCredentialsPayloadGetDisplayNameAttributeTypeOk(o.DisplayName); ok {
		toSerialize["DisplayName"] = val
	}
	if val, ok := getCreateCredentialsPayloadGetPasswordAttributeTypeOk(o.Password); ok {
		toSerialize["Password"] = val
	}
	if val, ok := getCreateCredentialsPayloadGetUsernameAttributeTypeOk(o.Username); ok {
		toSerialize["Username"] = val
	}
	return toSerialize, nil
}

type NullableCreateCredentialsPayload struct {
	value *CreateCredentialsPayload
	isSet bool
}

func (v NullableCreateCredentialsPayload) Get() *CreateCredentialsPayload {
	return v.value
}

func (v *NullableCreateCredentialsPayload) Set(val *CreateCredentialsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCredentialsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCredentialsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCredentialsPayload(val *CreateCredentialsPayload) *NullableCreateCredentialsPayload {
	return &NullableCreateCredentialsPayload{value: val, isSet: true}
}

func (v NullableCreateCredentialsPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCredentialsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
