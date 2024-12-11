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

// checks if the CreateScrapeConfigPayloadHttpSdConfigsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateScrapeConfigPayloadHttpSdConfigsInner{}

// CreateScrapeConfigPayloadHttpSdConfigsInner struct for CreateScrapeConfigPayloadHttpSdConfigsInner
type CreateScrapeConfigPayloadHttpSdConfigsInner struct {
	BasicAuth *CreateScrapeConfigPayloadBasicAuth                `json:"basicAuth,omitempty"`
	Oauth2    *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 `json:"oauth2,omitempty"`
	// Refresh interval to re-query the endpoint. E.g. 60s `Additional Validators:` * must be a valid time format* must be >= 60s
	RefreshInterval *string                                                     `json:"refreshInterval,omitempty"`
	TlsConfig       *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig `json:"tlsConfig,omitempty"`
	// URL from which the targets are fetched.
	// REQUIRED
	Url *string `json:"url"`
}

type _CreateScrapeConfigPayloadHttpSdConfigsInner CreateScrapeConfigPayloadHttpSdConfigsInner

// NewCreateScrapeConfigPayloadHttpSdConfigsInner instantiates a new CreateScrapeConfigPayloadHttpSdConfigsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateScrapeConfigPayloadHttpSdConfigsInner(url *string) *CreateScrapeConfigPayloadHttpSdConfigsInner {
	this := CreateScrapeConfigPayloadHttpSdConfigsInner{}
	var refreshInterval string = "60s"
	this.RefreshInterval = &refreshInterval
	this.Url = url
	return &this
}

// NewCreateScrapeConfigPayloadHttpSdConfigsInnerWithDefaults instantiates a new CreateScrapeConfigPayloadHttpSdConfigsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateScrapeConfigPayloadHttpSdConfigsInnerWithDefaults() *CreateScrapeConfigPayloadHttpSdConfigsInner {
	this := CreateScrapeConfigPayloadHttpSdConfigsInner{}
	var refreshInterval string = "60s"
	this.RefreshInterval = &refreshInterval
	return &this
}

// GetBasicAuth returns the BasicAuth field value if set, zero value otherwise.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) GetBasicAuth() *CreateScrapeConfigPayloadBasicAuth {
	if o == nil || IsNil(o.BasicAuth) {
		var ret *CreateScrapeConfigPayloadBasicAuth
		return ret
	}
	return o.BasicAuth
}

// GetBasicAuthOk returns a tuple with the BasicAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) GetBasicAuthOk() (*CreateScrapeConfigPayloadBasicAuth, bool) {
	if o == nil || IsNil(o.BasicAuth) {
		return nil, false
	}
	return o.BasicAuth, true
}

// HasBasicAuth returns a boolean if a field has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) HasBasicAuth() bool {
	if o != nil && !IsNil(o.BasicAuth) && !IsNil(o.BasicAuth) {
		return true
	}

	return false
}

// SetBasicAuth gets a reference to the given CreateScrapeConfigPayloadBasicAuth and assigns it to the BasicAuth field.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) SetBasicAuth(v *CreateScrapeConfigPayloadBasicAuth) {
	o.BasicAuth = v
}

// GetOauth2 returns the Oauth2 field value if set, zero value otherwise.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) GetOauth2() *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 {
	if o == nil || IsNil(o.Oauth2) {
		var ret *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2
		return ret
	}
	return o.Oauth2
}

// GetOauth2Ok returns a tuple with the Oauth2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) GetOauth2Ok() (*CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2, bool) {
	if o == nil || IsNil(o.Oauth2) {
		return nil, false
	}
	return o.Oauth2, true
}

// HasOauth2 returns a boolean if a field has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) HasOauth2() bool {
	if o != nil && !IsNil(o.Oauth2) && !IsNil(o.Oauth2) {
		return true
	}

	return false
}

// SetOauth2 gets a reference to the given CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 and assigns it to the Oauth2 field.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) SetOauth2(v *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2) {
	o.Oauth2 = v
}

// GetRefreshInterval returns the RefreshInterval field value if set, zero value otherwise.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) GetRefreshInterval() *string {
	if o == nil || IsNil(o.RefreshInterval) {
		var ret *string
		return ret
	}
	return o.RefreshInterval
}

// GetRefreshIntervalOk returns a tuple with the RefreshInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) GetRefreshIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshInterval) {
		return nil, false
	}
	return o.RefreshInterval, true
}

// HasRefreshInterval returns a boolean if a field has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) HasRefreshInterval() bool {
	if o != nil && !IsNil(o.RefreshInterval) && !IsNil(o.RefreshInterval) {
		return true
	}

	return false
}

// SetRefreshInterval gets a reference to the given string and assigns it to the RefreshInterval field.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) SetRefreshInterval(v *string) {
	o.RefreshInterval = v
}

// GetTlsConfig returns the TlsConfig field value if set, zero value otherwise.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) GetTlsConfig() *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig {
	if o == nil || IsNil(o.TlsConfig) {
		var ret *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig
		return ret
	}
	return o.TlsConfig
}

// GetTlsConfigOk returns a tuple with the TlsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) GetTlsConfigOk() (*CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig, bool) {
	if o == nil || IsNil(o.TlsConfig) {
		return nil, false
	}
	return o.TlsConfig, true
}

// HasTlsConfig returns a boolean if a field has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) HasTlsConfig() bool {
	if o != nil && !IsNil(o.TlsConfig) && !IsNil(o.TlsConfig) {
		return true
	}

	return false
}

// SetTlsConfig gets a reference to the given CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig and assigns it to the TlsConfig field.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) SetTlsConfig(v *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig) {
	o.TlsConfig = v
}

// GetUrl returns the Url field value
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) GetUrl() *string {
	if o == nil || IsNil(o.Url) {
		var ret *string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url, true
}

// SetUrl sets field value
func (o *CreateScrapeConfigPayloadHttpSdConfigsInner) SetUrl(v *string) {
	o.Url = v
}

func (o CreateScrapeConfigPayloadHttpSdConfigsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BasicAuth) {
		toSerialize["basicAuth"] = o.BasicAuth
	}
	if !IsNil(o.Oauth2) {
		toSerialize["oauth2"] = o.Oauth2
	}
	if !IsNil(o.RefreshInterval) {
		toSerialize["refreshInterval"] = o.RefreshInterval
	}
	if !IsNil(o.TlsConfig) {
		toSerialize["tlsConfig"] = o.TlsConfig
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableCreateScrapeConfigPayloadHttpSdConfigsInner struct {
	value *CreateScrapeConfigPayloadHttpSdConfigsInner
	isSet bool
}

func (v NullableCreateScrapeConfigPayloadHttpSdConfigsInner) Get() *CreateScrapeConfigPayloadHttpSdConfigsInner {
	return v.value
}

func (v *NullableCreateScrapeConfigPayloadHttpSdConfigsInner) Set(val *CreateScrapeConfigPayloadHttpSdConfigsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateScrapeConfigPayloadHttpSdConfigsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateScrapeConfigPayloadHttpSdConfigsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateScrapeConfigPayloadHttpSdConfigsInner(val *CreateScrapeConfigPayloadHttpSdConfigsInner) *NullableCreateScrapeConfigPayloadHttpSdConfigsInner {
	return &NullableCreateScrapeConfigPayloadHttpSdConfigsInner{value: val, isSet: true}
}

func (v NullableCreateScrapeConfigPayloadHttpSdConfigsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateScrapeConfigPayloadHttpSdConfigsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
