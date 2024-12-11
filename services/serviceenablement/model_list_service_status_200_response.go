/*
STACKIT Service Enablement API

STACKIT Service Enablement API

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceenablement

import (
	"encoding/json"
)

// checks if the ListServiceStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListServiceStatus200Response{}

// ListServiceStatus200Response struct for ListServiceStatus200Response
type ListServiceStatus200Response struct {
	Items      *[]ServiceStatus `json:"items,omitempty"`
	NextCursor *string          `json:"nextCursor,omitempty"`
}

// NewListServiceStatus200Response instantiates a new ListServiceStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListServiceStatus200Response() *ListServiceStatus200Response {
	this := ListServiceStatus200Response{}
	return &this
}

// NewListServiceStatus200ResponseWithDefaults instantiates a new ListServiceStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListServiceStatus200ResponseWithDefaults() *ListServiceStatus200Response {
	this := ListServiceStatus200Response{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListServiceStatus200Response) GetItems() *[]ServiceStatus {
	if o == nil || IsNil(o.Items) {
		var ret *[]ServiceStatus
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceStatus200Response) GetItemsOk() (*[]ServiceStatus, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListServiceStatus200Response) HasItems() bool {
	if o != nil && !IsNil(o.Items) && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ServiceStatus and assigns it to the Items field.
func (o *ListServiceStatus200Response) SetItems(v *[]ServiceStatus) {
	o.Items = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *ListServiceStatus200Response) GetNextCursor() *string {
	if o == nil || IsNil(o.NextCursor) {
		var ret *string
		return ret
	}
	return o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListServiceStatus200Response) GetNextCursorOk() (*string, bool) {
	if o == nil || IsNil(o.NextCursor) {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *ListServiceStatus200Response) HasNextCursor() bool {
	if o != nil && !IsNil(o.NextCursor) && !IsNil(o.NextCursor) {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *ListServiceStatus200Response) SetNextCursor(v *string) {
	o.NextCursor = v
}

func (o ListServiceStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.NextCursor) {
		toSerialize["nextCursor"] = o.NextCursor
	}
	return toSerialize, nil
}

type NullableListServiceStatus200Response struct {
	value *ListServiceStatus200Response
	isSet bool
}

func (v NullableListServiceStatus200Response) Get() *ListServiceStatus200Response {
	return v.value
}

func (v *NullableListServiceStatus200Response) Set(val *ListServiceStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListServiceStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListServiceStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListServiceStatus200Response(val *ListServiceStatus200Response) *NullableListServiceStatus200Response {
	return &NullableListServiceStatus200Response{value: val, isSet: true}
}

func (v NullableListServiceStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListServiceStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
