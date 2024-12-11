/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

import (
	"encoding/json"
)

// checks if the ACL type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACL{}

// ACL struct for ACL
type ACL struct {
	Items *[]string `json:"items,omitempty"`
}

// NewACL instantiates a new ACL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACL() *ACL {
	this := ACL{}
	return &this
}

// NewACLWithDefaults instantiates a new ACL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACLWithDefaults() *ACL {
	this := ACL{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ACL) GetItems() *[]string {
	if o == nil || IsNil(o.Items) {
		var ret *[]string
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACL) GetItemsOk() (*[]string, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ACL) HasItems() bool {
	if o != nil && !IsNil(o.Items) && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []string and assigns it to the Items field.
func (o *ACL) SetItems(v *[]string) {
	o.Items = v
}

func (o ACL) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableACL struct {
	value *ACL
	isSet bool
}

func (v NullableACL) Get() *ACL {
	return v.value
}

func (v *NullableACL) Set(val *ACL) {
	v.value = val
	v.isSet = true
}

func (v NullableACL) IsSet() bool {
	return v.isSet
}

func (v *NullableACL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACL(val *ACL) *NullableACL {
	return &NullableACL{value: val, isSet: true}
}

func (v NullableACL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
