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

// checks if the CreateBackupSchedulePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBackupSchedulePayload{}

// CreateBackupSchedulePayload struct for CreateBackupSchedulePayload
type CreateBackupSchedulePayload struct {
	BackupProperties *BackupProperties `json:"backupProperties,omitempty"`
	// REQUIRED
	Enabled *bool `json:"enabled"`
	// Max 255 characters
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Rrule *string `json:"rrule"`
}

type _CreateBackupSchedulePayload CreateBackupSchedulePayload

// NewCreateBackupSchedulePayload instantiates a new CreateBackupSchedulePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBackupSchedulePayload(enabled *bool, name *string, rrule *string) *CreateBackupSchedulePayload {
	this := CreateBackupSchedulePayload{}
	this.Enabled = enabled
	this.Name = name
	this.Rrule = rrule
	return &this
}

// NewCreateBackupSchedulePayloadWithDefaults instantiates a new CreateBackupSchedulePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBackupSchedulePayloadWithDefaults() *CreateBackupSchedulePayload {
	this := CreateBackupSchedulePayload{}
	return &this
}

// GetBackupProperties returns the BackupProperties field value if set, zero value otherwise.
func (o *CreateBackupSchedulePayload) GetBackupProperties() *BackupProperties {
	if o == nil || IsNil(o.BackupProperties) {
		var ret *BackupProperties
		return ret
	}
	return o.BackupProperties
}

// GetBackupPropertiesOk returns a tuple with the BackupProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBackupSchedulePayload) GetBackupPropertiesOk() (*BackupProperties, bool) {
	if o == nil || IsNil(o.BackupProperties) {
		return nil, false
	}
	return o.BackupProperties, true
}

// HasBackupProperties returns a boolean if a field has been set.
func (o *CreateBackupSchedulePayload) HasBackupProperties() bool {
	if o != nil && !IsNil(o.BackupProperties) {
		return true
	}

	return false
}

// SetBackupProperties gets a reference to the given BackupProperties and assigns it to the BackupProperties field.
func (o *CreateBackupSchedulePayload) SetBackupProperties(v *BackupProperties) {
	o.BackupProperties = v
}

// GetEnabled returns the Enabled field value
func (o *CreateBackupSchedulePayload) GetEnabled() *bool {
	if o == nil {
		var ret *bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateBackupSchedulePayload) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateBackupSchedulePayload) SetEnabled(v *bool) {
	o.Enabled = v
}

// GetName returns the Name field value
func (o *CreateBackupSchedulePayload) GetName() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateBackupSchedulePayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *CreateBackupSchedulePayload) SetName(v *string) {
	o.Name = v
}

// GetRrule returns the Rrule field value
func (o *CreateBackupSchedulePayload) GetRrule() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value
// and a boolean to check if the value has been set.
func (o *CreateBackupSchedulePayload) GetRruleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rrule, true
}

// SetRrule sets field value
func (o *CreateBackupSchedulePayload) SetRrule(v *string) {
	o.Rrule = v
}

func (o CreateBackupSchedulePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackupProperties) {
		toSerialize["backupProperties"] = o.BackupProperties
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["name"] = o.Name
	toSerialize["rrule"] = o.Rrule
	return toSerialize, nil
}

type NullableCreateBackupSchedulePayload struct {
	value *CreateBackupSchedulePayload
	isSet bool
}

func (v NullableCreateBackupSchedulePayload) Get() *CreateBackupSchedulePayload {
	return v.value
}

func (v *NullableCreateBackupSchedulePayload) Set(val *CreateBackupSchedulePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBackupSchedulePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBackupSchedulePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBackupSchedulePayload(val *CreateBackupSchedulePayload) *NullableCreateBackupSchedulePayload {
	return &NullableCreateBackupSchedulePayload{value: val, isSet: true}
}

func (v NullableCreateBackupSchedulePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBackupSchedulePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
