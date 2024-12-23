/*
STACKIT Server Backup Management API

API endpoints for Server Backup Operations on STACKIT Servers.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverbackup

import (
	"encoding/json"
)

// checks if the GetBackupSchedulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBackupSchedulesResponse{}

// GetBackupSchedulesResponse struct for GetBackupSchedulesResponse
type GetBackupSchedulesResponse struct {
	Items *[]BackupSchedule `json:"items,omitempty"`
}

// NewGetBackupSchedulesResponse instantiates a new GetBackupSchedulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBackupSchedulesResponse() *GetBackupSchedulesResponse {
	this := GetBackupSchedulesResponse{}
	return &this
}

// NewGetBackupSchedulesResponseWithDefaults instantiates a new GetBackupSchedulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBackupSchedulesResponseWithDefaults() *GetBackupSchedulesResponse {
	this := GetBackupSchedulesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetBackupSchedulesResponse) GetItems() *[]BackupSchedule {
	if o == nil || IsNil(o.Items) {
		var ret *[]BackupSchedule
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBackupSchedulesResponse) GetItemsOk() (*[]BackupSchedule, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetBackupSchedulesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BackupSchedule and assigns it to the Items field.
func (o *GetBackupSchedulesResponse) SetItems(v *[]BackupSchedule) {
	o.Items = v
}

func (o GetBackupSchedulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableGetBackupSchedulesResponse struct {
	value *GetBackupSchedulesResponse
	isSet bool
}

func (v NullableGetBackupSchedulesResponse) Get() *GetBackupSchedulesResponse {
	return v.value
}

func (v *NullableGetBackupSchedulesResponse) Set(val *GetBackupSchedulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBackupSchedulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBackupSchedulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBackupSchedulesResponse(val *GetBackupSchedulesResponse) *NullableGetBackupSchedulesResponse {
	return &NullableGetBackupSchedulesResponse{value: val, isSet: true}
}

func (v NullableGetBackupSchedulesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBackupSchedulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
