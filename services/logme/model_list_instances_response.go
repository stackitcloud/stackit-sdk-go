/*
STACKIT LogMe API

The STACKIT LogMe API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logme

import (
	"encoding/json"
)

// checks if the ListInstancesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInstancesResponse{}

/*
	types and functions for instances
*/

// isArray
type ListInstancesResponseGetInstancesAttributeType = *[]Instance
type ListInstancesResponseGetInstancesArgType = []Instance
type ListInstancesResponseGetInstancesRetType = []Instance

func getListInstancesResponseGetInstancesAttributeTypeOk(arg ListInstancesResponseGetInstancesAttributeType) (ret ListInstancesResponseGetInstancesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListInstancesResponseGetInstancesAttributeType(arg *ListInstancesResponseGetInstancesAttributeType, val ListInstancesResponseGetInstancesRetType) {
	*arg = &val
}

// ListInstancesResponse struct for ListInstancesResponse
type ListInstancesResponse struct {
	// REQUIRED
	Instances ListInstancesResponseGetInstancesAttributeType `json:"instances"`
}

type _ListInstancesResponse ListInstancesResponse

// NewListInstancesResponse instantiates a new ListInstancesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInstancesResponse(instances ListInstancesResponseGetInstancesArgType) *ListInstancesResponse {
	this := ListInstancesResponse{}
	setListInstancesResponseGetInstancesAttributeType(&this.Instances, instances)
	return &this
}

// NewListInstancesResponseWithDefaults instantiates a new ListInstancesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInstancesResponseWithDefaults() *ListInstancesResponse {
	this := ListInstancesResponse{}
	return &this
}

// GetInstances returns the Instances field value
func (o *ListInstancesResponse) GetInstances() (ret ListInstancesResponseGetInstancesRetType) {
	ret, _ = o.GetInstancesOk()
	return ret
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponse) GetInstancesOk() (ret ListInstancesResponseGetInstancesRetType, ok bool) {
	return getListInstancesResponseGetInstancesAttributeTypeOk(o.Instances)
}

// SetInstances sets field value
func (o *ListInstancesResponse) SetInstances(v ListInstancesResponseGetInstancesRetType) {
	setListInstancesResponseGetInstancesAttributeType(&o.Instances, v)
}

func (o ListInstancesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListInstancesResponseGetInstancesAttributeTypeOk(o.Instances); ok {
		toSerialize["Instances"] = val
	}
	return toSerialize, nil
}

type NullableListInstancesResponse struct {
	value *ListInstancesResponse
	isSet bool
}

func (v NullableListInstancesResponse) Get() *ListInstancesResponse {
	return v.value
}

func (v *NullableListInstancesResponse) Set(val *ListInstancesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListInstancesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListInstancesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInstancesResponse(val *ListInstancesResponse) *NullableListInstancesResponse {
	return &NullableListInstancesResponse{value: val, isSet: true}
}

func (v NullableListInstancesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInstancesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
