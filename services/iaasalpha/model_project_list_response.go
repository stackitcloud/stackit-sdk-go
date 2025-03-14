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

// checks if the ProjectListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectListResponse{}

/*
	types and functions for items
*/

// isArray
type ProjectListResponseGetItemsAttributeType = *[]string
type ProjectListResponseGetItemsArgType = []string
type ProjectListResponseGetItemsRetType = []string

func getProjectListResponseGetItemsAttributeTypeOk(arg ProjectListResponseGetItemsAttributeType) (ret ProjectListResponseGetItemsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setProjectListResponseGetItemsAttributeType(arg *ProjectListResponseGetItemsAttributeType, val ProjectListResponseGetItemsRetType) {
	*arg = &val
}

// ProjectListResponse Project list response.
type ProjectListResponse struct {
	// A list of STACKIT projects.
	// REQUIRED
	Items ProjectListResponseGetItemsAttributeType `json:"items"`
}

type _ProjectListResponse ProjectListResponse

// NewProjectListResponse instantiates a new ProjectListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectListResponse(items ProjectListResponseGetItemsArgType) *ProjectListResponse {
	this := ProjectListResponse{}
	setProjectListResponseGetItemsAttributeType(&this.Items, items)
	return &this
}

// NewProjectListResponseWithDefaults instantiates a new ProjectListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectListResponseWithDefaults() *ProjectListResponse {
	this := ProjectListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *ProjectListResponse) GetItems() (ret ProjectListResponseGetItemsRetType) {
	ret, _ = o.GetItemsOk()
	return ret
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ProjectListResponse) GetItemsOk() (ret ProjectListResponseGetItemsRetType, ok bool) {
	return getProjectListResponseGetItemsAttributeTypeOk(o.Items)
}

// SetItems sets field value
func (o *ProjectListResponse) SetItems(v ProjectListResponseGetItemsRetType) {
	setProjectListResponseGetItemsAttributeType(&o.Items, v)
}

func (o ProjectListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getProjectListResponseGetItemsAttributeTypeOk(o.Items); ok {
		toSerialize["Items"] = val
	}
	return toSerialize, nil
}

type NullableProjectListResponse struct {
	value *ProjectListResponse
	isSet bool
}

func (v NullableProjectListResponse) Get() *ProjectListResponse {
	return v.value
}

func (v *NullableProjectListResponse) Set(val *ProjectListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectListResponse(val *ProjectListResponse) *NullableProjectListResponse {
	return &NullableProjectListResponse{value: val, isSet: true}
}

func (v NullableProjectListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
