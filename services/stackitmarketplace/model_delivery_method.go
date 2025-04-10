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

// DeliveryMethod The product delivery method/type. For reference: SAAS - Software as a Service, SAI - STACKIT Application Image
type DeliveryMethod string

// List of deliveryMethod
const (
	DELIVERYMETHOD_SAAS                 DeliveryMethod = "SAAS"
	DELIVERYMETHOD_KUBERNETES           DeliveryMethod = "KUBERNETES"
	DELIVERYMETHOD_SAI                  DeliveryMethod = "SAI"
	DELIVERYMETHOD_PROFESSIONAL_SERVICE DeliveryMethod = "PROFESSIONAL_SERVICE"
)

// All allowed values of DeliveryMethod enum
var AllowedDeliveryMethodEnumValues = []DeliveryMethod{
	"SAAS",
	"KUBERNETES",
	"SAI",
	"PROFESSIONAL_SERVICE",
}

func (v *DeliveryMethod) UnmarshalJSON(src []byte) error {
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
	enumTypeValue := DeliveryMethod(value)
	for _, existing := range AllowedDeliveryMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeliveryMethod", value)
}

// NewDeliveryMethodFromValue returns a pointer to a valid DeliveryMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeliveryMethodFromValue(v string) (*DeliveryMethod, error) {
	ev := DeliveryMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeliveryMethod: valid values are %v", v, AllowedDeliveryMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeliveryMethod) IsValid() bool {
	for _, existing := range AllowedDeliveryMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to deliveryMethod value
func (v DeliveryMethod) Ptr() *DeliveryMethod {
	return &v
}

type NullableDeliveryMethod struct {
	value *DeliveryMethod
	isSet bool
}

func (v NullableDeliveryMethod) Get() *DeliveryMethod {
	return v.value
}

func (v *NullableDeliveryMethod) Set(val *DeliveryMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryMethod(val *DeliveryMethod) *NullableDeliveryMethod {
	return &NullableDeliveryMethod{value: val, isSet: true}
}

func (v NullableDeliveryMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
