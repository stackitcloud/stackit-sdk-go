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

// checks if the UpdateBackupPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBackupPayload{}

// UpdateBackupPayload Object that represents a backup.
type UpdateBackupPayload struct {
	// Object that represents an availability zone.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	// Date-time when resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name *string `json:"name,omitempty"`
	// Size in Gigabyte.
	Size *int64 `json:"size,omitempty"`
	// Universally Unique Identifier (UUID).
	SnapshotId *string `json:"snapshotId,omitempty"`
	// The status of a backup object. Possible values: `AVAILABLE`, `CREATING`, `DELETED`, `DELETING`, `ERROR`, `RESTORING`.
	Status *string `json:"status,omitempty"`
	// Date-time when resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Universally Unique Identifier (UUID).
	VolumeId *string `json:"volumeId,omitempty"`
}

// NewUpdateBackupPayload instantiates a new UpdateBackupPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBackupPayload() *UpdateBackupPayload {
	this := UpdateBackupPayload{}
	return &this
}

// NewUpdateBackupPayloadWithDefaults instantiates a new UpdateBackupPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBackupPayloadWithDefaults() *UpdateBackupPayload {
	this := UpdateBackupPayload{}
	return &this
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *UpdateBackupPayload) GetAvailabilityZone() *string {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret *string
		return ret
	}
	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBackupPayload) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityZone) {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *UpdateBackupPayload) HasAvailabilityZone() bool {
	if o != nil && !IsNil(o.AvailabilityZone) {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *UpdateBackupPayload) SetAvailabilityZone(v *string) {
	o.AvailabilityZone = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UpdateBackupPayload) GetCreatedAt() *time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret *time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBackupPayload) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UpdateBackupPayload) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UpdateBackupPayload) SetCreatedAt(v *time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateBackupPayload) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBackupPayload) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateBackupPayload) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateBackupPayload) SetId(v *string) {
	o.Id = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateBackupPayload) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBackupPayload) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *UpdateBackupPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *UpdateBackupPayload) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateBackupPayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBackupPayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateBackupPayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateBackupPayload) SetName(v *string) {
	o.Name = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *UpdateBackupPayload) GetSize() *int64 {
	if o == nil || IsNil(o.Size) {
		var ret *int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBackupPayload) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *UpdateBackupPayload) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *UpdateBackupPayload) SetSize(v *int64) {
	o.Size = v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *UpdateBackupPayload) GetSnapshotId() *string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret *string
		return ret
	}
	return o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBackupPayload) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *UpdateBackupPayload) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *UpdateBackupPayload) SetSnapshotId(v *string) {
	o.SnapshotId = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateBackupPayload) GetStatus() *string {
	if o == nil || IsNil(o.Status) {
		var ret *string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBackupPayload) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateBackupPayload) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateBackupPayload) SetStatus(v *string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UpdateBackupPayload) GetUpdatedAt() *time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret *time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBackupPayload) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UpdateBackupPayload) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *UpdateBackupPayload) SetUpdatedAt(v *time.Time) {
	o.UpdatedAt = v
}

// GetVolumeId returns the VolumeId field value if set, zero value otherwise.
func (o *UpdateBackupPayload) GetVolumeId() *string {
	if o == nil || IsNil(o.VolumeId) {
		var ret *string
		return ret
	}
	return o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBackupPayload) GetVolumeIdOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeId) {
		return nil, false
	}
	return o.VolumeId, true
}

// HasVolumeId returns a boolean if a field has been set.
func (o *UpdateBackupPayload) HasVolumeId() bool {
	if o != nil && !IsNil(o.VolumeId) {
		return true
	}

	return false
}

// SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.
func (o *UpdateBackupPayload) SetVolumeId(v *string) {
	o.VolumeId = v
}

func (o UpdateBackupPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailabilityZone) {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.SnapshotId) {
		toSerialize["snapshotId"] = o.SnapshotId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.VolumeId) {
		toSerialize["volumeId"] = o.VolumeId
	}
	return toSerialize, nil
}

type NullableUpdateBackupPayload struct {
	value *UpdateBackupPayload
	isSet bool
}

func (v NullableUpdateBackupPayload) Get() *UpdateBackupPayload {
	return v.value
}

func (v *NullableUpdateBackupPayload) Set(val *UpdateBackupPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBackupPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBackupPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBackupPayload(val *UpdateBackupPayload) *NullableUpdateBackupPayload {
	return &NullableUpdateBackupPayload{value: val, isSet: true}
}

func (v NullableUpdateBackupPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBackupPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
