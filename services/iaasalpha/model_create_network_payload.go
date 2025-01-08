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

// checks if the CreateNetworkPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkPayload{}

// CreateNetworkPayload Object that represents the request body for a network create.
type CreateNetworkPayload struct {
	AddressFamily *CreateNetworkAddressFamily `json:"addressFamily,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	// REQUIRED
	Name *string `json:"name"`
	// Shows if the network is routed and therefore accessible from other networks.
	Routed *bool `json:"routed,omitempty"`
}

type _CreateNetworkPayload CreateNetworkPayload

// NewCreateNetworkPayload instantiates a new CreateNetworkPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkPayload(name *string) *CreateNetworkPayload {
	this := CreateNetworkPayload{}
	this.Name = name
	return &this
}

// NewCreateNetworkPayloadWithDefaults instantiates a new CreateNetworkPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkPayloadWithDefaults() *CreateNetworkPayload {
	this := CreateNetworkPayload{}
	return &this
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *CreateNetworkPayload) GetAddressFamily() *CreateNetworkAddressFamily {
	if o == nil || IsNil(o.AddressFamily) {
		var ret *CreateNetworkAddressFamily
		return ret
	}
	return o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkPayload) GetAddressFamilyOk() (*CreateNetworkAddressFamily, bool) {
	if o == nil || IsNil(o.AddressFamily) {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *CreateNetworkPayload) HasAddressFamily() bool {
	if o != nil && !IsNil(o.AddressFamily) {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given CreateNetworkAddressFamily and assigns it to the AddressFamily field.
func (o *CreateNetworkPayload) SetAddressFamily(v *CreateNetworkAddressFamily) {
	o.AddressFamily = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CreateNetworkPayload) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkPayload) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CreateNetworkPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *CreateNetworkPayload) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value
func (o *CreateNetworkPayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *CreateNetworkPayload) SetName(v *string) {
	o.Name = v
}

// GetRouted returns the Routed field value if set, zero value otherwise.
func (o *CreateNetworkPayload) GetRouted() *bool {
	if o == nil || IsNil(o.Routed) {
		var ret *bool
		return ret
	}
	return o.Routed
}

// GetRoutedOk returns a tuple with the Routed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkPayload) GetRoutedOk() (*bool, bool) {
	if o == nil || IsNil(o.Routed) {
		return nil, false
	}
	return o.Routed, true
}

// HasRouted returns a boolean if a field has been set.
func (o *CreateNetworkPayload) HasRouted() bool {
	if o != nil && !IsNil(o.Routed) {
		return true
	}

	return false
}

// SetRouted gets a reference to the given bool and assigns it to the Routed field.
func (o *CreateNetworkPayload) SetRouted(v *bool) {
	o.Routed = v
}

func (o CreateNetworkPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressFamily) {
		toSerialize["addressFamily"] = o.AddressFamily
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Routed) {
		toSerialize["routed"] = o.Routed
	}
	return toSerialize, nil
}

type NullableCreateNetworkPayload struct {
	value *CreateNetworkPayload
	isSet bool
}

func (v NullableCreateNetworkPayload) Get() *CreateNetworkPayload {
	return v.value
}

func (v *NullableCreateNetworkPayload) Set(val *CreateNetworkPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkPayload(val *CreateNetworkPayload) *NullableCreateNetworkPayload {
	return &NullableCreateNetworkPayload{value: val, isSet: true}
}

func (v NullableCreateNetworkPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}