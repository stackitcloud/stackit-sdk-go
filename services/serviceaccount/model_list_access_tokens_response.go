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

// checks if the ListAccessTokensResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAccessTokensResponse{}

/*
	types and functions for items
*/

// isArray
type ListAccessTokensResponseGetItemsAttributeType = *[]AccessTokenMetadata
type ListAccessTokensResponseGetItemsArgType = []AccessTokenMetadata
type ListAccessTokensResponseGetItemsRetType = []AccessTokenMetadata

func getListAccessTokensResponseGetItemsAttributeTypeOk(arg ListAccessTokensResponseGetItemsAttributeType) (ret ListAccessTokensResponseGetItemsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListAccessTokensResponseGetItemsAttributeType(arg *ListAccessTokensResponseGetItemsAttributeType, val ListAccessTokensResponseGetItemsRetType) {
	*arg = &val
}

// ListAccessTokensResponse struct for ListAccessTokensResponse
type ListAccessTokensResponse struct {
	Items ListAccessTokensResponseGetItemsAttributeType `json:"items,omitempty"`
}

// NewListAccessTokensResponse instantiates a new ListAccessTokensResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAccessTokensResponse() *ListAccessTokensResponse {
	this := ListAccessTokensResponse{}
	return &this
}

// NewListAccessTokensResponseWithDefaults instantiates a new ListAccessTokensResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAccessTokensResponseWithDefaults() *ListAccessTokensResponse {
	this := ListAccessTokensResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListAccessTokensResponse) GetItems() (res ListAccessTokensResponseGetItemsRetType) {
	res, _ = o.GetItemsOk()
	return
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccessTokensResponse) GetItemsOk() (ret ListAccessTokensResponseGetItemsRetType, ok bool) {
	return getListAccessTokensResponseGetItemsAttributeTypeOk(o.Items)
}

// HasItems returns a boolean if a field has been set.
func (o *ListAccessTokensResponse) HasItems() bool {
	_, ok := o.GetItemsOk()
	return ok
}

// SetItems gets a reference to the given []AccessTokenMetadata and assigns it to the Items field.
func (o *ListAccessTokensResponse) SetItems(v ListAccessTokensResponseGetItemsRetType) {
	setListAccessTokensResponseGetItemsAttributeType(&o.Items, v)
}

func (o ListAccessTokensResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListAccessTokensResponseGetItemsAttributeTypeOk(o.Items); ok {
		toSerialize["Items"] = val
	}
	return toSerialize, nil
}

type NullableListAccessTokensResponse struct {
	value *ListAccessTokensResponse
	isSet bool
}

func (v NullableListAccessTokensResponse) Get() *ListAccessTokensResponse {
	return v.value
}

func (v *NullableListAccessTokensResponse) Set(val *ListAccessTokensResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccessTokensResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccessTokensResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccessTokensResponse(val *ListAccessTokensResponse) *NullableListAccessTokensResponse {
	return &NullableListAccessTokensResponse{value: val, isSet: true}
}

func (v NullableListAccessTokensResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccessTokensResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
