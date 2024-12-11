/*
STACKIT Service Enablement API

STACKIT Service Enablement API

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceenablement

import (
	"encoding/json"
)

// checks if the Dependencies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dependencies{}

// Dependencies struct for Dependencies
type Dependencies struct {
	// a list of service IDs which this service depend on. If the service is enabled, those service are enabled as well automatically.
	Hard *[]string `json:"hard,omitempty"`
	// a list of service IDs which this service depend on. When they are disabled a notification is sent.
	Soft *[]string `json:"soft,omitempty"`
}

// NewDependencies instantiates a new Dependencies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDependencies() *Dependencies {
	this := Dependencies{}
	return &this
}

// NewDependenciesWithDefaults instantiates a new Dependencies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDependenciesWithDefaults() *Dependencies {
	this := Dependencies{}
	return &this
}

// GetHard returns the Hard field value if set, zero value otherwise.
func (o *Dependencies) GetHard() *[]string {
	if o == nil || IsNil(o.Hard) {
		var ret *[]string
		return ret
	}
	return o.Hard
}

// GetHardOk returns a tuple with the Hard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependencies) GetHardOk() (*[]string, bool) {
	if o == nil || IsNil(o.Hard) {
		return nil, false
	}
	return o.Hard, true
}

// HasHard returns a boolean if a field has been set.
func (o *Dependencies) HasHard() bool {
	if o != nil && !IsNil(o.Hard) && !IsNil(o.Hard) {
		return true
	}

	return false
}

// SetHard gets a reference to the given []string and assigns it to the Hard field.
func (o *Dependencies) SetHard(v *[]string) {
	o.Hard = v
}

// GetSoft returns the Soft field value if set, zero value otherwise.
func (o *Dependencies) GetSoft() *[]string {
	if o == nil || IsNil(o.Soft) {
		var ret *[]string
		return ret
	}
	return o.Soft
}

// GetSoftOk returns a tuple with the Soft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependencies) GetSoftOk() (*[]string, bool) {
	if o == nil || IsNil(o.Soft) {
		return nil, false
	}
	return o.Soft, true
}

// HasSoft returns a boolean if a field has been set.
func (o *Dependencies) HasSoft() bool {
	if o != nil && !IsNil(o.Soft) && !IsNil(o.Soft) {
		return true
	}

	return false
}

// SetSoft gets a reference to the given []string and assigns it to the Soft field.
func (o *Dependencies) SetSoft(v *[]string) {
	o.Soft = v
}

func (o Dependencies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hard) {
		toSerialize["hard"] = o.Hard
	}
	if !IsNil(o.Soft) {
		toSerialize["soft"] = o.Soft
	}
	return toSerialize, nil
}

type NullableDependencies struct {
	value *Dependencies
	isSet bool
}

func (v NullableDependencies) Get() *Dependencies {
	return v.value
}

func (v *NullableDependencies) Set(val *Dependencies) {
	v.value = val
	v.isSet = true
}

func (v NullableDependencies) IsSet() bool {
	return v.isSet
}

func (v *NullableDependencies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDependencies(val *Dependencies) *NullableDependencies {
	return &NullableDependencies{value: val, isSet: true}
}

func (v NullableDependencies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDependencies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
