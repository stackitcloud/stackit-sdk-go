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

// checks if the CreateSnapshotPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSnapshotPayload{}

// CreateSnapshotPayload Object that represents a snapshot.
type CreateSnapshotPayload struct {
	// Date-time when resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name *string `json:"name,omitempty"`
	// Size in Gigabyte.
	Size *int64 `json:"size,omitempty"`
	// The status of a snapshot object.
	Status *string `json:"status,omitempty"`
	// Date-time when resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Universally Unique Identifier (UUID).
	// REQUIRED
	VolumeId *string `json:"volumeId"`
}

type _CreateSnapshotPayload CreateSnapshotPayload

// NewCreateSnapshotPayload instantiates a new CreateSnapshotPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSnapshotPayload(volumeId *string) *CreateSnapshotPayload {
	this := CreateSnapshotPayload{}
	this.VolumeId = volumeId
	return &this
}

// NewCreateSnapshotPayloadWithDefaults instantiates a new CreateSnapshotPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSnapshotPayloadWithDefaults() *CreateSnapshotPayload {
	this := CreateSnapshotPayload{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateSnapshotPayload) GetCreatedAt() *time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret *time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotPayload) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateSnapshotPayload) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CreateSnapshotPayload) SetCreatedAt(v *time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateSnapshotPayload) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotPayload) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateSnapshotPayload) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateSnapshotPayload) SetId(v *string) {
	o.Id = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CreateSnapshotPayload) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotPayload) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CreateSnapshotPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *CreateSnapshotPayload) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateSnapshotPayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotPayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateSnapshotPayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateSnapshotPayload) SetName(v *string) {
	o.Name = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CreateSnapshotPayload) GetSize() *int64 {
	if o == nil || IsNil(o.Size) {
		var ret *int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotPayload) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CreateSnapshotPayload) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *CreateSnapshotPayload) SetSize(v *int64) {
	o.Size = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateSnapshotPayload) GetStatus() *string {
	if o == nil || IsNil(o.Status) {
		var ret *string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotPayload) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateSnapshotPayload) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateSnapshotPayload) SetStatus(v *string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CreateSnapshotPayload) GetUpdatedAt() *time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret *time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotPayload) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CreateSnapshotPayload) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CreateSnapshotPayload) SetUpdatedAt(v *time.Time) {
	o.UpdatedAt = v
}

// GetVolumeId returns the VolumeId field value
func (o *CreateSnapshotPayload) GetVolumeId() *string {
	if o == nil || IsNil(o.VolumeId) {
		var ret *string
		return ret
	}

	return o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value
// and a boolean to check if the value has been set.
func (o *CreateSnapshotPayload) GetVolumeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeId, true
}

// SetVolumeId sets field value
func (o *CreateSnapshotPayload) SetVolumeId(v *string) {
	o.VolumeId = v
}

func (o CreateSnapshotPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	toSerialize["volumeId"] = o.VolumeId
	return toSerialize, nil
}

type NullableCreateSnapshotPayload struct {
	value *CreateSnapshotPayload
	isSet bool
}

func (v NullableCreateSnapshotPayload) Get() *CreateSnapshotPayload {
	return v.value
}

func (v *NullableCreateSnapshotPayload) Set(val *CreateSnapshotPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSnapshotPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSnapshotPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSnapshotPayload(val *CreateSnapshotPayload) *NullableCreateSnapshotPayload {
	return &NullableCreateSnapshotPayload{value: val, isSet: true}
}

func (v NullableCreateSnapshotPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSnapshotPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
