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

// checks if the AddRolesPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddRolesPayload{}

// AddRolesPayload struct for AddRolesPayload
type AddRolesPayload struct {
	// REQUIRED
	ResourceType *string `json:"resourceType"`
	// REQUIRED
	Roles *[]AddRolesPayloadItem `json:"roles"`
}

type _AddRolesPayload AddRolesPayload

// NewAddRolesPayload instantiates a new AddRolesPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddRolesPayload(resourceType *string, roles *[]AddRolesPayloadItem) *AddRolesPayload {
	this := AddRolesPayload{}
	this.ResourceType = resourceType
	this.Roles = roles
	return &this
}

// NewAddRolesPayloadWithDefaults instantiates a new AddRolesPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddRolesPayloadWithDefaults() *AddRolesPayload {
	this := AddRolesPayload{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *AddRolesPayload) GetResourceType() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *AddRolesPayload) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// SetResourceType sets field value
func (o *AddRolesPayload) SetResourceType(v *string) {
	o.ResourceType = v
}

// GetRoles returns the Roles field value
func (o *AddRolesPayload) GetRoles() *[]AddRolesPayloadItem {
	if o == nil {
		var ret *[]AddRolesPayloadItem
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *AddRolesPayload) GetRolesOk() (*[]AddRolesPayloadItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *AddRolesPayload) SetRoles(v *[]AddRolesPayloadItem) {
	o.Roles = v
}

func (o AddRolesPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

type NullableAddRolesPayload struct {
	value *AddRolesPayload
	isSet bool
}

func (v NullableAddRolesPayload) Get() *AddRolesPayload {
	return v.value
}

func (v *NullableAddRolesPayload) Set(val *AddRolesPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAddRolesPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAddRolesPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddRolesPayload(val *AddRolesPayload) *NullableAddRolesPayload {
	return &NullableAddRolesPayload{value: val, isSet: true}
}

func (v NullableAddRolesPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddRolesPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
