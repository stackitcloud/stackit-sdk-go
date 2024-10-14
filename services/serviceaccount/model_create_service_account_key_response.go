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

// checks if the CreateServiceAccountKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServiceAccountKeyResponse{}

// CreateServiceAccountKeyResponse struct for CreateServiceAccountKeyResponse
type CreateServiceAccountKeyResponse struct {
	// REQUIRED
	Active *bool `json:"active"`
	// Creation time of the key
	// REQUIRED
	CreatedAt *time.Time `json:"createdAt"`
	// REQUIRED
	Credentials *CreateServiceAccountKeyResponseCredentials `json:"credentials"`
	// Unique ID of the key.
	// REQUIRED
	Id *string `json:"id"`
	// REQUIRED
	KeyAlgorithm *string `json:"keyAlgorithm"`
	// REQUIRED
	KeyOrigin *string `json:"keyOrigin"`
	// REQUIRED
	KeyType *string `json:"keyType"`
	// Public key, that was provider, or was generated by the service account API
	// REQUIRED
	PublicKey *string `json:"publicKey"`
	// If specified, the timestamp until the key is active. May be null
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

type _CreateServiceAccountKeyResponse CreateServiceAccountKeyResponse

// NewCreateServiceAccountKeyResponse instantiates a new CreateServiceAccountKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServiceAccountKeyResponse(active *bool, createdAt *time.Time, credentials *CreateServiceAccountKeyResponseCredentials, id *string, keyAlgorithm *string, keyOrigin *string, keyType *string, publicKey *string) *CreateServiceAccountKeyResponse {
	this := CreateServiceAccountKeyResponse{}
	this.Active = active
	this.CreatedAt = createdAt
	this.Credentials = credentials
	this.Id = id
	this.KeyAlgorithm = keyAlgorithm
	this.KeyOrigin = keyOrigin
	this.KeyType = keyType
	this.PublicKey = publicKey
	return &this
}

// NewCreateServiceAccountKeyResponseWithDefaults instantiates a new CreateServiceAccountKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServiceAccountKeyResponseWithDefaults() *CreateServiceAccountKeyResponse {
	this := CreateServiceAccountKeyResponse{}
	return &this
}

// GetActive returns the Active field value
func (o *CreateServiceAccountKeyResponse) GetActive() *bool {
	if o == nil {
		var ret *bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountKeyResponse) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active, true
}

// SetActive sets field value
func (o *CreateServiceAccountKeyResponse) SetActive(v *bool) {
	o.Active = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CreateServiceAccountKeyResponse) GetCreatedAt() *time.Time {
	if o == nil {
		var ret *time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountKeyResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CreateServiceAccountKeyResponse) SetCreatedAt(v *time.Time) {
	o.CreatedAt = v
}

// GetCredentials returns the Credentials field value
func (o *CreateServiceAccountKeyResponse) GetCredentials() *CreateServiceAccountKeyResponseCredentials {
	if o == nil {
		var ret *CreateServiceAccountKeyResponseCredentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountKeyResponse) GetCredentialsOk() (*CreateServiceAccountKeyResponseCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials, true
}

// SetCredentials sets field value
func (o *CreateServiceAccountKeyResponse) SetCredentials(v *CreateServiceAccountKeyResponseCredentials) {
	o.Credentials = v
}

// GetId returns the Id field value
func (o *CreateServiceAccountKeyResponse) GetId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountKeyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *CreateServiceAccountKeyResponse) SetId(v *string) {
	o.Id = v
}

// GetKeyAlgorithm returns the KeyAlgorithm field value
func (o *CreateServiceAccountKeyResponse) GetKeyAlgorithm() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.KeyAlgorithm
}

// GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountKeyResponse) GetKeyAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyAlgorithm, true
}

// SetKeyAlgorithm sets field value
func (o *CreateServiceAccountKeyResponse) SetKeyAlgorithm(v *string) {
	o.KeyAlgorithm = v
}

// GetKeyOrigin returns the KeyOrigin field value
func (o *CreateServiceAccountKeyResponse) GetKeyOrigin() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.KeyOrigin
}

// GetKeyOriginOk returns a tuple with the KeyOrigin field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountKeyResponse) GetKeyOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyOrigin, true
}

// SetKeyOrigin sets field value
func (o *CreateServiceAccountKeyResponse) SetKeyOrigin(v *string) {
	o.KeyOrigin = v
}

// GetKeyType returns the KeyType field value
func (o *CreateServiceAccountKeyResponse) GetKeyType() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountKeyResponse) GetKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyType, true
}

// SetKeyType sets field value
func (o *CreateServiceAccountKeyResponse) SetKeyType(v *string) {
	o.KeyType = v
}

// GetPublicKey returns the PublicKey field value
func (o *CreateServiceAccountKeyResponse) GetPublicKey() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountKeyResponse) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// SetPublicKey sets field value
func (o *CreateServiceAccountKeyResponse) SetPublicKey(v *string) {
	o.PublicKey = v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *CreateServiceAccountKeyResponse) GetValidUntil() *time.Time {
	if o == nil || IsNil(o.ValidUntil) {
		var ret *time.Time
		return ret
	}
	return o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountKeyResponse) GetValidUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidUntil) {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *CreateServiceAccountKeyResponse) HasValidUntil() bool {
	if o != nil && !IsNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.
func (o *CreateServiceAccountKeyResponse) SetValidUntil(v *time.Time) {
	o.ValidUntil = v
}

func (o CreateServiceAccountKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["credentials"] = o.Credentials
	toSerialize["id"] = o.Id
	toSerialize["keyAlgorithm"] = o.KeyAlgorithm
	toSerialize["keyOrigin"] = o.KeyOrigin
	toSerialize["keyType"] = o.KeyType
	toSerialize["publicKey"] = o.PublicKey
	if !IsNil(o.ValidUntil) {
		toSerialize["validUntil"] = o.ValidUntil
	}
	return toSerialize, nil
}

type NullableCreateServiceAccountKeyResponse struct {
	value *CreateServiceAccountKeyResponse
	isSet bool
}

func (v NullableCreateServiceAccountKeyResponse) Get() *CreateServiceAccountKeyResponse {
	return v.value
}

func (v *NullableCreateServiceAccountKeyResponse) Set(val *CreateServiceAccountKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServiceAccountKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServiceAccountKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServiceAccountKeyResponse(val *CreateServiceAccountKeyResponse) *NullableCreateServiceAccountKeyResponse {
	return &NullableCreateServiceAccountKeyResponse{value: val, isSet: true}
}

func (v NullableCreateServiceAccountKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServiceAccountKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
