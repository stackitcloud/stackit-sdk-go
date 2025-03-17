/*
STACKIT RabbitMQ API

The STACKIT RabbitMQ API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rabbitmq

import (
	"encoding/json"
)

// checks if the ListCredentialsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCredentialsResponse{}

/*
	types and functions for credentialsList
*/

// isArray
type ListCredentialsResponseGetCredentialsListAttributeType = *[]CredentialsListItem
type ListCredentialsResponseGetCredentialsListArgType = []CredentialsListItem
type ListCredentialsResponseGetCredentialsListRetType = []CredentialsListItem

func getListCredentialsResponseGetCredentialsListAttributeTypeOk(arg ListCredentialsResponseGetCredentialsListAttributeType) (ret ListCredentialsResponseGetCredentialsListRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListCredentialsResponseGetCredentialsListAttributeType(arg *ListCredentialsResponseGetCredentialsListAttributeType, val ListCredentialsResponseGetCredentialsListRetType) {
	*arg = &val
}

// ListCredentialsResponse struct for ListCredentialsResponse
type ListCredentialsResponse struct {
	// REQUIRED
	CredentialsList ListCredentialsResponseGetCredentialsListAttributeType `json:"credentialsList"`
}

type _ListCredentialsResponse ListCredentialsResponse

// NewListCredentialsResponse instantiates a new ListCredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCredentialsResponse(credentialsList ListCredentialsResponseGetCredentialsListArgType) *ListCredentialsResponse {
	this := ListCredentialsResponse{}
	setListCredentialsResponseGetCredentialsListAttributeType(&this.CredentialsList, credentialsList)
	return &this
}

// NewListCredentialsResponseWithDefaults instantiates a new ListCredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCredentialsResponseWithDefaults() *ListCredentialsResponse {
	this := ListCredentialsResponse{}
	return &this
}

// GetCredentialsList returns the CredentialsList field value
func (o *ListCredentialsResponse) GetCredentialsList() (ret ListCredentialsResponseGetCredentialsListRetType) {
	ret, _ = o.GetCredentialsListOk()
	return ret
}

// GetCredentialsListOk returns a tuple with the CredentialsList field value
// and a boolean to check if the value has been set.
func (o *ListCredentialsResponse) GetCredentialsListOk() (ret ListCredentialsResponseGetCredentialsListRetType, ok bool) {
	return getListCredentialsResponseGetCredentialsListAttributeTypeOk(o.CredentialsList)
}

// SetCredentialsList sets field value
func (o *ListCredentialsResponse) SetCredentialsList(v ListCredentialsResponseGetCredentialsListRetType) {
	setListCredentialsResponseGetCredentialsListAttributeType(&o.CredentialsList, v)
}

func (o ListCredentialsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListCredentialsResponseGetCredentialsListAttributeTypeOk(o.CredentialsList); ok {
		toSerialize["CredentialsList"] = val
	}
	return toSerialize, nil
}

type NullableListCredentialsResponse struct {
	value *ListCredentialsResponse
	isSet bool
}

func (v NullableListCredentialsResponse) Get() *ListCredentialsResponse {
	return v.value
}

func (v *NullableListCredentialsResponse) Set(val *ListCredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCredentialsResponse(val *ListCredentialsResponse) *NullableListCredentialsResponse {
	return &NullableListCredentialsResponse{value: val, isSet: true}
}

func (v NullableListCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
