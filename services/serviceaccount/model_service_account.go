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

// checks if the ServiceAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccount{}

// ServiceAccount struct for ServiceAccount
type ServiceAccount struct {
	// Unique identifier of the service account in format of an email address generated by the service containing the prefix provided by the user during creation.
	// REQUIRED
	Email *string `json:"email"`
	// Unique ID of the service account. It is also used in the 'sub' field of the service accounts access tokens.
	// REQUIRED
	Id *string `json:"id"`
	// Flag indicating internal service accounts
	// REQUIRED
	Internal *bool `json:"internal"`
	// ID of the related project
	// REQUIRED
	ProjectId *string `json:"projectId"`
}

type _ServiceAccount ServiceAccount

// NewServiceAccount instantiates a new ServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccount(email *string, id *string, internal *bool, projectId *string) *ServiceAccount {
	this := ServiceAccount{}
	this.Email = email
	this.Id = id
	this.Internal = internal
	this.ProjectId = projectId
	return &this
}

// NewServiceAccountWithDefaults instantiates a new ServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountWithDefaults() *ServiceAccount {
	this := ServiceAccount{}
	return &this
}

// GetEmail returns the Email field value
func (o *ServiceAccount) GetEmail() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email, true
}

// SetEmail sets field value
func (o *ServiceAccount) SetEmail(v *string) {
	o.Email = v
}

// GetId returns the Id field value
func (o *ServiceAccount) GetId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *ServiceAccount) SetId(v *string) {
	o.Id = v
}

// GetInternal returns the Internal field value
func (o *ServiceAccount) GetInternal() *bool {
	if o == nil {
		var ret *bool
		return ret
	}

	return o.Internal
}

// GetInternalOk returns a tuple with the Internal field value
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetInternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Internal, true
}

// SetInternal sets field value
func (o *ServiceAccount) SetInternal(v *bool) {
	o.Internal = v
}

// GetProjectId returns the ProjectId field value
func (o *ServiceAccount) GetProjectId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// SetProjectId sets field value
func (o *ServiceAccount) SetProjectId(v *string) {
	o.ProjectId = v
}

func (o ServiceAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["id"] = o.Id
	toSerialize["internal"] = o.Internal
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableServiceAccount struct {
	value *ServiceAccount
	isSet bool
}

func (v NullableServiceAccount) Get() *ServiceAccount {
	return v.value
}

func (v *NullableServiceAccount) Set(val *ServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccount(val *ServiceAccount) *NullableServiceAccount {
	return &NullableServiceAccount{value: val, isSet: true}
}

func (v NullableServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
