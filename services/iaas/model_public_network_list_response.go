/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the PublicNetworkListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicNetworkListResponse{}

// PublicNetworkListResponse Public network list response.
type PublicNetworkListResponse struct {
	// A list of public networks.
	// REQUIRED
	Items *[]PublicNetwork `json:"items"`
}

type _PublicNetworkListResponse PublicNetworkListResponse

// NewPublicNetworkListResponse instantiates a new PublicNetworkListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicNetworkListResponse(items *[]PublicNetwork) *PublicNetworkListResponse {
	this := PublicNetworkListResponse{}
	this.Items = items
	return &this
}

// NewPublicNetworkListResponseWithDefaults instantiates a new PublicNetworkListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicNetworkListResponseWithDefaults() *PublicNetworkListResponse {
	this := PublicNetworkListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *PublicNetworkListResponse) GetItems() *[]PublicNetwork {
	if o == nil || IsNil(o.Items) {
		var ret *[]PublicNetwork
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PublicNetworkListResponse) GetItemsOk() (*[]PublicNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PublicNetworkListResponse) SetItems(v *[]PublicNetwork) {
	o.Items = v
}

func (o PublicNetworkListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullablePublicNetworkListResponse struct {
	value *PublicNetworkListResponse
	isSet bool
}

func (v NullablePublicNetworkListResponse) Get() *PublicNetworkListResponse {
	return v.value
}

func (v *NullablePublicNetworkListResponse) Set(val *PublicNetworkListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicNetworkListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicNetworkListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicNetworkListResponse(val *PublicNetworkListResponse) *NullablePublicNetworkListResponse {
	return &NullablePublicNetworkListResponse{value: val, isSet: true}
}

func (v NullablePublicNetworkListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicNetworkListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
