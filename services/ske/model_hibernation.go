/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

import (
	"encoding/json"
)

// checks if the Hibernation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Hibernation{}

/*
	types and functions for schedules
*/

// isArray
type HibernationGetSchedulesAttributeType = *[]HibernationSchedule
type HibernationGetSchedulesArgType = []HibernationSchedule
type HibernationGetSchedulesRetType = []HibernationSchedule

func getHibernationGetSchedulesAttributeTypeOk(arg HibernationGetSchedulesAttributeType) (ret HibernationGetSchedulesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setHibernationGetSchedulesAttributeType(arg *HibernationGetSchedulesAttributeType, val HibernationGetSchedulesRetType) {
	*arg = &val
}

// Hibernation struct for Hibernation
type Hibernation struct {
	// REQUIRED
	Schedules HibernationGetSchedulesAttributeType `json:"schedules" required:"true"`
}

type _Hibernation Hibernation

// NewHibernation instantiates a new Hibernation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHibernation(schedules HibernationGetSchedulesArgType) *Hibernation {
	this := Hibernation{}
	setHibernationGetSchedulesAttributeType(&this.Schedules, schedules)
	return &this
}

// NewHibernationWithDefaults instantiates a new Hibernation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHibernationWithDefaults() *Hibernation {
	this := Hibernation{}
	return &this
}

// GetSchedules returns the Schedules field value
func (o *Hibernation) GetSchedules() (ret HibernationGetSchedulesRetType) {
	ret, _ = o.GetSchedulesOk()
	return ret
}

// GetSchedulesOk returns a tuple with the Schedules field value
// and a boolean to check if the value has been set.
func (o *Hibernation) GetSchedulesOk() (ret HibernationGetSchedulesRetType, ok bool) {
	return getHibernationGetSchedulesAttributeTypeOk(o.Schedules)
}

// SetSchedules sets field value
func (o *Hibernation) SetSchedules(v HibernationGetSchedulesRetType) {
	setHibernationGetSchedulesAttributeType(&o.Schedules, v)
}

func (o Hibernation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getHibernationGetSchedulesAttributeTypeOk(o.Schedules); ok {
		toSerialize["Schedules"] = val
	}
	return toSerialize, nil
}

type NullableHibernation struct {
	value *Hibernation
	isSet bool
}

func (v NullableHibernation) Get() *Hibernation {
	return v.value
}

func (v *NullableHibernation) Set(val *Hibernation) {
	v.value = val
	v.isSet = true
}

func (v NullableHibernation) IsSet() bool {
	return v.isSet
}

func (v *NullableHibernation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHibernation(val *Hibernation) *NullableHibernation {
	return &NullableHibernation{value: val, isSet: true}
}

func (v NullableHibernation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHibernation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
