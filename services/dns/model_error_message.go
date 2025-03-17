/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the ErrorMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorMessage{}

/*
	types and functions for error
*/

// isNotNullableString
type ErrorMessageGetErrorAttributeType = *string

func getErrorMessageGetErrorAttributeTypeOk(arg ErrorMessageGetErrorAttributeType) (ret ErrorMessageGetErrorRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setErrorMessageGetErrorAttributeType(arg *ErrorMessageGetErrorAttributeType, val ErrorMessageGetErrorRetType) {
	*arg = &val
}

type ErrorMessageGetErrorArgType = string
type ErrorMessageGetErrorRetType = string

/*
	types and functions for message
*/

// isNotNullableString
type ErrorMessageGetMessageAttributeType = *string

func getErrorMessageGetMessageAttributeTypeOk(arg ErrorMessageGetMessageAttributeType) (ret ErrorMessageGetMessageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setErrorMessageGetMessageAttributeType(arg *ErrorMessageGetMessageAttributeType, val ErrorMessageGetMessageRetType) {
	*arg = &val
}

type ErrorMessageGetMessageArgType = string
type ErrorMessageGetMessageRetType = string

// ErrorMessage struct for ErrorMessage
type ErrorMessage struct {
	Error   ErrorMessageGetErrorAttributeType   `json:"error,omitempty"`
	Message ErrorMessageGetMessageAttributeType `json:"message,omitempty"`
}

// NewErrorMessage instantiates a new ErrorMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorMessage() *ErrorMessage {
	this := ErrorMessage{}
	return &this
}

// NewErrorMessageWithDefaults instantiates a new ErrorMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorMessageWithDefaults() *ErrorMessage {
	this := ErrorMessage{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ErrorMessage) GetError() (res ErrorMessageGetErrorRetType) {
	res, _ = o.GetErrorOk()
	return
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessage) GetErrorOk() (ret ErrorMessageGetErrorRetType, ok bool) {
	return getErrorMessageGetErrorAttributeTypeOk(o.Error)
}

// HasError returns a boolean if a field has been set.
func (o *ErrorMessage) HasError() bool {
	_, ok := o.GetErrorOk()
	return ok
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ErrorMessage) SetError(v ErrorMessageGetErrorRetType) {
	setErrorMessageGetErrorAttributeType(&o.Error, v)
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorMessage) GetMessage() (res ErrorMessageGetMessageRetType) {
	res, _ = o.GetMessageOk()
	return
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessage) GetMessageOk() (ret ErrorMessageGetMessageRetType, ok bool) {
	return getErrorMessageGetMessageAttributeTypeOk(o.Message)
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorMessage) HasMessage() bool {
	_, ok := o.GetMessageOk()
	return ok
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorMessage) SetMessage(v ErrorMessageGetMessageRetType) {
	setErrorMessageGetMessageAttributeType(&o.Message, v)
}

func (o ErrorMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getErrorMessageGetErrorAttributeTypeOk(o.Error); ok {
		toSerialize["Error"] = val
	}
	if val, ok := getErrorMessageGetMessageAttributeTypeOk(o.Message); ok {
		toSerialize["Message"] = val
	}
	return toSerialize, nil
}

type NullableErrorMessage struct {
	value *ErrorMessage
	isSet bool
}

func (v NullableErrorMessage) Get() *ErrorMessage {
	return v.value
}

func (v *NullableErrorMessage) Set(val *ErrorMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorMessage(val *ErrorMessage) *NullableErrorMessage {
	return &NullableErrorMessage{value: val, isSet: true}
}

func (v NullableErrorMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
