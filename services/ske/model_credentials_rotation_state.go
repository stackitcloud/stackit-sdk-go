/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the CredentialsRotationState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsRotationState{}

/*
	types and functions for lastCompletionTime
*/

// isDateTime
type CredentialsRotationStateGetLastCompletionTimeAttributeType = *time.Time
type CredentialsRotationStateGetLastCompletionTimeArgType = time.Time
type CredentialsRotationStateGetLastCompletionTimeRetType = time.Time

func getCredentialsRotationStateGetLastCompletionTimeAttributeTypeOk(arg CredentialsRotationStateGetLastCompletionTimeAttributeType) (ret CredentialsRotationStateGetLastCompletionTimeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsRotationStateGetLastCompletionTimeAttributeType(arg *CredentialsRotationStateGetLastCompletionTimeAttributeType, val CredentialsRotationStateGetLastCompletionTimeRetType) {
	*arg = &val
}

/*
	types and functions for lastInitiationTime
*/

// isDateTime
type CredentialsRotationStateGetLastInitiationTimeAttributeType = *time.Time
type CredentialsRotationStateGetLastInitiationTimeArgType = time.Time
type CredentialsRotationStateGetLastInitiationTimeRetType = time.Time

func getCredentialsRotationStateGetLastInitiationTimeAttributeTypeOk(arg CredentialsRotationStateGetLastInitiationTimeAttributeType) (ret CredentialsRotationStateGetLastInitiationTimeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsRotationStateGetLastInitiationTimeAttributeType(arg *CredentialsRotationStateGetLastInitiationTimeAttributeType, val CredentialsRotationStateGetLastInitiationTimeRetType) {
	*arg = &val
}

/*
	types and functions for phase
*/

// isEnum

// CredentialsRotationStatePhase Phase of the credentials rotation. `NEVER` indicates that no credentials rotation has been performed using the new credentials rotation endpoints yet. Using the deprecated [rotate-credentials](#tag/Credentials/operation/SkeService_GetClusterCredentials) endpoint will not update this status field.
// value type for enums
type CredentialsRotationStatePhase string

// List of Phase
const (
	CREDENTIALSROTATIONSTATEPHASE_NEVER      CredentialsRotationStatePhase = "NEVER"
	CREDENTIALSROTATIONSTATEPHASE_PREPARING  CredentialsRotationStatePhase = "PREPARING"
	CREDENTIALSROTATIONSTATEPHASE_PREPARED   CredentialsRotationStatePhase = "PREPARED"
	CREDENTIALSROTATIONSTATEPHASE_COMPLETING CredentialsRotationStatePhase = "COMPLETING"
	CREDENTIALSROTATIONSTATEPHASE_COMPLETED  CredentialsRotationStatePhase = "COMPLETED"
)

// All allowed values of CredentialsRotationState enum
var AllowedCredentialsRotationStatePhaseEnumValues = []CredentialsRotationStatePhase{
	"NEVER",
	"PREPARING",
	"PREPARED",
	"COMPLETING",
	"COMPLETED",
}

func (v *CredentialsRotationStatePhase) UnmarshalJSON(src []byte) error {
	// use a type alias to prevent infinite recursion during unmarshal,
	// see https://biscuit.ninja/posts/go-avoid-an-infitine-loop-with-custom-json-unmarshallers
	type TmpJson CredentialsRotationStatePhase
	var value TmpJson
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// Allow unmarshalling zero value for testing purposes
	var zeroValue TmpJson
	if value == zeroValue {
		return nil
	}
	enumTypeValue := CredentialsRotationStatePhase(value)
	for _, existing := range AllowedCredentialsRotationStatePhaseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CredentialsRotationState", value)
}

// NewCredentialsRotationStatePhaseFromValue returns a pointer to a valid CredentialsRotationStatePhase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCredentialsRotationStatePhaseFromValue(v CredentialsRotationStatePhase) (*CredentialsRotationStatePhase, error) {
	ev := CredentialsRotationStatePhase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CredentialsRotationStatePhase: valid values are %v", v, AllowedCredentialsRotationStatePhaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CredentialsRotationStatePhase) IsValid() bool {
	for _, existing := range AllowedCredentialsRotationStatePhaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PhasePhase value
func (v CredentialsRotationStatePhase) Ptr() *CredentialsRotationStatePhase {
	return &v
}

type NullableCredentialsRotationStatePhase struct {
	value *CredentialsRotationStatePhase
	isSet bool
}

func (v NullableCredentialsRotationStatePhase) Get() *CredentialsRotationStatePhase {
	return v.value
}

func (v *NullableCredentialsRotationStatePhase) Set(val *CredentialsRotationStatePhase) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsRotationStatePhase) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsRotationStatePhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsRotationStatePhase(val *CredentialsRotationStatePhase) *NullableCredentialsRotationStatePhase {
	return &NullableCredentialsRotationStatePhase{value: val, isSet: true}
}

func (v NullableCredentialsRotationStatePhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsRotationStatePhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type CredentialsRotationStateGetPhaseAttributeType = *CredentialsRotationStatePhase
type CredentialsRotationStateGetPhaseArgType = CredentialsRotationStatePhase
type CredentialsRotationStateGetPhaseRetType = CredentialsRotationStatePhase

func getCredentialsRotationStateGetPhaseAttributeTypeOk(arg CredentialsRotationStateGetPhaseAttributeType) (ret CredentialsRotationStateGetPhaseRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsRotationStateGetPhaseAttributeType(arg *CredentialsRotationStateGetPhaseAttributeType, val CredentialsRotationStateGetPhaseRetType) {
	*arg = &val
}

// CredentialsRotationState struct for CredentialsRotationState
type CredentialsRotationState struct {
	// Format: `2024-02-15T11:06:29Z`
	LastCompletionTime CredentialsRotationStateGetLastCompletionTimeAttributeType `json:"lastCompletionTime,omitempty"`
	// Format: `2024-02-15T11:06:29Z`
	LastInitiationTime CredentialsRotationStateGetLastInitiationTimeAttributeType `json:"lastInitiationTime,omitempty"`
	// Phase of the credentials rotation. `NEVER` indicates that no credentials rotation has been performed using the new credentials rotation endpoints yet. Using the deprecated [rotate-credentials](#tag/Credentials/operation/SkeService_GetClusterCredentials) endpoint will not update this status field.
	Phase CredentialsRotationStateGetPhaseAttributeType `json:"phase,omitempty"`
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
func (o *CredentialsRotationState) GetLastCompletionTime() (res CredentialsRotationStateGetLastCompletionTimeRetType) {
	res, _ = o.GetLastCompletionTimeOk()
	return
}

// GetLastCompletionTimeOk returns a tuple with the LastCompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsRotationState) GetLastCompletionTimeOk() (ret CredentialsRotationStateGetLastCompletionTimeRetType, ok bool) {
	return getCredentialsRotationStateGetLastCompletionTimeAttributeTypeOk(o.LastCompletionTime)
}

// HasLastCompletionTime returns a boolean if a field has been set.
func (o *CredentialsRotationState) HasLastCompletionTime() bool {
	_, ok := o.GetLastCompletionTimeOk()
	return ok
}

// SetLastCompletionTime gets a reference to the given time.Time and assigns it to the LastCompletionTime field.
func (o *CredentialsRotationState) SetLastCompletionTime(v CredentialsRotationStateGetLastCompletionTimeRetType) {
	setCredentialsRotationStateGetLastCompletionTimeAttributeType(&o.LastCompletionTime, v)
}

// GetLastInitiationTime returns the LastInitiationTime field value if set, zero value otherwise.
func (o *CredentialsRotationState) GetLastInitiationTime() (res CredentialsRotationStateGetLastInitiationTimeRetType) {
	res, _ = o.GetLastInitiationTimeOk()
	return
}

// GetLastInitiationTimeOk returns a tuple with the LastInitiationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsRotationState) GetLastInitiationTimeOk() (ret CredentialsRotationStateGetLastInitiationTimeRetType, ok bool) {
	return getCredentialsRotationStateGetLastInitiationTimeAttributeTypeOk(o.LastInitiationTime)
}

// HasLastInitiationTime returns a boolean if a field has been set.
func (o *CredentialsRotationState) HasLastInitiationTime() bool {
	_, ok := o.GetLastInitiationTimeOk()
	return ok
}

// SetLastInitiationTime gets a reference to the given time.Time and assigns it to the LastInitiationTime field.
func (o *CredentialsRotationState) SetLastInitiationTime(v CredentialsRotationStateGetLastInitiationTimeRetType) {
	setCredentialsRotationStateGetLastInitiationTimeAttributeType(&o.LastInitiationTime, v)
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *CredentialsRotationState) GetPhase() (res CredentialsRotationStateGetPhaseRetType) {
	res, _ = o.GetPhaseOk()
	return
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsRotationState) GetPhaseOk() (ret CredentialsRotationStateGetPhaseRetType, ok bool) {
	return getCredentialsRotationStateGetPhaseAttributeTypeOk(o.Phase)
}

// HasPhase returns a boolean if a field has been set.
func (o *CredentialsRotationState) HasPhase() bool {
	_, ok := o.GetPhaseOk()
	return ok
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *CredentialsRotationState) SetPhase(v CredentialsRotationStateGetPhaseRetType) {
	setCredentialsRotationStateGetPhaseAttributeType(&o.Phase, v)
}

func (o CredentialsRotationState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCredentialsRotationStateGetLastCompletionTimeAttributeTypeOk(o.LastCompletionTime); ok {
		toSerialize["LastCompletionTime"] = val
	}
	if val, ok := getCredentialsRotationStateGetLastInitiationTimeAttributeTypeOk(o.LastInitiationTime); ok {
		toSerialize["LastInitiationTime"] = val
	}
	if val, ok := getCredentialsRotationStateGetPhaseAttributeTypeOk(o.Phase); ok {
		toSerialize["Phase"] = val
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
