/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"encoding/json"
)

// checks if the AlertConfigReceiversResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertConfigReceiversResponse{}

/*
	types and functions for data
*/

// isArray
type AlertConfigReceiversResponseGetDataAttributeType = *[]Receivers
type AlertConfigReceiversResponseGetDataArgType = []Receivers
type AlertConfigReceiversResponseGetDataRetType = []Receivers

func getAlertConfigReceiversResponseGetDataAttributeTypeOk(arg AlertConfigReceiversResponseGetDataAttributeType) (ret AlertConfigReceiversResponseGetDataRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertConfigReceiversResponseGetDataAttributeType(arg *AlertConfigReceiversResponseGetDataAttributeType, val AlertConfigReceiversResponseGetDataRetType) {
	*arg = &val
}

/*
	types and functions for message
*/

// isNotNullableString
type AlertConfigReceiversResponseGetMessageAttributeType = *string

func getAlertConfigReceiversResponseGetMessageAttributeTypeOk(arg AlertConfigReceiversResponseGetMessageAttributeType) (ret AlertConfigReceiversResponseGetMessageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertConfigReceiversResponseGetMessageAttributeType(arg *AlertConfigReceiversResponseGetMessageAttributeType, val AlertConfigReceiversResponseGetMessageRetType) {
	*arg = &val
}

type AlertConfigReceiversResponseGetMessageArgType = string
type AlertConfigReceiversResponseGetMessageRetType = string

// AlertConfigReceiversResponse struct for AlertConfigReceiversResponse
type AlertConfigReceiversResponse struct {
	// REQUIRED
	Data AlertConfigReceiversResponseGetDataAttributeType `json:"data"`
	// REQUIRED
	Message AlertConfigReceiversResponseGetMessageAttributeType `json:"message"`
}

type _AlertConfigReceiversResponse AlertConfigReceiversResponse

// NewAlertConfigReceiversResponse instantiates a new AlertConfigReceiversResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertConfigReceiversResponse(data AlertConfigReceiversResponseGetDataArgType, message AlertConfigReceiversResponseGetMessageArgType) *AlertConfigReceiversResponse {
	this := AlertConfigReceiversResponse{}
	setAlertConfigReceiversResponseGetDataAttributeType(&this.Data, data)
	setAlertConfigReceiversResponseGetMessageAttributeType(&this.Message, message)
	return &this
}

// NewAlertConfigReceiversResponseWithDefaults instantiates a new AlertConfigReceiversResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertConfigReceiversResponseWithDefaults() *AlertConfigReceiversResponse {
	this := AlertConfigReceiversResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AlertConfigReceiversResponse) GetData() (ret AlertConfigReceiversResponseGetDataRetType) {
	ret, _ = o.GetDataOk()
	return ret
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AlertConfigReceiversResponse) GetDataOk() (ret AlertConfigReceiversResponseGetDataRetType, ok bool) {
	return getAlertConfigReceiversResponseGetDataAttributeTypeOk(o.Data)
}

// SetData sets field value
func (o *AlertConfigReceiversResponse) SetData(v AlertConfigReceiversResponseGetDataRetType) {
	setAlertConfigReceiversResponseGetDataAttributeType(&o.Data, v)
}

// GetMessage returns the Message field value
func (o *AlertConfigReceiversResponse) GetMessage() (ret AlertConfigReceiversResponseGetMessageRetType) {
	ret, _ = o.GetMessageOk()
	return ret
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlertConfigReceiversResponse) GetMessageOk() (ret AlertConfigReceiversResponseGetMessageRetType, ok bool) {
	return getAlertConfigReceiversResponseGetMessageAttributeTypeOk(o.Message)
}

// SetMessage sets field value
func (o *AlertConfigReceiversResponse) SetMessage(v AlertConfigReceiversResponseGetMessageRetType) {
	setAlertConfigReceiversResponseGetMessageAttributeType(&o.Message, v)
}

func (o AlertConfigReceiversResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getAlertConfigReceiversResponseGetDataAttributeTypeOk(o.Data); ok {
		toSerialize["Data"] = val
	}
	if val, ok := getAlertConfigReceiversResponseGetMessageAttributeTypeOk(o.Message); ok {
		toSerialize["Message"] = val
	}
	return toSerialize, nil
}

type NullableAlertConfigReceiversResponse struct {
	value *AlertConfigReceiversResponse
	isSet bool
}

func (v NullableAlertConfigReceiversResponse) Get() *AlertConfigReceiversResponse {
	return v.value
}

func (v *NullableAlertConfigReceiversResponse) Set(val *AlertConfigReceiversResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertConfigReceiversResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertConfigReceiversResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertConfigReceiversResponse(val *AlertConfigReceiversResponse) *NullableAlertConfigReceiversResponse {
	return &NullableAlertConfigReceiversResponse{value: val, isSet: true}
}

func (v NullableAlertConfigReceiversResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertConfigReceiversResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
