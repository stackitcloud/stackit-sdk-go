/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.7.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

import (
	"encoding/json"
)

// checks if the LoadBalancer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancer{}

// LoadBalancer struct for LoadBalancer
type LoadBalancer struct {
	// Reports all errors a load balancer has.
	Errors *[]LoadBalancerError `json:"errors,omitempty"`
	// External load balancer IP address where this load balancer is exposed. Not changeable after creation.
	ExternalAddress *string `json:"externalAddress,omitempty"`
	// There is a maximum listener count of 20.  Port and protocol limitations:  - UDP listeners cannot have the same port. - TCP-derived listeners cannot have the same port. A TCP-derived listener is any listener that listens on a TCP port. As of now those are: TCP, TCP_PROXY, and PROTOCOL_TLS_PASSTHROUGH. The only exception is, if all listeners for the same port are PROTOCOL_TLS_PASSTHROUGH. - PROTOCOL_TLS_PASSTHROUGH listeners cannot have the same port and at least one common domain name. - PROTOCOL_TLS_PASSTHROUGH listeners can have the same domain name and different ports though (e.g. ports 443 and 8443 for domain example.com). - PROTOCOL_TLS_PASSTHROUGH listeners without a domain name serve as a default listener and you can have only one default listener.
	Listeners *[]Listener `json:"listeners,omitempty"`
	// Load balancer name. Not changeable after creation.
	Name *string `json:"name,omitempty"`
	// List of networks that listeners and targets reside in. Currently limited to one. Not changeable after creation.
	Networks *[]Network           `json:"networks,omitempty"`
	Options  *LoadBalancerOptions `json:"options,omitempty"`
	// Service Plan configures the size of the Load Balancer. Currently supported plans are p10, p50, p250 and p750. This list can change in the future where plan ids will be removed and new plans by added. That is the reason this is not an enum.
	PlanId *string `json:"planId,omitempty"`
	// Transient private load balancer IP address that can change any time.
	PrivateAddress *string `json:"privateAddress,omitempty"`
	Status         *string `json:"status,omitempty"`
	// List of all target pools which will be used in the load balancer. Limited to 20.
	TargetPools *[]TargetPool `json:"targetPools,omitempty"`
	// Load balancer resource version. Must be empty or unset for creating load balancers, non-empty for updating load balancers. Semantics: While retrieving load balancers, this is the current version of this load balancer resource that changes during updates of the load balancers. On updates this field specified the load balancer version you calculated your update for instead of the future version to enable concurrency safe updates. Update calls will then report the new version in their result as you would see with a load balancer retrieval call later. There exist no total order of the version, so you can only compare it for equality, but not for less/greater than another version. Since the creation of load balancer is always intended to create the first version of it, there should be no existing version. That's why this field must by empty of not present in that case.
	Version *string `json:"version,omitempty"`
}

// NewLoadBalancer instantiates a new LoadBalancer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancer() *LoadBalancer {
	this := LoadBalancer{}
	return &this
}

// NewLoadBalancerWithDefaults instantiates a new LoadBalancer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerWithDefaults() *LoadBalancer {
	this := LoadBalancer{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *LoadBalancer) GetErrors() *[]LoadBalancerError {
	if o == nil || IsNil(o.Errors) {
		var ret *[]LoadBalancerError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetErrorsOk() (*[]LoadBalancerError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *LoadBalancer) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []LoadBalancerError and assigns it to the Errors field.
func (o *LoadBalancer) SetErrors(v *[]LoadBalancerError) {
	o.Errors = v
}

// GetExternalAddress returns the ExternalAddress field value if set, zero value otherwise.
func (o *LoadBalancer) GetExternalAddress() *string {
	if o == nil || IsNil(o.ExternalAddress) {
		var ret *string
		return ret
	}
	return o.ExternalAddress
}

// GetExternalAddressOk returns a tuple with the ExternalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetExternalAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalAddress) {
		return nil, false
	}
	return o.ExternalAddress, true
}

// HasExternalAddress returns a boolean if a field has been set.
func (o *LoadBalancer) HasExternalAddress() bool {
	if o != nil && !IsNil(o.ExternalAddress) && !IsNil(o.ExternalAddress) {
		return true
	}

	return false
}

// SetExternalAddress gets a reference to the given string and assigns it to the ExternalAddress field.
func (o *LoadBalancer) SetExternalAddress(v *string) {
	o.ExternalAddress = v
}

// GetListeners returns the Listeners field value if set, zero value otherwise.
func (o *LoadBalancer) GetListeners() *[]Listener {
	if o == nil || IsNil(o.Listeners) {
		var ret *[]Listener
		return ret
	}
	return o.Listeners
}

// GetListenersOk returns a tuple with the Listeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetListenersOk() (*[]Listener, bool) {
	if o == nil || IsNil(o.Listeners) {
		return nil, false
	}
	return o.Listeners, true
}

