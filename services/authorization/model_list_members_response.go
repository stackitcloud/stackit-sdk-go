/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorization

import (
	"encoding/json"
)

// checks if the ListMembersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMembersResponse{}

// ListMembersResponse struct for ListMembersResponse
type ListMembersResponse struct {
	// REQUIRED
	Members *[]Member `json:"members"`
	// REQUIRED
	ResourceId *string `json:"resourceId"`
	// REQUIRED
	ResourceType *string `json:"resourceType"`
}

type _ListMembersResponse ListMembersResponse

// NewListMembersResponse instantiates a new ListMembersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMembersResponse(members *[]Member, resourceId *string, resourceType *string) *ListMembersResponse {
	this := ListMembersResponse{}
	this.Members = members
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	return &this
}

// NewListMembersResponseWithDefaults instantiates a new ListMembersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMembersResponseWithDefaults() *ListMembersResponse {
	this := ListMembersResponse{}
	return &this
}

// GetMembers returns the Members field value
func (o *ListMembersResponse) GetMembers() *[]Member {
	if o == nil || IsNil(o.Members) {
		var ret *[]Member
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *ListMembersResponse) GetMembersOk() (*[]Member, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *ListMembersResponse) SetMembers(v *[]Member) {
	o.Members = v
}

// GetResourceId returns the ResourceId field value
func (o *ListMembersResponse) GetResourceId() *string {
	if o == nil || IsNil(o.ResourceId) {
		var ret *string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ListMembersResponse) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// SetResourceId sets field value
func (o *ListMembersResponse) SetResourceId(v *string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value
func (o *ListMembersResponse) GetResourceType() *string {
	if o == nil || IsNil(o.ResourceType) {
		var ret *string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ListMembersResponse) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// SetResourceType sets field value
func (o *ListMembersResponse) SetResourceType(v *string) {
	o.ResourceType = v
}

func (o ListMembersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["members"] = o.Members
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["resourceType"] = o.ResourceType
	return toSerialize, nil
}

type NullableListMembersResponse struct {
	value *ListMembersResponse
	isSet bool
}

func (v NullableListMembersResponse) Get() *ListMembersResponse {
	return v.value
}

func (v *NullableListMembersResponse) Set(val *ListMembersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListMembersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListMembersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMembersResponse(val *ListMembersResponse) *NullableListMembersResponse {
	return &NullableListMembersResponse{value: val, isSet: true}
}

func (v NullableListMembersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMembersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
