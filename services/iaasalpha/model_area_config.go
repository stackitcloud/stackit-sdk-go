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

// AreaConfig The basic network area object.
type AreaConfig struct {
	DefaultNameservers *[]string `json:"defaultNameservers,omitempty"`
	// A list of network ranges.
	// REQUIRED
	NetworkRanges *[]NetworkRange `json:"networkRanges"`
	// A list of routes.
	Routes *[]Route `json:"routes,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	// REQUIRED
	TransferNetwork *string `json:"transferNetwork"`
}

type _AreaConfig AreaConfig

// NewAreaConfig instantiates a new AreaConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAreaConfig(networkRanges *[]NetworkRange, transferNetwork *string) *AreaConfig {
	this := AreaConfig{}
	this.NetworkRanges = networkRanges
	this.TransferNetwork = transferNetwork
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
func (o *AreaConfig) GetDefaultNameservers() *[]string {
	if o == nil || IsNil(o.DefaultNameservers) {
		var ret *[]string
		return ret
	}
	return o.DefaultNameservers
}

// GetDefaultNameserversOk returns a tuple with the DefaultNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaConfig) GetDefaultNameserversOk() (*[]string, bool) {
	if o == nil || IsNil(o.DefaultNameservers) {
		return nil, false
	}
	return o.DefaultNameservers, true
}

// HasDefaultNameservers returns a boolean if a field has been set.
func (o *AreaConfig) HasDefaultNameservers() bool {
	if o != nil && !IsNil(o.DefaultNameservers) {
		return true
	}

	return false
}

// SetDefaultNameservers gets a reference to the given []string and assigns it to the DefaultNameservers field.
func (o *AreaConfig) SetDefaultNameservers(v *[]string) {
	o.DefaultNameservers = v
}

// GetNetworkRanges returns the NetworkRanges field value
func (o *AreaConfig) GetNetworkRanges() *[]NetworkRange {
	if o == nil || IsNil(o.NetworkRanges) {
		var ret *[]NetworkRange
		return ret
	}

	return o.NetworkRanges
}

// GetNetworkRangesOk returns a tuple with the NetworkRanges field value
// and a boolean to check if the value has been set.
func (o *AreaConfig) GetNetworkRangesOk() (*[]NetworkRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkRanges, true
}

// SetNetworkRanges sets field value
func (o *AreaConfig) SetNetworkRanges(v *[]NetworkRange) {
	o.NetworkRanges = v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *AreaConfig) GetRoutes() *[]Route {
	if o == nil || IsNil(o.Routes) {
		var ret *[]Route
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaConfig) GetRoutesOk() (*[]Route, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *AreaConfig) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []Route and assigns it to the Routes field.
func (o *AreaConfig) SetRoutes(v *[]Route) {
	o.Routes = v
}

// GetTransferNetwork returns the TransferNetwork field value
func (o *AreaConfig) GetTransferNetwork() *string {
	if o == nil || IsNil(o.TransferNetwork) {
		var ret *string
		return ret
	}

	return o.TransferNetwork
}

// GetTransferNetworkOk returns a tuple with the TransferNetwork field value
// and a boolean to check if the value has been set.
func (o *AreaConfig) GetTransferNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransferNetwork, true
}

// SetTransferNetwork sets field value
func (o *AreaConfig) SetTransferNetwork(v *string) {
	o.TransferNetwork = v
}

func (o AreaConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultNameservers) {
		toSerialize["defaultNameservers"] = o.DefaultNameservers
	}
	toSerialize["networkRanges"] = o.NetworkRanges
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	toSerialize["transferNetwork"] = o.TransferNetwork
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
