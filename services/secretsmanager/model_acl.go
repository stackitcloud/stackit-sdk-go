/*
STACKIT Secrets Manager API

This API provides endpoints for managing the Secrets-Manager.

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretsmanager

import (
	"encoding/json"
)

// checks if the ACL type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACL{}

// ACL struct for ACL
type ACL struct {
	// The given IP/IP Range that is permitted to access.
	// REQUIRED
	Cidr *string `json:"cidr"`
	// A auto generated unique id which identifies the acl.
	// REQUIRED
	Id *string `json:"id"`
}

type _ACL ACL

// NewACL instantiates a new ACL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACL(cidr *string, id *string) *ACL {
	this := ACL{}
	this.Cidr = cidr
	this.Id = id
	return &this
}

// NewACLWithDefaults instantiates a new ACL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACLWithDefaults() *ACL {
	this := ACL{}
	return &this
}

// GetCidr returns the Cidr field value
func (o *ACL) GetCidr() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *ACL) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cidr, true
}

// SetCidr sets field value
func (o *ACL) SetCidr(v *string) {
	o.Cidr = v
}

// GetId returns the Id field value
func (o *ACL) GetId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ACL) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *ACL) SetId(v *string) {
	o.Id = v
}

func (o ACL) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cidr"] = o.Cidr
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableACL struct {
	value *ACL
	isSet bool
}

func (v NullableACL) Get() *ACL {
	return v.value
}

func (v *NullableACL) Set(val *ACL) {
	v.value = val
	v.isSet = true
}

func (v NullableACL) IsSet() bool {
	return v.isSet
}

func (v *NullableACL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACL(val *ACL) *NullableACL {
	return &NullableACL{value: val, isSet: true}
}

func (v NullableACL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
