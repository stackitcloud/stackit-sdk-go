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

// checks if the ImageShare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageShare{}

// ImageShare Share details of an Image. ParentOrganization and Projects are mutually exclusive.
type ImageShare struct {
	ParentOrganization *bool     `json:"parentOrganization,omitempty"`
	Projects           *[]string `json:"projects,omitempty"`
}

// NewImageShare instantiates a new ImageShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageShare() *ImageShare {
	this := ImageShare{}
	return &this
}

// NewImageShareWithDefaults instantiates a new ImageShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageShareWithDefaults() *ImageShare {
	this := ImageShare{}
	return &this
}

// GetParentOrganization returns the ParentOrganization field value if set, zero value otherwise.
func (o *ImageShare) GetParentOrganization() *bool {
	if o == nil || IsNil(o.ParentOrganization) {
		var ret *bool
		return ret
	}
	return o.ParentOrganization
}

// GetParentOrganizationOk returns a tuple with the ParentOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageShare) GetParentOrganizationOk() (*bool, bool) {
	if o == nil || IsNil(o.ParentOrganization) {
		return nil, false
	}
	return o.ParentOrganization, true
}

// HasParentOrganization returns a boolean if a field has been set.
func (o *ImageShare) HasParentOrganization() bool {
	if o != nil && !IsNil(o.ParentOrganization) {
		return true
	}

	return false
}

// SetParentOrganization gets a reference to the given bool and assigns it to the ParentOrganization field.
func (o *ImageShare) SetParentOrganization(v *bool) {
	o.ParentOrganization = v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *ImageShare) GetProjects() *[]string {
	if o == nil || IsNil(o.Projects) {
		var ret *[]string
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageShare) GetProjectsOk() (*[]string, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *ImageShare) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []string and assigns it to the Projects field.
func (o *ImageShare) SetProjects(v *[]string) {
	o.Projects = v
}

func (o ImageShare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParentOrganization) {
		toSerialize["parentOrganization"] = o.ParentOrganization
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	return toSerialize, nil
}

type NullableImageShare struct {
	value *ImageShare
	isSet bool
}

func (v NullableImageShare) Get() *ImageShare {
	return v.value
}

func (v *NullableImageShare) Set(val *ImageShare) {
	v.value = val
	v.isSet = true
}

func (v NullableImageShare) IsSet() bool {
	return v.isSet
}

func (v *NullableImageShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageShare(val *ImageShare) *NullableImageShare {
	return &NullableImageShare{value: val, isSet: true}
}

func (v NullableImageShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}