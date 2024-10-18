/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
	"fmt"
)

// AllowedAddressesInner - struct for AllowedAddressesInner
type AllowedAddressesInner struct {
	String *string
}

// stringAsAllowedAddressesInner is a convenience function that returns string wrapped in AllowedAddressesInner
func StringAsAllowedAddressesInner(v *string) AllowedAddressesInner {
	return AllowedAddressesInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AllowedAddressesInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AllowedAddressesInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AllowedAddressesInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AllowedAddressesInner) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AllowedAddressesInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableAllowedAddressesInner struct {
	value *AllowedAddressesInner
	isSet bool
}

func (v NullableAllowedAddressesInner) Get() *AllowedAddressesInner {
	return v.value
}

func (v *NullableAllowedAddressesInner) Set(val *AllowedAddressesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedAddressesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedAddressesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedAddressesInner(val *AllowedAddressesInner) *NullableAllowedAddressesInner {
	return &NullableAllowedAddressesInner{value: val, isSet: true}
}

func (v NullableAllowedAddressesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedAddressesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
