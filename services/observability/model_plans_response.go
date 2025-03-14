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

// checks if the PlansResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlansResponse{}

/*
	types and functions for message
*/

// isNotNullableString
type PlansResponseGetMessageAttributeType = *string

func getPlansResponseGetMessageAttributeTypeOk(arg PlansResponseGetMessageAttributeType) (ret PlansResponseGetMessageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlansResponseGetMessageAttributeType(arg *PlansResponseGetMessageAttributeType, val PlansResponseGetMessageRetType) {
	*arg = &val
}

type PlansResponseGetMessageArgType = string
type PlansResponseGetMessageRetType = string

/*
	types and functions for plans
*/

// isArray
type PlansResponseGetPlansAttributeType = *[]Plan
type PlansResponseGetPlansArgType = []Plan
type PlansResponseGetPlansRetType = []Plan

func getPlansResponseGetPlansAttributeTypeOk(arg PlansResponseGetPlansAttributeType) (ret PlansResponseGetPlansRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPlansResponseGetPlansAttributeType(arg *PlansResponseGetPlansAttributeType, val PlansResponseGetPlansRetType) {
	*arg = &val
}

// PlansResponse struct for PlansResponse
type PlansResponse struct {
	// REQUIRED
	Message PlansResponseGetMessageAttributeType `json:"message"`
	// REQUIRED
	Plans PlansResponseGetPlansAttributeType `json:"plans"`
}

type _PlansResponse PlansResponse

// NewPlansResponse instantiates a new PlansResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlansResponse(message PlansResponseGetMessageArgType, plans PlansResponseGetPlansArgType) *PlansResponse {
	this := PlansResponse{}
	setPlansResponseGetMessageAttributeType(&this.Message, message)
	setPlansResponseGetPlansAttributeType(&this.Plans, plans)
	return &this
}

// NewPlansResponseWithDefaults instantiates a new PlansResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlansResponseWithDefaults() *PlansResponse {
	this := PlansResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *PlansResponse) GetMessage() (ret PlansResponseGetMessageRetType) {
	ret, _ = o.GetMessageOk()
	return ret
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PlansResponse) GetMessageOk() (ret PlansResponseGetMessageRetType, ok bool) {
	return getPlansResponseGetMessageAttributeTypeOk(o.Message)
}

// SetMessage sets field value
func (o *PlansResponse) SetMessage(v PlansResponseGetMessageRetType) {
	setPlansResponseGetMessageAttributeType(&o.Message, v)
}

// GetPlans returns the Plans field value
func (o *PlansResponse) GetPlans() (ret PlansResponseGetPlansRetType) {
	ret, _ = o.GetPlansOk()
	return ret
}

// GetPlansOk returns a tuple with the Plans field value
// and a boolean to check if the value has been set.
func (o *PlansResponse) GetPlansOk() (ret PlansResponseGetPlansRetType, ok bool) {
	return getPlansResponseGetPlansAttributeTypeOk(o.Plans)
}

// SetPlans sets field value
func (o *PlansResponse) SetPlans(v PlansResponseGetPlansRetType) {
	setPlansResponseGetPlansAttributeType(&o.Plans, v)
}

func (o PlansResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getPlansResponseGetMessageAttributeTypeOk(o.Message); ok {
		toSerialize["Message"] = val
	}
	if val, ok := getPlansResponseGetPlansAttributeTypeOk(o.Plans); ok {
		toSerialize["Plans"] = val
	}
	return toSerialize, nil
}

type NullablePlansResponse struct {
	value *PlansResponse
	isSet bool
}

func (v NullablePlansResponse) Get() *PlansResponse {
	return v.value
}

func (v *NullablePlansResponse) Set(val *PlansResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlansResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlansResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlansResponse(val *PlansResponse) *NullablePlansResponse {
	return &NullablePlansResponse{value: val, isSet: true}
}

func (v NullablePlansResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlansResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
