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

// checks if the VolumeListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeListResponse{}

/*
	types and functions for items
*/

// isArray
type VolumeListResponseGetItemsAttributeType = *[]Volume
type VolumeListResponseGetItemsArgType = []Volume
type VolumeListResponseGetItemsRetType = []Volume

func getVolumeListResponseGetItemsAttributeTypeOk(arg VolumeListResponseGetItemsAttributeType) (ret VolumeListResponseGetItemsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setVolumeListResponseGetItemsAttributeType(arg *VolumeListResponseGetItemsAttributeType, val VolumeListResponseGetItemsRetType) {
	*arg = &val
}

// VolumeListResponse Volume list response.
type VolumeListResponse struct {
	// A list containing volume objects.
	// REQUIRED
	Items VolumeListResponseGetItemsAttributeType `json:"items" required:"true"`
}

type _VolumeListResponse VolumeListResponse

// NewVolumeListResponse instantiates a new VolumeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeListResponse(items VolumeListResponseGetItemsArgType) *VolumeListResponse {
	this := VolumeListResponse{}
	setVolumeListResponseGetItemsAttributeType(&this.Items, items)
	return &this
}

// NewVolumeListResponseWithDefaults instantiates a new VolumeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeListResponseWithDefaults() *VolumeListResponse {
	this := VolumeListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *VolumeListResponse) GetItems() (ret VolumeListResponseGetItemsRetType) {
	ret, _ = o.GetItemsOk()
	return ret
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *VolumeListResponse) GetItemsOk() (ret VolumeListResponseGetItemsRetType, ok bool) {
	return getVolumeListResponseGetItemsAttributeTypeOk(o.Items)
}

// SetItems sets field value
func (o *VolumeListResponse) SetItems(v VolumeListResponseGetItemsRetType) {
	setVolumeListResponseGetItemsAttributeType(&o.Items, v)
}

func (o VolumeListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getVolumeListResponseGetItemsAttributeTypeOk(o.Items); ok {
		toSerialize["Items"] = val
	}
	return toSerialize, nil
}

type NullableVolumeListResponse struct {
	value *VolumeListResponse
	isSet bool
}

func (v NullableVolumeListResponse) Get() *VolumeListResponse {
	return v.value
}

func (v *NullableVolumeListResponse) Set(val *VolumeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeListResponse(val *VolumeListResponse) *NullableVolumeListResponse {
	return &NullableVolumeListResponse{value: val, isSet: true}
}

func (v NullableVolumeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
