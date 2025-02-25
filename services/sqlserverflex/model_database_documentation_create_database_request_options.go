/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
)

// checks if the DatabaseDocumentationCreateDatabaseRequestOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseDocumentationCreateDatabaseRequestOptions{}

// DatabaseDocumentationCreateDatabaseRequestOptions struct for DatabaseDocumentationCreateDatabaseRequestOptions
type DatabaseDocumentationCreateDatabaseRequestOptions struct {
	// Collation of the database
	Collation *string `json:"collation,omitempty"`
	// CompatibilityLevel of the Database.
	CompatibilityLevel *string `json:"compatibilityLevel,omitempty"`
	// Name of the owner of the database.
	// REQUIRED
	Owner *string `json:"owner"`
}

type _DatabaseDocumentationCreateDatabaseRequestOptions DatabaseDocumentationCreateDatabaseRequestOptions

// NewDatabaseDocumentationCreateDatabaseRequestOptions instantiates a new DatabaseDocumentationCreateDatabaseRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseDocumentationCreateDatabaseRequestOptions(owner *string) *DatabaseDocumentationCreateDatabaseRequestOptions {
	this := DatabaseDocumentationCreateDatabaseRequestOptions{}
	this.Owner = owner
	return &this
}

// NewDatabaseDocumentationCreateDatabaseRequestOptionsWithDefaults instantiates a new DatabaseDocumentationCreateDatabaseRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseDocumentationCreateDatabaseRequestOptionsWithDefaults() *DatabaseDocumentationCreateDatabaseRequestOptions {
	this := DatabaseDocumentationCreateDatabaseRequestOptions{}
	return &this
}

// GetCollation returns the Collation field value if set, zero value otherwise.
func (o *DatabaseDocumentationCreateDatabaseRequestOptions) GetCollation() *string {
	if o == nil || IsNil(o.Collation) {
		var ret *string
		return ret
	}
	return o.Collation
}

// GetCollationOk returns a tuple with the Collation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseDocumentationCreateDatabaseRequestOptions) GetCollationOk() (*string, bool) {
	if o == nil || IsNil(o.Collation) {
		return nil, false
	}
	return o.Collation, true
}

// HasCollation returns a boolean if a field has been set.
func (o *DatabaseDocumentationCreateDatabaseRequestOptions) HasCollation() bool {
	if o != nil && !IsNil(o.Collation) {
		return true
	}

	return false
}

// SetCollation gets a reference to the given string and assigns it to the Collation field.
func (o *DatabaseDocumentationCreateDatabaseRequestOptions) SetCollation(v *string) {
	o.Collation = v
}

// GetCompatibilityLevel returns the CompatibilityLevel field value if set, zero value otherwise.
func (o *DatabaseDocumentationCreateDatabaseRequestOptions) GetCompatibilityLevel() *string {
	if o == nil || IsNil(o.CompatibilityLevel) {
		var ret *string
		return ret
	}
	return o.CompatibilityLevel
}

// GetCompatibilityLevelOk returns a tuple with the CompatibilityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseDocumentationCreateDatabaseRequestOptions) GetCompatibilityLevelOk() (*string, bool) {
	if o == nil || IsNil(o.CompatibilityLevel) {
		return nil, false
	}
	return o.CompatibilityLevel, true
}

// HasCompatibilityLevel returns a boolean if a field has been set.
func (o *DatabaseDocumentationCreateDatabaseRequestOptions) HasCompatibilityLevel() bool {
	if o != nil && !IsNil(o.CompatibilityLevel) {
		return true
	}

	return false
}

// SetCompatibilityLevel gets a reference to the given string and assigns it to the CompatibilityLevel field.
func (o *DatabaseDocumentationCreateDatabaseRequestOptions) SetCompatibilityLevel(v *string) {
	o.CompatibilityLevel = v
}

// GetOwner returns the Owner field value
func (o *DatabaseDocumentationCreateDatabaseRequestOptions) GetOwner() *string {
	if o == nil || IsNil(o.Owner) {
		var ret *string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *DatabaseDocumentationCreateDatabaseRequestOptions) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner, true
}

// SetOwner sets field value
func (o *DatabaseDocumentationCreateDatabaseRequestOptions) SetOwner(v *string) {
	o.Owner = v
}

func (o DatabaseDocumentationCreateDatabaseRequestOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collation) {
		toSerialize["collation"] = o.Collation
	}
	if !IsNil(o.CompatibilityLevel) {
		toSerialize["compatibilityLevel"] = o.CompatibilityLevel
	}
	toSerialize["owner"] = o.Owner
	return toSerialize, nil
}

type NullableDatabaseDocumentationCreateDatabaseRequestOptions struct {
	value *DatabaseDocumentationCreateDatabaseRequestOptions
	isSet bool
}

func (v NullableDatabaseDocumentationCreateDatabaseRequestOptions) Get() *DatabaseDocumentationCreateDatabaseRequestOptions {
	return v.value
}

func (v *NullableDatabaseDocumentationCreateDatabaseRequestOptions) Set(val *DatabaseDocumentationCreateDatabaseRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseDocumentationCreateDatabaseRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseDocumentationCreateDatabaseRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseDocumentationCreateDatabaseRequestOptions(val *DatabaseDocumentationCreateDatabaseRequestOptions) *NullableDatabaseDocumentationCreateDatabaseRequestOptions {
	return &NullableDatabaseDocumentationCreateDatabaseRequestOptions{value: val, isSet: true}
}

func (v NullableDatabaseDocumentationCreateDatabaseRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseDocumentationCreateDatabaseRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
