/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the Label type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Label{}

/*
	types and functions for key
*/

// isNotNullableString
type LabelGetKeyAttributeType = *string

func getLabelGetKeyAttributeTypeOk(arg LabelGetKeyAttributeType) (ret LabelGetKeyRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setLabelGetKeyAttributeType(arg *LabelGetKeyAttributeType, val LabelGetKeyRetType) {
	*arg = &val
}

type LabelGetKeyArgType = string
type LabelGetKeyRetType = string

/*
	types and functions for value
*/

// isNotNullableString
type LabelGetValueAttributeType = *string

func getLabelGetValueAttributeTypeOk(arg LabelGetValueAttributeType) (ret LabelGetValueRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setLabelGetValueAttributeType(arg *LabelGetValueAttributeType, val LabelGetValueRetType) {
	*arg = &val
}

type LabelGetValueArgType = string
type LabelGetValueRetType = string

// Label struct for Label
type Label struct {
	// REQUIRED
	Key LabelGetKeyAttributeType `json:"key"`
	// REQUIRED
	Value LabelGetValueAttributeType `json:"value"`
}

type _Label Label

// NewLabel instantiates a new Label object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabel(key LabelGetKeyArgType, value LabelGetValueArgType) *Label {
	this := Label{}
	setLabelGetKeyAttributeType(&this.Key, key)
	setLabelGetValueAttributeType(&this.Value, value)
	return &this
}

// NewLabelWithDefaults instantiates a new Label object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelWithDefaults() *Label {
	this := Label{}
	return &this
}

// GetKey returns the Key field value
func (o *Label) GetKey() (ret LabelGetKeyRetType) {
	ret, _ = o.GetKeyOk()
	return ret
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Label) GetKeyOk() (ret LabelGetKeyRetType, ok bool) {
	return getLabelGetKeyAttributeTypeOk(o.Key)
}

// SetKey sets field value
func (o *Label) SetKey(v LabelGetKeyRetType) {
	setLabelGetKeyAttributeType(&o.Key, v)
}

// GetValue returns the Value field value
func (o *Label) GetValue() (ret LabelGetValueRetType) {
	ret, _ = o.GetValueOk()
	return ret
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Label) GetValueOk() (ret LabelGetValueRetType, ok bool) {
	return getLabelGetValueAttributeTypeOk(o.Value)
}

// SetValue sets field value
func (o *Label) SetValue(v LabelGetValueRetType) {
	setLabelGetValueAttributeType(&o.Value, v)
}

func (o Label) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getLabelGetKeyAttributeTypeOk(o.Key); ok {
		toSerialize["Key"] = val
	}
	if val, ok := getLabelGetValueAttributeTypeOk(o.Value); ok {
		toSerialize["Value"] = val
	}
	return toSerialize, nil
}

type NullableLabel struct {
	value *Label
	isSet bool
}

func (v NullableLabel) Get() *Label {
	return v.value
}

func (v *NullableLabel) Set(val *Label) {
	v.value = val
	v.isSet = true
}

func (v NullableLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabel(val *Label) *NullableLabel {
	return &NullableLabel{value: val, isSet: true}
}

func (v NullableLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
