/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

import (
	"encoding/json"
)

// checks if the SessionPersistence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionPersistence{}

/*
	types and functions for useSourceIpAddress
*/

// isBoolean
type SessionPersistencegetUseSourceIpAddressAttributeType = *bool
type SessionPersistencegetUseSourceIpAddressArgType = bool
type SessionPersistencegetUseSourceIpAddressRetType = bool

func getSessionPersistencegetUseSourceIpAddressAttributeTypeOk(arg SessionPersistencegetUseSourceIpAddressAttributeType) (ret SessionPersistencegetUseSourceIpAddressRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSessionPersistencegetUseSourceIpAddressAttributeType(arg *SessionPersistencegetUseSourceIpAddressAttributeType, val SessionPersistencegetUseSourceIpAddressRetType) {
	*arg = &val
}

// SessionPersistence struct for SessionPersistence
type SessionPersistence struct {
	// If enabled then all connections from one source IP address are redirected to the same target. This setting changes the load balancing algorithm to Maglev.
	UseSourceIpAddress SessionPersistencegetUseSourceIpAddressAttributeType `json:"useSourceIpAddress,omitempty"`
}

// NewSessionPersistence instantiates a new SessionPersistence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionPersistence() *SessionPersistence {
	this := SessionPersistence{}
	return &this
}

// NewSessionPersistenceWithDefaults instantiates a new SessionPersistence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionPersistenceWithDefaults() *SessionPersistence {
	this := SessionPersistence{}
	return &this
}

// GetUseSourceIpAddress returns the UseSourceIpAddress field value if set, zero value otherwise.
func (o *SessionPersistence) GetUseSourceIpAddress() (res SessionPersistencegetUseSourceIpAddressRetType) {
	res, _ = o.GetUseSourceIpAddressOk()
	return
}

// GetUseSourceIpAddressOk returns a tuple with the UseSourceIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionPersistence) GetUseSourceIpAddressOk() (ret SessionPersistencegetUseSourceIpAddressRetType, ok bool) {
	return getSessionPersistencegetUseSourceIpAddressAttributeTypeOk(o.UseSourceIpAddress)
}

// HasUseSourceIpAddress returns a boolean if a field has been set.
func (o *SessionPersistence) HasUseSourceIpAddress() bool {
	_, ok := o.GetUseSourceIpAddressOk()
	return ok
}

// SetUseSourceIpAddress gets a reference to the given bool and assigns it to the UseSourceIpAddress field.
func (o *SessionPersistence) SetUseSourceIpAddress(v SessionPersistencegetUseSourceIpAddressRetType) {
	setSessionPersistencegetUseSourceIpAddressAttributeType(&o.UseSourceIpAddress, v)
}

func (o SessionPersistence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getSessionPersistencegetUseSourceIpAddressAttributeTypeOk(o.UseSourceIpAddress); ok {
		toSerialize["UseSourceIpAddress"] = val
	}
	return toSerialize, nil
}

type NullableSessionPersistence struct {
	value *SessionPersistence
	isSet bool
}

func (v NullableSessionPersistence) Get() *SessionPersistence {
	return v.value
}

func (v *NullableSessionPersistence) Set(val *SessionPersistence) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionPersistence) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionPersistence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionPersistence(val *SessionPersistence) *NullableSessionPersistence {
	return &NullableSessionPersistence{value: val, isSet: true}
}

func (v NullableSessionPersistence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionPersistence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
