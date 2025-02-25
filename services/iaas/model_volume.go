/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
	"time"
)

// checks if the Volume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Volume{}

// Volume Object that represents a volume and its parameters. Used for Creating and returning (get/list).
type Volume struct {
	// Object that represents an availability zone.
	// REQUIRED
	AvailabilityZone *string `json:"availabilityZone"`
	// Indicates if a volume is bootable.
	Bootable *bool `json:"bootable,omitempty"`
	// Date-time when resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Description Object. Allows string up to 127 Characters.
	Description *string `json:"description,omitempty"`
	// Universally Unique Identifier (UUID).
	Id          *string      `json:"id,omitempty"`
	ImageConfig *ImageConfig `json:"imageConfig,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name *string `json:"name,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	PerformanceClass *string `json:"performanceClass,omitempty"`
	// Universally Unique Identifier (UUID).
	ServerId *string `json:"serverId,omitempty"`
	// Size in Gigabyte.
	Size   *int64        `json:"size,omitempty"`
	Source *VolumeSource `json:"source,omitempty"`
	// The status of a volume object. Possible values: `ATTACHED`, `ATTACHING`, `AVAILABLE`, `AWAITING-TRANSFER`, `BACKING-UP`, `CREATING`, `DELETED`, `DELETING`, `DETACHING`, `DOWNLOADING`, `ERROR`, `ERROR_BACKING-UP`, `ERROR_DELETING`, `ERROR_RESIZING`, `ERROR_RESTORING-BACKUP`, `MAINTENANCE`, `RESERVED`, `RESIZING`, `RESTORING-BACKUP`, `RETYPING`, `UPLOADING`.
	Status *string `json:"status,omitempty"`
	// Date-time when resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type _Volume Volume

// NewVolume instantiates a new Volume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolume(availabilityZone *string) *Volume {
	this := Volume{}
	this.AvailabilityZone = availabilityZone
	return &this
}

// NewVolumeWithDefaults instantiates a new Volume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeWithDefaults() *Volume {
	this := Volume{}
	return &this
}

// GetAvailabilityZone returns the AvailabilityZone field value
func (o *Volume) GetAvailabilityZone() *string {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret *string
		return ret
	}

	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
func (o *Volume) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *Volume) SetAvailabilityZone(v *string) {
	o.AvailabilityZone = v
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *Volume) GetBootable() *bool {
	if o == nil || IsNil(o.Bootable) {
		var ret *bool
		return ret
	}
	return o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetBootableOk() (*bool, bool) {
	if o == nil || IsNil(o.Bootable) {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *Volume) HasBootable() bool {
	if o != nil && !IsNil(o.Bootable) {
		return true
	}

	return false
}

// SetBootable gets a reference to the given bool and assigns it to the Bootable field.
func (o *Volume) SetBootable(v *bool) {
	o.Bootable = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Volume) GetCreatedAt() *time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret *time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Volume) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Volume) SetCreatedAt(v *time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Volume) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Volume) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Volume) SetDescription(v *string) {
	o.Description = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Volume) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Volume) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Volume) SetId(v *string) {
	o.Id = v
}

// GetImageConfig returns the ImageConfig field value if set, zero value otherwise.
func (o *Volume) GetImageConfig() *ImageConfig {
	if o == nil || IsNil(o.ImageConfig) {
		var ret *ImageConfig
		return ret
	}
	return o.ImageConfig
}

// GetImageConfigOk returns a tuple with the ImageConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetImageConfigOk() (*ImageConfig, bool) {
	if o == nil || IsNil(o.ImageConfig) {
		return nil, false
	}
	return o.ImageConfig, true
}

// HasImageConfig returns a boolean if a field has been set.
func (o *Volume) HasImageConfig() bool {
	if o != nil && !IsNil(o.ImageConfig) {
		return true
	}

	return false
}

// SetImageConfig gets a reference to the given ImageConfig and assigns it to the ImageConfig field.
func (o *Volume) SetImageConfig(v *ImageConfig) {
	o.ImageConfig = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Volume) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Volume) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *Volume) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Volume) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Volume) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Volume) SetName(v *string) {
	o.Name = v
}

// GetPerformanceClass returns the PerformanceClass field value if set, zero value otherwise.
func (o *Volume) GetPerformanceClass() *string {
	if o == nil || IsNil(o.PerformanceClass) {
		var ret *string
		return ret
	}
	return o.PerformanceClass
}

// GetPerformanceClassOk returns a tuple with the PerformanceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetPerformanceClassOk() (*string, bool) {
	if o == nil || IsNil(o.PerformanceClass) {
		return nil, false
	}
	return o.PerformanceClass, true
}

// HasPerformanceClass returns a boolean if a field has been set.
func (o *Volume) HasPerformanceClass() bool {
	if o != nil && !IsNil(o.PerformanceClass) {
		return true
	}

	return false
}

// SetPerformanceClass gets a reference to the given string and assigns it to the PerformanceClass field.
func (o *Volume) SetPerformanceClass(v *string) {
	o.PerformanceClass = v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *Volume) GetServerId() *string {
	if o == nil || IsNil(o.ServerId) {
		var ret *string
		return ret
	}
	return o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *Volume) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *Volume) SetServerId(v *string) {
	o.ServerId = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Volume) GetSize() *int64 {
	if o == nil || IsNil(o.Size) {
		var ret *int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Volume) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *Volume) SetSize(v *int64) {
	o.Size = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Volume) GetSource() *VolumeSource {
	if o == nil || IsNil(o.Source) {
		var ret *VolumeSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetSourceOk() (*VolumeSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Volume) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given VolumeSource and assigns it to the Source field.
func (o *Volume) SetSource(v *VolumeSource) {
	o.Source = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Volume) GetStatus() *string {
	if o == nil || IsNil(o.Status) {
		var ret *string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Volume) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Volume) SetStatus(v *string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Volume) GetUpdatedAt() *time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret *time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Volume) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Volume) SetUpdatedAt(v *time.Time) {
	o.UpdatedAt = v
}

func (o Volume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["availabilityZone"] = o.AvailabilityZone
	if !IsNil(o.Bootable) {
		toSerialize["bootable"] = o.Bootable
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ImageConfig) {
		toSerialize["imageConfig"] = o.ImageConfig
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PerformanceClass) {
		toSerialize["performanceClass"] = o.PerformanceClass
	}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableVolume struct {
	value *Volume
	isSet bool
}

func (v NullableVolume) Get() *Volume {
	return v.value
}

func (v *NullableVolume) Set(val *Volume) {
	v.value = val
	v.isSet = true
}

func (v NullableVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolume(val *Volume) *NullableVolume {
	return &NullableVolume{value: val, isSet: true}
}

func (v NullableVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
