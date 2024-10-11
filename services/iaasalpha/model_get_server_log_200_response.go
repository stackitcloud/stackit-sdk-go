/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
)

// checks if the GetServerLog200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServerLog200Response{}

// GetServerLog200Response struct for GetServerLog200Response
type GetServerLog200Response struct {
	Output *string `json:"output,omitempty"`
}

// NewGetServerLog200Response instantiates a new GetServerLog200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServerLog200Response() *GetServerLog200Response {
	this := GetServerLog200Response{}
	return &this
}

// NewGetServerLog200ResponseWithDefaults instantiates a new GetServerLog200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServerLog200ResponseWithDefaults() *GetServerLog200Response {
	this := GetServerLog200Response{}
	return &this
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *GetServerLog200Response) GetOutput() *string {
	if o == nil || IsNil(o.Output) {
		var ret *string
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerLog200Response) GetOutputOk() (*string, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *GetServerLog200Response) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *GetServerLog200Response) SetOutput(v *string) {
	o.Output = v
}

func (o GetServerLog200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	return toSerialize, nil
}

type NullableGetServerLog200Response struct {
	value *GetServerLog200Response
	isSet bool
}

func (v NullableGetServerLog200Response) Get() *GetServerLog200Response {
	return v.value
}

func (v *NullableGetServerLog200Response) Set(val *GetServerLog200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServerLog200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServerLog200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServerLog200Response(val *GetServerLog200Response) *NullableGetServerLog200Response {
	return &NullableGetServerLog200Response{value: val, isSet: true}
}

func (v NullableGetServerLog200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServerLog200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
