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

// checks if the CreateProjectPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProjectPayload{}

/*
	types and functions for containerParentId
*/

// isNotNullableString
type CreateProjectPayloadGetContainerParentIdAttributeType = *string

func getCreateProjectPayloadGetContainerParentIdAttributeTypeOk(arg CreateProjectPayloadGetContainerParentIdAttributeType) (ret CreateProjectPayloadGetContainerParentIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateProjectPayloadGetContainerParentIdAttributeType(arg *CreateProjectPayloadGetContainerParentIdAttributeType, val CreateProjectPayloadGetContainerParentIdRetType) {
	*arg = &val
}

type CreateProjectPayloadGetContainerParentIdArgType = string
type CreateProjectPayloadGetContainerParentIdRetType = string

/*
	types and functions for labels
*/

// isContainer
type CreateProjectPayloadGetLabelsAttributeType = *map[string]string
type CreateProjectPayloadGetLabelsArgType = map[string]string
type CreateProjectPayloadGetLabelsRetType = map[string]string

func getCreateProjectPayloadGetLabelsAttributeTypeOk(arg CreateProjectPayloadGetLabelsAttributeType) (ret CreateProjectPayloadGetLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateProjectPayloadGetLabelsAttributeType(arg *CreateProjectPayloadGetLabelsAttributeType, val CreateProjectPayloadGetLabelsRetType) {
	*arg = &val
}

/*
	types and functions for members
*/

// isArray
type CreateProjectPayloadGetMembersAttributeType = *[]Member
type CreateProjectPayloadGetMembersArgType = []Member
type CreateProjectPayloadGetMembersRetType = []Member

func getCreateProjectPayloadGetMembersAttributeTypeOk(arg CreateProjectPayloadGetMembersAttributeType) (ret CreateProjectPayloadGetMembersRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateProjectPayloadGetMembersAttributeType(arg *CreateProjectPayloadGetMembersAttributeType, val CreateProjectPayloadGetMembersRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type CreateProjectPayloadGetNameAttributeType = *string

func getCreateProjectPayloadGetNameAttributeTypeOk(arg CreateProjectPayloadGetNameAttributeType) (ret CreateProjectPayloadGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateProjectPayloadGetNameAttributeType(arg *CreateProjectPayloadGetNameAttributeType, val CreateProjectPayloadGetNameRetType) {
	*arg = &val
}

type CreateProjectPayloadGetNameArgType = string
type CreateProjectPayloadGetNameRetType = string

// CreateProjectPayload struct for CreateProjectPayload
type CreateProjectPayload struct {
	// Identifier of the parent resource container - containerId as well as UUID identifier is supported.
	// REQUIRED
	ContainerParentId CreateProjectPayloadGetContainerParentIdAttributeType `json:"containerParentId" required:"true"`
	// Labels are key-value string pairs that can be attached to a resource container. Some labels may be enforced via policies.  - A label key must match the regex `[A-ZÄÜÖa-zäüöß0-9_-]{1,64}`. - A label value must match the regex `^$|[A-ZÄÜÖa-zäüöß0-9_-]{1,64}`.
	Labels CreateProjectPayloadGetLabelsAttributeType `json:"labels,omitempty"`
	// The initial members assigned to the project. At least one subject needs to be a user, and not a client or service account.
	// REQUIRED
	Members CreateProjectPayloadGetMembersAttributeType `json:"members" required:"true"`
	// Project name matching the regex `^[a-zA-ZäüöÄÜÖ0-9]( ?[a-zA-ZäüöÄÜÖß0-9_+&-]){0,39}$`.
	// REQUIRED
	Name CreateProjectPayloadGetNameAttributeType `json:"name" required:"true"`
}

type _CreateProjectPayload CreateProjectPayload

// NewCreateProjectPayload instantiates a new CreateProjectPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectPayload(containerParentId CreateProjectPayloadGetContainerParentIdArgType, members CreateProjectPayloadGetMembersArgType, name CreateProjectPayloadGetNameArgType) *CreateProjectPayload {
	this := CreateProjectPayload{}
	setCreateProjectPayloadGetContainerParentIdAttributeType(&this.ContainerParentId, containerParentId)
	setCreateProjectPayloadGetMembersAttributeType(&this.Members, members)
	setCreateProjectPayloadGetNameAttributeType(&this.Name, name)
	return &this
}

// NewCreateProjectPayloadWithDefaults instantiates a new CreateProjectPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectPayloadWithDefaults() *CreateProjectPayload {
	this := CreateProjectPayload{}
	return &this
}

// GetContainerParentId returns the ContainerParentId field value
func (o *CreateProjectPayload) GetContainerParentId() (ret CreateProjectPayloadGetContainerParentIdRetType) {
	ret, _ = o.GetContainerParentIdOk()
	return ret
}

// GetContainerParentIdOk returns a tuple with the ContainerParentId field value
// and a boolean to check if the value has been set.
func (o *CreateProjectPayload) GetContainerParentIdOk() (ret CreateProjectPayloadGetContainerParentIdRetType, ok bool) {
	return getCreateProjectPayloadGetContainerParentIdAttributeTypeOk(o.ContainerParentId)
}

// SetContainerParentId sets field value
func (o *CreateProjectPayload) SetContainerParentId(v CreateProjectPayloadGetContainerParentIdRetType) {
	setCreateProjectPayloadGetContainerParentIdAttributeType(&o.ContainerParentId, v)
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CreateProjectPayload) GetLabels() (res CreateProjectPayloadGetLabelsRetType) {
	res, _ = o.GetLabelsOk()
	return
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectPayload) GetLabelsOk() (ret CreateProjectPayloadGetLabelsRetType, ok bool) {
	return getCreateProjectPayloadGetLabelsAttributeTypeOk(o.Labels)
}

// HasLabels returns a boolean if a field has been set.
func (o *CreateProjectPayload) HasLabels() bool {
	_, ok := o.GetLabelsOk()
	return ok
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *CreateProjectPayload) SetLabels(v CreateProjectPayloadGetLabelsRetType) {
	setCreateProjectPayloadGetLabelsAttributeType(&o.Labels, v)
}

// GetMembers returns the Members field value
func (o *CreateProjectPayload) GetMembers() (ret CreateProjectPayloadGetMembersRetType) {
	ret, _ = o.GetMembersOk()
	return ret
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *CreateProjectPayload) GetMembersOk() (ret CreateProjectPayloadGetMembersRetType, ok bool) {
	return getCreateProjectPayloadGetMembersAttributeTypeOk(o.Members)
}

// SetMembers sets field value
func (o *CreateProjectPayload) SetMembers(v CreateProjectPayloadGetMembersRetType) {
	setCreateProjectPayloadGetMembersAttributeType(&o.Members, v)
}

// GetName returns the Name field value
func (o *CreateProjectPayload) GetName() (ret CreateProjectPayloadGetNameRetType) {
	ret, _ = o.GetNameOk()
	return ret
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProjectPayload) GetNameOk() (ret CreateProjectPayloadGetNameRetType, ok bool) {
	return getCreateProjectPayloadGetNameAttributeTypeOk(o.Name)
}

// SetName sets field value
func (o *CreateProjectPayload) SetName(v CreateProjectPayloadGetNameRetType) {
	setCreateProjectPayloadGetNameAttributeType(&o.Name, v)
}

func (o CreateProjectPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateProjectPayloadGetContainerParentIdAttributeTypeOk(o.ContainerParentId); ok {
		toSerialize["ContainerParentId"] = val
	}
	if val, ok := getCreateProjectPayloadGetLabelsAttributeTypeOk(o.Labels); ok {
		toSerialize["Labels"] = val
	}
	if val, ok := getCreateProjectPayloadGetMembersAttributeTypeOk(o.Members); ok {
		toSerialize["Members"] = val
	}
	if val, ok := getCreateProjectPayloadGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	return toSerialize, nil
}

type NullableCreateProjectPayload struct {
	value *CreateProjectPayload
	isSet bool
}

func (v NullableCreateProjectPayload) Get() *CreateProjectPayload {
	return v.value
}

func (v *NullableCreateProjectPayload) Set(val *CreateProjectPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectPayload(val *CreateProjectPayload) *NullableCreateProjectPayload {
	return &NullableCreateProjectPayload{value: val, isSet: true}
}

func (v NullableCreateProjectPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
