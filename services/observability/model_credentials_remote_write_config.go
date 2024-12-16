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

// checks if the CredentialsRemoteWriteConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsRemoteWriteConfig{}

// CredentialsRemoteWriteConfig struct for CredentialsRemoteWriteConfig
type CredentialsRemoteWriteConfig struct {
	// REQUIRED
	CredentialsMaxLimit *int64 `json:"credentialsMaxLimit"`
	// REQUIRED
	MaxLimit *int64 `json:"maxLimit"`
	// REQUIRED
	Message *string `json:"message"`
}

type _CredentialsRemoteWriteConfig CredentialsRemoteWriteConfig

// NewCredentialsRemoteWriteConfig instantiates a new CredentialsRemoteWriteConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsRemoteWriteConfig(credentialsMaxLimit *int64, maxLimit *int64, message *string) *CredentialsRemoteWriteConfig {
	this := CredentialsRemoteWriteConfig{}
	this.CredentialsMaxLimit = credentialsMaxLimit
	this.MaxLimit = maxLimit
	this.Message = message
	return &this
}

// NewCredentialsRemoteWriteConfigWithDefaults instantiates a new CredentialsRemoteWriteConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsRemoteWriteConfigWithDefaults() *CredentialsRemoteWriteConfig {
	this := CredentialsRemoteWriteConfig{}
	return &this
}

// GetCredentialsMaxLimit returns the CredentialsMaxLimit field value
func (o *CredentialsRemoteWriteConfig) GetCredentialsMaxLimit() *int64 {
	if o == nil || IsNil(o.CredentialsMaxLimit) {
		var ret *int64
		return ret
	}

	return o.CredentialsMaxLimit
}

// GetCredentialsMaxLimitOk returns a tuple with the CredentialsMaxLimit field value
// and a boolean to check if the value has been set.
func (o *CredentialsRemoteWriteConfig) GetCredentialsMaxLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CredentialsMaxLimit, true
}

// SetCredentialsMaxLimit sets field value
func (o *CredentialsRemoteWriteConfig) SetCredentialsMaxLimit(v *int64) {
	o.CredentialsMaxLimit = v
}

// GetMaxLimit returns the MaxLimit field value
func (o *CredentialsRemoteWriteConfig) GetMaxLimit() *int64 {
	if o == nil || IsNil(o.MaxLimit) {
		var ret *int64
		return ret
	}

	return o.MaxLimit
}

// GetMaxLimitOk returns a tuple with the MaxLimit field value
// and a boolean to check if the value has been set.
func (o *CredentialsRemoteWriteConfig) GetMaxLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxLimit, true
}

// SetMaxLimit sets field value
func (o *CredentialsRemoteWriteConfig) SetMaxLimit(v *int64) {
	o.MaxLimit = v
}

// GetMessage returns the Message field value
func (o *CredentialsRemoteWriteConfig) GetMessage() *string {
	if o == nil || IsNil(o.Message) {
		var ret *string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CredentialsRemoteWriteConfig) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *CredentialsRemoteWriteConfig) SetMessage(v *string) {
	o.Message = v
}

func (o CredentialsRemoteWriteConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentialsMaxLimit"] = o.CredentialsMaxLimit
	toSerialize["maxLimit"] = o.MaxLimit
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableCredentialsRemoteWriteConfig struct {
	value *CredentialsRemoteWriteConfig
	isSet bool
}

func (v NullableCredentialsRemoteWriteConfig) Get() *CredentialsRemoteWriteConfig {
	return v.value
}

func (v *NullableCredentialsRemoteWriteConfig) Set(val *CredentialsRemoteWriteConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsRemoteWriteConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsRemoteWriteConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsRemoteWriteConfig(val *CredentialsRemoteWriteConfig) *NullableCredentialsRemoteWriteConfig {
	return &NullableCredentialsRemoteWriteConfig{value: val, isSet: true}
}

func (v NullableCredentialsRemoteWriteConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsRemoteWriteConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
