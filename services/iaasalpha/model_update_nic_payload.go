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

// checks if the UpdateNicPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNicPayload{}

// UpdateNicPayload Object that represents a network interface update.
type UpdateNicPayload struct {
	// A list of IPs or CIDR notations.
	AllowedAddresses *[]AllowedAddressesInner `json:"allowedAddresses,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name *string `json:"name,omitempty"`
	// If this is set to false, then no security groups will apply to this network interface.
	NicSecurity *bool `json:"nicSecurity,omitempty"`
	// A list of UUIDs.
	SecurityGroups *[]string `json:"securityGroups,omitempty"`
}

// NewUpdateNicPayload instantiates a new UpdateNicPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNicPayload() *UpdateNicPayload {
	this := UpdateNicPayload{}
	return &this
}

// NewUpdateNicPayloadWithDefaults instantiates a new UpdateNicPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNicPayloadWithDefaults() *UpdateNicPayload {
	this := UpdateNicPayload{}
	return &this
}

// GetAllowedAddresses returns the AllowedAddresses field value if set, zero value otherwise.
func (o *UpdateNicPayload) GetAllowedAddresses() *[]AllowedAddressesInner {
	if o == nil || IsNil(o.AllowedAddresses) {
		var ret *[]AllowedAddressesInner
		return ret
	}
	return o.AllowedAddresses
}

// GetAllowedAddressesOk returns a tuple with the AllowedAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNicPayload) GetAllowedAddressesOk() (*[]AllowedAddressesInner, bool) {
	if o == nil || IsNil(o.AllowedAddresses) {
		return nil, false
	}
	return o.AllowedAddresses, true
}

// HasAllowedAddresses returns a boolean if a field has been set.
func (o *UpdateNicPayload) HasAllowedAddresses() bool {
	if o != nil && !IsNil(o.AllowedAddresses) {
		return true
	}

	return false
}

// SetAllowedAddresses gets a reference to the given []AllowedAddressesInner and assigns it to the AllowedAddresses field.
func (o *UpdateNicPayload) SetAllowedAddresses(v *[]AllowedAddressesInner) {
	o.AllowedAddresses = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateNicPayload) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNicPayload) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *UpdateNicPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *UpdateNicPayload) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNicPayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNicPayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNicPayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNicPayload) SetName(v *string) {
	o.Name = v
}

// GetNicSecurity returns the NicSecurity field value if set, zero value otherwise.
func (o *UpdateNicPayload) GetNicSecurity() *bool {
	if o == nil || IsNil(o.NicSecurity) {
		var ret *bool
		return ret
	}
	return o.NicSecurity
}

// GetNicSecurityOk returns a tuple with the NicSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNicPayload) GetNicSecurityOk() (*bool, bool) {
	if o == nil || IsNil(o.NicSecurity) {
		return nil, false
	}
	return o.NicSecurity, true
}

// HasNicSecurity returns a boolean if a field has been set.
func (o *UpdateNicPayload) HasNicSecurity() bool {
	if o != nil && !IsNil(o.NicSecurity) {
		return true
	}

	return false
}

// SetNicSecurity gets a reference to the given bool and assigns it to the NicSecurity field.
func (o *UpdateNicPayload) SetNicSecurity(v *bool) {
	o.NicSecurity = v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *UpdateNicPayload) GetSecurityGroups() *[]string {
	if o == nil || IsNil(o.SecurityGroups) {
		var ret *[]string
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNicPayload) GetSecurityGroupsOk() (*[]string, bool) {
	if o == nil || IsNil(o.SecurityGroups) {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *UpdateNicPayload) HasSecurityGroups() bool {
	if o != nil && !IsNil(o.SecurityGroups) {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *UpdateNicPayload) SetSecurityGroups(v *[]string) {
	o.SecurityGroups = v
}

func (o UpdateNicPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedAddresses) {
		toSerialize["allowedAddresses"] = o.AllowedAddresses
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NicSecurity) {
		toSerialize["nicSecurity"] = o.NicSecurity
	}
	if !IsNil(o.SecurityGroups) {
		toSerialize["securityGroups"] = o.SecurityGroups
	}
	return toSerialize, nil
}

type NullableUpdateNicPayload struct {
	value *UpdateNicPayload
	isSet bool
}

func (v NullableUpdateNicPayload) Get() *UpdateNicPayload {
	return v.value
}

func (v *NullableUpdateNicPayload) Set(val *UpdateNicPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNicPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNicPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNicPayload(val *UpdateNicPayload) *NullableUpdateNicPayload {
	return &NullableUpdateNicPayload{value: val, isSet: true}
}

func (v NullableUpdateNicPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNicPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
