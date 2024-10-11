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

// checks if the ListServiceAccountKeysResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListServiceAccountKeysResponse{}

// ListServiceAccountKeysResponse struct for ListServiceAccountKeysResponse
type ListServiceAccountKeysResponse struct {
	// REQUIRED
	Items *[]ServiceAccountKeyListResponse `json:"items"`
}

type _ListServiceAccountKeysResponse ListServiceAccountKeysResponse

// NewListServiceAccountKeysResponse instantiates a new ListServiceAccountKeysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListServiceAccountKeysResponse(items *[]ServiceAccountKeyListResponse) *ListServiceAccountKeysResponse {
	this := ListServiceAccountKeysResponse{}
	this.Items = items
	return &this
}

// NewListServiceAccountKeysResponseWithDefaults instantiates a new ListServiceAccountKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListServiceAccountKeysResponseWithDefaults() *ListServiceAccountKeysResponse {
	this := ListServiceAccountKeysResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *ListServiceAccountKeysResponse) GetItems() *[]ServiceAccountKeyListResponse {
	if o == nil {
		var ret *[]ServiceAccountKeyListResponse
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListServiceAccountKeysResponse) GetItemsOk() (*[]ServiceAccountKeyListResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListServiceAccountKeysResponse) SetItems(v *[]ServiceAccountKeyListResponse) {
	o.Items = v
}

func (o ListServiceAccountKeysResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableListServiceAccountKeysResponse struct {
	value *ListServiceAccountKeysResponse
	isSet bool
}

func (v NullableListServiceAccountKeysResponse) Get() *ListServiceAccountKeysResponse {
	return v.value
}

func (v *NullableListServiceAccountKeysResponse) Set(val *ListServiceAccountKeysResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListServiceAccountKeysResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListServiceAccountKeysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListServiceAccountKeysResponse(val *ListServiceAccountKeysResponse) *NullableListServiceAccountKeysResponse {
	return &NullableListServiceAccountKeysResponse{value: val, isSet: true}
}

func (v NullableListServiceAccountKeysResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListServiceAccountKeysResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
