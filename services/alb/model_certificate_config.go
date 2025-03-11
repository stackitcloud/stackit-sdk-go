/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 2beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alb

import (
	"encoding/json"
)

// checks if the CertificateConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateConfig{}

// CertificateConfig struct for CertificateConfig
type CertificateConfig struct {
	// Certificate IDs for TLS termination
	CertificateIds *[]string `json:"certificateIds,omitempty"`
}

// NewCertificateConfig instantiates a new CertificateConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateConfig() *CertificateConfig {
	this := CertificateConfig{}
	return &this
}

// NewCertificateConfigWithDefaults instantiates a new CertificateConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateConfigWithDefaults() *CertificateConfig {
	this := CertificateConfig{}
	return &this
}

// GetCertificateIds returns the CertificateIds field value if set, zero value otherwise.
func (o *CertificateConfig) GetCertificateIds() *[]string {
	if o == nil || IsNil(o.CertificateIds) {
		var ret *[]string
		return ret
	}
	return o.CertificateIds
}

// GetCertificateIdsOk returns a tuple with the CertificateIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateConfig) GetCertificateIdsOk() (*[]string, bool) {
	if o == nil || IsNil(o.CertificateIds) {
		return nil, false
	}
	return o.CertificateIds, true
}

// HasCertificateIds returns a boolean if a field has been set.
func (o *CertificateConfig) HasCertificateIds() bool {
	if o != nil && !IsNil(o.CertificateIds) {
		return true
	}

	return false
}

// SetCertificateIds gets a reference to the given []string and assigns it to the CertificateIds field.
func (o *CertificateConfig) SetCertificateIds(v *[]string) {
	o.CertificateIds = v
}

func (o CertificateConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertificateIds) {
		toSerialize["certificateIds"] = o.CertificateIds
	}
	return toSerialize, nil
}

type NullableCertificateConfig struct {
	value *CertificateConfig
	isSet bool
}

func (v NullableCertificateConfig) Get() *CertificateConfig {
	return v.value
}

func (v *NullableCertificateConfig) Set(val *CertificateConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateConfig(val *CertificateConfig) *NullableCertificateConfig {
	return &NullableCertificateConfig{value: val, isSet: true}
}

func (v NullableCertificateConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
