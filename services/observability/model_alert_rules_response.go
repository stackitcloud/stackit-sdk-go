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

// checks if the AlertRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertRulesResponse{}

/*
	types and functions for data
*/

// isArray
type AlertRulesResponseGetDataAttributeType = *[]AlertRule
type AlertRulesResponseGetDataArgType = []AlertRule
type AlertRulesResponseGetDataRetType = []AlertRule

func getAlertRulesResponseGetDataAttributeTypeOk(arg AlertRulesResponseGetDataAttributeType) (ret AlertRulesResponseGetDataRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertRulesResponseGetDataAttributeType(arg *AlertRulesResponseGetDataAttributeType, val AlertRulesResponseGetDataRetType) {
	*arg = &val
}

/*
	types and functions for message
*/

// isNotNullableString
type AlertRulesResponseGetMessageAttributeType = *string

func getAlertRulesResponseGetMessageAttributeTypeOk(arg AlertRulesResponseGetMessageAttributeType) (ret AlertRulesResponseGetMessageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertRulesResponseGetMessageAttributeType(arg *AlertRulesResponseGetMessageAttributeType, val AlertRulesResponseGetMessageRetType) {
	*arg = &val
}

type AlertRulesResponseGetMessageArgType = string
type AlertRulesResponseGetMessageRetType = string

// AlertRulesResponse struct for AlertRulesResponse
type AlertRulesResponse struct {
	// REQUIRED
	Data AlertRulesResponseGetDataAttributeType `json:"data"`
	// REQUIRED
	Message AlertRulesResponseGetMessageAttributeType `json:"message"`
}

type _AlertRulesResponse AlertRulesResponse

// NewAlertRulesResponse instantiates a new AlertRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertRulesResponse(data AlertRulesResponseGetDataArgType, message AlertRulesResponseGetMessageArgType) *AlertRulesResponse {
	this := AlertRulesResponse{}
	setAlertRulesResponseGetDataAttributeType(&this.Data, data)
	setAlertRulesResponseGetMessageAttributeType(&this.Message, message)
	return &this
}

// NewAlertRulesResponseWithDefaults instantiates a new AlertRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertRulesResponseWithDefaults() *AlertRulesResponse {
	this := AlertRulesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AlertRulesResponse) GetData() (ret AlertRulesResponseGetDataRetType) {
	ret, _ = o.GetDataOk()
	return ret
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AlertRulesResponse) GetDataOk() (ret AlertRulesResponseGetDataRetType, ok bool) {
	return getAlertRulesResponseGetDataAttributeTypeOk(o.Data)
}

// SetData sets field value
func (o *AlertRulesResponse) SetData(v AlertRulesResponseGetDataRetType) {
	setAlertRulesResponseGetDataAttributeType(&o.Data, v)
}

// GetMessage returns the Message field value
func (o *AlertRulesResponse) GetMessage() (ret AlertRulesResponseGetMessageRetType) {
	ret, _ = o.GetMessageOk()
	return ret
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlertRulesResponse) GetMessageOk() (ret AlertRulesResponseGetMessageRetType, ok bool) {
	return getAlertRulesResponseGetMessageAttributeTypeOk(o.Message)
}

// SetMessage sets field value
func (o *AlertRulesResponse) SetMessage(v AlertRulesResponseGetMessageRetType) {
	setAlertRulesResponseGetMessageAttributeType(&o.Message, v)
}

func (o AlertRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getAlertRulesResponseGetDataAttributeTypeOk(o.Data); ok {
		toSerialize["Data"] = val
	}
	if val, ok := getAlertRulesResponseGetMessageAttributeTypeOk(o.Message); ok {
		toSerialize["Message"] = val
	}
	return toSerialize, nil
}

type NullableAlertRulesResponse struct {
	value *AlertRulesResponse
	isSet bool
}

func (v NullableAlertRulesResponse) Get() *AlertRulesResponse {
	return v.value
}

func (v *NullableAlertRulesResponse) Set(val *AlertRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertRulesResponse(val *AlertRulesResponse) *NullableAlertRulesResponse {
	return &NullableAlertRulesResponse{value: val, isSet: true}
}

func (v NullableAlertRulesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
