/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 2beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alb

import (
	"encoding/json"
)

// checks if the QueryParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryParameters{}

/*
	types and functions for exactMatch
*/

// isNotNullableString
type QueryParametersGetExactMatchAttributeType = *string

func getQueryParametersGetExactMatchAttributeTypeOk(arg QueryParametersGetExactMatchAttributeType) (ret QueryParametersGetExactMatchRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQueryParametersGetExactMatchAttributeType(arg *QueryParametersGetExactMatchAttributeType, val QueryParametersGetExactMatchRetType) {
	*arg = &val
}

type QueryParametersGetExactMatchArgType = string
type QueryParametersGetExactMatchRetType = string

/*
	types and functions for name
*/

// isNotNullableString
type QueryParametersGetNameAttributeType = *string

func getQueryParametersGetNameAttributeTypeOk(arg QueryParametersGetNameAttributeType) (ret QueryParametersGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setQueryParametersGetNameAttributeType(arg *QueryParametersGetNameAttributeType, val QueryParametersGetNameRetType) {
	*arg = &val
}

type QueryParametersGetNameArgType = string
type QueryParametersGetNameRetType = string

// QueryParameters struct for QueryParameters
type QueryParameters struct {
	// Exact match for the parameter value
	ExactMatch QueryParametersGetExactMatchAttributeType `json:"exactMatch,omitempty"`
	// Parameter name
	Name QueryParametersGetNameAttributeType `json:"name,omitempty"`
}

// NewQueryParameters instantiates a new QueryParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryParameters() *QueryParameters {
	this := QueryParameters{}
	return &this
}

// NewQueryParametersWithDefaults instantiates a new QueryParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryParametersWithDefaults() *QueryParameters {
	this := QueryParameters{}
	return &this
}

// GetExactMatch returns the ExactMatch field value if set, zero value otherwise.
func (o *QueryParameters) GetExactMatch() (res QueryParametersGetExactMatchRetType) {
	res, _ = o.GetExactMatchOk()
	return
}

// GetExactMatchOk returns a tuple with the ExactMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryParameters) GetExactMatchOk() (ret QueryParametersGetExactMatchRetType, ok bool) {
	return getQueryParametersGetExactMatchAttributeTypeOk(o.ExactMatch)
}

// HasExactMatch returns a boolean if a field has been set.
func (o *QueryParameters) HasExactMatch() bool {
	_, ok := o.GetExactMatchOk()
	return ok
}

// SetExactMatch gets a reference to the given string and assigns it to the ExactMatch field.
func (o *QueryParameters) SetExactMatch(v QueryParametersGetExactMatchRetType) {
	setQueryParametersGetExactMatchAttributeType(&o.ExactMatch, v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QueryParameters) GetName() (res QueryParametersGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryParameters) GetNameOk() (ret QueryParametersGetNameRetType, ok bool) {
	return getQueryParametersGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *QueryParameters) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QueryParameters) SetName(v QueryParametersGetNameRetType) {
	setQueryParametersGetNameAttributeType(&o.Name, v)
}

func (o QueryParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getQueryParametersGetExactMatchAttributeTypeOk(o.ExactMatch); ok {
		toSerialize["ExactMatch"] = val
	}
	if val, ok := getQueryParametersGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	return toSerialize, nil
}

type NullableQueryParameters struct {
	value *QueryParameters
	isSet bool
}

func (v NullableQueryParameters) Get() *QueryParameters {
	return v.value
}

func (v *NullableQueryParameters) Set(val *QueryParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryParameters(val *QueryParameters) *NullableQueryParameters {
	return &NullableQueryParameters{value: val, isSet: true}
}

func (v NullableQueryParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
