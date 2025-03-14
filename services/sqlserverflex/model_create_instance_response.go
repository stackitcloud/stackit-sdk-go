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

// checks if the CreateInstanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInstanceResponse{}

/*
	types and functions for id
*/

// isNotNullableString
type CreateInstanceResponseGetIdAttributeType = *string

func getCreateInstanceResponseGetIdAttributeTypeOk(arg CreateInstanceResponseGetIdAttributeType) (ret CreateInstanceResponseGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateInstanceResponseGetIdAttributeType(arg *CreateInstanceResponseGetIdAttributeType, val CreateInstanceResponseGetIdRetType) {
	*arg = &val
}

type CreateInstanceResponseGetIdArgType = string
type CreateInstanceResponseGetIdRetType = string

// CreateInstanceResponse struct for CreateInstanceResponse
type CreateInstanceResponse struct {
	Id CreateInstanceResponseGetIdAttributeType `json:"id,omitempty"`
}

// NewCreateInstanceResponse instantiates a new CreateInstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInstanceResponse() *CreateInstanceResponse {
	this := CreateInstanceResponse{}
	return &this
}

// NewCreateInstanceResponseWithDefaults instantiates a new CreateInstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInstanceResponseWithDefaults() *CreateInstanceResponse {
	this := CreateInstanceResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateInstanceResponse) GetId() (res CreateInstanceResponseGetIdRetType) {
	res, _ = o.GetIdOk()
	return
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceResponse) GetIdOk() (ret CreateInstanceResponseGetIdRetType, ok bool) {
	return getCreateInstanceResponseGetIdAttributeTypeOk(o.Id)
}

// HasId returns a boolean if a field has been set.
func (o *CreateInstanceResponse) HasId() bool {
	_, ok := o.GetIdOk()
	return ok
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateInstanceResponse) SetId(v CreateInstanceResponseGetIdRetType) {
	setCreateInstanceResponseGetIdAttributeType(&o.Id, v)
}

func (o CreateInstanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateInstanceResponseGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	return toSerialize, nil
}

type NullableCreateInstanceResponse struct {
	value *CreateInstanceResponse
	isSet bool
}

func (v NullableCreateInstanceResponse) Get() *CreateInstanceResponse {
	return v.value
}

func (v *NullableCreateInstanceResponse) Set(val *CreateInstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInstanceResponse(val *CreateInstanceResponse) *NullableCreateInstanceResponse {
	return &NullableCreateInstanceResponse{value: val, isSet: true}
}

func (v NullableCreateInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
