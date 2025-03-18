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

// checks if the BackupSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupSource{}

/*
	types and functions for id
*/

// isNotNullableString
type BackupSourceGetIdAttributeType = *string

func getBackupSourceGetIdAttributeTypeOk(arg BackupSourceGetIdAttributeType) (ret BackupSourceGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupSourceGetIdAttributeType(arg *BackupSourceGetIdAttributeType, val BackupSourceGetIdRetType) {
	*arg = &val
}

type BackupSourceGetIdArgType = string
type BackupSourceGetIdRetType = string

/*
	types and functions for type
*/

// isNotNullableString
type BackupSourceGetTypeAttributeType = *string

func getBackupSourceGetTypeAttributeTypeOk(arg BackupSourceGetTypeAttributeType) (ret BackupSourceGetTypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setBackupSourceGetTypeAttributeType(arg *BackupSourceGetTypeAttributeType, val BackupSourceGetTypeRetType) {
	*arg = &val
}

type BackupSourceGetTypeArgType = string
type BackupSourceGetTypeRetType = string

// BackupSource The source object of a backup.
type BackupSource struct {
	// Universally Unique Identifier (UUID).
	// REQUIRED
	Id BackupSourceGetIdAttributeType `json:"id"`
	// The source types of a backup. Possible values: `volume`, `snapshot`.
	// REQUIRED
	Type BackupSourceGetTypeAttributeType `json:"type"`
}

type _BackupSource BackupSource

// NewBackupSource instantiates a new BackupSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupSource(id BackupSourceGetIdArgType, type_ BackupSourceGetTypeArgType) *BackupSource {
	this := BackupSource{}
	setBackupSourceGetIdAttributeType(&this.Id, id)
	setBackupSourceGetTypeAttributeType(&this.Type, type_)
	return &this
}

// NewBackupSourceWithDefaults instantiates a new BackupSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupSourceWithDefaults() *BackupSource {
	this := BackupSource{}
	return &this
}

// GetId returns the Id field value
func (o *BackupSource) GetId() (ret BackupSourceGetIdRetType) {
	ret, _ = o.GetIdOk()
	return ret
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BackupSource) GetIdOk() (ret BackupSourceGetIdRetType, ok bool) {
	return getBackupSourceGetIdAttributeTypeOk(o.Id)
}

// SetId sets field value
func (o *BackupSource) SetId(v BackupSourceGetIdRetType) {
	setBackupSourceGetIdAttributeType(&o.Id, v)
}

// GetType returns the Type field value
func (o *BackupSource) GetType() (ret BackupSourceGetTypeRetType) {
	ret, _ = o.GetTypeOk()
	return ret
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BackupSource) GetTypeOk() (ret BackupSourceGetTypeRetType, ok bool) {
	return getBackupSourceGetTypeAttributeTypeOk(o.Type)
}

// SetType sets field value
func (o *BackupSource) SetType(v BackupSourceGetTypeRetType) {
	setBackupSourceGetTypeAttributeType(&o.Type, v)
}

func (o BackupSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getBackupSourceGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getBackupSourceGetTypeAttributeTypeOk(o.Type); ok {
		toSerialize["Type"] = val
	}
	return toSerialize, nil
}

type NullableBackupSource struct {
	value *BackupSource
	isSet bool
}

func (v NullableBackupSource) Get() *BackupSource {
	return v.value
}

func (v *NullableBackupSource) Set(val *BackupSource) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupSource) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupSource(val *BackupSource) *NullableBackupSource {
	return &NullableBackupSource{value: val, isSet: true}
}

func (v NullableBackupSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
