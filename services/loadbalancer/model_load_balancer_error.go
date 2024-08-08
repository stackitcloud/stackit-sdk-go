/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.7.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

import (
	"encoding/json"
)

// checks if the LoadBalancerError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerError{}

// LoadBalancerError struct for LoadBalancerError
type LoadBalancerError struct {
	// The error description contains additional helpful user information to fix the error state of the load balancer. For example the IP 45.135.247.139 does not exist in the project, then the description will report: Floating IP \"45.135.247.139\" could not be found.
	Description *string `json:"description,omitempty"`
	// The error type specifies which part of the load balancer encountered the error. I.e. the API will not check if a provided public IP is actually available in the project. Instead the load balancer with try to use the provided IP and if not available reports TYPE_FIP_NOT_CONFIGURED error.
	Type *string `json:"type,omitempty"`
}

// NewLoadBalancerError instantiates a new LoadBalancerError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerError() *LoadBalancerError {
	this := LoadBalancerError{}
	return &this
}

// NewLoadBalancerErrorWithDefaults instantiates a new LoadBalancerError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerErrorWithDefaults() *LoadBalancerError {
	this := LoadBalancerError{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LoadBalancerError) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerError) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LoadBalancerError) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LoadBalancerError) SetDescription(v *string) {
	o.Description = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LoadBalancerError) GetType() *string {
	if o == nil || IsNil(o.Type) {
		var ret *string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerError) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LoadBalancerError) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LoadBalancerError) SetType(v *string) {
	o.Type = v
}

func (o LoadBalancerError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableLoadBalancerError struct {
	value *LoadBalancerError
	isSet bool
}

func (v NullableLoadBalancerError) Get() *LoadBalancerError {
	return v.value
}

func (v *NullableLoadBalancerError) Set(val *LoadBalancerError) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerError) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerError(val *LoadBalancerError) *NullableLoadBalancerError {
	return &NullableLoadBalancerError{value: val, isSet: true}
}

func (v NullableLoadBalancerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