// HasListeners returns a boolean if a field has been set.
func (o *LoadBalancer) HasListeners() bool {
	if o != nil && !IsNil(o.Listeners) && !IsNil(o.Listeners) {
		return true
	}

	return false
}

// SetListeners gets a reference to the given []Listener and assigns it to the Listeners field.
func (o *LoadBalancer) SetListeners(v *[]Listener) {
	o.Listeners = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoadBalancer) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoadBalancer) HasName() bool {
	if o != nil && !IsNil(o.Name) && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoadBalancer) SetName(v *string) {
	o.Name = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *LoadBalancer) GetNetworks() *[]Network {
	if o == nil || IsNil(o.Networks) {
		var ret *[]Network
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetNetworksOk() (*[]Network, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *LoadBalancer) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []Network and assigns it to the Networks field.
func (o *LoadBalancer) SetNetworks(v *[]Network) {
	o.Networks = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *LoadBalancer) GetOptions() *LoadBalancerOptions {
	if o == nil || IsNil(o.Options) {
		var ret *LoadBalancerOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetOptionsOk() (*LoadBalancerOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *LoadBalancer) HasOptions() bool {
	if o != nil && !IsNil(o.Options) && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given LoadBalancerOptions and assigns it to the Options field.
func (o *LoadBalancer) SetOptions(v *LoadBalancerOptions) {
	o.Options = v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *LoadBalancer) GetPlanId() *string {
	if o == nil || IsNil(o.PlanId) {
		var ret *string
		return ret
	}
	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *LoadBalancer) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *LoadBalancer) SetPlanId(v *string) {
	o.PlanId = v
}

// GetPrivateAddress returns the PrivateAddress field value if set, zero value otherwise.
func (o *LoadBalancer) GetPrivateAddress() *string {
	if o == nil || IsNil(o.PrivateAddress) {
		var ret *string
		return ret
	}
	return o.PrivateAddress
}

// GetPrivateAddressOk returns a tuple with the PrivateAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetPrivateAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateAddress) {
		return nil, false
	}
	return o.PrivateAddress, true
}

// HasPrivateAddress returns a boolean if a field has been set.
func (o *LoadBalancer) HasPrivateAddress() bool {
	if o != nil && !IsNil(o.PrivateAddress) && !IsNil(o.PrivateAddress) {
		return true
	}

	return false
}

// SetPrivateAddress gets a reference to the given string and assigns it to the PrivateAddress field.
func (o *LoadBalancer) SetPrivateAddress(v *string) {
	o.PrivateAddress = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LoadBalancer) GetStatus() *string {
	if o == nil || IsNil(o.Status) {
		var ret *string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LoadBalancer) HasStatus() bool {
	if o != nil && !IsNil(o.Status) && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LoadBalancer) SetStatus(v *string) {
	o.Status = v
}

// GetTargetPools returns the TargetPools field value if set, zero value otherwise.
func (o *LoadBalancer) GetTargetPools() *[]TargetPool {
	if o == nil || IsNil(o.TargetPools) {
		var ret *[]TargetPool
		return ret
	}
	return o.TargetPools
}

// GetTargetPoolsOk returns a tuple with the TargetPools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetTargetPoolsOk() (*[]TargetPool, bool) {
	if o == nil || IsNil(o.TargetPools) {
		return nil, false
	}
	return o.TargetPools, true
}

// HasTargetPools returns a boolean if a field has been set.
func (o *LoadBalancer) HasTargetPools() bool {
	if o != nil && !IsNil(o.TargetPools) && !IsNil(o.TargetPools) {
		return true
	}

	return false
}

// SetTargetPools gets a reference to the given []TargetPool and assigns it to the TargetPools field.
func (o *LoadBalancer) SetTargetPools(v *[]TargetPool) {
	o.TargetPools = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LoadBalancer) GetVersion() *string {
	if o == nil || IsNil(o.Version) {
		var ret *string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LoadBalancer) HasVersion() bool {
	if o != nil && !IsNil(o.Version) && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LoadBalancer) SetVersion(v *string) {
	o.Version = v
}

func (o LoadBalancer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.ExternalAddress) {
		toSerialize["externalAddress"] = o.ExternalAddress
	}
	if !IsNil(o.Listeners) {
		toSerialize["listeners"] = o.Listeners
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.PlanId) {
		toSerialize["planId"] = o.PlanId
	}
	if !IsNil(o.PrivateAddress) {
		toSerialize["privateAddress"] = o.PrivateAddress
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TargetPools) {
		toSerialize["targetPools"] = o.TargetPools
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableLoadBalancer struct {
	value *LoadBalancer
	isSet bool
}

func (v NullableLoadBalancer) Get() *LoadBalancer {
	return v.value
}

func (v *NullableLoadBalancer) Set(val *LoadBalancer) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancer) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancer(val *LoadBalancer) *NullableLoadBalancer {
	return &NullableLoadBalancer{value: val, isSet: true}
}

func (v NullableLoadBalancer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
