/*
STACKIT Git API

Manage STACKIT Git instances.

API version: 1beta.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package git

import (
	"encoding/json"
)

// checks if the ListInstances type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInstances{}

/*
	types and functions for instances
*/

// isArray
type ListInstancesGetInstancesAttributeType = *[]Instance
type ListInstancesGetInstancesArgType = []Instance
type ListInstancesGetInstancesRetType = []Instance

func getListInstancesGetInstancesAttributeTypeOk(arg ListInstancesGetInstancesAttributeType) (ret ListInstancesGetInstancesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListInstancesGetInstancesAttributeType(arg *ListInstancesGetInstancesAttributeType, val ListInstancesGetInstancesRetType) {
	*arg = &val
}

// ListInstances A list of STACKIT Git instances.
type ListInstances struct {
	// REQUIRED
	Instances ListInstancesGetInstancesAttributeType `json:"instances"`
}

type _ListInstances ListInstances

// NewListInstances instantiates a new ListInstances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInstances(instances ListInstancesGetInstancesArgType) *ListInstances {
	this := ListInstances{}
	setListInstancesGetInstancesAttributeType(&this.Instances, instances)
	return &this
}

// NewListInstancesWithDefaults instantiates a new ListInstances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInstancesWithDefaults() *ListInstances {
	this := ListInstances{}
	return &this
}

// GetInstances returns the Instances field value
func (o *ListInstances) GetInstances() (ret ListInstancesGetInstancesRetType) {
	ret, _ = o.GetInstancesOk()
	return ret
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *ListInstances) GetInstancesOk() (ret ListInstancesGetInstancesRetType, ok bool) {
	return getListInstancesGetInstancesAttributeTypeOk(o.Instances)
}

// SetInstances sets field value
func (o *ListInstances) SetInstances(v ListInstancesGetInstancesRetType) {
	setListInstancesGetInstancesAttributeType(&o.Instances, v)
}

func (o ListInstances) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListInstancesGetInstancesAttributeTypeOk(o.Instances); ok {
		toSerialize["Instances"] = val
	}
	return toSerialize, nil
}

type NullableListInstances struct {
	value *ListInstances
	isSet bool
}

func (v NullableListInstances) Get() *ListInstances {
	return v.value
}

func (v *NullableListInstances) Set(val *ListInstances) {
	v.value = val
	v.isSet = true
}

func (v NullableListInstances) IsSet() bool {
	return v.isSet
}

func (v *NullableListInstances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInstances(val *ListInstances) *NullableListInstances {
	return &NullableListInstances{value: val, isSet: true}
}

func (v NullableListInstances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInstances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
