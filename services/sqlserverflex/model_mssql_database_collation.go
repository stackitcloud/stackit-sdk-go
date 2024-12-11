/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
)

// checks if the MssqlDatabaseCollation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MssqlDatabaseCollation{}

// MssqlDatabaseCollation struct for MssqlDatabaseCollation
type MssqlDatabaseCollation struct {
	CollationName *string `json:"collation_name,omitempty"`
	Description   *string `json:"description,omitempty"`
}

// NewMssqlDatabaseCollation instantiates a new MssqlDatabaseCollation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMssqlDatabaseCollation() *MssqlDatabaseCollation {
	this := MssqlDatabaseCollation{}
	return &this
}

// NewMssqlDatabaseCollationWithDefaults instantiates a new MssqlDatabaseCollation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMssqlDatabaseCollationWithDefaults() *MssqlDatabaseCollation {
	this := MssqlDatabaseCollation{}
	return &this
}

// GetCollationName returns the CollationName field value if set, zero value otherwise.
func (o *MssqlDatabaseCollation) GetCollationName() *string {
	if o == nil || IsNil(o.CollationName) {
		var ret *string
		return ret
	}
	return o.CollationName
}

// GetCollationNameOk returns a tuple with the CollationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MssqlDatabaseCollation) GetCollationNameOk() (*string, bool) {
	if o == nil || IsNil(o.CollationName) {
		return nil, false
	}
	return o.CollationName, true
}

// HasCollationName returns a boolean if a field has been set.
func (o *MssqlDatabaseCollation) HasCollationName() bool {
	if o != nil && !IsNil(o.CollationName) && !IsNil(o.CollationName) {
		return true
	}

	return false
}

// SetCollationName gets a reference to the given string and assigns it to the CollationName field.
func (o *MssqlDatabaseCollation) SetCollationName(v *string) {
	o.CollationName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MssqlDatabaseCollation) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MssqlDatabaseCollation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MssqlDatabaseCollation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MssqlDatabaseCollation) SetDescription(v *string) {
	o.Description = v
}

func (o MssqlDatabaseCollation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollationName) {
		toSerialize["collation_name"] = o.CollationName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableMssqlDatabaseCollation struct {
	value *MssqlDatabaseCollation
	isSet bool
}

func (v NullableMssqlDatabaseCollation) Get() *MssqlDatabaseCollation {
	return v.value
}

func (v *NullableMssqlDatabaseCollation) Set(val *MssqlDatabaseCollation) {
	v.value = val
	v.isSet = true
}

func (v NullableMssqlDatabaseCollation) IsSet() bool {
	return v.isSet
}

func (v *NullableMssqlDatabaseCollation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMssqlDatabaseCollation(val *MssqlDatabaseCollation) *NullableMssqlDatabaseCollation {
	return &NullableMssqlDatabaseCollation{value: val, isSet: true}
}

func (v NullableMssqlDatabaseCollation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMssqlDatabaseCollation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
