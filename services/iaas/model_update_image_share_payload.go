/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the UpdateImageSharePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateImageSharePayload{}

/*
	types and functions for parentOrganization
*/

// isBoolean
type UpdateImageSharePayloadgetParentOrganizationAttributeType = *bool
type UpdateImageSharePayloadgetParentOrganizationArgType = bool
type UpdateImageSharePayloadgetParentOrganizationRetType = bool

func getUpdateImageSharePayloadgetParentOrganizationAttributeTypeOk(arg UpdateImageSharePayloadgetParentOrganizationAttributeType) (ret UpdateImageSharePayloadgetParentOrganizationRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateImageSharePayloadgetParentOrganizationAttributeType(arg *UpdateImageSharePayloadgetParentOrganizationAttributeType, val UpdateImageSharePayloadgetParentOrganizationRetType) {
	*arg = &val
}

/*
	types and functions for projects
*/

// isArray
type UpdateImageSharePayloadGetProjectsAttributeType = *[]string
type UpdateImageSharePayloadGetProjectsArgType = []string
type UpdateImageSharePayloadGetProjectsRetType = []string

func getUpdateImageSharePayloadGetProjectsAttributeTypeOk(arg UpdateImageSharePayloadGetProjectsAttributeType) (ret UpdateImageSharePayloadGetProjectsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateImageSharePayloadGetProjectsAttributeType(arg *UpdateImageSharePayloadGetProjectsAttributeType, val UpdateImageSharePayloadGetProjectsRetType) {
	*arg = &val
}

// UpdateImageSharePayload Share details of an Image. For requests ParentOrganization and Projects are mutually exclusive.
type UpdateImageSharePayload struct {
	// Image is shared with all projects inside the image owners organization.
	ParentOrganization UpdateImageSharePayloadgetParentOrganizationAttributeType `json:"parentOrganization,omitempty"`
	// List of all projects the Image is shared with.
	Projects UpdateImageSharePayloadGetProjectsAttributeType `json:"projects,omitempty"`
}

// NewUpdateImageSharePayload instantiates a new UpdateImageSharePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateImageSharePayload() *UpdateImageSharePayload {
	this := UpdateImageSharePayload{}
	return &this
}

// NewUpdateImageSharePayloadWithDefaults instantiates a new UpdateImageSharePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateImageSharePayloadWithDefaults() *UpdateImageSharePayload {
	this := UpdateImageSharePayload{}
	return &this
}

// GetParentOrganization returns the ParentOrganization field value if set, zero value otherwise.
func (o *UpdateImageSharePayload) GetParentOrganization() (res UpdateImageSharePayloadgetParentOrganizationRetType) {
	res, _ = o.GetParentOrganizationOk()
	return
}

// GetParentOrganizationOk returns a tuple with the ParentOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImageSharePayload) GetParentOrganizationOk() (ret UpdateImageSharePayloadgetParentOrganizationRetType, ok bool) {
	return getUpdateImageSharePayloadgetParentOrganizationAttributeTypeOk(o.ParentOrganization)
}

// HasParentOrganization returns a boolean if a field has been set.
func (o *UpdateImageSharePayload) HasParentOrganization() bool {
	_, ok := o.GetParentOrganizationOk()
	return ok
}

// SetParentOrganization gets a reference to the given bool and assigns it to the ParentOrganization field.
func (o *UpdateImageSharePayload) SetParentOrganization(v UpdateImageSharePayloadgetParentOrganizationRetType) {
	setUpdateImageSharePayloadgetParentOrganizationAttributeType(&o.ParentOrganization, v)
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *UpdateImageSharePayload) GetProjects() (res UpdateImageSharePayloadGetProjectsRetType) {
	res, _ = o.GetProjectsOk()
	return
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImageSharePayload) GetProjectsOk() (ret UpdateImageSharePayloadGetProjectsRetType, ok bool) {
	return getUpdateImageSharePayloadGetProjectsAttributeTypeOk(o.Projects)
}

// HasProjects returns a boolean if a field has been set.
func (o *UpdateImageSharePayload) HasProjects() bool {
	_, ok := o.GetProjectsOk()
	return ok
}

// SetProjects gets a reference to the given []string and assigns it to the Projects field.
func (o *UpdateImageSharePayload) SetProjects(v UpdateImageSharePayloadGetProjectsRetType) {
	setUpdateImageSharePayloadGetProjectsAttributeType(&o.Projects, v)
}

func (o UpdateImageSharePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getUpdateImageSharePayloadgetParentOrganizationAttributeTypeOk(o.ParentOrganization); ok {
		toSerialize["ParentOrganization"] = val
	}
	if val, ok := getUpdateImageSharePayloadGetProjectsAttributeTypeOk(o.Projects); ok {
		toSerialize["Projects"] = val
	}
	return toSerialize, nil
}

type NullableUpdateImageSharePayload struct {
	value *UpdateImageSharePayload
	isSet bool
}

func (v NullableUpdateImageSharePayload) Get() *UpdateImageSharePayload {
	return v.value
}

func (v *NullableUpdateImageSharePayload) Set(val *UpdateImageSharePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateImageSharePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateImageSharePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateImageSharePayload(val *UpdateImageSharePayload) *NullableUpdateImageSharePayload {
	return &NullableUpdateImageSharePayload{value: val, isSet: true}
}

func (v NullableUpdateImageSharePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateImageSharePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
