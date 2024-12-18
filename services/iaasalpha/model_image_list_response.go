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

// checks if the ImageListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageListResponse{}

// ImageListResponse Image list response.
type ImageListResponse struct {
	// A list containing image objects.
	// REQUIRED
	Items *[]Image `json:"items"`
}

type _ImageListResponse ImageListResponse

// NewImageListResponse instantiates a new ImageListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageListResponse(items *[]Image) *ImageListResponse {
	this := ImageListResponse{}
	this.Items = items
	return &this
}

// NewImageListResponseWithDefaults instantiates a new ImageListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageListResponseWithDefaults() *ImageListResponse {
	this := ImageListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *ImageListResponse) GetItems() *[]Image {
	if o == nil || IsNil(o.Items) {
		var ret *[]Image
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ImageListResponse) GetItemsOk() (*[]Image, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ImageListResponse) SetItems(v *[]Image) {
	o.Items = v
}

func (o ImageListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableImageListResponse struct {
	value *ImageListResponse
	isSet bool
}

func (v NullableImageListResponse) Get() *ImageListResponse {
	return v.value
}

func (v *NullableImageListResponse) Set(val *ImageListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableImageListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableImageListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageListResponse(val *ImageListResponse) *NullableImageListResponse {
	return &NullableImageListResponse{value: val, isSet: true}
}

func (v NullableImageListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
