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

// checks if the AddMemberToVirtualIPPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddMemberToVirtualIPPayload{}

// AddMemberToVirtualIPPayload Object that represents a virtual IP member.
type AddMemberToVirtualIPPayload struct {
	// Universally Unique Identifier (UUID).
	// REQUIRED
	Member *string `json:"member"`
}

type _AddMemberToVirtualIPPayload AddMemberToVirtualIPPayload

// NewAddMemberToVirtualIPPayload instantiates a new AddMemberToVirtualIPPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddMemberToVirtualIPPayload(member *string) *AddMemberToVirtualIPPayload {
	this := AddMemberToVirtualIPPayload{}
	this.Member = member
	return &this
}

// NewAddMemberToVirtualIPPayloadWithDefaults instantiates a new AddMemberToVirtualIPPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddMemberToVirtualIPPayloadWithDefaults() *AddMemberToVirtualIPPayload {
	this := AddMemberToVirtualIPPayload{}
	return &this
}

// GetMember returns the Member field value
func (o *AddMemberToVirtualIPPayload) GetMember() *string {
	if o == nil || IsNil(o.Member) {
		var ret *string
		return ret
	}

	return o.Member
}

// GetMemberOk returns a tuple with the Member field value
// and a boolean to check if the value has been set.
func (o *AddMemberToVirtualIPPayload) GetMemberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Member, true
}

// SetMember sets field value
func (o *AddMemberToVirtualIPPayload) SetMember(v *string) {
	o.Member = v
}

func (o AddMemberToVirtualIPPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["member"] = o.Member
	return toSerialize, nil
}

type NullableAddMemberToVirtualIPPayload struct {
	value *AddMemberToVirtualIPPayload
	isSet bool
}

func (v NullableAddMemberToVirtualIPPayload) Get() *AddMemberToVirtualIPPayload {
	return v.value
}

func (v *NullableAddMemberToVirtualIPPayload) Set(val *AddMemberToVirtualIPPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAddMemberToVirtualIPPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAddMemberToVirtualIPPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddMemberToVirtualIPPayload(val *AddMemberToVirtualIPPayload) *NullableAddMemberToVirtualIPPayload {
	return &NullableAddMemberToVirtualIPPayload{value: val, isSet: true}
}

func (v NullableAddMemberToVirtualIPPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddMemberToVirtualIPPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
