/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
	"time"
)

// checks if the ServerMaintenance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerMaintenance{}

// ServerMaintenance Object that represents the information about the next planned server maintenance window.
type ServerMaintenance struct {
	Details *string `json:"details,omitempty"`
	// End of the maintenance window.
	// REQUIRED
	EndsAt *time.Time `json:"endsAt"`
	// Start of the maintenance window.
	// REQUIRED
	StartsAt *time.Time `json:"startsAt"`
	// REQUIRED
	Status *string `json:"status"`
}

type _ServerMaintenance ServerMaintenance

// NewServerMaintenance instantiates a new ServerMaintenance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerMaintenance(endsAt *time.Time, startsAt *time.Time, status *string) *ServerMaintenance {
	this := ServerMaintenance{}
	this.EndsAt = endsAt
	this.StartsAt = startsAt
	this.Status = status
	return &this
}

// NewServerMaintenanceWithDefaults instantiates a new ServerMaintenance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerMaintenanceWithDefaults() *ServerMaintenance {
	this := ServerMaintenance{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ServerMaintenance) GetDetails() *string {
	if o == nil || IsNil(o.Details) {
		var ret *string
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerMaintenance) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ServerMaintenance) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ServerMaintenance) SetDetails(v *string) {
	o.Details = v
}

// GetEndsAt returns the EndsAt field value
func (o *ServerMaintenance) GetEndsAt() *time.Time {
	if o == nil || IsNil(o.EndsAt) {
		var ret *time.Time
		return ret
	}

	return o.EndsAt
}

// GetEndsAtOk returns a tuple with the EndsAt field value
// and a boolean to check if the value has been set.
func (o *ServerMaintenance) GetEndsAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndsAt, true
}

// SetEndsAt sets field value
func (o *ServerMaintenance) SetEndsAt(v *time.Time) {
	o.EndsAt = v
}

// GetStartsAt returns the StartsAt field value
func (o *ServerMaintenance) GetStartsAt() *time.Time {
	if o == nil || IsNil(o.StartsAt) {
		var ret *time.Time
		return ret
	}

	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value
// and a boolean to check if the value has been set.
func (o *ServerMaintenance) GetStartsAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartsAt, true
}

// SetStartsAt sets field value
func (o *ServerMaintenance) SetStartsAt(v *time.Time) {
	o.StartsAt = v
}

// GetStatus returns the Status field value
func (o *ServerMaintenance) GetStatus() *string {
	if o == nil || IsNil(o.Status) {
		var ret *string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ServerMaintenance) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *ServerMaintenance) SetStatus(v *string) {
	o.Status = v
}

func (o ServerMaintenance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	toSerialize["endsAt"] = o.EndsAt
	toSerialize["startsAt"] = o.StartsAt
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableServerMaintenance struct {
	value *ServerMaintenance
	isSet bool
}

func (v NullableServerMaintenance) Get() *ServerMaintenance {
	return v.value
}

func (v *NullableServerMaintenance) Set(val *ServerMaintenance) {
	v.value = val
	v.isSet = true
}

func (v NullableServerMaintenance) IsSet() bool {
	return v.isSet
}

func (v *NullableServerMaintenance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerMaintenance(val *ServerMaintenance) *NullableServerMaintenance {
	return &NullableServerMaintenance{value: val, isSet: true}
}

func (v NullableServerMaintenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerMaintenance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
