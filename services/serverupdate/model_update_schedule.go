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

// checks if the UpdateSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSchedule{}

// UpdateSchedule struct for UpdateSchedule
type UpdateSchedule struct {
	// REQUIRED
	Enabled *bool `json:"enabled"`
	// Can be cast to int32 without loss of precision.
	// REQUIRED
	MaintenanceWindow *int64 `json:"maintenanceWindow"`
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Rrule *string `json:"rrule"`
}

type _UpdateSchedule UpdateSchedule

// NewUpdateSchedule instantiates a new UpdateSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSchedule(enabled *bool, maintenanceWindow *int64, name *string, rrule *string) *UpdateSchedule {
	this := UpdateSchedule{}
	this.Enabled = enabled
	this.MaintenanceWindow = maintenanceWindow
	this.Name = name
	this.Rrule = rrule
	return &this
}

// NewUpdateScheduleWithDefaults instantiates a new UpdateSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateScheduleWithDefaults() *UpdateSchedule {
	this := UpdateSchedule{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *UpdateSchedule) GetEnabled() *bool {
	if o == nil || IsNil(o.Enabled) {
		var ret *bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UpdateSchedule) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled, true
}

// SetEnabled sets field value
func (o *UpdateSchedule) SetEnabled(v *bool) {
	o.Enabled = v
}

// GetMaintenanceWindow returns the MaintenanceWindow field value
func (o *UpdateSchedule) GetMaintenanceWindow() *int64 {
	if o == nil || IsNil(o.MaintenanceWindow) {
		var ret *int64
		return ret
	}

	return o.MaintenanceWindow
}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value
// and a boolean to check if the value has been set.
func (o *UpdateSchedule) GetMaintenanceWindowOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaintenanceWindow, true
}

// SetMaintenanceWindow sets field value
func (o *UpdateSchedule) SetMaintenanceWindow(v *int64) {
	o.MaintenanceWindow = v
}

// GetName returns the Name field value
func (o *UpdateSchedule) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateSchedule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *UpdateSchedule) SetName(v *string) {
	o.Name = v
}

// GetRrule returns the Rrule field value
func (o *UpdateSchedule) GetRrule() *string {
	if o == nil || IsNil(o.Rrule) {
		var ret *string
		return ret
	}

	return o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value
// and a boolean to check if the value has been set.
func (o *UpdateSchedule) GetRruleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rrule, true
}

// SetRrule sets field value
func (o *UpdateSchedule) SetRrule(v *string) {
	o.Rrule = v
}

func (o UpdateSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	toSerialize["name"] = o.Name
	toSerialize["rrule"] = o.Rrule
	return toSerialize, nil
}

type NullableUpdateSchedule struct {
	value *UpdateSchedule
	isSet bool
}

func (v NullableUpdateSchedule) Get() *UpdateSchedule {
	return v.value
}

func (v *NullableUpdateSchedule) Set(val *UpdateSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSchedule(val *UpdateSchedule) *NullableUpdateSchedule {
	return &NullableUpdateSchedule{value: val, isSet: true}
}

func (v NullableUpdateSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
