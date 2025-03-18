/*
STACKIT LogMe API

The STACKIT LogMe API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logme

import (
	"encoding/json"
)

// checks if the CredentialsListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsListItem{}

/*
	types and functions for id
*/

// isNotNullableString
type CredentialsListItemGetIdAttributeType = *string

func getCredentialsListItemGetIdAttributeTypeOk(arg CredentialsListItemGetIdAttributeType) (ret CredentialsListItemGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsListItemGetIdAttributeType(arg *CredentialsListItemGetIdAttributeType, val CredentialsListItemGetIdRetType) {
	*arg = &val
}

type CredentialsListItemGetIdArgType = string
type CredentialsListItemGetIdRetType = string

// CredentialsListItem struct for CredentialsListItem
type CredentialsListItem struct {
	// REQUIRED
	Id CredentialsListItemGetIdAttributeType `json:"id"`
}

type _CredentialsListItem CredentialsListItem

// NewCredentialsListItem instantiates a new CredentialsListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsListItem(id CredentialsListItemGetIdArgType) *CredentialsListItem {
	this := CredentialsListItem{}
	setCredentialsListItemGetIdAttributeType(&this.Id, id)
	return &this
}

// NewCredentialsListItemWithDefaults instantiates a new CredentialsListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsListItemWithDefaults() *CredentialsListItem {
	this := CredentialsListItem{}
	return &this
}

// GetId returns the Id field value
func (o *CredentialsListItem) GetId() (ret CredentialsListItemGetIdRetType) {
	ret, _ = o.GetIdOk()
	return ret
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CredentialsListItem) GetIdOk() (ret CredentialsListItemGetIdRetType, ok bool) {
	return getCredentialsListItemGetIdAttributeTypeOk(o.Id)
}

// SetId sets field value
func (o *CredentialsListItem) SetId(v CredentialsListItemGetIdRetType) {
	setCredentialsListItemGetIdAttributeType(&o.Id, v)
}

func (o CredentialsListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCredentialsListItemGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	return toSerialize, nil
}

type NullableCredentialsListItem struct {
	value *CredentialsListItem
	isSet bool
}

func (v NullableCredentialsListItem) Get() *CredentialsListItem {
	return v.value
}

func (v *NullableCredentialsListItem) Set(val *CredentialsListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsListItem(val *CredentialsListItem) *NullableCredentialsListItem {
	return &NullableCredentialsListItem{value: val, isSet: true}
}

func (v NullableCredentialsListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
