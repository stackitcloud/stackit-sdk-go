/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the CreateZonePayloadExtensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateZonePayloadExtensions{}

// CreateZonePayloadExtensions optional extensions
type CreateZonePayloadExtensions struct {
	ObservabilityExtension *ZoneObservabilityExtension `json:"observabilityExtension,omitempty"`
}

// NewCreateZonePayloadExtensions instantiates a new CreateZonePayloadExtensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateZonePayloadExtensions() *CreateZonePayloadExtensions {
	this := CreateZonePayloadExtensions{}
	return &this
}

// NewCreateZonePayloadExtensionsWithDefaults instantiates a new CreateZonePayloadExtensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateZonePayloadExtensionsWithDefaults() *CreateZonePayloadExtensions {
	this := CreateZonePayloadExtensions{}
	return &this
}

// GetObservabilityExtension returns the ObservabilityExtension field value if set, zero value otherwise.
func (o *CreateZonePayloadExtensions) GetObservabilityExtension() *ZoneObservabilityExtension {
	if o == nil || IsNil(o.ObservabilityExtension) {
		var ret *ZoneObservabilityExtension
		return ret
	}
	return o.ObservabilityExtension
}

// GetObservabilityExtensionOk returns a tuple with the ObservabilityExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZonePayloadExtensions) GetObservabilityExtensionOk() (*ZoneObservabilityExtension, bool) {
	if o == nil || IsNil(o.ObservabilityExtension) {
		return nil, false
	}
	return o.ObservabilityExtension, true
}

// HasObservabilityExtension returns a boolean if a field has been set.
func (o *CreateZonePayloadExtensions) HasObservabilityExtension() bool {
	if o != nil && !IsNil(o.ObservabilityExtension) {
		return true
	}

	return false
}

// SetObservabilityExtension gets a reference to the given ZoneObservabilityExtension and assigns it to the ObservabilityExtension field.
func (o *CreateZonePayloadExtensions) SetObservabilityExtension(v *ZoneObservabilityExtension) {
	o.ObservabilityExtension = v
}

func (o CreateZonePayloadExtensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObservabilityExtension) {
		toSerialize["observabilityExtension"] = o.ObservabilityExtension
	}
	return toSerialize, nil
}

type NullableCreateZonePayloadExtensions struct {
	value *CreateZonePayloadExtensions
	isSet bool
}

func (v NullableCreateZonePayloadExtensions) Get() *CreateZonePayloadExtensions {
	return v.value
}

func (v *NullableCreateZonePayloadExtensions) Set(val *CreateZonePayloadExtensions) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateZonePayloadExtensions) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateZonePayloadExtensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateZonePayloadExtensions(val *CreateZonePayloadExtensions) *NullableCreateZonePayloadExtensions {
	return &NullableCreateZonePayloadExtensions{value: val, isSet: true}
}

func (v NullableCreateZonePayloadExtensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateZonePayloadExtensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
