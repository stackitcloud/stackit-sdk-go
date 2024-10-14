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

// checks if the CloneZonePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloneZonePayload{}

// CloneZonePayload metadata to clone a zone.
type CloneZonePayload struct {
	// Adjust record set content and replace the dns name of the original zone with the new dns name of the cloned zone
	AdjustRecords *bool `json:"adjustRecords,omitempty"`
	// New Description for the cloned zone. Leave empty to use the same description as the original zone
	Description *string `json:"description,omitempty"`
	// DnsName is the dns name of the zone to clone
	// REQUIRED
	DnsName *string `json:"dnsName"`
	// New Name for the cloned zone. Leave empty to use the same name as the original zone
	Name *string `json:"name,omitempty"`
}

type _CloneZonePayload CloneZonePayload

// NewCloneZonePayload instantiates a new CloneZonePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloneZonePayload(dnsName *string) *CloneZonePayload {
	this := CloneZonePayload{}
	this.DnsName = dnsName
	return &this
}

// NewCloneZonePayloadWithDefaults instantiates a new CloneZonePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloneZonePayloadWithDefaults() *CloneZonePayload {
	this := CloneZonePayload{}
	return &this
}

// GetAdjustRecords returns the AdjustRecords field value if set, zero value otherwise.
func (o *CloneZonePayload) GetAdjustRecords() *bool {
	if o == nil || IsNil(o.AdjustRecords) {
		var ret *bool
		return ret
	}
	return o.AdjustRecords
}

// GetAdjustRecordsOk returns a tuple with the AdjustRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloneZonePayload) GetAdjustRecordsOk() (*bool, bool) {
	if o == nil || IsNil(o.AdjustRecords) {
		return nil, false
	}
	return o.AdjustRecords, true
}

// HasAdjustRecords returns a boolean if a field has been set.
func (o *CloneZonePayload) HasAdjustRecords() bool {
	if o != nil && !IsNil(o.AdjustRecords) {
		return true
	}

	return false
}

// SetAdjustRecords gets a reference to the given bool and assigns it to the AdjustRecords field.
func (o *CloneZonePayload) SetAdjustRecords(v *bool) {
	o.AdjustRecords = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloneZonePayload) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloneZonePayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloneZonePayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloneZonePayload) SetDescription(v *string) {
	o.Description = v
}

// GetDnsName returns the DnsName field value
func (o *CloneZonePayload) GetDnsName() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value
// and a boolean to check if the value has been set.
func (o *CloneZonePayload) GetDnsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsName, true
}

// SetDnsName sets field value
func (o *CloneZonePayload) SetDnsName(v *string) {
	o.DnsName = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloneZonePayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloneZonePayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloneZonePayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloneZonePayload) SetName(v *string) {
	o.Name = v
}

func (o CloneZonePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdjustRecords) {
		toSerialize["adjustRecords"] = o.AdjustRecords
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["dnsName"] = o.DnsName
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCloneZonePayload struct {
	value *CloneZonePayload
	isSet bool
}

func (v NullableCloneZonePayload) Get() *CloneZonePayload {
	return v.value
}

func (v *NullableCloneZonePayload) Set(val *CloneZonePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneZonePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneZonePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneZonePayload(val *CloneZonePayload) *NullableCloneZonePayload {
	return &NullableCloneZonePayload{value: val, isSet: true}
}

func (v NullableCloneZonePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneZonePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
