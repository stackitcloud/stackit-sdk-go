/*
Service Account API

API to manage Service Accounts and their Access Tokens.  ### System for Cross-domain Identity Management (SCIM) Service Account Service offers SCIM APIs to query state. The SCIM protocol was created as standard for  automating the exchange of user identity information between identity domains, or IT systems. Service accounts  are be handled as indentites similar to SCIM users. A custom SCIM schema has been created: `/ServiceAccounts`  #### Syntax ##### Attribute operators | OPERATOR | DESCRIPTION              | |----------|--------------------------| | eq       | equal                    | | ne       | not equal                | | co       | contains                 | | sw       | starts with              | | ew       | ends with                |  ##### Logical operators | OPERATOR | DESCRIPTION              | |----------|--------------------------| | and      | logical \"and\"            | | or       | logical \"or\"             |  ##### Grouping operators | OPERATOR | DESCRIPTION              | |----------|--------------------------| | ()       | precending grouping      |  ##### Example ``` filter=email eq \"my-service-account-aBc2defg@sa.stackit.cloud\" filter=email ne \"my-service-account-aBc2defg@sa.stackit.cloud\" filter=email co \"my-service-account\" filter=name sw \"my\" filter=name ew \"account\" filter=email co \"my-service-account\" and name sw \"my\" filter=email co \"my-service-account\" and (name sw \"my\" or name ew \"account\") ```  #### Sorting  > Sorting is optional  | PARAMETER | DESCRIPTION                          | |-----------|--------------------------------------| | sortBy    | attribute response is ordered by     | | sortOrder | 'ASCENDING' (default) or 'DESCENDING'|  #### Pagination  | PARAMETER    | DESCRIPTION                                  | |--------------|----------------------------------------------| | startIndex   | index of first query result, default: 1      | | count        | maximum number of query results, default: 100|

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceaccount

import (
	"encoding/json"
	"time"
)

// checks if the AccessTokenMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenMetadata{}

// AccessTokenMetadata Does not contain the actual token.
type AccessTokenMetadata struct {
	// If true, access token can be used for authorized API calls, if false, the token is not usable anymore.
	// REQUIRED
	Active *bool `json:"active"`
	// Creation time of the access token.
	// REQUIRED
	CreatedAt *time.Time `json:"createdAt"`
	// Unique ID of the access token. Also used as JTI field.
	// REQUIRED
	Id *string `json:"id"`
	// Approximate expiration time of the access token. Check the JWT for actual validity date.
	// REQUIRED
	ValidUntil *time.Time `json:"validUntil"`
}

type _AccessTokenMetadata AccessTokenMetadata

// NewAccessTokenMetadata instantiates a new AccessTokenMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenMetadata(active *bool, createdAt *time.Time, id *string, validUntil *time.Time) *AccessTokenMetadata {
	this := AccessTokenMetadata{}
	this.Active = active
	this.CreatedAt = createdAt
	this.Id = id
	this.ValidUntil = validUntil
	return &this
}

// NewAccessTokenMetadataWithDefaults instantiates a new AccessTokenMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenMetadataWithDefaults() *AccessTokenMetadata {
	this := AccessTokenMetadata{}
	return &this
}

// GetActive returns the Active field value
func (o *AccessTokenMetadata) GetActive() *bool {
	if o == nil {
		var ret *bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *AccessTokenMetadata) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active, true
}

// SetActive sets field value
func (o *AccessTokenMetadata) SetActive(v *bool) {
	o.Active = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AccessTokenMetadata) GetCreatedAt() *time.Time {
	if o == nil {
		var ret *time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AccessTokenMetadata) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AccessTokenMetadata) SetCreatedAt(v *time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *AccessTokenMetadata) GetId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessTokenMetadata) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *AccessTokenMetadata) SetId(v *string) {
	o.Id = v
}

// GetValidUntil returns the ValidUntil field value
func (o *AccessTokenMetadata) GetValidUntil() *time.Time {
	if o == nil {
		var ret *time.Time
		return ret
	}

	return o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value
// and a boolean to check if the value has been set.
func (o *AccessTokenMetadata) GetValidUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidUntil, true
}

// SetValidUntil sets field value
func (o *AccessTokenMetadata) SetValidUntil(v *time.Time) {
	o.ValidUntil = v
}

func (o AccessTokenMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["validUntil"] = o.ValidUntil
	return toSerialize, nil
}

type NullableAccessTokenMetadata struct {
	value *AccessTokenMetadata
	isSet bool
}

func (v NullableAccessTokenMetadata) Get() *AccessTokenMetadata {
	return v.value
}

func (v *NullableAccessTokenMetadata) Set(val *AccessTokenMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenMetadata(val *AccessTokenMetadata) *NullableAccessTokenMetadata {
	return &NullableAccessTokenMetadata{value: val, isSet: true}
}

func (v NullableAccessTokenMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
