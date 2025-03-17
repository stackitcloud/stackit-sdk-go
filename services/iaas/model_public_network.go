/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the PublicNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicNetwork{}

/*
	types and functions for cidr
*/

// isNotNullableString
type PublicNetworkGetCidrAttributeType = *string

func getPublicNetworkGetCidrAttributeTypeOk(arg PublicNetworkGetCidrAttributeType) (ret PublicNetworkGetCidrRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPublicNetworkGetCidrAttributeType(arg *PublicNetworkGetCidrAttributeType, val PublicNetworkGetCidrRetType) {
	*arg = &val
}

type PublicNetworkGetCidrArgType = string
type PublicNetworkGetCidrRetType = string

// PublicNetwork Public network.
type PublicNetwork struct {
	// Classless Inter-Domain Routing (CIDR).
	// REQUIRED
	Cidr PublicNetworkGetCidrAttributeType `json:"cidr"`
}

type _PublicNetwork PublicNetwork

// NewPublicNetwork instantiates a new PublicNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicNetwork(cidr PublicNetworkGetCidrArgType) *PublicNetwork {
	this := PublicNetwork{}
	setPublicNetworkGetCidrAttributeType(&this.Cidr, cidr)
	return &this
}

// NewPublicNetworkWithDefaults instantiates a new PublicNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicNetworkWithDefaults() *PublicNetwork {
	this := PublicNetwork{}
	return &this
}

// GetCidr returns the Cidr field value
func (o *PublicNetwork) GetCidr() (ret PublicNetworkGetCidrRetType) {
	ret, _ = o.GetCidrOk()
	return ret
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *PublicNetwork) GetCidrOk() (ret PublicNetworkGetCidrRetType, ok bool) {
	return getPublicNetworkGetCidrAttributeTypeOk(o.Cidr)
}

// SetCidr sets field value
func (o *PublicNetwork) SetCidr(v PublicNetworkGetCidrRetType) {
	setPublicNetworkGetCidrAttributeType(&o.Cidr, v)
}

func (o PublicNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getPublicNetworkGetCidrAttributeTypeOk(o.Cidr); ok {
		toSerialize["Cidr"] = val
	}
	return toSerialize, nil
}

type NullablePublicNetwork struct {
	value *PublicNetwork
	isSet bool
}

func (v NullablePublicNetwork) Get() *PublicNetwork {
	return v.value
}

func (v *NullablePublicNetwork) Set(val *PublicNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicNetwork(val *PublicNetwork) *NullablePublicNetwork {
	return &NullablePublicNetwork{value: val, isSet: true}
}

func (v NullablePublicNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
