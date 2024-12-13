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

// checks if the ServiceAccountMailListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountMailListResponse{}

// ServiceAccountMailListResponse Service account mail list response.
type ServiceAccountMailListResponse struct {
	// A list of service account mails.
	// REQUIRED
	Items *[]string `json:"items"`
}

type _ServiceAccountMailListResponse ServiceAccountMailListResponse

// NewServiceAccountMailListResponse instantiates a new ServiceAccountMailListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountMailListResponse(items *[]string) *ServiceAccountMailListResponse {
	this := ServiceAccountMailListResponse{}
	this.Items = items
	return &this
}

// NewServiceAccountMailListResponseWithDefaults instantiates a new ServiceAccountMailListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountMailListResponseWithDefaults() *ServiceAccountMailListResponse {
	this := ServiceAccountMailListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *ServiceAccountMailListResponse) GetItems() *[]string {
	if o == nil || IsNil(o.Items) {
		var ret *[]string
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountMailListResponse) GetItemsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ServiceAccountMailListResponse) SetItems(v *[]string) {
	o.Items = v
}

func (o ServiceAccountMailListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableServiceAccountMailListResponse struct {
	value *ServiceAccountMailListResponse
	isSet bool
}

func (v NullableServiceAccountMailListResponse) Get() *ServiceAccountMailListResponse {
	return v.value
}

func (v *NullableServiceAccountMailListResponse) Set(val *ServiceAccountMailListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountMailListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountMailListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountMailListResponse(val *ServiceAccountMailListResponse) *NullableServiceAccountMailListResponse {
	return &NullableServiceAccountMailListResponse{value: val, isSet: true}
}

func (v NullableServiceAccountMailListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountMailListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
