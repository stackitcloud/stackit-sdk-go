/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

import (
	"encoding/json"
)

// checks if the Kubernetes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Kubernetes{}

/*
	types and functions for allowPrivilegedContainers
*/

// isBoolean
type KubernetesgetAllowPrivilegedContainersAttributeType = *bool
type KubernetesgetAllowPrivilegedContainersArgType = bool
type KubernetesgetAllowPrivilegedContainersRetType = bool

func getKubernetesgetAllowPrivilegedContainersAttributeTypeOk(arg KubernetesgetAllowPrivilegedContainersAttributeType) (ret KubernetesgetAllowPrivilegedContainersRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setKubernetesgetAllowPrivilegedContainersAttributeType(arg *KubernetesgetAllowPrivilegedContainersAttributeType, val KubernetesgetAllowPrivilegedContainersRetType) {
	*arg = &val
}

/*
	types and functions for version
*/

// isNotNullableString
type KubernetesGetVersionAttributeType = *string

func getKubernetesGetVersionAttributeTypeOk(arg KubernetesGetVersionAttributeType) (ret KubernetesGetVersionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setKubernetesGetVersionAttributeType(arg *KubernetesGetVersionAttributeType, val KubernetesGetVersionRetType) {
	*arg = &val
}

type KubernetesGetVersionArgType = string
type KubernetesGetVersionRetType = string

// Kubernetes For valid versions please take a look at [provider-options](#tag/ProviderOptions/operation/SkeService_GetProviderOptions) `kubernetesVersions`.
type Kubernetes struct {
	// DEPRECATED as of Kubernetes 1.25+ Flag to specify if privileged mode for containers is enabled or not. This should be used with care since it also disables a couple of other features like the use of some volume type (e.g. PVCs). By default this is set to true.
	AllowPrivilegedContainers KubernetesgetAllowPrivilegedContainersAttributeType `json:"allowPrivilegedContainers,omitempty"`
	// REQUIRED
	Version KubernetesGetVersionAttributeType `json:"version"`
}

type _Kubernetes Kubernetes

// NewKubernetes instantiates a new Kubernetes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetes(version KubernetesGetVersionArgType) *Kubernetes {
	this := Kubernetes{}
	setKubernetesGetVersionAttributeType(&this.Version, version)
	return &this
}

// NewKubernetesWithDefaults instantiates a new Kubernetes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesWithDefaults() *Kubernetes {
	this := Kubernetes{}
	return &this
}

// GetAllowPrivilegedContainers returns the AllowPrivilegedContainers field value if set, zero value otherwise.
func (o *Kubernetes) GetAllowPrivilegedContainers() (res KubernetesgetAllowPrivilegedContainersRetType) {
	res, _ = o.GetAllowPrivilegedContainersOk()
	return
}

// GetAllowPrivilegedContainersOk returns a tuple with the AllowPrivilegedContainers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kubernetes) GetAllowPrivilegedContainersOk() (ret KubernetesgetAllowPrivilegedContainersRetType, ok bool) {
	return getKubernetesgetAllowPrivilegedContainersAttributeTypeOk(o.AllowPrivilegedContainers)
}

// HasAllowPrivilegedContainers returns a boolean if a field has been set.
func (o *Kubernetes) HasAllowPrivilegedContainers() bool {
	_, ok := o.GetAllowPrivilegedContainersOk()
	return ok
}

// SetAllowPrivilegedContainers gets a reference to the given bool and assigns it to the AllowPrivilegedContainers field.
func (o *Kubernetes) SetAllowPrivilegedContainers(v KubernetesgetAllowPrivilegedContainersRetType) {
	setKubernetesgetAllowPrivilegedContainersAttributeType(&o.AllowPrivilegedContainers, v)
}

// GetVersion returns the Version field value
func (o *Kubernetes) GetVersion() (ret KubernetesGetVersionRetType) {
	ret, _ = o.GetVersionOk()
	return ret
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Kubernetes) GetVersionOk() (ret KubernetesGetVersionRetType, ok bool) {
	return getKubernetesGetVersionAttributeTypeOk(o.Version)
}

// SetVersion sets field value
func (o *Kubernetes) SetVersion(v KubernetesGetVersionRetType) {
	setKubernetesGetVersionAttributeType(&o.Version, v)
}

func (o Kubernetes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getKubernetesgetAllowPrivilegedContainersAttributeTypeOk(o.AllowPrivilegedContainers); ok {
		toSerialize["AllowPrivilegedContainers"] = val
	}
	if val, ok := getKubernetesGetVersionAttributeTypeOk(o.Version); ok {
		toSerialize["Version"] = val
	}
	return toSerialize, nil
}

type NullableKubernetes struct {
	value *Kubernetes
	isSet bool
}

func (v NullableKubernetes) Get() *Kubernetes {
	return v.value
}

func (v *NullableKubernetes) Set(val *Kubernetes) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetes) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetes(val *Kubernetes) *NullableKubernetes {
	return &NullableKubernetes{value: val, isSet: true}
}

func (v NullableKubernetes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
