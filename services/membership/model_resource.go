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

// checks if the Resource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Resource{}

// Resource struct for Resource
type Resource struct {
	ResourceAlias *string `json:"resourceAlias,omitempty"`
	// REQUIRED
	ResourceId *string `json:"resourceId"`
	// REQUIRED
	ResourceType *string `json:"resourceType"`
}

type _Resource Resource

// NewResource instantiates a new Resource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResource(resourceId *string, resourceType *string) *Resource {
	this := Resource{}
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	return &this
}

// NewResourceWithDefaults instantiates a new Resource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithDefaults() *Resource {
	this := Resource{}
	return &this
}

// GetResourceAlias returns the ResourceAlias field value if set, zero value otherwise.
func (o *Resource) GetResourceAlias() *string {
	if o == nil || IsNil(o.ResourceAlias) {
		var ret *string
		return ret
	}
	return o.ResourceAlias
}

// GetResourceAliasOk returns a tuple with the ResourceAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetResourceAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceAlias) {
		return nil, false
	}
	return o.ResourceAlias, true
}

// HasResourceAlias returns a boolean if a field has been set.
func (o *Resource) HasResourceAlias() bool {
	if o != nil && !IsNil(o.ResourceAlias) && !IsNil(o.ResourceAlias) {
		return true
	}

	return false
}

// SetResourceAlias gets a reference to the given string and assigns it to the ResourceAlias field.
func (o *Resource) SetResourceAlias(v *string) {
	o.ResourceAlias = v
}

// GetResourceId returns the ResourceId field value
func (o *Resource) GetResourceId() *string {
	if o == nil || IsNil(o.ResourceId) {
		var ret *string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *Resource) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// SetResourceId sets field value
func (o *Resource) SetResourceId(v *string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value
func (o *Resource) GetResourceType() *string {
	if o == nil || IsNil(o.ResourceType) {
		var ret *string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *Resource) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// SetResourceType sets field value
func (o *Resource) SetResourceType(v *string) {
	o.ResourceType = v
}

func (o Resource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceAlias) {
		toSerialize["resourceAlias"] = o.ResourceAlias
	}
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["resourceType"] = o.ResourceType
	return toSerialize, nil
}

type NullableResource struct {
	value *Resource
	isSet bool
}

func (v NullableResource) Get() *Resource {
	return v.value
}

func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

func (v NullableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

func (v NullableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
