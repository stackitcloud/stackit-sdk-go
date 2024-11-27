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

// checks if the CreateResourcePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateResourcePayload{}

// CreateResourcePayload struct for CreateResourcePayload
type CreateResourcePayload struct {
	Members *[]Member `json:"members,omitempty"`
	// REQUIRED
	ParentId *string `json:"parentId"`
	// REQUIRED
	ParentType    *string `json:"parentType"`
	ResourceAlias *string `json:"resourceAlias,omitempty"`
	// REQUIRED
	ResourceId *string `json:"resourceId"`
}

type _CreateResourcePayload CreateResourcePayload

// NewCreateResourcePayload instantiates a new CreateResourcePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateResourcePayload(parentId *string, parentType *string, resourceId *string) *CreateResourcePayload {
	this := CreateResourcePayload{}
	this.ParentId = parentId
	this.ParentType = parentType
	this.ResourceId = resourceId
	return &this
}

// NewCreateResourcePayloadWithDefaults instantiates a new CreateResourcePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateResourcePayloadWithDefaults() *CreateResourcePayload {
	this := CreateResourcePayload{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *CreateResourcePayload) GetMembers() *[]Member {
	if o == nil || IsNil(o.Members) {
		var ret *[]Member
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourcePayload) GetMembersOk() (*[]Member, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *CreateResourcePayload) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []Member and assigns it to the Members field.
func (o *CreateResourcePayload) SetMembers(v *[]Member) {
	o.Members = v
}

// GetParentId returns the ParentId field value
func (o *CreateResourcePayload) GetParentId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *CreateResourcePayload) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId, true
}

// SetParentId sets field value
func (o *CreateResourcePayload) SetParentId(v *string) {
	o.ParentId = v
}

// GetParentType returns the ParentType field value
func (o *CreateResourcePayload) GetParentType() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.ParentType
}

// GetParentTypeOk returns a tuple with the ParentType field value
// and a boolean to check if the value has been set.
func (o *CreateResourcePayload) GetParentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentType, true
}

// SetParentType sets field value
func (o *CreateResourcePayload) SetParentType(v *string) {
	o.ParentType = v
}

// GetResourceAlias returns the ResourceAlias field value if set, zero value otherwise.
func (o *CreateResourcePayload) GetResourceAlias() *string {
	if o == nil || IsNil(o.ResourceAlias) {
		var ret *string
		return ret
	}
	return o.ResourceAlias
}

// GetResourceAliasOk returns a tuple with the ResourceAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourcePayload) GetResourceAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceAlias) {
		return nil, false
	}
	return o.ResourceAlias, true
}

// HasResourceAlias returns a boolean if a field has been set.
func (o *CreateResourcePayload) HasResourceAlias() bool {
	if o != nil && !IsNil(o.ResourceAlias) {
		return true
	}

	return false
}

// SetResourceAlias gets a reference to the given string and assigns it to the ResourceAlias field.
func (o *CreateResourcePayload) SetResourceAlias(v *string) {
	o.ResourceAlias = v
}

// GetResourceId returns the ResourceId field value
func (o *CreateResourcePayload) GetResourceId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *CreateResourcePayload) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// SetResourceId sets field value
func (o *CreateResourcePayload) SetResourceId(v *string) {
	o.ResourceId = v
}

func (o CreateResourcePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	toSerialize["parentId"] = o.ParentId
	toSerialize["parentType"] = o.ParentType
	if !IsNil(o.ResourceAlias) {
		toSerialize["resourceAlias"] = o.ResourceAlias
	}
	toSerialize["resourceId"] = o.ResourceId
	return toSerialize, nil
}

type NullableCreateResourcePayload struct {
	value *CreateResourcePayload
	isSet bool
}

func (v NullableCreateResourcePayload) Get() *CreateResourcePayload {
	return v.value
}

func (v *NullableCreateResourcePayload) Set(val *CreateResourcePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateResourcePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateResourcePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateResourcePayload(val *CreateResourcePayload) *NullableCreateResourcePayload {
	return &NullableCreateResourcePayload{value: val, isSet: true}
}

func (v NullableCreateResourcePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateResourcePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
