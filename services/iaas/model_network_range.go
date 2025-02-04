/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
	"time"
)

// checks if the NetworkRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkRange{}

// NetworkRange Object that represents a network range.
type NetworkRange struct {
	// Date-time when resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Universally Unique Identifier (UUID).
	NetworkRangeId *string `json:"networkRangeId,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	// REQUIRED
	Prefix *string `json:"prefix"`
	// Date-time when resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type _NetworkRange NetworkRange

// NewNetworkRange instantiates a new NetworkRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRange(prefix *string) *NetworkRange {
	this := NetworkRange{}
	this.Prefix = prefix
	return &this
}

// NewNetworkRangeWithDefaults instantiates a new NetworkRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkRangeWithDefaults() *NetworkRange {
	this := NetworkRange{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NetworkRange) GetCreatedAt() *time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret *time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRange) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NetworkRange) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NetworkRange) SetCreatedAt(v *time.Time) {
	o.CreatedAt = v
}

// GetNetworkRangeId returns the NetworkRangeId field value if set, zero value otherwise.
func (o *NetworkRange) GetNetworkRangeId() *string {
	if o == nil || IsNil(o.NetworkRangeId) {
		var ret *string
		return ret
	}
	return o.NetworkRangeId
}

// GetNetworkRangeIdOk returns a tuple with the NetworkRangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRange) GetNetworkRangeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkRangeId) {
		return nil, false
	}
	return o.NetworkRangeId, true
}

// HasNetworkRangeId returns a boolean if a field has been set.
func (o *NetworkRange) HasNetworkRangeId() bool {
	if o != nil && !IsNil(o.NetworkRangeId) {
		return true
	}

	return false
}

// SetNetworkRangeId gets a reference to the given string and assigns it to the NetworkRangeId field.
func (o *NetworkRange) SetNetworkRangeId(v *string) {
	o.NetworkRangeId = v
}

// GetPrefix returns the Prefix field value
func (o *NetworkRange) GetPrefix() *string {
	if o == nil || IsNil(o.Prefix) {
		var ret *string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *NetworkRange) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix, true
}

// SetPrefix sets field value
func (o *NetworkRange) SetPrefix(v *string) {
	o.Prefix = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NetworkRange) GetUpdatedAt() *time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret *time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRange) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NetworkRange) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NetworkRange) SetUpdatedAt(v *time.Time) {
	o.UpdatedAt = v
}

func (o NetworkRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.NetworkRangeId) {
		toSerialize["networkRangeId"] = o.NetworkRangeId
	}
	toSerialize["prefix"] = o.Prefix
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableNetworkRange struct {
	value *NetworkRange
	isSet bool
}

func (v NullableNetworkRange) Get() *NetworkRange {
	return v.value
}

func (v *NullableNetworkRange) Set(val *NetworkRange) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkRange) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkRange(val *NetworkRange) *NullableNetworkRange {
	return &NullableNetworkRange{value: val, isSet: true}
}

func (v NullableNetworkRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
