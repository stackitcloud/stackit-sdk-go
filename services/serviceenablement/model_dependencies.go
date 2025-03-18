/*
STACKIT Service Enablement API

STACKIT Service Enablement API

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceenablement

import (
	"encoding/json"
)

// checks if the Dependencies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dependencies{}

/*
	types and functions for hard
*/

// isArray
type DependenciesGetHardAttributeType = *[]string
type DependenciesGetHardArgType = []string
type DependenciesGetHardRetType = []string

func getDependenciesGetHardAttributeTypeOk(arg DependenciesGetHardAttributeType) (ret DependenciesGetHardRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setDependenciesGetHardAttributeType(arg *DependenciesGetHardAttributeType, val DependenciesGetHardRetType) {
	*arg = &val
}

/*
	types and functions for soft
*/

// isArray
type DependenciesGetSoftAttributeType = *[]string
type DependenciesGetSoftArgType = []string
type DependenciesGetSoftRetType = []string

func getDependenciesGetSoftAttributeTypeOk(arg DependenciesGetSoftAttributeType) (ret DependenciesGetSoftRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setDependenciesGetSoftAttributeType(arg *DependenciesGetSoftAttributeType, val DependenciesGetSoftRetType) {
	*arg = &val
}

// Dependencies struct for Dependencies
type Dependencies struct {
	// a list of service IDs which this service depend on. If the service is enabled, those service are enabled as well automatically.
	Hard DependenciesGetHardAttributeType `json:"hard,omitempty"`
	// a list of service IDs which this service depend on. When they are disabled a notification is sent.
	Soft DependenciesGetSoftAttributeType `json:"soft,omitempty"`
}

// NewDependencies instantiates a new Dependencies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDependencies() *Dependencies {
	this := Dependencies{}
	return &this
}

// NewDependenciesWithDefaults instantiates a new Dependencies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDependenciesWithDefaults() *Dependencies {
	this := Dependencies{}
	return &this
}

// GetHard returns the Hard field value if set, zero value otherwise.
func (o *Dependencies) GetHard() (res DependenciesGetHardRetType) {
	res, _ = o.GetHardOk()
	return
}

// GetHardOk returns a tuple with the Hard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependencies) GetHardOk() (ret DependenciesGetHardRetType, ok bool) {
	return getDependenciesGetHardAttributeTypeOk(o.Hard)
}

// HasHard returns a boolean if a field has been set.
func (o *Dependencies) HasHard() bool {
	_, ok := o.GetHardOk()
	return ok
}

// SetHard gets a reference to the given []string and assigns it to the Hard field.
func (o *Dependencies) SetHard(v DependenciesGetHardRetType) {
	setDependenciesGetHardAttributeType(&o.Hard, v)
}

// GetSoft returns the Soft field value if set, zero value otherwise.
func (o *Dependencies) GetSoft() (res DependenciesGetSoftRetType) {
	res, _ = o.GetSoftOk()
	return
}

// GetSoftOk returns a tuple with the Soft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependencies) GetSoftOk() (ret DependenciesGetSoftRetType, ok bool) {
	return getDependenciesGetSoftAttributeTypeOk(o.Soft)
}

// HasSoft returns a boolean if a field has been set.
func (o *Dependencies) HasSoft() bool {
	_, ok := o.GetSoftOk()
	return ok
}

// SetSoft gets a reference to the given []string and assigns it to the Soft field.
func (o *Dependencies) SetSoft(v DependenciesGetSoftRetType) {
	setDependenciesGetSoftAttributeType(&o.Soft, v)
}

func (o Dependencies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getDependenciesGetHardAttributeTypeOk(o.Hard); ok {
		toSerialize["Hard"] = val
	}
	if val, ok := getDependenciesGetSoftAttributeTypeOk(o.Soft); ok {
		toSerialize["Soft"] = val
	}
	return toSerialize, nil
}

type NullableDependencies struct {
	value *Dependencies
	isSet bool
}

func (v NullableDependencies) Get() *Dependencies {
	return v.value
}

func (v *NullableDependencies) Set(val *Dependencies) {
	v.value = val
	v.isSet = true
}

func (v NullableDependencies) IsSet() bool {
	return v.isSet
}

func (v *NullableDependencies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDependencies(val *Dependencies) *NullableDependencies {
	return &NullableDependencies{value: val, isSet: true}
}

func (v NullableDependencies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDependencies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
