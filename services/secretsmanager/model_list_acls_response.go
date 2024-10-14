/*
STACKIT Secrets Manager API

This API provides endpoints for managing the Secrets-Manager.

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretsmanager

import (
	"encoding/json"
)

// checks if the ListACLsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListACLsResponse{}

// ListACLsResponse struct for ListACLsResponse
type ListACLsResponse struct {
	// REQUIRED
	Acls *[]ACL `json:"acls"`
}

type _ListACLsResponse ListACLsResponse

// NewListACLsResponse instantiates a new ListACLsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListACLsResponse(acls *[]ACL) *ListACLsResponse {
	this := ListACLsResponse{}
	this.Acls = acls
	return &this
}

// NewListACLsResponseWithDefaults instantiates a new ListACLsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListACLsResponseWithDefaults() *ListACLsResponse {
	this := ListACLsResponse{}
	return &this
}

// GetAcls returns the Acls field value
func (o *ListACLsResponse) GetAcls() *[]ACL {
	if o == nil {
		var ret *[]ACL
		return ret
	}

	return o.Acls
}

// GetAclsOk returns a tuple with the Acls field value
// and a boolean to check if the value has been set.
func (o *ListACLsResponse) GetAclsOk() (*[]ACL, bool) {
	if o == nil {
		return nil, false
	}
	return o.Acls, true
}

// SetAcls sets field value
func (o *ListACLsResponse) SetAcls(v *[]ACL) {
	o.Acls = v
}

func (o ListACLsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acls"] = o.Acls
	return toSerialize, nil
}

type NullableListACLsResponse struct {
	value *ListACLsResponse
	isSet bool
}

func (v NullableListACLsResponse) Get() *ListACLsResponse {
	return v.value
}

func (v *NullableListACLsResponse) Set(val *ListACLsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListACLsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListACLsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListACLsResponse(val *ListACLsResponse) *NullableListACLsResponse {
	return &NullableListACLsResponse{value: val, isSet: true}
}

func (v NullableListACLsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListACLsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
