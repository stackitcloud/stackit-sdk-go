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

// checks if the NetworkRangeListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkRangeListResponse{}

/*
	types and functions for items
*/

// isArray
type NetworkRangeListResponseGetItemsAttributeType = *[]NetworkRange
type NetworkRangeListResponseGetItemsArgType = []NetworkRange
type NetworkRangeListResponseGetItemsRetType = []NetworkRange

func getNetworkRangeListResponseGetItemsAttributeTypeOk(arg NetworkRangeListResponseGetItemsAttributeType) (ret NetworkRangeListResponseGetItemsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNetworkRangeListResponseGetItemsAttributeType(arg *NetworkRangeListResponseGetItemsAttributeType, val NetworkRangeListResponseGetItemsRetType) {
	*arg = &val
}

// NetworkRangeListResponse Network Range list response.
type NetworkRangeListResponse struct {
	// A list of network ranges.
	// REQUIRED
	Items NetworkRangeListResponseGetItemsAttributeType `json:"items"`
}

type _NetworkRangeListResponse NetworkRangeListResponse

// NewNetworkRangeListResponse instantiates a new NetworkRangeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRangeListResponse(items NetworkRangeListResponseGetItemsArgType) *NetworkRangeListResponse {
	this := NetworkRangeListResponse{}
	setNetworkRangeListResponseGetItemsAttributeType(&this.Items, items)
	return &this
}

// NewNetworkRangeListResponseWithDefaults instantiates a new NetworkRangeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkRangeListResponseWithDefaults() *NetworkRangeListResponse {
	this := NetworkRangeListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *NetworkRangeListResponse) GetItems() (ret NetworkRangeListResponseGetItemsRetType) {
	ret, _ = o.GetItemsOk()
	return ret
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *NetworkRangeListResponse) GetItemsOk() (ret NetworkRangeListResponseGetItemsRetType, ok bool) {
	return getNetworkRangeListResponseGetItemsAttributeTypeOk(o.Items)
}

// SetItems sets field value
func (o *NetworkRangeListResponse) SetItems(v NetworkRangeListResponseGetItemsRetType) {
	setNetworkRangeListResponseGetItemsAttributeType(&o.Items, v)
}

func (o NetworkRangeListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getNetworkRangeListResponseGetItemsAttributeTypeOk(o.Items); ok {
		toSerialize["Items"] = val
	}
	return toSerialize, nil
}

type NullableNetworkRangeListResponse struct {
	value *NetworkRangeListResponse
	isSet bool
}

func (v NullableNetworkRangeListResponse) Get() *NetworkRangeListResponse {
	return v.value
}

func (v *NullableNetworkRangeListResponse) Set(val *NetworkRangeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkRangeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkRangeListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkRangeListResponse(val *NetworkRangeListResponse) *NullableNetworkRangeListResponse {
	return &NullableNetworkRangeListResponse{value: val, isSet: true}
}

func (v NullableNetworkRangeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkRangeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
