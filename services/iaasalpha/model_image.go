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

// checks if the Image type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Image{}

// Image Object that represents an Image and its parameters. Used for Creating and returning (get/list).
type Image struct {
	Checksum *ImageChecksum `json:"checksum,omitempty"`
	Config   *ImageConfig   `json:"config,omitempty"`
	// Date-time when resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Object that represents a disk format.
	// REQUIRED
	DiskFormat *string `json:"diskFormat"`
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// Size in Gigabyte.
	MinDiskSize *int64 `json:"minDiskSize,omitempty"`
	// Size in Megabyte.
	MinRam *int64 `json:"minRam,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	// REQUIRED
	Name *string `json:"name"`
	// Universally Unique Identifier (UUID).
	Owner     *string `json:"owner,omitempty"`
	Protected *bool   `json:"protected,omitempty"`
	// Scope of an Image.
	Scope *string `json:"scope,omitempty"`
	// The status of an image object.
	Status *string `json:"status,omitempty"`
	// Date-time when resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type _Image Image

// NewImage instantiates a new Image object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImage(diskFormat *string, name *string) *Image {
	this := Image{}
	this.DiskFormat = diskFormat
	this.Name = name
	return &this
}

// NewImageWithDefaults instantiates a new Image object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWithDefaults() *Image {
	this := Image{}
	return &this
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *Image) GetChecksum() *ImageChecksum {
	if o == nil || IsNil(o.Checksum) {
		var ret *ImageChecksum
		return ret
	}
	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetChecksumOk() (*ImageChecksum, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *Image) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given ImageChecksum and assigns it to the Checksum field.
func (o *Image) SetChecksum(v *ImageChecksum) {
	o.Checksum = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Image) GetConfig() *ImageConfig {
	if o == nil || IsNil(o.Config) {
		var ret *ImageConfig
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetConfigOk() (*ImageConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Image) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ImageConfig and assigns it to the Config field.
func (o *Image) SetConfig(v *ImageConfig) {
	o.Config = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Image) GetCreatedAt() *time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret *time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Image) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Image) SetCreatedAt(v *time.Time) {
	o.CreatedAt = v
}

// GetDiskFormat returns the DiskFormat field value
func (o *Image) GetDiskFormat() *string {
	if o == nil || IsNil(o.DiskFormat) {
		var ret *string
		return ret
	}

	return o.DiskFormat
}

// GetDiskFormatOk returns a tuple with the DiskFormat field value
// and a boolean to check if the value has been set.
func (o *Image) GetDiskFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskFormat, true
}

// SetDiskFormat sets field value
func (o *Image) SetDiskFormat(v *string) {
	o.DiskFormat = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Image) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Image) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Image) SetId(v *string) {
	o.Id = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Image) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Image) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *Image) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetMinDiskSize returns the MinDiskSize field value if set, zero value otherwise.
func (o *Image) GetMinDiskSize() *int64 {
	if o == nil || IsNil(o.MinDiskSize) {
		var ret *int64
		return ret
	}
	return o.MinDiskSize
}

// GetMinDiskSizeOk returns a tuple with the MinDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetMinDiskSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MinDiskSize) {
		return nil, false
	}
	return o.MinDiskSize, true
}

// HasMinDiskSize returns a boolean if a field has been set.
func (o *Image) HasMinDiskSize() bool {
	if o != nil && !IsNil(o.MinDiskSize) {
		return true
	}

	return false
}

// SetMinDiskSize gets a reference to the given int64 and assigns it to the MinDiskSize field.
func (o *Image) SetMinDiskSize(v *int64) {
	o.MinDiskSize = v
}

// GetMinRam returns the MinRam field value if set, zero value otherwise.
func (o *Image) GetMinRam() *int64 {
	if o == nil || IsNil(o.MinRam) {
		var ret *int64
		return ret
	}
	return o.MinRam
}

// GetMinRamOk returns a tuple with the MinRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetMinRamOk() (*int64, bool) {
	if o == nil || IsNil(o.MinRam) {
		return nil, false
	}
	return o.MinRam, true
}

// HasMinRam returns a boolean if a field has been set.
func (o *Image) HasMinRam() bool {
	if o != nil && !IsNil(o.MinRam) {
		return true
	}

	return false
}

// SetMinRam gets a reference to the given int64 and assigns it to the MinRam field.
func (o *Image) SetMinRam(v *int64) {
	o.MinRam = v
}

// GetName returns the Name field value
func (o *Image) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Image) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *Image) SetName(v *string) {
	o.Name = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Image) GetOwner() *string {
	if o == nil || IsNil(o.Owner) {
		var ret *string
		return ret
	}
	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Image) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *Image) SetOwner(v *string) {
	o.Owner = v
}

// GetProtected returns the Protected field value if set, zero value otherwise.
func (o *Image) GetProtected() *bool {
	if o == nil || IsNil(o.Protected) {
		var ret *bool
		return ret
	}
	return o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Protected) {
		return nil, false
	}
	return o.Protected, true
}

// HasProtected returns a boolean if a field has been set.
func (o *Image) HasProtected() bool {
	if o != nil && !IsNil(o.Protected) {
		return true
	}

	return false
}

// SetProtected gets a reference to the given bool and assigns it to the Protected field.
func (o *Image) SetProtected(v *bool) {
	o.Protected = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *Image) GetScope() *string {
	if o == nil || IsNil(o.Scope) {
		var ret *string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *Image) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *Image) SetScope(v *string) {
	o.Scope = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Image) GetStatus() *string {
	if o == nil || IsNil(o.Status) {
		var ret *string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Image) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Image) SetStatus(v *string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Image) GetUpdatedAt() *time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret *time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Image) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Image) SetUpdatedAt(v *time.Time) {
	o.UpdatedAt = v
}

func (o Image) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Checksum) {
		toSerialize["checksum"] = o.Checksum
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	toSerialize["diskFormat"] = o.DiskFormat
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.MinDiskSize) {
		toSerialize["minDiskSize"] = o.MinDiskSize
	}
	if !IsNil(o.MinRam) {
		toSerialize["minRam"] = o.MinRam
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Protected) {
		toSerialize["protected"] = o.Protected
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableImage struct {
	value *Image
	isSet bool
}

func (v NullableImage) Get() *Image {
	return v.value
}

func (v *NullableImage) Set(val *Image) {
	v.value = val
	v.isSet = true
}

func (v NullableImage) IsSet() bool {
	return v.isSet
}

func (v *NullableImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImage(val *Image) *NullableImage {
	return &NullableImage{value: val, isSet: true}
}

func (v NullableImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
