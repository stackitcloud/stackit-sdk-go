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

// checks if the AreaConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AreaConfig{}

/*
	types and functions for defaultNameservers
*/

// isArray
type AreaConfigGetDefaultNameserversAttributeType = *[]string
type AreaConfigGetDefaultNameserversArgType = []string
type AreaConfigGetDefaultNameserversRetType = []string

func getAreaConfigGetDefaultNameserversAttributeTypeOk(arg AreaConfigGetDefaultNameserversAttributeType) (ret AreaConfigGetDefaultNameserversRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAreaConfigGetDefaultNameserversAttributeType(arg *AreaConfigGetDefaultNameserversAttributeType, val AreaConfigGetDefaultNameserversRetType) {
	*arg = &val
}

/*
	types and functions for networkRanges
*/

// isArray
type AreaConfigGetNetworkRangesAttributeType = *[]NetworkRange
type AreaConfigGetNetworkRangesArgType = []NetworkRange
type AreaConfigGetNetworkRangesRetType = []NetworkRange

func getAreaConfigGetNetworkRangesAttributeTypeOk(arg AreaConfigGetNetworkRangesAttributeType) (ret AreaConfigGetNetworkRangesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAreaConfigGetNetworkRangesAttributeType(arg *AreaConfigGetNetworkRangesAttributeType, val AreaConfigGetNetworkRangesRetType) {
	*arg = &val
}

/*
	types and functions for routes
*/

// isArray
type AreaConfigGetRoutesAttributeType = *[]Route
type AreaConfigGetRoutesArgType = []Route
type AreaConfigGetRoutesRetType = []Route

func getAreaConfigGetRoutesAttributeTypeOk(arg AreaConfigGetRoutesAttributeType) (ret AreaConfigGetRoutesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAreaConfigGetRoutesAttributeType(arg *AreaConfigGetRoutesAttributeType, val AreaConfigGetRoutesRetType) {
	*arg = &val
}

/*
	types and functions for transferNetwork
*/

// isNotNullableString
type AreaConfigGetTransferNetworkAttributeType = *string

func getAreaConfigGetTransferNetworkAttributeTypeOk(arg AreaConfigGetTransferNetworkAttributeType) (ret AreaConfigGetTransferNetworkRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAreaConfigGetTransferNetworkAttributeType(arg *AreaConfigGetTransferNetworkAttributeType, val AreaConfigGetTransferNetworkRetType) {
	*arg = &val
}

type AreaConfigGetTransferNetworkArgType = string
type AreaConfigGetTransferNetworkRetType = string

// AreaConfig The basic network area object.
type AreaConfig struct {
	DefaultNameservers AreaConfigGetDefaultNameserversAttributeType `json:"defaultNameservers,omitempty"`
	// A list of network ranges.
	// REQUIRED
	NetworkRanges AreaConfigGetNetworkRangesAttributeType `json:"networkRanges"`
	// A list of routes.
	Routes AreaConfigGetRoutesAttributeType `json:"routes,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	// REQUIRED
	TransferNetwork AreaConfigGetTransferNetworkAttributeType `json:"transferNetwork"`
}

type _AreaConfig AreaConfig

// NewAreaConfig instantiates a new AreaConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAreaConfig(networkRanges AreaConfigGetNetworkRangesArgType, transferNetwork AreaConfigGetTransferNetworkArgType) *AreaConfig {
	this := AreaConfig{}
	setAreaConfigGetNetworkRangesAttributeType(&this.NetworkRanges, networkRanges)
	setAreaConfigGetTransferNetworkAttributeType(&this.TransferNetwork, transferNetwork)
	return &this
}

// NewAreaConfigWithDefaults instantiates a new AreaConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAreaConfigWithDefaults() *AreaConfig {
	this := AreaConfig{}
	return &this
}

// GetDefaultNameservers returns the DefaultNameservers field value if set, zero value otherwise.
func (o *AreaConfig) GetDefaultNameservers() (res AreaConfigGetDefaultNameserversRetType) {
	res, _ = o.GetDefaultNameserversOk()
	return
}

// GetDefaultNameserversOk returns a tuple with the DefaultNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaConfig) GetDefaultNameserversOk() (ret AreaConfigGetDefaultNameserversRetType, ok bool) {
	return getAreaConfigGetDefaultNameserversAttributeTypeOk(o.DefaultNameservers)
}

// HasDefaultNameservers returns a boolean if a field has been set.
func (o *AreaConfig) HasDefaultNameservers() bool {
	_, ok := o.GetDefaultNameserversOk()
	return ok
}

// SetDefaultNameservers gets a reference to the given []string and assigns it to the DefaultNameservers field.
func (o *AreaConfig) SetDefaultNameservers(v AreaConfigGetDefaultNameserversRetType) {
	setAreaConfigGetDefaultNameserversAttributeType(&o.DefaultNameservers, v)
}

// GetNetworkRanges returns the NetworkRanges field value
func (o *AreaConfig) GetNetworkRanges() (ret AreaConfigGetNetworkRangesRetType) {
	ret, _ = o.GetNetworkRangesOk()
	return ret
}

// GetNetworkRangesOk returns a tuple with the NetworkRanges field value
// and a boolean to check if the value has been set.
func (o *AreaConfig) GetNetworkRangesOk() (ret AreaConfigGetNetworkRangesRetType, ok bool) {
	return getAreaConfigGetNetworkRangesAttributeTypeOk(o.NetworkRanges)
}

// SetNetworkRanges sets field value
func (o *AreaConfig) SetNetworkRanges(v AreaConfigGetNetworkRangesRetType) {
	setAreaConfigGetNetworkRangesAttributeType(&o.NetworkRanges, v)
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *AreaConfig) GetRoutes() (res AreaConfigGetRoutesRetType) {
	res, _ = o.GetRoutesOk()
	return
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaConfig) GetRoutesOk() (ret AreaConfigGetRoutesRetType, ok bool) {
	return getAreaConfigGetRoutesAttributeTypeOk(o.Routes)
}

// HasRoutes returns a boolean if a field has been set.
func (o *AreaConfig) HasRoutes() bool {
	_, ok := o.GetRoutesOk()
	return ok
}

// SetRoutes gets a reference to the given []Route and assigns it to the Routes field.
func (o *AreaConfig) SetRoutes(v AreaConfigGetRoutesRetType) {
	setAreaConfigGetRoutesAttributeType(&o.Routes, v)
}

// GetTransferNetwork returns the TransferNetwork field value
func (o *AreaConfig) GetTransferNetwork() (ret AreaConfigGetTransferNetworkRetType) {
	ret, _ = o.GetTransferNetworkOk()
	return ret
}

// GetTransferNetworkOk returns a tuple with the TransferNetwork field value
// and a boolean to check if the value has been set.
func (o *AreaConfig) GetTransferNetworkOk() (ret AreaConfigGetTransferNetworkRetType, ok bool) {
	return getAreaConfigGetTransferNetworkAttributeTypeOk(o.TransferNetwork)
}

// SetTransferNetwork sets field value
func (o *AreaConfig) SetTransferNetwork(v AreaConfigGetTransferNetworkRetType) {
	setAreaConfigGetTransferNetworkAttributeType(&o.TransferNetwork, v)
}

func (o AreaConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getAreaConfigGetDefaultNameserversAttributeTypeOk(o.DefaultNameservers); ok {
		toSerialize["DefaultNameservers"] = val
	}
	if val, ok := getAreaConfigGetNetworkRangesAttributeTypeOk(o.NetworkRanges); ok {
		toSerialize["NetworkRanges"] = val
	}
	if val, ok := getAreaConfigGetRoutesAttributeTypeOk(o.Routes); ok {
		toSerialize["Routes"] = val
	}
	if val, ok := getAreaConfigGetTransferNetworkAttributeTypeOk(o.TransferNetwork); ok {
		toSerialize["TransferNetwork"] = val
	}
	return toSerialize, nil
}

type NullableAreaConfig struct {
	value *AreaConfig
	isSet bool
}

func (v NullableAreaConfig) Get() *AreaConfig {
	return v.value
}

func (v *NullableAreaConfig) Set(val *AreaConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAreaConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAreaConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAreaConfig(val *AreaConfig) *NullableAreaConfig {
	return &NullableAreaConfig{value: val, isSet: true}
}

func (v NullableAreaConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAreaConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
