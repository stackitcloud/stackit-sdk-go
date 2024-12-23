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

// checks if the RestoreBackupPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreBackupPayload{}

// RestoreBackupPayload struct for RestoreBackupPayload
type RestoreBackupPayload struct {
	// REQUIRED
	StartServerAfterRestore *bool     `json:"startServerAfterRestore"`
	VolumeIds               *[]string `json:"volumeIds,omitempty"`
}

type _RestoreBackupPayload RestoreBackupPayload

// NewRestoreBackupPayload instantiates a new RestoreBackupPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreBackupPayload(startServerAfterRestore *bool) *RestoreBackupPayload {
	this := RestoreBackupPayload{}
	this.StartServerAfterRestore = startServerAfterRestore
	return &this
}

// NewRestoreBackupPayloadWithDefaults instantiates a new RestoreBackupPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreBackupPayloadWithDefaults() *RestoreBackupPayload {
	this := RestoreBackupPayload{}
	return &this
}

// GetStartServerAfterRestore returns the StartServerAfterRestore field value
func (o *RestoreBackupPayload) GetStartServerAfterRestore() *bool {
	if o == nil || IsNil(o.StartServerAfterRestore) {
		var ret *bool
		return ret
	}

	return o.StartServerAfterRestore
}

// GetStartServerAfterRestoreOk returns a tuple with the StartServerAfterRestore field value
// and a boolean to check if the value has been set.
func (o *RestoreBackupPayload) GetStartServerAfterRestoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartServerAfterRestore, true
}

// SetStartServerAfterRestore sets field value
func (o *RestoreBackupPayload) SetStartServerAfterRestore(v *bool) {
	o.StartServerAfterRestore = v
}

// GetVolumeIds returns the VolumeIds field value if set, zero value otherwise.
func (o *RestoreBackupPayload) GetVolumeIds() *[]string {
	if o == nil || IsNil(o.VolumeIds) {
		var ret *[]string
		return ret
	}
	return o.VolumeIds
}

// GetVolumeIdsOk returns a tuple with the VolumeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreBackupPayload) GetVolumeIdsOk() (*[]string, bool) {
	if o == nil || IsNil(o.VolumeIds) {
		return nil, false
	}
	return o.VolumeIds, true
}

// HasVolumeIds returns a boolean if a field has been set.
func (o *RestoreBackupPayload) HasVolumeIds() bool {
	if o != nil && !IsNil(o.VolumeIds) {
		return true
	}

	return false
}

// SetVolumeIds gets a reference to the given []string and assigns it to the VolumeIds field.
func (o *RestoreBackupPayload) SetVolumeIds(v *[]string) {
	o.VolumeIds = v
}

func (o RestoreBackupPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startServerAfterRestore"] = o.StartServerAfterRestore
	if !IsNil(o.VolumeIds) {
		toSerialize["volumeIds"] = o.VolumeIds
	}
	return toSerialize, nil
}

type NullableRestoreBackupPayload struct {
	value *RestoreBackupPayload
	isSet bool
}

func (v NullableRestoreBackupPayload) Get() *RestoreBackupPayload {
	return v.value
}

func (v *NullableRestoreBackupPayload) Set(val *RestoreBackupPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreBackupPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreBackupPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreBackupPayload(val *RestoreBackupPayload) *NullableRestoreBackupPayload {
	return &NullableRestoreBackupPayload{value: val, isSet: true}
}

func (v NullableRestoreBackupPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreBackupPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
