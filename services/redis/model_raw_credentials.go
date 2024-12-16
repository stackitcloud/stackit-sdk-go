/*
STACKIT Redis API

The STACKIT Redis API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redis

import (
	"encoding/json"
)

// checks if the RawCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RawCredentials{}

// RawCredentials struct for RawCredentials
type RawCredentials struct {
	// REQUIRED
	Credentials *Credentials `json:"credentials"`
}

type _RawCredentials RawCredentials

// NewRawCredentials instantiates a new RawCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRawCredentials(credentials *Credentials) *RawCredentials {
	this := RawCredentials{}
	this.Credentials = credentials
	return &this
}

// NewRawCredentialsWithDefaults instantiates a new RawCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawCredentialsWithDefaults() *RawCredentials {
	this := RawCredentials{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *RawCredentials) GetCredentials() *Credentials {
	if o == nil || IsNil(o.Credentials) {
		var ret *Credentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *RawCredentials) GetCredentialsOk() (*Credentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials, true
}

// SetCredentials sets field value
func (o *RawCredentials) SetCredentials(v *Credentials) {
	o.Credentials = v
}

func (o RawCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentials"] = o.Credentials
	return toSerialize, nil
}

type NullableRawCredentials struct {
	value *RawCredentials
	isSet bool
}

func (v NullableRawCredentials) Get() *RawCredentials {
	return v.value
}

func (v *NullableRawCredentials) Set(val *RawCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableRawCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableRawCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRawCredentials(val *RawCredentials) *NullableRawCredentials {
	return &NullableRawCredentials{value: val, isSet: true}
}

func (v NullableRawCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRawCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
