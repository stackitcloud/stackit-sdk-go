/*
STACKIT Run Commands Service API

API endpoints for the STACKIT Run Commands Service API

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package runcommand

import (
	"encoding/json"
)

// checks if the CreateCommandPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCommandPayload{}

// CreateCommandPayload struct for CreateCommandPayload
type CreateCommandPayload struct {
	// REQUIRED
	CommandTemplateName *string            `json:"commandTemplateName"`
	Parameters          *map[string]string `json:"parameters,omitempty"`
}

type _CreateCommandPayload CreateCommandPayload

// NewCreateCommandPayload instantiates a new CreateCommandPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCommandPayload(commandTemplateName *string) *CreateCommandPayload {
	this := CreateCommandPayload{}
	this.CommandTemplateName = commandTemplateName
	return &this
}

// NewCreateCommandPayloadWithDefaults instantiates a new CreateCommandPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCommandPayloadWithDefaults() *CreateCommandPayload {
	this := CreateCommandPayload{}
	return &this
}

// GetCommandTemplateName returns the CommandTemplateName field value
func (o *CreateCommandPayload) GetCommandTemplateName() *string {
	if o == nil || IsNil(o.CommandTemplateName) {
		var ret *string
		return ret
	}

	return o.CommandTemplateName
}

// GetCommandTemplateNameOk returns a tuple with the CommandTemplateName field value
// and a boolean to check if the value has been set.
func (o *CreateCommandPayload) GetCommandTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommandTemplateName, true
}

// SetCommandTemplateName sets field value
func (o *CreateCommandPayload) SetCommandTemplateName(v *string) {
	o.CommandTemplateName = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *CreateCommandPayload) GetParameters() *map[string]string {
	if o == nil || IsNil(o.Parameters) {
		var ret *map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCommandPayload) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *CreateCommandPayload) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *CreateCommandPayload) SetParameters(v *map[string]string) {
	o.Parameters = v
}

func (o CreateCommandPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commandTemplateName"] = o.CommandTemplateName
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableCreateCommandPayload struct {
	value *CreateCommandPayload
	isSet bool
}

func (v NullableCreateCommandPayload) Get() *CreateCommandPayload {
	return v.value
}

func (v *NullableCreateCommandPayload) Set(val *CreateCommandPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCommandPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCommandPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCommandPayload(val *CreateCommandPayload) *NullableCreateCommandPayload {
	return &NullableCreateCommandPayload{value: val, isSet: true}
}

func (v NullableCreateCommandPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCommandPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
