/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lbapplication

import (
	"encoding/json"
)

// checks if the GoogleProtobufAny type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleProtobufAny{}

/*
	types and functions for @type
*/

// isNotNullableString
type GoogleProtobufAnyGetTypeAttributeType = *string

func getGoogleProtobufAnyGetTypeAttributeTypeOk(arg GoogleProtobufAnyGetTypeAttributeType) (ret GoogleProtobufAnyGetTypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGoogleProtobufAnyGetTypeAttributeType(arg *GoogleProtobufAnyGetTypeAttributeType, val GoogleProtobufAnyGetTypeRetType) {
	*arg = &val
}

type GoogleProtobufAnyGetTypeArgType = string
type GoogleProtobufAnyGetTypeRetType = string

// GoogleProtobufAny Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type GoogleProtobufAny struct {
	// The type of the serialized message.
	Type                 GoogleProtobufAnyGetTypeAttributeType `json:"@type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GoogleProtobufAny GoogleProtobufAny

// NewGoogleProtobufAny instantiates a new GoogleProtobufAny object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleProtobufAny() *GoogleProtobufAny {
	this := GoogleProtobufAny{}
	return &this
}

// NewGoogleProtobufAnyWithDefaults instantiates a new GoogleProtobufAny object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleProtobufAnyWithDefaults() *GoogleProtobufAny {
	this := GoogleProtobufAny{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GoogleProtobufAny) GetType() (res GoogleProtobufAnyGetTypeRetType) {
	res, _ = o.GetTypeOk()
	return
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleProtobufAny) GetTypeOk() (ret GoogleProtobufAnyGetTypeRetType, ok bool) {
	return getGoogleProtobufAnyGetTypeAttributeTypeOk(o.Type)
}

// HasType returns a boolean if a field has been set.
func (o *GoogleProtobufAny) HasType() bool {
	_, ok := o.GetTypeOk()
	return ok
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GoogleProtobufAny) SetType(v GoogleProtobufAnyGetTypeRetType) {
	setGoogleProtobufAnyGetTypeAttributeType(&o.Type, v)
}

func (o GoogleProtobufAny) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getGoogleProtobufAnyGetTypeAttributeTypeOk(o.Type); ok {
		toSerialize["Type"] = val
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableGoogleProtobufAny struct {
	value *GoogleProtobufAny
	isSet bool
}

func (v NullableGoogleProtobufAny) Get() *GoogleProtobufAny {
	return v.value
}

func (v *NullableGoogleProtobufAny) Set(val *GoogleProtobufAny) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleProtobufAny) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleProtobufAny) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleProtobufAny(val *GoogleProtobufAny) *NullableGoogleProtobufAny {
	return &NullableGoogleProtobufAny{value: val, isSet: true}
}

func (v NullableGoogleProtobufAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleProtobufAny) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
