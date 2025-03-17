/*
STACKIT Server Update Management API

API endpoints for Server Update Operations on STACKIT Servers.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverupdate

import (
	"encoding/json"
)

// checks if the GetUpdateSchedulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUpdateSchedulesResponse{}

/*
	types and functions for items
*/

// isArray
type GetUpdateSchedulesResponseGetItemsAttributeType = *[]UpdateSchedule
type GetUpdateSchedulesResponseGetItemsArgType = []UpdateSchedule
type GetUpdateSchedulesResponseGetItemsRetType = []UpdateSchedule

func getGetUpdateSchedulesResponseGetItemsAttributeTypeOk(arg GetUpdateSchedulesResponseGetItemsAttributeType) (ret GetUpdateSchedulesResponseGetItemsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetUpdateSchedulesResponseGetItemsAttributeType(arg *GetUpdateSchedulesResponseGetItemsAttributeType, val GetUpdateSchedulesResponseGetItemsRetType) {
	*arg = &val
}

// GetUpdateSchedulesResponse struct for GetUpdateSchedulesResponse
type GetUpdateSchedulesResponse struct {
	Items GetUpdateSchedulesResponseGetItemsAttributeType `json:"items,omitempty"`
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
func (o *GetUpdateSchedulesResponse) GetItems() (res GetUpdateSchedulesResponseGetItemsRetType) {
	res, _ = o.GetItemsOk()
	return
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUpdateSchedulesResponse) GetItemsOk() (ret GetUpdateSchedulesResponseGetItemsRetType, ok bool) {
	return getGetUpdateSchedulesResponseGetItemsAttributeTypeOk(o.Items)
}

// HasItems returns a boolean if a field has been set.
func (o *GetUpdateSchedulesResponse) HasItems() bool {
	_, ok := o.GetItemsOk()
	return ok
}

// SetItems gets a reference to the given []UpdateSchedule and assigns it to the Items field.
func (o *GetUpdateSchedulesResponse) SetItems(v GetUpdateSchedulesResponseGetItemsRetType) {
	setGetUpdateSchedulesResponseGetItemsAttributeType(&o.Items, v)
}

func (o GetUpdateSchedulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getGetUpdateSchedulesResponseGetItemsAttributeTypeOk(o.Items); ok {
		toSerialize["Items"] = val
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
