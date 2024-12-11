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

// checks if the ServerNameIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerNameIndicator{}

// ServerNameIndicator struct for ServerNameIndicator
type ServerNameIndicator struct {
	// The domain name for this SNI config.
	Name *string `json:"name,omitempty"`
}

// NewServerNameIndicator instantiates a new ServerNameIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerNameIndicator() *ServerNameIndicator {
	this := ServerNameIndicator{}
	return &this
}

// NewServerNameIndicatorWithDefaults instantiates a new ServerNameIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerNameIndicatorWithDefaults() *ServerNameIndicator {
	this := ServerNameIndicator{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServerNameIndicator) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerNameIndicator) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServerNameIndicator) HasName() bool {
	if o != nil && !IsNil(o.Name) && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServerNameIndicator) SetName(v *string) {
	o.Name = v
}

func (o ServerNameIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableServerNameIndicator struct {
	value *ServerNameIndicator
	isSet bool
}

func (v NullableServerNameIndicator) Get() *ServerNameIndicator {
	return v.value
}

func (v *NullableServerNameIndicator) Set(val *ServerNameIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableServerNameIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableServerNameIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerNameIndicator(val *ServerNameIndicator) *NullableServerNameIndicator {
	return &NullableServerNameIndicator{value: val, isSet: true}
}

func (v NullableServerNameIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerNameIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
