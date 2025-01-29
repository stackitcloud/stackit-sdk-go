/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 2.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

import (
	"encoding/json"
)

// checks if the ErrorMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorMessage{}

// ErrorMessage struct for ErrorMessage
type ErrorMessage struct {
	// REQUIRED
	Detail *[]DetailedError `json:"detail"`
}

type _ErrorMessage ErrorMessage

// NewErrorMessage instantiates a new ErrorMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorMessage(detail *[]DetailedError) *ErrorMessage {
	this := ErrorMessage{}
	this.Detail = detail
	return &this
}

// NewErrorMessageWithDefaults instantiates a new ErrorMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorMessageWithDefaults() *ErrorMessage {
	this := ErrorMessage{}
	return &this
}

// GetDetail returns the Detail field value
func (o *ErrorMessage) GetDetail() *[]DetailedError {
	if o == nil || IsNil(o.Detail) {
		var ret *[]DetailedError
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *ErrorMessage) GetDetailOk() (*[]DetailedError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail, true
}

// SetDetail sets field value
func (o *ErrorMessage) SetDetail(v *[]DetailedError) {
	o.Detail = v
}

func (o ErrorMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detail"] = o.Detail
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
