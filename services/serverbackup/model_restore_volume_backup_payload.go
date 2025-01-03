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

// checks if the RestoreVolumeBackupPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreVolumeBackupPayload{}

// RestoreVolumeBackupPayload struct for RestoreVolumeBackupPayload
type RestoreVolumeBackupPayload struct {
	// REQUIRED
	RestoreVolumeId *string `json:"restoreVolumeId"`
}

type _RestoreVolumeBackupPayload RestoreVolumeBackupPayload

// NewRestoreVolumeBackupPayload instantiates a new RestoreVolumeBackupPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreVolumeBackupPayload(restoreVolumeId *string) *RestoreVolumeBackupPayload {
	this := RestoreVolumeBackupPayload{}
	this.RestoreVolumeId = restoreVolumeId
	return &this
}

// NewRestoreVolumeBackupPayloadWithDefaults instantiates a new RestoreVolumeBackupPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreVolumeBackupPayloadWithDefaults() *RestoreVolumeBackupPayload {
	this := RestoreVolumeBackupPayload{}
	return &this
}

// GetRestoreVolumeId returns the RestoreVolumeId field value
func (o *RestoreVolumeBackupPayload) GetRestoreVolumeId() *string {
	if o == nil || IsNil(o.RestoreVolumeId) {
		var ret *string
		return ret
	}

	return o.RestoreVolumeId
}

// GetRestoreVolumeIdOk returns a tuple with the RestoreVolumeId field value
// and a boolean to check if the value has been set.
func (o *RestoreVolumeBackupPayload) GetRestoreVolumeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RestoreVolumeId, true
}

// SetRestoreVolumeId sets field value
func (o *RestoreVolumeBackupPayload) SetRestoreVolumeId(v *string) {
	o.RestoreVolumeId = v
}

func (o RestoreVolumeBackupPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["restoreVolumeId"] = o.RestoreVolumeId
	return toSerialize, nil
}

type NullableRestoreVolumeBackupPayload struct {
	value *RestoreVolumeBackupPayload
	isSet bool
}

func (v NullableRestoreVolumeBackupPayload) Get() *RestoreVolumeBackupPayload {
	return v.value
}

func (v *NullableRestoreVolumeBackupPayload) Set(val *RestoreVolumeBackupPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreVolumeBackupPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreVolumeBackupPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreVolumeBackupPayload(val *RestoreVolumeBackupPayload) *NullableRestoreVolumeBackupPayload {
	return &NullableRestoreVolumeBackupPayload{value: val, isSet: true}
}

func (v NullableRestoreVolumeBackupPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreVolumeBackupPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
