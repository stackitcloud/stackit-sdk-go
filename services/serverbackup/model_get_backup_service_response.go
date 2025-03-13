/*
STACKIT Server Backup Management API

API endpoints for Server Backup Operations on STACKIT Servers.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverbackup

import (
	"encoding/json"
)

// checks if the GetBackupServiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBackupServiceResponse{}

// GetBackupServiceResponse struct for GetBackupServiceResponse
type GetBackupServiceResponse struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGetBackupServiceResponse instantiates a new GetBackupServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBackupServiceResponse() *GetBackupServiceResponse {
	this := GetBackupServiceResponse{}
	return &this
}

// NewGetBackupServiceResponseWithDefaults instantiates a new GetBackupServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBackupServiceResponseWithDefaults() *GetBackupServiceResponse {
	this := GetBackupServiceResponse{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetBackupServiceResponse) GetEnabled() *bool {
	if o == nil || IsNil(o.Enabled) {
		var ret *bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBackupServiceResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetBackupServiceResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetBackupServiceResponse) SetEnabled(v *bool) {
	o.Enabled = v
}

func (o GetBackupServiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableGetBackupServiceResponse struct {
	value *GetBackupServiceResponse
	isSet bool
}

func (v NullableGetBackupServiceResponse) Get() *GetBackupServiceResponse {
	return v.value
}

func (v *NullableGetBackupServiceResponse) Set(val *GetBackupServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBackupServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBackupServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBackupServiceResponse(val *GetBackupServiceResponse) *NullableGetBackupServiceResponse {
	return &NullableGetBackupServiceResponse{value: val, isSet: true}
}

func (v NullableGetBackupServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBackupServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
