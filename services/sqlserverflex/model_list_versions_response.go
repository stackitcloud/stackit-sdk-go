/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
)

// checks if the ListVersionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListVersionsResponse{}

// ListVersionsResponse struct for ListVersionsResponse
type ListVersionsResponse struct {
	Versions *[]string `json:"versions,omitempty"`
}

// NewListVersionsResponse instantiates a new ListVersionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVersionsResponse() *ListVersionsResponse {
	this := ListVersionsResponse{}
	return &this
}

// NewListVersionsResponseWithDefaults instantiates a new ListVersionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVersionsResponseWithDefaults() *ListVersionsResponse {
	this := ListVersionsResponse{}
	return &this
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ListVersionsResponse) GetVersions() *[]string {
	if o == nil || IsNil(o.Versions) {
		var ret *[]string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVersionsResponse) GetVersionsOk() (*[]string, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ListVersionsResponse) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *ListVersionsResponse) SetVersions(v *[]string) {
	o.Versions = v
}

func (o ListVersionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableListVersionsResponse struct {
	value *ListVersionsResponse
	isSet bool
}

func (v NullableListVersionsResponse) Get() *ListVersionsResponse {
	return v.value
}

func (v *NullableListVersionsResponse) Set(val *ListVersionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListVersionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListVersionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVersionsResponse(val *ListVersionsResponse) *NullableListVersionsResponse {
	return &NullableListVersionsResponse{value: val, isSet: true}
}

func (v NullableListVersionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVersionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
