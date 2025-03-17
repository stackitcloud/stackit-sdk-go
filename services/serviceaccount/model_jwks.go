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

// checks if the JWKS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JWKS{}

/*
	types and functions for keys
*/

// isArray
type JWKSGetKeysAttributeType = *[]JWK
type JWKSGetKeysArgType = []JWK
type JWKSGetKeysRetType = []JWK

func getJWKSGetKeysAttributeTypeOk(arg JWKSGetKeysAttributeType) (ret JWKSGetKeysRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setJWKSGetKeysAttributeType(arg *JWKSGetKeysAttributeType, val JWKSGetKeysRetType) {
	*arg = &val
}

// JWKS struct for JWKS
type JWKS struct {
	// REQUIRED
	Keys JWKSGetKeysAttributeType `json:"keys"`
}

type _JWKS JWKS

// NewJWKS instantiates a new JWKS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJWKS(keys JWKSGetKeysArgType) *JWKS {
	this := JWKS{}
	setJWKSGetKeysAttributeType(&this.Keys, keys)
	return &this
}

// NewJWKSWithDefaults instantiates a new JWKS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJWKSWithDefaults() *JWKS {
	this := JWKS{}
	return &this
}

// GetKeys returns the Keys field value
func (o *JWKS) GetKeys() (ret JWKSGetKeysRetType) {
	ret, _ = o.GetKeysOk()
	return ret
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *JWKS) GetKeysOk() (ret JWKSGetKeysRetType, ok bool) {
	return getJWKSGetKeysAttributeTypeOk(o.Keys)
}

// SetKeys sets field value
func (o *JWKS) SetKeys(v JWKSGetKeysRetType) {
	setJWKSGetKeysAttributeType(&o.Keys, v)
}

func (o JWKS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getJWKSGetKeysAttributeTypeOk(o.Keys); ok {
		toSerialize["Keys"] = val
	}
	return toSerialize, nil
}

type NullableJWKS struct {
	value *JWKS
	isSet bool
}

func (v NullableJWKS) Get() *JWKS {
	return v.value
}

func (v *NullableJWKS) Set(val *JWKS) {
	v.value = val
	v.isSet = true
}

func (v NullableJWKS) IsSet() bool {
	return v.isSet
}

func (v *NullableJWKS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJWKS(val *JWKS) *NullableJWKS {
	return &NullableJWKS{value: val, isSet: true}
}

func (v NullableJWKS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJWKS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
