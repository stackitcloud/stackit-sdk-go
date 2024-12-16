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

// checks if the AuthError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthError{}

// AuthError struct for AuthError
type AuthError struct {
	// REQUIRED
	Error *AuthErrorError `json:"error"`
}

type _AuthError AuthError

// NewAuthError instantiates a new AuthError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthError(error_ *AuthErrorError) *AuthError {
	this := AuthError{}
	this.Error = error_
	return &this
}

// NewAuthErrorWithDefaults instantiates a new AuthError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthErrorWithDefaults() *AuthError {
	this := AuthError{}
	return &this
}

// GetError returns the Error field value
func (o *AuthError) GetError() *AuthErrorError {
	if o == nil || IsNil(o.Error) {
		var ret *AuthErrorError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AuthError) GetErrorOk() (*AuthErrorError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error, true
}

// SetError sets field value
func (o *AuthError) SetError(v *AuthErrorError) {
	o.Error = v
}

func (o AuthError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

type NullableAuthError struct {
	value *AuthError
	isSet bool
}

func (v NullableAuthError) Get() *AuthError {
	return v.value
}

func (v *NullableAuthError) Set(val *AuthError) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthError) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthError(val *AuthError) *NullableAuthError {
	return &NullableAuthError{value: val, isSet: true}
}

func (v NullableAuthError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
