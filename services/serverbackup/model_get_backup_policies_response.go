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

// checks if the GetBackupPoliciesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBackupPoliciesResponse{}

/*
	types and functions for items
*/

// isArray
type GetBackupPoliciesResponseGetItemsAttributeType = *[]BackupPolicy
type GetBackupPoliciesResponseGetItemsArgType = []BackupPolicy
type GetBackupPoliciesResponseGetItemsRetType = []BackupPolicy

func getGetBackupPoliciesResponseGetItemsAttributeTypeOk(arg GetBackupPoliciesResponseGetItemsAttributeType) (ret GetBackupPoliciesResponseGetItemsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetBackupPoliciesResponseGetItemsAttributeType(arg *GetBackupPoliciesResponseGetItemsAttributeType, val GetBackupPoliciesResponseGetItemsRetType) {
	*arg = &val
}

// GetBackupPoliciesResponse struct for GetBackupPoliciesResponse
type GetBackupPoliciesResponse struct {
	Items GetBackupPoliciesResponseGetItemsAttributeType `json:"items,omitempty"`
}

// NewGetBackupPoliciesResponse instantiates a new GetBackupPoliciesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBackupPoliciesResponse() *GetBackupPoliciesResponse {
	this := GetBackupPoliciesResponse{}
	return &this
}

// NewGetBackupPoliciesResponseWithDefaults instantiates a new GetBackupPoliciesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBackupPoliciesResponseWithDefaults() *GetBackupPoliciesResponse {
	this := GetBackupPoliciesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetBackupPoliciesResponse) GetItems() (res GetBackupPoliciesResponseGetItemsRetType) {
	res, _ = o.GetItemsOk()
	return
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBackupPoliciesResponse) GetItemsOk() (ret GetBackupPoliciesResponseGetItemsRetType, ok bool) {
	return getGetBackupPoliciesResponseGetItemsAttributeTypeOk(o.Items)
}

// HasItems returns a boolean if a field has been set.
func (o *GetBackupPoliciesResponse) HasItems() bool {
	_, ok := o.GetItemsOk()
	return ok
}

// SetItems gets a reference to the given []BackupPolicy and assigns it to the Items field.
func (o *GetBackupPoliciesResponse) SetItems(v GetBackupPoliciesResponseGetItemsRetType) {
	setGetBackupPoliciesResponseGetItemsAttributeType(&o.Items, v)
}

func (o GetBackupPoliciesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getGetBackupPoliciesResponseGetItemsAttributeTypeOk(o.Items); ok {
		toSerialize["Items"] = val
	}
	return toSerialize, nil
}

type NullableGetBackupPoliciesResponse struct {
	value *GetBackupPoliciesResponse
	isSet bool
}

func (v NullableGetBackupPoliciesResponse) Get() *GetBackupPoliciesResponse {
	return v.value
}

func (v *NullableGetBackupPoliciesResponse) Set(val *GetBackupPoliciesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBackupPoliciesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBackupPoliciesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBackupPoliciesResponse(val *GetBackupPoliciesResponse) *NullableGetBackupPoliciesResponse {
	return &NullableGetBackupPoliciesResponse{value: val, isSet: true}
}

func (v NullableGetBackupPoliciesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBackupPoliciesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
