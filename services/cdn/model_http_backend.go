/*
CDN API

API used to create and manage your CDN distributions.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the HttpBackend type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpBackend{}

/*
	types and functions for originRequestHeaders
*/

// isContainer
type HttpBackendGetOriginRequestHeadersAttributeType = *map[string]string
type HttpBackendGetOriginRequestHeadersArgType = map[string]string
type HttpBackendGetOriginRequestHeadersRetType = map[string]string

func getHttpBackendGetOriginRequestHeadersAttributeTypeOk(arg HttpBackendGetOriginRequestHeadersAttributeType) (ret HttpBackendGetOriginRequestHeadersRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setHttpBackendGetOriginRequestHeadersAttributeType(arg *HttpBackendGetOriginRequestHeadersAttributeType, val HttpBackendGetOriginRequestHeadersRetType) {
	*arg = &val
}

/*
	types and functions for originUrl
*/

// isNotNullableString
type HttpBackendGetOriginUrlAttributeType = *string

func getHttpBackendGetOriginUrlAttributeTypeOk(arg HttpBackendGetOriginUrlAttributeType) (ret HttpBackendGetOriginUrlRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setHttpBackendGetOriginUrlAttributeType(arg *HttpBackendGetOriginUrlAttributeType, val HttpBackendGetOriginUrlRetType) {
	*arg = &val
}

type HttpBackendGetOriginUrlArgType = string
type HttpBackendGetOriginUrlRetType = string

/*
	types and functions for type
*/

// isNotNullableString
type HttpBackendGetTypeAttributeType = *string

func getHttpBackendGetTypeAttributeTypeOk(arg HttpBackendGetTypeAttributeType) (ret HttpBackendGetTypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setHttpBackendGetTypeAttributeType(arg *HttpBackendGetTypeAttributeType, val HttpBackendGetTypeRetType) {
	*arg = &val
}

type HttpBackendGetTypeArgType = string
type HttpBackendGetTypeRetType = string

// HttpBackend struct for HttpBackend
type HttpBackend struct {
	// Headers that will be sent with every request to the configured origin.  **WARNING**: Do not store sensitive values in the headers.  The configuration is stored as plain text.
	// REQUIRED
	OriginRequestHeaders HttpBackendGetOriginRequestHeadersAttributeType `json:"originRequestHeaders"`
	// REQUIRED
	OriginUrl HttpBackendGetOriginUrlAttributeType `json:"originUrl"`
	// REQUIRED
	Type HttpBackendGetTypeAttributeType `json:"type"`
}

type _HttpBackend HttpBackend

// NewHttpBackend instantiates a new HttpBackend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpBackend(originRequestHeaders HttpBackendGetOriginRequestHeadersArgType, originUrl HttpBackendGetOriginUrlArgType, type_ HttpBackendGetTypeArgType) *HttpBackend {
	this := HttpBackend{}
	setHttpBackendGetOriginRequestHeadersAttributeType(&this.OriginRequestHeaders, originRequestHeaders)
	setHttpBackendGetOriginUrlAttributeType(&this.OriginUrl, originUrl)
	setHttpBackendGetTypeAttributeType(&this.Type, type_)
	return &this
}

// NewHttpBackendWithDefaults instantiates a new HttpBackend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpBackendWithDefaults() *HttpBackend {
	this := HttpBackend{}
	return &this
}

// GetOriginRequestHeaders returns the OriginRequestHeaders field value
func (o *HttpBackend) GetOriginRequestHeaders() (ret HttpBackendGetOriginRequestHeadersRetType) {
	ret, _ = o.GetOriginRequestHeadersOk()
	return ret
}

// GetOriginRequestHeadersOk returns a tuple with the OriginRequestHeaders field value
// and a boolean to check if the value has been set.
func (o *HttpBackend) GetOriginRequestHeadersOk() (ret HttpBackendGetOriginRequestHeadersRetType, ok bool) {
	return getHttpBackendGetOriginRequestHeadersAttributeTypeOk(o.OriginRequestHeaders)
}

// SetOriginRequestHeaders sets field value
func (o *HttpBackend) SetOriginRequestHeaders(v HttpBackendGetOriginRequestHeadersRetType) {
	setHttpBackendGetOriginRequestHeadersAttributeType(&o.OriginRequestHeaders, v)
}

// GetOriginUrl returns the OriginUrl field value
func (o *HttpBackend) GetOriginUrl() (ret HttpBackendGetOriginUrlRetType) {
	ret, _ = o.GetOriginUrlOk()
	return ret
}

// GetOriginUrlOk returns a tuple with the OriginUrl field value
// and a boolean to check if the value has been set.
func (o *HttpBackend) GetOriginUrlOk() (ret HttpBackendGetOriginUrlRetType, ok bool) {
	return getHttpBackendGetOriginUrlAttributeTypeOk(o.OriginUrl)
}

// SetOriginUrl sets field value
func (o *HttpBackend) SetOriginUrl(v HttpBackendGetOriginUrlRetType) {
	setHttpBackendGetOriginUrlAttributeType(&o.OriginUrl, v)
}

// GetType returns the Type field value
func (o *HttpBackend) GetType() (ret HttpBackendGetTypeRetType) {
	ret, _ = o.GetTypeOk()
	return ret
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HttpBackend) GetTypeOk() (ret HttpBackendGetTypeRetType, ok bool) {
	return getHttpBackendGetTypeAttributeTypeOk(o.Type)
}

// SetType sets field value
func (o *HttpBackend) SetType(v HttpBackendGetTypeRetType) {
	setHttpBackendGetTypeAttributeType(&o.Type, v)
}

func (o HttpBackend) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getHttpBackendGetOriginRequestHeadersAttributeTypeOk(o.OriginRequestHeaders); ok {
		toSerialize["OriginRequestHeaders"] = val
	}
	if val, ok := getHttpBackendGetOriginUrlAttributeTypeOk(o.OriginUrl); ok {
		toSerialize["OriginUrl"] = val
	}
	if val, ok := getHttpBackendGetTypeAttributeTypeOk(o.Type); ok {
		toSerialize["Type"] = val
	}
	return toSerialize, nil
}

type NullableHttpBackend struct {
	value *HttpBackend
	isSet bool
}

func (v NullableHttpBackend) Get() *HttpBackend {
	return v.value
}

func (v *NullableHttpBackend) Set(val *HttpBackend) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpBackend) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpBackend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpBackend(val *HttpBackend) *NullableHttpBackend {
	return &NullableHttpBackend{value: val, isSet: true}
}

func (v NullableHttpBackend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpBackend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
