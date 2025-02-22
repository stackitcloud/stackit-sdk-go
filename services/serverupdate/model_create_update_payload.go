/*
STACKIT Server Update Management API

API endpoints for Server Update Operations on STACKIT Servers.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverupdate

import (
	"encoding/json"
)

// checks if the CreateUpdatePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUpdatePayload{}

// CreateUpdatePayload struct for CreateUpdatePayload
type CreateUpdatePayload struct {
	BackupBeforeUpdate *bool `json:"backupBeforeUpdate,omitempty"`
	// Can be cast to int32 without loss of precision.
	// REQUIRED
	MaintenanceWindow *int64 `json:"maintenanceWindow"`
}

type _CreateUpdatePayload CreateUpdatePayload

// NewCreateUpdatePayload instantiates a new CreateUpdatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUpdatePayload(maintenanceWindow *int64) *CreateUpdatePayload {
	this := CreateUpdatePayload{}
	this.MaintenanceWindow = maintenanceWindow
	return &this
}

// NewCreateUpdatePayloadWithDefaults instantiates a new CreateUpdatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUpdatePayloadWithDefaults() *CreateUpdatePayload {
	this := CreateUpdatePayload{}
	return &this
}

// GetBackupBeforeUpdate returns the BackupBeforeUpdate field value if set, zero value otherwise.
func (o *CreateUpdatePayload) GetBackupBeforeUpdate() *bool {
	if o == nil || IsNil(o.BackupBeforeUpdate) {
		var ret *bool
		return ret
	}
	return o.BackupBeforeUpdate
}

// GetBackupBeforeUpdateOk returns a tuple with the BackupBeforeUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdatePayload) GetBackupBeforeUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.BackupBeforeUpdate) {
		return nil, false
	}
	return o.BackupBeforeUpdate, true
}

// HasBackupBeforeUpdate returns a boolean if a field has been set.
func (o *CreateUpdatePayload) HasBackupBeforeUpdate() bool {
	if o != nil && !IsNil(o.BackupBeforeUpdate) {
		return true
	}

	return false
}

// SetBackupBeforeUpdate gets a reference to the given bool and assigns it to the BackupBeforeUpdate field.
func (o *CreateUpdatePayload) SetBackupBeforeUpdate(v *bool) {
	o.BackupBeforeUpdate = v
}

// GetMaintenanceWindow returns the MaintenanceWindow field value
func (o *CreateUpdatePayload) GetMaintenanceWindow() *int64 {
	if o == nil || IsNil(o.MaintenanceWindow) {
		var ret *int64
		return ret
	}

	return o.MaintenanceWindow
}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value
// and a boolean to check if the value has been set.
func (o *CreateUpdatePayload) GetMaintenanceWindowOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaintenanceWindow, true
}

// SetMaintenanceWindow sets field value
func (o *CreateUpdatePayload) SetMaintenanceWindow(v *int64) {
	o.MaintenanceWindow = v
}

func (o CreateUpdatePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackupBeforeUpdate) {
		toSerialize["backupBeforeUpdate"] = o.BackupBeforeUpdate
	}
	toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	return toSerialize, nil
}

type NullableCreateUpdatePayload struct {
	value *CreateUpdatePayload
	isSet bool
}

func (v NullableCreateUpdatePayload) Get() *CreateUpdatePayload {
	return v.value
}

func (v *NullableCreateUpdatePayload) Set(val *CreateUpdatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUpdatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUpdatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUpdatePayload(val *CreateUpdatePayload) *NullableCreateUpdatePayload {
	return &NullableCreateUpdatePayload{value: val, isSet: true}
}

func (v NullableCreateUpdatePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUpdatePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
