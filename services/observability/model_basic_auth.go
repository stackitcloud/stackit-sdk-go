/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"encoding/json"
)

// checks if the BasicAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasicAuth{}

// BasicAuth struct for BasicAuth
type BasicAuth struct {
	// REQUIRED
	Password *string `json:"password"`
	// REQUIRED
	Username *string `json:"username"`
}

type _BasicAuth BasicAuth

// NewBasicAuth instantiates a new BasicAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicAuth(password *string, username *string) *BasicAuth {
	this := BasicAuth{}
	this.Password = password
	this.Username = username
	return &this
}

// NewBasicAuthWithDefaults instantiates a new BasicAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicAuthWithDefaults() *BasicAuth {
	this := BasicAuth{}
	return &this
}

// GetPassword returns the Password field value
func (o *BasicAuth) GetPassword() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *BasicAuth) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password, true
}

// SetPassword sets field value
func (o *BasicAuth) SetPassword(v *string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *BasicAuth) GetUsername() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *BasicAuth) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username, true
}

// SetUsername sets field value
func (o *BasicAuth) SetUsername(v *string) {
	o.Username = v
}

func (o BasicAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableBasicAuth struct {
	value *BasicAuth
	isSet bool
}

func (v NullableBasicAuth) Get() *BasicAuth {
	return v.value
}

func (v *NullableBasicAuth) Set(val *BasicAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicAuth(val *BasicAuth) *NullableBasicAuth {
	return &NullableBasicAuth{value: val, isSet: true}
}

func (v NullableBasicAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
