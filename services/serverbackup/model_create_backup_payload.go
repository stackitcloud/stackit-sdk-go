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

// checks if the CreateBackupPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBackupPayload{}

// CreateBackupPayload struct for CreateBackupPayload
type CreateBackupPayload struct {
	// Max 255 characters
	// REQUIRED
	Name *string `json:"name"`
	// Values are set in days (1-36500)
	// REQUIRED
	RetentionPeriod *int64    `json:"retentionPeriod"`
	VolumeIds       *[]string `json:"volumeIds,omitempty"`
}

type _CreateBackupPayload CreateBackupPayload

// NewCreateBackupPayload instantiates a new CreateBackupPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBackupPayload(name *string, retentionPeriod *int64) *CreateBackupPayload {
	this := CreateBackupPayload{}
	this.Name = name
	this.RetentionPeriod = retentionPeriod
	return &this
}

// NewCreateBackupPayloadWithDefaults instantiates a new CreateBackupPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBackupPayloadWithDefaults() *CreateBackupPayload {
	this := CreateBackupPayload{}
	return &this
}

// GetName returns the Name field value
func (o *CreateBackupPayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateBackupPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *CreateBackupPayload) SetName(v *string) {
	o.Name = v
}

// GetRetentionPeriod returns the RetentionPeriod field value
func (o *CreateBackupPayload) GetRetentionPeriod() *int64 {
	if o == nil || IsNil(o.RetentionPeriod) {
		var ret *int64
		return ret
	}

	return o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value
// and a boolean to check if the value has been set.
func (o *CreateBackupPayload) GetRetentionPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// SetRetentionPeriod sets field value
func (o *CreateBackupPayload) SetRetentionPeriod(v *int64) {
	o.RetentionPeriod = v
}

// GetVolumeIds returns the VolumeIds field value if set, zero value otherwise.
func (o *CreateBackupPayload) GetVolumeIds() *[]string {
	if o == nil || IsNil(o.VolumeIds) {
		var ret *[]string
		return ret
	}
	return o.VolumeIds
}

// GetVolumeIdsOk returns a tuple with the VolumeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBackupPayload) GetVolumeIdsOk() (*[]string, bool) {
	if o == nil || IsNil(o.VolumeIds) {
		return nil, false
	}
	return o.VolumeIds, true
}

// HasVolumeIds returns a boolean if a field has been set.
func (o *CreateBackupPayload) HasVolumeIds() bool {
	if o != nil && !IsNil(o.VolumeIds) && !IsNil(o.VolumeIds) {
		return true
	}

	return false
}

// SetVolumeIds gets a reference to the given []string and assigns it to the VolumeIds field.
func (o *CreateBackupPayload) SetVolumeIds(v *[]string) {
	o.VolumeIds = v
}

func (o CreateBackupPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["retentionPeriod"] = o.RetentionPeriod
	if !IsNil(o.VolumeIds) {
		toSerialize["volumeIds"] = o.VolumeIds
	}
	return toSerialize, nil
}

type NullableCreateBackupPayload struct {
	value *CreateBackupPayload
	isSet bool
}

func (v NullableCreateBackupPayload) Get() *CreateBackupPayload {
	return v.value
}

func (v *NullableCreateBackupPayload) Set(val *CreateBackupPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBackupPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBackupPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBackupPayload(val *CreateBackupPayload) *NullableCreateBackupPayload {
	return &NullableCreateBackupPayload{value: val, isSet: true}
}

func (v NullableCreateBackupPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBackupPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
