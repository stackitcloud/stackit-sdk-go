/*
STACKIT Marketplace API

API to manage STACKIT Marketplace.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackitmarketplace

import (
	"encoding/json"
	"fmt"
)

// ProductLifecycleState The lifecycle state of the product.
type ProductLifecycleState string

// List of productLifecycleState
const (
	PRODUCTLIFECYCLESTATE_LIVE    ProductLifecycleState = "PRODUCT_LIVE"
	PRODUCTLIFECYCLESTATE_PREVIEW ProductLifecycleState = "PRODUCT_PREVIEW"
)

// All allowed values of ProductLifecycleState enum
var AllowedProductLifecycleStateEnumValues = []ProductLifecycleState{
	"PRODUCT_LIVE",
	"PRODUCT_PREVIEW",
}

func (v *ProductLifecycleState) UnmarshalJSON(src []byte) error {
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
	enumTypeValue := ProductLifecycleState(value)
	for _, existing := range AllowedProductLifecycleStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductLifecycleState", value)
}

// NewProductLifecycleStateFromValue returns a pointer to a valid ProductLifecycleState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductLifecycleStateFromValue(v string) (*ProductLifecycleState, error) {
	ev := ProductLifecycleState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductLifecycleState: valid values are %v", v, AllowedProductLifecycleStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductLifecycleState) IsValid() bool {
	for _, existing := range AllowedProductLifecycleStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to productLifecycleState value
func (v ProductLifecycleState) Ptr() *ProductLifecycleState {
	return &v
}

type NullableProductLifecycleState struct {
	value *ProductLifecycleState
	isSet bool
}

func (v NullableProductLifecycleState) Get() *ProductLifecycleState {
	return v.value
}

func (v *NullableProductLifecycleState) Set(val *ProductLifecycleState) {
	v.value = val
	v.isSet = true
}

func (v NullableProductLifecycleState) IsSet() bool {
	return v.isSet
}

func (v *NullableProductLifecycleState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductLifecycleState(val *ProductLifecycleState) *NullableProductLifecycleState {
	return &NullableProductLifecycleState{value: val, isSet: true}
}

func (v NullableProductLifecycleState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductLifecycleState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
