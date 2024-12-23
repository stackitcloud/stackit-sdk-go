/*
Resource Manager API

API v2 to manage resource containers - organizations, folders, projects incl. labels  ### Resource Management STACKIT resource management handles the terms _Organization_, _Folder_, _Project_, _Label_, and the hierarchical structure between them. Technically, organizations,  folders, and projects are _Resource Containers_ to which a _Label_ can be attached to. The STACKIT _Resource Manager_ provides CRUD endpoints to query and to modify the state.  ### Organizations STACKIT organizations are the base element to create and to use cloud-resources. An organization is bound to one customer account. Organizations have a lifecycle. - Organizations are always the root node in resource hierarchy and do not have a parent  ### Projects STACKIT projects are needed to use cloud-resources. Projects serve as wrapper for underlying technical structures and processes. Projects have a lifecycle. Projects compared to folders may have different policies. - Projects are optional, but mandatory for cloud-resource usage - A project can be created having either an organization, or a folder as parent - A project must not have a project as parent - Project names under the same parent must not be unique - Root organization cannot be changed  ### Label STACKIT labels are key-value pairs including a resource container reference. Labels can be defined and attached freely to resource containers by which resources can be organized and queried. - Policy-based, immutable labels may exists

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resourcemanager

import (
	"encoding/json"
	"time"
)

// checks if the GetProjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectResponse{}

// GetProjectResponse struct for GetProjectResponse
type GetProjectResponse struct {
	// Globally unique identifier.
	// REQUIRED
	ContainerId *string `json:"containerId"`
	// Timestamp at which the project was created.
	// REQUIRED
	CreationTime *time.Time `json:"creationTime"`
	// Labels are key-value string pairs that can be attached to a resource container. Some labels may be enforced via policies.  - A label key must match the regex `[A-ZÄÜÖa-zäüöß0-9_-]{1,64}`. - A label value must match the regex `^$|[A-ZÄÜÖa-zäüöß0-9_-]{1,64}`.
	Labels *map[string]string `json:"labels,omitempty"`
	// REQUIRED
	LifecycleState *LifecycleState `json:"lifecycleState"`
	// Project name.
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Parent  *Parent            `json:"parent"`
	Parents *[]ParentListInner `json:"parents,omitempty"`
	// Globally unique identifier.
	// REQUIRED
	ProjectId *string `json:"projectId"`
	// Timestamp at which the project was last modified.
	// REQUIRED
	UpdateTime *time.Time `json:"updateTime"`
}

type _GetProjectResponse GetProjectResponse

// NewGetProjectResponse instantiates a new GetProjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectResponse(containerId *string, creationTime *time.Time, lifecycleState *LifecycleState, name *string, parent *Parent, projectId *string, updateTime *time.Time) *GetProjectResponse {
	this := GetProjectResponse{}
	this.ContainerId = containerId
	this.CreationTime = creationTime
	this.LifecycleState = lifecycleState
	this.Name = name
	this.Parent = parent
	this.ProjectId = projectId
	this.UpdateTime = updateTime
	return &this
}

// NewGetProjectResponseWithDefaults instantiates a new GetProjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectResponseWithDefaults() *GetProjectResponse {
	this := GetProjectResponse{}
	return &this
}

// GetContainerId returns the ContainerId field value
func (o *GetProjectResponse) GetContainerId() *string {
	if o == nil || IsNil(o.ContainerId) {
		var ret *string
		return ret
	}

	return o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value
// and a boolean to check if the value has been set.
func (o *GetProjectResponse) GetContainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerId, true
}

// SetContainerId sets field value
func (o *GetProjectResponse) SetContainerId(v *string) {
	o.ContainerId = v
}

// GetCreationTime returns the CreationTime field value
func (o *GetProjectResponse) GetCreationTime() *time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret *time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *GetProjectResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// SetCreationTime sets field value
func (o *GetProjectResponse) SetCreationTime(v *time.Time) {
	o.CreationTime = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *GetProjectResponse) GetLabels() *map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectResponse) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *GetProjectResponse) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *GetProjectResponse) SetLabels(v *map[string]string) {
	o.Labels = v
}

// GetLifecycleState returns the LifecycleState field value
func (o *GetProjectResponse) GetLifecycleState() *LifecycleState {
	if o == nil || IsNil(o.LifecycleState) {
		var ret *LifecycleState
		return ret
	}

	return o.LifecycleState
}

// GetLifecycleStateOk returns a tuple with the LifecycleState field value
// and a boolean to check if the value has been set.
func (o *GetProjectResponse) GetLifecycleStateOk() (*LifecycleState, bool) {
	if o == nil {
		return nil, false
	}
	return o.LifecycleState, true
}

// SetLifecycleState sets field value
func (o *GetProjectResponse) SetLifecycleState(v *LifecycleState) {
	o.LifecycleState = v
}

// GetName returns the Name field value
func (o *GetProjectResponse) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetProjectResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *GetProjectResponse) SetName(v *string) {
	o.Name = v
}

// GetParent returns the Parent field value
func (o *GetProjectResponse) GetParent() *Parent {
	if o == nil || IsNil(o.Parent) {
		var ret *Parent
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *GetProjectResponse) GetParentOk() (*Parent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent, true
}

// SetParent sets field value
func (o *GetProjectResponse) SetParent(v *Parent) {
	o.Parent = v
}

// GetParents returns the Parents field value if set, zero value otherwise.
func (o *GetProjectResponse) GetParents() *[]ParentListInner {
	if o == nil || IsNil(o.Parents) {
		var ret *[]ParentListInner
		return ret
	}
	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectResponse) GetParentsOk() (*[]ParentListInner, bool) {
	if o == nil || IsNil(o.Parents) {
		return nil, false
	}
	return o.Parents, true
}

// HasParents returns a boolean if a field has been set.
func (o *GetProjectResponse) HasParents() bool {
	if o != nil && !IsNil(o.Parents) {
		return true
	}

	return false
}

// SetParents gets a reference to the given []ParentListInner and assigns it to the Parents field.
func (o *GetProjectResponse) SetParents(v *[]ParentListInner) {
	o.Parents = v
}

// GetProjectId returns the ProjectId field value
func (o *GetProjectResponse) GetProjectId() *string {
	if o == nil || IsNil(o.ProjectId) {
		var ret *string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetProjectResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetProjectResponse) SetProjectId(v *string) {
	o.ProjectId = v
}

// GetUpdateTime returns the UpdateTime field value
func (o *GetProjectResponse) GetUpdateTime() *time.Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret *time.Time
		return ret
	}

	return o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value
// and a boolean to check if the value has been set.
func (o *GetProjectResponse) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// SetUpdateTime sets field value
func (o *GetProjectResponse) SetUpdateTime(v *time.Time) {
	o.UpdateTime = v
}

func (o GetProjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containerId"] = o.ContainerId
	toSerialize["creationTime"] = o.CreationTime
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["lifecycleState"] = o.LifecycleState
	toSerialize["name"] = o.Name
	toSerialize["parent"] = o.Parent
	if !IsNil(o.Parents) {
		toSerialize["parents"] = o.Parents
	}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["updateTime"] = o.UpdateTime
	return toSerialize, nil
}

type NullableGetProjectResponse struct {
	value *GetProjectResponse
	isSet bool
}

func (v NullableGetProjectResponse) Get() *GetProjectResponse {
	return v.value
}

func (v *NullableGetProjectResponse) Set(val *GetProjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectResponse(val *GetProjectResponse) *NullableGetProjectResponse {
	return &NullableGetProjectResponse{value: val, isSet: true}
}

func (v NullableGetProjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
