/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"encoding/json"
)

// checks if the PermissionDenied type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionDenied{}

// PermissionDenied struct for PermissionDenied
type PermissionDenied struct {
	// REQUIRED
	Detail *string `json:"detail"`
}

type _PermissionDenied PermissionDenied

// NewPermissionDenied instantiates a new PermissionDenied object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionDenied(detail *string) *PermissionDenied {
	this := PermissionDenied{}
	this.Detail = detail
	return &this
}

// NewPermissionDeniedWithDefaults instantiates a new PermissionDenied object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionDeniedWithDefaults() *PermissionDenied {
	this := PermissionDenied{}
	return &this
}

// GetDetail returns the Detail field value
func (o *PermissionDenied) GetDetail() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *PermissionDenied) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail, true
}

// SetDetail sets field value
func (o *PermissionDenied) SetDetail(v *string) {
	o.Detail = v
}

func (o PermissionDenied) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detail"] = o.Detail
	return toSerialize, nil
}

type NullablePermissionDenied struct {
	value *PermissionDenied
	isSet bool
}

func (v NullablePermissionDenied) Get() *PermissionDenied {
	return v.value
}

func (v *NullablePermissionDenied) Set(val *PermissionDenied) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionDenied) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionDenied) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionDenied(val *PermissionDenied) *NullablePermissionDenied {
	return &NullablePermissionDenied{value: val, isSet: true}
}

func (v NullablePermissionDenied) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionDenied) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
