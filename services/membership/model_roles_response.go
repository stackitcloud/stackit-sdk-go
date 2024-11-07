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

// checks if the RolesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolesResponse{}

// RolesResponse struct for RolesResponse
type RolesResponse struct {
	// REQUIRED
	ResourceId *string `json:"resourceId"`
	// REQUIRED
	ResourceType *string `json:"resourceType"`
	// REQUIRED
	Roles *[]Role `json:"roles"`
}

type _RolesResponse RolesResponse

// NewRolesResponse instantiates a new RolesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolesResponse(resourceId *string, resourceType *string, roles *[]Role) *RolesResponse {
	this := RolesResponse{}
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	this.Roles = roles
	return &this
}

// NewRolesResponseWithDefaults instantiates a new RolesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolesResponseWithDefaults() *RolesResponse {
	this := RolesResponse{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *RolesResponse) GetResourceId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *RolesResponse) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// SetResourceId sets field value
func (o *RolesResponse) SetResourceId(v *string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value
func (o *RolesResponse) GetResourceType() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *RolesResponse) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// SetResourceType sets field value
func (o *RolesResponse) SetResourceType(v *string) {
	o.ResourceType = v
}

// GetRoles returns the Roles field value
func (o *RolesResponse) GetRoles() *[]Role {
	if o == nil {
		var ret *[]Role
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *RolesResponse) GetRolesOk() (*[]Role, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *RolesResponse) SetRoles(v *[]Role) {
	o.Roles = v
}

func (o RolesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

type NullableRolesResponse struct {
	value *RolesResponse
	isSet bool
}

func (v NullableRolesResponse) Get() *RolesResponse {
	return v.value
}

func (v *NullableRolesResponse) Set(val *RolesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRolesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRolesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolesResponse(val *RolesResponse) *NullableRolesResponse {
	return &NullableRolesResponse{value: val, isSet: true}
}

func (v NullableRolesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
