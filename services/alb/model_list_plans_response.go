/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 2beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alb

import (
	"encoding/json"
)

// checks if the ListPlansResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPlansResponse{}

/*
	types and functions for validPlans
*/

// isArray
type ListPlansResponseGetValidPlansAttributeType = *[]PlanDetails
type ListPlansResponseGetValidPlansArgType = []PlanDetails
type ListPlansResponseGetValidPlansRetType = []PlanDetails

func getListPlansResponseGetValidPlansAttributeTypeOk(arg ListPlansResponseGetValidPlansAttributeType) (ret ListPlansResponseGetValidPlansRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListPlansResponseGetValidPlansAttributeType(arg *ListPlansResponseGetValidPlansAttributeType, val ListPlansResponseGetValidPlansRetType) {
	*arg = &val
}

// ListPlansResponse struct for ListPlansResponse
type ListPlansResponse struct {
	ValidPlans ListPlansResponseGetValidPlansAttributeType `json:"validPlans,omitempty"`
}

// NewListPlansResponse instantiates a new ListPlansResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPlansResponse() *ListPlansResponse {
	this := ListPlansResponse{}
	return &this
}

// NewListPlansResponseWithDefaults instantiates a new ListPlansResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPlansResponseWithDefaults() *ListPlansResponse {
	this := ListPlansResponse{}
	return &this
}

// GetValidPlans returns the ValidPlans field value if set, zero value otherwise.
func (o *ListPlansResponse) GetValidPlans() (res ListPlansResponseGetValidPlansRetType) {
	res, _ = o.GetValidPlansOk()
	return
}

// GetValidPlansOk returns a tuple with the ValidPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlansResponse) GetValidPlansOk() (ret ListPlansResponseGetValidPlansRetType, ok bool) {
	return getListPlansResponseGetValidPlansAttributeTypeOk(o.ValidPlans)
}

// HasValidPlans returns a boolean if a field has been set.
func (o *ListPlansResponse) HasValidPlans() bool {
	_, ok := o.GetValidPlansOk()
	return ok
}

// SetValidPlans gets a reference to the given []PlanDetails and assigns it to the ValidPlans field.
func (o *ListPlansResponse) SetValidPlans(v ListPlansResponseGetValidPlansRetType) {
	setListPlansResponseGetValidPlansAttributeType(&o.ValidPlans, v)
}

func (o ListPlansResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListPlansResponseGetValidPlansAttributeTypeOk(o.ValidPlans); ok {
		toSerialize["ValidPlans"] = val
	}
	return toSerialize, nil
}

type NullableListPlansResponse struct {
	value *ListPlansResponse
	isSet bool
}

func (v NullableListPlansResponse) Get() *ListPlansResponse {
	return v.value
}

func (v *NullableListPlansResponse) Set(val *ListPlansResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPlansResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPlansResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPlansResponse(val *ListPlansResponse) *NullableListPlansResponse {
	return &NullableListPlansResponse{value: val, isSet: true}
}

func (v NullableListPlansResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPlansResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
