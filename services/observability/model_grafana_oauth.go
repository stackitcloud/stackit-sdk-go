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

// checks if the GrafanaOauth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrafanaOauth{}

// GrafanaOauth struct for GrafanaOauth
type GrafanaOauth struct {
	// REQUIRED
	ApiUrl *string `json:"apiUrl"`
	// REQUIRED
	AuthUrl *string `json:"authUrl"`
	// REQUIRED
	Enabled *bool   `json:"enabled"`
	Name    *string `json:"name,omitempty"`
	// REQUIRED
	OauthClientId *string `json:"oauthClientId"`
	// REQUIRED
	OauthClientSecret *string `json:"oauthClientSecret"`
	// REQUIRED
	RoleAttributePath   *string `json:"roleAttributePath"`
	RoleAttributeStrict *bool   `json:"roleAttributeStrict,omitempty"`
	Scopes              *string `json:"scopes,omitempty"`
	// REQUIRED
	TokenUrl *string `json:"tokenUrl"`
	UsePkce  *bool   `json:"usePkce,omitempty"`
}

type _GrafanaOauth GrafanaOauth

// NewGrafanaOauth instantiates a new GrafanaOauth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrafanaOauth(apiUrl *string, authUrl *string, enabled *bool, oauthClientId *string, oauthClientSecret *string, roleAttributePath *string, tokenUrl *string) *GrafanaOauth {
	this := GrafanaOauth{}
	this.ApiUrl = apiUrl
	this.AuthUrl = authUrl
	this.Enabled = enabled
	this.OauthClientId = oauthClientId
	this.OauthClientSecret = oauthClientSecret
	this.RoleAttributePath = roleAttributePath
	var roleAttributeStrict bool = true
	this.RoleAttributeStrict = &roleAttributeStrict
	var scopes string = "openid profile email"
	this.Scopes = &scopes
	this.TokenUrl = tokenUrl
	return &this
}

// NewGrafanaOauthWithDefaults instantiates a new GrafanaOauth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrafanaOauthWithDefaults() *GrafanaOauth {
	this := GrafanaOauth{}
	var roleAttributeStrict bool = true
	this.RoleAttributeStrict = &roleAttributeStrict
	var scopes string = "openid profile email"
	this.Scopes = &scopes
	return &this
}

// GetApiUrl returns the ApiUrl field value
func (o *GrafanaOauth) GetApiUrl() *string {
	if o == nil || IsNil(o.ApiUrl) {
		var ret *string
		return ret
	}

	return o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value
// and a boolean to check if the value has been set.
func (o *GrafanaOauth) GetApiUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiUrl, true
}

// SetApiUrl sets field value
func (o *GrafanaOauth) SetApiUrl(v *string) {
	o.ApiUrl = v
}

// GetAuthUrl returns the AuthUrl field value
func (o *GrafanaOauth) GetAuthUrl() *string {
	if o == nil || IsNil(o.AuthUrl) {
		var ret *string
		return ret
	}

	return o.AuthUrl
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value
// and a boolean to check if the value has been set.
func (o *GrafanaOauth) GetAuthUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthUrl, true
}

// SetAuthUrl sets field value
func (o *GrafanaOauth) SetAuthUrl(v *string) {
	o.AuthUrl = v
}

// GetEnabled returns the Enabled field value
func (o *GrafanaOauth) GetEnabled() *bool {
	if o == nil || IsNil(o.Enabled) {
		var ret *bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GrafanaOauth) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled, true
}

// SetEnabled sets field value
func (o *GrafanaOauth) SetEnabled(v *bool) {
	o.Enabled = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GrafanaOauth) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaOauth) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GrafanaOauth) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GrafanaOauth) SetName(v *string) {
	o.Name = v
}

// GetOauthClientId returns the OauthClientId field value
func (o *GrafanaOauth) GetOauthClientId() *string {
	if o == nil || IsNil(o.OauthClientId) {
		var ret *string
		return ret
	}

	return o.OauthClientId
}

// GetOauthClientIdOk returns a tuple with the OauthClientId field value
// and a boolean to check if the value has been set.
func (o *GrafanaOauth) GetOauthClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OauthClientId, true
}

// SetOauthClientId sets field value
func (o *GrafanaOauth) SetOauthClientId(v *string) {
	o.OauthClientId = v
}

// GetOauthClientSecret returns the OauthClientSecret field value
func (o *GrafanaOauth) GetOauthClientSecret() *string {
	if o == nil || IsNil(o.OauthClientSecret) {
		var ret *string
		return ret
	}

	return o.OauthClientSecret
}

// GetOauthClientSecretOk returns a tuple with the OauthClientSecret field value
// and a boolean to check if the value has been set.
func (o *GrafanaOauth) GetOauthClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OauthClientSecret, true
}

