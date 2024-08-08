/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

import (
	"encoding/json"
)

// checks if the ProjectInstanceFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectInstanceFull{}

// ProjectInstanceFull struct for ProjectInstanceFull
type ProjectInstanceFull struct {
	Error *NullableString `json:"error,omitempty"`
	// REQUIRED
	Id *string `json:"id"`
	// REQUIRED
	Instance *string `json:"instance"`
	Name     *string `json:"name,omitempty"`
	// REQUIRED
	PlanName *string `json:"planName"`
	// REQUIRED
	ServiceName *string `json:"serviceName"`
	// REQUIRED
	Status *string `json:"status"`
}

type _ProjectInstanceFull ProjectInstanceFull

// NewProjectInstanceFull instantiates a new ProjectInstanceFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectInstanceFull(id *string, instance *string, planName *string, serviceName *string, status *string) *ProjectInstanceFull {
	this := ProjectInstanceFull{}
	this.Id = id
	this.Instance = instance
	this.PlanName = planName
	this.ServiceName = serviceName
	this.Status = status
	return &this
}

// NewProjectInstanceFullWithDefaults instantiates a new ProjectInstanceFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectInstanceFullWithDefaults() *ProjectInstanceFull {
	this := ProjectInstanceFull{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectInstanceFull) GetError() *string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret *string
		return ret
	}
	return o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectInstanceFull) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ProjectInstanceFull) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ProjectInstanceFull) SetError(v *string) {
	o.Error.Set(v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *ProjectInstanceFull) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ProjectInstanceFull) UnsetError() {
	o.Error.Unset()
}

// GetId returns the Id field value
func (o *ProjectInstanceFull) GetId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectInstanceFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *ProjectInstanceFull) SetId(v *string) {
	o.Id = v
}

// GetInstance returns the Instance field value
func (o *ProjectInstanceFull) GetInstance() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value
// and a boolean to check if the value has been set.
func (o *ProjectInstanceFull) GetInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instance, true
}

// SetInstance sets field value
func (o *ProjectInstanceFull) SetInstance(v *string) {
	o.Instance = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectInstanceFull) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInstanceFull) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectInstanceFull) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectInstanceFull) SetName(v *string) {
	o.Name = v
}

// GetPlanName returns the PlanName field value
func (o *ProjectInstanceFull) GetPlanName() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value
// and a boolean to check if the value has been set.
func (o *ProjectInstanceFull) GetPlanNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanName, true
}

// SetPlanName sets field value
func (o *ProjectInstanceFull) SetPlanName(v *string) {
	o.PlanName = v
}

// GetServiceName returns the ServiceName field value
func (o *ProjectInstanceFull) GetServiceName() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *ProjectInstanceFull) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// SetServiceName sets field value
func (o *ProjectInstanceFull) SetServiceName(v *string) {
	o.ServiceName = v
}

// GetStatus returns the Status field value
func (o *ProjectInstanceFull) GetStatus() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProjectInstanceFull) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *ProjectInstanceFull) SetStatus(v *string) {
	o.Status = v
}

func (o ProjectInstanceFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["instance"] = o.Instance
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["planName"] = o.PlanName
	toSerialize["serviceName"] = o.ServiceName
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableProjectInstanceFull struct {
	value *ProjectInstanceFull
	isSet bool
}

func (v NullableProjectInstanceFull) Get() *ProjectInstanceFull {
	return v.value
}

func (v *NullableProjectInstanceFull) Set(val *ProjectInstanceFull) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectInstanceFull) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectInstanceFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectInstanceFull(val *ProjectInstanceFull) *NullableProjectInstanceFull {
	return &NullableProjectInstanceFull{value: val, isSet: true}
}

func (v NullableProjectInstanceFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectInstanceFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
