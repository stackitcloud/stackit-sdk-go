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

// checks if the CreateAreaIPv4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAreaIPv4{}

// CreateAreaIPv4 The config object for a IPv4 network area.
type CreateAreaIPv4 struct {
	DefaultNameservers *[]string `json:"defaultNameservers,omitempty"`
	// A list of network ranges.
	// REQUIRED
	NetworkRanges *[]NetworkRange `json:"networkRanges"`
	// A list of routes.
	Routes *[]Route `json:"routes,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	// REQUIRED
	TransferNetwork *string `json:"transferNetwork"`
	// The default prefix length for networks in the network area.
	DefaultPrefixLen *int64 `json:"defaultPrefixLen,omitempty"`
	// The maximal prefix length for networks in the network area.
	MaxPrefixLen *int64 `json:"maxPrefixLen,omitempty"`
	// The minimal prefix length for networks in the network area.
	MinPrefixLen *int64 `json:"minPrefixLen,omitempty"`
}

type _CreateAreaIPv4 CreateAreaIPv4

// NewCreateAreaIPv4 instantiates a new CreateAreaIPv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAreaIPv4(networkRanges *[]NetworkRange, transferNetwork *string) *CreateAreaIPv4 {
	this := CreateAreaIPv4{}
	this.NetworkRanges = networkRanges
	this.TransferNetwork = transferNetwork
	var defaultPrefixLen int64 = 25
	this.DefaultPrefixLen = &defaultPrefixLen
	var maxPrefixLen int64 = 29
	this.MaxPrefixLen = &maxPrefixLen
	var minPrefixLen int64 = 24
	this.MinPrefixLen = &minPrefixLen
	return &this
}

// NewCreateAreaIPv4WithDefaults instantiates a new CreateAreaIPv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAreaIPv4WithDefaults() *CreateAreaIPv4 {
	this := CreateAreaIPv4{}
	var defaultPrefixLen int64 = 25
	this.DefaultPrefixLen = &defaultPrefixLen
	var maxPrefixLen int64 = 29
	this.MaxPrefixLen = &maxPrefixLen
	var minPrefixLen int64 = 24
	this.MinPrefixLen = &minPrefixLen
	return &this
}

// GetDefaultNameservers returns the DefaultNameservers field value if set, zero value otherwise.
func (o *CreateAreaIPv4) GetDefaultNameservers() *[]string {
	if o == nil || IsNil(o.DefaultNameservers) {
		var ret *[]string
		return ret
	}
	return o.DefaultNameservers
}

// GetDefaultNameserversOk returns a tuple with the DefaultNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAreaIPv4) GetDefaultNameserversOk() (*[]string, bool) {
	if o == nil || IsNil(o.DefaultNameservers) {
		return nil, false
	}
	return o.DefaultNameservers, true
}

// HasDefaultNameservers returns a boolean if a field has been set.
func (o *CreateAreaIPv4) HasDefaultNameservers() bool {
	if o != nil && !IsNil(o.DefaultNameservers) && !IsNil(o.DefaultNameservers) {
		return true
	}

	return false
}

// SetDefaultNameservers gets a reference to the given []string and assigns it to the DefaultNameservers field.
func (o *CreateAreaIPv4) SetDefaultNameservers(v *[]string) {
	o.DefaultNameservers = v
}

// GetNetworkRanges returns the NetworkRanges field value
func (o *CreateAreaIPv4) GetNetworkRanges() *[]NetworkRange {
	if o == nil || IsNil(o.NetworkRanges) {
		var ret *[]NetworkRange
		return ret
	}

	return o.NetworkRanges
}

// GetNetworkRangesOk returns a tuple with the NetworkRanges field value
// and a boolean to check if the value has been set.
func (o *CreateAreaIPv4) GetNetworkRangesOk() (*[]NetworkRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkRanges, true
}

// SetNetworkRanges sets field value
func (o *CreateAreaIPv4) SetNetworkRanges(v *[]NetworkRange) {
	o.NetworkRanges = v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *CreateAreaIPv4) GetRoutes() *[]Route {
	if o == nil || IsNil(o.Routes) {
		var ret *[]Route
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAreaIPv4) GetRoutesOk() (*[]Route, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *CreateAreaIPv4) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []Route and assigns it to the Routes field.
func (o *CreateAreaIPv4) SetRoutes(v *[]Route) {
	o.Routes = v
}

// GetTransferNetwork returns the TransferNetwork field value
func (o *CreateAreaIPv4) GetTransferNetwork() *string {
	if o == nil || IsNil(o.TransferNetwork) {
		var ret *string
		return ret
	}

	return o.TransferNetwork
}

// GetTransferNetworkOk returns a tuple with the TransferNetwork field value
// and a boolean to check if the value has been set.
func (o *CreateAreaIPv4) GetTransferNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransferNetwork, true
}

// SetTransferNetwork sets field value
func (o *CreateAreaIPv4) SetTransferNetwork(v *string) {
	o.TransferNetwork = v
}

// GetDefaultPrefixLen returns the DefaultPrefixLen field value if set, zero value otherwise.
func (o *CreateAreaIPv4) GetDefaultPrefixLen() *int64 {
	if o == nil || IsNil(o.DefaultPrefixLen) {
		var ret *int64
		return ret
	}
	return o.DefaultPrefixLen
}

// GetDefaultPrefixLenOk returns a tuple with the DefaultPrefixLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAreaIPv4) GetDefaultPrefixLenOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultPrefixLen) {
		return nil, false
	}
	return o.DefaultPrefixLen, true
}

// HasDefaultPrefixLen returns a boolean if a field has been set.
func (o *CreateAreaIPv4) HasDefaultPrefixLen() bool {
	if o != nil && !IsNil(o.DefaultPrefixLen) && !IsNil(o.DefaultPrefixLen) {
		return true
	}

	return false
}

// SetDefaultPrefixLen gets a reference to the given int64 and assigns it to the DefaultPrefixLen field.
func (o *CreateAreaIPv4) SetDefaultPrefixLen(v *int64) {
	o.DefaultPrefixLen = v
}

// GetMaxPrefixLen returns the MaxPrefixLen field value if set, zero value otherwise.
func (o *CreateAreaIPv4) GetMaxPrefixLen() *int64 {
	if o == nil || IsNil(o.MaxPrefixLen) {
		var ret *int64
		return ret
	}
	return o.MaxPrefixLen
}

// GetMaxPrefixLenOk returns a tuple with the MaxPrefixLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAreaIPv4) GetMaxPrefixLenOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxPrefixLen) {
		return nil, false
	}
	return o.MaxPrefixLen, true
}

// HasMaxPrefixLen returns a boolean if a field has been set.
func (o *CreateAreaIPv4) HasMaxPrefixLen() bool {
	if o != nil && !IsNil(o.MaxPrefixLen) && !IsNil(o.MaxPrefixLen) {
		return true
	}

	return false
}

// SetMaxPrefixLen gets a reference to the given int64 and assigns it to the MaxPrefixLen field.
func (o *CreateAreaIPv4) SetMaxPrefixLen(v *int64) {
	o.MaxPrefixLen = v
}

// GetMinPrefixLen returns the MinPrefixLen field value if set, zero value otherwise.
func (o *CreateAreaIPv4) GetMinPrefixLen() *int64 {
	if o == nil || IsNil(o.MinPrefixLen) {
		var ret *int64
		return ret
	}
	return o.MinPrefixLen
}

// GetMinPrefixLenOk returns a tuple with the MinPrefixLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAreaIPv4) GetMinPrefixLenOk() (*int64, bool) {
	if o == nil || IsNil(o.MinPrefixLen) {
		return nil, false
	}
	return o.MinPrefixLen, true
}

// HasMinPrefixLen returns a boolean if a field has been set.
func (o *CreateAreaIPv4) HasMinPrefixLen() bool {
	if o != nil && !IsNil(o.MinPrefixLen) && !IsNil(o.MinPrefixLen) {
		return true
	}

	return false
}

// SetMinPrefixLen gets a reference to the given int64 and assigns it to the MinPrefixLen field.
func (o *CreateAreaIPv4) SetMinPrefixLen(v *int64) {
	o.MinPrefixLen = v
}

func (o CreateAreaIPv4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultNameservers) {
		toSerialize["defaultNameservers"] = o.DefaultNameservers
	}
	toSerialize["networkRanges"] = o.NetworkRanges
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	toSerialize["transferNetwork"] = o.TransferNetwork
	if !IsNil(o.DefaultPrefixLen) {
		toSerialize["defaultPrefixLen"] = o.DefaultPrefixLen
	}
	if !IsNil(o.MaxPrefixLen) {
		toSerialize["maxPrefixLen"] = o.MaxPrefixLen
	}
	if !IsNil(o.MinPrefixLen) {
		toSerialize["minPrefixLen"] = o.MinPrefixLen
	}
	return toSerialize, nil
}

type NullableCreateAreaIPv4 struct {
	value *CreateAreaIPv4
	isSet bool
}

func (v NullableCreateAreaIPv4) Get() *CreateAreaIPv4 {
	return v.value
}

func (v *NullableCreateAreaIPv4) Set(val *CreateAreaIPv4) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAreaIPv4) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAreaIPv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAreaIPv4(val *CreateAreaIPv4) *NullableCreateAreaIPv4 {
	return &NullableCreateAreaIPv4{value: val, isSet: true}
}

func (v NullableCreateAreaIPv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAreaIPv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
