/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

import (
	"encoding/json"
)

// checks if the ExtensionsExtensionListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionsExtensionListResponse{}

/*
	types and functions for list
*/

// isArray
type ExtensionsExtensionListResponseGetListAttributeType = *[]ApiExtensionList
type ExtensionsExtensionListResponseGetListArgType = []ApiExtensionList
type ExtensionsExtensionListResponseGetListRetType = []ApiExtensionList

func getExtensionsExtensionListResponseGetListAttributeTypeOk(arg ExtensionsExtensionListResponseGetListAttributeType) (ret ExtensionsExtensionListResponseGetListRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setExtensionsExtensionListResponseGetListAttributeType(arg *ExtensionsExtensionListResponseGetListAttributeType, val ExtensionsExtensionListResponseGetListRetType) {
	*arg = &val
}

// ExtensionsExtensionListResponse struct for ExtensionsExtensionListResponse
type ExtensionsExtensionListResponse struct {
	List ExtensionsExtensionListResponseGetListAttributeType `json:"list,omitempty"`
}

// NewExtensionsExtensionListResponse instantiates a new ExtensionsExtensionListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionsExtensionListResponse() *ExtensionsExtensionListResponse {
	this := ExtensionsExtensionListResponse{}
	return &this
}

// NewExtensionsExtensionListResponseWithDefaults instantiates a new ExtensionsExtensionListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionsExtensionListResponseWithDefaults() *ExtensionsExtensionListResponse {
	this := ExtensionsExtensionListResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ExtensionsExtensionListResponse) GetList() (res ExtensionsExtensionListResponseGetListRetType) {
	res, _ = o.GetListOk()
	return
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionsExtensionListResponse) GetListOk() (ret ExtensionsExtensionListResponseGetListRetType, ok bool) {
	return getExtensionsExtensionListResponseGetListAttributeTypeOk(o.List)
}

// HasList returns a boolean if a field has been set.
func (o *ExtensionsExtensionListResponse) HasList() bool {
	_, ok := o.GetListOk()
	return ok
}

// SetList gets a reference to the given []ApiExtensionList and assigns it to the List field.
func (o *ExtensionsExtensionListResponse) SetList(v ExtensionsExtensionListResponseGetListRetType) {
	setExtensionsExtensionListResponseGetListAttributeType(&o.List, v)
}

func (o ExtensionsExtensionListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getExtensionsExtensionListResponseGetListAttributeTypeOk(o.List); ok {
		toSerialize["List"] = val
	}
	return toSerialize, nil
}

type NullableExtensionsExtensionListResponse struct {
	value *ExtensionsExtensionListResponse
	isSet bool
}

func (v NullableExtensionsExtensionListResponse) Get() *ExtensionsExtensionListResponse {
	return v.value
}

func (v *NullableExtensionsExtensionListResponse) Set(val *ExtensionsExtensionListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionsExtensionListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionsExtensionListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionsExtensionListResponse(val *ExtensionsExtensionListResponse) *NullableExtensionsExtensionListResponse {
	return &NullableExtensionsExtensionListResponse{value: val, isSet: true}
}

func (v NullableExtensionsExtensionListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionsExtensionListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
