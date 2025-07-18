/*
CDN API

API used to create and manage your CDN distributions.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the Optimizer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Optimizer{}

/*
	types and functions for enabled
*/

// isBoolean
type OptimizergetEnabledAttributeType = *bool
type OptimizergetEnabledArgType = bool
type OptimizergetEnabledRetType = bool

func getOptimizergetEnabledAttributeTypeOk(arg OptimizergetEnabledAttributeType) (ret OptimizergetEnabledRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setOptimizergetEnabledAttributeType(arg *OptimizergetEnabledAttributeType, val OptimizergetEnabledRetType) {
	*arg = &val
}

// Optimizer Optimizer is paid feature, a real-time on the fly image manipulation and optimization service that automatically optimizes your images for faster image delivery.
type Optimizer struct {
	// Determines if the optimizer should be enabled for this distribution and incurs a monthly fee
	// REQUIRED
	Enabled OptimizergetEnabledAttributeType `json:"enabled" required:"true"`
}

type _Optimizer Optimizer

// NewOptimizer instantiates a new Optimizer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizer(enabled OptimizergetEnabledArgType) *Optimizer {
	this := Optimizer{}
	setOptimizergetEnabledAttributeType(&this.Enabled, enabled)
	return &this
}

// NewOptimizerWithDefaults instantiates a new Optimizer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizerWithDefaults() *Optimizer {
	this := Optimizer{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *Optimizer) GetEnabled() (ret OptimizergetEnabledRetType) {
	ret, _ = o.GetEnabledOk()
	return ret
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Optimizer) GetEnabledOk() (ret OptimizergetEnabledRetType, ok bool) {
	return getOptimizergetEnabledAttributeTypeOk(o.Enabled)
}

// SetEnabled sets field value
func (o *Optimizer) SetEnabled(v OptimizergetEnabledRetType) {
	setOptimizergetEnabledAttributeType(&o.Enabled, v)
}

func (o Optimizer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getOptimizergetEnabledAttributeTypeOk(o.Enabled); ok {
		toSerialize["Enabled"] = val
	}
	return toSerialize, nil
}

type NullableOptimizer struct {
	value *Optimizer
	isSet bool
}

func (v NullableOptimizer) Get() *Optimizer {
	return v.value
}

func (v *NullableOptimizer) Set(val *Optimizer) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizer) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizer(val *Optimizer) *NullableOptimizer {
	return &NullableOptimizer{value: val, isSet: true}
}

func (v NullableOptimizer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptimizer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
