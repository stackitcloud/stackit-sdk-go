/*
STACKIT Secrets Manager API

This API provides endpoints for managing the Secrets-Manager.

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretsmanager

import (
	"encoding/json"
)

// checks if the CreateACLPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateACLPayload{}

// CreateACLPayload struct for CreateACLPayload
type CreateACLPayload struct {
	// The given IP/IP Range that is permitted to access.
	// REQUIRED
	Cidr *string `json:"cidr"`
}

type _CreateACLPayload CreateACLPayload

// NewCreateACLPayload instantiates a new CreateACLPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateACLPayload(cidr *string) *CreateACLPayload {
	this := CreateACLPayload{}
	this.Cidr = cidr
	return &this
}

// NewCreateACLPayloadWithDefaults instantiates a new CreateACLPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateACLPayloadWithDefaults() *CreateACLPayload {
	this := CreateACLPayload{}
	return &this
}

// GetCidr returns the Cidr field value
func (o *CreateACLPayload) GetCidr() *string {
	if o == nil || IsNil(o.Cidr) {
		var ret *string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *CreateACLPayload) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cidr, true
}

// SetCidr sets field value
func (o *CreateACLPayload) SetCidr(v *string) {
	o.Cidr = v
}

func (o CreateACLPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cidr"] = o.Cidr
	return toSerialize, nil
}

type NullableCreateACLPayload struct {
	value *CreateACLPayload
	isSet bool
}

func (v NullableCreateACLPayload) Get() *CreateACLPayload {
	return v.value
}

func (v *NullableCreateACLPayload) Set(val *CreateACLPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateACLPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateACLPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateACLPayload(val *CreateACLPayload) *NullableCreateACLPayload {
	return &NullableCreateACLPayload{value: val, isSet: true}
}

func (v NullableCreateACLPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateACLPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
