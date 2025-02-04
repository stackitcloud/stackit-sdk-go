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

// checks if the BackupProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupProperties{}

// BackupProperties struct for BackupProperties
type BackupProperties struct {
	// Max 255 characters
	// REQUIRED
	Name *string `json:"name"`
	// Values are set in days (1-36500)
	// REQUIRED
	RetentionPeriod *int64    `json:"retentionPeriod"`
	VolumeIds       *[]string `json:"volumeIds,omitempty"`
}

type _BackupProperties BackupProperties

// NewBackupProperties instantiates a new BackupProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupProperties(name *string, retentionPeriod *int64) *BackupProperties {
	this := BackupProperties{}
	this.Name = name
	this.RetentionPeriod = retentionPeriod
	return &this
}

// NewBackupPropertiesWithDefaults instantiates a new BackupProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupPropertiesWithDefaults() *BackupProperties {
	this := BackupProperties{}
	return &this
}

// GetName returns the Name field value
func (o *BackupProperties) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BackupProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *BackupProperties) SetName(v *string) {
	o.Name = v
}

// GetRetentionPeriod returns the RetentionPeriod field value
func (o *BackupProperties) GetRetentionPeriod() *int64 {
	if o == nil || IsNil(o.RetentionPeriod) {
		var ret *int64
		return ret
	}

	return o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value
// and a boolean to check if the value has been set.
func (o *BackupProperties) GetRetentionPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// SetRetentionPeriod sets field value
func (o *BackupProperties) SetRetentionPeriod(v *int64) {
	o.RetentionPeriod = v
}

// GetVolumeIds returns the VolumeIds field value if set, zero value otherwise.
func (o *BackupProperties) GetVolumeIds() *[]string {
	if o == nil || IsNil(o.VolumeIds) {
		var ret *[]string
		return ret
	}
	return o.VolumeIds
}

// GetVolumeIdsOk returns a tuple with the VolumeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupProperties) GetVolumeIdsOk() (*[]string, bool) {
	if o == nil || IsNil(o.VolumeIds) {
		return nil, false
	}
	return o.VolumeIds, true
}

// HasVolumeIds returns a boolean if a field has been set.
func (o *BackupProperties) HasVolumeIds() bool {
	if o != nil && !IsNil(o.VolumeIds) {
		return true
	}

	return false
}

// SetVolumeIds gets a reference to the given []string and assigns it to the VolumeIds field.
func (o *BackupProperties) SetVolumeIds(v *[]string) {
	o.VolumeIds = v
}

func (o BackupProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["retentionPeriod"] = o.RetentionPeriod
	if !IsNil(o.VolumeIds) {
		toSerialize["volumeIds"] = o.VolumeIds
	}
	return toSerialize, nil
}

type NullableBackupProperties struct {
	value *BackupProperties
	isSet bool
}

func (v NullableBackupProperties) Get() *BackupProperties {
	return v.value
}

func (v *NullableBackupProperties) Set(val *BackupProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupProperties(val *BackupProperties) *NullableBackupProperties {
	return &NullableBackupProperties{value: val, isSet: true}
}

func (v NullableBackupProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
