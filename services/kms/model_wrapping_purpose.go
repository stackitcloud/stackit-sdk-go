/*
STACKIT Key Management Service API

This API provides endpoints for managing keys and key rings.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kms

import (
	"encoding/json"
	"fmt"
)

// WrappingPurpose The wrapping purpose for the wrapping key.
type WrappingPurpose string

// List of wrappingPurpose
const (
	WRAPPINGPURPOSE_SYMMETRIC_KEY  WrappingPurpose = "wrap_symmetric_key"
	WRAPPINGPURPOSE_ASYMMETRIC_KEY WrappingPurpose = "wrap_asymmetric_key"
)

// All allowed values of WrappingPurpose enum
var AllowedWrappingPurposeEnumValues = []WrappingPurpose{
	"wrap_symmetric_key",
	"wrap_asymmetric_key",
}

func (v *WrappingPurpose) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// Allow unmarshalling zero value for testing purposes
	var zeroValue string
	if value == zeroValue {
		return nil
	}
	enumTypeValue := WrappingPurpose(value)
	for _, existing := range AllowedWrappingPurposeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WrappingPurpose", value)
}

// NewWrappingPurposeFromValue returns a pointer to a valid WrappingPurpose
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWrappingPurposeFromValue(v string) (*WrappingPurpose, error) {
	ev := WrappingPurpose(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WrappingPurpose: valid values are %v", v, AllowedWrappingPurposeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WrappingPurpose) IsValid() bool {
	for _, existing := range AllowedWrappingPurposeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to wrappingPurpose value
func (v WrappingPurpose) Ptr() *WrappingPurpose {
	return &v
}

type NullableWrappingPurpose struct {
	value *WrappingPurpose
	isSet bool
}

func (v NullableWrappingPurpose) Get() *WrappingPurpose {
	return v.value
}

func (v *NullableWrappingPurpose) Set(val *WrappingPurpose) {
	v.value = val
	v.isSet = true
}

func (v NullableWrappingPurpose) IsSet() bool {
	return v.isSet
}

func (v *NullableWrappingPurpose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWrappingPurpose(val *WrappingPurpose) *NullableWrappingPurpose {
	return &NullableWrappingPurpose{value: val, isSet: true}
}

func (v NullableWrappingPurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWrappingPurpose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
