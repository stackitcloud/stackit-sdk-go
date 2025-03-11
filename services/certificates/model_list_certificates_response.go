/*
Load Balancer Certificates API

This API offers the ability to store TLS certificates, which can be used by load balancing servers in STACKIT. They can be between consumer and load balancing server and/or between load balancing server and endpoint server.

API version: 2beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package certificates

import (
	"encoding/json"
)

// checks if the ListCertificatesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCertificatesResponse{}

// ListCertificatesResponse ListCertificateResponse returns a list of certificate responses
type ListCertificatesResponse struct {
	Items *[]GetCertificateResponse `json:"items,omitempty"`
	// Continue token from the ListCertificatesResponse with Limit option
	NextPageId *string `json:"nextPageId,omitempty"`
}

// NewListCertificatesResponse instantiates a new ListCertificatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCertificatesResponse() *ListCertificatesResponse {
	this := ListCertificatesResponse{}
	return &this
}

// NewListCertificatesResponseWithDefaults instantiates a new ListCertificatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCertificatesResponseWithDefaults() *ListCertificatesResponse {
	this := ListCertificatesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListCertificatesResponse) GetItems() *[]GetCertificateResponse {
	if o == nil || IsNil(o.Items) {
		var ret *[]GetCertificateResponse
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCertificatesResponse) GetItemsOk() (*[]GetCertificateResponse, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListCertificatesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []GetCertificateResponse and assigns it to the Items field.
func (o *ListCertificatesResponse) SetItems(v *[]GetCertificateResponse) {
	o.Items = v
}

// GetNextPageId returns the NextPageId field value if set, zero value otherwise.
func (o *ListCertificatesResponse) GetNextPageId() *string {
	if o == nil || IsNil(o.NextPageId) {
		var ret *string
		return ret
	}
	return o.NextPageId
}

// GetNextPageIdOk returns a tuple with the NextPageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCertificatesResponse) GetNextPageIdOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageId) {
		return nil, false
	}
	return o.NextPageId, true
}

// HasNextPageId returns a boolean if a field has been set.
func (o *ListCertificatesResponse) HasNextPageId() bool {
	if o != nil && !IsNil(o.NextPageId) {
		return true
	}

	return false
}

// SetNextPageId gets a reference to the given string and assigns it to the NextPageId field.
func (o *ListCertificatesResponse) SetNextPageId(v *string) {
	o.NextPageId = v
}

func (o ListCertificatesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.NextPageId) {
		toSerialize["nextPageId"] = o.NextPageId
	}
	return toSerialize, nil
}

type NullableListCertificatesResponse struct {
	value *ListCertificatesResponse
	isSet bool
}

func (v NullableListCertificatesResponse) Get() *ListCertificatesResponse {
	return v.value
}

func (v *NullableListCertificatesResponse) Set(val *ListCertificatesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCertificatesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCertificatesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCertificatesResponse(val *ListCertificatesResponse) *NullableListCertificatesResponse {
	return &NullableListCertificatesResponse{value: val, isSet: true}
}

func (v NullableListCertificatesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCertificatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