// SetOauthClientSecret sets field value
func (o *GrafanaOauth) SetOauthClientSecret(v *string) {
	o.OauthClientSecret = v
}

// GetRoleAttributePath returns the RoleAttributePath field value
func (o *GrafanaOauth) GetRoleAttributePath() *string {
	if o == nil || IsNil(o.RoleAttributePath) {
		var ret *string
		return ret
	}

	return o.RoleAttributePath
}

// GetRoleAttributePathOk returns a tuple with the RoleAttributePath field value
// and a boolean to check if the value has been set.
func (o *GrafanaOauth) GetRoleAttributePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleAttributePath, true
}

// SetRoleAttributePath sets field value
func (o *GrafanaOauth) SetRoleAttributePath(v *string) {
	o.RoleAttributePath = v
}

// GetRoleAttributeStrict returns the RoleAttributeStrict field value if set, zero value otherwise.
func (o *GrafanaOauth) GetRoleAttributeStrict() *bool {
	if o == nil || IsNil(o.RoleAttributeStrict) {
		var ret *bool
		return ret
	}
	return o.RoleAttributeStrict
}

// GetRoleAttributeStrictOk returns a tuple with the RoleAttributeStrict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaOauth) GetRoleAttributeStrictOk() (*bool, bool) {
	if o == nil || IsNil(o.RoleAttributeStrict) {
		return nil, false
	}
	return o.RoleAttributeStrict, true
}

// HasRoleAttributeStrict returns a boolean if a field has been set.
func (o *GrafanaOauth) HasRoleAttributeStrict() bool {
	if o != nil && !IsNil(o.RoleAttributeStrict) {
		return true
	}

	return false
}

// SetRoleAttributeStrict gets a reference to the given bool and assigns it to the RoleAttributeStrict field.
func (o *GrafanaOauth) SetRoleAttributeStrict(v *bool) {
	o.RoleAttributeStrict = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *GrafanaOauth) GetScopes() *string {
	if o == nil || IsNil(o.Scopes) {
		var ret *string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaOauth) GetScopesOk() (*string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *GrafanaOauth) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given string and assigns it to the Scopes field.
func (o *GrafanaOauth) SetScopes(v *string) {
	o.Scopes = v
}

// GetTokenUrl returns the TokenUrl field value
func (o *GrafanaOauth) GetTokenUrl() *string {
	if o == nil || IsNil(o.TokenUrl) {
		var ret *string
		return ret
	}

	return o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value
// and a boolean to check if the value has been set.
func (o *GrafanaOauth) GetTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// SetTokenUrl sets field value
func (o *GrafanaOauth) SetTokenUrl(v *string) {
	o.TokenUrl = v
}

// GetUsePkce returns the UsePkce field value if set, zero value otherwise.
func (o *GrafanaOauth) GetUsePkce() *bool {
	if o == nil || IsNil(o.UsePkce) {
		var ret *bool
		return ret
	}
	return o.UsePkce
}

// GetUsePkceOk returns a tuple with the UsePkce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaOauth) GetUsePkceOk() (*bool, bool) {
	if o == nil || IsNil(o.UsePkce) {
		return nil, false
	}
	return o.UsePkce, true
}

// HasUsePkce returns a boolean if a field has been set.
func (o *GrafanaOauth) HasUsePkce() bool {
	if o != nil && !IsNil(o.UsePkce) {
		return true
	}

	return false
}

// SetUsePkce gets a reference to the given bool and assigns it to the UsePkce field.
func (o *GrafanaOauth) SetUsePkce(v *bool) {
	o.UsePkce = v
}

func (o GrafanaOauth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiUrl"] = o.ApiUrl
	toSerialize["authUrl"] = o.AuthUrl
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["oauthClientId"] = o.OauthClientId
	toSerialize["oauthClientSecret"] = o.OauthClientSecret
	toSerialize["roleAttributePath"] = o.RoleAttributePath
	if !IsNil(o.RoleAttributeStrict) {
		toSerialize["roleAttributeStrict"] = o.RoleAttributeStrict
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	toSerialize["tokenUrl"] = o.TokenUrl
	if !IsNil(o.UsePkce) {
		toSerialize["usePkce"] = o.UsePkce
	}
	return toSerialize, nil
}

type NullableGrafanaOauth struct {
	value *GrafanaOauth
	isSet bool
}

func (v NullableGrafanaOauth) Get() *GrafanaOauth {
	return v.value
}

func (v *NullableGrafanaOauth) Set(val *GrafanaOauth) {
	v.value = val
	v.isSet = true
}

func (v NullableGrafanaOauth) IsSet() bool {
	return v.isSet
}

func (v *NullableGrafanaOauth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrafanaOauth(val *GrafanaOauth) *NullableGrafanaOauth {
	return &NullableGrafanaOauth{value: val, isSet: true}
}

func (v NullableGrafanaOauth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrafanaOauth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
