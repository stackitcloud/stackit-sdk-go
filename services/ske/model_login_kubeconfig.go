/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

import (
	"encoding/json"
)

// checks if the LoginKubeconfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginKubeconfig{}

// LoginKubeconfig struct for LoginKubeconfig
type LoginKubeconfig struct {
	Kubeconfig *string `json:"kubeconfig,omitempty"`
}

// NewLoginKubeconfig instantiates a new LoginKubeconfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginKubeconfig() *LoginKubeconfig {
	this := LoginKubeconfig{}
	return &this
}

// NewLoginKubeconfigWithDefaults instantiates a new LoginKubeconfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginKubeconfigWithDefaults() *LoginKubeconfig {
	this := LoginKubeconfig{}
	return &this
}

// GetKubeconfig returns the Kubeconfig field value if set, zero value otherwise.
func (o *LoginKubeconfig) GetKubeconfig() *string {
	if o == nil || IsNil(o.Kubeconfig) {
		var ret *string
		return ret
	}
	return o.Kubeconfig
}

// GetKubeconfigOk returns a tuple with the Kubeconfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginKubeconfig) GetKubeconfigOk() (*string, bool) {
	if o == nil || IsNil(o.Kubeconfig) {
		return nil, false
	}
	return o.Kubeconfig, true
}

// HasKubeconfig returns a boolean if a field has been set.
func (o *LoginKubeconfig) HasKubeconfig() bool {
	if o != nil && !IsNil(o.Kubeconfig) && !IsNil(o.Kubeconfig) {
		return true
	}

	return false
}

// SetKubeconfig gets a reference to the given string and assigns it to the Kubeconfig field.
func (o *LoginKubeconfig) SetKubeconfig(v *string) {
	o.Kubeconfig = v
}

func (o LoginKubeconfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kubeconfig) {
		toSerialize["kubeconfig"] = o.Kubeconfig
	}
	return toSerialize, nil
}

type NullableLoginKubeconfig struct {
	value *LoginKubeconfig
	isSet bool
}

func (v NullableLoginKubeconfig) Get() *LoginKubeconfig {
	return v.value
}

func (v *NullableLoginKubeconfig) Set(val *LoginKubeconfig) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginKubeconfig) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginKubeconfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginKubeconfig(val *LoginKubeconfig) *NullableLoginKubeconfig {
	return &NullableLoginKubeconfig{value: val, isSet: true}
}

func (v NullableLoginKubeconfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginKubeconfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
