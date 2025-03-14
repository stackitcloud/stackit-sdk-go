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

/*
	types and functions for gateway
*/

// isNullableString
type CreateNetworkIPv4BodyGetGatewayAttributeType = *NullableString

func getCreateNetworkIPv4BodyGetGatewayAttributeTypeOk(arg CreateNetworkIPv4BodyGetGatewayAttributeType) (ret CreateNetworkIPv4BodyGetGatewayRetType, ok bool) {
	if arg == nil {
		return nil, false
	}
	return arg.Get(), true
}

func setCreateNetworkIPv4BodyGetGatewayAttributeType(arg *CreateNetworkIPv4BodyGetGatewayAttributeType, val CreateNetworkIPv4BodyGetGatewayRetType) {
	if IsNil(*arg) {
		*arg = NewNullableString(val)
	} else {
		(*arg).Set(val)
	}
}

type CreateNetworkIPv4BodyGetGatewayArgType = *string
type CreateNetworkIPv4BodyGetGatewayRetType = *string

/*
	types and functions for nameservers
*/

// isArray
type CreateNetworkIPv4BodyGetNameserversAttributeType = *[]string
type CreateNetworkIPv4BodyGetNameserversArgType = []string
type CreateNetworkIPv4BodyGetNameserversRetType = []string

