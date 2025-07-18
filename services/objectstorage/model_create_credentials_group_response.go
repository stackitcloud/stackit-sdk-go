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

// checks if the CreateCredentialsGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCredentialsGroupResponse{}

/*
	types and functions for credentialsGroup
*/

// isModel
type CreateCredentialsGroupResponseGetCredentialsGroupAttributeType = *CredentialsGroup
type CreateCredentialsGroupResponseGetCredentialsGroupArgType = CredentialsGroup
type CreateCredentialsGroupResponseGetCredentialsGroupRetType = CredentialsGroup

func getCreateCredentialsGroupResponseGetCredentialsGroupAttributeTypeOk(arg CreateCredentialsGroupResponseGetCredentialsGroupAttributeType) (ret CreateCredentialsGroupResponseGetCredentialsGroupRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateCredentialsGroupResponseGetCredentialsGroupAttributeType(arg *CreateCredentialsGroupResponseGetCredentialsGroupAttributeType, val CreateCredentialsGroupResponseGetCredentialsGroupRetType) {
	*arg = &val
}

/*
	types and functions for project
*/

// isNotNullableString
type CreateCredentialsGroupResponseGetProjectAttributeType = *string

func getCreateCredentialsGroupResponseGetProjectAttributeTypeOk(arg CreateCredentialsGroupResponseGetProjectAttributeType) (ret CreateCredentialsGroupResponseGetProjectRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateCredentialsGroupResponseGetProjectAttributeType(arg *CreateCredentialsGroupResponseGetProjectAttributeType, val CreateCredentialsGroupResponseGetProjectRetType) {
	*arg = &val
}

type CreateCredentialsGroupResponseGetProjectArgType = string
type CreateCredentialsGroupResponseGetProjectRetType = string

// CreateCredentialsGroupResponse struct for CreateCredentialsGroupResponse
type CreateCredentialsGroupResponse struct {
	// REQUIRED
	CredentialsGroup CreateCredentialsGroupResponseGetCredentialsGroupAttributeType `json:"credentialsGroup" required:"true"`
	// Project ID
	// REQUIRED
	Project CreateCredentialsGroupResponseGetProjectAttributeType `json:"project" required:"true"`
}

type _CreateCredentialsGroupResponse CreateCredentialsGroupResponse

// NewCreateCredentialsGroupResponse instantiates a new CreateCredentialsGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCredentialsGroupResponse(credentialsGroup CreateCredentialsGroupResponseGetCredentialsGroupArgType, project CreateCredentialsGroupResponseGetProjectArgType) *CreateCredentialsGroupResponse {
	this := CreateCredentialsGroupResponse{}
	setCreateCredentialsGroupResponseGetCredentialsGroupAttributeType(&this.CredentialsGroup, credentialsGroup)
	setCreateCredentialsGroupResponseGetProjectAttributeType(&this.Project, project)
	return &this
}

// NewCreateCredentialsGroupResponseWithDefaults instantiates a new CreateCredentialsGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCredentialsGroupResponseWithDefaults() *CreateCredentialsGroupResponse {
	this := CreateCredentialsGroupResponse{}
	return &this
}

// GetCredentialsGroup returns the CredentialsGroup field value
func (o *CreateCredentialsGroupResponse) GetCredentialsGroup() (ret CreateCredentialsGroupResponseGetCredentialsGroupRetType) {
	ret, _ = o.GetCredentialsGroupOk()
	return ret
}

// GetCredentialsGroupOk returns a tuple with the CredentialsGroup field value
// and a boolean to check if the value has been set.
func (o *CreateCredentialsGroupResponse) GetCredentialsGroupOk() (ret CreateCredentialsGroupResponseGetCredentialsGroupRetType, ok bool) {
	return getCreateCredentialsGroupResponseGetCredentialsGroupAttributeTypeOk(o.CredentialsGroup)
}

// SetCredentialsGroup sets field value
func (o *CreateCredentialsGroupResponse) SetCredentialsGroup(v CreateCredentialsGroupResponseGetCredentialsGroupRetType) {
	setCreateCredentialsGroupResponseGetCredentialsGroupAttributeType(&o.CredentialsGroup, v)
}

// GetProject returns the Project field value
func (o *CreateCredentialsGroupResponse) GetProject() (ret CreateCredentialsGroupResponseGetProjectRetType) {
	ret, _ = o.GetProjectOk()
	return ret
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *CreateCredentialsGroupResponse) GetProjectOk() (ret CreateCredentialsGroupResponseGetProjectRetType, ok bool) {
	return getCreateCredentialsGroupResponseGetProjectAttributeTypeOk(o.Project)
}

// SetProject sets field value
func (o *CreateCredentialsGroupResponse) SetProject(v CreateCredentialsGroupResponseGetProjectRetType) {
	setCreateCredentialsGroupResponseGetProjectAttributeType(&o.Project, v)
}

func (o CreateCredentialsGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateCredentialsGroupResponseGetCredentialsGroupAttributeTypeOk(o.CredentialsGroup); ok {
		toSerialize["CredentialsGroup"] = val
	}
	if val, ok := getCreateCredentialsGroupResponseGetProjectAttributeTypeOk(o.Project); ok {
		toSerialize["Project"] = val
	}
	return toSerialize, nil
}

type NullableCreateCredentialsGroupResponse struct {
	value *CreateCredentialsGroupResponse
	isSet bool
}

func (v NullableCreateCredentialsGroupResponse) Get() *CreateCredentialsGroupResponse {
	return v.value
}

func (v *NullableCreateCredentialsGroupResponse) Set(val *CreateCredentialsGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCredentialsGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCredentialsGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCredentialsGroupResponse(val *CreateCredentialsGroupResponse) *NullableCreateCredentialsGroupResponse {
	return &NullableCreateCredentialsGroupResponse{value: val, isSet: true}
}

func (v NullableCreateCredentialsGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCredentialsGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
