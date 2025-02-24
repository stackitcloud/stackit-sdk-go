/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
)

// checks if the UpdateImagePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateImagePayload{}

// UpdateImagePayload Object that represents an update request body of an Image.
type UpdateImagePayload struct {
	Config *ImageConfig `json:"config,omitempty"`
	// Object that represents a disk format. Possible values: `raw`, `qcow2`, `iso`.
	DiskFormat *string `json:"diskFormat,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// Size in Gigabyte.
	MinDiskSize *int64 `json:"minDiskSize,omitempty"`
	// Size in Megabyte.
	MinRam *int64 `json:"minRam,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name      *string `json:"name,omitempty"`
	Protected *bool   `json:"protected,omitempty"`
}

// NewUpdateImagePayload instantiates a new UpdateImagePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateImagePayload() *UpdateImagePayload {
	this := UpdateImagePayload{}
	return &this
}

// NewUpdateImagePayloadWithDefaults instantiates a new UpdateImagePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateImagePayloadWithDefaults() *UpdateImagePayload {
	this := UpdateImagePayload{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *UpdateImagePayload) GetConfig() *ImageConfig {
	if o == nil || IsNil(o.Config) {
		var ret *ImageConfig
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImagePayload) GetConfigOk() (*ImageConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *UpdateImagePayload) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ImageConfig and assigns it to the Config field.
func (o *UpdateImagePayload) SetConfig(v *ImageConfig) {
	o.Config = v
}

// GetDiskFormat returns the DiskFormat field value if set, zero value otherwise.
func (o *UpdateImagePayload) GetDiskFormat() *string {
	if o == nil || IsNil(o.DiskFormat) {
		var ret *string
		return ret
	}
	return o.DiskFormat
}

// GetDiskFormatOk returns a tuple with the DiskFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImagePayload) GetDiskFormatOk() (*string, bool) {
	if o == nil || IsNil(o.DiskFormat) {
		return nil, false
	}
	return o.DiskFormat, true
}

// HasDiskFormat returns a boolean if a field has been set.
func (o *UpdateImagePayload) HasDiskFormat() bool {
	if o != nil && !IsNil(o.DiskFormat) {
		return true
	}

	return false
}

// SetDiskFormat gets a reference to the given string and assigns it to the DiskFormat field.
func (o *UpdateImagePayload) SetDiskFormat(v *string) {
	o.DiskFormat = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateImagePayload) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImagePayload) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *UpdateImagePayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *UpdateImagePayload) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetMinDiskSize returns the MinDiskSize field value if set, zero value otherwise.
func (o *UpdateImagePayload) GetMinDiskSize() *int64 {
	if o == nil || IsNil(o.MinDiskSize) {
		var ret *int64
		return ret
	}
	return o.MinDiskSize
}

// GetMinDiskSizeOk returns a tuple with the MinDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImagePayload) GetMinDiskSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MinDiskSize) {
		return nil, false
	}
	return o.MinDiskSize, true
}

// HasMinDiskSize returns a boolean if a field has been set.
func (o *UpdateImagePayload) HasMinDiskSize() bool {
	if o != nil && !IsNil(o.MinDiskSize) {
		return true
	}

	return false
}

// SetMinDiskSize gets a reference to the given int64 and assigns it to the MinDiskSize field.
func (o *UpdateImagePayload) SetMinDiskSize(v *int64) {
	o.MinDiskSize = v
}

// GetMinRam returns the MinRam field value if set, zero value otherwise.
func (o *UpdateImagePayload) GetMinRam() *int64 {
	if o == nil || IsNil(o.MinRam) {
		var ret *int64
		return ret
	}
	return o.MinRam
}

// GetMinRamOk returns a tuple with the MinRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImagePayload) GetMinRamOk() (*int64, bool) {
	if o == nil || IsNil(o.MinRam) {
		return nil, false
	}
	return o.MinRam, true
}

// HasMinRam returns a boolean if a field has been set.
func (o *UpdateImagePayload) HasMinRam() bool {
	if o != nil && !IsNil(o.MinRam) {
		return true
	}

	return false
}

// SetMinRam gets a reference to the given int64 and assigns it to the MinRam field.
func (o *UpdateImagePayload) SetMinRam(v *int64) {
	o.MinRam = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateImagePayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImagePayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateImagePayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateImagePayload) SetName(v *string) {
	o.Name = v
}

// GetProtected returns the Protected field value if set, zero value otherwise.
func (o *UpdateImagePayload) GetProtected() *bool {
	if o == nil || IsNil(o.Protected) {
		var ret *bool
		return ret
	}
	return o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImagePayload) GetProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Protected) {
		return nil, false
	}
	return o.Protected, true
}

// HasProtected returns a boolean if a field has been set.
func (o *UpdateImagePayload) HasProtected() bool {
	if o != nil && !IsNil(o.Protected) {
		return true
	}

	return false
}

// SetProtected gets a reference to the given bool and assigns it to the Protected field.
func (o *UpdateImagePayload) SetProtected(v *bool) {
	o.Protected = v
}

func (o UpdateImagePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.DiskFormat) {
		toSerialize["diskFormat"] = o.DiskFormat
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Protected) {
		toSerialize["protected"] = o.Protected
	}
	return toSerialize, nil
}

type NullableUpdateImagePayload struct {
	value *UpdateImagePayload
	isSet bool
}

func (v NullableUpdateImagePayload) Get() *UpdateImagePayload {
	return v.value
}

func (v *NullableUpdateImagePayload) Set(val *UpdateImagePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateImagePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateImagePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateImagePayload(val *UpdateImagePayload) *NullableUpdateImagePayload {
	return &NullableUpdateImagePayload{value: val, isSet: true}
}

func (v NullableUpdateImagePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateImagePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
