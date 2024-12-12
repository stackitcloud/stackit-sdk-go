/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
)

// checks if the NetworkAreaListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkAreaListResponse{}

// NetworkAreaListResponse Network area list response.
type NetworkAreaListResponse struct {
	// A list of network areas.
	// REQUIRED
	Items *[]NetworkArea `json:"items"`
}

type _NetworkAreaListResponse NetworkAreaListResponse

// NewNetworkAreaListResponse instantiates a new NetworkAreaListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAreaListResponse(items *[]NetworkArea) *NetworkAreaListResponse {
	this := NetworkAreaListResponse{}
	this.Items = items
	return &this
}

// NewNetworkAreaListResponseWithDefaults instantiates a new NetworkAreaListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAreaListResponseWithDefaults() *NetworkAreaListResponse {
	this := NetworkAreaListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *NetworkAreaListResponse) GetItems() *[]NetworkArea {
	if o == nil || IsNil(o.Items) {
		var ret *[]NetworkArea
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *NetworkAreaListResponse) GetItemsOk() (*[]NetworkArea, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *NetworkAreaListResponse) SetItems(v *[]NetworkArea) {
	o.Items = v
}

func (o NetworkAreaListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableNetworkAreaListResponse struct {
	value *NetworkAreaListResponse
	isSet bool
}

func (v NullableNetworkAreaListResponse) Get() *NetworkAreaListResponse {
	return v.value
}

func (v *NullableNetworkAreaListResponse) Set(val *NetworkAreaListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAreaListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAreaListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAreaListResponse(val *NetworkAreaListResponse) *NullableNetworkAreaListResponse {
	return &NullableNetworkAreaListResponse{value: val, isSet: true}
}

func (v NullableNetworkAreaListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAreaListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
