/*
STACKIT Server Update Management API

API endpoints for Server Update Operations on STACKIT Servers.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverupdate

import (
	"encoding/json"
)

// checks if the GetUpdateSchedulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUpdateSchedulesResponse{}

// GetUpdateSchedulesResponse struct for GetUpdateSchedulesResponse
type GetUpdateSchedulesResponse struct {
	Items *[]UpdateSchedule `json:"items,omitempty"`
}

// NewGetUpdateSchedulesResponse instantiates a new GetUpdateSchedulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUpdateSchedulesResponse() *GetUpdateSchedulesResponse {
	this := GetUpdateSchedulesResponse{}
	return &this
}

// NewGetUpdateSchedulesResponseWithDefaults instantiates a new GetUpdateSchedulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUpdateSchedulesResponseWithDefaults() *GetUpdateSchedulesResponse {
	this := GetUpdateSchedulesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetUpdateSchedulesResponse) GetItems() *[]UpdateSchedule {
	if o == nil || IsNil(o.Items) {
		var ret *[]UpdateSchedule
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUpdateSchedulesResponse) GetItemsOk() (*[]UpdateSchedule, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetUpdateSchedulesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []UpdateSchedule and assigns it to the Items field.
func (o *GetUpdateSchedulesResponse) SetItems(v *[]UpdateSchedule) {
	o.Items = v
}

func (o GetUpdateSchedulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableGetUpdateSchedulesResponse struct {
	value *GetUpdateSchedulesResponse
	isSet bool
}

func (v NullableGetUpdateSchedulesResponse) Get() *GetUpdateSchedulesResponse {
	return v.value
}

func (v *NullableGetUpdateSchedulesResponse) Set(val *GetUpdateSchedulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUpdateSchedulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUpdateSchedulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUpdateSchedulesResponse(val *GetUpdateSchedulesResponse) *NullableGetUpdateSchedulesResponse {
	return &NullableGetUpdateSchedulesResponse{value: val, isSet: true}
}

func (v NullableGetUpdateSchedulesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUpdateSchedulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
