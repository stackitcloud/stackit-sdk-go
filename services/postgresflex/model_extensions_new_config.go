/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

import (
	"encoding/json"
)

// checks if the ExtensionsNewConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionsNewConfig{}

// ExtensionsNewConfig struct for ExtensionsNewConfig
type ExtensionsNewConfig struct {
	Configuration *[]ExtensionsConfiguration `json:"configuration,omitempty"`
}

// NewExtensionsNewConfig instantiates a new ExtensionsNewConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionsNewConfig() *ExtensionsNewConfig {
	this := ExtensionsNewConfig{}
	return &this
}

// NewExtensionsNewConfigWithDefaults instantiates a new ExtensionsNewConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionsNewConfigWithDefaults() *ExtensionsNewConfig {
	this := ExtensionsNewConfig{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ExtensionsNewConfig) GetConfiguration() *[]ExtensionsConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret *[]ExtensionsConfiguration
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionsNewConfig) GetConfigurationOk() (*[]ExtensionsConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ExtensionsNewConfig) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given []ExtensionsConfiguration and assigns it to the Configuration field.
func (o *ExtensionsNewConfig) SetConfiguration(v *[]ExtensionsConfiguration) {
	o.Configuration = v
}

func (o ExtensionsNewConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	return toSerialize, nil
}

type NullableExtensionsNewConfig struct {
	value *ExtensionsNewConfig
	isSet bool
}

func (v NullableExtensionsNewConfig) Get() *ExtensionsNewConfig {
	return v.value
}

func (v *NullableExtensionsNewConfig) Set(val *ExtensionsNewConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionsNewConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionsNewConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionsNewConfig(val *ExtensionsNewConfig) *NullableExtensionsNewConfig {
	return &NullableExtensionsNewConfig{value: val, isSet: true}
}

func (v NullableExtensionsNewConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionsNewConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
