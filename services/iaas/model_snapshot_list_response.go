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

// checks if the SnapshotListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotListResponse{}

// SnapshotListResponse Snapshot list response.
type SnapshotListResponse struct {
	// A list containing snapshot objects.
	// REQUIRED
	Items *[]Snapshot `json:"items"`
}

type _SnapshotListResponse SnapshotListResponse

// NewSnapshotListResponse instantiates a new SnapshotListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotListResponse(items *[]Snapshot) *SnapshotListResponse {
	this := SnapshotListResponse{}
	this.Items = items
	return &this
}

// NewSnapshotListResponseWithDefaults instantiates a new SnapshotListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotListResponseWithDefaults() *SnapshotListResponse {
	this := SnapshotListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *SnapshotListResponse) GetItems() *[]Snapshot {
	if o == nil || IsNil(o.Items) {
		var ret *[]Snapshot
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SnapshotListResponse) GetItemsOk() (*[]Snapshot, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SnapshotListResponse) SetItems(v *[]Snapshot) {
	o.Items = v
}

func (o SnapshotListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableSnapshotListResponse struct {
	value *SnapshotListResponse
	isSet bool
}

func (v NullableSnapshotListResponse) Get() *SnapshotListResponse {
	return v.value
}

func (v *NullableSnapshotListResponse) Set(val *SnapshotListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotListResponse(val *SnapshotListResponse) *NullableSnapshotListResponse {
	return &NullableSnapshotListResponse{value: val, isSet: true}
}

func (v NullableSnapshotListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}