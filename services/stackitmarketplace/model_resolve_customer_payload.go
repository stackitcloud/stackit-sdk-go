/*
STACKIT Marketplace API

API to manage STACKIT Marketplace.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackitmarketplace

import (
	"encoding/json"
)

// checks if the ResolveCustomerPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResolveCustomerPayload{}

// ResolveCustomerPayload struct for ResolveCustomerPayload
type ResolveCustomerPayload struct {
	// Opaque token exchangeable for subscription details.
	// REQUIRED
	Token *string `json:"token"`
}

type _ResolveCustomerPayload ResolveCustomerPayload

// NewResolveCustomerPayload instantiates a new ResolveCustomerPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolveCustomerPayload(token *string) *ResolveCustomerPayload {
	this := ResolveCustomerPayload{}
	this.Token = token
	return &this
}

// NewResolveCustomerPayloadWithDefaults instantiates a new ResolveCustomerPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolveCustomerPayloadWithDefaults() *ResolveCustomerPayload {
	this := ResolveCustomerPayload{}
	return &this
}

// GetToken returns the Token field value
func (o *ResolveCustomerPayload) GetToken() *string {
	if o == nil || IsNil(o.Token) {
		var ret *string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ResolveCustomerPayload) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token, true
}

// SetToken sets field value
func (o *ResolveCustomerPayload) SetToken(v *string) {
	o.Token = v
}

func (o ResolveCustomerPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableResolveCustomerPayload struct {
	value *ResolveCustomerPayload
	isSet bool
}

func (v NullableResolveCustomerPayload) Get() *ResolveCustomerPayload {
	return v.value
}

func (v *NullableResolveCustomerPayload) Set(val *ResolveCustomerPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableResolveCustomerPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableResolveCustomerPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolveCustomerPayload(val *ResolveCustomerPayload) *NullableResolveCustomerPayload {
	return &NullableResolveCustomerPayload{value: val, isSet: true}
}

func (v NullableResolveCustomerPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolveCustomerPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}