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

// checks if the ApiInstalledListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInstalledListResponse{}

// ApiInstalledListResponse struct for ApiInstalledListResponse
type ApiInstalledListResponse struct {
	Installed *[]ApiExtensionList `json:"installed,omitempty"`
}

// NewApiInstalledListResponse instantiates a new ApiInstalledListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInstalledListResponse() *ApiInstalledListResponse {
	this := ApiInstalledListResponse{}
	return &this
}

// NewApiInstalledListResponseWithDefaults instantiates a new ApiInstalledListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInstalledListResponseWithDefaults() *ApiInstalledListResponse {
	this := ApiInstalledListResponse{}
	return &this
}

// GetInstalled returns the Installed field value if set, zero value otherwise.
func (o *ApiInstalledListResponse) GetInstalled() *[]ApiExtensionList {
	if o == nil || IsNil(o.Installed) {
		var ret *[]ApiExtensionList
		return ret
	}
	return o.Installed
}

// GetInstalledOk returns a tuple with the Installed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstalledListResponse) GetInstalledOk() (*[]ApiExtensionList, bool) {
	if o == nil || IsNil(o.Installed) {
		return nil, false
	}
	return o.Installed, true
}

// HasInstalled returns a boolean if a field has been set.
func (o *ApiInstalledListResponse) HasInstalled() bool {
	if o != nil && !IsNil(o.Installed) {
		return true
	}

	return false
}

// SetInstalled gets a reference to the given []ApiExtensionList and assigns it to the Installed field.
func (o *ApiInstalledListResponse) SetInstalled(v *[]ApiExtensionList) {
	o.Installed = v
}

func (o ApiInstalledListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Installed) {
		toSerialize["installed"] = o.Installed
	}
	return toSerialize, nil
}

type NullableApiInstalledListResponse struct {
	value *ApiInstalledListResponse
	isSet bool
}

func (v NullableApiInstalledListResponse) Get() *ApiInstalledListResponse {
	return v.value
}

func (v *NullableApiInstalledListResponse) Set(val *ApiInstalledListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInstalledListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInstalledListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInstalledListResponse(val *ApiInstalledListResponse) *NullableApiInstalledListResponse {
	return &NullableApiInstalledListResponse{value: val, isSet: true}
}

func (v NullableApiInstalledListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInstalledListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
