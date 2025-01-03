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

// checks if the ServiceKeysList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceKeysList{}

// ServiceKeysList struct for ServiceKeysList
type ServiceKeysList struct {
	CredentialsInfo *map[string]string `json:"credentialsInfo,omitempty"`
	// REQUIRED
	Id *string `json:"id"`
	// REQUIRED
	Name *string `json:"name"`
}

type _ServiceKeysList ServiceKeysList

// NewServiceKeysList instantiates a new ServiceKeysList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceKeysList(id *string, name *string) *ServiceKeysList {
	this := ServiceKeysList{}
	this.Id = id
	this.Name = name
	return &this
}

// NewServiceKeysListWithDefaults instantiates a new ServiceKeysList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceKeysListWithDefaults() *ServiceKeysList {
	this := ServiceKeysList{}
	return &this
}

// GetCredentialsInfo returns the CredentialsInfo field value if set, zero value otherwise.
func (o *ServiceKeysList) GetCredentialsInfo() *map[string]string {
	if o == nil || IsNil(o.CredentialsInfo) {
		var ret *map[string]string
		return ret
	}
	return o.CredentialsInfo
}

// GetCredentialsInfoOk returns a tuple with the CredentialsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceKeysList) GetCredentialsInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CredentialsInfo) {
		return nil, false
	}
	return o.CredentialsInfo, true
}

// HasCredentialsInfo returns a boolean if a field has been set.
func (o *ServiceKeysList) HasCredentialsInfo() bool {
	if o != nil && !IsNil(o.CredentialsInfo) {
		return true
	}

	return false
}

// SetCredentialsInfo gets a reference to the given map[string]string and assigns it to the CredentialsInfo field.
func (o *ServiceKeysList) SetCredentialsInfo(v *map[string]string) {
	o.CredentialsInfo = v
}

// GetId returns the Id field value
func (o *ServiceKeysList) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceKeysList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *ServiceKeysList) SetId(v *string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ServiceKeysList) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceKeysList) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *ServiceKeysList) SetName(v *string) {
	o.Name = v
}

func (o ServiceKeysList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CredentialsInfo) {
		toSerialize["credentialsInfo"] = o.CredentialsInfo
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableServiceKeysList struct {
	value *ServiceKeysList
	isSet bool
}

func (v NullableServiceKeysList) Get() *ServiceKeysList {
	return v.value
}

func (v *NullableServiceKeysList) Set(val *ServiceKeysList) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceKeysList) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceKeysList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceKeysList(val *ServiceKeysList) *NullableServiceKeysList {
	return &NullableServiceKeysList{value: val, isSet: true}
}

func (v NullableServiceKeysList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceKeysList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
