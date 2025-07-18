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

// checks if the AlertGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertGroup{}

/*
	types and functions for interval
*/

// isNotNullableString
type AlertGroupGetIntervalAttributeType = *string

func getAlertGroupGetIntervalAttributeTypeOk(arg AlertGroupGetIntervalAttributeType) (ret AlertGroupGetIntervalRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertGroupGetIntervalAttributeType(arg *AlertGroupGetIntervalAttributeType, val AlertGroupGetIntervalRetType) {
	*arg = &val
}

type AlertGroupGetIntervalArgType = string
type AlertGroupGetIntervalRetType = string

/*
	types and functions for name
*/

// isNotNullableString
type AlertGroupGetNameAttributeType = *string

func getAlertGroupGetNameAttributeTypeOk(arg AlertGroupGetNameAttributeType) (ret AlertGroupGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertGroupGetNameAttributeType(arg *AlertGroupGetNameAttributeType, val AlertGroupGetNameRetType) {
	*arg = &val
}

type AlertGroupGetNameArgType = string
type AlertGroupGetNameRetType = string

/*
	types and functions for rules
*/

// isArray
type AlertGroupGetRulesAttributeType = *[]AlertRuleRecord
type AlertGroupGetRulesArgType = []AlertRuleRecord
type AlertGroupGetRulesRetType = []AlertRuleRecord

func getAlertGroupGetRulesAttributeTypeOk(arg AlertGroupGetRulesAttributeType) (ret AlertGroupGetRulesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertGroupGetRulesAttributeType(arg *AlertGroupGetRulesAttributeType, val AlertGroupGetRulesRetType) {
	*arg = &val
}

// AlertGroup struct for AlertGroup
type AlertGroup struct {
	Interval AlertGroupGetIntervalAttributeType `json:"interval,omitempty"`
	// REQUIRED
	Name AlertGroupGetNameAttributeType `json:"name" required:"true"`
	// REQUIRED
	Rules AlertGroupGetRulesAttributeType `json:"rules" required:"true"`
}

type _AlertGroup AlertGroup

// NewAlertGroup instantiates a new AlertGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertGroup(name AlertGroupGetNameArgType, rules AlertGroupGetRulesArgType) *AlertGroup {
	this := AlertGroup{}
	setAlertGroupGetNameAttributeType(&this.Name, name)
	setAlertGroupGetRulesAttributeType(&this.Rules, rules)
	return &this
}

// NewAlertGroupWithDefaults instantiates a new AlertGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertGroupWithDefaults() *AlertGroup {
	this := AlertGroup{}
	var interval string = "60s"
	this.Interval = &interval
	return &this
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *AlertGroup) GetInterval() (res AlertGroupGetIntervalRetType) {
	res, _ = o.GetIntervalOk()
	return
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGroup) GetIntervalOk() (ret AlertGroupGetIntervalRetType, ok bool) {
	return getAlertGroupGetIntervalAttributeTypeOk(o.Interval)
}

// HasInterval returns a boolean if a field has been set.
func (o *AlertGroup) HasInterval() bool {
	_, ok := o.GetIntervalOk()
	return ok
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *AlertGroup) SetInterval(v AlertGroupGetIntervalRetType) {
	setAlertGroupGetIntervalAttributeType(&o.Interval, v)
}

// GetName returns the Name field value
func (o *AlertGroup) GetName() (ret AlertGroupGetNameRetType) {
	ret, _ = o.GetNameOk()
	return ret
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AlertGroup) GetNameOk() (ret AlertGroupGetNameRetType, ok bool) {
	return getAlertGroupGetNameAttributeTypeOk(o.Name)
}

// SetName sets field value
func (o *AlertGroup) SetName(v AlertGroupGetNameRetType) {
	setAlertGroupGetNameAttributeType(&o.Name, v)
}

// GetRules returns the Rules field value
func (o *AlertGroup) GetRules() (ret AlertGroupGetRulesRetType) {
	ret, _ = o.GetRulesOk()
	return ret
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *AlertGroup) GetRulesOk() (ret AlertGroupGetRulesRetType, ok bool) {
	return getAlertGroupGetRulesAttributeTypeOk(o.Rules)
}

// SetRules sets field value
func (o *AlertGroup) SetRules(v AlertGroupGetRulesRetType) {
	setAlertGroupGetRulesAttributeType(&o.Rules, v)
}

func (o AlertGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getAlertGroupGetIntervalAttributeTypeOk(o.Interval); ok {
		toSerialize["Interval"] = val
	}
	if val, ok := getAlertGroupGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getAlertGroupGetRulesAttributeTypeOk(o.Rules); ok {
		toSerialize["Rules"] = val
	}
	return toSerialize, nil
}

type NullableAlertGroup struct {
	value *AlertGroup
	isSet bool
}

func (v NullableAlertGroup) Get() *AlertGroup {
	return v.value
}

func (v *NullableAlertGroup) Set(val *AlertGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertGroup(val *AlertGroup) *NullableAlertGroup {
	return &NullableAlertGroup{value: val, isSet: true}
}

func (v NullableAlertGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
