/*
STACKIT Redis API

The STACKIT Redis API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redis

import (
	"encoding/json"
)

// checks if the Restore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Restore{}

// Restore struct for Restore
type Restore struct {
	// REQUIRED
	BackupId *int64 `json:"backup_id"`
	// REQUIRED
	FinishedAt *string `json:"finished_at"`
	// REQUIRED
	Id *int64 `json:"id"`
	// REQUIRED
	Status      *string `json:"status"`
	TriggeredAt *string `json:"triggered_at,omitempty"`
}

type _Restore Restore

// NewRestore instantiates a new Restore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestore(backupId *int64, finishedAt *string, id *int64, status *string) *Restore {
	this := Restore{}
	this.BackupId = backupId
	this.FinishedAt = finishedAt
	this.Id = id
	this.Status = status
	return &this
}

// NewRestoreWithDefaults instantiates a new Restore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreWithDefaults() *Restore {
	this := Restore{}
	return &this
}

// GetBackupId returns the BackupId field value
func (o *Restore) GetBackupId() *int64 {
	if o == nil || IsNil(o.BackupId) {
		var ret *int64
		return ret
	}

	return o.BackupId
}

// GetBackupIdOk returns a tuple with the BackupId field value
// and a boolean to check if the value has been set.
func (o *Restore) GetBackupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupId, true
}

// SetBackupId sets field value
func (o *Restore) SetBackupId(v *int64) {
	o.BackupId = v
}

// GetFinishedAt returns the FinishedAt field value
func (o *Restore) GetFinishedAt() *string {
	if o == nil || IsNil(o.FinishedAt) {
		var ret *string
		return ret
	}

	return o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value
// and a boolean to check if the value has been set.
func (o *Restore) GetFinishedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// SetFinishedAt sets field value
func (o *Restore) SetFinishedAt(v *string) {
	o.FinishedAt = v
}

// GetId returns the Id field value
func (o *Restore) GetId() *int64 {
	if o == nil || IsNil(o.Id) {
		var ret *int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Restore) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *Restore) SetId(v *int64) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *Restore) GetStatus() *string {
	if o == nil || IsNil(o.Status) {
		var ret *string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Restore) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *Restore) SetStatus(v *string) {
	o.Status = v
}

// GetTriggeredAt returns the TriggeredAt field value if set, zero value otherwise.
func (o *Restore) GetTriggeredAt() *string {
	if o == nil || IsNil(o.TriggeredAt) {
		var ret *string
		return ret
	}
	return o.TriggeredAt
}

// GetTriggeredAtOk returns a tuple with the TriggeredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Restore) GetTriggeredAtOk() (*string, bool) {
	if o == nil || IsNil(o.TriggeredAt) {
		return nil, false
	}
	return o.TriggeredAt, true
}

// HasTriggeredAt returns a boolean if a field has been set.
func (o *Restore) HasTriggeredAt() bool {
	if o != nil && !IsNil(o.TriggeredAt) {
		return true
	}

	return false
}

// SetTriggeredAt gets a reference to the given string and assigns it to the TriggeredAt field.
func (o *Restore) SetTriggeredAt(v *string) {
	o.TriggeredAt = v
}

func (o Restore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["backup_id"] = o.BackupId
	toSerialize["finished_at"] = o.FinishedAt
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	if !IsNil(o.TriggeredAt) {
		toSerialize["triggered_at"] = o.TriggeredAt
	}
	return toSerialize, nil
}

type NullableRestore struct {
	value *Restore
	isSet bool
}

func (v NullableRestore) Get() *Restore {
	return v.value
}

func (v *NullableRestore) Set(val *Restore) {
	v.value = val
	v.isSet = true
}

func (v NullableRestore) IsSet() bool {
	return v.isSet
}

func (v *NullableRestore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestore(val *Restore) *NullableRestore {
	return &NullableRestore{value: val, isSet: true}
}

func (v NullableRestore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
