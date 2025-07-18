/*
STACKIT Server Update Management API

API endpoints for Server Update Operations on STACKIT Servers.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverupdate

import (
	"encoding/json"
)

// checks if the CreateUpdateSchedulePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUpdateSchedulePayload{}

/*
	types and functions for enabled
*/

// isBoolean
type CreateUpdateSchedulePayloadgetEnabledAttributeType = *bool
type CreateUpdateSchedulePayloadgetEnabledArgType = bool
type CreateUpdateSchedulePayloadgetEnabledRetType = bool

func getCreateUpdateSchedulePayloadgetEnabledAttributeTypeOk(arg CreateUpdateSchedulePayloadgetEnabledAttributeType) (ret CreateUpdateSchedulePayloadgetEnabledRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateUpdateSchedulePayloadgetEnabledAttributeType(arg *CreateUpdateSchedulePayloadgetEnabledAttributeType, val CreateUpdateSchedulePayloadgetEnabledRetType) {
	*arg = &val
}

/*
	types and functions for maintenanceWindow
*/

// isInteger
type CreateUpdateSchedulePayloadGetMaintenanceWindowAttributeType = *int64
type CreateUpdateSchedulePayloadGetMaintenanceWindowArgType = int64
type CreateUpdateSchedulePayloadGetMaintenanceWindowRetType = int64

func getCreateUpdateSchedulePayloadGetMaintenanceWindowAttributeTypeOk(arg CreateUpdateSchedulePayloadGetMaintenanceWindowAttributeType) (ret CreateUpdateSchedulePayloadGetMaintenanceWindowRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateUpdateSchedulePayloadGetMaintenanceWindowAttributeType(arg *CreateUpdateSchedulePayloadGetMaintenanceWindowAttributeType, val CreateUpdateSchedulePayloadGetMaintenanceWindowRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type CreateUpdateSchedulePayloadGetNameAttributeType = *string

func getCreateUpdateSchedulePayloadGetNameAttributeTypeOk(arg CreateUpdateSchedulePayloadGetNameAttributeType) (ret CreateUpdateSchedulePayloadGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateUpdateSchedulePayloadGetNameAttributeType(arg *CreateUpdateSchedulePayloadGetNameAttributeType, val CreateUpdateSchedulePayloadGetNameRetType) {
	*arg = &val
}

type CreateUpdateSchedulePayloadGetNameArgType = string
type CreateUpdateSchedulePayloadGetNameRetType = string

/*
	types and functions for rrule
*/

// isNotNullableString
type CreateUpdateSchedulePayloadGetRruleAttributeType = *string

func getCreateUpdateSchedulePayloadGetRruleAttributeTypeOk(arg CreateUpdateSchedulePayloadGetRruleAttributeType) (ret CreateUpdateSchedulePayloadGetRruleRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateUpdateSchedulePayloadGetRruleAttributeType(arg *CreateUpdateSchedulePayloadGetRruleAttributeType, val CreateUpdateSchedulePayloadGetRruleRetType) {
	*arg = &val
}

type CreateUpdateSchedulePayloadGetRruleArgType = string
type CreateUpdateSchedulePayloadGetRruleRetType = string

// CreateUpdateSchedulePayload struct for CreateUpdateSchedulePayload
type CreateUpdateSchedulePayload struct {
	// REQUIRED
	Enabled CreateUpdateSchedulePayloadgetEnabledAttributeType `json:"enabled" required:"true"`
	// Can be cast to int32 without loss of precision.
	// REQUIRED
	MaintenanceWindow CreateUpdateSchedulePayloadGetMaintenanceWindowAttributeType `json:"maintenanceWindow" required:"true"`
	// REQUIRED
	Name CreateUpdateSchedulePayloadGetNameAttributeType `json:"name" required:"true"`
	// REQUIRED
	Rrule CreateUpdateSchedulePayloadGetRruleAttributeType `json:"rrule" required:"true"`
}

type _CreateUpdateSchedulePayload CreateUpdateSchedulePayload

// NewCreateUpdateSchedulePayload instantiates a new CreateUpdateSchedulePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUpdateSchedulePayload(enabled CreateUpdateSchedulePayloadgetEnabledArgType, maintenanceWindow CreateUpdateSchedulePayloadGetMaintenanceWindowArgType, name CreateUpdateSchedulePayloadGetNameArgType, rrule CreateUpdateSchedulePayloadGetRruleArgType) *CreateUpdateSchedulePayload {
	this := CreateUpdateSchedulePayload{}
	setCreateUpdateSchedulePayloadgetEnabledAttributeType(&this.Enabled, enabled)
	setCreateUpdateSchedulePayloadGetMaintenanceWindowAttributeType(&this.MaintenanceWindow, maintenanceWindow)
	setCreateUpdateSchedulePayloadGetNameAttributeType(&this.Name, name)
	setCreateUpdateSchedulePayloadGetRruleAttributeType(&this.Rrule, rrule)
	return &this
}

// NewCreateUpdateSchedulePayloadWithDefaults instantiates a new CreateUpdateSchedulePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUpdateSchedulePayloadWithDefaults() *CreateUpdateSchedulePayload {
	this := CreateUpdateSchedulePayload{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *CreateUpdateSchedulePayload) GetEnabled() (ret CreateUpdateSchedulePayloadgetEnabledRetType) {
	ret, _ = o.GetEnabledOk()
	return ret
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateUpdateSchedulePayload) GetEnabledOk() (ret CreateUpdateSchedulePayloadgetEnabledRetType, ok bool) {
	return getCreateUpdateSchedulePayloadgetEnabledAttributeTypeOk(o.Enabled)
}

// SetEnabled sets field value
func (o *CreateUpdateSchedulePayload) SetEnabled(v CreateUpdateSchedulePayloadgetEnabledRetType) {
	setCreateUpdateSchedulePayloadgetEnabledAttributeType(&o.Enabled, v)
}

// GetMaintenanceWindow returns the MaintenanceWindow field value
func (o *CreateUpdateSchedulePayload) GetMaintenanceWindow() (ret CreateUpdateSchedulePayloadGetMaintenanceWindowRetType) {
	ret, _ = o.GetMaintenanceWindowOk()
	return ret
}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value
// and a boolean to check if the value has been set.
func (o *CreateUpdateSchedulePayload) GetMaintenanceWindowOk() (ret CreateUpdateSchedulePayloadGetMaintenanceWindowRetType, ok bool) {
	return getCreateUpdateSchedulePayloadGetMaintenanceWindowAttributeTypeOk(o.MaintenanceWindow)
}

// SetMaintenanceWindow sets field value
func (o *CreateUpdateSchedulePayload) SetMaintenanceWindow(v CreateUpdateSchedulePayloadGetMaintenanceWindowRetType) {
	setCreateUpdateSchedulePayloadGetMaintenanceWindowAttributeType(&o.MaintenanceWindow, v)
}

// GetName returns the Name field value
func (o *CreateUpdateSchedulePayload) GetName() (ret CreateUpdateSchedulePayloadGetNameRetType) {
	ret, _ = o.GetNameOk()
	return ret
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateUpdateSchedulePayload) GetNameOk() (ret CreateUpdateSchedulePayloadGetNameRetType, ok bool) {
	return getCreateUpdateSchedulePayloadGetNameAttributeTypeOk(o.Name)
}

// SetName sets field value
func (o *CreateUpdateSchedulePayload) SetName(v CreateUpdateSchedulePayloadGetNameRetType) {
	setCreateUpdateSchedulePayloadGetNameAttributeType(&o.Name, v)
}

// GetRrule returns the Rrule field value
func (o *CreateUpdateSchedulePayload) GetRrule() (ret CreateUpdateSchedulePayloadGetRruleRetType) {
	ret, _ = o.GetRruleOk()
	return ret
}

// GetRruleOk returns a tuple with the Rrule field value
// and a boolean to check if the value has been set.
func (o *CreateUpdateSchedulePayload) GetRruleOk() (ret CreateUpdateSchedulePayloadGetRruleRetType, ok bool) {
	return getCreateUpdateSchedulePayloadGetRruleAttributeTypeOk(o.Rrule)
}

// SetRrule sets field value
func (o *CreateUpdateSchedulePayload) SetRrule(v CreateUpdateSchedulePayloadGetRruleRetType) {
	setCreateUpdateSchedulePayloadGetRruleAttributeType(&o.Rrule, v)
}

func (o CreateUpdateSchedulePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateUpdateSchedulePayloadgetEnabledAttributeTypeOk(o.Enabled); ok {
		toSerialize["Enabled"] = val
	}
	if val, ok := getCreateUpdateSchedulePayloadGetMaintenanceWindowAttributeTypeOk(o.MaintenanceWindow); ok {
		toSerialize["MaintenanceWindow"] = val
	}
	if val, ok := getCreateUpdateSchedulePayloadGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getCreateUpdateSchedulePayloadGetRruleAttributeTypeOk(o.Rrule); ok {
		toSerialize["Rrule"] = val
	}
	return toSerialize, nil
}

type NullableCreateUpdateSchedulePayload struct {
	value *CreateUpdateSchedulePayload
	isSet bool
}

func (v NullableCreateUpdateSchedulePayload) Get() *CreateUpdateSchedulePayload {
	return v.value
}

func (v *NullableCreateUpdateSchedulePayload) Set(val *CreateUpdateSchedulePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUpdateSchedulePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUpdateSchedulePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUpdateSchedulePayload(val *CreateUpdateSchedulePayload) *NullableCreateUpdateSchedulePayload {
	return &NullableCreateUpdateSchedulePayload{value: val, isSet: true}
}

func (v NullableCreateUpdateSchedulePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUpdateSchedulePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
