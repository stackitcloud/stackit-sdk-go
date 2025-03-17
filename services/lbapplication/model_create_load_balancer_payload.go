/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lbapplication

import (
	"encoding/json"
)

// checks if the CreateLoadBalancerPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLoadBalancerPayload{}

/*
	types and functions for errors
*/

// isArray
type CreateLoadBalancerPayloadGetErrorsAttributeType = *[]LoadBalancerError
type CreateLoadBalancerPayloadGetErrorsArgType = []LoadBalancerError
type CreateLoadBalancerPayloadGetErrorsRetType = []LoadBalancerError

func getCreateLoadBalancerPayloadGetErrorsAttributeTypeOk(arg CreateLoadBalancerPayloadGetErrorsAttributeType) (ret CreateLoadBalancerPayloadGetErrorsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateLoadBalancerPayloadGetErrorsAttributeType(arg *CreateLoadBalancerPayloadGetErrorsAttributeType, val CreateLoadBalancerPayloadGetErrorsRetType) {
	*arg = &val
}

/*
	types and functions for externalAddress
*/

// isNotNullableString
type CreateLoadBalancerPayloadGetExternalAddressAttributeType = *string

func getCreateLoadBalancerPayloadGetExternalAddressAttributeTypeOk(arg CreateLoadBalancerPayloadGetExternalAddressAttributeType) (ret CreateLoadBalancerPayloadGetExternalAddressRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateLoadBalancerPayloadGetExternalAddressAttributeType(arg *CreateLoadBalancerPayloadGetExternalAddressAttributeType, val CreateLoadBalancerPayloadGetExternalAddressRetType) {
	*arg = &val
}

type CreateLoadBalancerPayloadGetExternalAddressArgType = string
type CreateLoadBalancerPayloadGetExternalAddressRetType = string

/*
	types and functions for listeners
*/

// isArray
type CreateLoadBalancerPayloadGetListenersAttributeType = *[]Listener
type CreateLoadBalancerPayloadGetListenersArgType = []Listener
type CreateLoadBalancerPayloadGetListenersRetType = []Listener

func getCreateLoadBalancerPayloadGetListenersAttributeTypeOk(arg CreateLoadBalancerPayloadGetListenersAttributeType) (ret CreateLoadBalancerPayloadGetListenersRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateLoadBalancerPayloadGetListenersAttributeType(arg *CreateLoadBalancerPayloadGetListenersAttributeType, val CreateLoadBalancerPayloadGetListenersRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type CreateLoadBalancerPayloadGetNameAttributeType = *string

func getCreateLoadBalancerPayloadGetNameAttributeTypeOk(arg CreateLoadBalancerPayloadGetNameAttributeType) (ret CreateLoadBalancerPayloadGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateLoadBalancerPayloadGetNameAttributeType(arg *CreateLoadBalancerPayloadGetNameAttributeType, val CreateLoadBalancerPayloadGetNameRetType) {
	*arg = &val
}

type CreateLoadBalancerPayloadGetNameArgType = string
type CreateLoadBalancerPayloadGetNameRetType = string

/*
	types and functions for networks
*/

// isArray
type CreateLoadBalancerPayloadGetNetworksAttributeType = *[]Network
type CreateLoadBalancerPayloadGetNetworksArgType = []Network
type CreateLoadBalancerPayloadGetNetworksRetType = []Network

func getCreateLoadBalancerPayloadGetNetworksAttributeTypeOk(arg CreateLoadBalancerPayloadGetNetworksAttributeType) (ret CreateLoadBalancerPayloadGetNetworksRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateLoadBalancerPayloadGetNetworksAttributeType(arg *CreateLoadBalancerPayloadGetNetworksAttributeType, val CreateLoadBalancerPayloadGetNetworksRetType) {
	*arg = &val
}

/*
	types and functions for options
*/

// isModel
type CreateLoadBalancerPayloadGetOptionsAttributeType = *LoadBalancerOptions
type CreateLoadBalancerPayloadGetOptionsArgType = LoadBalancerOptions
type CreateLoadBalancerPayloadGetOptionsRetType = LoadBalancerOptions

func getCreateLoadBalancerPayloadGetOptionsAttributeTypeOk(arg CreateLoadBalancerPayloadGetOptionsAttributeType) (ret CreateLoadBalancerPayloadGetOptionsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateLoadBalancerPayloadGetOptionsAttributeType(arg *CreateLoadBalancerPayloadGetOptionsAttributeType, val CreateLoadBalancerPayloadGetOptionsRetType) {
	*arg = &val
}

/*
	types and functions for planId
*/

// isNotNullableString
type CreateLoadBalancerPayloadGetPlanIdAttributeType = *string

func getCreateLoadBalancerPayloadGetPlanIdAttributeTypeOk(arg CreateLoadBalancerPayloadGetPlanIdAttributeType) (ret CreateLoadBalancerPayloadGetPlanIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateLoadBalancerPayloadGetPlanIdAttributeType(arg *CreateLoadBalancerPayloadGetPlanIdAttributeType, val CreateLoadBalancerPayloadGetPlanIdRetType) {
	*arg = &val
}

type CreateLoadBalancerPayloadGetPlanIdArgType = string
type CreateLoadBalancerPayloadGetPlanIdRetType = string

/*
	types and functions for privateAddress
*/

// isNotNullableString
type CreateLoadBalancerPayloadGetPrivateAddressAttributeType = *string

func getCreateLoadBalancerPayloadGetPrivateAddressAttributeTypeOk(arg CreateLoadBalancerPayloadGetPrivateAddressAttributeType) (ret CreateLoadBalancerPayloadGetPrivateAddressRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateLoadBalancerPayloadGetPrivateAddressAttributeType(arg *CreateLoadBalancerPayloadGetPrivateAddressAttributeType, val CreateLoadBalancerPayloadGetPrivateAddressRetType) {
	*arg = &val
}

type CreateLoadBalancerPayloadGetPrivateAddressArgType = string
type CreateLoadBalancerPayloadGetPrivateAddressRetType = string

/*
	types and functions for status
*/

// isEnumRef
type CreateLoadBalancerPayloadGetStatusAttributeType = *string
type CreateLoadBalancerPayloadGetStatusArgType = string
type CreateLoadBalancerPayloadGetStatusRetType = string

func getCreateLoadBalancerPayloadGetStatusAttributeTypeOk(arg CreateLoadBalancerPayloadGetStatusAttributeType) (ret CreateLoadBalancerPayloadGetStatusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateLoadBalancerPayloadGetStatusAttributeType(arg *CreateLoadBalancerPayloadGetStatusAttributeType, val CreateLoadBalancerPayloadGetStatusRetType) {
	*arg = &val
}

/*
	types and functions for targetPools
*/

// isArray
type CreateLoadBalancerPayloadGetTargetPoolsAttributeType = *[]TargetPool
type CreateLoadBalancerPayloadGetTargetPoolsArgType = []TargetPool
type CreateLoadBalancerPayloadGetTargetPoolsRetType = []TargetPool

func getCreateLoadBalancerPayloadGetTargetPoolsAttributeTypeOk(arg CreateLoadBalancerPayloadGetTargetPoolsAttributeType) (ret CreateLoadBalancerPayloadGetTargetPoolsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateLoadBalancerPayloadGetTargetPoolsAttributeType(arg *CreateLoadBalancerPayloadGetTargetPoolsAttributeType, val CreateLoadBalancerPayloadGetTargetPoolsRetType) {
	*arg = &val
}

/*
	types and functions for version
*/

// isNotNullableString
type CreateLoadBalancerPayloadGetVersionAttributeType = *string

func getCreateLoadBalancerPayloadGetVersionAttributeTypeOk(arg CreateLoadBalancerPayloadGetVersionAttributeType) (ret CreateLoadBalancerPayloadGetVersionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateLoadBalancerPayloadGetVersionAttributeType(arg *CreateLoadBalancerPayloadGetVersionAttributeType, val CreateLoadBalancerPayloadGetVersionRetType) {
	*arg = &val
}

type CreateLoadBalancerPayloadGetVersionArgType = string
type CreateLoadBalancerPayloadGetVersionRetType = string

// CreateLoadBalancerPayload struct for CreateLoadBalancerPayload
type CreateLoadBalancerPayload struct {
	// Reports all errors a application load balancer has.
	Errors CreateLoadBalancerPayloadGetErrorsAttributeType `json:"errors,omitempty"`
	// External application load balancer IP address where this application load balancer is exposed. Not changeable after creation.
	ExternalAddress CreateLoadBalancerPayloadGetExternalAddressAttributeType `json:"externalAddress,omitempty"`
	// There is a maximum listener count of 20.
	Listeners CreateLoadBalancerPayloadGetListenersAttributeType `json:"listeners,omitempty"`
	// Application Load Balancer name. Not changeable after creation.
	Name CreateLoadBalancerPayloadGetNameAttributeType `json:"name,omitempty"`
	// List of networks that listeners and targets reside in. Currently limited to one. Not changeable after creation.
	Networks CreateLoadBalancerPayloadGetNetworksAttributeType `json:"networks,omitempty"`
	Options  CreateLoadBalancerPayloadGetOptionsAttributeType  `json:"options,omitempty"`
	// Service Plan configures the size of the Application Load Balancer. Currently supported plans are p10, p50, p250 and p750. This list can change in the future where plan ids will be removed and new plans by added. That is the reason this is not an enum.
	PlanId CreateLoadBalancerPayloadGetPlanIdAttributeType `json:"planId,omitempty"`
	// Transient private application load balancer IP address that can change any time.
	PrivateAddress CreateLoadBalancerPayloadGetPrivateAddressAttributeType `json:"privateAddress,omitempty"`
	Status         CreateLoadBalancerPayloadGetStatusAttributeType         `json:"status,omitempty"`
	// List of all target pools which will be used in the application load balancer. Limited to 20.
	TargetPools CreateLoadBalancerPayloadGetTargetPoolsAttributeType `json:"targetPools,omitempty"`
	// Application Load Balancer resource version. Must be empty or unset for creating load balancers, non-empty for updating load balancers. Semantics: While retrieving load balancers, this is the current version of this application load balancer resource that changes during updates of the load balancers. On updates this field specified the application load balancer version you calculated your update for instead of the future version to enable concurrency safe updates. Update calls will then report the new version in their result as you would see with a application load balancer retrieval call later. There exist no total order of the version, so you can only compare it for equality, but not for less/greater than another version. Since the creation of application load balancer is always intended to create the first version of it, there should be no existing version. That's why this field must by empty of not present in that case.
	Version CreateLoadBalancerPayloadGetVersionAttributeType `json:"version,omitempty"`
}

// NewCreateLoadBalancerPayload instantiates a new CreateLoadBalancerPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLoadBalancerPayload() *CreateLoadBalancerPayload {
	this := CreateLoadBalancerPayload{}
	return &this
}

// NewCreateLoadBalancerPayloadWithDefaults instantiates a new CreateLoadBalancerPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLoadBalancerPayloadWithDefaults() *CreateLoadBalancerPayload {
	this := CreateLoadBalancerPayload{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateLoadBalancerPayload) GetErrors() (res CreateLoadBalancerPayloadGetErrorsRetType) {
	res, _ = o.GetErrorsOk()
	return
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPayload) GetErrorsOk() (ret CreateLoadBalancerPayloadGetErrorsRetType, ok bool) {
	return getCreateLoadBalancerPayloadGetErrorsAttributeTypeOk(o.Errors)
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateLoadBalancerPayload) HasErrors() bool {
	_, ok := o.GetErrorsOk()
	return ok
}

// SetErrors gets a reference to the given []LoadBalancerError and assigns it to the Errors field.
func (o *CreateLoadBalancerPayload) SetErrors(v CreateLoadBalancerPayloadGetErrorsRetType) {
	setCreateLoadBalancerPayloadGetErrorsAttributeType(&o.Errors, v)
}

// GetExternalAddress returns the ExternalAddress field value if set, zero value otherwise.
func (o *CreateLoadBalancerPayload) GetExternalAddress() (res CreateLoadBalancerPayloadGetExternalAddressRetType) {
	res, _ = o.GetExternalAddressOk()
	return
}

// GetExternalAddressOk returns a tuple with the ExternalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPayload) GetExternalAddressOk() (ret CreateLoadBalancerPayloadGetExternalAddressRetType, ok bool) {
	return getCreateLoadBalancerPayloadGetExternalAddressAttributeTypeOk(o.ExternalAddress)
}

// HasExternalAddress returns a boolean if a field has been set.
func (o *CreateLoadBalancerPayload) HasExternalAddress() bool {
	_, ok := o.GetExternalAddressOk()
	return ok
}

// SetExternalAddress gets a reference to the given string and assigns it to the ExternalAddress field.
func (o *CreateLoadBalancerPayload) SetExternalAddress(v CreateLoadBalancerPayloadGetExternalAddressRetType) {
	setCreateLoadBalancerPayloadGetExternalAddressAttributeType(&o.ExternalAddress, v)
}

// GetListeners returns the Listeners field value if set, zero value otherwise.
func (o *CreateLoadBalancerPayload) GetListeners() (res CreateLoadBalancerPayloadGetListenersRetType) {
	res, _ = o.GetListenersOk()
	return
}

// GetListenersOk returns a tuple with the Listeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPayload) GetListenersOk() (ret CreateLoadBalancerPayloadGetListenersRetType, ok bool) {
	return getCreateLoadBalancerPayloadGetListenersAttributeTypeOk(o.Listeners)
}

// HasListeners returns a boolean if a field has been set.
func (o *CreateLoadBalancerPayload) HasListeners() bool {
	_, ok := o.GetListenersOk()
	return ok
}

// SetListeners gets a reference to the given []Listener and assigns it to the Listeners field.
func (o *CreateLoadBalancerPayload) SetListeners(v CreateLoadBalancerPayloadGetListenersRetType) {
	setCreateLoadBalancerPayloadGetListenersAttributeType(&o.Listeners, v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateLoadBalancerPayload) GetName() (res CreateLoadBalancerPayloadGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPayload) GetNameOk() (ret CreateLoadBalancerPayloadGetNameRetType, ok bool) {
	return getCreateLoadBalancerPayloadGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *CreateLoadBalancerPayload) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateLoadBalancerPayload) SetName(v CreateLoadBalancerPayloadGetNameRetType) {
	setCreateLoadBalancerPayloadGetNameAttributeType(&o.Name, v)
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *CreateLoadBalancerPayload) GetNetworks() (res CreateLoadBalancerPayloadGetNetworksRetType) {
	res, _ = o.GetNetworksOk()
	return
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPayload) GetNetworksOk() (ret CreateLoadBalancerPayloadGetNetworksRetType, ok bool) {
	return getCreateLoadBalancerPayloadGetNetworksAttributeTypeOk(o.Networks)
}

// HasNetworks returns a boolean if a field has been set.
func (o *CreateLoadBalancerPayload) HasNetworks() bool {
	_, ok := o.GetNetworksOk()
	return ok
}

// SetNetworks gets a reference to the given []Network and assigns it to the Networks field.
func (o *CreateLoadBalancerPayload) SetNetworks(v CreateLoadBalancerPayloadGetNetworksRetType) {
	setCreateLoadBalancerPayloadGetNetworksAttributeType(&o.Networks, v)
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CreateLoadBalancerPayload) GetOptions() (res CreateLoadBalancerPayloadGetOptionsRetType) {
	res, _ = o.GetOptionsOk()
	return
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPayload) GetOptionsOk() (ret CreateLoadBalancerPayloadGetOptionsRetType, ok bool) {
	return getCreateLoadBalancerPayloadGetOptionsAttributeTypeOk(o.Options)
}

// HasOptions returns a boolean if a field has been set.
func (o *CreateLoadBalancerPayload) HasOptions() bool {
	_, ok := o.GetOptionsOk()
	return ok
}

// SetOptions gets a reference to the given LoadBalancerOptions and assigns it to the Options field.
func (o *CreateLoadBalancerPayload) SetOptions(v CreateLoadBalancerPayloadGetOptionsRetType) {
	setCreateLoadBalancerPayloadGetOptionsAttributeType(&o.Options, v)
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *CreateLoadBalancerPayload) GetPlanId() (res CreateLoadBalancerPayloadGetPlanIdRetType) {
	res, _ = o.GetPlanIdOk()
	return
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPayload) GetPlanIdOk() (ret CreateLoadBalancerPayloadGetPlanIdRetType, ok bool) {
	return getCreateLoadBalancerPayloadGetPlanIdAttributeTypeOk(o.PlanId)
}

// HasPlanId returns a boolean if a field has been set.
func (o *CreateLoadBalancerPayload) HasPlanId() bool {
	_, ok := o.GetPlanIdOk()
	return ok
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *CreateLoadBalancerPayload) SetPlanId(v CreateLoadBalancerPayloadGetPlanIdRetType) {
	setCreateLoadBalancerPayloadGetPlanIdAttributeType(&o.PlanId, v)
}

// GetPrivateAddress returns the PrivateAddress field value if set, zero value otherwise.
func (o *CreateLoadBalancerPayload) GetPrivateAddress() (res CreateLoadBalancerPayloadGetPrivateAddressRetType) {
	res, _ = o.GetPrivateAddressOk()
	return
}

// GetPrivateAddressOk returns a tuple with the PrivateAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPayload) GetPrivateAddressOk() (ret CreateLoadBalancerPayloadGetPrivateAddressRetType, ok bool) {
	return getCreateLoadBalancerPayloadGetPrivateAddressAttributeTypeOk(o.PrivateAddress)
}

// HasPrivateAddress returns a boolean if a field has been set.
func (o *CreateLoadBalancerPayload) HasPrivateAddress() bool {
	_, ok := o.GetPrivateAddressOk()
	return ok
}

// SetPrivateAddress gets a reference to the given string and assigns it to the PrivateAddress field.
func (o *CreateLoadBalancerPayload) SetPrivateAddress(v CreateLoadBalancerPayloadGetPrivateAddressRetType) {
	setCreateLoadBalancerPayloadGetPrivateAddressAttributeType(&o.PrivateAddress, v)
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateLoadBalancerPayload) GetStatus() (res CreateLoadBalancerPayloadGetStatusRetType) {
	res, _ = o.GetStatusOk()
	return
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPayload) GetStatusOk() (ret CreateLoadBalancerPayloadGetStatusRetType, ok bool) {
	return getCreateLoadBalancerPayloadGetStatusAttributeTypeOk(o.Status)
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateLoadBalancerPayload) HasStatus() bool {
	_, ok := o.GetStatusOk()
	return ok
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateLoadBalancerPayload) SetStatus(v CreateLoadBalancerPayloadGetStatusRetType) {
	setCreateLoadBalancerPayloadGetStatusAttributeType(&o.Status, v)
}

// GetTargetPools returns the TargetPools field value if set, zero value otherwise.
func (o *CreateLoadBalancerPayload) GetTargetPools() (res CreateLoadBalancerPayloadGetTargetPoolsRetType) {
	res, _ = o.GetTargetPoolsOk()
	return
}

// GetTargetPoolsOk returns a tuple with the TargetPools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPayload) GetTargetPoolsOk() (ret CreateLoadBalancerPayloadGetTargetPoolsRetType, ok bool) {
	return getCreateLoadBalancerPayloadGetTargetPoolsAttributeTypeOk(o.TargetPools)
}

// HasTargetPools returns a boolean if a field has been set.
func (o *CreateLoadBalancerPayload) HasTargetPools() bool {
	_, ok := o.GetTargetPoolsOk()
	return ok
}

// SetTargetPools gets a reference to the given []TargetPool and assigns it to the TargetPools field.
func (o *CreateLoadBalancerPayload) SetTargetPools(v CreateLoadBalancerPayloadGetTargetPoolsRetType) {
	setCreateLoadBalancerPayloadGetTargetPoolsAttributeType(&o.TargetPools, v)
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CreateLoadBalancerPayload) GetVersion() (res CreateLoadBalancerPayloadGetVersionRetType) {
	res, _ = o.GetVersionOk()
	return
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPayload) GetVersionOk() (ret CreateLoadBalancerPayloadGetVersionRetType, ok bool) {
	return getCreateLoadBalancerPayloadGetVersionAttributeTypeOk(o.Version)
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateLoadBalancerPayload) HasVersion() bool {
	_, ok := o.GetVersionOk()
	return ok
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CreateLoadBalancerPayload) SetVersion(v CreateLoadBalancerPayloadGetVersionRetType) {
	setCreateLoadBalancerPayloadGetVersionAttributeType(&o.Version, v)
}

func (o CreateLoadBalancerPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateLoadBalancerPayloadGetErrorsAttributeTypeOk(o.Errors); ok {
		toSerialize["Errors"] = val
	}
	if val, ok := getCreateLoadBalancerPayloadGetExternalAddressAttributeTypeOk(o.ExternalAddress); ok {
		toSerialize["ExternalAddress"] = val
	}
	if val, ok := getCreateLoadBalancerPayloadGetListenersAttributeTypeOk(o.Listeners); ok {
		toSerialize["Listeners"] = val
	}
	if val, ok := getCreateLoadBalancerPayloadGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getCreateLoadBalancerPayloadGetNetworksAttributeTypeOk(o.Networks); ok {
		toSerialize["Networks"] = val
	}
	if val, ok := getCreateLoadBalancerPayloadGetOptionsAttributeTypeOk(o.Options); ok {
		toSerialize["Options"] = val
	}
	if val, ok := getCreateLoadBalancerPayloadGetPlanIdAttributeTypeOk(o.PlanId); ok {
		toSerialize["PlanId"] = val
	}
	if val, ok := getCreateLoadBalancerPayloadGetPrivateAddressAttributeTypeOk(o.PrivateAddress); ok {
		toSerialize["PrivateAddress"] = val
	}
	if val, ok := getCreateLoadBalancerPayloadGetStatusAttributeTypeOk(o.Status); ok {
		toSerialize["Status"] = val
	}
	if val, ok := getCreateLoadBalancerPayloadGetTargetPoolsAttributeTypeOk(o.TargetPools); ok {
		toSerialize["TargetPools"] = val
	}
	if val, ok := getCreateLoadBalancerPayloadGetVersionAttributeTypeOk(o.Version); ok {
		toSerialize["Version"] = val
	}
	return toSerialize, nil
}

type NullableCreateLoadBalancerPayload struct {
	value *CreateLoadBalancerPayload
	isSet bool
}

func (v NullableCreateLoadBalancerPayload) Get() *CreateLoadBalancerPayload {
	return v.value
}

func (v *NullableCreateLoadBalancerPayload) Set(val *CreateLoadBalancerPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLoadBalancerPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLoadBalancerPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLoadBalancerPayload(val *CreateLoadBalancerPayload) *NullableCreateLoadBalancerPayload {
	return &NullableCreateLoadBalancerPayload{value: val, isSet: true}
}

func (v NullableCreateLoadBalancerPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLoadBalancerPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
