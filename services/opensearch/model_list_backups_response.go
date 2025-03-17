/*
STACKIT Opensearch API

The STACKIT Opensearch API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opensearch

import (
	"encoding/json"
)

// checks if the ListBackupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListBackupsResponse{}

/*
	types and functions for instanceBackups
*/

// isArray
type ListBackupsResponseGetInstanceBackupsAttributeType = *[]Backup
type ListBackupsResponseGetInstanceBackupsArgType = []Backup
type ListBackupsResponseGetInstanceBackupsRetType = []Backup

func getListBackupsResponseGetInstanceBackupsAttributeTypeOk(arg ListBackupsResponseGetInstanceBackupsAttributeType) (ret ListBackupsResponseGetInstanceBackupsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListBackupsResponseGetInstanceBackupsAttributeType(arg *ListBackupsResponseGetInstanceBackupsAttributeType, val ListBackupsResponseGetInstanceBackupsRetType) {
	*arg = &val
}

// ListBackupsResponse struct for ListBackupsResponse
type ListBackupsResponse struct {
	// REQUIRED
	InstanceBackups ListBackupsResponseGetInstanceBackupsAttributeType `json:"instanceBackups"`
}

type _ListBackupsResponse ListBackupsResponse

// NewListBackupsResponse instantiates a new ListBackupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBackupsResponse(instanceBackups ListBackupsResponseGetInstanceBackupsArgType) *ListBackupsResponse {
	this := ListBackupsResponse{}
	setListBackupsResponseGetInstanceBackupsAttributeType(&this.InstanceBackups, instanceBackups)
	return &this
}

// NewListBackupsResponseWithDefaults instantiates a new ListBackupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBackupsResponseWithDefaults() *ListBackupsResponse {
	this := ListBackupsResponse{}
	return &this
}

// GetInstanceBackups returns the InstanceBackups field value
func (o *ListBackupsResponse) GetInstanceBackups() (ret ListBackupsResponseGetInstanceBackupsRetType) {
	ret, _ = o.GetInstanceBackupsOk()
	return ret
}

// GetInstanceBackupsOk returns a tuple with the InstanceBackups field value
// and a boolean to check if the value has been set.
func (o *ListBackupsResponse) GetInstanceBackupsOk() (ret ListBackupsResponseGetInstanceBackupsRetType, ok bool) {
	return getListBackupsResponseGetInstanceBackupsAttributeTypeOk(o.InstanceBackups)
}

// SetInstanceBackups sets field value
func (o *ListBackupsResponse) SetInstanceBackups(v ListBackupsResponseGetInstanceBackupsRetType) {
	setListBackupsResponseGetInstanceBackupsAttributeType(&o.InstanceBackups, v)
}

func (o ListBackupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListBackupsResponseGetInstanceBackupsAttributeTypeOk(o.InstanceBackups); ok {
		toSerialize["InstanceBackups"] = val
	}
	return toSerialize, nil
}

type NullableListBackupsResponse struct {
	value *ListBackupsResponse
	isSet bool
}

func (v NullableListBackupsResponse) Get() *ListBackupsResponse {
	return v.value
}

func (v *NullableListBackupsResponse) Set(val *ListBackupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListBackupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListBackupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBackupsResponse(val *ListBackupsResponse) *NullableListBackupsResponse {
	return &NullableListBackupsResponse{value: val, isSet: true}
}

func (v NullableListBackupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBackupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
