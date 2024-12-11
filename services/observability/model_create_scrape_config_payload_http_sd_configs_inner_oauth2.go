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

// checks if the CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2{}

// CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 OAuth 2.0 authentication using the client credentials grant type. Prometheus fetches an access token from the specified endpoint with the given client access and secret keys. `Additional Validators:` * if oauth2 is in the body no other authentication method should be in the body
type CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 struct {
	// clientId
	// REQUIRED
	ClientId *string `json:"clientId"`
	// clientSecret
	// REQUIRED
	ClientSecret *string `json:"clientSecret"`
	// The URL to fetch the token from.
	Scopes    *[]string                                                   `json:"scopes,omitempty"`
	TlsConfig *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig `json:"tlsConfig,omitempty"`
	// The URL to fetch the token from.
	// REQUIRED
	TokenUrl *string `json:"tokenUrl"`
}

type _CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2

// NewCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 instantiates a new CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2(clientId *string, clientSecret *string, tokenUrl *string) *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 {
	this := CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.TokenUrl = tokenUrl
	return &this
}

// NewCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2WithDefaults instantiates a new CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2WithDefaults() *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 {
	this := CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) GetClientId() *string {
	if o == nil || IsNil(o.ClientId) {
		var ret *string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId, true
}

// SetClientId sets field value
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) SetClientId(v *string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) GetClientSecret() *string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret *string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) SetClientSecret(v *string) {
	o.ClientSecret = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) GetScopes() *[]string {
	if o == nil || IsNil(o.Scopes) {
		var ret *[]string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) GetScopesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) SetScopes(v *[]string) {
	o.Scopes = v
}

// GetTlsConfig returns the TlsConfig field value if set, zero value otherwise.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) GetTlsConfig() *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig {
	if o == nil || IsNil(o.TlsConfig) {
		var ret *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig
		return ret
	}
	return o.TlsConfig
}

// GetTlsConfigOk returns a tuple with the TlsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) GetTlsConfigOk() (*CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig, bool) {
	if o == nil || IsNil(o.TlsConfig) {
		return nil, false
	}
	return o.TlsConfig, true
}

// HasTlsConfig returns a boolean if a field has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) HasTlsConfig() bool {
	if o != nil && !IsNil(o.TlsConfig) && !IsNil(o.TlsConfig) {
		return true
	}

	return false
}

// SetTlsConfig gets a reference to the given CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig and assigns it to the TlsConfig field.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) SetTlsConfig(v *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig) {
	o.TlsConfig = v
}

// GetTokenUrl returns the TokenUrl field value
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) GetTokenUrl() *string {
	if o == nil || IsNil(o.TokenUrl) {
		var ret *string
		return ret
	}

	return o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value
// and a boolean to check if the value has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) GetTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// SetTokenUrl sets field value
func (o *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) SetTokenUrl(v *string) {
	o.TokenUrl = v
}

func (o CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.TlsConfig) {
		toSerialize["tlsConfig"] = o.TlsConfig
	}
	toSerialize["tokenUrl"] = o.TokenUrl
	return toSerialize, nil
}

type NullableCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 struct {
	value *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2
	isSet bool
}

func (v NullableCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) Get() *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 {
	return v.value
}

func (v *NullableCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) Set(val *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2(val *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) *NullableCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 {
	return &NullableCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2{value: val, isSet: true}
}

func (v NullableCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
