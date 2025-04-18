/*
CDN API

API used to create and manage your CDN distributions.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the ConfigPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigPatch{}

/*
	types and functions for backend
*/

// isModel
type ConfigPatchGetBackendAttributeType = *ConfigPatchBackend
type ConfigPatchGetBackendArgType = ConfigPatchBackend
type ConfigPatchGetBackendRetType = ConfigPatchBackend

func getConfigPatchGetBackendAttributeTypeOk(arg ConfigPatchGetBackendAttributeType) (ret ConfigPatchGetBackendRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setConfigPatchGetBackendAttributeType(arg *ConfigPatchGetBackendAttributeType, val ConfigPatchGetBackendRetType) {
	*arg = &val
}

/*
	types and functions for regions
*/

// isArray
type ConfigPatchGetRegionsAttributeType = *[]Region
type ConfigPatchGetRegionsArgType = []Region
type ConfigPatchGetRegionsRetType = []Region

func getConfigPatchGetRegionsAttributeTypeOk(arg ConfigPatchGetRegionsAttributeType) (ret ConfigPatchGetRegionsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setConfigPatchGetRegionsAttributeType(arg *ConfigPatchGetRegionsAttributeType, val ConfigPatchGetRegionsRetType) {
	*arg = &val
}

// ConfigPatch struct for ConfigPatch
type ConfigPatch struct {
	Backend ConfigPatchGetBackendAttributeType `json:"backend,omitempty"`
	Regions ConfigPatchGetRegionsAttributeType `json:"regions,omitempty"`
}

// NewConfigPatch instantiates a new ConfigPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigPatch() *ConfigPatch {
	this := ConfigPatch{}
	return &this
}

// NewConfigPatchWithDefaults instantiates a new ConfigPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigPatchWithDefaults() *ConfigPatch {
	this := ConfigPatch{}
	return &this
}

// GetBackend returns the Backend field value if set, zero value otherwise.
func (o *ConfigPatch) GetBackend() (res ConfigPatchGetBackendRetType) {
	res, _ = o.GetBackendOk()
	return
}

// GetBackendOk returns a tuple with the Backend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPatch) GetBackendOk() (ret ConfigPatchGetBackendRetType, ok bool) {
	return getConfigPatchGetBackendAttributeTypeOk(o.Backend)
}

// HasBackend returns a boolean if a field has been set.
func (o *ConfigPatch) HasBackend() bool {
	_, ok := o.GetBackendOk()
	return ok
}

// SetBackend gets a reference to the given ConfigPatchBackend and assigns it to the Backend field.
func (o *ConfigPatch) SetBackend(v ConfigPatchGetBackendRetType) {
	setConfigPatchGetBackendAttributeType(&o.Backend, v)
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *ConfigPatch) GetRegions() (res ConfigPatchGetRegionsRetType) {
	res, _ = o.GetRegionsOk()
	return
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigPatch) GetRegionsOk() (ret ConfigPatchGetRegionsRetType, ok bool) {
	return getConfigPatchGetRegionsAttributeTypeOk(o.Regions)
}

// HasRegions returns a boolean if a field has been set.
func (o *ConfigPatch) HasRegions() bool {
	_, ok := o.GetRegionsOk()
	return ok
}

// SetRegions gets a reference to the given []Region and assigns it to the Regions field.
func (o *ConfigPatch) SetRegions(v ConfigPatchGetRegionsRetType) {
	setConfigPatchGetRegionsAttributeType(&o.Regions, v)
}

func (o ConfigPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getConfigPatchGetBackendAttributeTypeOk(o.Backend); ok {
		toSerialize["Backend"] = val
	}
	if val, ok := getConfigPatchGetRegionsAttributeTypeOk(o.Regions); ok {
		toSerialize["Regions"] = val
	}
	return toSerialize, nil
}

type NullableConfigPatch struct {
	value *ConfigPatch
	isSet bool
}

func (v NullableConfigPatch) Get() *ConfigPatch {
	return v.value
}

func (v *NullableConfigPatch) Set(val *ConfigPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigPatch(val *ConfigPatch) *NullableConfigPatch {
	return &NullableConfigPatch{value: val, isSet: true}
}

func (v NullableConfigPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
