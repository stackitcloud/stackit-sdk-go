/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
	"time"
)

// checks if the NetworkArea type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkArea{}

// NetworkArea Object that represents a network area.
type NetworkArea struct {
	// Universally Unique Identifier (UUID).
	// REQUIRED
	AreaId *string `json:"areaId"`
	// Date-time when resource was created.
	CreatedAt *time.Time       `json:"createdAt,omitempty"`
	Ipv4      *NetworkAreaIPv4 `json:"ipv4,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// REQUIRED
	Name *string `json:"name"`
	// The amount of projects currently referencing a specific area.
	// REQUIRED
	ProjectCount *int64 `json:"projectCount"`
	// The state of a resource object.
	// REQUIRED
	State *string `json:"state"`
	// Date-time when resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type _NetworkArea NetworkArea

// NewNetworkArea instantiates a new NetworkArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkArea(areaId *string, name *string, projectCount *int64, state *string) *NetworkArea {
	this := NetworkArea{}
	this.AreaId = areaId
	this.Name = name
	this.ProjectCount = projectCount
	this.State = state
	return &this
}

// NewNetworkAreaWithDefaults instantiates a new NetworkArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAreaWithDefaults() *NetworkArea {
	this := NetworkArea{}
	return &this
}

// GetAreaId returns the AreaId field value
func (o *NetworkArea) GetAreaId() *string {
	if o == nil || IsNil(o.AreaId) {
		var ret *string
		return ret
	}

	return o.AreaId
}

// GetAreaIdOk returns a tuple with the AreaId field value
// and a boolean to check if the value has been set.
func (o *NetworkArea) GetAreaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AreaId, true
}

// SetAreaId sets field value
func (o *NetworkArea) SetAreaId(v *string) {
	o.AreaId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NetworkArea) GetCreatedAt() *time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret *time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkArea) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NetworkArea) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NetworkArea) SetCreatedAt(v *time.Time) {
	o.CreatedAt = v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *NetworkArea) GetIpv4() *NetworkAreaIPv4 {
	if o == nil || IsNil(o.Ipv4) {
		var ret *NetworkAreaIPv4
		return ret
	}
	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkArea) GetIpv4Ok() (*NetworkAreaIPv4, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *NetworkArea) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given NetworkAreaIPv4 and assigns it to the Ipv4 field.
func (o *NetworkArea) SetIpv4(v *NetworkAreaIPv4) {
	o.Ipv4 = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *NetworkArea) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkArea) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *NetworkArea) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *NetworkArea) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value
func (o *NetworkArea) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkArea) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *NetworkArea) SetName(v *string) {
	o.Name = v
}

// GetProjectCount returns the ProjectCount field value
func (o *NetworkArea) GetProjectCount() *int64 {
	if o == nil || IsNil(o.ProjectCount) {
		var ret *int64
		return ret
	}

	return o.ProjectCount
}

// GetProjectCountOk returns a tuple with the ProjectCount field value
// and a boolean to check if the value has been set.
func (o *NetworkArea) GetProjectCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectCount, true
}

// SetProjectCount sets field value
func (o *NetworkArea) SetProjectCount(v *int64) {
	o.ProjectCount = v
}

// GetState returns the State field value
func (o *NetworkArea) GetState() *string {
	if o == nil || IsNil(o.State) {
		var ret *string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *NetworkArea) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State, true
}

// SetState sets field value
func (o *NetworkArea) SetState(v *string) {
	o.State = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NetworkArea) GetUpdatedAt() *time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret *time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkArea) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NetworkArea) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NetworkArea) SetUpdatedAt(v *time.Time) {
	o.UpdatedAt = v
}

func (o NetworkArea) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["areaId"] = o.AreaId
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["name"] = o.Name
	toSerialize["projectCount"] = o.ProjectCount
	toSerialize["state"] = o.State
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableNetworkArea struct {
	value *NetworkArea
	isSet bool
}

func (v NullableNetworkArea) Get() *NetworkArea {
	return v.value
}

func (v *NullableNetworkArea) Set(val *NetworkArea) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkArea) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkArea(val *NetworkArea) *NullableNetworkArea {
	return &NullableNetworkArea{value: val, isSet: true}
}

func (v NullableNetworkArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
