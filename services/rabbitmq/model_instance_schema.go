/*
STACKIT RabbitMQ API

The STACKIT RabbitMQ API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rabbitmq

import (
	"encoding/json"
)

// checks if the InstanceSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceSchema{}

// InstanceSchema struct for InstanceSchema
type InstanceSchema struct {
	// REQUIRED
	Create *Schema `json:"create"`
	// REQUIRED
	Update *Schema `json:"update"`
}

type _InstanceSchema InstanceSchema

// NewInstanceSchema instantiates a new InstanceSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceSchema(create *Schema, update *Schema) *InstanceSchema {
	this := InstanceSchema{}
	this.Create = create
	this.Update = update
	return &this
}

// NewInstanceSchemaWithDefaults instantiates a new InstanceSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceSchemaWithDefaults() *InstanceSchema {
	this := InstanceSchema{}
	return &this
}

// GetCreate returns the Create field value
func (o *InstanceSchema) GetCreate() *Schema {
	if o == nil {
		var ret *Schema
		return ret
	}

	return o.Create
}

// GetCreateOk returns a tuple with the Create field value
// and a boolean to check if the value has been set.
func (o *InstanceSchema) GetCreateOk() (*Schema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Create, true
}

// SetCreate sets field value
func (o *InstanceSchema) SetCreate(v *Schema) {
	o.Create = v
}

// GetUpdate returns the Update field value
func (o *InstanceSchema) GetUpdate() *Schema {
	if o == nil {
		var ret *Schema
		return ret
	}

	return o.Update
}

// GetUpdateOk returns a tuple with the Update field value
// and a boolean to check if the value has been set.
func (o *InstanceSchema) GetUpdateOk() (*Schema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Update, true
}

// SetUpdate sets field value
func (o *InstanceSchema) SetUpdate(v *Schema) {
	o.Update = v
}

func (o InstanceSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["create"] = o.Create
	toSerialize["update"] = o.Update
	return toSerialize, nil
}

type NullableInstanceSchema struct {
	value *InstanceSchema
	isSet bool
}

func (v NullableInstanceSchema) Get() *InstanceSchema {
	return v.value
}

func (v *NullableInstanceSchema) Set(val *InstanceSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceSchema(val *InstanceSchema) *NullableInstanceSchema {
	return &NullableInstanceSchema{value: val, isSet: true}
}

func (v NullableInstanceSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
