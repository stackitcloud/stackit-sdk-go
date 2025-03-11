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

// checks if the GetUpdateServiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUpdateServiceResponse{}

// GetUpdateServiceResponse struct for GetUpdateServiceResponse
type GetUpdateServiceResponse struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGetUpdateServiceResponse instantiates a new GetUpdateServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUpdateServiceResponse() *GetUpdateServiceResponse {
	this := GetUpdateServiceResponse{}
	return &this
}

// NewGetUpdateServiceResponseWithDefaults instantiates a new GetUpdateServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUpdateServiceResponseWithDefaults() *GetUpdateServiceResponse {
	this := GetUpdateServiceResponse{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetUpdateServiceResponse) GetEnabled() *bool {
	if o == nil || IsNil(o.Enabled) {
		var ret *bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUpdateServiceResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetUpdateServiceResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetUpdateServiceResponse) SetEnabled(v *bool) {
	o.Enabled = v
}

func (o GetUpdateServiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableGetUpdateServiceResponse struct {
	value *GetUpdateServiceResponse
	isSet bool
}

func (v NullableGetUpdateServiceResponse) Get() *GetUpdateServiceResponse {
	return v.value
}

func (v *NullableGetUpdateServiceResponse) Set(val *GetUpdateServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUpdateServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUpdateServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUpdateServiceResponse(val *GetUpdateServiceResponse) *NullableGetUpdateServiceResponse {
	return &NullableGetUpdateServiceResponse{value: val, isSet: true}
}

func (v NullableGetUpdateServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUpdateServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
