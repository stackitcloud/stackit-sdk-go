/*
STACKIT Server Backup Management API

API endpoints for Server Backup Operations on STACKIT Servers.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverbackup

import (
	"encoding/json"
)

// checks if the BackupPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupPolicy{}

// BackupPolicy struct for BackupPolicy
type BackupPolicy struct {
	BackupProperties *BackupPolicyBackupProperties `json:"backupProperties,omitempty"`
	Default          *bool                         `json:"default,omitempty"`
	Description      *string                       `json:"description,omitempty"`
	Enabled          *bool                         `json:"enabled,omitempty"`
	Id               *string                       `json:"id,omitempty"`
	Name             *string                       `json:"name,omitempty"`
	Rrule            *string                       `json:"rrule,omitempty"`
}

// NewBackupPolicy instantiates a new BackupPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupPolicy() *BackupPolicy {
	this := BackupPolicy{}
	return &this
}

// NewBackupPolicyWithDefaults instantiates a new BackupPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupPolicyWithDefaults() *BackupPolicy {
	this := BackupPolicy{}
	return &this
}

// GetBackupProperties returns the BackupProperties field value if set, zero value otherwise.
func (o *BackupPolicy) GetBackupProperties() *BackupPolicyBackupProperties {
	if o == nil || IsNil(o.BackupProperties) {
		var ret *BackupPolicyBackupProperties
		return ret
	}
	return o.BackupProperties
}

// GetBackupPropertiesOk returns a tuple with the BackupProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetBackupPropertiesOk() (*BackupPolicyBackupProperties, bool) {
	if o == nil || IsNil(o.BackupProperties) {
		return nil, false
	}
	return o.BackupProperties, true
}

// HasBackupProperties returns a boolean if a field has been set.
func (o *BackupPolicy) HasBackupProperties() bool {
	if o != nil && !IsNil(o.BackupProperties) && !IsNil(o.BackupProperties) {
		return true
	}

	return false
}

// SetBackupProperties gets a reference to the given BackupPolicyBackupProperties and assigns it to the BackupProperties field.
func (o *BackupPolicy) SetBackupProperties(v *BackupPolicyBackupProperties) {
	o.BackupProperties = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *BackupPolicy) GetDefault() *bool {
	if o == nil || IsNil(o.Default) {
		var ret *bool
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *BackupPolicy) HasDefault() bool {
	if o != nil && !IsNil(o.Default) && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *BackupPolicy) SetDefault(v *bool) {
	o.Default = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BackupPolicy) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BackupPolicy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BackupPolicy) SetDescription(v *string) {
	o.Description = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BackupPolicy) GetEnabled() *bool {
	if o == nil || IsNil(o.Enabled) {
		var ret *bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BackupPolicy) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BackupPolicy) SetEnabled(v *bool) {
	o.Enabled = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BackupPolicy) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BackupPolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BackupPolicy) SetId(v *string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BackupPolicy) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BackupPolicy) HasName() bool {
	if o != nil && !IsNil(o.Name) && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BackupPolicy) SetName(v *string) {
	o.Name = v
}

// GetRrule returns the Rrule field value if set, zero value otherwise.
func (o *BackupPolicy) GetRrule() *string {
	if o == nil || IsNil(o.Rrule) {
		var ret *string
		return ret
	}
	return o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetRruleOk() (*string, bool) {
	if o == nil || IsNil(o.Rrule) {
		return nil, false
	}
	return o.Rrule, true
}

// HasRrule returns a boolean if a field has been set.
func (o *BackupPolicy) HasRrule() bool {
	if o != nil && !IsNil(o.Rrule) && !IsNil(o.Rrule) {
		return true
	}

	return false
}

// SetRrule gets a reference to the given string and assigns it to the Rrule field.
func (o *BackupPolicy) SetRrule(v *string) {
	o.Rrule = v
}

func (o BackupPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackupProperties) {
		toSerialize["backupProperties"] = o.BackupProperties
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Rrule) {
		toSerialize["rrule"] = o.Rrule
	}
	return toSerialize, nil
}

type NullableBackupPolicy struct {
	value *BackupPolicy
	isSet bool
}

func (v NullableBackupPolicy) Get() *BackupPolicy {
	return v.value
}

func (v *NullableBackupPolicy) Set(val *BackupPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupPolicy(val *BackupPolicy) *NullableBackupPolicy {
	return &NullableBackupPolicy{value: val, isSet: true}
}

func (v NullableBackupPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
