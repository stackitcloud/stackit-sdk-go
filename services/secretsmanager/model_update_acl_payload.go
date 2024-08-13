/*
STACKIT Secrets Manager API

This API provides endpoints for managing the Secrets-Manager.

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretsmanager

import (
	"encoding/json"
)

// checks if the UpdateACLPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateACLPayload{}

// UpdateACLPayload struct for UpdateACLPayload
type UpdateACLPayload struct {
	// The given IP/IP Range that is permitted to access.
	// REQUIRED
	Cidr *string `json:"cidr"`
}

type _UpdateACLPayload UpdateACLPayload

// NewUpdateACLPayload instantiates a new UpdateACLPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateACLPayload(cidr *string) *UpdateACLPayload {
	this := UpdateACLPayload{}
	this.Cidr = cidr
	return &this
}

// NewUpdateACLPayloadWithDefaults instantiates a new UpdateACLPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateACLPayloadWithDefaults() *UpdateACLPayload {
	this := UpdateACLPayload{}
	return &this
}

// GetCidr returns the Cidr field value
func (o *UpdateACLPayload) GetCidr() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *UpdateACLPayload) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cidr, true
}

// SetCidr sets field value
func (o *UpdateACLPayload) SetCidr(v *string) {
	o.Cidr = v
}

func (o UpdateACLPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cidr"] = o.Cidr
	return toSerialize, nil
}

type NullableUpdateACLPayload struct {
	value *UpdateACLPayload
	isSet bool
}

func (v NullableUpdateACLPayload) Get() *UpdateACLPayload {
	return v.value
}

func (v *NullableUpdateACLPayload) Set(val *UpdateACLPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateACLPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateACLPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateACLPayload(val *UpdateACLPayload) *NullableUpdateACLPayload {
	return &NullableUpdateACLPayload{value: val, isSet: true}
}

func (v NullableUpdateACLPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateACLPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
