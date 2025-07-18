/*
STACKIT Marketplace API

API to manage STACKIT Marketplace.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackitmarketplace

import (
	"encoding/json"
)

// checks if the LocalizedVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalizedVersion{}

/*
	types and functions for de
*/

// isNotNullableString
type LocalizedVersionGetDeAttributeType = *string

func getLocalizedVersionGetDeAttributeTypeOk(arg LocalizedVersionGetDeAttributeType) (ret LocalizedVersionGetDeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setLocalizedVersionGetDeAttributeType(arg *LocalizedVersionGetDeAttributeType, val LocalizedVersionGetDeRetType) {
	*arg = &val
}

type LocalizedVersionGetDeArgType = string
type LocalizedVersionGetDeRetType = string

/*
	types and functions for en
*/

// isNotNullableString
type LocalizedVersionGetEnAttributeType = *string

func getLocalizedVersionGetEnAttributeTypeOk(arg LocalizedVersionGetEnAttributeType) (ret LocalizedVersionGetEnRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setLocalizedVersionGetEnAttributeType(arg *LocalizedVersionGetEnAttributeType, val LocalizedVersionGetEnRetType) {
	*arg = &val
}

type LocalizedVersionGetEnArgType = string
type LocalizedVersionGetEnRetType = string

// LocalizedVersion The localized version (file name) of a file.
type LocalizedVersion struct {
	// The file version matching the file name (localized).
	De LocalizedVersionGetDeAttributeType `json:"de,omitempty"`
	// The file version matching the file name (localized).
	En LocalizedVersionGetEnAttributeType `json:"en,omitempty"`
}

// NewLocalizedVersion instantiates a new LocalizedVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalizedVersion() *LocalizedVersion {
	this := LocalizedVersion{}
	return &this
}

// NewLocalizedVersionWithDefaults instantiates a new LocalizedVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalizedVersionWithDefaults() *LocalizedVersion {
	this := LocalizedVersion{}
	return &this
}

// GetDe returns the De field value if set, zero value otherwise.
func (o *LocalizedVersion) GetDe() (res LocalizedVersionGetDeRetType) {
	res, _ = o.GetDeOk()
	return
}

// GetDeOk returns a tuple with the De field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalizedVersion) GetDeOk() (ret LocalizedVersionGetDeRetType, ok bool) {
	return getLocalizedVersionGetDeAttributeTypeOk(o.De)
}

// HasDe returns a boolean if a field has been set.
func (o *LocalizedVersion) HasDe() bool {
	_, ok := o.GetDeOk()
	return ok
}

// SetDe gets a reference to the given string and assigns it to the De field.
func (o *LocalizedVersion) SetDe(v LocalizedVersionGetDeRetType) {
	setLocalizedVersionGetDeAttributeType(&o.De, v)
}

// GetEn returns the En field value if set, zero value otherwise.
func (o *LocalizedVersion) GetEn() (res LocalizedVersionGetEnRetType) {
	res, _ = o.GetEnOk()
	return
}

// GetEnOk returns a tuple with the En field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalizedVersion) GetEnOk() (ret LocalizedVersionGetEnRetType, ok bool) {
	return getLocalizedVersionGetEnAttributeTypeOk(o.En)
}

// HasEn returns a boolean if a field has been set.
func (o *LocalizedVersion) HasEn() bool {
	_, ok := o.GetEnOk()
	return ok
}

// SetEn gets a reference to the given string and assigns it to the En field.
func (o *LocalizedVersion) SetEn(v LocalizedVersionGetEnRetType) {
	setLocalizedVersionGetEnAttributeType(&o.En, v)
}

func (o LocalizedVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getLocalizedVersionGetDeAttributeTypeOk(o.De); ok {
		toSerialize["De"] = val
	}
	if val, ok := getLocalizedVersionGetEnAttributeTypeOk(o.En); ok {
		toSerialize["En"] = val
	}
	return toSerialize, nil
}

type NullableLocalizedVersion struct {
	value *LocalizedVersion
	isSet bool
}

func (v NullableLocalizedVersion) Get() *LocalizedVersion {
	return v.value
}

func (v *NullableLocalizedVersion) Set(val *LocalizedVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalizedVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalizedVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalizedVersion(val *LocalizedVersion) *NullableLocalizedVersion {
	return &NullableLocalizedVersion{value: val, isSet: true}
}

func (v NullableLocalizedVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalizedVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
