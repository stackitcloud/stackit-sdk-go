/*
STACKIT LogMe API

The STACKIT LogMe API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logme

import (
	"encoding/json"
)

// checks if the Backup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Backup{}

/*
	types and functions for downloadable
*/

// isBoolean
type BackupgetDownloadableAttributeType = *bool
type BackupgetDownloadableArgType = bool
type BackupgetDownloadableRetType = bool

func getBackupgetDownloadableAttributeTypeOk(arg BackupgetDownloadableAttributeType) (ret BackupgetDownloadableRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupgetDownloadableAttributeType(arg *BackupgetDownloadableAttributeType, val BackupgetDownloadableRetType) {
	*arg = &val
}

/*
	types and functions for finished_at
*/

// isNotNullableString
type BackupGetFinishedAtAttributeType = *string

func getBackupGetFinishedAtAttributeTypeOk(arg BackupGetFinishedAtAttributeType) (ret BackupGetFinishedAtRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetFinishedAtAttributeType(arg *BackupGetFinishedAtAttributeType, val BackupGetFinishedAtRetType) {
	*arg = &val
}

type BackupGetFinishedAtArgType = string
type BackupGetFinishedAtRetType = string

/*
	types and functions for id
*/

// isInteger
type BackupGetIdAttributeType = *int64
type BackupGetIdArgType = int64
type BackupGetIdRetType = int64

func getBackupGetIdAttributeTypeOk(arg BackupGetIdAttributeType) (ret BackupGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetIdAttributeType(arg *BackupGetIdAttributeType, val BackupGetIdRetType) {
	*arg = &val
}

/*
	types and functions for size
*/

// isInteger
type BackupGetSizeAttributeType = *int64
type BackupGetSizeArgType = int64
type BackupGetSizeRetType = int64

func getBackupGetSizeAttributeTypeOk(arg BackupGetSizeAttributeType) (ret BackupGetSizeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetSizeAttributeType(arg *BackupGetSizeAttributeType, val BackupGetSizeRetType) {
	*arg = &val
}

/*
	types and functions for status
*/

// isNotNullableString
type BackupGetStatusAttributeType = *string

func getBackupGetStatusAttributeTypeOk(arg BackupGetStatusAttributeType) (ret BackupGetStatusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetStatusAttributeType(arg *BackupGetStatusAttributeType, val BackupGetStatusRetType) {
	*arg = &val
}

type BackupGetStatusArgType = string
type BackupGetStatusRetType = string

/*
	types and functions for triggered_at
*/

// isNotNullableString
type BackupGetTriggeredAtAttributeType = *string

func getBackupGetTriggeredAtAttributeTypeOk(arg BackupGetTriggeredAtAttributeType) (ret BackupGetTriggeredAtRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupGetTriggeredAtAttributeType(arg *BackupGetTriggeredAtAttributeType, val BackupGetTriggeredAtRetType) {
	*arg = &val
}

type BackupGetTriggeredAtArgType = string
type BackupGetTriggeredAtRetType = string

// Backup struct for Backup
type Backup struct {
	Downloadable BackupgetDownloadableAttributeType `json:"downloadable,omitempty"`
	// REQUIRED
	FinishedAt BackupGetFinishedAtAttributeType `json:"finished_at"`
	// REQUIRED
	Id   BackupGetIdAttributeType   `json:"id"`
	Size BackupGetSizeAttributeType `json:"size,omitempty"`
	// REQUIRED
	Status      BackupGetStatusAttributeType      `json:"status"`
	TriggeredAt BackupGetTriggeredAtAttributeType `json:"triggered_at,omitempty"`
}

type _Backup Backup

// NewBackup instantiates a new Backup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackup(finishedAt BackupGetFinishedAtArgType, id BackupGetIdArgType, status BackupGetStatusArgType) *Backup {
	this := Backup{}
	setBackupGetFinishedAtAttributeType(&this.FinishedAt, finishedAt)
	setBackupGetIdAttributeType(&this.Id, id)
	setBackupGetStatusAttributeType(&this.Status, status)
	return &this
}

// NewBackupWithDefaults instantiates a new Backup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupWithDefaults() *Backup {
	this := Backup{}
	return &this
}

// GetDownloadable returns the Downloadable field value if set, zero value otherwise.
func (o *Backup) GetDownloadable() (res BackupgetDownloadableRetType) {
	res, _ = o.GetDownloadableOk()
	return
}

// GetDownloadableOk returns a tuple with the Downloadable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetDownloadableOk() (ret BackupgetDownloadableRetType, ok bool) {
	return getBackupgetDownloadableAttributeTypeOk(o.Downloadable)
}

// HasDownloadable returns a boolean if a field has been set.
func (o *Backup) HasDownloadable() bool {
	_, ok := o.GetDownloadableOk()
	return ok
}

// SetDownloadable gets a reference to the given bool and assigns it to the Downloadable field.
func (o *Backup) SetDownloadable(v BackupgetDownloadableRetType) {
	setBackupgetDownloadableAttributeType(&o.Downloadable, v)
}

// GetFinishedAt returns the FinishedAt field value
func (o *Backup) GetFinishedAt() (ret BackupGetFinishedAtRetType) {
	ret, _ = o.GetFinishedAtOk()
	return ret
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value
// and a boolean to check if the value has been set.
func (o *Backup) GetFinishedAtOk() (ret BackupGetFinishedAtRetType, ok bool) {
	return getBackupGetFinishedAtAttributeTypeOk(o.FinishedAt)
}

// SetFinishedAt sets field value
func (o *Backup) SetFinishedAt(v BackupGetFinishedAtRetType) {
	setBackupGetFinishedAtAttributeType(&o.FinishedAt, v)
}

// GetId returns the Id field value
func (o *Backup) GetId() (ret BackupGetIdRetType) {
	ret, _ = o.GetIdOk()
	return ret
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Backup) GetIdOk() (ret BackupGetIdRetType, ok bool) {
	return getBackupGetIdAttributeTypeOk(o.Id)
}

// SetId sets field value
func (o *Backup) SetId(v BackupGetIdRetType) {
	setBackupGetIdAttributeType(&o.Id, v)
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Backup) GetSize() (res BackupGetSizeRetType) {
	res, _ = o.GetSizeOk()
	return
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetSizeOk() (ret BackupGetSizeRetType, ok bool) {
	return getBackupGetSizeAttributeTypeOk(o.Size)
}

// HasSize returns a boolean if a field has been set.
func (o *Backup) HasSize() bool {
	_, ok := o.GetSizeOk()
	return ok
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *Backup) SetSize(v BackupGetSizeRetType) {
	setBackupGetSizeAttributeType(&o.Size, v)
}

// GetStatus returns the Status field value
func (o *Backup) GetStatus() (ret BackupGetStatusRetType) {
	ret, _ = o.GetStatusOk()
	return ret
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Backup) GetStatusOk() (ret BackupGetStatusRetType, ok bool) {
	return getBackupGetStatusAttributeTypeOk(o.Status)
}

// SetStatus sets field value
func (o *Backup) SetStatus(v BackupGetStatusRetType) {
	setBackupGetStatusAttributeType(&o.Status, v)
}

// GetTriggeredAt returns the TriggeredAt field value if set, zero value otherwise.
func (o *Backup) GetTriggeredAt() (res BackupGetTriggeredAtRetType) {
	res, _ = o.GetTriggeredAtOk()
	return
}

// GetTriggeredAtOk returns a tuple with the TriggeredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Backup) GetTriggeredAtOk() (ret BackupGetTriggeredAtRetType, ok bool) {
	return getBackupGetTriggeredAtAttributeTypeOk(o.TriggeredAt)
}

// HasTriggeredAt returns a boolean if a field has been set.
func (o *Backup) HasTriggeredAt() bool {
	_, ok := o.GetTriggeredAtOk()
	return ok
}

// SetTriggeredAt gets a reference to the given string and assigns it to the TriggeredAt field.
func (o *Backup) SetTriggeredAt(v BackupGetTriggeredAtRetType) {
	setBackupGetTriggeredAtAttributeType(&o.TriggeredAt, v)
}

func (o Backup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getBackupgetDownloadableAttributeTypeOk(o.Downloadable); ok {
		toSerialize["Downloadable"] = val
	}
	if val, ok := getBackupGetFinishedAtAttributeTypeOk(o.FinishedAt); ok {
		toSerialize["FinishedAt"] = val
	}
	if val, ok := getBackupGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getBackupGetSizeAttributeTypeOk(o.Size); ok {
		toSerialize["Size"] = val
	}
	if val, ok := getBackupGetStatusAttributeTypeOk(o.Status); ok {
		toSerialize["Status"] = val
	}
	if val, ok := getBackupGetTriggeredAtAttributeTypeOk(o.TriggeredAt); ok {
		toSerialize["TriggeredAt"] = val
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
