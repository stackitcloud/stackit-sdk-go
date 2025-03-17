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

// checks if the SecurityGroupRuleProtocol type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupRuleProtocol{}

/*
	types and functions for protocol
*/

// isModel
type SecurityGroupRuleProtocolGetProtocolAttributeType = *Protocol
type SecurityGroupRuleProtocolGetProtocolArgType = Protocol
type SecurityGroupRuleProtocolGetProtocolRetType = Protocol

func getSecurityGroupRuleProtocolGetProtocolAttributeTypeOk(arg SecurityGroupRuleProtocolGetProtocolAttributeType) (ret SecurityGroupRuleProtocolGetProtocolRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSecurityGroupRuleProtocolGetProtocolAttributeType(arg *SecurityGroupRuleProtocolGetProtocolAttributeType, val SecurityGroupRuleProtocolGetProtocolRetType) {
	*arg = &val
}

// SecurityGroupRuleProtocol The internet protocol which the rule matches.
type SecurityGroupRuleProtocol struct {
	Protocol SecurityGroupRuleProtocolGetProtocolAttributeType `json:"protocol,omitempty"`
}

// NewSecurityGroupRuleProtocol instantiates a new SecurityGroupRuleProtocol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupRuleProtocol() *SecurityGroupRuleProtocol {
	this := SecurityGroupRuleProtocol{}
	return &this
}

// NewSecurityGroupRuleProtocolWithDefaults instantiates a new SecurityGroupRuleProtocol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupRuleProtocolWithDefaults() *SecurityGroupRuleProtocol {
	this := SecurityGroupRuleProtocol{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SecurityGroupRuleProtocol) GetProtocol() (res SecurityGroupRuleProtocolGetProtocolRetType) {
	res, _ = o.GetProtocolOk()
	return
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRuleProtocol) GetProtocolOk() (ret SecurityGroupRuleProtocolGetProtocolRetType, ok bool) {
	return getSecurityGroupRuleProtocolGetProtocolAttributeTypeOk(o.Protocol)
}

// HasProtocol returns a boolean if a field has been set.
func (o *SecurityGroupRuleProtocol) HasProtocol() bool {
	_, ok := o.GetProtocolOk()
	return ok
}

// SetProtocol gets a reference to the given Protocol and assigns it to the Protocol field.
func (o *SecurityGroupRuleProtocol) SetProtocol(v SecurityGroupRuleProtocolGetProtocolRetType) {
	setSecurityGroupRuleProtocolGetProtocolAttributeType(&o.Protocol, v)
}

func (o SecurityGroupRuleProtocol) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getSecurityGroupRuleProtocolGetProtocolAttributeTypeOk(o.Protocol); ok {
		toSerialize["Protocol"] = val
	}
	return toSerialize, nil
}

type NullableSecurityGroupRuleProtocol struct {
	value *SecurityGroupRuleProtocol
	isSet bool
}

func (v NullableSecurityGroupRuleProtocol) Get() *SecurityGroupRuleProtocol {
	return v.value
}

func (v *NullableSecurityGroupRuleProtocol) Set(val *SecurityGroupRuleProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupRuleProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupRuleProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupRuleProtocol(val *SecurityGroupRuleProtocol) *NullableSecurityGroupRuleProtocol {
	return &NullableSecurityGroupRuleProtocol{value: val, isSet: true}
}

func (v NullableSecurityGroupRuleProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupRuleProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
