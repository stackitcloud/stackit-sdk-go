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

// checks if the RemoveRolesPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveRolesPayload{}

// RemoveRolesPayload struct for RemoveRolesPayload
type RemoveRolesPayload struct {
	// REQUIRED
	ResourceType *string `json:"resourceType"`
	// REQUIRED
	Roles *[]RemoveRoleRequest `json:"roles"`
}

type _RemoveRolesPayload RemoveRolesPayload

// NewRemoveRolesPayload instantiates a new RemoveRolesPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveRolesPayload(resourceType *string, roles *[]RemoveRoleRequest) *RemoveRolesPayload {
	this := RemoveRolesPayload{}
	this.ResourceType = resourceType
	this.Roles = roles
	return &this
}

// NewRemoveRolesPayloadWithDefaults instantiates a new RemoveRolesPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveRolesPayloadWithDefaults() *RemoveRolesPayload {
	this := RemoveRolesPayload{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *RemoveRolesPayload) GetResourceType() *string {
	if o == nil || IsNil(o.ResourceType) {
		var ret *string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *RemoveRolesPayload) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// SetResourceType sets field value
func (o *RemoveRolesPayload) SetResourceType(v *string) {
	o.ResourceType = v
}

// GetRoles returns the Roles field value
func (o *RemoveRolesPayload) GetRoles() *[]RemoveRoleRequest {
	if o == nil || IsNil(o.Roles) {
		var ret *[]RemoveRoleRequest
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *RemoveRolesPayload) GetRolesOk() (*[]RemoveRoleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *RemoveRolesPayload) SetRoles(v *[]RemoveRoleRequest) {
	o.Roles = v
}

func (o RemoveRolesPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

type NullableRemoveRolesPayload struct {
	value *RemoveRolesPayload
	isSet bool
}

func (v NullableRemoveRolesPayload) Get() *RemoveRolesPayload {
	return v.value
}

func (v *NullableRemoveRolesPayload) Set(val *RemoveRolesPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveRolesPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveRolesPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveRolesPayload(val *RemoveRolesPayload) *NullableRemoveRolesPayload {
	return &NullableRemoveRolesPayload{value: val, isSet: true}
}

func (v NullableRemoveRolesPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveRolesPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
