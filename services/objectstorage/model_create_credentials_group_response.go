/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

import (
	"encoding/json"
)

// checks if the CreateCredentialsGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCredentialsGroupResponse{}

// CreateCredentialsGroupResponse struct for CreateCredentialsGroupResponse
type CreateCredentialsGroupResponse struct {
	// REQUIRED
	CredentialsGroup *CredentialsGroup `json:"credentialsGroup"`
	// Project ID
	// REQUIRED
	Project *string `json:"project"`
}

type _CreateCredentialsGroupResponse CreateCredentialsGroupResponse

// NewCreateCredentialsGroupResponse instantiates a new CreateCredentialsGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCredentialsGroupResponse(credentialsGroup *CredentialsGroup, project *string) *CreateCredentialsGroupResponse {
	this := CreateCredentialsGroupResponse{}
	this.CredentialsGroup = credentialsGroup
	this.Project = project
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
func (o *CreateCredentialsGroupResponse) GetCredentialsGroup() *CredentialsGroup {
	if o == nil {
		var ret *CredentialsGroup
		return ret
	}

	return o.CredentialsGroup
}

// GetCredentialsGroupOk returns a tuple with the CredentialsGroup field value
// and a boolean to check if the value has been set.
func (o *CreateCredentialsGroupResponse) GetCredentialsGroupOk() (*CredentialsGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.CredentialsGroup, true
}

// SetCredentialsGroup sets field value
func (o *CreateCredentialsGroupResponse) SetCredentialsGroup(v *CredentialsGroup) {
	o.CredentialsGroup = v
}

// GetProject returns the Project field value
func (o *CreateCredentialsGroupResponse) GetProject() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *CreateCredentialsGroupResponse) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project, true
}

// SetProject sets field value
func (o *CreateCredentialsGroupResponse) SetProject(v *string) {
	o.Project = v
}

func (o CreateCredentialsGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentialsGroup"] = o.CredentialsGroup
	toSerialize["project"] = o.Project
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
