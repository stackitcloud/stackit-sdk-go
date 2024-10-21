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

// checks if the ListOrganizationContainersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationContainersResponse{}

// ListOrganizationContainersResponse struct for ListOrganizationContainersResponse
type ListOrganizationContainersResponse struct {
	// A pagination cursor is returned on the first call of the pagination process. If given, it will start from the end of the previous position. If not given, a new pagination is started.
	// REQUIRED
	Cursor *string `json:"cursor"`
	// REQUIRED
	Items *[]ListOrganizationContainersResponseItemsInner `json:"items"`
	// The maximum number of projects to return in the response. If not present, an appropriate default will be used.
	// REQUIRED
	Limit *float64 `json:"limit"`
}

type _ListOrganizationContainersResponse ListOrganizationContainersResponse

// NewListOrganizationContainersResponse instantiates a new ListOrganizationContainersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationContainersResponse(cursor *string, items *[]ListOrganizationContainersResponseItemsInner, limit *float64) *ListOrganizationContainersResponse {
	this := ListOrganizationContainersResponse{}
	this.Cursor = cursor
	this.Items = items
	this.Limit = limit
	return &this
}

// NewListOrganizationContainersResponseWithDefaults instantiates a new ListOrganizationContainersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationContainersResponseWithDefaults() *ListOrganizationContainersResponse {
	this := ListOrganizationContainersResponse{}
	var limit float64 = 50
	this.Limit = &limit
	return &this
}

// GetCursor returns the Cursor field value
func (o *ListOrganizationContainersResponse) GetCursor() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationContainersResponse) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cursor, true
}

// SetCursor sets field value
func (o *ListOrganizationContainersResponse) SetCursor(v *string) {
	o.Cursor = v
}

// GetItems returns the Items field value
func (o *ListOrganizationContainersResponse) GetItems() *[]ListOrganizationContainersResponseItemsInner {
	if o == nil {
		var ret *[]ListOrganizationContainersResponseItemsInner
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationContainersResponse) GetItemsOk() (*[]ListOrganizationContainersResponseItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListOrganizationContainersResponse) SetItems(v *[]ListOrganizationContainersResponseItemsInner) {
	o.Items = v
}

// GetLimit returns the Limit field value
func (o *ListOrganizationContainersResponse) GetLimit() *float64 {
	if o == nil {
		var ret *float64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationContainersResponse) GetLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit, true
}

// SetLimit sets field value
func (o *ListOrganizationContainersResponse) SetLimit(v *float64) {
	o.Limit = v
}

func (o ListOrganizationContainersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cursor"] = o.Cursor
	toSerialize["items"] = o.Items
	toSerialize["limit"] = o.Limit
	return toSerialize, nil
}

type NullableListOrganizationContainersResponse struct {
	value *ListOrganizationContainersResponse
	isSet bool
}

func (v NullableListOrganizationContainersResponse) Get() *ListOrganizationContainersResponse {
	return v.value
}

func (v *NullableListOrganizationContainersResponse) Set(val *ListOrganizationContainersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationContainersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationContainersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationContainersResponse(val *ListOrganizationContainersResponse) *NullableListOrganizationContainersResponse {
	return &NullableListOrganizationContainersResponse{value: val, isSet: true}
}

func (v NullableListOrganizationContainersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationContainersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
