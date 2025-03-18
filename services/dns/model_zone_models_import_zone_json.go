/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the ZoneModelsImportZoneJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneModelsImportZoneJson{}

/*
	types and functions for rrSets
*/

// isArray
type ZoneModelsImportZoneJsonGetRrSetsAttributeType = *[]ZoneModelsImportRecordModel
type ZoneModelsImportZoneJsonGetRrSetsArgType = []ZoneModelsImportRecordModel
type ZoneModelsImportZoneJsonGetRrSetsRetType = []ZoneModelsImportRecordModel

func getZoneModelsImportZoneJsonGetRrSetsAttributeTypeOk(arg ZoneModelsImportZoneJsonGetRrSetsAttributeType) (ret ZoneModelsImportZoneJsonGetRrSetsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setZoneModelsImportZoneJsonGetRrSetsAttributeType(arg *ZoneModelsImportZoneJsonGetRrSetsAttributeType, val ZoneModelsImportZoneJsonGetRrSetsRetType) {
	*arg = &val
}

// ZoneModelsImportZoneJson struct for ZoneModelsImportZoneJson
type ZoneModelsImportZoneJson struct {
	RrSets ZoneModelsImportZoneJsonGetRrSetsAttributeType `json:"rrSets,omitempty"`
}

// NewZoneModelsImportZoneJson instantiates a new ZoneModelsImportZoneJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneModelsImportZoneJson() *ZoneModelsImportZoneJson {
	this := ZoneModelsImportZoneJson{}
	return &this
}

// NewZoneModelsImportZoneJsonWithDefaults instantiates a new ZoneModelsImportZoneJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneModelsImportZoneJsonWithDefaults() *ZoneModelsImportZoneJson {
	this := ZoneModelsImportZoneJson{}
	return &this
}

// GetRrSets returns the RrSets field value if set, zero value otherwise.
func (o *ZoneModelsImportZoneJson) GetRrSets() (res ZoneModelsImportZoneJsonGetRrSetsRetType) {
	res, _ = o.GetRrSetsOk()
	return
}

// GetRrSetsOk returns a tuple with the RrSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneModelsImportZoneJson) GetRrSetsOk() (ret ZoneModelsImportZoneJsonGetRrSetsRetType, ok bool) {
	return getZoneModelsImportZoneJsonGetRrSetsAttributeTypeOk(o.RrSets)
}

// HasRrSets returns a boolean if a field has been set.
func (o *ZoneModelsImportZoneJson) HasRrSets() bool {
	_, ok := o.GetRrSetsOk()
	return ok
}

// SetRrSets gets a reference to the given []ZoneModelsImportRecordModel and assigns it to the RrSets field.
func (o *ZoneModelsImportZoneJson) SetRrSets(v ZoneModelsImportZoneJsonGetRrSetsRetType) {
	setZoneModelsImportZoneJsonGetRrSetsAttributeType(&o.RrSets, v)
}

func (o ZoneModelsImportZoneJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getZoneModelsImportZoneJsonGetRrSetsAttributeTypeOk(o.RrSets); ok {
		toSerialize["RrSets"] = val
	}
	return toSerialize, nil
}

type NullableZoneModelsImportZoneJson struct {
	value *ZoneModelsImportZoneJson
	isSet bool
}

func (v NullableZoneModelsImportZoneJson) Get() *ZoneModelsImportZoneJson {
	return v.value
}

func (v *NullableZoneModelsImportZoneJson) Set(val *ZoneModelsImportZoneJson) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneModelsImportZoneJson) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneModelsImportZoneJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneModelsImportZoneJson(val *ZoneModelsImportZoneJson) *NullableZoneModelsImportZoneJson {
	return &NullableZoneModelsImportZoneJson{value: val, isSet: true}
}

func (v NullableZoneModelsImportZoneJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneModelsImportZoneJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
