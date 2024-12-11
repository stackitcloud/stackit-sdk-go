/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.7.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

import (
	"encoding/json"
)

// checks if the OptionsUDP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionsUDP{}

// OptionsUDP struct for OptionsUDP
type OptionsUDP struct {
	// The connection idle timeout to be used with the protocol. The default value is set to 1 minute, and the maximum value is 2 minutes.
	IdleTimeout *string `json:"idleTimeout,omitempty"`
}

// NewOptionsUDP instantiates a new OptionsUDP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionsUDP() *OptionsUDP {
	this := OptionsUDP{}
	return &this
}

// NewOptionsUDPWithDefaults instantiates a new OptionsUDP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionsUDPWithDefaults() *OptionsUDP {
	this := OptionsUDP{}
	return &this
}

// GetIdleTimeout returns the IdleTimeout field value if set, zero value otherwise.
func (o *OptionsUDP) GetIdleTimeout() *string {
	if o == nil || IsNil(o.IdleTimeout) {
		var ret *string
		return ret
	}
	return o.IdleTimeout
}

// GetIdleTimeoutOk returns a tuple with the IdleTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsUDP) GetIdleTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.IdleTimeout) {
		return nil, false
	}
	return o.IdleTimeout, true
}

// HasIdleTimeout returns a boolean if a field has been set.
func (o *OptionsUDP) HasIdleTimeout() bool {
	if o != nil && !IsNil(o.IdleTimeout) && !IsNil(o.IdleTimeout) {
		return true
	}

	return false
}

// SetIdleTimeout gets a reference to the given string and assigns it to the IdleTimeout field.
func (o *OptionsUDP) SetIdleTimeout(v *string) {
	o.IdleTimeout = v
}

func (o OptionsUDP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdleTimeout) {
		toSerialize["idleTimeout"] = o.IdleTimeout
	}
	return toSerialize, nil
}

type NullableOptionsUDP struct {
	value *OptionsUDP
	isSet bool
}

func (v NullableOptionsUDP) Get() *OptionsUDP {
	return v.value
}

func (v *NullableOptionsUDP) Set(val *OptionsUDP) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionsUDP) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionsUDP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionsUDP(val *OptionsUDP) *NullableOptionsUDP {
	return &NullableOptionsUDP{value: val, isSet: true}
}

func (v NullableOptionsUDP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionsUDP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
