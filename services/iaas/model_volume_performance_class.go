/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the VolumePerformanceClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumePerformanceClass{}

// VolumePerformanceClass Object that represents a Volume performance class.
type VolumePerformanceClass struct {
	// Description Object. Allows string up to 127 Characters.
	Description *string `json:"description,omitempty"`
	// Input/Output Operations per second.
	Iops *int64 `json:"iops,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	// REQUIRED
	Name *string `json:"name"`
	// Throughput in Megabyte per second.
	Throughput *int64 `json:"throughput,omitempty"`
}

type _VolumePerformanceClass VolumePerformanceClass

// NewVolumePerformanceClass instantiates a new VolumePerformanceClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumePerformanceClass(name *string) *VolumePerformanceClass {
	this := VolumePerformanceClass{}
	this.Name = name
	return &this
}

// NewVolumePerformanceClassWithDefaults instantiates a new VolumePerformanceClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumePerformanceClassWithDefaults() *VolumePerformanceClass {
	this := VolumePerformanceClass{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VolumePerformanceClass) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePerformanceClass) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VolumePerformanceClass) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VolumePerformanceClass) SetDescription(v *string) {
	o.Description = v
}

// GetIops returns the Iops field value if set, zero value otherwise.
func (o *VolumePerformanceClass) GetIops() *int64 {
	if o == nil || IsNil(o.Iops) {
		var ret *int64
		return ret
	}
	return o.Iops
}

// GetIopsOk returns a tuple with the Iops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePerformanceClass) GetIopsOk() (*int64, bool) {
	if o == nil || IsNil(o.Iops) {
		return nil, false
	}
	return o.Iops, true
}

// HasIops returns a boolean if a field has been set.
func (o *VolumePerformanceClass) HasIops() bool {
	if o != nil && !IsNil(o.Iops) {
		return true
	}

	return false
}

// SetIops gets a reference to the given int64 and assigns it to the Iops field.
func (o *VolumePerformanceClass) SetIops(v *int64) {
	o.Iops = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *VolumePerformanceClass) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePerformanceClass) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *VolumePerformanceClass) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *VolumePerformanceClass) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value
func (o *VolumePerformanceClass) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VolumePerformanceClass) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *VolumePerformanceClass) SetName(v *string) {
	o.Name = v
}

// GetThroughput returns the Throughput field value if set, zero value otherwise.
func (o *VolumePerformanceClass) GetThroughput() *int64 {
	if o == nil || IsNil(o.Throughput) {
		var ret *int64
		return ret
	}
	return o.Throughput
}

// GetThroughputOk returns a tuple with the Throughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePerformanceClass) GetThroughputOk() (*int64, bool) {
	if o == nil || IsNil(o.Throughput) {
		return nil, false
	}
	return o.Throughput, true
}

// HasThroughput returns a boolean if a field has been set.
func (o *VolumePerformanceClass) HasThroughput() bool {
	if o != nil && !IsNil(o.Throughput) {
		return true
	}

	return false
}

// SetThroughput gets a reference to the given int64 and assigns it to the Throughput field.
func (o *VolumePerformanceClass) SetThroughput(v *int64) {
	o.Throughput = v
}

func (o VolumePerformanceClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Iops) {
		toSerialize["iops"] = o.Iops
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Throughput) {
		toSerialize["throughput"] = o.Throughput
	}
	return toSerialize, nil
}

type NullableVolumePerformanceClass struct {
	value *VolumePerformanceClass
	isSet bool
}

func (v NullableVolumePerformanceClass) Get() *VolumePerformanceClass {
	return v.value
}

func (v *NullableVolumePerformanceClass) Set(val *VolumePerformanceClass) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumePerformanceClass) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumePerformanceClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumePerformanceClass(val *VolumePerformanceClass) *NullableVolumePerformanceClass {
	return &NullableVolumePerformanceClass{value: val, isSet: true}
}

func (v NullableVolumePerformanceClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumePerformanceClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
