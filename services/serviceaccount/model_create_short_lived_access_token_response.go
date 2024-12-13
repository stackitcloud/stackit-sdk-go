/*
Service Account API

API to manage Service Accounts and their Access Tokens.  ### System for Cross-domain Identity Management (SCIM) Service Account Service offers SCIM APIs to query state. The SCIM protocol was created as standard for  automating the exchange of user identity information between identity domains, or IT systems. Service accounts  are be handled as indentites similar to SCIM users. A custom SCIM schema has been created: `/ServiceAccounts`  #### Syntax ##### Attribute operators | OPERATOR | DESCRIPTION              | |----------|--------------------------| | eq       | equal                    | | ne       | not equal                | | co       | contains                 | | sw       | starts with              | | ew       | ends with                |  ##### Logical operators | OPERATOR | DESCRIPTION              | |----------|--------------------------| | and      | logical \"and\"            | | or       | logical \"or\"             |  ##### Grouping operators | OPERATOR | DESCRIPTION              | |----------|--------------------------| | ()       | precending grouping      |  ##### Example ``` filter=email eq \"my-service-account-aBc2defg@sa.stackit.cloud\" filter=email ne \"my-service-account-aBc2defg@sa.stackit.cloud\" filter=email co \"my-service-account\" filter=name sw \"my\" filter=name ew \"account\" filter=email co \"my-service-account\" and name sw \"my\" filter=email co \"my-service-account\" and (name sw \"my\" or name ew \"account\") ```  #### Sorting  > Sorting is optional  | PARAMETER | DESCRIPTION                          | |-----------|--------------------------------------| | sortBy    | attribute response is ordered by     | | sortOrder | 'ASCENDING' (default) or 'DESCENDING'|  #### Pagination  | PARAMETER    | DESCRIPTION                                  | |--------------|----------------------------------------------| | startIndex   | index of first query result, default: 1      | | count        | maximum number of query results, default: 100|

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceaccount

import (
	"encoding/json"
)

// checks if the CreateShortLivedAccessTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShortLivedAccessTokenResponse{}

// CreateShortLivedAccessTokenResponse struct for CreateShortLivedAccessTokenResponse
type CreateShortLivedAccessTokenResponse struct {
	// The short lived token that can be used for API access
	// REQUIRED
	AccessToken *string `json:"access_token"`
	// REQUIRED
	ExpiresIn *int64 `json:"expires_in"`
	// Refresh token that can be used to request a new access token when it expires (and before refresh token expires). Tokens are rotated.
	// REQUIRED
	RefreshToken *string `json:"refresh_token"`
	// scope field of the self signed token
	// REQUIRED
	Scope *string `json:"scope"`
	// REQUIRED
	TokenType *string `json:"token_type"`
}

type _CreateShortLivedAccessTokenResponse CreateShortLivedAccessTokenResponse

// NewCreateShortLivedAccessTokenResponse instantiates a new CreateShortLivedAccessTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShortLivedAccessTokenResponse(accessToken *string, expiresIn *int64, refreshToken *string, scope *string, tokenType *string) *CreateShortLivedAccessTokenResponse {
	this := CreateShortLivedAccessTokenResponse{}
	this.AccessToken = accessToken
	this.ExpiresIn = expiresIn
	this.RefreshToken = refreshToken
	this.Scope = scope
	this.TokenType = tokenType
	return &this
}

// NewCreateShortLivedAccessTokenResponseWithDefaults instantiates a new CreateShortLivedAccessTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShortLivedAccessTokenResponseWithDefaults() *CreateShortLivedAccessTokenResponse {
	this := CreateShortLivedAccessTokenResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *CreateShortLivedAccessTokenResponse) GetAccessToken() *string {
	if o == nil || IsNil(o.AccessToken) {
		var ret *string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *CreateShortLivedAccessTokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// SetAccessToken sets field value
func (o *CreateShortLivedAccessTokenResponse) SetAccessToken(v *string) {
	o.AccessToken = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *CreateShortLivedAccessTokenResponse) GetExpiresIn() *int64 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret *int64
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *CreateShortLivedAccessTokenResponse) GetExpiresInOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *CreateShortLivedAccessTokenResponse) SetExpiresIn(v *int64) {
	o.ExpiresIn = v
}

// GetRefreshToken returns the RefreshToken field value
func (o *CreateShortLivedAccessTokenResponse) GetRefreshToken() *string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret *string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *CreateShortLivedAccessTokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *CreateShortLivedAccessTokenResponse) SetRefreshToken(v *string) {
	o.RefreshToken = v
}

// GetScope returns the Scope field value
func (o *CreateShortLivedAccessTokenResponse) GetScope() *string {
	if o == nil || IsNil(o.Scope) {
		var ret *string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *CreateShortLivedAccessTokenResponse) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope, true
}

// SetScope sets field value
func (o *CreateShortLivedAccessTokenResponse) SetScope(v *string) {
	o.Scope = v
}

// GetTokenType returns the TokenType field value
func (o *CreateShortLivedAccessTokenResponse) GetTokenType() *string {
	if o == nil || IsNil(o.TokenType) {
		var ret *string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *CreateShortLivedAccessTokenResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenType, true
}

// SetTokenType sets field value
func (o *CreateShortLivedAccessTokenResponse) SetTokenType(v *string) {
	o.TokenType = v
}

func (o CreateShortLivedAccessTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["expires_in"] = o.ExpiresIn
	toSerialize["refresh_token"] = o.RefreshToken
	toSerialize["scope"] = o.Scope
	toSerialize["token_type"] = o.TokenType
	return toSerialize, nil
}

type NullableCreateShortLivedAccessTokenResponse struct {
	value *CreateShortLivedAccessTokenResponse
	isSet bool
}

func (v NullableCreateShortLivedAccessTokenResponse) Get() *CreateShortLivedAccessTokenResponse {
	return v.value
}

func (v *NullableCreateShortLivedAccessTokenResponse) Set(val *CreateShortLivedAccessTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShortLivedAccessTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShortLivedAccessTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShortLivedAccessTokenResponse(val *CreateShortLivedAccessTokenResponse) *NullableCreateShortLivedAccessTokenResponse {
	return &NullableCreateShortLivedAccessTokenResponse{value: val, isSet: true}
}

func (v NullableCreateShortLivedAccessTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShortLivedAccessTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
