/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the CreateServerPayloadBootVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServerPayloadBootVolume{}

// CreateServerPayloadBootVolume struct for CreateServerPayloadBootVolume
type CreateServerPayloadBootVolume struct {
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

// NewCreateServerPayloadBootVolume instantiates a new CreateServerPayloadBootVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServerPayloadBootVolume() *CreateServerPayloadBootVolume {
	this := CreateServerPayloadBootVolume{}
	var deleteOnTermination bool = false
	this.DeleteOnTermination = &deleteOnTermination
	return &this
}

// NewCreateServerPayloadBootVolumeWithDefaults instantiates a new CreateServerPayloadBootVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServerPayloadBootVolumeWithDefaults() *CreateServerPayloadBootVolume {
	this := CreateServerPayloadBootVolume{}
	var deleteOnTermination bool = false
	this.DeleteOnTermination = &deleteOnTermination
	return &this
}

// GetDeleteOnTermination returns the DeleteOnTermination field value if set, zero value otherwise.
func (o *CreateServerPayloadBootVolume) GetDeleteOnTermination() *bool {
	if o == nil || IsNil(o.DeleteOnTermination) {
		var ret *bool
		return ret
	}
	return o.DeleteOnTermination
}

// GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerPayloadBootVolume) GetDeleteOnTerminationOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteOnTermination) {
		return nil, false
	}
	return o.DeleteOnTermination, true
}

// HasDeleteOnTermination returns a boolean if a field has been set.
func (o *CreateServerPayloadBootVolume) HasDeleteOnTermination() bool {
	if o != nil && !IsNil(o.DeleteOnTermination) {
		return true
	}

	return false
}

// SetDeleteOnTermination gets a reference to the given bool and assigns it to the DeleteOnTermination field.
func (o *CreateServerPayloadBootVolume) SetDeleteOnTermination(v *bool) {
	o.DeleteOnTermination = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateServerPayloadBootVolume) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerPayloadBootVolume) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateServerPayloadBootVolume) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateServerPayloadBootVolume) SetId(v *string) {
	o.Id = v
}

// GetPerformanceClass returns the PerformanceClass field value if set, zero value otherwise.
func (o *CreateServerPayloadBootVolume) GetPerformanceClass() *string {
	if o == nil || IsNil(o.PerformanceClass) {
		var ret *string
		return ret
	}
	return o.PerformanceClass
}

// GetPerformanceClassOk returns a tuple with the PerformanceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerPayloadBootVolume) GetPerformanceClassOk() (*string, bool) {
	if o == nil || IsNil(o.PerformanceClass) {
		return nil, false
	}
	return o.PerformanceClass, true
}

// HasPerformanceClass returns a boolean if a field has been set.
func (o *CreateServerPayloadBootVolume) HasPerformanceClass() bool {
	if o != nil && !IsNil(o.PerformanceClass) {
		return true
	}

	return false
}

// SetPerformanceClass gets a reference to the given string and assigns it to the PerformanceClass field.
func (o *CreateServerPayloadBootVolume) SetPerformanceClass(v *string) {
	o.PerformanceClass = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CreateServerPayloadBootVolume) GetSize() *int64 {
	if o == nil || IsNil(o.Size) {
		var ret *int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerPayloadBootVolume) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CreateServerPayloadBootVolume) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *CreateServerPayloadBootVolume) SetSize(v *int64) {
	o.Size = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CreateServerPayloadBootVolume) GetSource() *BootVolumeSource {
	if o == nil || IsNil(o.Source) {
		var ret *BootVolumeSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerPayloadBootVolume) GetSourceOk() (*BootVolumeSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CreateServerPayloadBootVolume) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given BootVolumeSource and assigns it to the Source field.
func (o *CreateServerPayloadBootVolume) SetSource(v *BootVolumeSource) {
	o.Source = v
}

func (o CreateServerPayloadBootVolume) ToMap() (map[string]interface{}, error) {
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

type NullableCreateServerPayloadBootVolume struct {
	value *CreateServerPayloadBootVolume
	isSet bool
}

func (v NullableCreateServerPayloadBootVolume) Get() *CreateServerPayloadBootVolume {
	return v.value
}

func (v *NullableCreateServerPayloadBootVolume) Set(val *CreateServerPayloadBootVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServerPayloadBootVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServerPayloadBootVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServerPayloadBootVolume(val *CreateServerPayloadBootVolume) *NullableCreateServerPayloadBootVolume {
	return &NullableCreateServerPayloadBootVolume{value: val, isSet: true}
}

func (v NullableCreateServerPayloadBootVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServerPayloadBootVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
