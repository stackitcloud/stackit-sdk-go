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

// checks if the GetQuotaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetQuotaResponse{}

// GetQuotaResponse struct for GetQuotaResponse
type GetQuotaResponse struct {
	// The maximum number of load balancing servers in this project. Unlimited if set to -1.
	// Can be cast to int32 without loss of precision.
	MaxLoadBalancers *int64 `json:"maxLoadBalancers,omitempty"`
	// Project identifier
	ProjectId *string `json:"projectId,omitempty"`
}

// NewGetQuotaResponse instantiates a new GetQuotaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetQuotaResponse() *GetQuotaResponse {
	this := GetQuotaResponse{}
	return &this
}

// NewGetQuotaResponseWithDefaults instantiates a new GetQuotaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetQuotaResponseWithDefaults() *GetQuotaResponse {
	this := GetQuotaResponse{}
	return &this
}

// GetMaxLoadBalancers returns the MaxLoadBalancers field value if set, zero value otherwise.
func (o *GetQuotaResponse) GetMaxLoadBalancers() *int64 {
	if o == nil || IsNil(o.MaxLoadBalancers) {
		var ret *int64
		return ret
	}
	return o.MaxLoadBalancers
}

// GetMaxLoadBalancersOk returns a tuple with the MaxLoadBalancers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuotaResponse) GetMaxLoadBalancersOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxLoadBalancers) {
		return nil, false
	}
	return o.MaxLoadBalancers, true
}

// HasMaxLoadBalancers returns a boolean if a field has been set.
func (o *GetQuotaResponse) HasMaxLoadBalancers() bool {
	if o != nil && !IsNil(o.MaxLoadBalancers) && !IsNil(o.MaxLoadBalancers) {
		return true
	}

	return false
}

// SetMaxLoadBalancers gets a reference to the given int64 and assigns it to the MaxLoadBalancers field.
func (o *GetQuotaResponse) SetMaxLoadBalancers(v *int64) {
	o.MaxLoadBalancers = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GetQuotaResponse) GetProjectId() *string {
	if o == nil || IsNil(o.ProjectId) {
		var ret *string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuotaResponse) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GetQuotaResponse) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GetQuotaResponse) SetProjectId(v *string) {
	o.ProjectId = v
}

func (o GetQuotaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxLoadBalancers) {
		toSerialize["maxLoadBalancers"] = o.MaxLoadBalancers
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableGetQuotaResponse struct {
	value *GetQuotaResponse
	isSet bool
}

func (v NullableGetQuotaResponse) Get() *GetQuotaResponse {
	return v.value
}

func (v *NullableGetQuotaResponse) Set(val *GetQuotaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQuotaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQuotaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQuotaResponse(val *GetQuotaResponse) *NullableGetQuotaResponse {
	return &NullableGetQuotaResponse{value: val, isSet: true}
}

func (v NullableGetQuotaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQuotaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
