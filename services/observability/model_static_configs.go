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

/*
	types and functions for labels
*/

// isContainer
type StaticConfigsGetLabelsAttributeType = *map[string]string
type StaticConfigsGetLabelsArgType = map[string]string
type StaticConfigsGetLabelsRetType = map[string]string

func getStaticConfigsGetLabelsAttributeTypeOk(arg StaticConfigsGetLabelsAttributeType) (ret StaticConfigsGetLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setStaticConfigsGetLabelsAttributeType(arg *StaticConfigsGetLabelsAttributeType, val StaticConfigsGetLabelsRetType) {
	*arg = &val
}

/*
	types and functions for targets
*/

// isArray
type StaticConfigsGetTargetsAttributeType = *[]string
type StaticConfigsGetTargetsArgType = []string
type StaticConfigsGetTargetsRetType = []string

func getStaticConfigsGetTargetsAttributeTypeOk(arg StaticConfigsGetTargetsAttributeType) (ret StaticConfigsGetTargetsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setStaticConfigsGetTargetsAttributeType(arg *StaticConfigsGetTargetsAttributeType, val StaticConfigsGetTargetsRetType) {
	*arg = &val
}

// StaticConfigs struct for StaticConfigs
type StaticConfigs struct {
	Labels StaticConfigsGetLabelsAttributeType `json:"labels,omitempty"`
	// REQUIRED
	Targets StaticConfigsGetTargetsAttributeType `json:"targets"`
}

type _StaticConfigs StaticConfigs

// NewStaticConfigs instantiates a new StaticConfigs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticConfigs(targets StaticConfigsGetTargetsArgType) *StaticConfigs {
	this := StaticConfigs{}
	setStaticConfigsGetTargetsAttributeType(&this.Targets, targets)
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
func (o *StaticConfigs) GetLabels() (res StaticConfigsGetLabelsRetType) {
	res, _ = o.GetLabelsOk()
	return
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticConfigs) GetLabelsOk() (ret StaticConfigsGetLabelsRetType, ok bool) {
	return getStaticConfigsGetLabelsAttributeTypeOk(o.Labels)
}

// HasLabels returns a boolean if a field has been set.
func (o *StaticConfigs) HasLabels() bool {
	_, ok := o.GetLabelsOk()
	return ok
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *StaticConfigs) SetLabels(v StaticConfigsGetLabelsRetType) {
	setStaticConfigsGetLabelsAttributeType(&o.Labels, v)
}

// GetTargets returns the Targets field value
func (o *StaticConfigs) GetTargets() (ret StaticConfigsGetTargetsRetType) {
	ret, _ = o.GetTargetsOk()
	return ret
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *StaticConfigs) GetTargetsOk() (ret StaticConfigsGetTargetsRetType, ok bool) {
	return getStaticConfigsGetTargetsAttributeTypeOk(o.Targets)
}

// SetTargets sets field value
func (o *StaticConfigs) SetTargets(v StaticConfigsGetTargetsRetType) {
	setStaticConfigsGetTargetsAttributeType(&o.Targets, v)
}

func (o StaticConfigs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getStaticConfigsGetLabelsAttributeTypeOk(o.Labels); ok {
		toSerialize["Labels"] = val
	}
	if val, ok := getStaticConfigsGetTargetsAttributeTypeOk(o.Targets); ok {
		toSerialize["Targets"] = val
	}
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
