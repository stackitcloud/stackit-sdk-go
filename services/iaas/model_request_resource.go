/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the RequestResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestResource{}

/*
	types and functions for id
*/

// isNotNullableString
type RequestResourceGetIdAttributeType = *string

func getRequestResourceGetIdAttributeTypeOk(arg RequestResourceGetIdAttributeType) (ret RequestResourceGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRequestResourceGetIdAttributeType(arg *RequestResourceGetIdAttributeType, val RequestResourceGetIdRetType) {
	*arg = &val
}

type RequestResourceGetIdArgType = string
type RequestResourceGetIdRetType = string

/*
	types and functions for status
*/

// isNotNullableString
type RequestResourceGetStatusAttributeType = *string

func getRequestResourceGetStatusAttributeTypeOk(arg RequestResourceGetStatusAttributeType) (ret RequestResourceGetStatusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRequestResourceGetStatusAttributeType(arg *RequestResourceGetStatusAttributeType, val RequestResourceGetStatusRetType) {
	*arg = &val
}

type RequestResourceGetStatusArgType = string
type RequestResourceGetStatusRetType = string

/*
	types and functions for type
*/

// isNotNullableString
type RequestResourceGetTypeAttributeType = *string

func getRequestResourceGetTypeAttributeTypeOk(arg RequestResourceGetTypeAttributeType) (ret RequestResourceGetTypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRequestResourceGetTypeAttributeType(arg *RequestResourceGetTypeAttributeType, val RequestResourceGetTypeRetType) {
	*arg = &val
}

type RequestResourceGetTypeArgType = string
type RequestResourceGetTypeRetType = string

// RequestResource Object that represents a resource as part of a request.
type RequestResource struct {
	// Universally Unique Identifier (UUID).
	// REQUIRED
	Id RequestResourceGetIdAttributeType `json:"id" required:"true"`
	// The state of a resource object. Possible values: `CREATING`, `CREATED`, `DELETING`, `DELETED`, `FAILED`, `UPDATED`, `UPDATING`.
	// REQUIRED
	Status RequestResourceGetStatusAttributeType `json:"status" required:"true"`
	// Object that represents a resource type. Possible values: `BACKUP`, `IMAGE`, `NETWORK`, `NETWORKAREA`, `NIC`, `PROJECT`, `ROUTE`, `SERVER`, `SERVICEACCOUNT`, `SNAPSHOT`, `VIRTUALIP`, `VOLUME`.
	// REQUIRED
	Type RequestResourceGetTypeAttributeType `json:"type" required:"true"`
}

type _RequestResource RequestResource

// NewRequestResource instantiates a new RequestResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestResource(id RequestResourceGetIdArgType, status RequestResourceGetStatusArgType, types RequestResourceGetTypeArgType) *RequestResource {
	this := RequestResource{}
	setRequestResourceGetIdAttributeType(&this.Id, id)
	setRequestResourceGetStatusAttributeType(&this.Status, status)
	setRequestResourceGetTypeAttributeType(&this.Type, types)
	return &this
}

// NewRequestResourceWithDefaults instantiates a new RequestResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestResourceWithDefaults() *RequestResource {
	this := RequestResource{}
	return &this
}

// GetId returns the Id field value
func (o *RequestResource) GetId() (ret RequestResourceGetIdRetType) {
	ret, _ = o.GetIdOk()
	return ret
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestResource) GetIdOk() (ret RequestResourceGetIdRetType, ok bool) {
	return getRequestResourceGetIdAttributeTypeOk(o.Id)
}

// SetId sets field value
func (o *RequestResource) SetId(v RequestResourceGetIdRetType) {
	setRequestResourceGetIdAttributeType(&o.Id, v)
}

// GetStatus returns the Status field value
func (o *RequestResource) GetStatus() (ret RequestResourceGetStatusRetType) {
	ret, _ = o.GetStatusOk()
	return ret
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestResource) GetStatusOk() (ret RequestResourceGetStatusRetType, ok bool) {
	return getRequestResourceGetStatusAttributeTypeOk(o.Status)
}

// SetStatus sets field value
func (o *RequestResource) SetStatus(v RequestResourceGetStatusRetType) {
	setRequestResourceGetStatusAttributeType(&o.Status, v)
}

// GetType returns the Type field value
func (o *RequestResource) GetType() (ret RequestResourceGetTypeRetType) {
	ret, _ = o.GetTypeOk()
	return ret
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestResource) GetTypeOk() (ret RequestResourceGetTypeRetType, ok bool) {
	return getRequestResourceGetTypeAttributeTypeOk(o.Type)
}

// SetType sets field value
func (o *RequestResource) SetType(v RequestResourceGetTypeRetType) {
	setRequestResourceGetTypeAttributeType(&o.Type, v)
}

func (o RequestResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getRequestResourceGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getRequestResourceGetStatusAttributeTypeOk(o.Status); ok {
		toSerialize["Status"] = val
	}
	if val, ok := getRequestResourceGetTypeAttributeTypeOk(o.Type); ok {
		toSerialize["Type"] = val
	}
	return toSerialize, nil
}

type NullableRequestResource struct {
	value *RequestResource
	isSet bool
}

func (v NullableRequestResource) Get() *RequestResource {
	return v.value
}

func (v *NullableRequestResource) Set(val *RequestResource) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestResource) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestResource(val *RequestResource) *NullableRequestResource {
	return &NullableRequestResource{value: val, isSet: true}
}

func (v NullableRequestResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
