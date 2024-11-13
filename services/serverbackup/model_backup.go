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

// checks if the Backup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Backup{}

// Backup struct for Backup
type Backup struct {
	// REQUIRED
	CreatedAt *string `json:"createdAt"`
	// REQUIRED
	ExpireAt *string `json:"expireAt"`
	// REQUIRED
	Id             *string `json:"id"`
	LastRestoredAt *string `json:"lastRestoredAt,omitempty"`
	// REQUIRED
	Name *string `json:"name"`
	Size *int64  `json:"size,omitempty"`
	// REQUIRED
	Status        *string                     `json:"status"`
	VolumeBackups *[]BackupVolumeBackupsInner `json:"volumeBackups,omitempty"`
}

type _Backup Backup

// NewBackup instantiates a new Backup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackup(createdAt *string, expireAt *string, id *string, name *string, status *string) *Backup {
	this := Backup{}
	this.CreatedAt = createdAt
	this.ExpireAt = expireAt
	this.Id = id
	this.Name = name
	this.Status = status
	return &this
}

// NewBackupWithDefaults instantiates a new Backup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupWithDefaults() *Backup {
	this := Backup{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *Backup) GetCreatedAt() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Backup) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Backup) SetCreatedAt(v *string) {
	o.CreatedAt = v
}

// GetExpireAt returns the ExpireAt field value
func (o *Backup) GetExpireAt() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.ExpireAt
}

// GetExpireAtOk returns a tuple with the ExpireAt field value
// and a boolean to check if the value has been set.
func (o *Backup) GetExpireAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpireAt, true
}

// SetExpireAt sets field value
func (o *Backup) SetExpireAt(v *string) {
	o.ExpireAt = v
}

// GetId returns the Id field value
func (o *Backup) GetId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Backup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *Backup) SetId(v *string) {
	o.Id = v
}

// GetLastRestoredAt returns the LastRestoredAt field value if set, zero value otherwise.
func (o *Backup) GetLastRestoredAt() *string {
	if o == nil || IsNil(o.LastRestoredAt) {
		var ret *string
		return ret
	}
	return o.LastRestoredAt
}

// GetLastRestoredAtOk returns a tuple with the LastRestoredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetLastRestoredAtOk() (*string, bool) {
	if o == nil || IsNil(o.LastRestoredAt) {
		return nil, false
	}
	return o.LastRestoredAt, true
}

// HasLastRestoredAt returns a boolean if a field has been set.
func (o *Backup) HasLastRestoredAt() bool {
	if o != nil && !IsNil(o.LastRestoredAt) {
		return true
	}

	return false
}

// SetLastRestoredAt gets a reference to the given string and assigns it to the LastRestoredAt field.
func (o *Backup) SetLastRestoredAt(v *string) {
	o.LastRestoredAt = v
}

// GetName returns the Name field value
func (o *Backup) GetName() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Backup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *Backup) SetName(v *string) {
	o.Name = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Backup) GetSize() *int64 {
	if o == nil || IsNil(o.Size) {
		var ret *int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Backup) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *Backup) SetSize(v *int64) {
	o.Size = v
}

// GetStatus returns the Status field value
func (o *Backup) GetStatus() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Backup) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *Backup) SetStatus(v *string) {
	o.Status = v
}

// GetVolumeBackups returns the VolumeBackups field value if set, zero value otherwise.
func (o *Backup) GetVolumeBackups() *[]BackupVolumeBackupsInner {
	if o == nil || IsNil(o.VolumeBackups) {
		var ret *[]BackupVolumeBackupsInner
		return ret
	}
	return o.VolumeBackups
}

// GetVolumeBackupsOk returns a tuple with the VolumeBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetVolumeBackupsOk() (*[]BackupVolumeBackupsInner, bool) {
	if o == nil || IsNil(o.VolumeBackups) {
		return nil, false
	}
	return o.VolumeBackups, true
}

// HasVolumeBackups returns a boolean if a field has been set.
func (o *Backup) HasVolumeBackups() bool {
	if o != nil && !IsNil(o.VolumeBackups) {
		return true
	}

	return false
}

// SetVolumeBackups gets a reference to the given []BackupVolumeBackupsInner and assigns it to the VolumeBackups field.
func (o *Backup) SetVolumeBackups(v *[]BackupVolumeBackupsInner) {
	o.VolumeBackups = v
}

func (o Backup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["expireAt"] = o.ExpireAt
	toSerialize["id"] = o.Id
	if !IsNil(o.LastRestoredAt) {
		toSerialize["lastRestoredAt"] = o.LastRestoredAt
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.VolumeBackups) {
		toSerialize["volumeBackups"] = o.VolumeBackups
	}
	return toSerialize, nil
}

type NullableBackup struct {
	value *Backup
	isSet bool
}

func (v NullableBackup) Get() *Backup {
	return v.value
}

func (v *NullableBackup) Set(val *Backup) {
	v.value = val
	v.isSet = true
}

func (v NullableBackup) IsSet() bool {
	return v.isSet
}

func (v *NullableBackup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackup(val *Backup) *NullableBackup {
	return &NullableBackup{value: val, isSet: true}
}

func (v NullableBackup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
