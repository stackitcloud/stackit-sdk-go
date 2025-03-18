/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
)

// checks if the ServerConsoleUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerConsoleUrl{}

/*
	types and functions for url
*/

// isNotNullableString
type ServerConsoleUrlGetUrlAttributeType = *string

func getServerConsoleUrlGetUrlAttributeTypeOk(arg ServerConsoleUrlGetUrlAttributeType) (ret ServerConsoleUrlGetUrlRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setServerConsoleUrlGetUrlAttributeType(arg *ServerConsoleUrlGetUrlAttributeType, val ServerConsoleUrlGetUrlRetType) {
	*arg = &val
}

type ServerConsoleUrlGetUrlArgType = string
type ServerConsoleUrlGetUrlRetType = string

// ServerConsoleUrl Object that represents a server console URL.
type ServerConsoleUrl struct {
	// REQUIRED
	Url ServerConsoleUrlGetUrlAttributeType `json:"url"`
}

type _ServerConsoleUrl ServerConsoleUrl

// NewServerConsoleUrl instantiates a new ServerConsoleUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConsoleUrl(url ServerConsoleUrlGetUrlArgType) *ServerConsoleUrl {
	this := ServerConsoleUrl{}
	setServerConsoleUrlGetUrlAttributeType(&this.Url, url)
	return &this
}

// NewServerConsoleUrlWithDefaults instantiates a new ServerConsoleUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConsoleUrlWithDefaults() *ServerConsoleUrl {
	this := ServerConsoleUrl{}
	return &this
}

// GetUrl returns the Url field value
func (o *ServerConsoleUrl) GetUrl() (ret ServerConsoleUrlGetUrlRetType) {
	ret, _ = o.GetUrlOk()
	return ret
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ServerConsoleUrl) GetUrlOk() (ret ServerConsoleUrlGetUrlRetType, ok bool) {
	return getServerConsoleUrlGetUrlAttributeTypeOk(o.Url)
}

// SetUrl sets field value
func (o *ServerConsoleUrl) SetUrl(v ServerConsoleUrlGetUrlRetType) {
	setServerConsoleUrlGetUrlAttributeType(&o.Url, v)
}

func (o ServerConsoleUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getServerConsoleUrlGetUrlAttributeTypeOk(o.Url); ok {
		toSerialize["Url"] = val
	}
	return toSerialize, nil
}

type NullableServerConsoleUrl struct {
	value *ServerConsoleUrl
	isSet bool
}

func (v NullableServerConsoleUrl) Get() *ServerConsoleUrl {
	return v.value
}

func (v *NullableServerConsoleUrl) Set(val *ServerConsoleUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConsoleUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConsoleUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConsoleUrl(val *ServerConsoleUrl) *NullableServerConsoleUrl {
	return &NullableServerConsoleUrl{value: val, isSet: true}
}

func (v NullableServerConsoleUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConsoleUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
