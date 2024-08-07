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

// checks if the ListSubjectIdsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSubjectIdsResponse{}

// ListSubjectIdsResponse struct for ListSubjectIdsResponse
type ListSubjectIdsResponse struct {
	Cursor *string `json:"cursor,omitempty"`
	// REQUIRED
	Items *[]string `json:"items"`
	// Can be cast to int32 without loss of precision.
	Limit *int64 `json:"limit,omitempty"`
}

type _ListSubjectIdsResponse ListSubjectIdsResponse

// NewListSubjectIdsResponse instantiates a new ListSubjectIdsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSubjectIdsResponse(items *[]string) *ListSubjectIdsResponse {
	this := ListSubjectIdsResponse{}
	this.Items = items
	return &this
}

// NewListSubjectIdsResponseWithDefaults instantiates a new ListSubjectIdsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSubjectIdsResponseWithDefaults() *ListSubjectIdsResponse {
	this := ListSubjectIdsResponse{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ListSubjectIdsResponse) GetCursor() *string {
	if o == nil || IsNil(o.Cursor) {
		var ret *string
		return ret
	}
	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSubjectIdsResponse) GetCursorOk() (*string, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ListSubjectIdsResponse) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *ListSubjectIdsResponse) SetCursor(v *string) {
	o.Cursor = v
}

// GetItems returns the Items field value
func (o *ListSubjectIdsResponse) GetItems() *[]string {
	if o == nil {
		var ret *[]string
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListSubjectIdsResponse) GetItemsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListSubjectIdsResponse) SetItems(v *[]string) {
	o.Items = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListSubjectIdsResponse) GetLimit() *int64 {
	if o == nil || IsNil(o.Limit) {
		var ret *int64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSubjectIdsResponse) GetLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListSubjectIdsResponse) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ListSubjectIdsResponse) SetLimit(v *int64) {
	o.Limit = v
}

func (o ListSubjectIdsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	toSerialize["items"] = o.Items
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}

type NullableListSubjectIdsResponse struct {
	value *ListSubjectIdsResponse
	isSet bool
}

func (v NullableListSubjectIdsResponse) Get() *ListSubjectIdsResponse {
	return v.value
}

func (v *NullableListSubjectIdsResponse) Set(val *ListSubjectIdsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListSubjectIdsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListSubjectIdsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSubjectIdsResponse(val *ListSubjectIdsResponse) *NullableListSubjectIdsResponse {
	return &NullableListSubjectIdsResponse{value: val, isSet: true}
}

func (v NullableListSubjectIdsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSubjectIdsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
