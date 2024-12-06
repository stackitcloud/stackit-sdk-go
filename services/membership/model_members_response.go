/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package membership

import (
	"encoding/json"
)

// checks if the MembersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembersResponse{}

// MembersResponse struct for MembersResponse
type MembersResponse struct {
	// REQUIRED
	Members *[]Member `json:"members"`
	// REQUIRED
	ResourceId *string `json:"resourceId"`
	// REQUIRED
	ResourceType *string `json:"resourceType"`
	WrittenAt    *Zookie `json:"writtenAt,omitempty"`
}

type _MembersResponse MembersResponse

// NewMembersResponse instantiates a new MembersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembersResponse(members *[]Member, resourceId *string, resourceType *string) *MembersResponse {
	this := MembersResponse{}
	this.Members = members
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	return &this
}

// NewMembersResponseWithDefaults instantiates a new MembersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembersResponseWithDefaults() *MembersResponse {
	this := MembersResponse{}
	return &this
}

// GetMembers returns the Members field value
func (o *MembersResponse) GetMembers() *[]Member {
	if o == nil {
		var ret *[]Member
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *MembersResponse) GetMembersOk() (*[]Member, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *MembersResponse) SetMembers(v *[]Member) {
	o.Members = v
}

// GetResourceId returns the ResourceId field value
func (o *MembersResponse) GetResourceId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *MembersResponse) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// SetResourceId sets field value
func (o *MembersResponse) SetResourceId(v *string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value
func (o *MembersResponse) GetResourceType() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *MembersResponse) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// SetResourceType sets field value
func (o *MembersResponse) SetResourceType(v *string) {
	o.ResourceType = v
}

// GetWrittenAt returns the WrittenAt field value if set, zero value otherwise.
func (o *MembersResponse) GetWrittenAt() *Zookie {
	if o == nil || IsNil(o.WrittenAt) {
		var ret *Zookie
		return ret
	}
	return o.WrittenAt
}

// GetWrittenAtOk returns a tuple with the WrittenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembersResponse) GetWrittenAtOk() (*Zookie, bool) {
	if o == nil || IsNil(o.WrittenAt) {
		return nil, false
	}
	return o.WrittenAt, true
}

// HasWrittenAt returns a boolean if a field has been set.
func (o *MembersResponse) HasWrittenAt() bool {
	if o != nil && !IsNil(o.WrittenAt) {
		return true
	}

	return false
}

// SetWrittenAt gets a reference to the given Zookie and assigns it to the WrittenAt field.
func (o *MembersResponse) SetWrittenAt(v *Zookie) {
	o.WrittenAt = v
}

func (o MembersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["members"] = o.Members
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["resourceType"] = o.ResourceType
	if !IsNil(o.WrittenAt) {
		toSerialize["writtenAt"] = o.WrittenAt
	}
	return toSerialize, nil
}

type NullableMembersResponse struct {
	value *MembersResponse
	isSet bool
}

func (v NullableMembersResponse) Get() *MembersResponse {
	return v.value
}

func (v *NullableMembersResponse) Set(val *MembersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMembersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMembersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembersResponse(val *MembersResponse) *NullableMembersResponse {
	return &NullableMembersResponse{value: val, isSet: true}
}

func (v NullableMembersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
