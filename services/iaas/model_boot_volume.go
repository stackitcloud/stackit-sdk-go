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

// checks if the BootVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BootVolume{}

// BootVolume The boot device for the server.
type BootVolume struct {
	// Delete the volume during the termination of the server. Defaults to false.
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty"`
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	PerformanceClass *string `json:"performanceClass,omitempty"`
	// Size in Gigabyte.
	Size   *int64            `json:"size,omitempty"`
	Source *BootVolumeSource `json:"source,omitempty"`
}

// NewBootVolume instantiates a new BootVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootVolume() *BootVolume {
	this := BootVolume{}
	var deleteOnTermination bool = false
	this.DeleteOnTermination = &deleteOnTermination
	return &this
}

// NewBootVolumeWithDefaults instantiates a new BootVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootVolumeWithDefaults() *BootVolume {
	this := BootVolume{}
	var deleteOnTermination bool = false
	this.DeleteOnTermination = &deleteOnTermination
	return &this
}

// GetDeleteOnTermination returns the DeleteOnTermination field value if set, zero value otherwise.
func (o *BootVolume) GetDeleteOnTermination() *bool {
	if o == nil || IsNil(o.DeleteOnTermination) {
		var ret *bool
		return ret
	}
	return o.DeleteOnTermination
}

// GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVolume) GetDeleteOnTerminationOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteOnTermination) {
		return nil, false
	}
	return o.DeleteOnTermination, true
}

// HasDeleteOnTermination returns a boolean if a field has been set.
func (o *BootVolume) HasDeleteOnTermination() bool {
	if o != nil && !IsNil(o.DeleteOnTermination) {
		return true
	}

	return false
}

// SetDeleteOnTermination gets a reference to the given bool and assigns it to the DeleteOnTermination field.
func (o *BootVolume) SetDeleteOnTermination(v *bool) {
	o.DeleteOnTermination = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BootVolume) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVolume) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BootVolume) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BootVolume) SetId(v *string) {
	o.Id = v
}

// GetPerformanceClass returns the PerformanceClass field value if set, zero value otherwise.
func (o *BootVolume) GetPerformanceClass() *string {
	if o == nil || IsNil(o.PerformanceClass) {
		var ret *string
		return ret
	}
	return o.PerformanceClass
}

// GetPerformanceClassOk returns a tuple with the PerformanceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVolume) GetPerformanceClassOk() (*string, bool) {
	if o == nil || IsNil(o.PerformanceClass) {
		return nil, false
	}
	return o.PerformanceClass, true
}

// HasPerformanceClass returns a boolean if a field has been set.
func (o *BootVolume) HasPerformanceClass() bool {
	if o != nil && !IsNil(o.PerformanceClass) {
		return true
	}

	return false
}

// SetPerformanceClass gets a reference to the given string and assigns it to the PerformanceClass field.
func (o *BootVolume) SetPerformanceClass(v *string) {
	o.PerformanceClass = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BootVolume) GetSize() *int64 {
	if o == nil || IsNil(o.Size) {
		var ret *int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVolume) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BootVolume) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *BootVolume) SetSize(v *int64) {
	o.Size = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BootVolume) GetSource() *BootVolumeSource {
	if o == nil || IsNil(o.Source) {
		var ret *BootVolumeSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVolume) GetSourceOk() (*BootVolumeSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BootVolume) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given BootVolumeSource and assigns it to the Source field.
func (o *BootVolume) SetSource(v *BootVolumeSource) {
	o.Source = v
}

func (o BootVolume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeleteOnTermination) {
		toSerialize["deleteOnTermination"] = o.DeleteOnTermination
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PerformanceClass) {
		toSerialize["performanceClass"] = o.PerformanceClass
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableBootVolume struct {
	value *BootVolume
	isSet bool
}

func (v NullableBootVolume) Get() *BootVolume {
	return v.value
}

func (v *NullableBootVolume) Set(val *BootVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableBootVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableBootVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootVolume(val *BootVolume) *NullableBootVolume {
	return &NullableBootVolume{value: val, isSet: true}
}

func (v NullableBootVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
