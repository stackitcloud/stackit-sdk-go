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

// checks if the ListCollationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCollationsResponse{}

// ListCollationsResponse struct for ListCollationsResponse
type ListCollationsResponse struct {
	Collations *[]MssqlDatabaseCollation `json:"collations,omitempty"`
}

// NewListCollationsResponse instantiates a new ListCollationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCollationsResponse() *ListCollationsResponse {
	this := ListCollationsResponse{}
	return &this
}

// NewListCollationsResponseWithDefaults instantiates a new ListCollationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCollationsResponseWithDefaults() *ListCollationsResponse {
	this := ListCollationsResponse{}
	return &this
}

// GetCollations returns the Collations field value if set, zero value otherwise.
func (o *ListCollationsResponse) GetCollations() *[]MssqlDatabaseCollation {
	if o == nil || IsNil(o.Collations) {
		var ret *[]MssqlDatabaseCollation
		return ret
	}
	return o.Collations
}

// GetCollationsOk returns a tuple with the Collations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCollationsResponse) GetCollationsOk() (*[]MssqlDatabaseCollation, bool) {
	if o == nil || IsNil(o.Collations) {
		return nil, false
	}
	return o.Collations, true
}

// HasCollations returns a boolean if a field has been set.
func (o *ListCollationsResponse) HasCollations() bool {
	if o != nil && !IsNil(o.Collations) {
		return true
	}

	return false
}

// SetCollations gets a reference to the given []MssqlDatabaseCollation and assigns it to the Collations field.
func (o *ListCollationsResponse) SetCollations(v *[]MssqlDatabaseCollation) {
	o.Collations = v
}

func (o ListCollationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collations) {
		toSerialize["collations"] = o.Collations
	}
	return toSerialize, nil
}

type NullableListCollationsResponse struct {
	value *ListCollationsResponse
	isSet bool
}

func (v NullableListCollationsResponse) Get() *ListCollationsResponse {
	return v.value
}

func (v *NullableListCollationsResponse) Set(val *ListCollationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCollationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCollationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCollationsResponse(val *ListCollationsResponse) *NullableListCollationsResponse {
	return &NullableListCollationsResponse{value: val, isSet: true}
}

func (v NullableListCollationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCollationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
