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

// checks if the NICListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NICListResponse{}

// NICListResponse NIC list response.
type NICListResponse struct {
	// A list of network interfaces.
	// REQUIRED
	Items *[]NIC `json:"items"`
}

type _NICListResponse NICListResponse

// NewNICListResponse instantiates a new NICListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNICListResponse(items *[]NIC) *NICListResponse {
	this := NICListResponse{}
	this.Items = items
	return &this
}

// NewNICListResponseWithDefaults instantiates a new NICListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNICListResponseWithDefaults() *NICListResponse {
	this := NICListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *NICListResponse) GetItems() *[]NIC {
	if o == nil || IsNil(o.Items) {
		var ret *[]NIC
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *NICListResponse) GetItemsOk() (*[]NIC, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *NICListResponse) SetItems(v *[]NIC) {
	o.Items = v
}

func (o NICListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableNICListResponse struct {
	value *NICListResponse
	isSet bool
}

func (v NullableNICListResponse) Get() *NICListResponse {
	return v.value
}

func (v *NullableNICListResponse) Set(val *NICListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNICListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNICListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNICListResponse(val *NICListResponse) *NullableNICListResponse {
	return &NullableNICListResponse{value: val, isSet: true}
}

func (v NullableNICListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNICListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
