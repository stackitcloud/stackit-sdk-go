/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package membership

import (
	"encoding/json"
)

// checks if the WriteSchemaPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WriteSchemaPayload{}

// WriteSchemaPayload struct for WriteSchemaPayload
type WriteSchemaPayload struct {
	// REQUIRED
	Schema *string `json:"schema"`
}

type _WriteSchemaPayload WriteSchemaPayload

// NewWriteSchemaPayload instantiates a new WriteSchemaPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteSchemaPayload(schema *string) *WriteSchemaPayload {
	this := WriteSchemaPayload{}
	this.Schema = schema
	return &this
}

// NewWriteSchemaPayloadWithDefaults instantiates a new WriteSchemaPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteSchemaPayloadWithDefaults() *WriteSchemaPayload {
	this := WriteSchemaPayload{}
	return &this
}

// GetSchema returns the Schema field value
func (o *WriteSchemaPayload) GetSchema() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *WriteSchemaPayload) GetSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schema, true
}

// SetSchema sets field value
func (o *WriteSchemaPayload) SetSchema(v *string) {
	o.Schema = v
}

func (o WriteSchemaPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schema"] = o.Schema
	return toSerialize, nil
}

type NullableWriteSchemaPayload struct {
	value *WriteSchemaPayload
	isSet bool
}

func (v NullableWriteSchemaPayload) Get() *WriteSchemaPayload {
	return v.value
}

func (v *NullableWriteSchemaPayload) Set(val *WriteSchemaPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteSchemaPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteSchemaPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteSchemaPayload(val *WriteSchemaPayload) *NullableWriteSchemaPayload {
	return &NullableWriteSchemaPayload{value: val, isSet: true}
}

func (v NullableWriteSchemaPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteSchemaPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
