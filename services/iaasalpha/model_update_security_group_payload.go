/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
)

// checks if the UpdateSecurityGroupPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSecurityGroupPayload{}

/*
	types and functions for description
*/

// isNotNullableString
type UpdateSecurityGroupPayloadGetDescriptionAttributeType = *string

func getUpdateSecurityGroupPayloadGetDescriptionAttributeTypeOk(arg UpdateSecurityGroupPayloadGetDescriptionAttributeType) (ret UpdateSecurityGroupPayloadGetDescriptionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateSecurityGroupPayloadGetDescriptionAttributeType(arg *UpdateSecurityGroupPayloadGetDescriptionAttributeType, val UpdateSecurityGroupPayloadGetDescriptionRetType) {
	*arg = &val
}

type UpdateSecurityGroupPayloadGetDescriptionArgType = string
type UpdateSecurityGroupPayloadGetDescriptionRetType = string

/*
	types and functions for labels
*/

// isFreeform
type UpdateSecurityGroupPayloadGetLabelsAttributeType = *map[string]interface{}
type UpdateSecurityGroupPayloadGetLabelsArgType = map[string]interface{}
type UpdateSecurityGroupPayloadGetLabelsRetType = map[string]interface{}

func getUpdateSecurityGroupPayloadGetLabelsAttributeTypeOk(arg UpdateSecurityGroupPayloadGetLabelsAttributeType) (ret UpdateSecurityGroupPayloadGetLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateSecurityGroupPayloadGetLabelsAttributeType(arg *UpdateSecurityGroupPayloadGetLabelsAttributeType, val UpdateSecurityGroupPayloadGetLabelsRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type UpdateSecurityGroupPayloadGetNameAttributeType = *string

func getUpdateSecurityGroupPayloadGetNameAttributeTypeOk(arg UpdateSecurityGroupPayloadGetNameAttributeType) (ret UpdateSecurityGroupPayloadGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateSecurityGroupPayloadGetNameAttributeType(arg *UpdateSecurityGroupPayloadGetNameAttributeType, val UpdateSecurityGroupPayloadGetNameRetType) {
	*arg = &val
}

type UpdateSecurityGroupPayloadGetNameArgType = string
type UpdateSecurityGroupPayloadGetNameRetType = string

// UpdateSecurityGroupPayload Object that represents an update request body of a security group.
type UpdateSecurityGroupPayload struct {
	// Description Object. Allows string up to 127 Characters.
	Description UpdateSecurityGroupPayloadGetDescriptionAttributeType `json:"description,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels UpdateSecurityGroupPayloadGetLabelsAttributeType `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name UpdateSecurityGroupPayloadGetNameAttributeType `json:"name,omitempty"`
}

// NewUpdateSecurityGroupPayload instantiates a new UpdateSecurityGroupPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSecurityGroupPayload() *UpdateSecurityGroupPayload {
	this := UpdateSecurityGroupPayload{}
	return &this
}

// NewUpdateSecurityGroupPayloadWithDefaults instantiates a new UpdateSecurityGroupPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSecurityGroupPayloadWithDefaults() *UpdateSecurityGroupPayload {
	this := UpdateSecurityGroupPayload{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateSecurityGroupPayload) GetDescription() (res UpdateSecurityGroupPayloadGetDescriptionRetType) {
	res, _ = o.GetDescriptionOk()
	return
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSecurityGroupPayload) GetDescriptionOk() (ret UpdateSecurityGroupPayloadGetDescriptionRetType, ok bool) {
	return getUpdateSecurityGroupPayloadGetDescriptionAttributeTypeOk(o.Description)
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateSecurityGroupPayload) HasDescription() bool {
	_, ok := o.GetDescriptionOk()
	return ok
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateSecurityGroupPayload) SetDescription(v UpdateSecurityGroupPayloadGetDescriptionRetType) {
	setUpdateSecurityGroupPayloadGetDescriptionAttributeType(&o.Description, v)
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateSecurityGroupPayload) GetLabels() (res UpdateSecurityGroupPayloadGetLabelsRetType) {
	res, _ = o.GetLabelsOk()
	return
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSecurityGroupPayload) GetLabelsOk() (ret UpdateSecurityGroupPayloadGetLabelsRetType, ok bool) {
	return getUpdateSecurityGroupPayloadGetLabelsAttributeTypeOk(o.Labels)
}

// HasLabels returns a boolean if a field has been set.
func (o *UpdateSecurityGroupPayload) HasLabels() bool {
	_, ok := o.GetLabelsOk()
	return ok
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *UpdateSecurityGroupPayload) SetLabels(v UpdateSecurityGroupPayloadGetLabelsRetType) {
	setUpdateSecurityGroupPayloadGetLabelsAttributeType(&o.Labels, v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateSecurityGroupPayload) GetName() (res UpdateSecurityGroupPayloadGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSecurityGroupPayload) GetNameOk() (ret UpdateSecurityGroupPayloadGetNameRetType, ok bool) {
	return getUpdateSecurityGroupPayloadGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *UpdateSecurityGroupPayload) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateSecurityGroupPayload) SetName(v UpdateSecurityGroupPayloadGetNameRetType) {
	setUpdateSecurityGroupPayloadGetNameAttributeType(&o.Name, v)
}

func (o UpdateSecurityGroupPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getUpdateSecurityGroupPayloadGetDescriptionAttributeTypeOk(o.Description); ok {
		toSerialize["Description"] = val
	}
	if val, ok := getUpdateSecurityGroupPayloadGetLabelsAttributeTypeOk(o.Labels); ok {
		toSerialize["Labels"] = val
	}
	if val, ok := getUpdateSecurityGroupPayloadGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	return toSerialize, nil
}

type NullableUpdateSecurityGroupPayload struct {
	value *UpdateSecurityGroupPayload
	isSet bool
}

func (v NullableUpdateSecurityGroupPayload) Get() *UpdateSecurityGroupPayload {
	return v.value
}

func (v *NullableUpdateSecurityGroupPayload) Set(val *UpdateSecurityGroupPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSecurityGroupPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSecurityGroupPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSecurityGroupPayload(val *UpdateSecurityGroupPayload) *NullableUpdateSecurityGroupPayload {
	return &NullableUpdateSecurityGroupPayload{value: val, isSet: true}
}

func (v NullableUpdateSecurityGroupPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSecurityGroupPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
