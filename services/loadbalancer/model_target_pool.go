/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.7.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

import (
	"encoding/json"
)

// checks if the TargetPool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetPool{}

// TargetPool struct for TargetPool
type TargetPool struct {
	ActiveHealthCheck *ActiveHealthCheck `json:"activeHealthCheck,omitempty"`
	// Target pool name
	Name               *string             `json:"name,omitempty"`
	SessionPersistence *SessionPersistence `json:"sessionPersistence,omitempty"`
	// The number identifying the port where each target listens for traffic.
	TargetPort *int64 `json:"targetPort,omitempty"`
	// List of all targets which will be used in the pool. Limited to 250.
	Targets *[]Target `json:"targets,omitempty"`
}

// NewTargetPool instantiates a new TargetPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetPool() *TargetPool {
	this := TargetPool{}
	return &this
}

// NewTargetPoolWithDefaults instantiates a new TargetPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetPoolWithDefaults() *TargetPool {
	this := TargetPool{}
	return &this
}

// GetActiveHealthCheck returns the ActiveHealthCheck field value if set, zero value otherwise.
func (o *TargetPool) GetActiveHealthCheck() *ActiveHealthCheck {
	if o == nil || IsNil(o.ActiveHealthCheck) {
		var ret *ActiveHealthCheck
		return ret
	}
	return o.ActiveHealthCheck
}

// GetActiveHealthCheckOk returns a tuple with the ActiveHealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPool) GetActiveHealthCheckOk() (*ActiveHealthCheck, bool) {
	if o == nil || IsNil(o.ActiveHealthCheck) {
		return nil, false
	}
	return o.ActiveHealthCheck, true
}

// HasActiveHealthCheck returns a boolean if a field has been set.
func (o *TargetPool) HasActiveHealthCheck() bool {
	if o != nil && !IsNil(o.ActiveHealthCheck) {
		return true
	}

	return false
}

// SetActiveHealthCheck gets a reference to the given ActiveHealthCheck and assigns it to the ActiveHealthCheck field.
func (o *TargetPool) SetActiveHealthCheck(v *ActiveHealthCheck) {
	o.ActiveHealthCheck = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TargetPool) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPool) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TargetPool) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TargetPool) SetName(v *string) {
	o.Name = v
}

// GetSessionPersistence returns the SessionPersistence field value if set, zero value otherwise.
func (o *TargetPool) GetSessionPersistence() *SessionPersistence {
	if o == nil || IsNil(o.SessionPersistence) {
		var ret *SessionPersistence
		return ret
	}
	return o.SessionPersistence
}

// GetSessionPersistenceOk returns a tuple with the SessionPersistence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPool) GetSessionPersistenceOk() (*SessionPersistence, bool) {
	if o == nil || IsNil(o.SessionPersistence) {
		return nil, false
	}
	return o.SessionPersistence, true
}

// HasSessionPersistence returns a boolean if a field has been set.
func (o *TargetPool) HasSessionPersistence() bool {
	if o != nil && !IsNil(o.SessionPersistence) {
		return true
	}

	return false
}

// SetSessionPersistence gets a reference to the given SessionPersistence and assigns it to the SessionPersistence field.
func (o *TargetPool) SetSessionPersistence(v *SessionPersistence) {
	o.SessionPersistence = v
}

// GetTargetPort returns the TargetPort field value if set, zero value otherwise.
func (o *TargetPool) GetTargetPort() *int64 {
	if o == nil || IsNil(o.TargetPort) {
		var ret *int64
		return ret
	}
	return o.TargetPort
}

// GetTargetPortOk returns a tuple with the TargetPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPool) GetTargetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.TargetPort) {
		return nil, false
	}
	return o.TargetPort, true
}

// HasTargetPort returns a boolean if a field has been set.
func (o *TargetPool) HasTargetPort() bool {
	if o != nil && !IsNil(o.TargetPort) {
		return true
	}

	return false
}

// SetTargetPort gets a reference to the given int64 and assigns it to the TargetPort field.
func (o *TargetPool) SetTargetPort(v *int64) {
	o.TargetPort = v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *TargetPool) GetTargets() *[]Target {
	if o == nil || IsNil(o.Targets) {
		var ret *[]Target
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetPool) GetTargetsOk() (*[]Target, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *TargetPool) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []Target and assigns it to the Targets field.
func (o *TargetPool) SetTargets(v *[]Target) {
	o.Targets = v
}

func (o TargetPool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveHealthCheck) {
		toSerialize["activeHealthCheck"] = o.ActiveHealthCheck
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SessionPersistence) {
		toSerialize["sessionPersistence"] = o.SessionPersistence
	}
	if !IsNil(o.TargetPort) {
		toSerialize["targetPort"] = o.TargetPort
	}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	return toSerialize, nil
}

type NullableTargetPool struct {
	value *TargetPool
	isSet bool
}

func (v NullableTargetPool) Get() *TargetPool {
	return v.value
}

func (v *NullableTargetPool) Set(val *TargetPool) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetPool) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetPool(val *TargetPool) *NullableTargetPool {
	return &NullableTargetPool{value: val, isSet: true}
}

func (v NullableTargetPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
