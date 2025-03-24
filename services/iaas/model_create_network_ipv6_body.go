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

// checks if the CreateNetworkIPv6Body type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkIPv6Body{}

/*
	types and functions for gateway
*/

// isNullableString
type CreateNetworkIPv6BodyGetGatewayAttributeType = *NullableString

func getCreateNetworkIPv6BodyGetGatewayAttributeTypeOk(arg CreateNetworkIPv6BodyGetGatewayAttributeType) (ret CreateNetworkIPv6BodyGetGatewayRetType, ok bool) {
	if arg == nil {
		return nil, false
	}
	return arg.Get(), true
}

func setCreateNetworkIPv6BodyGetGatewayAttributeType(arg *CreateNetworkIPv6BodyGetGatewayAttributeType, val CreateNetworkIPv6BodyGetGatewayRetType) {
	if IsNil(*arg) {
		*arg = NewNullableString(val)
	} else {
		(*arg).Set(val)
	}
}

type CreateNetworkIPv6BodyGetGatewayArgType = *string
type CreateNetworkIPv6BodyGetGatewayRetType = *string

/*
	types and functions for nameservers
*/

// isArray
type CreateNetworkIPv6BodyGetNameserversAttributeType = *[]string
type CreateNetworkIPv6BodyGetNameserversArgType = []string
type CreateNetworkIPv6BodyGetNameserversRetType = []string

