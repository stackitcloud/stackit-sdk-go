/*
STACKIT Marketplace API

API to manage STACKIT Marketplace.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackitmarketplace

import (
	"encoding/json"
)

// checks if the ApproveSubscriptionPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApproveSubscriptionPayload{}

// ApproveSubscriptionPayload struct for ApproveSubscriptionPayload
type ApproveSubscriptionPayload struct {
	// The target URL of the user instance, used to redirect the user to the instance after the subscription is active.
	InstanceTarget *string `json:"instanceTarget,omitempty"`
}

// NewApproveSubscriptionPayload instantiates a new ApproveSubscriptionPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproveSubscriptionPayload() *ApproveSubscriptionPayload {
	this := ApproveSubscriptionPayload{}
	return &this
}

// NewApproveSubscriptionPayloadWithDefaults instantiates a new ApproveSubscriptionPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApproveSubscriptionPayloadWithDefaults() *ApproveSubscriptionPayload {
	this := ApproveSubscriptionPayload{}
	return &this
}

// GetInstanceTarget returns the InstanceTarget field value if set, zero value otherwise.
func (o *ApproveSubscriptionPayload) GetInstanceTarget() *string {
	if o == nil || IsNil(o.InstanceTarget) {
		var ret *string
		return ret
	}
	return o.InstanceTarget
}

// GetInstanceTargetOk returns a tuple with the InstanceTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApproveSubscriptionPayload) GetInstanceTargetOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceTarget) {
		return nil, false
	}
	return o.InstanceTarget, true
}

// HasInstanceTarget returns a boolean if a field has been set.
func (o *ApproveSubscriptionPayload) HasInstanceTarget() bool {
	if o != nil && !IsNil(o.InstanceTarget) {
		return true
	}

	return false
}

// SetInstanceTarget gets a reference to the given string and assigns it to the InstanceTarget field.
func (o *ApproveSubscriptionPayload) SetInstanceTarget(v *string) {
	o.InstanceTarget = v
}

func (o ApproveSubscriptionPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceTarget) {
		toSerialize["instanceTarget"] = o.InstanceTarget
	}
	return toSerialize, nil
}

type NullableApproveSubscriptionPayload struct {
	value *ApproveSubscriptionPayload
	isSet bool
}

func (v NullableApproveSubscriptionPayload) Get() *ApproveSubscriptionPayload {
	return v.value
}

func (v *NullableApproveSubscriptionPayload) Set(val *ApproveSubscriptionPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableApproveSubscriptionPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableApproveSubscriptionPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproveSubscriptionPayload(val *ApproveSubscriptionPayload) *NullableApproveSubscriptionPayload {
	return &NullableApproveSubscriptionPayload{value: val, isSet: true}
}

func (v NullableApproveSubscriptionPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproveSubscriptionPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
