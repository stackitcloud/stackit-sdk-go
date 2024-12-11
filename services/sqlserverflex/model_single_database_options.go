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

// checks if the SingleDatabaseOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleDatabaseOptions{}

// SingleDatabaseOptions Database specific options
type SingleDatabaseOptions struct {
	// Name of the collation of the database
	CollationName *string `json:"collationName,omitempty"`
	// CompatibilityLevel of the Database.
	CompatibilityLevel *int64 `json:"compatibilityLevel,omitempty"`
	// Name of the owner of the database.
	Owner *string `json:"owner,omitempty"`
}

// NewSingleDatabaseOptions instantiates a new SingleDatabaseOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleDatabaseOptions() *SingleDatabaseOptions {
	this := SingleDatabaseOptions{}
	return &this
}

// NewSingleDatabaseOptionsWithDefaults instantiates a new SingleDatabaseOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleDatabaseOptionsWithDefaults() *SingleDatabaseOptions {
	this := SingleDatabaseOptions{}
	return &this
}

// GetCollationName returns the CollationName field value if set, zero value otherwise.
func (o *SingleDatabaseOptions) GetCollationName() *string {
	if o == nil || IsNil(o.CollationName) {
		var ret *string
		return ret
	}
	return o.CollationName
}

// GetCollationNameOk returns a tuple with the CollationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleDatabaseOptions) GetCollationNameOk() (*string, bool) {
	if o == nil || IsNil(o.CollationName) {
		return nil, false
	}
	return o.CollationName, true
}

// HasCollationName returns a boolean if a field has been set.
func (o *SingleDatabaseOptions) HasCollationName() bool {
	if o != nil && !IsNil(o.CollationName) && !IsNil(o.CollationName) {
		return true
	}

	return false
}

// SetCollationName gets a reference to the given string and assigns it to the CollationName field.
func (o *SingleDatabaseOptions) SetCollationName(v *string) {
	o.CollationName = v
}

// GetCompatibilityLevel returns the CompatibilityLevel field value if set, zero value otherwise.
func (o *SingleDatabaseOptions) GetCompatibilityLevel() *int64 {
	if o == nil || IsNil(o.CompatibilityLevel) {
		var ret *int64
		return ret
	}
	return o.CompatibilityLevel
}

// GetCompatibilityLevelOk returns a tuple with the CompatibilityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleDatabaseOptions) GetCompatibilityLevelOk() (*int64, bool) {
	if o == nil || IsNil(o.CompatibilityLevel) {
		return nil, false
	}
	return o.CompatibilityLevel, true
}

// HasCompatibilityLevel returns a boolean if a field has been set.
func (o *SingleDatabaseOptions) HasCompatibilityLevel() bool {
	if o != nil && !IsNil(o.CompatibilityLevel) && !IsNil(o.CompatibilityLevel) {
		return true
	}

	return false
}

// SetCompatibilityLevel gets a reference to the given int64 and assigns it to the CompatibilityLevel field.
func (o *SingleDatabaseOptions) SetCompatibilityLevel(v *int64) {
	o.CompatibilityLevel = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *SingleDatabaseOptions) GetOwner() *string {
	if o == nil || IsNil(o.Owner) {
		var ret *string
		return ret
	}
	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleDatabaseOptions) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *SingleDatabaseOptions) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *SingleDatabaseOptions) SetOwner(v *string) {
	o.Owner = v
}

func (o SingleDatabaseOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollationName) {
		toSerialize["collationName"] = o.CollationName
	}
	if !IsNil(o.CompatibilityLevel) {
		toSerialize["compatibilityLevel"] = o.CompatibilityLevel
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableSingleDatabaseOptions struct {
	value *SingleDatabaseOptions
	isSet bool
}

func (v NullableSingleDatabaseOptions) Get() *SingleDatabaseOptions {
	return v.value
}

func (v *NullableSingleDatabaseOptions) Set(val *SingleDatabaseOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleDatabaseOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleDatabaseOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleDatabaseOptions(val *SingleDatabaseOptions) *NullableSingleDatabaseOptions {
	return &NullableSingleDatabaseOptions{value: val, isSet: true}
}

func (v NullableSingleDatabaseOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleDatabaseOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
