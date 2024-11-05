/*
Resource Manager API

API v2 to manage resource containers - organizations, folders, projects incl. labels  ### Resource Management STACKIT resource management handles the terms _Organization_, _Folder_, _Project_, _Label_, and the hierarchical structure between them. Technically, organizations,  folders, and projects are _Resource Containers_ to which a _Label_ can be attached to. The STACKIT _Resource Manager_ provides CRUD endpoints to query and to modify the state.  ### Organizations STACKIT organizations are the base element to create and to use cloud-resources. An organization is bound to one customer account. Organizations have a lifecycle. - Organizations are always the root node in resource hierarchy and do not have a parent  ### Projects STACKIT projects are needed to use cloud-resources. Projects serve as wrapper for underlying technical structures and processes. Projects have a lifecycle. Projects compared to folders may have different policies. - Projects are optional, but mandatory for cloud-resource usage - A project can be created having either an organization, or a folder as parent - A project must not have a project as parent - Project names under the same parent must not be unique - Root organization cannot be changed  ### Label STACKIT labels are key-value pairs including a resource container reference. Labels can be defined and attached freely to resource containers by which resources can be organized and queried. - Policy-based, immutable labels may exists

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resourcemanager

import (
	"encoding/json"
)

// checks if the ListOrganizationContainersResponseItemsInnerAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationContainersResponseItemsInnerAnyOf{}

// ListOrganizationContainersResponseItemsInnerAnyOf struct for ListOrganizationContainersResponseItemsInnerAnyOf
type ListOrganizationContainersResponseItemsInnerAnyOf struct {
	// REQUIRED
	Item *FolderResponse `json:"item"`
	// Resource container type.
	// REQUIRED
	Type *string `json:"type"`
}

type _ListOrganizationContainersResponseItemsInnerAnyOf ListOrganizationContainersResponseItemsInnerAnyOf

// NewListOrganizationContainersResponseItemsInnerAnyOf instantiates a new ListOrganizationContainersResponseItemsInnerAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationContainersResponseItemsInnerAnyOf(item *FolderResponse, type_ *string) *ListOrganizationContainersResponseItemsInnerAnyOf {
	this := ListOrganizationContainersResponseItemsInnerAnyOf{}
	this.Item = item
	this.Type = type_
	return &this
}

// NewListOrganizationContainersResponseItemsInnerAnyOfWithDefaults instantiates a new ListOrganizationContainersResponseItemsInnerAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationContainersResponseItemsInnerAnyOfWithDefaults() *ListOrganizationContainersResponseItemsInnerAnyOf {
	this := ListOrganizationContainersResponseItemsInnerAnyOf{}
	return &this
}

// GetItem returns the Item field value
func (o *ListOrganizationContainersResponseItemsInnerAnyOf) GetItem() *FolderResponse {
	if o == nil {
		var ret *FolderResponse
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationContainersResponseItemsInnerAnyOf) GetItemOk() (*FolderResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Item, true
}

// SetItem sets field value
func (o *ListOrganizationContainersResponseItemsInnerAnyOf) SetItem(v *FolderResponse) {
	o.Item = v
}

// GetType returns the Type field value
func (o *ListOrganizationContainersResponseItemsInnerAnyOf) GetType() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationContainersResponseItemsInnerAnyOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *ListOrganizationContainersResponseItemsInnerAnyOf) SetType(v *string) {
	o.Type = v
}

func (o ListOrganizationContainersResponseItemsInnerAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["item"] = o.Item
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableListOrganizationContainersResponseItemsInnerAnyOf struct {
	value *ListOrganizationContainersResponseItemsInnerAnyOf
	isSet bool
}

func (v NullableListOrganizationContainersResponseItemsInnerAnyOf) Get() *ListOrganizationContainersResponseItemsInnerAnyOf {
	return v.value
}

func (v *NullableListOrganizationContainersResponseItemsInnerAnyOf) Set(val *ListOrganizationContainersResponseItemsInnerAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationContainersResponseItemsInnerAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationContainersResponseItemsInnerAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationContainersResponseItemsInnerAnyOf(val *ListOrganizationContainersResponseItemsInnerAnyOf) *NullableListOrganizationContainersResponseItemsInnerAnyOf {
	return &NullableListOrganizationContainersResponseItemsInnerAnyOf{value: val, isSet: true}
}

func (v NullableListOrganizationContainersResponseItemsInnerAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationContainersResponseItemsInnerAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}