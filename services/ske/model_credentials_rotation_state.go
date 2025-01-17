/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

import (
	"encoding/json"
	"time"
)

// checks if the CredentialsRotationState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsRotationState{}

// CredentialsRotationState struct for CredentialsRotationState
type CredentialsRotationState struct {
	// Format: `2024-02-15T11:06:29Z`
	LastCompletionTime *time.Time `json:"lastCompletionTime,omitempty"`
	// Format: `2024-02-15T11:06:29Z`
	LastInitiationTime *time.Time `json:"lastInitiationTime,omitempty"`
	// Phase of the credentials rotation. `NEVER` indicates that no credentials rotation has been performed using the new credentials rotation endpoints yet.
	Phase *string `json:"phase,omitempty"`
}

// NewCredentialsRotationState instantiates a new CredentialsRotationState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsRotationState() *CredentialsRotationState {
	this := CredentialsRotationState{}
	return &this
}

// NewCredentialsRotationStateWithDefaults instantiates a new CredentialsRotationState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsRotationStateWithDefaults() *CredentialsRotationState {
	this := CredentialsRotationState{}
	return &this
}

// GetLastCompletionTime returns the LastCompletionTime field value if set, zero value otherwise.
func (o *CredentialsRotationState) GetLastCompletionTime() *time.Time {
	if o == nil || IsNil(o.LastCompletionTime) {
		var ret *time.Time
		return ret
	}
	return o.LastCompletionTime
}

// GetLastCompletionTimeOk returns a tuple with the LastCompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsRotationState) GetLastCompletionTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastCompletionTime) {
		return nil, false
	}
	return o.LastCompletionTime, true
}

// HasLastCompletionTime returns a boolean if a field has been set.
func (o *CredentialsRotationState) HasLastCompletionTime() bool {
	if o != nil && !IsNil(o.LastCompletionTime) {
		return true
	}

	return false
}

// SetLastCompletionTime gets a reference to the given time.Time and assigns it to the LastCompletionTime field.
func (o *CredentialsRotationState) SetLastCompletionTime(v *time.Time) {
	o.LastCompletionTime = v
}

// GetLastInitiationTime returns the LastInitiationTime field value if set, zero value otherwise.
func (o *CredentialsRotationState) GetLastInitiationTime() *time.Time {
	if o == nil || IsNil(o.LastInitiationTime) {
		var ret *time.Time
		return ret
	}
	return o.LastInitiationTime
}

// GetLastInitiationTimeOk returns a tuple with the LastInitiationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsRotationState) GetLastInitiationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastInitiationTime) {
		return nil, false
	}
	return o.LastInitiationTime, true
}

// HasLastInitiationTime returns a boolean if a field has been set.
func (o *CredentialsRotationState) HasLastInitiationTime() bool {
	if o != nil && !IsNil(o.LastInitiationTime) {
		return true
	}

	return false
}

// SetLastInitiationTime gets a reference to the given time.Time and assigns it to the LastInitiationTime field.
func (o *CredentialsRotationState) SetLastInitiationTime(v *time.Time) {
	o.LastInitiationTime = v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *CredentialsRotationState) GetPhase() *string {
	if o == nil || IsNil(o.Phase) {
		var ret *string
		return ret
	}
	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsRotationState) GetPhaseOk() (*string, bool) {
	if o == nil || IsNil(o.Phase) {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *CredentialsRotationState) HasPhase() bool {
	if o != nil && !IsNil(o.Phase) {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *CredentialsRotationState) SetPhase(v *string) {
	o.Phase = v
}

func (o CredentialsRotationState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastCompletionTime) {
		toSerialize["lastCompletionTime"] = o.LastCompletionTime
	}
	if !IsNil(o.LastInitiationTime) {
		toSerialize["lastInitiationTime"] = o.LastInitiationTime
	}
	if !IsNil(o.Phase) {
		toSerialize["phase"] = o.Phase
	}
	return toSerialize, nil
}

type NullableCredentialsRotationState struct {
	value *CredentialsRotationState
	isSet bool
}

func (v NullableCredentialsRotationState) Get() *CredentialsRotationState {
	return v.value
}

func (v *NullableCredentialsRotationState) Set(val *CredentialsRotationState) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsRotationState) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsRotationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsRotationState(val *CredentialsRotationState) *NullableCredentialsRotationState {
	return &NullableCredentialsRotationState{value: val, isSet: true}
}

func (v NullableCredentialsRotationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsRotationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
