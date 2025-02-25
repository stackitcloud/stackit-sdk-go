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

// checks if the HttpBackendPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpBackendPatch{}

// HttpBackendPatch A partial HTTP Backend
type HttpBackendPatch struct {
	// Headers that will be sent with every request to the configured origin.  **WARNING**: Do not store sensitive values in the headers.  The configuration is stored as plain text.
	OriginRequestHeaders *map[string]string `json:"originRequestHeaders,omitempty"`
	OriginUrl            *string            `json:"originUrl,omitempty"`
	// This property is required to determine the used backend type.
	// REQUIRED
	Type *string `json:"type"`
}

type _HttpBackendPatch HttpBackendPatch

// NewHttpBackendPatch instantiates a new HttpBackendPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpBackendPatch(type_ *string) *HttpBackendPatch {
	this := HttpBackendPatch{}
	this.Type = type_
	return &this
}

// NewHttpBackendPatchWithDefaults instantiates a new HttpBackendPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpBackendPatchWithDefaults() *HttpBackendPatch {
	this := HttpBackendPatch{}
	return &this
}

// GetOriginRequestHeaders returns the OriginRequestHeaders field value if set, zero value otherwise.
func (o *HttpBackendPatch) GetOriginRequestHeaders() *map[string]string {
	if o == nil || IsNil(o.OriginRequestHeaders) {
		var ret *map[string]string
		return ret
	}
	return o.OriginRequestHeaders
}

// GetOriginRequestHeadersOk returns a tuple with the OriginRequestHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpBackendPatch) GetOriginRequestHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.OriginRequestHeaders) {
		return nil, false
	}
	return o.OriginRequestHeaders, true
}

// HasOriginRequestHeaders returns a boolean if a field has been set.
func (o *HttpBackendPatch) HasOriginRequestHeaders() bool {
	if o != nil && !IsNil(o.OriginRequestHeaders) {
		return true
	}

	return false
}

// SetOriginRequestHeaders gets a reference to the given map[string]string and assigns it to the OriginRequestHeaders field.
func (o *HttpBackendPatch) SetOriginRequestHeaders(v *map[string]string) {
	o.OriginRequestHeaders = v
}

// GetOriginUrl returns the OriginUrl field value if set, zero value otherwise.
func (o *HttpBackendPatch) GetOriginUrl() *string {
	if o == nil || IsNil(o.OriginUrl) {
		var ret *string
		return ret
	}
	return o.OriginUrl
}

// GetOriginUrlOk returns a tuple with the OriginUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpBackendPatch) GetOriginUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OriginUrl) {
		return nil, false
	}
	return o.OriginUrl, true
}

// HasOriginUrl returns a boolean if a field has been set.
func (o *HttpBackendPatch) HasOriginUrl() bool {
	if o != nil && !IsNil(o.OriginUrl) {
		return true
	}

	return false
}

// SetOriginUrl gets a reference to the given string and assigns it to the OriginUrl field.
func (o *HttpBackendPatch) SetOriginUrl(v *string) {
	o.OriginUrl = v
}

// GetType returns the Type field value
func (o *HttpBackendPatch) GetType() *string {
	if o == nil || IsNil(o.Type) {
		var ret *string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HttpBackendPatch) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *HttpBackendPatch) SetType(v *string) {
	o.Type = v
}

func (o HttpBackendPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OriginRequestHeaders) {
		toSerialize["originRequestHeaders"] = o.OriginRequestHeaders
	}
	if !IsNil(o.OriginUrl) {
		toSerialize["originUrl"] = o.OriginUrl
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableHttpBackendPatch struct {
	value *HttpBackendPatch
	isSet bool
}

func (v NullableHttpBackendPatch) Get() *HttpBackendPatch {
	return v.value
}

func (v *NullableHttpBackendPatch) Set(val *HttpBackendPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpBackendPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpBackendPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpBackendPatch(val *HttpBackendPatch) *NullableHttpBackendPatch {
	return &NullableHttpBackendPatch{value: val, isSet: true}
}

func (v NullableHttpBackendPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpBackendPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
