/*
STACKIT Run Commands Service API

API endpoints for the STACKIT Run Commands Service API

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package runcommand

import (
	"encoding/json"
)

// checks if the CommandTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommandTemplate{}

// CommandTemplate struct for CommandTemplate
type CommandTemplate struct {
	Name   *string   `json:"name,omitempty"`
	OsType *[]string `json:"osType,omitempty"`
	Title  *string   `json:"title,omitempty"`
}

// NewCommandTemplate instantiates a new CommandTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandTemplate() *CommandTemplate {
	this := CommandTemplate{}
	return &this
}

// NewCommandTemplateWithDefaults instantiates a new CommandTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandTemplateWithDefaults() *CommandTemplate {
	this := CommandTemplate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CommandTemplate) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CommandTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CommandTemplate) SetName(v *string) {
	o.Name = v
}

// GetOsType returns the OsType field value if set, zero value otherwise.
func (o *CommandTemplate) GetOsType() *[]string {
	if o == nil || IsNil(o.OsType) {
		var ret *[]string
		return ret
	}
	return o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandTemplate) GetOsTypeOk() (*[]string, bool) {
	if o == nil || IsNil(o.OsType) {
		return nil, false
	}
	return o.OsType, true
}

// HasOsType returns a boolean if a field has been set.
func (o *CommandTemplate) HasOsType() bool {
	if o != nil && !IsNil(o.OsType) {
		return true
	}

	return false
}

// SetOsType gets a reference to the given []string and assigns it to the OsType field.
func (o *CommandTemplate) SetOsType(v *[]string) {
	o.OsType = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CommandTemplate) GetTitle() *string {
	if o == nil || IsNil(o.Title) {
		var ret *string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandTemplate) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CommandTemplate) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CommandTemplate) SetTitle(v *string) {
	o.Title = v
}

func (o CommandTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OsType) {
		toSerialize["osType"] = o.OsType
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableCommandTemplate struct {
	value *CommandTemplate
	isSet bool
}

func (v NullableCommandTemplate) Get() *CommandTemplate {
	return v.value
}

func (v *NullableCommandTemplate) Set(val *CommandTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandTemplate(val *CommandTemplate) *NullableCommandTemplate {
	return &NullableCommandTemplate{value: val, isSet: true}
}

func (v NullableCommandTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
