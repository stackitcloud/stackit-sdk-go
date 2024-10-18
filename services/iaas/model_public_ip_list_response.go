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

// checks if the PublicIpListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicIpListResponse{}

// PublicIpListResponse Public IP list response.
type PublicIpListResponse struct {
	// A list of public IPs.
	// REQUIRED
	Items *[]PublicIp `json:"items"`
}

type _PublicIpListResponse PublicIpListResponse

// NewPublicIpListResponse instantiates a new PublicIpListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIpListResponse(items *[]PublicIp) *PublicIpListResponse {
	this := PublicIpListResponse{}
	this.Items = items
	return &this
}

// NewPublicIpListResponseWithDefaults instantiates a new PublicIpListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIpListResponseWithDefaults() *PublicIpListResponse {
	this := PublicIpListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *PublicIpListResponse) GetItems() *[]PublicIp {
	if o == nil {
		var ret *[]PublicIp
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PublicIpListResponse) GetItemsOk() (*[]PublicIp, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PublicIpListResponse) SetItems(v *[]PublicIp) {
	o.Items = v
}

func (o PublicIpListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullablePublicIpListResponse struct {
	value *PublicIpListResponse
	isSet bool
}

func (v NullablePublicIpListResponse) Get() *PublicIpListResponse {
	return v.value
}

func (v *NullablePublicIpListResponse) Set(val *PublicIpListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicIpListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicIpListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicIpListResponse(val *PublicIpListResponse) *NullablePublicIpListResponse {
	return &NullablePublicIpListResponse{value: val, isSet: true}
}

func (v NullablePublicIpListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicIpListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
