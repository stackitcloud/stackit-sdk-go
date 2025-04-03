/*
STACKIT Key Management Service API

This API provides endpoints for managing keys and key rings.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kms

import (
	"encoding/json"
	"time"
)

// checks if the Version type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Version{}

/*
	types and functions for createdAt
*/

// isDateTime
type VersionGetCreatedAtAttributeType = *time.Time
type VersionGetCreatedAtArgType = time.Time
type VersionGetCreatedAtRetType = time.Time

func getVersionGetCreatedAtAttributeTypeOk(arg VersionGetCreatedAtAttributeType) (ret VersionGetCreatedAtRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setVersionGetCreatedAtAttributeType(arg *VersionGetCreatedAtAttributeType, val VersionGetCreatedAtRetType) {
	*arg = &val
}

/*
	types and functions for disabled
*/

// isBoolean
type VersiongetDisabledAttributeType = *bool
type VersiongetDisabledArgType = bool
type VersiongetDisabledRetType = bool

func getVersiongetDisabledAttributeTypeOk(arg VersiongetDisabledAttributeType) (ret VersiongetDisabledRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setVersiongetDisabledAttributeType(arg *VersiongetDisabledAttributeType, val VersiongetDisabledRetType) {
	*arg = &val
}

/*
	types and functions for keyId
*/

// isNotNullableString
type VersionGetKeyIdAttributeType = *string

func getVersionGetKeyIdAttributeTypeOk(arg VersionGetKeyIdAttributeType) (ret VersionGetKeyIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setVersionGetKeyIdAttributeType(arg *VersionGetKeyIdAttributeType, val VersionGetKeyIdRetType) {
	*arg = &val
}

type VersionGetKeyIdArgType = string
type VersionGetKeyIdRetType = string

/*
	types and functions for keyRingId
*/

// isNotNullableString
type VersionGetKeyRingIdAttributeType = *string

func getVersionGetKeyRingIdAttributeTypeOk(arg VersionGetKeyRingIdAttributeType) (ret VersionGetKeyRingIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setVersionGetKeyRingIdAttributeType(arg *VersionGetKeyRingIdAttributeType, val VersionGetKeyRingIdRetType) {
	*arg = &val
}

type VersionGetKeyRingIdArgType = string
type VersionGetKeyRingIdRetType = string

/*
	types and functions for number
*/

// isLong
type VersionGetNumberAttributeType = *int64
type VersionGetNumberArgType = int64
type VersionGetNumberRetType = int64

func getVersionGetNumberAttributeTypeOk(arg VersionGetNumberAttributeType) (ret VersionGetNumberRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setVersionGetNumberAttributeType(arg *VersionGetNumberAttributeType, val VersionGetNumberRetType) {
	*arg = &val
}

/*
	types and functions for publicKey
*/

// isNotNullableString
type VersionGetPublicKeyAttributeType = *string

func getVersionGetPublicKeyAttributeTypeOk(arg VersionGetPublicKeyAttributeType) (ret VersionGetPublicKeyRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setVersionGetPublicKeyAttributeType(arg *VersionGetPublicKeyAttributeType, val VersionGetPublicKeyRetType) {
	*arg = &val
}

type VersionGetPublicKeyArgType = string
type VersionGetPublicKeyRetType = string

/*
	types and functions for state
*/

// isEnumRef
type VersionGetStateAttributeType = *string
type VersionGetStateArgType = string
type VersionGetStateRetType = string

func getVersionGetStateAttributeTypeOk(arg VersionGetStateAttributeType) (ret VersionGetStateRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setVersionGetStateAttributeType(arg *VersionGetStateAttributeType, val VersionGetStateRetType) {
	*arg = &val
}

// Version struct for Version
type Version struct {
	// The date and time the creation of the key was triggered.
	// REQUIRED
	CreatedAt VersionGetCreatedAtAttributeType `json:"createdAt"`
	// States whether versions is enabled or disabled.
	Disabled VersiongetDisabledAttributeType `json:"disabled,omitempty"`
	// The unique id of the key this version is assigned to.
	// REQUIRED
	KeyId VersionGetKeyIdAttributeType `json:"keyId"`
	// The unique id of the key ring the key of this version is assigned to.
	// REQUIRED
	KeyRingId VersionGetKeyRingIdAttributeType `json:"keyRingId"`
	// A sequential number which identifies the key versions.
	// REQUIRED
	Number VersionGetNumberAttributeType `json:"number"`
	// The public key of the key version. Only present in asymmetric keys.
	PublicKey VersionGetPublicKeyAttributeType `json:"publicKey,omitempty"`
	// The current state of the key.
	// REQUIRED
	State VersionGetStateAttributeType `json:"state"`
}

type _Version Version

// NewVersion instantiates a new Version object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersion(createdAt VersionGetCreatedAtArgType, keyId VersionGetKeyIdArgType, keyRingId VersionGetKeyRingIdArgType, number VersionGetNumberArgType, state VersionGetStateArgType) *Version {
	this := Version{}
	setVersionGetCreatedAtAttributeType(&this.CreatedAt, createdAt)
	setVersionGetKeyIdAttributeType(&this.KeyId, keyId)
	setVersionGetKeyRingIdAttributeType(&this.KeyRingId, keyRingId)
	setVersionGetNumberAttributeType(&this.Number, number)
	setVersionGetStateAttributeType(&this.State, state)
	return &this
}

// NewVersionWithDefaults instantiates a new Version object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionWithDefaults() *Version {
	this := Version{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *Version) GetCreatedAt() (ret VersionGetCreatedAtRetType) {
	ret, _ = o.GetCreatedAtOk()
	return ret
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Version) GetCreatedAtOk() (ret VersionGetCreatedAtRetType, ok bool) {
	return getVersionGetCreatedAtAttributeTypeOk(o.CreatedAt)
}

// SetCreatedAt sets field value
func (o *Version) SetCreatedAt(v VersionGetCreatedAtRetType) {
	setVersionGetCreatedAtAttributeType(&o.CreatedAt, v)
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *Version) GetDisabled() (res VersiongetDisabledRetType) {
	res, _ = o.GetDisabledOk()
	return
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetDisabledOk() (ret VersiongetDisabledRetType, ok bool) {
	return getVersiongetDisabledAttributeTypeOk(o.Disabled)
}

// HasDisabled returns a boolean if a field has been set.
func (o *Version) HasDisabled() bool {
	_, ok := o.GetDisabledOk()
	return ok
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *Version) SetDisabled(v VersiongetDisabledRetType) {
	setVersiongetDisabledAttributeType(&o.Disabled, v)
}

// GetKeyId returns the KeyId field value
func (o *Version) GetKeyId() (ret VersionGetKeyIdRetType) {
	ret, _ = o.GetKeyIdOk()
	return ret
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *Version) GetKeyIdOk() (ret VersionGetKeyIdRetType, ok bool) {
	return getVersionGetKeyIdAttributeTypeOk(o.KeyId)
}

// SetKeyId sets field value
func (o *Version) SetKeyId(v VersionGetKeyIdRetType) {
	setVersionGetKeyIdAttributeType(&o.KeyId, v)
}

// GetKeyRingId returns the KeyRingId field value
func (o *Version) GetKeyRingId() (ret VersionGetKeyRingIdRetType) {
	ret, _ = o.GetKeyRingIdOk()
	return ret
}

// GetKeyRingIdOk returns a tuple with the KeyRingId field value
// and a boolean to check if the value has been set.
func (o *Version) GetKeyRingIdOk() (ret VersionGetKeyRingIdRetType, ok bool) {
	return getVersionGetKeyRingIdAttributeTypeOk(o.KeyRingId)
}

// SetKeyRingId sets field value
func (o *Version) SetKeyRingId(v VersionGetKeyRingIdRetType) {
	setVersionGetKeyRingIdAttributeType(&o.KeyRingId, v)
}

// GetNumber returns the Number field value
func (o *Version) GetNumber() (ret VersionGetNumberRetType) {
	ret, _ = o.GetNumberOk()
	return ret
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *Version) GetNumberOk() (ret VersionGetNumberRetType, ok bool) {
	return getVersionGetNumberAttributeTypeOk(o.Number)
}

// SetNumber sets field value
func (o *Version) SetNumber(v VersionGetNumberRetType) {
	setVersionGetNumberAttributeType(&o.Number, v)
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *Version) GetPublicKey() (res VersionGetPublicKeyRetType) {
	res, _ = o.GetPublicKeyOk()
	return
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetPublicKeyOk() (ret VersionGetPublicKeyRetType, ok bool) {
	return getVersionGetPublicKeyAttributeTypeOk(o.PublicKey)
}

// HasPublicKey returns a boolean if a field has been set.
func (o *Version) HasPublicKey() bool {
	_, ok := o.GetPublicKeyOk()
	return ok
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *Version) SetPublicKey(v VersionGetPublicKeyRetType) {
	setVersionGetPublicKeyAttributeType(&o.PublicKey, v)
}

// GetState returns the State field value
func (o *Version) GetState() (ret VersionGetStateRetType) {
	ret, _ = o.GetStateOk()
	return ret
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Version) GetStateOk() (ret VersionGetStateRetType, ok bool) {
	return getVersionGetStateAttributeTypeOk(o.State)
}

// SetState sets field value
func (o *Version) SetState(v VersionGetStateRetType) {
	setVersionGetStateAttributeType(&o.State, v)
}

func (o Version) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getVersionGetCreatedAtAttributeTypeOk(o.CreatedAt); ok {
		toSerialize["CreatedAt"] = val
	}
	if val, ok := getVersiongetDisabledAttributeTypeOk(o.Disabled); ok {
		toSerialize["Disabled"] = val
	}
	if val, ok := getVersionGetKeyIdAttributeTypeOk(o.KeyId); ok {
		toSerialize["KeyId"] = val
	}
	if val, ok := getVersionGetKeyRingIdAttributeTypeOk(o.KeyRingId); ok {
		toSerialize["KeyRingId"] = val
	}
	if val, ok := getVersionGetNumberAttributeTypeOk(o.Number); ok {
		toSerialize["Number"] = val
	}
	if val, ok := getVersionGetPublicKeyAttributeTypeOk(o.PublicKey); ok {
		toSerialize["PublicKey"] = val
	}
	if val, ok := getVersionGetStateAttributeTypeOk(o.State); ok {
		toSerialize["State"] = val
	}
	return toSerialize, nil
}

type NullableVersion struct {
	value *Version
	isSet bool
}

func (v NullableVersion) Get() *Version {
	return v.value
}

func (v *NullableVersion) Set(val *Version) {
	v.value = val
	v.isSet = true
}

func (v NullableVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersion(val *Version) *NullableVersion {
	return &NullableVersion{value: val, isSet: true}
}

func (v NullableVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
