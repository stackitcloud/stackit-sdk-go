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

// checks if the ProtocolOptionsHTTPS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtocolOptionsHTTPS{}

// ProtocolOptionsHTTPS struct for ProtocolOptionsHTTPS
type ProtocolOptionsHTTPS struct {
	CertificateConfig *CertificateConfig `json:"certificateConfig,omitempty"`
}

// NewProtocolOptionsHTTPS instantiates a new ProtocolOptionsHTTPS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolOptionsHTTPS() *ProtocolOptionsHTTPS {
	this := ProtocolOptionsHTTPS{}
	return &this
}

// NewProtocolOptionsHTTPSWithDefaults instantiates a new ProtocolOptionsHTTPS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolOptionsHTTPSWithDefaults() *ProtocolOptionsHTTPS {
	this := ProtocolOptionsHTTPS{}
	return &this
}

// GetCertificateConfig returns the CertificateConfig field value if set, zero value otherwise.
func (o *ProtocolOptionsHTTPS) GetCertificateConfig() *CertificateConfig {
	if o == nil || IsNil(o.CertificateConfig) {
		var ret *CertificateConfig
		return ret
	}
	return o.CertificateConfig
}

// GetCertificateConfigOk returns a tuple with the CertificateConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolOptionsHTTPS) GetCertificateConfigOk() (*CertificateConfig, bool) {
	if o == nil || IsNil(o.CertificateConfig) {
		return nil, false
	}
	return o.CertificateConfig, true
}

// HasCertificateConfig returns a boolean if a field has been set.
func (o *ProtocolOptionsHTTPS) HasCertificateConfig() bool {
	if o != nil && !IsNil(o.CertificateConfig) {
		return true
	}

	return false
}

// SetCertificateConfig gets a reference to the given CertificateConfig and assigns it to the CertificateConfig field.
func (o *ProtocolOptionsHTTPS) SetCertificateConfig(v *CertificateConfig) {
	o.CertificateConfig = v
}

func (o ProtocolOptionsHTTPS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertificateConfig) {
		toSerialize["certificateConfig"] = o.CertificateConfig
	}
	return toSerialize, nil
}

type NullableProtocolOptionsHTTPS struct {
	value *ProtocolOptionsHTTPS
	isSet bool
}

func (v NullableProtocolOptionsHTTPS) Get() *ProtocolOptionsHTTPS {
	return v.value
}

func (v *NullableProtocolOptionsHTTPS) Set(val *ProtocolOptionsHTTPS) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolOptionsHTTPS) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolOptionsHTTPS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolOptionsHTTPS(val *ProtocolOptionsHTTPS) *NullableProtocolOptionsHTTPS {
	return &NullableProtocolOptionsHTTPS{value: val, isSet: true}
}

func (v NullableProtocolOptionsHTTPS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolOptionsHTTPS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