func getCreateNetworkIPv4BodyGetNameserversAttributeTypeOk(arg CreateNetworkIPv4BodyGetNameserversAttributeType) (ret CreateNetworkIPv4BodyGetNameserversRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateNetworkIPv4BodyGetNameserversAttributeType(arg *CreateNetworkIPv4BodyGetNameserversAttributeType, val CreateNetworkIPv4BodyGetNameserversRetType) {
	*arg = &val
}

/*
	types and functions for prefix
*/

// isNotNullableString
type CreateNetworkIPv4BodyGetPrefixAttributeType = *string

func getCreateNetworkIPv4BodyGetPrefixAttributeTypeOk(arg CreateNetworkIPv4BodyGetPrefixAttributeType) (ret CreateNetworkIPv4BodyGetPrefixRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateNetworkIPv4BodyGetPrefixAttributeType(arg *CreateNetworkIPv4BodyGetPrefixAttributeType, val CreateNetworkIPv4BodyGetPrefixRetType) {
	*arg = &val
}

type CreateNetworkIPv4BodyGetPrefixArgType = string
type CreateNetworkIPv4BodyGetPrefixRetType = string

/*
	types and functions for prefixLength
*/

// isInteger
type CreateNetworkIPv4BodyGetPrefixLengthAttributeType = *int64
type CreateNetworkIPv4BodyGetPrefixLengthArgType = int64
type CreateNetworkIPv4BodyGetPrefixLengthRetType = int64

func getCreateNetworkIPv4BodyGetPrefixLengthAttributeTypeOk(arg CreateNetworkIPv4BodyGetPrefixLengthAttributeType) (ret CreateNetworkIPv4BodyGetPrefixLengthRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateNetworkIPv4BodyGetPrefixLengthAttributeType(arg *CreateNetworkIPv4BodyGetPrefixLengthAttributeType, val CreateNetworkIPv4BodyGetPrefixLengthRetType) {
	*arg = &val
}

// CreateNetworkIPv4Body The config object for an IPv4 network.
type CreateNetworkIPv4Body struct {
	// The gateway of a network. If not specified the first IP of the network will be assigned as the gateway. If 'null' is sent, then the network doesn't have a gateway.
	Gateway CreateNetworkIPv4BodyGetGatewayAttributeType `json:"gateway,omitempty"`
	// A list containing DNS Servers/Nameservers for IPv4.
	Nameservers CreateNetworkIPv4BodyGetNameserversAttributeType `json:"nameservers,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	Prefix       CreateNetworkIPv4BodyGetPrefixAttributeType       `json:"prefix,omitempty"`
	PrefixLength CreateNetworkIPv4BodyGetPrefixLengthAttributeType `json:"prefixLength,omitempty"`
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
func (o *CreateNetworkIPv4Body) GetGateway() (res CreateNetworkIPv4BodyGetGatewayRetType) {
	res, _ = o.GetGatewayOk()
	return
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateNetworkIPv4Body) GetGatewayOk() (ret CreateNetworkIPv4BodyGetGatewayRetType, ok bool) {
	return getCreateNetworkIPv4BodyGetGatewayAttributeTypeOk(o.Gateway)
}

// HasGateway returns a boolean if a field has been set.
func (o *CreateNetworkIPv4Body) HasGateway() bool {
	_, ok := o.GetGatewayOk()
	return ok
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *CreateNetworkIPv4Body) SetGateway(v CreateNetworkIPv4BodyGetGatewayRetType) {
	setCreateNetworkIPv4BodyGetGatewayAttributeType(&o.Gateway, v)
}

// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *CreateNetworkIPv4Body) SetGatewayNil() {
	o.Gateway = nil
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *CreateNetworkIPv4Body) UnsetGateway() {
	o.Gateway = nil
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *CreateNetworkIPv4Body) GetNameservers() (res CreateNetworkIPv4BodyGetNameserversRetType) {
	res, _ = o.GetNameserversOk()
	return
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkIPv4Body) GetNameserversOk() (ret CreateNetworkIPv4BodyGetNameserversRetType, ok bool) {
	return getCreateNetworkIPv4BodyGetNameserversAttributeTypeOk(o.Nameservers)
}

// HasNameservers returns a boolean if a field has been set.
func (o *CreateNetworkIPv4Body) HasNameservers() bool {
	_, ok := o.GetNameserversOk()
	return ok
}

// SetNameservers gets a reference to the given []string and assigns it to the Nameservers field.
func (o *CreateNetworkIPv4Body) SetNameservers(v CreateNetworkIPv4BodyGetNameserversRetType) {
	setCreateNetworkIPv4BodyGetNameserversAttributeType(&o.Nameservers, v)
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *CreateNetworkIPv4Body) GetPrefix() (res CreateNetworkIPv4BodyGetPrefixRetType) {
	res, _ = o.GetPrefixOk()
	return
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkIPv4Body) GetPrefixOk() (ret CreateNetworkIPv4BodyGetPrefixRetType, ok bool) {
	return getCreateNetworkIPv4BodyGetPrefixAttributeTypeOk(o.Prefix)
}

// HasPrefix returns a boolean if a field has been set.
func (o *CreateNetworkIPv4Body) HasPrefix() bool {
	_, ok := o.GetPrefixOk()
	return ok
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *CreateNetworkIPv4Body) SetPrefix(v CreateNetworkIPv4BodyGetPrefixRetType) {
	setCreateNetworkIPv4BodyGetPrefixAttributeType(&o.Prefix, v)
}

// GetPrefixLength returns the PrefixLength field value if set, zero value otherwise.
func (o *CreateNetworkIPv4Body) GetPrefixLength() (res CreateNetworkIPv4BodyGetPrefixLengthRetType) {
	res, _ = o.GetPrefixLengthOk()
	return
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkIPv4Body) GetPrefixLengthOk() (ret CreateNetworkIPv4BodyGetPrefixLengthRetType, ok bool) {
	return getCreateNetworkIPv4BodyGetPrefixLengthAttributeTypeOk(o.PrefixLength)
}

// HasPrefixLength returns a boolean if a field has been set.
func (o *CreateNetworkIPv4Body) HasPrefixLength() bool {
	_, ok := o.GetPrefixLengthOk()
	return ok
}

// SetPrefixLength gets a reference to the given int64 and assigns it to the PrefixLength field.
func (o *CreateNetworkIPv4Body) SetPrefixLength(v CreateNetworkIPv4BodyGetPrefixLengthRetType) {
	setCreateNetworkIPv4BodyGetPrefixLengthAttributeType(&o.PrefixLength, v)
}

func (o CreateNetworkIPv4Body) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateNetworkIPv4BodyGetGatewayAttributeTypeOk(o.Gateway); ok {
		toSerialize["Gateway"] = val
	}
	if val, ok := getCreateNetworkIPv4BodyGetNameserversAttributeTypeOk(o.Nameservers); ok {
		toSerialize["Nameservers"] = val
	}
	if val, ok := getCreateNetworkIPv4BodyGetPrefixAttributeTypeOk(o.Prefix); ok {
		toSerialize["Prefix"] = val
	}
	if val, ok := getCreateNetworkIPv4BodyGetPrefixLengthAttributeTypeOk(o.PrefixLength); ok {
		toSerialize["PrefixLength"] = val
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
