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

// checks if the LoadbalancerOptionAccessControl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadbalancerOptionAccessControl{}

/*
	types and functions for allowedSourceRanges
*/

// isArray
type LoadbalancerOptionAccessControlGetAllowedSourceRangesAttributeType = *[]string
type LoadbalancerOptionAccessControlGetAllowedSourceRangesArgType = []string
type LoadbalancerOptionAccessControlGetAllowedSourceRangesRetType = []string

func getLoadbalancerOptionAccessControlGetAllowedSourceRangesAttributeTypeOk(arg LoadbalancerOptionAccessControlGetAllowedSourceRangesAttributeType) (ret LoadbalancerOptionAccessControlGetAllowedSourceRangesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setLoadbalancerOptionAccessControlGetAllowedSourceRangesAttributeType(arg *LoadbalancerOptionAccessControlGetAllowedSourceRangesAttributeType, val LoadbalancerOptionAccessControlGetAllowedSourceRangesRetType) {
	*arg = &val
}

// LoadbalancerOptionAccessControl Use this option to limit the IP ranges that can use the application load balancer.
type LoadbalancerOptionAccessControl struct {
	// Application Load Balancer is accessible only from an IP address in this range
	AllowedSourceRanges LoadbalancerOptionAccessControlGetAllowedSourceRangesAttributeType `json:"allowedSourceRanges,omitempty"`
}

// NewLoadbalancerOptionAccessControl instantiates a new LoadbalancerOptionAccessControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadbalancerOptionAccessControl() *LoadbalancerOptionAccessControl {
	this := LoadbalancerOptionAccessControl{}
	return &this
}

// NewLoadbalancerOptionAccessControlWithDefaults instantiates a new LoadbalancerOptionAccessControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadbalancerOptionAccessControlWithDefaults() *LoadbalancerOptionAccessControl {
	this := LoadbalancerOptionAccessControl{}
	return &this
}

// GetAllowedSourceRanges returns the AllowedSourceRanges field value if set, zero value otherwise.
func (o *LoadbalancerOptionAccessControl) GetAllowedSourceRanges() (res LoadbalancerOptionAccessControlGetAllowedSourceRangesRetType) {
	res, _ = o.GetAllowedSourceRangesOk()
	return
}

// GetAllowedSourceRangesOk returns a tuple with the AllowedSourceRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadbalancerOptionAccessControl) GetAllowedSourceRangesOk() (ret LoadbalancerOptionAccessControlGetAllowedSourceRangesRetType, ok bool) {
	return getLoadbalancerOptionAccessControlGetAllowedSourceRangesAttributeTypeOk(o.AllowedSourceRanges)
}

// HasAllowedSourceRanges returns a boolean if a field has been set.
func (o *LoadbalancerOptionAccessControl) HasAllowedSourceRanges() bool {
	_, ok := o.GetAllowedSourceRangesOk()
	return ok
}

// SetAllowedSourceRanges gets a reference to the given []string and assigns it to the AllowedSourceRanges field.
func (o *LoadbalancerOptionAccessControl) SetAllowedSourceRanges(v LoadbalancerOptionAccessControlGetAllowedSourceRangesRetType) {
	setLoadbalancerOptionAccessControlGetAllowedSourceRangesAttributeType(&o.AllowedSourceRanges, v)
}

func (o LoadbalancerOptionAccessControl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getLoadbalancerOptionAccessControlGetAllowedSourceRangesAttributeTypeOk(o.AllowedSourceRanges); ok {
		toSerialize["AllowedSourceRanges"] = val
	}
	return toSerialize, nil
}

type NullableLoadbalancerOptionAccessControl struct {
	value *LoadbalancerOptionAccessControl
	isSet bool
}

func (v NullableLoadbalancerOptionAccessControl) Get() *LoadbalancerOptionAccessControl {
	return v.value
}

func (v *NullableLoadbalancerOptionAccessControl) Set(val *LoadbalancerOptionAccessControl) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadbalancerOptionAccessControl) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadbalancerOptionAccessControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadbalancerOptionAccessControl(val *LoadbalancerOptionAccessControl) *NullableLoadbalancerOptionAccessControl {
	return &NullableLoadbalancerOptionAccessControl{value: val, isSet: true}
}

func (v NullableLoadbalancerOptionAccessControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadbalancerOptionAccessControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
