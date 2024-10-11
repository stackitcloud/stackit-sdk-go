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

// checks if the ListRestoresResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRestoresResponse{}

// ListRestoresResponse struct for ListRestoresResponse
type ListRestoresResponse struct {
	// REQUIRED
	InstanceRestores *[]Restore `json:"instanceRestores"`
}

type _ListRestoresResponse ListRestoresResponse

// NewListRestoresResponse instantiates a new ListRestoresResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRestoresResponse(instanceRestores *[]Restore) *ListRestoresResponse {
	this := ListRestoresResponse{}
	this.InstanceRestores = instanceRestores
	return &this
}

// NewListRestoresResponseWithDefaults instantiates a new ListRestoresResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRestoresResponseWithDefaults() *ListRestoresResponse {
	this := ListRestoresResponse{}
	return &this
}

// GetInstanceRestores returns the InstanceRestores field value
func (o *ListRestoresResponse) GetInstanceRestores() *[]Restore {
	if o == nil {
		var ret *[]Restore
		return ret
	}

	return o.InstanceRestores
}

// GetInstanceRestoresOk returns a tuple with the InstanceRestores field value
// and a boolean to check if the value has been set.
func (o *ListRestoresResponse) GetInstanceRestoresOk() (*[]Restore, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceRestores, true
}

// SetInstanceRestores sets field value
func (o *ListRestoresResponse) SetInstanceRestores(v *[]Restore) {
	o.InstanceRestores = v
}

func (o ListRestoresResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceRestores"] = o.InstanceRestores
	return toSerialize, nil
}

type NullableListRestoresResponse struct {
	value *ListRestoresResponse
	isSet bool
}

func (v NullableListRestoresResponse) Get() *ListRestoresResponse {
	return v.value
}

func (v *NullableListRestoresResponse) Set(val *ListRestoresResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRestoresResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRestoresResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRestoresResponse(val *ListRestoresResponse) *NullableListRestoresResponse {
	return &NullableListRestoresResponse{value: val, isSet: true}
}

func (v NullableListRestoresResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRestoresResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
