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

// checks if the TraceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceConfig{}

/*
	types and functions for retention
*/

// isNotNullableString
type TraceConfigGetRetentionAttributeType = *string

func getTraceConfigGetRetentionAttributeTypeOk(arg TraceConfigGetRetentionAttributeType) (ret TraceConfigGetRetentionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setTraceConfigGetRetentionAttributeType(arg *TraceConfigGetRetentionAttributeType, val TraceConfigGetRetentionRetType) {
	*arg = &val
}

type TraceConfigGetRetentionArgType = string
type TraceConfigGetRetentionRetType = string

// TraceConfig struct for TraceConfig
type TraceConfig struct {
	// REQUIRED
	Retention TraceConfigGetRetentionAttributeType `json:"retention" required:"true"`
}

type _TraceConfig TraceConfig

// NewTraceConfig instantiates a new TraceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceConfig(retention TraceConfigGetRetentionArgType) *TraceConfig {
	this := TraceConfig{}
	setTraceConfigGetRetentionAttributeType(&this.Retention, retention)
	return &this
}

// NewTraceConfigWithDefaults instantiates a new TraceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceConfigWithDefaults() *TraceConfig {
	this := TraceConfig{}
	return &this
}

// GetRetention returns the Retention field value
func (o *TraceConfig) GetRetention() (ret TraceConfigGetRetentionRetType) {
	ret, _ = o.GetRetentionOk()
	return ret
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *TraceConfig) GetRetentionOk() (ret TraceConfigGetRetentionRetType, ok bool) {
	return getTraceConfigGetRetentionAttributeTypeOk(o.Retention)
}

// SetRetention sets field value
func (o *TraceConfig) SetRetention(v TraceConfigGetRetentionRetType) {
	setTraceConfigGetRetentionAttributeType(&o.Retention, v)
}

func (o TraceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getTraceConfigGetRetentionAttributeTypeOk(o.Retention); ok {
		toSerialize["Retention"] = val
	}
	return toSerialize, nil
}

type NullableTraceConfig struct {
	value *TraceConfig
	isSet bool
}

func (v NullableTraceConfig) Get() *TraceConfig {
	return v.value
}

func (v *NullableTraceConfig) Set(val *TraceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceConfig(val *TraceConfig) *NullableTraceConfig {
	return &NullableTraceConfig{value: val, isSet: true}
}

func (v NullableTraceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
