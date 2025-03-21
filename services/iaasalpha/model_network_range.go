/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
	"time"
)

// checks if the NetworkRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkRange{}

/*
	types and functions for createdAt
*/

// isDateTime
type NetworkRangeGetCreatedAtAttributeType = *time.Time
type NetworkRangeGetCreatedAtArgType = time.Time
type NetworkRangeGetCreatedAtRetType = time.Time

func getNetworkRangeGetCreatedAtAttributeTypeOk(arg NetworkRangeGetCreatedAtAttributeType) (ret NetworkRangeGetCreatedAtRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNetworkRangeGetCreatedAtAttributeType(arg *NetworkRangeGetCreatedAtAttributeType, val NetworkRangeGetCreatedAtRetType) {
	*arg = &val
}

/*
	types and functions for networkRangeId
*/

// isNotNullableString
type NetworkRangeGetNetworkRangeIdAttributeType = *string

func getNetworkRangeGetNetworkRangeIdAttributeTypeOk(arg NetworkRangeGetNetworkRangeIdAttributeType) (ret NetworkRangeGetNetworkRangeIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNetworkRangeGetNetworkRangeIdAttributeType(arg *NetworkRangeGetNetworkRangeIdAttributeType, val NetworkRangeGetNetworkRangeIdRetType) {
	*arg = &val
}

type NetworkRangeGetNetworkRangeIdArgType = string
type NetworkRangeGetNetworkRangeIdRetType = string

/*
	types and functions for prefix
*/

// isNotNullableString
type NetworkRangeGetPrefixAttributeType = *string

func getNetworkRangeGetPrefixAttributeTypeOk(arg NetworkRangeGetPrefixAttributeType) (ret NetworkRangeGetPrefixRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNetworkRangeGetPrefixAttributeType(arg *NetworkRangeGetPrefixAttributeType, val NetworkRangeGetPrefixRetType) {
	*arg = &val
}

type NetworkRangeGetPrefixArgType = string
type NetworkRangeGetPrefixRetType = string

/*
	types and functions for updatedAt
*/

// isDateTime
type NetworkRangeGetUpdatedAtAttributeType = *time.Time
type NetworkRangeGetUpdatedAtArgType = time.Time
type NetworkRangeGetUpdatedAtRetType = time.Time

func getNetworkRangeGetUpdatedAtAttributeTypeOk(arg NetworkRangeGetUpdatedAtAttributeType) (ret NetworkRangeGetUpdatedAtRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNetworkRangeGetUpdatedAtAttributeType(arg *NetworkRangeGetUpdatedAtAttributeType, val NetworkRangeGetUpdatedAtRetType) {
	*arg = &val
}

// NetworkRange Object that represents a network range.
type NetworkRange struct {
	// Date-time when resource was created.
	CreatedAt NetworkRangeGetCreatedAtAttributeType `json:"createdAt,omitempty"`
	// Universally Unique Identifier (UUID).
	NetworkRangeId NetworkRangeGetNetworkRangeIdAttributeType `json:"networkRangeId,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	// REQUIRED
	Prefix NetworkRangeGetPrefixAttributeType `json:"prefix"`
	// Date-time when resource was last updated.
	UpdatedAt NetworkRangeGetUpdatedAtAttributeType `json:"updatedAt,omitempty"`
}

type _NetworkRange NetworkRange

// NewNetworkRange instantiates a new NetworkRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRange(prefix NetworkRangeGetPrefixArgType) *NetworkRange {
	this := NetworkRange{}
	setNetworkRangeGetPrefixAttributeType(&this.Prefix, prefix)
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
func (o *NetworkRange) GetCreatedAt() (res NetworkRangeGetCreatedAtRetType) {
	res, _ = o.GetCreatedAtOk()
	return
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRange) GetCreatedAtOk() (ret NetworkRangeGetCreatedAtRetType, ok bool) {
	return getNetworkRangeGetCreatedAtAttributeTypeOk(o.CreatedAt)
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NetworkRange) HasCreatedAt() bool {
	_, ok := o.GetCreatedAtOk()
	return ok
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NetworkRange) SetCreatedAt(v NetworkRangeGetCreatedAtRetType) {
	setNetworkRangeGetCreatedAtAttributeType(&o.CreatedAt, v)
}

// GetNetworkRangeId returns the NetworkRangeId field value if set, zero value otherwise.
func (o *NetworkRange) GetNetworkRangeId() (res NetworkRangeGetNetworkRangeIdRetType) {
	res, _ = o.GetNetworkRangeIdOk()
	return
}

// GetNetworkRangeIdOk returns a tuple with the NetworkRangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRange) GetNetworkRangeIdOk() (ret NetworkRangeGetNetworkRangeIdRetType, ok bool) {
	return getNetworkRangeGetNetworkRangeIdAttributeTypeOk(o.NetworkRangeId)
}

// HasNetworkRangeId returns a boolean if a field has been set.
func (o *NetworkRange) HasNetworkRangeId() bool {
	_, ok := o.GetNetworkRangeIdOk()
	return ok
}

// SetNetworkRangeId gets a reference to the given string and assigns it to the NetworkRangeId field.
func (o *NetworkRange) SetNetworkRangeId(v NetworkRangeGetNetworkRangeIdRetType) {
	setNetworkRangeGetNetworkRangeIdAttributeType(&o.NetworkRangeId, v)
}

// GetPrefix returns the Prefix field value
func (o *NetworkRange) GetPrefix() (ret NetworkRangeGetPrefixRetType) {
	ret, _ = o.GetPrefixOk()
	return ret
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *NetworkRange) GetPrefixOk() (ret NetworkRangeGetPrefixRetType, ok bool) {
	return getNetworkRangeGetPrefixAttributeTypeOk(o.Prefix)
}

// SetPrefix sets field value
func (o *NetworkRange) SetPrefix(v NetworkRangeGetPrefixRetType) {
	setNetworkRangeGetPrefixAttributeType(&o.Prefix, v)
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NetworkRange) GetUpdatedAt() (res NetworkRangeGetUpdatedAtRetType) {
	res, _ = o.GetUpdatedAtOk()
	return
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRange) GetUpdatedAtOk() (ret NetworkRangeGetUpdatedAtRetType, ok bool) {
	return getNetworkRangeGetUpdatedAtAttributeTypeOk(o.UpdatedAt)
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NetworkRange) HasUpdatedAt() bool {
	_, ok := o.GetUpdatedAtOk()
	return ok
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NetworkRange) SetUpdatedAt(v NetworkRangeGetUpdatedAtRetType) {
	setNetworkRangeGetUpdatedAtAttributeType(&o.UpdatedAt, v)
}

func (o NetworkRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getNetworkRangeGetCreatedAtAttributeTypeOk(o.CreatedAt); ok {
		toSerialize["CreatedAt"] = val
	}
	if val, ok := getNetworkRangeGetNetworkRangeIdAttributeTypeOk(o.NetworkRangeId); ok {
		toSerialize["NetworkRangeId"] = val
	}
	if val, ok := getNetworkRangeGetPrefixAttributeTypeOk(o.Prefix); ok {
		toSerialize["Prefix"] = val
	}
	if val, ok := getNetworkRangeGetUpdatedAtAttributeTypeOk(o.UpdatedAt); ok {
		toSerialize["UpdatedAt"] = val
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
