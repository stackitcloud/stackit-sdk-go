/*
STACKIT LogMe API

The STACKIT LogMe API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logme

import (
	"encoding/json"
)

// checks if the CredentialsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsResponse{}

// CredentialsResponse struct for CredentialsResponse
type CredentialsResponse struct {
	// REQUIRED
	Id  *string         `json:"id"`
	Raw *RawCredentials `json:"raw,omitempty"`
	// REQUIRED
	Uri *string `json:"uri"`
}

type _CredentialsResponse CredentialsResponse

// NewCredentialsResponse instantiates a new CredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsResponse(id *string, uri *string) *CredentialsResponse {
	this := CredentialsResponse{}
	this.Id = id
	this.Uri = uri
	return &this
}

// NewCredentialsResponseWithDefaults instantiates a new CredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsResponseWithDefaults() *CredentialsResponse {
	this := CredentialsResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CredentialsResponse) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CredentialsResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *CredentialsResponse) SetId(v *string) {
	o.Id = v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *CredentialsResponse) GetRaw() *RawCredentials {
	if o == nil || IsNil(o.Raw) {
		var ret *RawCredentials
		return ret
	}
	return o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsResponse) GetRawOk() (*RawCredentials, bool) {
	if o == nil || IsNil(o.Raw) {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *CredentialsResponse) HasRaw() bool {
	if o != nil && !IsNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given RawCredentials and assigns it to the Raw field.
func (o *CredentialsResponse) SetRaw(v *RawCredentials) {
	o.Raw = v
}

// GetUri returns the Uri field value
func (o *CredentialsResponse) GetUri() *string {
	if o == nil || IsNil(o.Uri) {
		var ret *string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *CredentialsResponse) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri, true
}

// SetUri sets field value
func (o *CredentialsResponse) SetUri(v *string) {
	o.Uri = v
}

func (o CredentialsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Raw) {
		toSerialize["raw"] = o.Raw
	}
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

type NullableCredentialsResponse struct {
	value *CredentialsResponse
	isSet bool
}

func (v NullableCredentialsResponse) Get() *CredentialsResponse {
	return v.value
}

func (v *NullableCredentialsResponse) Set(val *CredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsResponse(val *CredentialsResponse) *NullableCredentialsResponse {
	return &NullableCredentialsResponse{value: val, isSet: true}
}

func (v NullableCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
