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

// checks if the UpdateVolumePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVolumePayload{}

// UpdateVolumePayload Object that represents an update request body of a  volume.
type UpdateVolumePayload struct {
	// Description Object. Allows string up to 127 Characters.
	Description *string `json:"description,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name *string `json:"name,omitempty"`
}

// NewUpdateVolumePayload instantiates a new UpdateVolumePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVolumePayload() *UpdateVolumePayload {
	this := UpdateVolumePayload{}
	return &this
}

// NewUpdateVolumePayloadWithDefaults instantiates a new UpdateVolumePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVolumePayloadWithDefaults() *UpdateVolumePayload {
	this := UpdateVolumePayload{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateVolumePayload) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumePayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateVolumePayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateVolumePayload) SetDescription(v *string) {
	o.Description = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateVolumePayload) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumePayload) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *UpdateVolumePayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *UpdateVolumePayload) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateVolumePayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumePayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateVolumePayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateVolumePayload) SetName(v *string) {
	o.Name = v
}

func (o UpdateVolumePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUpdateVolumePayload struct {
	value *UpdateVolumePayload
	isSet bool
}

func (v NullableUpdateVolumePayload) Get() *UpdateVolumePayload {
	return v.value
}

func (v *NullableUpdateVolumePayload) Set(val *UpdateVolumePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVolumePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVolumePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVolumePayload(val *UpdateVolumePayload) *NullableUpdateVolumePayload {
	return &NullableUpdateVolumePayload{value: val, isSet: true}
}

func (v NullableUpdateVolumePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVolumePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