func getCreateNetworkIPv6BodyGetNameserversAttributeTypeOk(arg CreateNetworkIPv6BodyGetNameserversAttributeType) (ret CreateNetworkIPv6BodyGetNameserversRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateNetworkIPv6BodyGetNameserversAttributeType(arg *CreateNetworkIPv6BodyGetNameserversAttributeType, val CreateNetworkIPv6BodyGetNameserversRetType) {
	*arg = &val
}

/*
	types and functions for prefix
*/

// isNotNullableString
type CreateNetworkIPv6BodyGetPrefixAttributeType = *string

func getCreateNetworkIPv6BodyGetPrefixAttributeTypeOk(arg CreateNetworkIPv6BodyGetPrefixAttributeType) (ret CreateNetworkIPv6BodyGetPrefixRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateNetworkIPv6BodyGetPrefixAttributeType(arg *CreateNetworkIPv6BodyGetPrefixAttributeType, val CreateNetworkIPv6BodyGetPrefixRetType) {
	*arg = &val
}

type CreateNetworkIPv6BodyGetPrefixArgType = string
type CreateNetworkIPv6BodyGetPrefixRetType = string

/*
	types and functions for prefixLength
*/

// isInteger
type CreateNetworkIPv6BodyGetPrefixLengthAttributeType = *int64
type CreateNetworkIPv6BodyGetPrefixLengthArgType = int64
type CreateNetworkIPv6BodyGetPrefixLengthRetType = int64

func getCreateNetworkIPv6BodyGetPrefixLengthAttributeTypeOk(arg CreateNetworkIPv6BodyGetPrefixLengthAttributeType) (ret CreateNetworkIPv6BodyGetPrefixLengthRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateNetworkIPv6BodyGetPrefixLengthAttributeType(arg *CreateNetworkIPv6BodyGetPrefixLengthAttributeType, val CreateNetworkIPv6BodyGetPrefixLengthRetType) {
	*arg = &val
}

// CreateNetworkIPv6Body The config object for an IPv6 network.
type CreateNetworkIPv6Body struct {
	// The gateway of a network. If not specified the first IP of the network will be assigned as the gateway. If 'null' is sent, then the network doesn't have a gateway.
	Gateway CreateNetworkIPv6BodyGetGatewayAttributeType `json:"gateway,omitempty"`
	// A list containing DNS Servers/Nameservers for IPv6.
	Nameservers CreateNetworkIPv6BodyGetNameserversAttributeType `json:"nameservers,omitempty"`
	// Classless Inter-Domain Routing (CIDR) for IPv6.
	Prefix       CreateNetworkIPv6BodyGetPrefixAttributeType       `json:"prefix,omitempty"`
	PrefixLength CreateNetworkIPv6BodyGetPrefixLengthAttributeType `json:"prefixLength,omitempty"`
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
func (o *CreateNetworkIPv6Body) GetGateway() (res CreateNetworkIPv6BodyGetGatewayRetType) {
	res, _ = o.GetGatewayOk()
	return
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateNetworkIPv6Body) GetGatewayOk() (ret CreateNetworkIPv6BodyGetGatewayRetType, ok bool) {
	return getCreateNetworkIPv6BodyGetGatewayAttributeTypeOk(o.Gateway)
}

// HasGateway returns a boolean if a field has been set.
func (o *CreateNetworkIPv6Body) HasGateway() bool {
	_, ok := o.GetGatewayOk()
	return ok
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *CreateNetworkIPv6Body) SetGateway(v CreateNetworkIPv6BodyGetGatewayRetType) {
	setCreateNetworkIPv6BodyGetGatewayAttributeType(&o.Gateway, v)
}

// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *CreateNetworkIPv6Body) SetGatewayNil() {
	o.Gateway = nil
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *CreateNetworkIPv6Body) UnsetGateway() {
	o.Gateway = nil
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *CreateNetworkIPv6Body) GetNameservers() (res CreateNetworkIPv6BodyGetNameserversRetType) {
	res, _ = o.GetNameserversOk()
	return
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkIPv6Body) GetNameserversOk() (ret CreateNetworkIPv6BodyGetNameserversRetType, ok bool) {
	return getCreateNetworkIPv6BodyGetNameserversAttributeTypeOk(o.Nameservers)
}

// HasNameservers returns a boolean if a field has been set.
func (o *CreateNetworkIPv6Body) HasNameservers() bool {
	_, ok := o.GetNameserversOk()
	return ok
}

// SetNameservers gets a reference to the given []string and assigns it to the Nameservers field.
func (o *CreateNetworkIPv6Body) SetNameservers(v CreateNetworkIPv6BodyGetNameserversRetType) {
	setCreateNetworkIPv6BodyGetNameserversAttributeType(&o.Nameservers, v)
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *CreateNetworkIPv6Body) GetPrefix() (res CreateNetworkIPv6BodyGetPrefixRetType) {
	res, _ = o.GetPrefixOk()
	return
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkIPv6Body) GetPrefixOk() (ret CreateNetworkIPv6BodyGetPrefixRetType, ok bool) {
	return getCreateNetworkIPv6BodyGetPrefixAttributeTypeOk(o.Prefix)
}

// HasPrefix returns a boolean if a field has been set.
func (o *CreateNetworkIPv6Body) HasPrefix() bool {
	_, ok := o.GetPrefixOk()
	return ok
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *CreateNetworkIPv6Body) SetPrefix(v CreateNetworkIPv6BodyGetPrefixRetType) {
	setCreateNetworkIPv6BodyGetPrefixAttributeType(&o.Prefix, v)
}

// GetPrefixLength returns the PrefixLength field value if set, zero value otherwise.
func (o *CreateNetworkIPv6Body) GetPrefixLength() (res CreateNetworkIPv6BodyGetPrefixLengthRetType) {
	res, _ = o.GetPrefixLengthOk()
	return
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkIPv6Body) GetPrefixLengthOk() (ret CreateNetworkIPv6BodyGetPrefixLengthRetType, ok bool) {
	return getCreateNetworkIPv6BodyGetPrefixLengthAttributeTypeOk(o.PrefixLength)
}

// HasPrefixLength returns a boolean if a field has been set.
func (o *CreateNetworkIPv6Body) HasPrefixLength() bool {
	_, ok := o.GetPrefixLengthOk()
	return ok
}

// SetPrefixLength gets a reference to the given int64 and assigns it to the PrefixLength field.
func (o *CreateNetworkIPv6Body) SetPrefixLength(v CreateNetworkIPv6BodyGetPrefixLengthRetType) {
	setCreateNetworkIPv6BodyGetPrefixLengthAttributeType(&o.PrefixLength, v)
}

func (o CreateNetworkIPv6Body) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateNetworkIPv6BodyGetGatewayAttributeTypeOk(o.Gateway); ok {
		toSerialize["Gateway"] = val
	}
	if val, ok := getCreateNetworkIPv6BodyGetNameserversAttributeTypeOk(o.Nameservers); ok {
		toSerialize["Nameservers"] = val
	}
	if val, ok := getCreateNetworkIPv6BodyGetPrefixAttributeTypeOk(o.Prefix); ok {
		toSerialize["Prefix"] = val
	}
	if val, ok := getCreateNetworkIPv6BodyGetPrefixLengthAttributeTypeOk(o.PrefixLength); ok {
		toSerialize["PrefixLength"] = val
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
