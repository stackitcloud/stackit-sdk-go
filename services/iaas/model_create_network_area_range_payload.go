/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the CreateNetworkAreaRangePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkAreaRangePayload{}

// CreateNetworkAreaRangePayload struct for CreateNetworkAreaRangePayload
type CreateNetworkAreaRangePayload struct {
	// A list of network ranges.
	Ipv4 *[]NetworkRange `json:"ipv4,omitempty"`
}

// NewCreateNetworkAreaRangePayload instantiates a new CreateNetworkAreaRangePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkAreaRangePayload() *CreateNetworkAreaRangePayload {
	this := CreateNetworkAreaRangePayload{}
	return &this
}

// NewCreateNetworkAreaRangePayloadWithDefaults instantiates a new CreateNetworkAreaRangePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkAreaRangePayloadWithDefaults() *CreateNetworkAreaRangePayload {
	this := CreateNetworkAreaRangePayload{}
	return &this
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *CreateNetworkAreaRangePayload) GetIpv4() *[]NetworkRange {
	if o == nil || IsNil(o.Ipv4) {
		var ret *[]NetworkRange
		return ret
	}
	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkAreaRangePayload) GetIpv4Ok() (*[]NetworkRange, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *CreateNetworkAreaRangePayload) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given []NetworkRange and assigns it to the Ipv4 field.
func (o *CreateNetworkAreaRangePayload) SetIpv4(v *[]NetworkRange) {
	o.Ipv4 = v
}

func (o CreateNetworkAreaRangePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	return toSerialize, nil
}

type NullableCreateNetworkAreaRangePayload struct {
	value *CreateNetworkAreaRangePayload
	isSet bool
}

func (v NullableCreateNetworkAreaRangePayload) Get() *CreateNetworkAreaRangePayload {
	return v.value
}

func (v *NullableCreateNetworkAreaRangePayload) Set(val *CreateNetworkAreaRangePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkAreaRangePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkAreaRangePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkAreaRangePayload(val *CreateNetworkAreaRangePayload) *NullableCreateNetworkAreaRangePayload {
	return &NullableCreateNetworkAreaRangePayload{value: val, isSet: true}
}

func (v NullableCreateNetworkAreaRangePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkAreaRangePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
