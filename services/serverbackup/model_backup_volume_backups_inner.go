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

// checks if the BackupVolumeBackupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupVolumeBackupsInner{}

// BackupVolumeBackupsInner struct for BackupVolumeBackupsInner
type BackupVolumeBackupsInner struct {
	Id                   *string `json:"id,omitempty"`
	LastRestoredAt       *string `json:"lastRestoredAt,omitempty"`
	LastRestoredVolumeId *string `json:"lastRestoredVolumeId,omitempty"`
	// Can be cast to int32 without loss of precision.
	Size     *int64  `json:"size,omitempty"`
	Status   *string `json:"status,omitempty"`
	VolumeId *string `json:"volumeId,omitempty"`
}

// NewBackupVolumeBackupsInner instantiates a new BackupVolumeBackupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupVolumeBackupsInner() *BackupVolumeBackupsInner {
	this := BackupVolumeBackupsInner{}
	return &this
}

// NewBackupVolumeBackupsInnerWithDefaults instantiates a new BackupVolumeBackupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupVolumeBackupsInnerWithDefaults() *BackupVolumeBackupsInner {
	this := BackupVolumeBackupsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BackupVolumeBackupsInner) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupVolumeBackupsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BackupVolumeBackupsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BackupVolumeBackupsInner) SetId(v *string) {
	o.Id = v
}

// GetLastRestoredAt returns the LastRestoredAt field value if set, zero value otherwise.
func (o *BackupVolumeBackupsInner) GetLastRestoredAt() *string {
	if o == nil || IsNil(o.LastRestoredAt) {
		var ret *string
		return ret
	}
	return o.LastRestoredAt
}

// GetLastRestoredAtOk returns a tuple with the LastRestoredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupVolumeBackupsInner) GetLastRestoredAtOk() (*string, bool) {
	if o == nil || IsNil(o.LastRestoredAt) {
		return nil, false
	}
	return o.LastRestoredAt, true
}

// HasLastRestoredAt returns a boolean if a field has been set.
func (o *BackupVolumeBackupsInner) HasLastRestoredAt() bool {
	if o != nil && !IsNil(o.LastRestoredAt) {
		return true
	}

	return false
}

// SetLastRestoredAt gets a reference to the given string and assigns it to the LastRestoredAt field.
func (o *BackupVolumeBackupsInner) SetLastRestoredAt(v *string) {
	o.LastRestoredAt = v
}

// GetLastRestoredVolumeId returns the LastRestoredVolumeId field value if set, zero value otherwise.
func (o *BackupVolumeBackupsInner) GetLastRestoredVolumeId() *string {
	if o == nil || IsNil(o.LastRestoredVolumeId) {
		var ret *string
		return ret
	}
	return o.LastRestoredVolumeId
}

// GetLastRestoredVolumeIdOk returns a tuple with the LastRestoredVolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupVolumeBackupsInner) GetLastRestoredVolumeIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastRestoredVolumeId) {
		return nil, false
	}
	return o.LastRestoredVolumeId, true
}

// HasLastRestoredVolumeId returns a boolean if a field has been set.
func (o *BackupVolumeBackupsInner) HasLastRestoredVolumeId() bool {
	if o != nil && !IsNil(o.LastRestoredVolumeId) {
		return true
	}

	return false
}

// SetLastRestoredVolumeId gets a reference to the given string and assigns it to the LastRestoredVolumeId field.
func (o *BackupVolumeBackupsInner) SetLastRestoredVolumeId(v *string) {
	o.LastRestoredVolumeId = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BackupVolumeBackupsInner) GetSize() *int64 {
	if o == nil || IsNil(o.Size) {
		var ret *int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupVolumeBackupsInner) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BackupVolumeBackupsInner) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *BackupVolumeBackupsInner) SetSize(v *int64) {
	o.Size = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BackupVolumeBackupsInner) GetStatus() *string {
	if o == nil || IsNil(o.Status) {
		var ret *string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupVolumeBackupsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BackupVolumeBackupsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BackupVolumeBackupsInner) SetStatus(v *string) {
	o.Status = v
}

// GetVolumeId returns the VolumeId field value if set, zero value otherwise.
func (o *BackupVolumeBackupsInner) GetVolumeId() *string {
	if o == nil || IsNil(o.VolumeId) {
		var ret *string
		return ret
	}
	return o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupVolumeBackupsInner) GetVolumeIdOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeId) {
		return nil, false
	}
	return o.VolumeId, true
}

// HasVolumeId returns a boolean if a field has been set.
func (o *BackupVolumeBackupsInner) HasVolumeId() bool {
	if o != nil && !IsNil(o.VolumeId) {
		return true
	}

	return false
}

// SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.
func (o *BackupVolumeBackupsInner) SetVolumeId(v *string) {
	o.VolumeId = v
}

func (o BackupVolumeBackupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastRestoredAt) {
		toSerialize["lastRestoredAt"] = o.LastRestoredAt
	}
	if !IsNil(o.LastRestoredVolumeId) {
		toSerialize["lastRestoredVolumeId"] = o.LastRestoredVolumeId
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.VolumeId) {
		toSerialize["volumeId"] = o.VolumeId
	}
	return toSerialize, nil
}

type NullableBackupVolumeBackupsInner struct {
	value *BackupVolumeBackupsInner
	isSet bool
}

func (v NullableBackupVolumeBackupsInner) Get() *BackupVolumeBackupsInner {
	return v.value
}

func (v *NullableBackupVolumeBackupsInner) Set(val *BackupVolumeBackupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupVolumeBackupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupVolumeBackupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupVolumeBackupsInner(val *BackupVolumeBackupsInner) *NullableBackupVolumeBackupsInner {
	return &NullableBackupVolumeBackupsInner{value: val, isSet: true}
}

func (v NullableBackupVolumeBackupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupVolumeBackupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
