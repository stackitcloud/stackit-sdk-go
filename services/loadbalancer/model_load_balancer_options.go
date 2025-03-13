/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

import (
	"encoding/json"
)

// checks if the LoadBalancerOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerOptions{}

// LoadBalancerOptions Defines any optional functionality you want to have enabled on your load balancer.
type LoadBalancerOptions struct {
	AccessControl    *LoadbalancerOptionAccessControl `json:"accessControl,omitempty"`
	EphemeralAddress *bool                            `json:"ephemeralAddress,omitempty"`
	Observability    *LoadbalancerOptionObservability `json:"observability,omitempty"`
	// Load Balancer is accessible only via a private network ip address. Not changeable after creation.
	PrivateNetworkOnly *bool `json:"privateNetworkOnly,omitempty"`
}

// NewLoadBalancerOptions instantiates a new LoadBalancerOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerOptions() *LoadBalancerOptions {
	this := LoadBalancerOptions{}
	return &this
}

// NewLoadBalancerOptionsWithDefaults instantiates a new LoadBalancerOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerOptionsWithDefaults() *LoadBalancerOptions {
	this := LoadBalancerOptions{}
	return &this
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise.
func (o *LoadBalancerOptions) GetAccessControl() *LoadbalancerOptionAccessControl {
	if o == nil || IsNil(o.AccessControl) {
		var ret *LoadbalancerOptionAccessControl
		return ret
	}
	return o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOptions) GetAccessControlOk() (*LoadbalancerOptionAccessControl, bool) {
	if o == nil || IsNil(o.AccessControl) {
		return nil, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *LoadBalancerOptions) HasAccessControl() bool {
	if o != nil && !IsNil(o.AccessControl) {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given LoadbalancerOptionAccessControl and assigns it to the AccessControl field.
func (o *LoadBalancerOptions) SetAccessControl(v *LoadbalancerOptionAccessControl) {
	o.AccessControl = v
}

// GetEphemeralAddress returns the EphemeralAddress field value if set, zero value otherwise.
func (o *LoadBalancerOptions) GetEphemeralAddress() *bool {
	if o == nil || IsNil(o.EphemeralAddress) {
		var ret *bool
		return ret
	}
	return o.EphemeralAddress
}

// GetEphemeralAddressOk returns a tuple with the EphemeralAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOptions) GetEphemeralAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.EphemeralAddress) {
		return nil, false
	}
	return o.EphemeralAddress, true
}

// HasEphemeralAddress returns a boolean if a field has been set.
func (o *LoadBalancerOptions) HasEphemeralAddress() bool {
	if o != nil && !IsNil(o.EphemeralAddress) {
		return true
	}

	return false
}

// SetEphemeralAddress gets a reference to the given bool and assigns it to the EphemeralAddress field.
func (o *LoadBalancerOptions) SetEphemeralAddress(v *bool) {
	o.EphemeralAddress = v
}

// GetObservability returns the Observability field value if set, zero value otherwise.
func (o *LoadBalancerOptions) GetObservability() *LoadbalancerOptionObservability {
	if o == nil || IsNil(o.Observability) {
		var ret *LoadbalancerOptionObservability
		return ret
	}
	return o.Observability
}

// GetObservabilityOk returns a tuple with the Observability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOptions) GetObservabilityOk() (*LoadbalancerOptionObservability, bool) {
	if o == nil || IsNil(o.Observability) {
		return nil, false
	}
	return o.Observability, true
}

// HasObservability returns a boolean if a field has been set.
func (o *LoadBalancerOptions) HasObservability() bool {
	if o != nil && !IsNil(o.Observability) {
		return true
	}

	return false
}

// SetObservability gets a reference to the given LoadbalancerOptionObservability and assigns it to the Observability field.
func (o *LoadBalancerOptions) SetObservability(v *LoadbalancerOptionObservability) {
	o.Observability = v
}

// GetPrivateNetworkOnly returns the PrivateNetworkOnly field value if set, zero value otherwise.
func (o *LoadBalancerOptions) GetPrivateNetworkOnly() *bool {
	if o == nil || IsNil(o.PrivateNetworkOnly) {
		var ret *bool
		return ret
	}
	return o.PrivateNetworkOnly
}

// GetPrivateNetworkOnlyOk returns a tuple with the PrivateNetworkOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOptions) GetPrivateNetworkOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.PrivateNetworkOnly) {
		return nil, false
	}
	return o.PrivateNetworkOnly, true
}

// HasPrivateNetworkOnly returns a boolean if a field has been set.
func (o *LoadBalancerOptions) HasPrivateNetworkOnly() bool {
	if o != nil && !IsNil(o.PrivateNetworkOnly) {
		return true
	}

	return false
}

// SetPrivateNetworkOnly gets a reference to the given bool and assigns it to the PrivateNetworkOnly field.
func (o *LoadBalancerOptions) SetPrivateNetworkOnly(v *bool) {
	o.PrivateNetworkOnly = v
}

func (o LoadBalancerOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessControl) {
		toSerialize["accessControl"] = o.AccessControl
	}
	if !IsNil(o.EphemeralAddress) {
		toSerialize["ephemeralAddress"] = o.EphemeralAddress
	}
	if !IsNil(o.Observability) {
		toSerialize["observability"] = o.Observability
	}
	if !IsNil(o.PrivateNetworkOnly) {
		toSerialize["privateNetworkOnly"] = o.PrivateNetworkOnly
	}
	return toSerialize, nil
}

type NullableLoadBalancerOptions struct {
	value *LoadBalancerOptions
	isSet bool
}

func (v NullableLoadBalancerOptions) Get() *LoadBalancerOptions {
	return v.value
}

func (v *NullableLoadBalancerOptions) Set(val *LoadBalancerOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerOptions(val *LoadBalancerOptions) *NullableLoadBalancerOptions {
	return &NullableLoadBalancerOptions{value: val, isSet: true}
}

func (v NullableLoadBalancerOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
