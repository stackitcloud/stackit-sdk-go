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

// checks if the VolumePerformanceClassListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumePerformanceClassListResponse{}

// VolumePerformanceClassListResponse Volume performance class list response.
type VolumePerformanceClassListResponse struct {
	// A list containing Volume performance classes.
	// REQUIRED
	Items *[]VolumePerformanceClass `json:"items"`
}

type _VolumePerformanceClassListResponse VolumePerformanceClassListResponse

// NewVolumePerformanceClassListResponse instantiates a new VolumePerformanceClassListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumePerformanceClassListResponse(items *[]VolumePerformanceClass) *VolumePerformanceClassListResponse {
	this := VolumePerformanceClassListResponse{}
	this.Items = items
	return &this
}

// NewVolumePerformanceClassListResponseWithDefaults instantiates a new VolumePerformanceClassListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumePerformanceClassListResponseWithDefaults() *VolumePerformanceClassListResponse {
	this := VolumePerformanceClassListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *VolumePerformanceClassListResponse) GetItems() *[]VolumePerformanceClass {
	if o == nil {
		var ret *[]VolumePerformanceClass
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *VolumePerformanceClassListResponse) GetItemsOk() (*[]VolumePerformanceClass, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *VolumePerformanceClassListResponse) SetItems(v *[]VolumePerformanceClass) {
	o.Items = v
}

func (o VolumePerformanceClassListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableVolumePerformanceClassListResponse struct {
	value *VolumePerformanceClassListResponse
	isSet bool
}

func (v NullableVolumePerformanceClassListResponse) Get() *VolumePerformanceClassListResponse {
	return v.value
}

func (v *NullableVolumePerformanceClassListResponse) Set(val *VolumePerformanceClassListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumePerformanceClassListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumePerformanceClassListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumePerformanceClassListResponse(val *VolumePerformanceClassListResponse) *NullableVolumePerformanceClassListResponse {
	return &NullableVolumePerformanceClassListResponse{value: val, isSet: true}
}

func (v NullableVolumePerformanceClassListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumePerformanceClassListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
