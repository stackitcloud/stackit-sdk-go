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

// checks if the GetTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTokenResponse{}

/*
	types and functions for message
*/

// isNotNullableString
type GetTokenResponseGetMessageAttributeType = *string

func getGetTokenResponseGetMessageAttributeTypeOk(arg GetTokenResponseGetMessageAttributeType) (ret GetTokenResponseGetMessageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetTokenResponseGetMessageAttributeType(arg *GetTokenResponseGetMessageAttributeType, val GetTokenResponseGetMessageRetType) {
	*arg = &val
}

type GetTokenResponseGetMessageArgType = string
type GetTokenResponseGetMessageRetType = string

/*
	types and functions for token
*/

// isModel
type GetTokenResponseGetTokenAttributeType = *Token
type GetTokenResponseGetTokenArgType = Token
type GetTokenResponseGetTokenRetType = Token

func getGetTokenResponseGetTokenAttributeTypeOk(arg GetTokenResponseGetTokenAttributeType) (ret GetTokenResponseGetTokenRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetTokenResponseGetTokenAttributeType(arg *GetTokenResponseGetTokenAttributeType, val GetTokenResponseGetTokenRetType) {
	*arg = &val
}

// GetTokenResponse struct for GetTokenResponse
type GetTokenResponse struct {
	Message GetTokenResponseGetMessageAttributeType `json:"message,omitempty"`
	// REQUIRED
	Token GetTokenResponseGetTokenAttributeType `json:"token"`
}

type _GetTokenResponse GetTokenResponse

// NewGetTokenResponse instantiates a new GetTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTokenResponse(token GetTokenResponseGetTokenArgType) *GetTokenResponse {
	this := GetTokenResponse{}
	setGetTokenResponseGetTokenAttributeType(&this.Token, token)
	return &this
}

// NewGetTokenResponseWithDefaults instantiates a new GetTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTokenResponseWithDefaults() *GetTokenResponse {
	this := GetTokenResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetTokenResponse) GetMessage() (res GetTokenResponseGetMessageRetType) {
	res, _ = o.GetMessageOk()
	return
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTokenResponse) GetMessageOk() (ret GetTokenResponseGetMessageRetType, ok bool) {
	return getGetTokenResponseGetMessageAttributeTypeOk(o.Message)
}

// HasMessage returns a boolean if a field has been set.
func (o *GetTokenResponse) HasMessage() bool {
	_, ok := o.GetMessageOk()
	return ok
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetTokenResponse) SetMessage(v GetTokenResponseGetMessageRetType) {
	setGetTokenResponseGetMessageAttributeType(&o.Message, v)
}

// GetToken returns the Token field value
func (o *GetTokenResponse) GetToken() (ret GetTokenResponseGetTokenRetType) {
	ret, _ = o.GetTokenOk()
	return ret
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *GetTokenResponse) GetTokenOk() (ret GetTokenResponseGetTokenRetType, ok bool) {
	return getGetTokenResponseGetTokenAttributeTypeOk(o.Token)
}

// SetToken sets field value
func (o *GetTokenResponse) SetToken(v GetTokenResponseGetTokenRetType) {
	setGetTokenResponseGetTokenAttributeType(&o.Token, v)
}

func (o GetTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getGetTokenResponseGetMessageAttributeTypeOk(o.Message); ok {
		toSerialize["Message"] = val
	}
	if val, ok := getGetTokenResponseGetTokenAttributeTypeOk(o.Token); ok {
		toSerialize["Token"] = val
	}
	return toSerialize, nil
}

type NullableGetTokenResponse struct {
	value *GetTokenResponse
	isSet bool
}

func (v NullableGetTokenResponse) Get() *GetTokenResponse {
	return v.value
}

func (v *NullableGetTokenResponse) Set(val *GetTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTokenResponse(val *GetTokenResponse) *NullableGetTokenResponse {
	return &NullableGetTokenResponse{value: val, isSet: true}
}

func (v NullableGetTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
