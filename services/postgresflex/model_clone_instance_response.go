/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

import (
	"encoding/json"
)

// checks if the CloneInstanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloneInstanceResponse{}

// CloneInstanceResponse struct for CloneInstanceResponse
type CloneInstanceResponse struct {
	InstanceId *string `json:"instanceId,omitempty"`
}

// NewCloneInstanceResponse instantiates a new CloneInstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloneInstanceResponse() *CloneInstanceResponse {
	this := CloneInstanceResponse{}
	return &this
}

// NewCloneInstanceResponseWithDefaults instantiates a new CloneInstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloneInstanceResponseWithDefaults() *CloneInstanceResponse {
	this := CloneInstanceResponse{}
	return &this
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *CloneInstanceResponse) GetInstanceId() *string {
	if o == nil || IsNil(o.InstanceId) {
		var ret *string
		return ret
	}
	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloneInstanceResponse) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *CloneInstanceResponse) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *CloneInstanceResponse) SetInstanceId(v *string) {
	o.InstanceId = v
}

func (o CloneInstanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	return toSerialize, nil
}

type NullableCloneInstanceResponse struct {
	value *CloneInstanceResponse
	isSet bool
}

func (v NullableCloneInstanceResponse) Get() *CloneInstanceResponse {
	return v.value
}

func (v *NullableCloneInstanceResponse) Set(val *CloneInstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneInstanceResponse(val *CloneInstanceResponse) *NullableCloneInstanceResponse {
	return &NullableCloneInstanceResponse{value: val, isSet: true}
}

func (v NullableCloneInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
