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

/*
	types and functions for name
*/

// isNotNullableString
type CreateCertificatePayloadGetNameAttributeType = *string

func getCreateCertificatePayloadGetNameAttributeTypeOk(arg CreateCertificatePayloadGetNameAttributeType) (ret CreateCertificatePayloadGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateCertificatePayloadGetNameAttributeType(arg *CreateCertificatePayloadGetNameAttributeType, val CreateCertificatePayloadGetNameRetType) {
	*arg = &val
}

type CreateCertificatePayloadGetNameArgType = string
type CreateCertificatePayloadGetNameRetType = string

/*
	types and functions for privateKey
*/

// isNotNullableString
type CreateCertificatePayloadGetPrivateKeyAttributeType = *string

func getCreateCertificatePayloadGetPrivateKeyAttributeTypeOk(arg CreateCertificatePayloadGetPrivateKeyAttributeType) (ret CreateCertificatePayloadGetPrivateKeyRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateCertificatePayloadGetPrivateKeyAttributeType(arg *CreateCertificatePayloadGetPrivateKeyAttributeType, val CreateCertificatePayloadGetPrivateKeyRetType) {
	*arg = &val
}

type CreateCertificatePayloadGetPrivateKeyArgType = string
type CreateCertificatePayloadGetPrivateKeyRetType = string

/*
	types and functions for projectId
*/

// isNotNullableString
type CreateCertificatePayloadGetProjectIdAttributeType = *string

func getCreateCertificatePayloadGetProjectIdAttributeTypeOk(arg CreateCertificatePayloadGetProjectIdAttributeType) (ret CreateCertificatePayloadGetProjectIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateCertificatePayloadGetProjectIdAttributeType(arg *CreateCertificatePayloadGetProjectIdAttributeType, val CreateCertificatePayloadGetProjectIdRetType) {
	*arg = &val
}

type CreateCertificatePayloadGetProjectIdArgType = string
type CreateCertificatePayloadGetProjectIdRetType = string

/*
	types and functions for publicKey
*/

// isNotNullableString
type CreateCertificatePayloadGetPublicKeyAttributeType = *string

func getCreateCertificatePayloadGetPublicKeyAttributeTypeOk(arg CreateCertificatePayloadGetPublicKeyAttributeType) (ret CreateCertificatePayloadGetPublicKeyRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateCertificatePayloadGetPublicKeyAttributeType(arg *CreateCertificatePayloadGetPublicKeyAttributeType, val CreateCertificatePayloadGetPublicKeyRetType) {
	*arg = &val
}

type CreateCertificatePayloadGetPublicKeyArgType = string
type CreateCertificatePayloadGetPublicKeyRetType = string

/*
	types and functions for region
*/

// isNotNullableString
type CreateCertificatePayloadGetRegionAttributeType = *string

func getCreateCertificatePayloadGetRegionAttributeTypeOk(arg CreateCertificatePayloadGetRegionAttributeType) (ret CreateCertificatePayloadGetRegionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateCertificatePayloadGetRegionAttributeType(arg *CreateCertificatePayloadGetRegionAttributeType, val CreateCertificatePayloadGetRegionRetType) {
	*arg = &val
}

type CreateCertificatePayloadGetRegionArgType = string
type CreateCertificatePayloadGetRegionRetType = string

// CreateCertificatePayload Uploads a PEM encoded X509 public/private key pair
type CreateCertificatePayload struct {
	// TLS certificate name
	Name CreateCertificatePayloadGetNameAttributeType `json:"name,omitempty"`
	// The PEM encoded private key part
	PrivateKey CreateCertificatePayloadGetPrivateKeyAttributeType `json:"privateKey,omitempty"`
	// Project identifier
	ProjectId CreateCertificatePayloadGetProjectIdAttributeType `json:"projectId,omitempty"`
	// The PEM encoded public key part
	PublicKey CreateCertificatePayloadGetPublicKeyAttributeType `json:"publicKey,omitempty"`
	// Region
	Region CreateCertificatePayloadGetRegionAttributeType `json:"region,omitempty"`
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
func (o *CreateCertificatePayload) GetName() (res CreateCertificatePayloadGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificatePayload) GetNameOk() (ret CreateCertificatePayloadGetNameRetType, ok bool) {
	return getCreateCertificatePayloadGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *CreateCertificatePayload) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateCertificatePayload) SetName(v CreateCertificatePayloadGetNameRetType) {
	setCreateCertificatePayloadGetNameAttributeType(&o.Name, v)
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *CreateCertificatePayload) GetPrivateKey() (res CreateCertificatePayloadGetPrivateKeyRetType) {
	res, _ = o.GetPrivateKeyOk()
	return
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificatePayload) GetPrivateKeyOk() (ret CreateCertificatePayloadGetPrivateKeyRetType, ok bool) {
	return getCreateCertificatePayloadGetPrivateKeyAttributeTypeOk(o.PrivateKey)
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *CreateCertificatePayload) HasPrivateKey() bool {
	_, ok := o.GetPrivateKeyOk()
	return ok
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *CreateCertificatePayload) SetPrivateKey(v CreateCertificatePayloadGetPrivateKeyRetType) {
	setCreateCertificatePayloadGetPrivateKeyAttributeType(&o.PrivateKey, v)
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CreateCertificatePayload) GetProjectId() (res CreateCertificatePayloadGetProjectIdRetType) {
	res, _ = o.GetProjectIdOk()
	return
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificatePayload) GetProjectIdOk() (ret CreateCertificatePayloadGetProjectIdRetType, ok bool) {
	return getCreateCertificatePayloadGetProjectIdAttributeTypeOk(o.ProjectId)
}

// HasProjectId returns a boolean if a field has been set.
func (o *CreateCertificatePayload) HasProjectId() bool {
	_, ok := o.GetProjectIdOk()
	return ok
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *CreateCertificatePayload) SetProjectId(v CreateCertificatePayloadGetProjectIdRetType) {
	setCreateCertificatePayloadGetProjectIdAttributeType(&o.ProjectId, v)
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *CreateCertificatePayload) GetPublicKey() (res CreateCertificatePayloadGetPublicKeyRetType) {
	res, _ = o.GetPublicKeyOk()
	return
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificatePayload) GetPublicKeyOk() (ret CreateCertificatePayloadGetPublicKeyRetType, ok bool) {
	return getCreateCertificatePayloadGetPublicKeyAttributeTypeOk(o.PublicKey)
}

// HasPublicKey returns a boolean if a field has been set.
func (o *CreateCertificatePayload) HasPublicKey() bool {
	_, ok := o.GetPublicKeyOk()
	return ok
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *CreateCertificatePayload) SetPublicKey(v CreateCertificatePayloadGetPublicKeyRetType) {
	setCreateCertificatePayloadGetPublicKeyAttributeType(&o.PublicKey, v)
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CreateCertificatePayload) GetRegion() (res CreateCertificatePayloadGetRegionRetType) {
	res, _ = o.GetRegionOk()
	return
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCertificatePayload) GetRegionOk() (ret CreateCertificatePayloadGetRegionRetType, ok bool) {
	return getCreateCertificatePayloadGetRegionAttributeTypeOk(o.Region)
}

// HasRegion returns a boolean if a field has been set.
func (o *CreateCertificatePayload) HasRegion() bool {
	_, ok := o.GetRegionOk()
	return ok
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CreateCertificatePayload) SetRegion(v CreateCertificatePayloadGetRegionRetType) {
	setCreateCertificatePayloadGetRegionAttributeType(&o.Region, v)
}

func (o CreateCertificatePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateCertificatePayloadGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getCreateCertificatePayloadGetPrivateKeyAttributeTypeOk(o.PrivateKey); ok {
		toSerialize["PrivateKey"] = val
	}
	if val, ok := getCreateCertificatePayloadGetProjectIdAttributeTypeOk(o.ProjectId); ok {
		toSerialize["ProjectId"] = val
	}
	if val, ok := getCreateCertificatePayloadGetPublicKeyAttributeTypeOk(o.PublicKey); ok {
		toSerialize["PublicKey"] = val
	}
	if val, ok := getCreateCertificatePayloadGetRegionAttributeTypeOk(o.Region); ok {
		toSerialize["Region"] = val
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
