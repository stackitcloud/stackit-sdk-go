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

// checks if the GetBackupPoliciesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBackupPoliciesResponse{}

// GetBackupPoliciesResponse struct for GetBackupPoliciesResponse
type GetBackupPoliciesResponse struct {
	Items *[]BackupPolicy `json:"items,omitempty"`
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
func (o *GetBackupPoliciesResponse) GetItems() *[]BackupPolicy {
	if o == nil || IsNil(o.Items) {
		var ret *[]BackupPolicy
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBackupPoliciesResponse) GetItemsOk() (*[]BackupPolicy, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetBackupPoliciesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BackupPolicy and assigns it to the Items field.
func (o *GetBackupPoliciesResponse) SetItems(v *[]BackupPolicy) {
	o.Items = v
}

func (o GetBackupPoliciesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
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
