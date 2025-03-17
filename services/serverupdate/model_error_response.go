/*
STACKIT Server Update Management API

API endpoints for Server Update Operations on STACKIT Servers.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverupdate

import (
	"encoding/json"
)

// checks if the ErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponse{}

/*
	types and functions for message
*/

// isNotNullableString
type ErrorResponseGetMessageAttributeType = *string

func getErrorResponseGetMessageAttributeTypeOk(arg ErrorResponseGetMessageAttributeType) (ret ErrorResponseGetMessageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setErrorResponseGetMessageAttributeType(arg *ErrorResponseGetMessageAttributeType, val ErrorResponseGetMessageRetType) {
	*arg = &val
}

type ErrorResponseGetMessageArgType = string
type ErrorResponseGetMessageRetType = string

/*
	types and functions for status
*/

// isNotNullableString
type ErrorResponseGetStatusAttributeType = *string

func getErrorResponseGetStatusAttributeTypeOk(arg ErrorResponseGetStatusAttributeType) (ret ErrorResponseGetStatusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setErrorResponseGetStatusAttributeType(arg *ErrorResponseGetStatusAttributeType, val ErrorResponseGetStatusRetType) {
	*arg = &val
}

type ErrorResponseGetStatusArgType = string
type ErrorResponseGetStatusRetType = string

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	// Details about the error
	// REQUIRED
	Message ErrorResponseGetMessageAttributeType `json:"message"`
	// The string representation of the http status code (i.e. Not Found, Bad Request, etc)
	// REQUIRED
	Status ErrorResponseGetStatusAttributeType `json:"status"`
}

type _ErrorResponse ErrorResponse

// NewErrorResponse instantiates a new ErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponse(message ErrorResponseGetMessageArgType, status ErrorResponseGetStatusArgType) *ErrorResponse {
	this := ErrorResponse{}
	setErrorResponseGetMessageAttributeType(&this.Message, message)
	setErrorResponseGetStatusAttributeType(&this.Status, status)
	return &this
}

// NewErrorResponseWithDefaults instantiates a new ErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseWithDefaults() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *ErrorResponse) GetMessage() (ret ErrorResponseGetMessageRetType) {
	ret, _ = o.GetMessageOk()
	return ret
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetMessageOk() (ret ErrorResponseGetMessageRetType, ok bool) {
	return getErrorResponseGetMessageAttributeTypeOk(o.Message)
}

// SetMessage sets field value
func (o *ErrorResponse) SetMessage(v ErrorResponseGetMessageRetType) {
	setErrorResponseGetMessageAttributeType(&o.Message, v)
}

// GetStatus returns the Status field value
func (o *ErrorResponse) GetStatus() (ret ErrorResponseGetStatusRetType) {
	ret, _ = o.GetStatusOk()
	return ret
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetStatusOk() (ret ErrorResponseGetStatusRetType, ok bool) {
	return getErrorResponseGetStatusAttributeTypeOk(o.Status)
}

// SetStatus sets field value
func (o *ErrorResponse) SetStatus(v ErrorResponseGetStatusRetType) {
	setErrorResponseGetStatusAttributeType(&o.Status, v)
}

func (o ErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getErrorResponseGetMessageAttributeTypeOk(o.Message); ok {
		toSerialize["Message"] = val
	}
	if val, ok := getErrorResponseGetStatusAttributeTypeOk(o.Status); ok {
		toSerialize["Status"] = val
	}
	return toSerialize, nil
}

type NullableErrorResponse struct {
	value *ErrorResponse
	isSet bool
}

func (v NullableErrorResponse) Get() *ErrorResponse {
	return v.value
}

func (v *NullableErrorResponse) Set(val *ErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponse(val *ErrorResponse) *NullableErrorResponse {
	return &NullableErrorResponse{value: val, isSet: true}
}

func (v NullableErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
