/*
STACKIT LogMe API

The STACKIT LogMe API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logme

import (
	"encoding/json"
)

// checks if the Credentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Credentials{}

// Credentials struct for Credentials
type Credentials struct {
	// REQUIRED
	Host *string `json:"host"`
	// REQUIRED
	Password       *string `json:"password"`
	Port           *int64  `json:"port,omitempty"`
	SyslogDrainUrl *string `json:"syslog_drain_url,omitempty"`
	Uri            *string `json:"uri,omitempty"`
	// REQUIRED
	Username *string `json:"username"`
}

type _Credentials Credentials

// NewCredentials instantiates a new Credentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentials(host *string, password *string, username *string) *Credentials {
	this := Credentials{}
	this.Host = host
	this.Password = password
	this.Username = username
	return &this
}

// NewCredentialsWithDefaults instantiates a new Credentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsWithDefaults() *Credentials {
	this := Credentials{}
	return &this
}

// GetHost returns the Host field value
func (o *Credentials) GetHost() *string {
	if o == nil || IsNil(o.Host) {
		var ret *string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *Credentials) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host, true
}

// SetHost sets field value
func (o *Credentials) SetHost(v *string) {
	o.Host = v
}

// GetPassword returns the Password field value
func (o *Credentials) GetPassword() *string {
	if o == nil || IsNil(o.Password) {
		var ret *string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *Credentials) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password, true
}

// SetPassword sets field value
func (o *Credentials) SetPassword(v *string) {
	o.Password = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Credentials) GetPort() *int64 {
	if o == nil || IsNil(o.Port) {
		var ret *int64
		return ret
	}
	return o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credentials) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Credentials) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *Credentials) SetPort(v *int64) {
	o.Port = v
}

// GetSyslogDrainUrl returns the SyslogDrainUrl field value if set, zero value otherwise.
func (o *Credentials) GetSyslogDrainUrl() *string {
	if o == nil || IsNil(o.SyslogDrainUrl) {
		var ret *string
		return ret
	}
	return o.SyslogDrainUrl
}

// GetSyslogDrainUrlOk returns a tuple with the SyslogDrainUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credentials) GetSyslogDrainUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SyslogDrainUrl) {
		return nil, false
	}
	return o.SyslogDrainUrl, true
}

// HasSyslogDrainUrl returns a boolean if a field has been set.
func (o *Credentials) HasSyslogDrainUrl() bool {
	if o != nil && !IsNil(o.SyslogDrainUrl) {
		return true
	}

	return false
}

// SetSyslogDrainUrl gets a reference to the given string and assigns it to the SyslogDrainUrl field.
func (o *Credentials) SetSyslogDrainUrl(v *string) {
	o.SyslogDrainUrl = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Credentials) GetUri() *string {
	if o == nil || IsNil(o.Uri) {
		var ret *string
		return ret
	}
	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credentials) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Credentials) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Credentials) SetUri(v *string) {
	o.Uri = v
}

// GetUsername returns the Username field value
func (o *Credentials) GetUsername() *string {
	if o == nil || IsNil(o.Username) {
		var ret *string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Credentials) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username, true
}

// SetUsername sets field value
func (o *Credentials) SetUsername(v *string) {
	o.Username = v
}

func (o Credentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["password"] = o.Password
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.SyslogDrainUrl) {
		toSerialize["syslog_drain_url"] = o.SyslogDrainUrl
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableCredentials struct {
	value *Credentials
	isSet bool
}

func (v NullableCredentials) Get() *Credentials {
	return v.value
}

func (v *NullableCredentials) Set(val *Credentials) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentials(val *Credentials) *NullableCredentials {
	return &NullableCredentials{value: val, isSet: true}
}

func (v NullableCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
