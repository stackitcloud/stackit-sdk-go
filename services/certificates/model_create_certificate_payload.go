/*
Load Balancer Certificates API

This API offers the ability to store TLS certificates, which can be used by load balancing servers in STACKIT. They can be between consumer and load balancing server and/or between load balancing server and endpoint server.

API version: 2beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package certificates

import (
	"encoding/json"
)

// checks if the CreateCertificatePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCertificatePayload{}

// CreateCertificatePayload Uploads a PEM encoded X509 public/private key pair
type CreateCertificatePayload struct {
	// TLS certificate name
	Name *string `json:"name,omitempty"`
	// The PEM encoded private key part
	PrivateKey *string `json:"privateKey,omitempty"`
	// Project identifier
	ProjectId *string `json:"projectId,omitempty"`
	// The PEM encoded public key part
	PublicKey *string `json:"publicKey,omitempty"`
	// Region
	Region *string `json:"region,omitempty"`
}

// NewCreateCertificatePayload instantiates a new CreateCertificatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCertificatePayload() *CreateCertificatePayload {
	this := CreateCertificatePayload{}
	return &this
}

// NewCreateCertificatePayloadWithDefaults instantiates a new CreateCertificatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCertificatePayloadWithDefaults() *CreateCertificatePayload {
	this := CreateCertificatePayload{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateCertificatePayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificatePayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateCertificatePayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateCertificatePayload) SetName(v *string) {
	o.Name = v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *CreateCertificatePayload) GetPrivateKey() *string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret *string
		return ret
	}
	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificatePayload) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *CreateCertificatePayload) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *CreateCertificatePayload) SetPrivateKey(v *string) {
	o.PrivateKey = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CreateCertificatePayload) GetProjectId() *string {
	if o == nil || IsNil(o.ProjectId) {
		var ret *string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificatePayload) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CreateCertificatePayload) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *CreateCertificatePayload) SetProjectId(v *string) {
	o.ProjectId = v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *CreateCertificatePayload) GetPublicKey() *string {
	if o == nil || IsNil(o.PublicKey) {
		var ret *string
		return ret
	}
	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificatePayload) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *CreateCertificatePayload) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *CreateCertificatePayload) SetPublicKey(v *string) {
	o.PublicKey = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CreateCertificatePayload) GetRegion() *string {
	if o == nil || IsNil(o.Region) {
		var ret *string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificatePayload) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CreateCertificatePayload) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CreateCertificatePayload) SetRegion(v *string) {
	o.Region = v
}

func (o CreateCertificatePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PrivateKey) {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}

type NullableCreateCertificatePayload struct {
	value *CreateCertificatePayload
	isSet bool
}

func (v NullableCreateCertificatePayload) Get() *CreateCertificatePayload {
	return v.value
}

func (v *NullableCreateCertificatePayload) Set(val *CreateCertificatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCertificatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCertificatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCertificatePayload(val *CreateCertificatePayload) *NullableCreateCertificatePayload {
	return &NullableCreateCertificatePayload{value: val, isSet: true}
}

func (v NullableCreateCertificatePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCertificatePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
