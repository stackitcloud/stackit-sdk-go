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

// checks if the CreateNetworkIPv4Body type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkIPv4Body{}

// CreateNetworkIPv4Body The config object for an IPv4 network.
type CreateNetworkIPv4Body struct {
	Gateway *NullableV1NetworkGateway `json:"gateway,omitempty"`
	// A list containing DNS Servers/Nameservers for IPv4.
	Nameservers *[]string `json:"nameservers,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	Prefix       *string `json:"prefix,omitempty"`
	PrefixLength *int64  `json:"prefixLength,omitempty"`
}

// NewCreateNetworkIPv4Body instantiates a new CreateNetworkIPv4Body object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkIPv4Body() *CreateNetworkIPv4Body {
	this := CreateNetworkIPv4Body{}
	return &this
}

// NewCreateNetworkIPv4BodyWithDefaults instantiates a new CreateNetworkIPv4Body object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkIPv4BodyWithDefaults() *CreateNetworkIPv4Body {
	this := CreateNetworkIPv4Body{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateNetworkIPv4Body) GetGateway() *V1NetworkGateway {
	if o == nil || IsNil(o.Gateway.Get()) {
		var ret *V1NetworkGateway
		return ret
	}
	return o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateNetworkIPv4Body) GetGatewayOk() (*V1NetworkGateway, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *CreateNetworkIPv4Body) HasGateway() bool {
	if o != nil && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given V1NetworkGateway and assigns it to the Gateway field.
func (o *CreateNetworkIPv4Body) SetGateway(v *V1NetworkGateway) {
	o.Gateway.Set(v)
}

// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *CreateNetworkIPv4Body) SetGatewayNil() {
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *CreateNetworkIPv4Body) UnsetGateway() {
	o.Gateway.Unset()
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *CreateNetworkIPv4Body) GetNameservers() *[]string {
	if o == nil || IsNil(o.Nameservers) {
		var ret *[]string
		return ret
	}
	return o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkIPv4Body) GetNameserversOk() (*[]string, bool) {
	if o == nil || IsNil(o.Nameservers) {
		return nil, false
	}
	return o.Nameservers, true
}

// HasNameservers returns a boolean if a field has been set.
func (o *CreateNetworkIPv4Body) HasNameservers() bool {
	if o != nil && !IsNil(o.Nameservers) {
		return true
	}

	return false
}

// SetNameservers gets a reference to the given []string and assigns it to the Nameservers field.
func (o *CreateNetworkIPv4Body) SetNameservers(v *[]string) {
	o.Nameservers = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *CreateNetworkIPv4Body) GetPrefix() *string {
	if o == nil || IsNil(o.Prefix) {
		var ret *string
		return ret
	}
	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkIPv4Body) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *CreateNetworkIPv4Body) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *CreateNetworkIPv4Body) SetPrefix(v *string) {
	o.Prefix = v
}

// GetPrefixLength returns the PrefixLength field value if set, zero value otherwise.
func (o *CreateNetworkIPv4Body) GetPrefixLength() *int64 {
	if o == nil || IsNil(o.PrefixLength) {
		var ret *int64
		return ret
	}
	return o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkIPv4Body) GetPrefixLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.PrefixLength) {
		return nil, false
	}
	return o.PrefixLength, true
}

// HasPrefixLength returns a boolean if a field has been set.
func (o *CreateNetworkIPv4Body) HasPrefixLength() bool {
	if o != nil && !IsNil(o.PrefixLength) {
		return true
	}

	return false
}

// SetPrefixLength gets a reference to the given int64 and assigns it to the PrefixLength field.
func (o *CreateNetworkIPv4Body) SetPrefixLength(v *int64) {
	o.PrefixLength = v
}

func (o CreateNetworkIPv4Body) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Gateway.IsSet() {
		toSerialize["gateway"] = o.Gateway.Get()
	}
	if !IsNil(o.Nameservers) {
		toSerialize["nameservers"] = o.Nameservers
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.PrefixLength) {
		toSerialize["prefixLength"] = o.PrefixLength
	}
	return toSerialize, nil
}

type NullableCreateNetworkIPv4Body struct {
	value *CreateNetworkIPv4Body
	isSet bool
}

func (v NullableCreateNetworkIPv4Body) Get() *CreateNetworkIPv4Body {
	return v.value
}

func (v *NullableCreateNetworkIPv4Body) Set(val *CreateNetworkIPv4Body) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkIPv4Body) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkIPv4Body) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkIPv4Body(val *CreateNetworkIPv4Body) *NullableCreateNetworkIPv4Body {
	return &NullableCreateNetworkIPv4Body{value: val, isSet: true}
}

func (v NullableCreateNetworkIPv4Body) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkIPv4Body) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
