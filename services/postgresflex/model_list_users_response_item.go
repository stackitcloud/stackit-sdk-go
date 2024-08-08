/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

import (
	"encoding/json"
)

// checks if the ListUsersResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUsersResponseItem{}

// ListUsersResponseItem struct for ListUsersResponseItem
type ListUsersResponseItem struct {
	Id       *string `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewListUsersResponseItem instantiates a new ListUsersResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUsersResponseItem() *ListUsersResponseItem {
	this := ListUsersResponseItem{}
	return &this
}

// NewListUsersResponseItemWithDefaults instantiates a new ListUsersResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUsersResponseItemWithDefaults() *ListUsersResponseItem {
	this := ListUsersResponseItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListUsersResponseItem) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsersResponseItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListUsersResponseItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListUsersResponseItem) SetId(v *string) {
	o.Id = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ListUsersResponseItem) GetUsername() *string {
	if o == nil || IsNil(o.Username) {
		var ret *string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsersResponseItem) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ListUsersResponseItem) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ListUsersResponseItem) SetUsername(v *string) {
	o.Username = v
}

func (o ListUsersResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableListUsersResponseItem struct {
	value *ListUsersResponseItem
	isSet bool
}

func (v NullableListUsersResponseItem) Get() *ListUsersResponseItem {
	return v.value
}

func (v *NullableListUsersResponseItem) Set(val *ListUsersResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableListUsersResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableListUsersResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUsersResponseItem(val *ListUsersResponseItem) *NullableListUsersResponseItem {
	return &NullableListUsersResponseItem{value: val, isSet: true}
}

func (v NullableListUsersResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUsersResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
