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

// checks if the UpdateBackupSchedulePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBackupSchedulePayload{}

// UpdateBackupSchedulePayload struct for UpdateBackupSchedulePayload
type UpdateBackupSchedulePayload struct {
	BackupProperties *BackupProperties `json:"backupProperties,omitempty"`
	// REQUIRED
	Enabled *bool `json:"enabled"`
	// Max 255 characters
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Rrule *string `json:"rrule"`
}

type _UpdateBackupSchedulePayload UpdateBackupSchedulePayload

// NewUpdateBackupSchedulePayload instantiates a new UpdateBackupSchedulePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBackupSchedulePayload(enabled *bool, name *string, rrule *string) *UpdateBackupSchedulePayload {
	this := UpdateBackupSchedulePayload{}
	this.Enabled = enabled
	this.Name = name
	this.Rrule = rrule
	return &this
}

// NewUpdateBackupSchedulePayloadWithDefaults instantiates a new UpdateBackupSchedulePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBackupSchedulePayloadWithDefaults() *UpdateBackupSchedulePayload {
	this := UpdateBackupSchedulePayload{}
	return &this
}

// GetBackupProperties returns the BackupProperties field value if set, zero value otherwise.
func (o *UpdateBackupSchedulePayload) GetBackupProperties() *BackupProperties {
	if o == nil || IsNil(o.BackupProperties) {
		var ret *BackupProperties
		return ret
	}
	return o.BackupProperties
}

// GetBackupPropertiesOk returns a tuple with the BackupProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBackupSchedulePayload) GetBackupPropertiesOk() (*BackupProperties, bool) {
	if o == nil || IsNil(o.BackupProperties) {
		return nil, false
	}
	return o.BackupProperties, true
}

// HasBackupProperties returns a boolean if a field has been set.
func (o *UpdateBackupSchedulePayload) HasBackupProperties() bool {
	if o != nil && !IsNil(o.BackupProperties) {
		return true
	}

	return false
}

// SetBackupProperties gets a reference to the given BackupProperties and assigns it to the BackupProperties field.
func (o *UpdateBackupSchedulePayload) SetBackupProperties(v *BackupProperties) {
	o.BackupProperties = v
}

// GetEnabled returns the Enabled field value
func (o *UpdateBackupSchedulePayload) GetEnabled() *bool {
	if o == nil || IsNil(o.Enabled) {
		var ret *bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UpdateBackupSchedulePayload) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled, true
}

// SetEnabled sets field value
func (o *UpdateBackupSchedulePayload) SetEnabled(v *bool) {
	o.Enabled = v
}

// GetName returns the Name field value
func (o *UpdateBackupSchedulePayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateBackupSchedulePayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *UpdateBackupSchedulePayload) SetName(v *string) {
	o.Name = v
}

// GetRrule returns the Rrule field value
func (o *UpdateBackupSchedulePayload) GetRrule() *string {
	if o == nil || IsNil(o.Rrule) {
		var ret *string
		return ret
	}

	return o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value
// and a boolean to check if the value has been set.
func (o *UpdateBackupSchedulePayload) GetRruleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rrule, true
}

// SetRrule sets field value
func (o *UpdateBackupSchedulePayload) SetRrule(v *string) {
	o.Rrule = v
}

func (o UpdateBackupSchedulePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackupProperties) {
		toSerialize["backupProperties"] = o.BackupProperties
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["name"] = o.Name
	toSerialize["rrule"] = o.Rrule
	return toSerialize, nil
}

type NullableUpdateBackupSchedulePayload struct {
	value *UpdateBackupSchedulePayload
	isSet bool
}

func (v NullableUpdateBackupSchedulePayload) Get() *UpdateBackupSchedulePayload {
	return v.value
}

func (v *NullableUpdateBackupSchedulePayload) Set(val *UpdateBackupSchedulePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBackupSchedulePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBackupSchedulePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBackupSchedulePayload(val *UpdateBackupSchedulePayload) *NullableUpdateBackupSchedulePayload {
	return &NullableUpdateBackupSchedulePayload{value: val, isSet: true}
}

func (v NullableUpdateBackupSchedulePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBackupSchedulePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
