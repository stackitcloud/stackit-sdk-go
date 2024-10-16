/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
)

// checks if the CreateKeyPairPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateKeyPairPayload{}

// CreateKeyPairPayload Object that represents the public key of an SSH keypair and its name.
type CreateKeyPairPayload struct {
	// Object that represents an SSH keypair MD5 fingerprint.
	Fingerprint *string `json:"fingerprint,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// The name of an SSH keypair. Allowed characters are letters [a-zA-Z], digits [0-9] and the following special characters: [@._-].
	Name *string `json:"name,omitempty"`
	// Object that represents a public SSH key.
	// REQUIRED
	PublicKey *string `json:"publicKey"`
}

type _CreateKeyPairPayload CreateKeyPairPayload

// NewCreateKeyPairPayload instantiates a new CreateKeyPairPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKeyPairPayload(publicKey *string) *CreateKeyPairPayload {
	this := CreateKeyPairPayload{}
	this.PublicKey = publicKey
	return &this
}

// NewCreateKeyPairPayloadWithDefaults instantiates a new CreateKeyPairPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKeyPairPayloadWithDefaults() *CreateKeyPairPayload {
	this := CreateKeyPairPayload{}
	return &this
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *CreateKeyPairPayload) GetFingerprint() *string {
	if o == nil || IsNil(o.Fingerprint) {
		var ret *string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKeyPairPayload) GetFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *CreateKeyPairPayload) HasFingerprint() bool {
	if o != nil && !IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *CreateKeyPairPayload) SetFingerprint(v *string) {
	o.Fingerprint = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CreateKeyPairPayload) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKeyPairPayload) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CreateKeyPairPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *CreateKeyPairPayload) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateKeyPairPayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKeyPairPayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateKeyPairPayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateKeyPairPayload) SetName(v *string) {
	o.Name = v
}

// GetPublicKey returns the PublicKey field value
func (o *CreateKeyPairPayload) GetPublicKey() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *CreateKeyPairPayload) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// SetPublicKey sets field value
func (o *CreateKeyPairPayload) SetPublicKey(v *string) {
	o.PublicKey = v
}

func (o CreateKeyPairPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["publicKey"] = o.PublicKey
	return toSerialize, nil
}

type NullableCreateKeyPairPayload struct {
	value *CreateKeyPairPayload
	isSet bool
}

func (v NullableCreateKeyPairPayload) Get() *CreateKeyPairPayload {
	return v.value
}

func (v *NullableCreateKeyPairPayload) Set(val *CreateKeyPairPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKeyPairPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKeyPairPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKeyPairPayload(val *CreateKeyPairPayload) *NullableCreateKeyPairPayload {
	return &NullableCreateKeyPairPayload{value: val, isSet: true}
}

func (v NullableCreateKeyPairPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKeyPairPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
