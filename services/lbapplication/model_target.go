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

// checks if the Target type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Target{}

/*
	types and functions for displayName
*/

// isNotNullableString
type TargetGetDisplayNameAttributeType = *string

func getTargetGetDisplayNameAttributeTypeOk(arg TargetGetDisplayNameAttributeType) (ret TargetGetDisplayNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setTargetGetDisplayNameAttributeType(arg *TargetGetDisplayNameAttributeType, val TargetGetDisplayNameRetType) {
	*arg = &val
}

type TargetGetDisplayNameArgType = string
type TargetGetDisplayNameRetType = string

/*
	types and functions for ip
*/

// isNotNullableString
type TargetGetIpAttributeType = *string

func getTargetGetIpAttributeTypeOk(arg TargetGetIpAttributeType) (ret TargetGetIpRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setTargetGetIpAttributeType(arg *TargetGetIpAttributeType, val TargetGetIpRetType) {
	*arg = &val
}

type TargetGetIpArgType = string
type TargetGetIpRetType = string

// Target struct for Target
type Target struct {
	// Target name
	DisplayName TargetGetDisplayNameAttributeType `json:"displayName,omitempty"`
	// Target IP. Must by unique within a target pool.
	Ip TargetGetIpAttributeType `json:"ip,omitempty"`
}

// NewTarget instantiates a new Target object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTarget() *Target {
	this := Target{}
	return &this
}

// NewTargetWithDefaults instantiates a new Target object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetWithDefaults() *Target {
	this := Target{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Target) GetDisplayName() (res TargetGetDisplayNameRetType) {
	res, _ = o.GetDisplayNameOk()
	return
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetDisplayNameOk() (ret TargetGetDisplayNameRetType, ok bool) {
	return getTargetGetDisplayNameAttributeTypeOk(o.DisplayName)
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Target) HasDisplayName() bool {
	_, ok := o.GetDisplayNameOk()
	return ok
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Target) SetDisplayName(v TargetGetDisplayNameRetType) {
	setTargetGetDisplayNameAttributeType(&o.DisplayName, v)
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *Target) GetIp() (res TargetGetIpRetType) {
	res, _ = o.GetIpOk()
	return
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetIpOk() (ret TargetGetIpRetType, ok bool) {
	return getTargetGetIpAttributeTypeOk(o.Ip)
}

// HasIp returns a boolean if a field has been set.
func (o *Target) HasIp() bool {
	_, ok := o.GetIpOk()
	return ok
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *Target) SetIp(v TargetGetIpRetType) {
	setTargetGetIpAttributeType(&o.Ip, v)
}

func (o Target) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getTargetGetDisplayNameAttributeTypeOk(o.DisplayName); ok {
		toSerialize["DisplayName"] = val
	}
	if val, ok := getTargetGetIpAttributeTypeOk(o.Ip); ok {
		toSerialize["Ip"] = val
	}
	return toSerialize, nil
}

type NullableTarget struct {
	value *Target
	isSet bool
}

func (v NullableTarget) Get() *Target {
	return v.value
}

func (v *NullableTarget) Set(val *Target) {
	v.value = val
	v.isSet = true
}

func (v NullableTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTarget(val *Target) *NullableTarget {
	return &NullableTarget{value: val, isSet: true}
}

func (v NullableTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
