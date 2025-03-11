/*
STACKIT Server Backup Management API

API endpoints for Server Backup Operations on STACKIT Servers.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverbackup

import (
	"encoding/json"
)

// checks if the GetBackupsListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBackupsListResponse{}

// GetBackupsListResponse struct for GetBackupsListResponse
type GetBackupsListResponse struct {
	Items *[]Backup `json:"items,omitempty"`
}

// NewGetBackupsListResponse instantiates a new GetBackupsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBackupsListResponse() *GetBackupsListResponse {
	this := GetBackupsListResponse{}
	return &this
}

// NewGetBackupsListResponseWithDefaults instantiates a new GetBackupsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBackupsListResponseWithDefaults() *GetBackupsListResponse {
	this := GetBackupsListResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetBackupsListResponse) GetItems() *[]Backup {
	if o == nil || IsNil(o.Items) {
		var ret *[]Backup
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBackupsListResponse) GetItemsOk() (*[]Backup, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetBackupsListResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Backup and assigns it to the Items field.
func (o *GetBackupsListResponse) SetItems(v *[]Backup) {
	o.Items = v
}

func (o GetBackupsListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableGetBackupsListResponse struct {
	value *GetBackupsListResponse
	isSet bool
}

func (v NullableGetBackupsListResponse) Get() *GetBackupsListResponse {
	return v.value
}

func (v *NullableGetBackupsListResponse) Set(val *GetBackupsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBackupsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBackupsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBackupsListResponse(val *GetBackupsListResponse) *NullableGetBackupsListResponse {
	return &NullableGetBackupsListResponse{value: val, isSet: true}
}

func (v NullableGetBackupsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBackupsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
