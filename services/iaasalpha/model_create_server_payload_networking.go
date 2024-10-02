/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
	"fmt"
)

// CreateServerPayloadNetworking - The initial networking setup for the server creation. A network, a nic or nothing can be given.
type CreateServerPayloadNetworking struct {
	CreateServerNetworking         *CreateServerNetworking
	CreateServerNetworkingWithNics *CreateServerNetworkingWithNics
}

// CreateServerNetworkingAsCreateServerPayloadNetworking is a convenience function that returns CreateServerNetworking wrapped in CreateServerPayloadNetworking
func CreateServerNetworkingAsCreateServerPayloadNetworking(v *CreateServerNetworking) CreateServerPayloadNetworking {
	return CreateServerPayloadNetworking{
		CreateServerNetworking: v,
	}
}

// CreateServerNetworkingWithNicsAsCreateServerPayloadNetworking is a convenience function that returns CreateServerNetworkingWithNics wrapped in CreateServerPayloadNetworking
func CreateServerNetworkingWithNicsAsCreateServerPayloadNetworking(v *CreateServerNetworkingWithNics) CreateServerPayloadNetworking {
	return CreateServerPayloadNetworking{
		CreateServerNetworkingWithNics: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateServerPayloadNetworking) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateServerNetworking
	err = newStrictDecoder(data).Decode(&dst.CreateServerNetworking)
	if err == nil {
		jsonCreateServerNetworking, _ := json.Marshal(dst.CreateServerNetworking)
		if string(jsonCreateServerNetworking) == "{}" { // empty struct
			dst.CreateServerNetworking = nil
		} else {
			match++
		}
	} else {
		dst.CreateServerNetworking = nil
	}

	// try to unmarshal data into CreateServerNetworkingWithNics
	err = newStrictDecoder(data).Decode(&dst.CreateServerNetworkingWithNics)
	if err == nil {
		jsonCreateServerNetworkingWithNics, _ := json.Marshal(dst.CreateServerNetworkingWithNics)
		if string(jsonCreateServerNetworkingWithNics) == "{}" { // empty struct
			dst.CreateServerNetworkingWithNics = nil
		} else {
			match++
		}
	} else {
		dst.CreateServerNetworkingWithNics = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateServerNetworking = nil
		dst.CreateServerNetworkingWithNics = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateServerPayloadNetworking)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateServerPayloadNetworking)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateServerPayloadNetworking) MarshalJSON() ([]byte, error) {
	if src.CreateServerNetworking != nil {
		return json.Marshal(&src.CreateServerNetworking)
	}

	if src.CreateServerNetworkingWithNics != nil {
		return json.Marshal(&src.CreateServerNetworkingWithNics)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateServerPayloadNetworking) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateServerNetworking != nil {
		return obj.CreateServerNetworking
	}

	if obj.CreateServerNetworkingWithNics != nil {
		return obj.CreateServerNetworkingWithNics
	}

	// all schemas are nil
	return nil
}

type NullableCreateServerPayloadNetworking struct {
	value *CreateServerPayloadNetworking
	isSet bool
}

func (v NullableCreateServerPayloadNetworking) Get() *CreateServerPayloadNetworking {
	return v.value
}

func (v *NullableCreateServerPayloadNetworking) Set(val *CreateServerPayloadNetworking) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServerPayloadNetworking) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServerPayloadNetworking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServerPayloadNetworking(val *CreateServerPayloadNetworking) *NullableCreateServerPayloadNetworking {
	return &NullableCreateServerPayloadNetworking{value: val, isSet: true}
}

func (v NullableCreateServerPayloadNetworking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServerPayloadNetworking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}