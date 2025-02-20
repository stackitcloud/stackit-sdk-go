/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
)

// checks if the CreateDatabaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDatabaseResponse{}

// CreateDatabaseResponse struct for CreateDatabaseResponse
type CreateDatabaseResponse struct {
	Id *string `json:"id,omitempty"`
}

// NewCreateDatabaseResponse instantiates a new CreateDatabaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDatabaseResponse() *CreateDatabaseResponse {
	this := CreateDatabaseResponse{}
	return &this
}

// NewCreateDatabaseResponseWithDefaults instantiates a new CreateDatabaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDatabaseResponseWithDefaults() *CreateDatabaseResponse {
	this := CreateDatabaseResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateDatabaseResponse) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabaseResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateDatabaseResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateDatabaseResponse) SetId(v *string) {
	o.Id = v
}

func (o CreateDatabaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCreateDatabaseResponse struct {
	value *CreateDatabaseResponse
	isSet bool
}

func (v NullableCreateDatabaseResponse) Get() *CreateDatabaseResponse {
	return v.value
}

func (v *NullableCreateDatabaseResponse) Set(val *CreateDatabaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDatabaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDatabaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDatabaseResponse(val *CreateDatabaseResponse) *NullableCreateDatabaseResponse {
	return &NullableCreateDatabaseResponse{value: val, isSet: true}
}

func (v NullableCreateDatabaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDatabaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
