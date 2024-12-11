/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
)

// checks if the ListInstancesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInstancesResponse{}

// ListInstancesResponse struct for ListInstancesResponse
type ListInstancesResponse struct {
	Count *int64                  `json:"count,omitempty"`
	Items *[]InstanceListInstance `json:"items,omitempty"`
}

// NewListInstancesResponse instantiates a new ListInstancesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInstancesResponse() *ListInstancesResponse {
	this := ListInstancesResponse{}
	return &this
}

// NewListInstancesResponseWithDefaults instantiates a new ListInstancesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInstancesResponseWithDefaults() *ListInstancesResponse {
	this := ListInstancesResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListInstancesResponse) GetCount() *int64 {
	if o == nil || IsNil(o.Count) {
		var ret *int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstancesResponse) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListInstancesResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ListInstancesResponse) SetCount(v *int64) {
	o.Count = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListInstancesResponse) GetItems() *[]InstanceListInstance {
	if o == nil || IsNil(o.Items) {
		var ret *[]InstanceListInstance
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstancesResponse) GetItemsOk() (*[]InstanceListInstance, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListInstancesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InstanceListInstance and assigns it to the Items field.
func (o *ListInstancesResponse) SetItems(v *[]InstanceListInstance) {
	o.Items = v
}

func (o ListInstancesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableListInstancesResponse struct {
	value *ListInstancesResponse
	isSet bool
}

func (v NullableListInstancesResponse) Get() *ListInstancesResponse {
	return v.value
}

func (v *NullableListInstancesResponse) Set(val *ListInstancesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListInstancesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListInstancesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInstancesResponse(val *ListInstancesResponse) *NullableListInstancesResponse {
	return &NullableListInstancesResponse{value: val, isSet: true}
}

func (v NullableListInstancesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInstancesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
