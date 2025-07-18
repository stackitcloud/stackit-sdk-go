/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 2.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

import (
	"encoding/json"
)

// checks if the ListAccessKeysResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAccessKeysResponse{}

/*
	types and functions for accessKeys
*/

// isArray
type ListAccessKeysResponseGetAccessKeysAttributeType = *[]AccessKey
type ListAccessKeysResponseGetAccessKeysArgType = []AccessKey
type ListAccessKeysResponseGetAccessKeysRetType = []AccessKey

func getListAccessKeysResponseGetAccessKeysAttributeTypeOk(arg ListAccessKeysResponseGetAccessKeysAttributeType) (ret ListAccessKeysResponseGetAccessKeysRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListAccessKeysResponseGetAccessKeysAttributeType(arg *ListAccessKeysResponseGetAccessKeysAttributeType, val ListAccessKeysResponseGetAccessKeysRetType) {
	*arg = &val
}

/*
	types and functions for project
*/

// isNotNullableString
type ListAccessKeysResponseGetProjectAttributeType = *string

func getListAccessKeysResponseGetProjectAttributeTypeOk(arg ListAccessKeysResponseGetProjectAttributeType) (ret ListAccessKeysResponseGetProjectRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListAccessKeysResponseGetProjectAttributeType(arg *ListAccessKeysResponseGetProjectAttributeType, val ListAccessKeysResponseGetProjectRetType) {
	*arg = &val
}

type ListAccessKeysResponseGetProjectArgType = string
type ListAccessKeysResponseGetProjectRetType = string

// ListAccessKeysResponse struct for ListAccessKeysResponse
type ListAccessKeysResponse struct {
	// REQUIRED
	AccessKeys ListAccessKeysResponseGetAccessKeysAttributeType `json:"accessKeys" required:"true"`
	// Project ID
	// REQUIRED
	Project ListAccessKeysResponseGetProjectAttributeType `json:"project" required:"true"`
}

type _ListAccessKeysResponse ListAccessKeysResponse

// NewListAccessKeysResponse instantiates a new ListAccessKeysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAccessKeysResponse(accessKeys ListAccessKeysResponseGetAccessKeysArgType, project ListAccessKeysResponseGetProjectArgType) *ListAccessKeysResponse {
	this := ListAccessKeysResponse{}
	setListAccessKeysResponseGetAccessKeysAttributeType(&this.AccessKeys, accessKeys)
	setListAccessKeysResponseGetProjectAttributeType(&this.Project, project)
	return &this
}

// NewListAccessKeysResponseWithDefaults instantiates a new ListAccessKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAccessKeysResponseWithDefaults() *ListAccessKeysResponse {
	this := ListAccessKeysResponse{}
	return &this
}

// GetAccessKeys returns the AccessKeys field value
func (o *ListAccessKeysResponse) GetAccessKeys() (ret ListAccessKeysResponseGetAccessKeysRetType) {
	ret, _ = o.GetAccessKeysOk()
	return ret
}

// GetAccessKeysOk returns a tuple with the AccessKeys field value
// and a boolean to check if the value has been set.
func (o *ListAccessKeysResponse) GetAccessKeysOk() (ret ListAccessKeysResponseGetAccessKeysRetType, ok bool) {
	return getListAccessKeysResponseGetAccessKeysAttributeTypeOk(o.AccessKeys)
}

// SetAccessKeys sets field value
func (o *ListAccessKeysResponse) SetAccessKeys(v ListAccessKeysResponseGetAccessKeysRetType) {
	setListAccessKeysResponseGetAccessKeysAttributeType(&o.AccessKeys, v)
}

// GetProject returns the Project field value
func (o *ListAccessKeysResponse) GetProject() (ret ListAccessKeysResponseGetProjectRetType) {
	ret, _ = o.GetProjectOk()
	return ret
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ListAccessKeysResponse) GetProjectOk() (ret ListAccessKeysResponseGetProjectRetType, ok bool) {
	return getListAccessKeysResponseGetProjectAttributeTypeOk(o.Project)
}

// SetProject sets field value
func (o *ListAccessKeysResponse) SetProject(v ListAccessKeysResponseGetProjectRetType) {
	setListAccessKeysResponseGetProjectAttributeType(&o.Project, v)
}

func (o ListAccessKeysResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListAccessKeysResponseGetAccessKeysAttributeTypeOk(o.AccessKeys); ok {
		toSerialize["AccessKeys"] = val
	}
	if val, ok := getListAccessKeysResponseGetProjectAttributeTypeOk(o.Project); ok {
		toSerialize["Project"] = val
	}
	return toSerialize, nil
}

type NullableListAccessKeysResponse struct {
	value *ListAccessKeysResponse
	isSet bool
}

func (v NullableListAccessKeysResponse) Get() *ListAccessKeysResponse {
	return v.value
}

func (v *NullableListAccessKeysResponse) Set(val *ListAccessKeysResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccessKeysResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccessKeysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccessKeysResponse(val *ListAccessKeysResponse) *NullableListAccessKeysResponse {
	return &NullableListAccessKeysResponse{value: val, isSet: true}
}

func (v NullableListAccessKeysResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccessKeysResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
