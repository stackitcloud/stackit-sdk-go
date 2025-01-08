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

// checks if the RouteListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteListResponse{}

// RouteListResponse Route list response.
type RouteListResponse struct {
	// A list of routes.
	// REQUIRED
	Items *[]Route `json:"items"`
}

type _RouteListResponse RouteListResponse

// NewRouteListResponse instantiates a new RouteListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteListResponse(items *[]Route) *RouteListResponse {
	this := RouteListResponse{}
	this.Items = items
	return &this
}

// NewRouteListResponseWithDefaults instantiates a new RouteListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteListResponseWithDefaults() *RouteListResponse {
	this := RouteListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *RouteListResponse) GetItems() *[]Route {
	if o == nil || IsNil(o.Items) {
		var ret *[]Route
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *RouteListResponse) GetItemsOk() (*[]Route, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *RouteListResponse) SetItems(v *[]Route) {
	o.Items = v
}

func (o RouteListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableRouteListResponse struct {
	value *RouteListResponse
	isSet bool
}

func (v NullableRouteListResponse) Get() *RouteListResponse {
	return v.value
}

func (v *NullableRouteListResponse) Set(val *RouteListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteListResponse(val *RouteListResponse) *NullableRouteListResponse {
	return &NullableRouteListResponse{value: val, isSet: true}
}

func (v NullableRouteListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}