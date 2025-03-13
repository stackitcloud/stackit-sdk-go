/*
STACKIT Server Update Management API

API endpoints for Server Update Operations on STACKIT Servers.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverupdate

import (
	"encoding/json"
)

// checks if the EnableServiceResourcePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableServiceResourcePayload{}

// EnableServiceResourcePayload struct for EnableServiceResourcePayload
type EnableServiceResourcePayload struct {
	UpdatePolicyId *string `json:"updatePolicyId,omitempty"`
}

// NewEnableServiceResourcePayload instantiates a new EnableServiceResourcePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableServiceResourcePayload() *EnableServiceResourcePayload {
	this := EnableServiceResourcePayload{}
	return &this
}

// NewEnableServiceResourcePayloadWithDefaults instantiates a new EnableServiceResourcePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableServiceResourcePayloadWithDefaults() *EnableServiceResourcePayload {
	this := EnableServiceResourcePayload{}
	return &this
}

// GetUpdatePolicyId returns the UpdatePolicyId field value if set, zero value otherwise.
func (o *EnableServiceResourcePayload) GetUpdatePolicyId() *string {
	if o == nil || IsNil(o.UpdatePolicyId) {
		var ret *string
		return ret
	}
	return o.UpdatePolicyId
}

// GetUpdatePolicyIdOk returns a tuple with the UpdatePolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableServiceResourcePayload) GetUpdatePolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatePolicyId) {
		return nil, false
	}
	return o.UpdatePolicyId, true
}

// HasUpdatePolicyId returns a boolean if a field has been set.
func (o *EnableServiceResourcePayload) HasUpdatePolicyId() bool {
	if o != nil && !IsNil(o.UpdatePolicyId) {
		return true
	}

	return false
}

// SetUpdatePolicyId gets a reference to the given string and assigns it to the UpdatePolicyId field.
func (o *EnableServiceResourcePayload) SetUpdatePolicyId(v *string) {
	o.UpdatePolicyId = v
}

func (o EnableServiceResourcePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpdatePolicyId) {
		toSerialize["updatePolicyId"] = o.UpdatePolicyId
	}
	return toSerialize, nil
}

type NullableEnableServiceResourcePayload struct {
	value *EnableServiceResourcePayload
	isSet bool
}

func (v NullableEnableServiceResourcePayload) Get() *EnableServiceResourcePayload {
	return v.value
}

func (v *NullableEnableServiceResourcePayload) Set(val *EnableServiceResourcePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableServiceResourcePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableServiceResourcePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableServiceResourcePayload(val *EnableServiceResourcePayload) *NullableEnableServiceResourcePayload {
	return &NullableEnableServiceResourcePayload{value: val, isSet: true}
}

func (v NullableEnableServiceResourcePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableServiceResourcePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
