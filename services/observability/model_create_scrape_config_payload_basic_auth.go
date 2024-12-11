/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"encoding/json"
)

// checks if the CreateScrapeConfigPayloadBasicAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateScrapeConfigPayloadBasicAuth{}

// CreateScrapeConfigPayloadBasicAuth Sets the 'Authorization' header on every scrape request with the configured username and password. `Additional Validators:` * if basicAuth is in the body no other authentication method should be in the body
type CreateScrapeConfigPayloadBasicAuth struct {
	// password
	Password *string `json:"password,omitempty"`
	// username
	Username *string `json:"username,omitempty"`
}

// NewCreateScrapeConfigPayloadBasicAuth instantiates a new CreateScrapeConfigPayloadBasicAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateScrapeConfigPayloadBasicAuth() *CreateScrapeConfigPayloadBasicAuth {
	this := CreateScrapeConfigPayloadBasicAuth{}
	return &this
}

// NewCreateScrapeConfigPayloadBasicAuthWithDefaults instantiates a new CreateScrapeConfigPayloadBasicAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateScrapeConfigPayloadBasicAuthWithDefaults() *CreateScrapeConfigPayloadBasicAuth {
	this := CreateScrapeConfigPayloadBasicAuth{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateScrapeConfigPayloadBasicAuth) GetPassword() *string {
	if o == nil || IsNil(o.Password) {
		var ret *string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScrapeConfigPayloadBasicAuth) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateScrapeConfigPayloadBasicAuth) HasPassword() bool {
	if o != nil && !IsNil(o.Password) && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateScrapeConfigPayloadBasicAuth) SetPassword(v *string) {
	o.Password = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateScrapeConfigPayloadBasicAuth) GetUsername() *string {
	if o == nil || IsNil(o.Username) {
		var ret *string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScrapeConfigPayloadBasicAuth) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateScrapeConfigPayloadBasicAuth) HasUsername() bool {
	if o != nil && !IsNil(o.Username) && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateScrapeConfigPayloadBasicAuth) SetUsername(v *string) {
	o.Username = v
}

func (o CreateScrapeConfigPayloadBasicAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableCreateScrapeConfigPayloadBasicAuth struct {
	value *CreateScrapeConfigPayloadBasicAuth
	isSet bool
}

func (v NullableCreateScrapeConfigPayloadBasicAuth) Get() *CreateScrapeConfigPayloadBasicAuth {
	return v.value
}

func (v *NullableCreateScrapeConfigPayloadBasicAuth) Set(val *CreateScrapeConfigPayloadBasicAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateScrapeConfigPayloadBasicAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateScrapeConfigPayloadBasicAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateScrapeConfigPayloadBasicAuth(val *CreateScrapeConfigPayloadBasicAuth) *NullableCreateScrapeConfigPayloadBasicAuth {
	return &NullableCreateScrapeConfigPayloadBasicAuth{value: val, isSet: true}
}

func (v NullableCreateScrapeConfigPayloadBasicAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateScrapeConfigPayloadBasicAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
