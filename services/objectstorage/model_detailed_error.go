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

// checks if the DetailedError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailedError{}

/*
	types and functions for key
*/

// isNotNullableString
type DetailedErrorGetKeyAttributeType = *string

func getDetailedErrorGetKeyAttributeTypeOk(arg DetailedErrorGetKeyAttributeType) (ret DetailedErrorGetKeyRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setDetailedErrorGetKeyAttributeType(arg *DetailedErrorGetKeyAttributeType, val DetailedErrorGetKeyRetType) {
	*arg = &val
}

type DetailedErrorGetKeyArgType = string
type DetailedErrorGetKeyRetType = string

/*
	types and functions for msg
*/

// isNotNullableString
type DetailedErrorGetMsgAttributeType = *string

func getDetailedErrorGetMsgAttributeTypeOk(arg DetailedErrorGetMsgAttributeType) (ret DetailedErrorGetMsgRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setDetailedErrorGetMsgAttributeType(arg *DetailedErrorGetMsgAttributeType, val DetailedErrorGetMsgRetType) {
	*arg = &val
}

type DetailedErrorGetMsgArgType = string
type DetailedErrorGetMsgRetType = string

// DetailedError struct for DetailedError
type DetailedError struct {
	// REQUIRED
	Key DetailedErrorGetKeyAttributeType `json:"key"`
	// REQUIRED
	Msg DetailedErrorGetMsgAttributeType `json:"msg"`
}

type _DetailedError DetailedError

// NewDetailedError instantiates a new DetailedError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedError(key DetailedErrorGetKeyArgType, msg DetailedErrorGetMsgArgType) *DetailedError {
	this := DetailedError{}
	setDetailedErrorGetKeyAttributeType(&this.Key, key)
	setDetailedErrorGetMsgAttributeType(&this.Msg, msg)
	return &this
}

// NewDetailedErrorWithDefaults instantiates a new DetailedError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedErrorWithDefaults() *DetailedError {
	this := DetailedError{}
	return &this
}

// GetKey returns the Key field value
func (o *DetailedError) GetKey() (ret DetailedErrorGetKeyRetType) {
	ret, _ = o.GetKeyOk()
	return ret
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DetailedError) GetKeyOk() (ret DetailedErrorGetKeyRetType, ok bool) {
	return getDetailedErrorGetKeyAttributeTypeOk(o.Key)
}

// SetKey sets field value
func (o *DetailedError) SetKey(v DetailedErrorGetKeyRetType) {
	setDetailedErrorGetKeyAttributeType(&o.Key, v)
}

// GetMsg returns the Msg field value
func (o *DetailedError) GetMsg() (ret DetailedErrorGetMsgRetType) {
	ret, _ = o.GetMsgOk()
	return ret
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *DetailedError) GetMsgOk() (ret DetailedErrorGetMsgRetType, ok bool) {
	return getDetailedErrorGetMsgAttributeTypeOk(o.Msg)
}

// SetMsg sets field value
func (o *DetailedError) SetMsg(v DetailedErrorGetMsgRetType) {
	setDetailedErrorGetMsgAttributeType(&o.Msg, v)
}

func (o DetailedError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getDetailedErrorGetKeyAttributeTypeOk(o.Key); ok {
		toSerialize["Key"] = val
	}
	if val, ok := getDetailedErrorGetMsgAttributeTypeOk(o.Msg); ok {
		toSerialize["Msg"] = val
	}
	return toSerialize, nil
}

type NullableDetailedError struct {
	value *DetailedError
	isSet bool
}

func (v NullableDetailedError) Get() *DetailedError {
	return v.value
}

func (v *NullableDetailedError) Set(val *DetailedError) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedError) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedError(val *DetailedError) *NullableDetailedError {
	return &NullableDetailedError{value: val, isSet: true}
}

func (v NullableDetailedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
