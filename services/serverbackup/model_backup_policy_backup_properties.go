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

// checks if the BackupPolicyBackupProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupPolicyBackupProperties{}

// BackupPolicyBackupProperties struct for BackupPolicyBackupProperties
type BackupPolicyBackupProperties struct {
	Name            *string `json:"name,omitempty"`
	RetentionPeriod *int64  `json:"retentionPeriod,omitempty"`
}

// NewBackupPolicyBackupProperties instantiates a new BackupPolicyBackupProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupPolicyBackupProperties() *BackupPolicyBackupProperties {
	this := BackupPolicyBackupProperties{}
	return &this
}

// NewBackupPolicyBackupPropertiesWithDefaults instantiates a new BackupPolicyBackupProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupPolicyBackupPropertiesWithDefaults() *BackupPolicyBackupProperties {
	this := BackupPolicyBackupProperties{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BackupPolicyBackupProperties) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicyBackupProperties) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BackupPolicyBackupProperties) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BackupPolicyBackupProperties) SetName(v *string) {
	o.Name = v
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *BackupPolicyBackupProperties) GetRetentionPeriod() *int64 {
	if o == nil || IsNil(o.RetentionPeriod) {
		var ret *int64
		return ret
	}
	return o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicyBackupProperties) GetRetentionPeriodOk() (*int64, bool) {
	if o == nil || IsNil(o.RetentionPeriod) {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *BackupPolicyBackupProperties) HasRetentionPeriod() bool {
	if o != nil && !IsNil(o.RetentionPeriod) {
		return true
	}

	return false
}

// SetRetentionPeriod gets a reference to the given int64 and assigns it to the RetentionPeriod field.
func (o *BackupPolicyBackupProperties) SetRetentionPeriod(v *int64) {
	o.RetentionPeriod = v
}

func (o BackupPolicyBackupProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RetentionPeriod) {
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	return toSerialize, nil
}

type NullableBackupPolicyBackupProperties struct {
	value *BackupPolicyBackupProperties
	isSet bool
}

func (v NullableBackupPolicyBackupProperties) Get() *BackupPolicyBackupProperties {
	return v.value
}

func (v *NullableBackupPolicyBackupProperties) Set(val *BackupPolicyBackupProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupPolicyBackupProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupPolicyBackupProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupPolicyBackupProperties(val *BackupPolicyBackupProperties) *NullableBackupPolicyBackupProperties {
	return &NullableBackupPolicyBackupProperties{value: val, isSet: true}
}

func (v NullableBackupPolicyBackupProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupPolicyBackupProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}