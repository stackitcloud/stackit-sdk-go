/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the Error type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Error{}

/*
	types and functions for code
*/

// isLong
type ErrorGetCodeAttributeType = *int64
type ErrorGetCodeArgType = int64
type ErrorGetCodeRetType = int64

func getErrorGetCodeAttributeTypeOk(arg ErrorGetCodeAttributeType) (ret ErrorGetCodeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setErrorGetCodeAttributeType(arg *ErrorGetCodeAttributeType, val ErrorGetCodeRetType) {
	*arg = &val
}

/*
	types and functions for msg
*/

// isNotNullableString
type ErrorGetMsgAttributeType = *string

func getErrorGetMsgAttributeTypeOk(arg ErrorGetMsgAttributeType) (ret ErrorGetMsgRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setErrorGetMsgAttributeType(arg *ErrorGetMsgAttributeType, val ErrorGetMsgRetType) {
	*arg = &val
}

type ErrorGetMsgArgType = string
type ErrorGetMsgRetType = string

// Error Error with HTTP error code and an error message.
type Error struct {
	// REQUIRED
	Code ErrorGetCodeAttributeType `json:"code"`
	// An error message.
	// REQUIRED
	Msg ErrorGetMsgAttributeType `json:"msg"`
}

type _Error Error

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError(code ErrorGetCodeArgType, msg ErrorGetMsgArgType) *Error {
	this := Error{}
	setErrorGetCodeAttributeType(&this.Code, code)
	setErrorGetMsgAttributeType(&this.Msg, msg)
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetCode returns the Code field value
func (o *Error) GetCode() (ret ErrorGetCodeRetType) {
	ret, _ = o.GetCodeOk()
	return ret
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Error) GetCodeOk() (ret ErrorGetCodeRetType, ok bool) {
	return getErrorGetCodeAttributeTypeOk(o.Code)
}

// SetCode sets field value
func (o *Error) SetCode(v ErrorGetCodeRetType) {
	setErrorGetCodeAttributeType(&o.Code, v)
}

// GetMsg returns the Msg field value
func (o *Error) GetMsg() (ret ErrorGetMsgRetType) {
	ret, _ = o.GetMsgOk()
	return ret
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *Error) GetMsgOk() (ret ErrorGetMsgRetType, ok bool) {
	return getErrorGetMsgAttributeTypeOk(o.Msg)
}

// SetMsg sets field value
func (o *Error) SetMsg(v ErrorGetMsgRetType) {
	setErrorGetMsgAttributeType(&o.Msg, v)
}

func (o Error) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getErrorGetCodeAttributeTypeOk(o.Code); ok {
		toSerialize["Code"] = val
	}
	if val, ok := getErrorGetMsgAttributeTypeOk(o.Msg); ok {
		toSerialize["Msg"] = val
	}
	return toSerialize, nil
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
