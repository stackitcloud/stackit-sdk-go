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

// checks if the Area type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Area{}

// Area The basic properties of a network area.
type Area struct {
	// A list containing DNS Servers/Nameservers for IPv4.
	DefaultNameservers *[]string `json:"defaultNameservers,omitempty"`
	// A list of network ranges.
	NetworkRanges *[]NetworkRange `json:"networkRanges,omitempty"`
	// A list of routes.
	Routes *[]Route `json:"routes,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	TransferNetwork *string `json:"transferNetwork,omitempty"`
}

// NewArea instantiates a new Area object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArea() *Area {
	this := Area{}
	return &this
}

// NewAreaWithDefaults instantiates a new Area object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAreaWithDefaults() *Area {
	this := Area{}
	return &this
}

// GetDefaultNameservers returns the DefaultNameservers field value if set, zero value otherwise.
func (o *Area) GetDefaultNameservers() *[]string {
	if o == nil || IsNil(o.DefaultNameservers) {
		var ret *[]string
		return ret
	}
	return o.DefaultNameservers
}

// GetDefaultNameserversOk returns a tuple with the DefaultNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Area) GetDefaultNameserversOk() (*[]string, bool) {
	if o == nil || IsNil(o.DefaultNameservers) {
		return nil, false
	}
	return o.DefaultNameservers, true
}

// HasDefaultNameservers returns a boolean if a field has been set.
func (o *Area) HasDefaultNameservers() bool {
	if o != nil && !IsNil(o.DefaultNameservers) && !IsNil(o.DefaultNameservers) {
		return true
	}

	return false
}

// SetDefaultNameservers gets a reference to the given []string and assigns it to the DefaultNameservers field.
func (o *Area) SetDefaultNameservers(v *[]string) {
	o.DefaultNameservers = v
}

// GetNetworkRanges returns the NetworkRanges field value if set, zero value otherwise.
func (o *Area) GetNetworkRanges() *[]NetworkRange {
	if o == nil || IsNil(o.NetworkRanges) {
		var ret *[]NetworkRange
		return ret
	}
	return o.NetworkRanges
}

// GetNetworkRangesOk returns a tuple with the NetworkRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Area) GetNetworkRangesOk() (*[]NetworkRange, bool) {
	if o == nil || IsNil(o.NetworkRanges) {
		return nil, false
	}
	return o.NetworkRanges, true
}

// HasNetworkRanges returns a boolean if a field has been set.
func (o *Area) HasNetworkRanges() bool {
	if o != nil && !IsNil(o.NetworkRanges) && !IsNil(o.NetworkRanges) {
		return true
	}

	return false
}

// SetNetworkRanges gets a reference to the given []NetworkRange and assigns it to the NetworkRanges field.
func (o *Area) SetNetworkRanges(v *[]NetworkRange) {
	o.NetworkRanges = v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *Area) GetRoutes() *[]Route {
	if o == nil || IsNil(o.Routes) {
		var ret *[]Route
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Area) GetRoutesOk() (*[]Route, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *Area) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []Route and assigns it to the Routes field.
func (o *Area) SetRoutes(v *[]Route) {
	o.Routes = v
}

// GetTransferNetwork returns the TransferNetwork field value if set, zero value otherwise.
func (o *Area) GetTransferNetwork() *string {
	if o == nil || IsNil(o.TransferNetwork) {
		var ret *string
		return ret
	}
	return o.TransferNetwork
}

// GetTransferNetworkOk returns a tuple with the TransferNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Area) GetTransferNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.TransferNetwork) {
		return nil, false
	}
	return o.TransferNetwork, true
}

// HasTransferNetwork returns a boolean if a field has been set.
func (o *Area) HasTransferNetwork() bool {
	if o != nil && !IsNil(o.TransferNetwork) && !IsNil(o.TransferNetwork) {
		return true
	}

	return false
}

// SetTransferNetwork gets a reference to the given string and assigns it to the TransferNetwork field.
func (o *Area) SetTransferNetwork(v *string) {
	o.TransferNetwork = v
}

func (o Area) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultNameservers) {
		toSerialize["defaultNameservers"] = o.DefaultNameservers
	}
	if !IsNil(o.NetworkRanges) {
		toSerialize["networkRanges"] = o.NetworkRanges
	}
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	if !IsNil(o.TransferNetwork) {
		toSerialize["transferNetwork"] = o.TransferNetwork
	}
	return toSerialize, nil
}

type NullableArea struct {
	value *Area
	isSet bool
}

func (v NullableArea) Get() *Area {
	return v.value
}

func (v *NullableArea) Set(val *Area) {
	v.value = val
	v.isSet = true
}

func (v NullableArea) IsSet() bool {
	return v.isSet
}

func (v *NullableArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArea(val *Area) *NullableArea {
	return &NullableArea{value: val, isSet: true}
}

func (v NullableArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
