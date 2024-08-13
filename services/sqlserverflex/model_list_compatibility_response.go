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

// checks if the ListCompatibilityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCompatibilityResponse{}

// ListCompatibilityResponse struct for ListCompatibilityResponse
type ListCompatibilityResponse struct {
	Compatibilities *[]MssqlDatabaseCompatibility `json:"compatibilities,omitempty"`
}

// NewListCompatibilityResponse instantiates a new ListCompatibilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCompatibilityResponse() *ListCompatibilityResponse {
	this := ListCompatibilityResponse{}
	return &this
}

// NewListCompatibilityResponseWithDefaults instantiates a new ListCompatibilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCompatibilityResponseWithDefaults() *ListCompatibilityResponse {
	this := ListCompatibilityResponse{}
	return &this
}

// GetCompatibilities returns the Compatibilities field value if set, zero value otherwise.
func (o *ListCompatibilityResponse) GetCompatibilities() *[]MssqlDatabaseCompatibility {
	if o == nil || IsNil(o.Compatibilities) {
		var ret *[]MssqlDatabaseCompatibility
		return ret
	}
	return o.Compatibilities
}

// GetCompatibilitiesOk returns a tuple with the Compatibilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCompatibilityResponse) GetCompatibilitiesOk() (*[]MssqlDatabaseCompatibility, bool) {
	if o == nil || IsNil(o.Compatibilities) {
		return nil, false
	}
	return o.Compatibilities, true
}

// HasCompatibilities returns a boolean if a field has been set.
func (o *ListCompatibilityResponse) HasCompatibilities() bool {
	if o != nil && !IsNil(o.Compatibilities) {
		return true
	}

	return false
}

// SetCompatibilities gets a reference to the given []MssqlDatabaseCompatibility and assigns it to the Compatibilities field.
func (o *ListCompatibilityResponse) SetCompatibilities(v *[]MssqlDatabaseCompatibility) {
	o.Compatibilities = v
}

func (o ListCompatibilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Compatibilities) {
		toSerialize["compatibilities"] = o.Compatibilities
	}
	return toSerialize, nil
}

type NullableListCompatibilityResponse struct {
	value *ListCompatibilityResponse
	isSet bool
}

func (v NullableListCompatibilityResponse) Get() *ListCompatibilityResponse {
	return v.value
}

func (v *NullableListCompatibilityResponse) Set(val *ListCompatibilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCompatibilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCompatibilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCompatibilityResponse(val *ListCompatibilityResponse) *NullableListCompatibilityResponse {
	return &NullableListCompatibilityResponse{value: val, isSet: true}
}

func (v NullableListCompatibilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCompatibilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
