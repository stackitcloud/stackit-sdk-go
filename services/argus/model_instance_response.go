/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

import (
	"encoding/json"
)

// checks if the InstanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceResponse{}

// InstanceResponse struct for InstanceResponse
type InstanceResponse struct {
	// REQUIRED
	Message *string `json:"message"`
}

type _InstanceResponse InstanceResponse

// NewInstanceResponse instantiates a new InstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceResponse(message *string) *InstanceResponse {
	this := InstanceResponse{}
	this.Message = message
	return &this
}

// NewInstanceResponseWithDefaults instantiates a new InstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceResponseWithDefaults() *InstanceResponse {
	this := InstanceResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *InstanceResponse) GetMessage() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InstanceResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *InstanceResponse) SetMessage(v *string) {
	o.Message = v
}

func (o InstanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableInstanceResponse struct {
	value *InstanceResponse
	isSet bool
}

func (v NullableInstanceResponse) Get() *InstanceResponse {
	return v.value
}

func (v *NullableInstanceResponse) Set(val *InstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceResponse(val *InstanceResponse) *NullableInstanceResponse {
	return &NullableInstanceResponse{value: val, isSet: true}
}

func (v NullableInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
