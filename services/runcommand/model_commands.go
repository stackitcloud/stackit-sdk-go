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

// checks if the Commands type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Commands{}

/*
	types and functions for commandTemplateName
*/

// isNotNullableString
type CommandsGetCommandTemplateNameAttributeType = *string

func getCommandsGetCommandTemplateNameAttributeTypeOk(arg CommandsGetCommandTemplateNameAttributeType) (ret CommandsGetCommandTemplateNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCommandsGetCommandTemplateNameAttributeType(arg *CommandsGetCommandTemplateNameAttributeType, val CommandsGetCommandTemplateNameRetType) {
	*arg = &val
}

type CommandsGetCommandTemplateNameArgType = string
type CommandsGetCommandTemplateNameRetType = string

/*
	types and functions for commandTemplateTitle
*/

// isNotNullableString
type CommandsGetCommandTemplateTitleAttributeType = *string

func getCommandsGetCommandTemplateTitleAttributeTypeOk(arg CommandsGetCommandTemplateTitleAttributeType) (ret CommandsGetCommandTemplateTitleRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCommandsGetCommandTemplateTitleAttributeType(arg *CommandsGetCommandTemplateTitleAttributeType, val CommandsGetCommandTemplateTitleRetType) {
	*arg = &val
}

type CommandsGetCommandTemplateTitleArgType = string
type CommandsGetCommandTemplateTitleRetType = string

/*
	types and functions for finishedAt
*/

// isNotNullableString
type CommandsGetFinishedAtAttributeType = *string

func getCommandsGetFinishedAtAttributeTypeOk(arg CommandsGetFinishedAtAttributeType) (ret CommandsGetFinishedAtRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCommandsGetFinishedAtAttributeType(arg *CommandsGetFinishedAtAttributeType, val CommandsGetFinishedAtRetType) {
	*arg = &val
}

type CommandsGetFinishedAtArgType = string
type CommandsGetFinishedAtRetType = string

/*
	types and functions for id
*/

// isInteger
type CommandsGetIdAttributeType = *int64
type CommandsGetIdArgType = int64
type CommandsGetIdRetType = int64

func getCommandsGetIdAttributeTypeOk(arg CommandsGetIdAttributeType) (ret CommandsGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCommandsGetIdAttributeType(arg *CommandsGetIdAttributeType, val CommandsGetIdRetType) {
	*arg = &val
}

/*
	types and functions for startedAt
*/

// isNotNullableString
type CommandsGetStartedAtAttributeType = *string

func getCommandsGetStartedAtAttributeTypeOk(arg CommandsGetStartedAtAttributeType) (ret CommandsGetStartedAtRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCommandsGetStartedAtAttributeType(arg *CommandsGetStartedAtAttributeType, val CommandsGetStartedAtRetType) {
	*arg = &val
}

type CommandsGetStartedAtArgType = string
type CommandsGetStartedAtRetType = string

/*
	types and functions for status
*/

// isEnumRef
type CommandsGetStatusAttributeType = *string
type CommandsGetStatusArgType = string
type CommandsGetStatusRetType = string

func getCommandsGetStatusAttributeTypeOk(arg CommandsGetStatusAttributeType) (ret CommandsGetStatusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCommandsGetStatusAttributeType(arg *CommandsGetStatusAttributeType, val CommandsGetStatusRetType) {
	*arg = &val
}

// Commands struct for Commands
type Commands struct {
	CommandTemplateName  CommandsGetCommandTemplateNameAttributeType  `json:"commandTemplateName,omitempty"`
	CommandTemplateTitle CommandsGetCommandTemplateTitleAttributeType `json:"commandTemplateTitle,omitempty"`
	FinishedAt           CommandsGetFinishedAtAttributeType           `json:"finishedAt,omitempty"`
	// Can be cast to int32 without loss of precision.
	Id        CommandsGetIdAttributeType        `json:"id,omitempty"`
	StartedAt CommandsGetStartedAtAttributeType `json:"startedAt,omitempty"`
	Status    CommandsGetStatusAttributeType    `json:"status,omitempty"`
}

// NewCommands instantiates a new Commands object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommands() *Commands {
	this := Commands{}
	return &this
}

// NewCommandsWithDefaults instantiates a new Commands object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandsWithDefaults() *Commands {
	this := Commands{}
	return &this
}

// GetCommandTemplateName returns the CommandTemplateName field value if set, zero value otherwise.
func (o *Commands) GetCommandTemplateName() (res CommandsGetCommandTemplateNameRetType) {
	res, _ = o.GetCommandTemplateNameOk()
	return
}

// GetCommandTemplateNameOk returns a tuple with the CommandTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commands) GetCommandTemplateNameOk() (ret CommandsGetCommandTemplateNameRetType, ok bool) {
	return getCommandsGetCommandTemplateNameAttributeTypeOk(o.CommandTemplateName)
}

// HasCommandTemplateName returns a boolean if a field has been set.
func (o *Commands) HasCommandTemplateName() bool {
	_, ok := o.GetCommandTemplateNameOk()
	return ok
}

// SetCommandTemplateName gets a reference to the given string and assigns it to the CommandTemplateName field.
func (o *Commands) SetCommandTemplateName(v CommandsGetCommandTemplateNameRetType) {
	setCommandsGetCommandTemplateNameAttributeType(&o.CommandTemplateName, v)
}

// GetCommandTemplateTitle returns the CommandTemplateTitle field value if set, zero value otherwise.
func (o *Commands) GetCommandTemplateTitle() (res CommandsGetCommandTemplateTitleRetType) {
	res, _ = o.GetCommandTemplateTitleOk()
	return
}

// GetCommandTemplateTitleOk returns a tuple with the CommandTemplateTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commands) GetCommandTemplateTitleOk() (ret CommandsGetCommandTemplateTitleRetType, ok bool) {
	return getCommandsGetCommandTemplateTitleAttributeTypeOk(o.CommandTemplateTitle)
}

// HasCommandTemplateTitle returns a boolean if a field has been set.
func (o *Commands) HasCommandTemplateTitle() bool {
	_, ok := o.GetCommandTemplateTitleOk()
	return ok
}

// SetCommandTemplateTitle gets a reference to the given string and assigns it to the CommandTemplateTitle field.
func (o *Commands) SetCommandTemplateTitle(v CommandsGetCommandTemplateTitleRetType) {
	setCommandsGetCommandTemplateTitleAttributeType(&o.CommandTemplateTitle, v)
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *Commands) GetFinishedAt() (res CommandsGetFinishedAtRetType) {
	res, _ = o.GetFinishedAtOk()
	return
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commands) GetFinishedAtOk() (ret CommandsGetFinishedAtRetType, ok bool) {
	return getCommandsGetFinishedAtAttributeTypeOk(o.FinishedAt)
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *Commands) HasFinishedAt() bool {
	_, ok := o.GetFinishedAtOk()
	return ok
}

// SetFinishedAt gets a reference to the given string and assigns it to the FinishedAt field.
func (o *Commands) SetFinishedAt(v CommandsGetFinishedAtRetType) {
	setCommandsGetFinishedAtAttributeType(&o.FinishedAt, v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Commands) GetId() (res CommandsGetIdRetType) {
	res, _ = o.GetIdOk()
	return
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commands) GetIdOk() (ret CommandsGetIdRetType, ok bool) {
	return getCommandsGetIdAttributeTypeOk(o.Id)
}

// HasId returns a boolean if a field has been set.
func (o *Commands) HasId() bool {
	_, ok := o.GetIdOk()
	return ok
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Commands) SetId(v CommandsGetIdRetType) {
	setCommandsGetIdAttributeType(&o.Id, v)
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *Commands) GetStartedAt() (res CommandsGetStartedAtRetType) {
	res, _ = o.GetStartedAtOk()
	return
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commands) GetStartedAtOk() (ret CommandsGetStartedAtRetType, ok bool) {
	return getCommandsGetStartedAtAttributeTypeOk(o.StartedAt)
}

// HasStartedAt returns a boolean if a field has been set.
func (o *Commands) HasStartedAt() bool {
	_, ok := o.GetStartedAtOk()
	return ok
}

// SetStartedAt gets a reference to the given string and assigns it to the StartedAt field.
func (o *Commands) SetStartedAt(v CommandsGetStartedAtRetType) {
	setCommandsGetStartedAtAttributeType(&o.StartedAt, v)
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Commands) GetStatus() (res CommandsGetStatusRetType) {
	res, _ = o.GetStatusOk()
	return
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commands) GetStatusOk() (ret CommandsGetStatusRetType, ok bool) {
	return getCommandsGetStatusAttributeTypeOk(o.Status)
}

// HasStatus returns a boolean if a field has been set.
func (o *Commands) HasStatus() bool {
	_, ok := o.GetStatusOk()
	return ok
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Commands) SetStatus(v CommandsGetStatusRetType) {
	setCommandsGetStatusAttributeType(&o.Status, v)
}

func (o Commands) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCommandsGetCommandTemplateNameAttributeTypeOk(o.CommandTemplateName); ok {
		toSerialize["CommandTemplateName"] = val
	}
	if val, ok := getCommandsGetCommandTemplateTitleAttributeTypeOk(o.CommandTemplateTitle); ok {
		toSerialize["CommandTemplateTitle"] = val
	}
	if val, ok := getCommandsGetFinishedAtAttributeTypeOk(o.FinishedAt); ok {
		toSerialize["FinishedAt"] = val
	}
	if val, ok := getCommandsGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getCommandsGetStartedAtAttributeTypeOk(o.StartedAt); ok {
		toSerialize["StartedAt"] = val
	}
	if val, ok := getCommandsGetStatusAttributeTypeOk(o.Status); ok {
		toSerialize["Status"] = val
	}
	return toSerialize, nil
}

type NullableCommands struct {
	value *Commands
	isSet bool
}

func (v NullableCommands) Get() *Commands {
	return v.value
}

func (v *NullableCommands) Set(val *Commands) {
	v.value = val
	v.isSet = true
}

func (v NullableCommands) IsSet() bool {
	return v.isSet
}

func (v *NullableCommands) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommands(val *Commands) *NullableCommands {
	return &NullableCommands{value: val, isSet: true}
}

func (v NullableCommands) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommands) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
