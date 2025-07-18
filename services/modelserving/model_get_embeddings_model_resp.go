/*
STACKIT Model Serving API

This API provides endpoints for the model serving api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package modelserving

import (
	"encoding/json"
)

// checks if the GetEmbeddingsModelResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEmbeddingsModelResp{}

/*
	types and functions for message
*/

// isNotNullableString
type GetEmbeddingsModelRespGetMessageAttributeType = *string

func getGetEmbeddingsModelRespGetMessageAttributeTypeOk(arg GetEmbeddingsModelRespGetMessageAttributeType) (ret GetEmbeddingsModelRespGetMessageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetEmbeddingsModelRespGetMessageAttributeType(arg *GetEmbeddingsModelRespGetMessageAttributeType, val GetEmbeddingsModelRespGetMessageRetType) {
	*arg = &val
}

type GetEmbeddingsModelRespGetMessageArgType = string
type GetEmbeddingsModelRespGetMessageRetType = string

/*
	types and functions for model
*/

// isModel
type GetEmbeddingsModelRespGetModelAttributeType = *EmbeddingModelDetails
type GetEmbeddingsModelRespGetModelArgType = EmbeddingModelDetails
type GetEmbeddingsModelRespGetModelRetType = EmbeddingModelDetails

func getGetEmbeddingsModelRespGetModelAttributeTypeOk(arg GetEmbeddingsModelRespGetModelAttributeType) (ret GetEmbeddingsModelRespGetModelRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetEmbeddingsModelRespGetModelAttributeType(arg *GetEmbeddingsModelRespGetModelAttributeType, val GetEmbeddingsModelRespGetModelRetType) {
	*arg = &val
}

// GetEmbeddingsModelResp struct for GetEmbeddingsModelResp
type GetEmbeddingsModelResp struct {
	Message GetEmbeddingsModelRespGetMessageAttributeType `json:"message,omitempty"`
	// REQUIRED
	Model GetEmbeddingsModelRespGetModelAttributeType `json:"model" required:"true"`
}

type _GetEmbeddingsModelResp GetEmbeddingsModelResp

// NewGetEmbeddingsModelResp instantiates a new GetEmbeddingsModelResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEmbeddingsModelResp(model GetEmbeddingsModelRespGetModelArgType) *GetEmbeddingsModelResp {
	this := GetEmbeddingsModelResp{}
	setGetEmbeddingsModelRespGetModelAttributeType(&this.Model, model)
	return &this
}

// NewGetEmbeddingsModelRespWithDefaults instantiates a new GetEmbeddingsModelResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEmbeddingsModelRespWithDefaults() *GetEmbeddingsModelResp {
	this := GetEmbeddingsModelResp{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetEmbeddingsModelResp) GetMessage() (res GetEmbeddingsModelRespGetMessageRetType) {
	res, _ = o.GetMessageOk()
	return
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEmbeddingsModelResp) GetMessageOk() (ret GetEmbeddingsModelRespGetMessageRetType, ok bool) {
	return getGetEmbeddingsModelRespGetMessageAttributeTypeOk(o.Message)
}

// HasMessage returns a boolean if a field has been set.
func (o *GetEmbeddingsModelResp) HasMessage() bool {
	_, ok := o.GetMessageOk()
	return ok
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetEmbeddingsModelResp) SetMessage(v GetEmbeddingsModelRespGetMessageRetType) {
	setGetEmbeddingsModelRespGetMessageAttributeType(&o.Message, v)
}

// GetModel returns the Model field value
func (o *GetEmbeddingsModelResp) GetModel() (ret GetEmbeddingsModelRespGetModelRetType) {
	ret, _ = o.GetModelOk()
	return ret
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *GetEmbeddingsModelResp) GetModelOk() (ret GetEmbeddingsModelRespGetModelRetType, ok bool) {
	return getGetEmbeddingsModelRespGetModelAttributeTypeOk(o.Model)
}

// SetModel sets field value
func (o *GetEmbeddingsModelResp) SetModel(v GetEmbeddingsModelRespGetModelRetType) {
	setGetEmbeddingsModelRespGetModelAttributeType(&o.Model, v)
}

func (o GetEmbeddingsModelResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getGetEmbeddingsModelRespGetMessageAttributeTypeOk(o.Message); ok {
		toSerialize["Message"] = val
	}
	if val, ok := getGetEmbeddingsModelRespGetModelAttributeTypeOk(o.Model); ok {
		toSerialize["Model"] = val
	}
	return toSerialize, nil
}

type NullableGetEmbeddingsModelResp struct {
	value *GetEmbeddingsModelResp
	isSet bool
}

func (v NullableGetEmbeddingsModelResp) Get() *GetEmbeddingsModelResp {
	return v.value
}

func (v *NullableGetEmbeddingsModelResp) Set(val *GetEmbeddingsModelResp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEmbeddingsModelResp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEmbeddingsModelResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEmbeddingsModelResp(val *GetEmbeddingsModelResp) *NullableGetEmbeddingsModelResp {
	return &NullableGetEmbeddingsModelResp{value: val, isSet: true}
}

func (v NullableGetEmbeddingsModelResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEmbeddingsModelResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
