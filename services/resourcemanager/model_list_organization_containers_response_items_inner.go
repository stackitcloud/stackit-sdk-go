/*
Resource Manager API

API v2 to manage resource containers - organizations, folders, projects incl. labels  ### Resource Management STACKIT resource management handles the terms _Organization_, _Folder_, _Project_, _Label_, and the hierarchical structure between them. Technically, organizations,  folders, and projects are _Resource Containers_ to which a _Label_ can be attached to. The STACKIT _Resource Manager_ provides CRUD endpoints to query and to modify the state.  ### Organizations STACKIT organizations are the base element to create and to use cloud-resources. An organization is bound to one customer account. Organizations have a lifecycle. - Organizations are always the root node in resource hierarchy and do not have a parent  ### Projects STACKIT projects are needed to use cloud-resources. Projects serve as wrapper for underlying technical structures and processes. Projects have a lifecycle. Projects compared to folders may have different policies. - Projects are optional, but mandatory for cloud-resource usage - A project can be created having either an organization, or a folder as parent - A project must not have a project as parent - Project names under the same parent must not be unique - Root organization cannot be changed  ### Label STACKIT labels are key-value pairs including a resource container reference. Labels can be defined and attached freely to resource containers by which resources can be organized and queried. - Policy-based, immutable labels may exists

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resourcemanager

import (
	"encoding/json"
	"fmt"
)

// ListOrganizationContainersResponseItemsInner struct for ListOrganizationContainersResponseItemsInner
type ListOrganizationContainersResponseItemsInner struct {
	ListOrganizationContainersResponseItemsInnerAnyOf  *ListOrganizationContainersResponseItemsInnerAnyOf
	ListOrganizationContainersResponseItemsInnerAnyOf1 *ListOrganizationContainersResponseItemsInnerAnyOf1
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ListOrganizationContainersResponseItemsInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ListOrganizationContainersResponseItemsInnerAnyOf
	err = json.Unmarshal(data, &dst.ListOrganizationContainersResponseItemsInnerAnyOf)
	if err == nil {
		jsonListOrganizationContainersResponseItemsInnerAnyOf, _ := json.Marshal(dst.ListOrganizationContainersResponseItemsInnerAnyOf)
		if string(jsonListOrganizationContainersResponseItemsInnerAnyOf) == "{}" { // empty struct
			dst.ListOrganizationContainersResponseItemsInnerAnyOf = nil
		} else {
			return nil // data stored in dst.ListOrganizationContainersResponseItemsInnerAnyOf, return on the first match
		}
	} else {
		dst.ListOrganizationContainersResponseItemsInnerAnyOf = nil
	}

	// try to unmarshal JSON data into ListOrganizationContainersResponseItemsInnerAnyOf1
	err = json.Unmarshal(data, &dst.ListOrganizationContainersResponseItemsInnerAnyOf1)
	if err == nil {
		jsonListOrganizationContainersResponseItemsInnerAnyOf1, _ := json.Marshal(dst.ListOrganizationContainersResponseItemsInnerAnyOf1)
		if string(jsonListOrganizationContainersResponseItemsInnerAnyOf1) == "{}" { // empty struct
			dst.ListOrganizationContainersResponseItemsInnerAnyOf1 = nil
		} else {
			return nil // data stored in dst.ListOrganizationContainersResponseItemsInnerAnyOf1, return on the first match
		}
	} else {
		dst.ListOrganizationContainersResponseItemsInnerAnyOf1 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ListOrganizationContainersResponseItemsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ListOrganizationContainersResponseItemsInner) MarshalJSON() ([]byte, error) {
	if src.ListOrganizationContainersResponseItemsInnerAnyOf != nil {
		return json.Marshal(&src.ListOrganizationContainersResponseItemsInnerAnyOf)
	}

	if src.ListOrganizationContainersResponseItemsInnerAnyOf1 != nil {
		return json.Marshal(&src.ListOrganizationContainersResponseItemsInnerAnyOf1)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableListOrganizationContainersResponseItemsInner struct {
	value *ListOrganizationContainersResponseItemsInner
	isSet bool
}

func (v NullableListOrganizationContainersResponseItemsInner) Get() *ListOrganizationContainersResponseItemsInner {
	return v.value
}

func (v *NullableListOrganizationContainersResponseItemsInner) Set(val *ListOrganizationContainersResponseItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationContainersResponseItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationContainersResponseItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationContainersResponseItemsInner(val *ListOrganizationContainersResponseItemsInner) *NullableListOrganizationContainersResponseItemsInner {
	return &NullableListOrganizationContainersResponseItemsInner{value: val, isSet: true}
}

func (v NullableListOrganizationContainersResponseItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationContainersResponseItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}