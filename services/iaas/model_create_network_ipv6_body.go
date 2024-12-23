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

// checks if the CreateNetworkIPv6Body type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkIPv6Body{}

// CreateNetworkIPv6Body The config object for an IPv6 network.
type CreateNetworkIPv6Body struct {
	// The gateway of a network. If not specified the first IP of the network will be assigned as the gateway. If 'null' is sent, then the network doesn't have a gateway.
	Gateway *NullableString `json:"gateway,omitempty"`
	// A list containing DNS Servers/Nameservers for IPv6.
	Nameservers *[]string `json:"nameservers,omitempty"`
	// Classless Inter-Domain Routing (CIDR) for IPv6.
	Prefix       *string `json:"prefix,omitempty"`
	PrefixLength *int64  `json:"prefixLength,omitempty"`
}

// NewCreateNetworkIPv6Body instantiates a new CreateNetworkIPv6Body object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkIPv6Body() *CreateNetworkIPv6Body {
	this := CreateNetworkIPv6Body{}
	return &this
}

// NewCreateNetworkIPv6BodyWithDefaults instantiates a new CreateNetworkIPv6Body object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkIPv6BodyWithDefaults() *CreateNetworkIPv6Body {
	this := CreateNetworkIPv6Body{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateNetworkIPv6Body) GetGateway() *string {
	if o == nil || IsNil(o.Gateway) || IsNil(o.Gateway.Get()) {
		var ret *string
		return ret
	}
	return o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateNetworkIPv6Body) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *CreateNetworkIPv6Body) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *CreateNetworkIPv6Body) SetGateway(v *string) {
	if IsNil(o.Gateway) {
		o.Gateway = new(NullableString)
	}
	o.Gateway.Set(v)
}

// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *CreateNetworkIPv6Body) SetGatewayNil() {
	if IsNil(o.Gateway) {
		o.Gateway = new(NullableString)
	}
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *CreateNetworkIPv6Body) UnsetGateway() {
	if IsNil(o.Gateway) {
		o.Gateway = new(NullableString)
	}
	o.Gateway.Unset()
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *CreateNetworkIPv6Body) GetNameservers() *[]string {
	if o == nil || IsNil(o.Nameservers) {
		var ret *[]string
		return ret
	}
	return o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkIPv6Body) GetNameserversOk() (*[]string, bool) {
	if o == nil || IsNil(o.Nameservers) {
		return nil, false
	}
	return o.Nameservers, true
}

// HasNameservers returns a boolean if a field has been set.
func (o *CreateNetworkIPv6Body) HasNameservers() bool {
	if o != nil && !IsNil(o.Nameservers) {
		return true
	}

	return false
}

// SetNameservers gets a reference to the given []string and assigns it to the Nameservers field.
func (o *CreateNetworkIPv6Body) SetNameservers(v *[]string) {
	o.Nameservers = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *CreateNetworkIPv6Body) GetPrefix() *string {
	if o == nil || IsNil(o.Prefix) {
		var ret *string
		return ret
	}
	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkIPv6Body) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *CreateNetworkIPv6Body) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *CreateNetworkIPv6Body) SetPrefix(v *string) {
	o.Prefix = v
}

// GetPrefixLength returns the PrefixLength field value if set, zero value otherwise.
func (o *CreateNetworkIPv6Body) GetPrefixLength() *int64 {
	if o == nil || IsNil(o.PrefixLength) {
		var ret *int64
		return ret
	}
	return o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkIPv6Body) GetPrefixLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.PrefixLength) {
		return nil, false
	}
	return o.PrefixLength, true
}

// HasPrefixLength returns a boolean if a field has been set.
func (o *CreateNetworkIPv6Body) HasPrefixLength() bool {
	if o != nil && !IsNil(o.PrefixLength) {
		return true
	}

	return false
}

// SetPrefixLength gets a reference to the given int64 and assigns it to the PrefixLength field.
func (o *CreateNetworkIPv6Body) SetPrefixLength(v *int64) {
	o.PrefixLength = v
}

func (o CreateNetworkIPv6Body) ToMap() (map[string]interface{}, error) {
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

type NullableCreateNetworkIPv6Body struct {
	value *CreateNetworkIPv6Body
	isSet bool
}

func (v NullableCreateNetworkIPv6Body) Get() *CreateNetworkIPv6Body {
	return v.value
}

func (v *NullableCreateNetworkIPv6Body) Set(val *CreateNetworkIPv6Body) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkIPv6Body) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkIPv6Body) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkIPv6Body(val *CreateNetworkIPv6Body) *NullableCreateNetworkIPv6Body {
	return &NullableCreateNetworkIPv6Body{value: val, isSet: true}
}

func (v NullableCreateNetworkIPv6Body) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkIPv6Body) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
