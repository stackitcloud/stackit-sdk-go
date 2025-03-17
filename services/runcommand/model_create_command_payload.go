/*
STACKIT Run Commands Service API

API endpoints for the STACKIT Run Commands Service API

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package runcommand

import (
	"encoding/json"
)

// checks if the CreateCommandPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCommandPayload{}

/*
	types and functions for commandTemplateName
*/

// isNotNullableString
type CreateCommandPayloadGetCommandTemplateNameAttributeType = *string

func getCreateCommandPayloadGetCommandTemplateNameAttributeTypeOk(arg CreateCommandPayloadGetCommandTemplateNameAttributeType) (ret CreateCommandPayloadGetCommandTemplateNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateCommandPayloadGetCommandTemplateNameAttributeType(arg *CreateCommandPayloadGetCommandTemplateNameAttributeType, val CreateCommandPayloadGetCommandTemplateNameRetType) {
	*arg = &val
}

type CreateCommandPayloadGetCommandTemplateNameArgType = string
type CreateCommandPayloadGetCommandTemplateNameRetType = string

/*
	types and functions for parameters
*/

// isContainer
type CreateCommandPayloadGetParametersAttributeType = *map[string]string
type CreateCommandPayloadGetParametersArgType = map[string]string
type CreateCommandPayloadGetParametersRetType = map[string]string

func getCreateCommandPayloadGetParametersAttributeTypeOk(arg CreateCommandPayloadGetParametersAttributeType) (ret CreateCommandPayloadGetParametersRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateCommandPayloadGetParametersAttributeType(arg *CreateCommandPayloadGetParametersAttributeType, val CreateCommandPayloadGetParametersRetType) {
	*arg = &val
}

// CreateCommandPayload struct for CreateCommandPayload
type CreateCommandPayload struct {
	// REQUIRED
	CommandTemplateName CreateCommandPayloadGetCommandTemplateNameAttributeType `json:"commandTemplateName"`
	Parameters          CreateCommandPayloadGetParametersAttributeType          `json:"parameters,omitempty"`
}

type _CreateCommandPayload CreateCommandPayload

// NewCreateCommandPayload instantiates a new CreateCommandPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCommandPayload(commandTemplateName CreateCommandPayloadGetCommandTemplateNameArgType) *CreateCommandPayload {
	this := CreateCommandPayload{}
	setCreateCommandPayloadGetCommandTemplateNameAttributeType(&this.CommandTemplateName, commandTemplateName)
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
func (o *CreateCommandPayload) GetCommandTemplateName() (ret CreateCommandPayloadGetCommandTemplateNameRetType) {
	ret, _ = o.GetCommandTemplateNameOk()
	return ret
}

// GetCommandTemplateNameOk returns a tuple with the CommandTemplateName field value
// and a boolean to check if the value has been set.
func (o *CreateCommandPayload) GetCommandTemplateNameOk() (ret CreateCommandPayloadGetCommandTemplateNameRetType, ok bool) {
	return getCreateCommandPayloadGetCommandTemplateNameAttributeTypeOk(o.CommandTemplateName)
}

// SetCommandTemplateName sets field value
func (o *CreateCommandPayload) SetCommandTemplateName(v CreateCommandPayloadGetCommandTemplateNameRetType) {
	setCreateCommandPayloadGetCommandTemplateNameAttributeType(&o.CommandTemplateName, v)
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *CreateCommandPayload) GetParameters() (res CreateCommandPayloadGetParametersRetType) {
	res, _ = o.GetParametersOk()
	return
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCommandPayload) GetParametersOk() (ret CreateCommandPayloadGetParametersRetType, ok bool) {
	return getCreateCommandPayloadGetParametersAttributeTypeOk(o.Parameters)
}

// HasParameters returns a boolean if a field has been set.
func (o *CreateCommandPayload) HasParameters() bool {
	_, ok := o.GetParametersOk()
	return ok
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *CreateCommandPayload) SetParameters(v CreateCommandPayloadGetParametersRetType) {
	setCreateCommandPayloadGetParametersAttributeType(&o.Parameters, v)
}

func (o CreateCommandPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateCommandPayloadGetCommandTemplateNameAttributeTypeOk(o.CommandTemplateName); ok {
		toSerialize["CommandTemplateName"] = val
	}
	if val, ok := getCreateCommandPayloadGetParametersAttributeTypeOk(o.Parameters); ok {
		toSerialize["Parameters"] = val
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
