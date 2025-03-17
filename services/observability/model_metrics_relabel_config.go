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

// checks if the MetricsRelabelConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsRelabelConfig{}

/*
	types and functions for action
*/

// isEnumRef
type MetricsRelabelConfigGetActionAttributeType = *string
type MetricsRelabelConfigGetActionArgType = string
type MetricsRelabelConfigGetActionRetType = string

func getMetricsRelabelConfigGetActionAttributeTypeOk(arg MetricsRelabelConfigGetActionAttributeType) (ret MetricsRelabelConfigGetActionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setMetricsRelabelConfigGetActionAttributeType(arg *MetricsRelabelConfigGetActionAttributeType, val MetricsRelabelConfigGetActionRetType) {
	*arg = &val
}

/*
	types and functions for modulus
*/

// isInteger
type MetricsRelabelConfigGetModulusAttributeType = *int64
type MetricsRelabelConfigGetModulusArgType = int64
type MetricsRelabelConfigGetModulusRetType = int64

func getMetricsRelabelConfigGetModulusAttributeTypeOk(arg MetricsRelabelConfigGetModulusAttributeType) (ret MetricsRelabelConfigGetModulusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setMetricsRelabelConfigGetModulusAttributeType(arg *MetricsRelabelConfigGetModulusAttributeType, val MetricsRelabelConfigGetModulusRetType) {
	*arg = &val
}

/*
	types and functions for regex
*/

// isNotNullableString
type MetricsRelabelConfigGetRegexAttributeType = *string

func getMetricsRelabelConfigGetRegexAttributeTypeOk(arg MetricsRelabelConfigGetRegexAttributeType) (ret MetricsRelabelConfigGetRegexRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setMetricsRelabelConfigGetRegexAttributeType(arg *MetricsRelabelConfigGetRegexAttributeType, val MetricsRelabelConfigGetRegexRetType) {
	*arg = &val
}

type MetricsRelabelConfigGetRegexArgType = string
type MetricsRelabelConfigGetRegexRetType = string

/*
	types and functions for replacement
*/

// isNotNullableString
type MetricsRelabelConfigGetReplacementAttributeType = *string

func getMetricsRelabelConfigGetReplacementAttributeTypeOk(arg MetricsRelabelConfigGetReplacementAttributeType) (ret MetricsRelabelConfigGetReplacementRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setMetricsRelabelConfigGetReplacementAttributeType(arg *MetricsRelabelConfigGetReplacementAttributeType, val MetricsRelabelConfigGetReplacementRetType) {
	*arg = &val
}

type MetricsRelabelConfigGetReplacementArgType = string
type MetricsRelabelConfigGetReplacementRetType = string

/*
	types and functions for separator
*/

// isNotNullableString
type MetricsRelabelConfigGetSeparatorAttributeType = *string

func getMetricsRelabelConfigGetSeparatorAttributeTypeOk(arg MetricsRelabelConfigGetSeparatorAttributeType) (ret MetricsRelabelConfigGetSeparatorRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setMetricsRelabelConfigGetSeparatorAttributeType(arg *MetricsRelabelConfigGetSeparatorAttributeType, val MetricsRelabelConfigGetSeparatorRetType) {
	*arg = &val
}

type MetricsRelabelConfigGetSeparatorArgType = string
type MetricsRelabelConfigGetSeparatorRetType = string

/*
	types and functions for sourceLabels
*/

// isArray
type MetricsRelabelConfigGetSourceLabelsAttributeType = *[]string
type MetricsRelabelConfigGetSourceLabelsArgType = []string
type MetricsRelabelConfigGetSourceLabelsRetType = []string

func getMetricsRelabelConfigGetSourceLabelsAttributeTypeOk(arg MetricsRelabelConfigGetSourceLabelsAttributeType) (ret MetricsRelabelConfigGetSourceLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setMetricsRelabelConfigGetSourceLabelsAttributeType(arg *MetricsRelabelConfigGetSourceLabelsAttributeType, val MetricsRelabelConfigGetSourceLabelsRetType) {
	*arg = &val
}

/*
	types and functions for targetLabel
*/

// isNotNullableString
type MetricsRelabelConfigGetTargetLabelAttributeType = *string

func getMetricsRelabelConfigGetTargetLabelAttributeTypeOk(arg MetricsRelabelConfigGetTargetLabelAttributeType) (ret MetricsRelabelConfigGetTargetLabelRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setMetricsRelabelConfigGetTargetLabelAttributeType(arg *MetricsRelabelConfigGetTargetLabelAttributeType, val MetricsRelabelConfigGetTargetLabelRetType) {
	*arg = &val
}

type MetricsRelabelConfigGetTargetLabelArgType = string
type MetricsRelabelConfigGetTargetLabelRetType = string

// MetricsRelabelConfig struct for MetricsRelabelConfig
type MetricsRelabelConfig struct {
	Action      MetricsRelabelConfigGetActionAttributeType      `json:"action,omitempty"`
	Modulus     MetricsRelabelConfigGetModulusAttributeType     `json:"modulus,omitempty"`
	Regex       MetricsRelabelConfigGetRegexAttributeType       `json:"regex,omitempty"`
	Replacement MetricsRelabelConfigGetReplacementAttributeType `json:"replacement,omitempty"`
	Separator   MetricsRelabelConfigGetSeparatorAttributeType   `json:"separator,omitempty"`
	// REQUIRED
	SourceLabels MetricsRelabelConfigGetSourceLabelsAttributeType `json:"sourceLabels"`
	TargetLabel  MetricsRelabelConfigGetTargetLabelAttributeType  `json:"targetLabel,omitempty"`
}

type _MetricsRelabelConfig MetricsRelabelConfig

// NewMetricsRelabelConfig instantiates a new MetricsRelabelConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsRelabelConfig(sourceLabels MetricsRelabelConfigGetSourceLabelsArgType) *MetricsRelabelConfig {
	this := MetricsRelabelConfig{}
	setMetricsRelabelConfigGetSourceLabelsAttributeType(&this.SourceLabels, sourceLabels)
	return &this
}

// NewMetricsRelabelConfigWithDefaults instantiates a new MetricsRelabelConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsRelabelConfigWithDefaults() *MetricsRelabelConfig {
	this := MetricsRelabelConfig{}
	var action string = "replace"
	this.Action = &action
	var regex string = ".*"
	this.Regex = &regex
	var replacement string = "$1"
	this.Replacement = &replacement
	var separator string = ";"
	this.Separator = &separator
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *MetricsRelabelConfig) GetAction() (res MetricsRelabelConfigGetActionRetType) {
	res, _ = o.GetActionOk()
	return
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRelabelConfig) GetActionOk() (ret MetricsRelabelConfigGetActionRetType, ok bool) {
	return getMetricsRelabelConfigGetActionAttributeTypeOk(o.Action)
}

// HasAction returns a boolean if a field has been set.
func (o *MetricsRelabelConfig) HasAction() bool {
	_, ok := o.GetActionOk()
	return ok
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *MetricsRelabelConfig) SetAction(v MetricsRelabelConfigGetActionRetType) {
	setMetricsRelabelConfigGetActionAttributeType(&o.Action, v)
}

// GetModulus returns the Modulus field value if set, zero value otherwise.
func (o *MetricsRelabelConfig) GetModulus() (res MetricsRelabelConfigGetModulusRetType) {
	res, _ = o.GetModulusOk()
	return
}

// GetModulusOk returns a tuple with the Modulus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRelabelConfig) GetModulusOk() (ret MetricsRelabelConfigGetModulusRetType, ok bool) {
	return getMetricsRelabelConfigGetModulusAttributeTypeOk(o.Modulus)
}

// HasModulus returns a boolean if a field has been set.
func (o *MetricsRelabelConfig) HasModulus() bool {
	_, ok := o.GetModulusOk()
	return ok
}

// SetModulus gets a reference to the given int64 and assigns it to the Modulus field.
func (o *MetricsRelabelConfig) SetModulus(v MetricsRelabelConfigGetModulusRetType) {
	setMetricsRelabelConfigGetModulusAttributeType(&o.Modulus, v)
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *MetricsRelabelConfig) GetRegex() (res MetricsRelabelConfigGetRegexRetType) {
	res, _ = o.GetRegexOk()
	return
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRelabelConfig) GetRegexOk() (ret MetricsRelabelConfigGetRegexRetType, ok bool) {
	return getMetricsRelabelConfigGetRegexAttributeTypeOk(o.Regex)
}

// HasRegex returns a boolean if a field has been set.
func (o *MetricsRelabelConfig) HasRegex() bool {
	_, ok := o.GetRegexOk()
	return ok
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *MetricsRelabelConfig) SetRegex(v MetricsRelabelConfigGetRegexRetType) {
	setMetricsRelabelConfigGetRegexAttributeType(&o.Regex, v)
}

// GetReplacement returns the Replacement field value if set, zero value otherwise.
func (o *MetricsRelabelConfig) GetReplacement() (res MetricsRelabelConfigGetReplacementRetType) {
	res, _ = o.GetReplacementOk()
	return
}

// GetReplacementOk returns a tuple with the Replacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRelabelConfig) GetReplacementOk() (ret MetricsRelabelConfigGetReplacementRetType, ok bool) {
	return getMetricsRelabelConfigGetReplacementAttributeTypeOk(o.Replacement)
}

// HasReplacement returns a boolean if a field has been set.
func (o *MetricsRelabelConfig) HasReplacement() bool {
	_, ok := o.GetReplacementOk()
	return ok
}

// SetReplacement gets a reference to the given string and assigns it to the Replacement field.
func (o *MetricsRelabelConfig) SetReplacement(v MetricsRelabelConfigGetReplacementRetType) {
	setMetricsRelabelConfigGetReplacementAttributeType(&o.Replacement, v)
}

// GetSeparator returns the Separator field value if set, zero value otherwise.
func (o *MetricsRelabelConfig) GetSeparator() (res MetricsRelabelConfigGetSeparatorRetType) {
	res, _ = o.GetSeparatorOk()
	return
}

// GetSeparatorOk returns a tuple with the Separator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRelabelConfig) GetSeparatorOk() (ret MetricsRelabelConfigGetSeparatorRetType, ok bool) {
	return getMetricsRelabelConfigGetSeparatorAttributeTypeOk(o.Separator)
}

// HasSeparator returns a boolean if a field has been set.
func (o *MetricsRelabelConfig) HasSeparator() bool {
	_, ok := o.GetSeparatorOk()
	return ok
}

// SetSeparator gets a reference to the given string and assigns it to the Separator field.
func (o *MetricsRelabelConfig) SetSeparator(v MetricsRelabelConfigGetSeparatorRetType) {
	setMetricsRelabelConfigGetSeparatorAttributeType(&o.Separator, v)
}

// GetSourceLabels returns the SourceLabels field value
func (o *MetricsRelabelConfig) GetSourceLabels() (ret MetricsRelabelConfigGetSourceLabelsRetType) {
	ret, _ = o.GetSourceLabelsOk()
	return ret
}

// GetSourceLabelsOk returns a tuple with the SourceLabels field value
// and a boolean to check if the value has been set.
func (o *MetricsRelabelConfig) GetSourceLabelsOk() (ret MetricsRelabelConfigGetSourceLabelsRetType, ok bool) {
	return getMetricsRelabelConfigGetSourceLabelsAttributeTypeOk(o.SourceLabels)
}

// SetSourceLabels sets field value
func (o *MetricsRelabelConfig) SetSourceLabels(v MetricsRelabelConfigGetSourceLabelsRetType) {
	setMetricsRelabelConfigGetSourceLabelsAttributeType(&o.SourceLabels, v)
}

// GetTargetLabel returns the TargetLabel field value if set, zero value otherwise.
func (o *MetricsRelabelConfig) GetTargetLabel() (res MetricsRelabelConfigGetTargetLabelRetType) {
	res, _ = o.GetTargetLabelOk()
	return
}

// GetTargetLabelOk returns a tuple with the TargetLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsRelabelConfig) GetTargetLabelOk() (ret MetricsRelabelConfigGetTargetLabelRetType, ok bool) {
	return getMetricsRelabelConfigGetTargetLabelAttributeTypeOk(o.TargetLabel)
}

// HasTargetLabel returns a boolean if a field has been set.
func (o *MetricsRelabelConfig) HasTargetLabel() bool {
	_, ok := o.GetTargetLabelOk()
	return ok
}

// SetTargetLabel gets a reference to the given string and assigns it to the TargetLabel field.
func (o *MetricsRelabelConfig) SetTargetLabel(v MetricsRelabelConfigGetTargetLabelRetType) {
	setMetricsRelabelConfigGetTargetLabelAttributeType(&o.TargetLabel, v)
}

func (o MetricsRelabelConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getMetricsRelabelConfigGetActionAttributeTypeOk(o.Action); ok {
		toSerialize["Action"] = val
	}
	if val, ok := getMetricsRelabelConfigGetModulusAttributeTypeOk(o.Modulus); ok {
		toSerialize["Modulus"] = val
	}
	if val, ok := getMetricsRelabelConfigGetRegexAttributeTypeOk(o.Regex); ok {
		toSerialize["Regex"] = val
	}
	if val, ok := getMetricsRelabelConfigGetReplacementAttributeTypeOk(o.Replacement); ok {
		toSerialize["Replacement"] = val
	}
	if val, ok := getMetricsRelabelConfigGetSeparatorAttributeTypeOk(o.Separator); ok {
		toSerialize["Separator"] = val
	}
	if val, ok := getMetricsRelabelConfigGetSourceLabelsAttributeTypeOk(o.SourceLabels); ok {
		toSerialize["SourceLabels"] = val
	}
	if val, ok := getMetricsRelabelConfigGetTargetLabelAttributeTypeOk(o.TargetLabel); ok {
		toSerialize["TargetLabel"] = val
	}
	return toSerialize, nil
}

type NullableMetricsRelabelConfig struct {
	value *MetricsRelabelConfig
	isSet bool
}

func (v NullableMetricsRelabelConfig) Get() *MetricsRelabelConfig {
	return v.value
}

func (v *NullableMetricsRelabelConfig) Set(val *MetricsRelabelConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsRelabelConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsRelabelConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsRelabelConfig(val *MetricsRelabelConfig) *NullableMetricsRelabelConfig {
	return &NullableMetricsRelabelConfig{value: val, isSet: true}
}

func (v NullableMetricsRelabelConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsRelabelConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
