/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 2alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
	"fmt"
)

// CreateNetworkIPv6 - The create request for an IPv6 network.
type CreateNetworkIPv6 struct {
	CreateNetworkIPv6WithPrefix       *CreateNetworkIPv6WithPrefix
	CreateNetworkIPv6WithPrefixLength *CreateNetworkIPv6WithPrefixLength
}

// CreateNetworkIPv6WithPrefixAsCreateNetworkIPv6 is a convenience function that returns CreateNetworkIPv6WithPrefix wrapped in CreateNetworkIPv6
func CreateNetworkIPv6WithPrefixAsCreateNetworkIPv6(v *CreateNetworkIPv6WithPrefix) CreateNetworkIPv6 {
	return CreateNetworkIPv6{
		CreateNetworkIPv6WithPrefix: v,
	}
}

// CreateNetworkIPv6WithPrefixLengthAsCreateNetworkIPv6 is a convenience function that returns CreateNetworkIPv6WithPrefixLength wrapped in CreateNetworkIPv6
func CreateNetworkIPv6WithPrefixLengthAsCreateNetworkIPv6(v *CreateNetworkIPv6WithPrefixLength) CreateNetworkIPv6 {
	return CreateNetworkIPv6{
		CreateNetworkIPv6WithPrefixLength: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateNetworkIPv6) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateNetworkIPv6WithPrefix
	err = json.Unmarshal(data, &dst.CreateNetworkIPv6WithPrefix)
	if err == nil {
		jsonCreateNetworkIPv6WithPrefix, _ := json.Marshal(dst.CreateNetworkIPv6WithPrefix)
		if string(jsonCreateNetworkIPv6WithPrefix) == "{}" { // empty struct
			dst.CreateNetworkIPv6WithPrefix = nil
		} else {
			match++
		}
	} else {
		dst.CreateNetworkIPv6WithPrefix = nil
	}

	// try to unmarshal data into CreateNetworkIPv6WithPrefixLength
	err = json.Unmarshal(data, &dst.CreateNetworkIPv6WithPrefixLength)
	if err == nil {
		jsonCreateNetworkIPv6WithPrefixLength, _ := json.Marshal(dst.CreateNetworkIPv6WithPrefixLength)
		if string(jsonCreateNetworkIPv6WithPrefixLength) == "{}" { // empty struct
			dst.CreateNetworkIPv6WithPrefixLength = nil
		} else {
			match++
		}
	} else {
		dst.CreateNetworkIPv6WithPrefixLength = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateNetworkIPv6WithPrefix = nil
		dst.CreateNetworkIPv6WithPrefixLength = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateNetworkIPv6)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateNetworkIPv6)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateNetworkIPv6) MarshalJSON() ([]byte, error) {
	if src.CreateNetworkIPv6WithPrefix != nil {
		return json.Marshal(&src.CreateNetworkIPv6WithPrefix)
	}

	if src.CreateNetworkIPv6WithPrefixLength != nil {
		return json.Marshal(&src.CreateNetworkIPv6WithPrefixLength)
	}

	return []byte("{}"), nil // no data in oneOf schemas => empty JSON object
}

// Get the actual instance
func (obj *CreateNetworkIPv6) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateNetworkIPv6WithPrefix != nil {
		return obj.CreateNetworkIPv6WithPrefix
	}

	if obj.CreateNetworkIPv6WithPrefixLength != nil {
		return obj.CreateNetworkIPv6WithPrefixLength
	}

	// all schemas are nil
	return nil
}

type NullableCreateNetworkIPv6 struct {
	value *CreateNetworkIPv6
	isSet bool
}

func (v NullableCreateNetworkIPv6) Get() *CreateNetworkIPv6 {
	return v.value
}

func (v *NullableCreateNetworkIPv6) Set(val *CreateNetworkIPv6) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkIPv6) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkIPv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkIPv6(val *CreateNetworkIPv6) *NullableCreateNetworkIPv6 {
	return &NullableCreateNetworkIPv6{value: val, isSet: true}
}

func (v NullableCreateNetworkIPv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkIPv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
