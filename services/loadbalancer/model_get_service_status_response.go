/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.7.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

import (
	"encoding/json"
)

// checks if the GetServiceStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServiceStatusResponse{}

/*
	types and functions for status
*/

// isEnumRef
type GetServiceStatusResponseGetStatusAttributeType = *string
type GetServiceStatusResponseGetStatusArgType = string
type GetServiceStatusResponseGetStatusRetType = string

func getGetServiceStatusResponseGetStatusAttributeTypeOk(arg GetServiceStatusResponseGetStatusAttributeType) (ret GetServiceStatusResponseGetStatusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetServiceStatusResponseGetStatusAttributeType(arg *GetServiceStatusResponseGetStatusAttributeType, val GetServiceStatusResponseGetStatusRetType) {
	*arg = &val
}

// GetServiceStatusResponse Response with customer project status.
type GetServiceStatusResponse struct {
	// status of the project
	Status GetServiceStatusResponseGetStatusAttributeType `json:"status,omitempty"`
}

// NewGetServiceStatusResponse instantiates a new GetServiceStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServiceStatusResponse() *GetServiceStatusResponse {
	this := GetServiceStatusResponse{}
	return &this
}

// NewGetServiceStatusResponseWithDefaults instantiates a new GetServiceStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServiceStatusResponseWithDefaults() *GetServiceStatusResponse {
	this := GetServiceStatusResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetServiceStatusResponse) GetStatus() (res GetServiceStatusResponseGetStatusRetType) {
	res, _ = o.GetStatusOk()
	return
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceStatusResponse) GetStatusOk() (ret GetServiceStatusResponseGetStatusRetType, ok bool) {
	return getGetServiceStatusResponseGetStatusAttributeTypeOk(o.Status)
}

// HasStatus returns a boolean if a field has been set.
func (o *GetServiceStatusResponse) HasStatus() bool {
	_, ok := o.GetStatusOk()
	return ok
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetServiceStatusResponse) SetStatus(v GetServiceStatusResponseGetStatusRetType) {
	setGetServiceStatusResponseGetStatusAttributeType(&o.Status, v)
}

func (o GetServiceStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getGetServiceStatusResponseGetStatusAttributeTypeOk(o.Status); ok {
		toSerialize["Status"] = val
	}
	return toSerialize, nil
}

type NullableGetServiceStatusResponse struct {
	value *GetServiceStatusResponse
	isSet bool
}

func (v NullableGetServiceStatusResponse) Get() *GetServiceStatusResponse {
	return v.value
}

func (v *NullableGetServiceStatusResponse) Set(val *GetServiceStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServiceStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServiceStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServiceStatusResponse(val *GetServiceStatusResponse) *NullableGetServiceStatusResponse {
	return &NullableGetServiceStatusResponse{value: val, isSet: true}
}

func (v NullableGetServiceStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServiceStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
