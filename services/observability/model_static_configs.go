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

// checks if the StaticConfigs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticConfigs{}

// StaticConfigs struct for StaticConfigs
type StaticConfigs struct {
	Labels *map[string]string `json:"labels,omitempty"`
	// REQUIRED
	Targets *[]string `json:"targets"`
}

type _StaticConfigs StaticConfigs

// NewStaticConfigs instantiates a new StaticConfigs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticConfigs(targets *[]string) *StaticConfigs {
	this := StaticConfigs{}
	this.Targets = targets
	return &this
}

// NewStaticConfigsWithDefaults instantiates a new StaticConfigs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticConfigsWithDefaults() *StaticConfigs {
	this := StaticConfigs{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *StaticConfigs) GetLabels() *map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticConfigs) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *StaticConfigs) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *StaticConfigs) SetLabels(v *map[string]string) {
	o.Labels = v
}

// GetTargets returns the Targets field value
func (o *StaticConfigs) GetTargets() *[]string {
	if o == nil {
		var ret *[]string
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *StaticConfigs) GetTargetsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Targets, true
}

// SetTargets sets field value
func (o *StaticConfigs) SetTargets(v *[]string) {
	o.Targets = v
}

func (o StaticConfigs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["targets"] = o.Targets
	return toSerialize, nil
}

type NullableStaticConfigs struct {
	value *StaticConfigs
	isSet bool
}

func (v NullableStaticConfigs) Get() *StaticConfigs {
	return v.value
}

func (v *NullableStaticConfigs) Set(val *StaticConfigs) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticConfigs) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticConfigs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticConfigs(val *StaticConfigs) *NullableStaticConfigs {
	return &NullableStaticConfigs{value: val, isSet: true}
}

func (v NullableStaticConfigs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticConfigs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
