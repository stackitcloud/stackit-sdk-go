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

// checks if the GetBackupSchedulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBackupSchedulesResponse{}

/*
	types and functions for items
*/

// isArray
type GetBackupSchedulesResponseGetItemsAttributeType = *[]BackupSchedule
type GetBackupSchedulesResponseGetItemsArgType = []BackupSchedule
type GetBackupSchedulesResponseGetItemsRetType = []BackupSchedule

func getGetBackupSchedulesResponseGetItemsAttributeTypeOk(arg GetBackupSchedulesResponseGetItemsAttributeType) (ret GetBackupSchedulesResponseGetItemsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetBackupSchedulesResponseGetItemsAttributeType(arg *GetBackupSchedulesResponseGetItemsAttributeType, val GetBackupSchedulesResponseGetItemsRetType) {
	*arg = &val
}

// GetBackupSchedulesResponse struct for GetBackupSchedulesResponse
type GetBackupSchedulesResponse struct {
	Items GetBackupSchedulesResponseGetItemsAttributeType `json:"items,omitempty"`
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
func (o *GetBackupSchedulesResponse) GetItems() (res GetBackupSchedulesResponseGetItemsRetType) {
	res, _ = o.GetItemsOk()
	return
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBackupSchedulesResponse) GetItemsOk() (ret GetBackupSchedulesResponseGetItemsRetType, ok bool) {
	return getGetBackupSchedulesResponseGetItemsAttributeTypeOk(o.Items)
}

// HasItems returns a boolean if a field has been set.
func (o *GetBackupSchedulesResponse) HasItems() bool {
	_, ok := o.GetItemsOk()
	return ok
}

// SetItems gets a reference to the given []BackupSchedule and assigns it to the Items field.
func (o *GetBackupSchedulesResponse) SetItems(v GetBackupSchedulesResponseGetItemsRetType) {
	setGetBackupSchedulesResponseGetItemsAttributeType(&o.Items, v)
}

func (o GetBackupSchedulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getGetBackupSchedulesResponseGetItemsAttributeTypeOk(o.Items); ok {
		toSerialize["Items"] = val
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
