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

// checks if the CreateServerNetworkingWithNics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServerNetworkingWithNics{}

// CreateServerNetworkingWithNics The initial networking setup for the server creation with a network interface.
type CreateServerNetworkingWithNics struct {
	// A list of UUIDs.
	NicIds *[]string `json:"nicIds,omitempty"`
}

// NewCreateServerNetworkingWithNics instantiates a new CreateServerNetworkingWithNics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServerNetworkingWithNics() *CreateServerNetworkingWithNics {
	this := CreateServerNetworkingWithNics{}
	return &this
}

// NewCreateServerNetworkingWithNicsWithDefaults instantiates a new CreateServerNetworkingWithNics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServerNetworkingWithNicsWithDefaults() *CreateServerNetworkingWithNics {
	this := CreateServerNetworkingWithNics{}
	return &this
}

// GetNicIds returns the NicIds field value if set, zero value otherwise.
func (o *CreateServerNetworkingWithNics) GetNicIds() *[]string {
	if o == nil || IsNil(o.NicIds) {
		var ret *[]string
		return ret
	}
	return o.NicIds
}

// GetNicIdsOk returns a tuple with the NicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServerNetworkingWithNics) GetNicIdsOk() (*[]string, bool) {
	if o == nil || IsNil(o.NicIds) {
		return nil, false
	}
	return o.NicIds, true
}

// HasNicIds returns a boolean if a field has been set.
func (o *CreateServerNetworkingWithNics) HasNicIds() bool {
	if o != nil && !IsNil(o.NicIds) && !IsNil(o.NicIds) {
		return true
	}

	return false
}

// SetNicIds gets a reference to the given []string and assigns it to the NicIds field.
func (o *CreateServerNetworkingWithNics) SetNicIds(v *[]string) {
	o.NicIds = v
}

func (o CreateServerNetworkingWithNics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NicIds) {
		toSerialize["nicIds"] = o.NicIds
	}
	return toSerialize, nil
}

type NullableCreateServerNetworkingWithNics struct {
	value *CreateServerNetworkingWithNics
	isSet bool
}

func (v NullableCreateServerNetworkingWithNics) Get() *CreateServerNetworkingWithNics {
	return v.value
}

func (v *NullableCreateServerNetworkingWithNics) Set(val *CreateServerNetworkingWithNics) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServerNetworkingWithNics) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServerNetworkingWithNics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServerNetworkingWithNics(val *CreateServerNetworkingWithNics) *NullableCreateServerNetworkingWithNics {
	return &NullableCreateServerNetworkingWithNics{value: val, isSet: true}
}

func (v NullableCreateServerNetworkingWithNics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServerNetworkingWithNics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
