/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
	"time"
)

// checks if the Keypair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Keypair{}

// Keypair Object that represents the public key of an SSH keypair and its name.
type Keypair struct {
	// Date-time when resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Object that represents an SSH keypair MD5 fingerprint.
	Fingerprint *string `json:"fingerprint,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// The name of an SSH keypair. Allowed characters are letters [a-zA-Z], digits [0-9] and the following special characters: [@._-].
	Name *string `json:"name,omitempty"`
	// Object that represents a public SSH key.
	// REQUIRED
	PublicKey *string `json:"publicKey"`
	// Date-time when resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type _Keypair Keypair

// NewKeypair instantiates a new Keypair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeypair(publicKey *string) *Keypair {
	this := Keypair{}
	this.PublicKey = publicKey
	return &this
}

// NewKeypairWithDefaults instantiates a new Keypair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeypairWithDefaults() *Keypair {
	this := Keypair{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Keypair) GetCreatedAt() *time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret *time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keypair) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Keypair) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Keypair) SetCreatedAt(v *time.Time) {
	o.CreatedAt = v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *Keypair) GetFingerprint() *string {
	if o == nil || IsNil(o.Fingerprint) {
		var ret *string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keypair) GetFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *Keypair) HasFingerprint() bool {
	if o != nil && !IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *Keypair) SetFingerprint(v *string) {
	o.Fingerprint = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Keypair) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keypair) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Keypair) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *Keypair) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Keypair) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keypair) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Keypair) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Keypair) SetName(v *string) {
	o.Name = v
}

// GetPublicKey returns the PublicKey field value
func (o *Keypair) GetPublicKey() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *Keypair) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// SetPublicKey sets field value
func (o *Keypair) SetPublicKey(v *string) {
	o.PublicKey = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Keypair) GetUpdatedAt() *time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret *time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Keypair) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Keypair) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Keypair) SetUpdatedAt(v *time.Time) {
	o.UpdatedAt = v
}

func (o Keypair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
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
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableKeypair struct {
	value *Keypair
	isSet bool
}

func (v NullableKeypair) Get() *Keypair {
	return v.value
}

func (v *NullableKeypair) Set(val *Keypair) {
	v.value = val
	v.isSet = true
}

func (v NullableKeypair) IsSet() bool {
	return v.isSet
}

func (v *NullableKeypair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeypair(val *Keypair) *NullableKeypair {
	return &NullableKeypair{value: val, isSet: true}
}

func (v NullableKeypair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeypair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
