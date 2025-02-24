/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
)

// checks if the Database type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Database{}

// Database struct for Database
type Database struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Database specific options
	Options *map[string]interface{} `json:"options,omitempty"`
}

// NewDatabase instantiates a new Database object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabase() *Database {
	this := Database{}
	return &this
}

// NewDatabaseWithDefaults instantiates a new Database object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseWithDefaults() *Database {
	this := Database{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Database) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Database) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Database) SetId(v *string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Database) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Database) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Database) SetName(v *string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Database) GetOptions() *map[string]interface{} {
	if o == nil || IsNil(o.Options) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetOptionsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return &map[string]interface{}{}, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Database) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *Database) SetOptions(v *map[string]interface{}) {
	o.Options = v
}

func (o Database) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableDatabase struct {
	value *Database
	isSet bool
}

func (v NullableDatabase) Get() *Database {
	return v.value
}

func (v *NullableDatabase) Set(val *Database) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabase(val *Database) *NullableDatabase {
	return &NullableDatabase{value: val, isSet: true}
}

func (v NullableDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
