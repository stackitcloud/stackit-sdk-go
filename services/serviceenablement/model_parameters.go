/*
STACKIT Service Enablement API

STACKIT Service Enablement API

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceenablement

import (
	"encoding/json"
)

// checks if the Parameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Parameters{}

// Parameters service parameters
type Parameters struct {
	General *ParametersGeneral `json:"general,omitempty"`
}

// NewParameters instantiates a new Parameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameters() *Parameters {
	this := Parameters{}
	return &this
}

// NewParametersWithDefaults instantiates a new Parameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParametersWithDefaults() *Parameters {
	this := Parameters{}
	return &this
}

// GetGeneral returns the General field value if set, zero value otherwise.
func (o *Parameters) GetGeneral() *ParametersGeneral {
	if o == nil || IsNil(o.General) {
		var ret *ParametersGeneral
		return ret
	}
	return o.General
}

// GetGeneralOk returns a tuple with the General field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Parameters) GetGeneralOk() (*ParametersGeneral, bool) {
	if o == nil || IsNil(o.General) {
		return nil, false
	}
	return o.General, true
}

// HasGeneral returns a boolean if a field has been set.
func (o *Parameters) HasGeneral() bool {
	if o != nil && !IsNil(o.General) {
		return true
	}

	return false
}

// SetGeneral gets a reference to the given ParametersGeneral and assigns it to the General field.
func (o *Parameters) SetGeneral(v *ParametersGeneral) {
	o.General = v
}

func (o Parameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.General) {
		toSerialize["general"] = o.General
	}
	return toSerialize, nil
}

type NullableParameters struct {
	value *Parameters
	isSet bool
}

func (v NullableParameters) Get() *Parameters {
	return v.value
}

func (v *NullableParameters) Set(val *Parameters) {
	v.value = val
	v.isSet = true
}

func (v NullableParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameters(val *Parameters) *NullableParameters {
	return &NullableParameters{value: val, isSet: true}
}

func (v NullableParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
