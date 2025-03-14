/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"encoding/json"
)

// checks if the ProjectInstanceFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectInstanceFull{}

/*
	types and functions for error
*/

// isNullableString
type ProjectInstanceFullGetErrorAttributeType = *NullableString

func getProjectInstanceFullGetErrorAttributeTypeOk(arg ProjectInstanceFullGetErrorAttributeType) (ret ProjectInstanceFullGetErrorRetType, ok bool) {
	if arg == nil {
		return nil, false
	}
	return arg.Get(), true
}

func setProjectInstanceFullGetErrorAttributeType(arg *ProjectInstanceFullGetErrorAttributeType, val ProjectInstanceFullGetErrorRetType) {
	if IsNil(*arg) {
		*arg = NewNullableString(val)
	} else {
		(*arg).Set(val)
	}
}

type ProjectInstanceFullGetErrorArgType = *string
type ProjectInstanceFullGetErrorRetType = *string

/*
	types and functions for id
*/

// isNotNullableString
type ProjectInstanceFullGetIdAttributeType = *string

func getProjectInstanceFullGetIdAttributeTypeOk(arg ProjectInstanceFullGetIdAttributeType) (ret ProjectInstanceFullGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setProjectInstanceFullGetIdAttributeType(arg *ProjectInstanceFullGetIdAttributeType, val ProjectInstanceFullGetIdRetType) {
	*arg = &val
}

type ProjectInstanceFullGetIdArgType = string
type ProjectInstanceFullGetIdRetType = string

/*
	types and functions for instance
*/

// isNotNullableString
type ProjectInstanceFullGetInstanceAttributeType = *string

func getProjectInstanceFullGetInstanceAttributeTypeOk(arg ProjectInstanceFullGetInstanceAttributeType) (ret ProjectInstanceFullGetInstanceRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setProjectInstanceFullGetInstanceAttributeType(arg *ProjectInstanceFullGetInstanceAttributeType, val ProjectInstanceFullGetInstanceRetType) {
	*arg = &val
}

type ProjectInstanceFullGetInstanceArgType = string
type ProjectInstanceFullGetInstanceRetType = string

/*
	types and functions for name
*/

// isNotNullableString
type ProjectInstanceFullGetNameAttributeType = *string

func getProjectInstanceFullGetNameAttributeTypeOk(arg ProjectInstanceFullGetNameAttributeType) (ret ProjectInstanceFullGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setProjectInstanceFullGetNameAttributeType(arg *ProjectInstanceFullGetNameAttributeType, val ProjectInstanceFullGetNameRetType) {
	*arg = &val
}

type ProjectInstanceFullGetNameArgType = string
type ProjectInstanceFullGetNameRetType = string

/*
	types and functions for planName
*/

// isNotNullableString
type ProjectInstanceFullGetPlanNameAttributeType = *string

func getProjectInstanceFullGetPlanNameAttributeTypeOk(arg ProjectInstanceFullGetPlanNameAttributeType) (ret ProjectInstanceFullGetPlanNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setProjectInstanceFullGetPlanNameAttributeType(arg *ProjectInstanceFullGetPlanNameAttributeType, val ProjectInstanceFullGetPlanNameRetType) {
	*arg = &val
}

type ProjectInstanceFullGetPlanNameArgType = string
type ProjectInstanceFullGetPlanNameRetType = string

/*
	types and functions for serviceName
*/

// isNotNullableString
type ProjectInstanceFullGetServiceNameAttributeType = *string

func getProjectInstanceFullGetServiceNameAttributeTypeOk(arg ProjectInstanceFullGetServiceNameAttributeType) (ret ProjectInstanceFullGetServiceNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setProjectInstanceFullGetServiceNameAttributeType(arg *ProjectInstanceFullGetServiceNameAttributeType, val ProjectInstanceFullGetServiceNameRetType) {
	*arg = &val
}

type ProjectInstanceFullGetServiceNameArgType = string
type ProjectInstanceFullGetServiceNameRetType = string

/*
	types and functions for status
*/

// isEnumRef
type ProjectInstanceFullGetStatusAttributeType = *string
type ProjectInstanceFullGetStatusArgType = string
type ProjectInstanceFullGetStatusRetType = string

func getProjectInstanceFullGetStatusAttributeTypeOk(arg ProjectInstanceFullGetStatusAttributeType) (ret ProjectInstanceFullGetStatusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setProjectInstanceFullGetStatusAttributeType(arg *ProjectInstanceFullGetStatusAttributeType, val ProjectInstanceFullGetStatusRetType) {
	*arg = &val
}

// ProjectInstanceFull struct for ProjectInstanceFull
type ProjectInstanceFull struct {
	Error ProjectInstanceFullGetErrorAttributeType `json:"error,omitempty"`
	// REQUIRED
	Id ProjectInstanceFullGetIdAttributeType `json:"id"`
	// REQUIRED
	Instance ProjectInstanceFullGetInstanceAttributeType `json:"instance"`
	Name     ProjectInstanceFullGetNameAttributeType     `json:"name,omitempty"`
	// REQUIRED
	PlanName ProjectInstanceFullGetPlanNameAttributeType `json:"planName"`
	// REQUIRED
	ServiceName ProjectInstanceFullGetServiceNameAttributeType `json:"serviceName"`
	// REQUIRED
	Status ProjectInstanceFullGetStatusAttributeType `json:"status"`
}

type _ProjectInstanceFull ProjectInstanceFull

// NewProjectInstanceFull instantiates a new ProjectInstanceFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectInstanceFull(id ProjectInstanceFullGetIdArgType, instance ProjectInstanceFullGetInstanceArgType, planName ProjectInstanceFullGetPlanNameArgType, serviceName ProjectInstanceFullGetServiceNameArgType, status ProjectInstanceFullGetStatusArgType) *ProjectInstanceFull {
	this := ProjectInstanceFull{}
	setProjectInstanceFullGetIdAttributeType(&this.Id, id)
	setProjectInstanceFullGetInstanceAttributeType(&this.Instance, instance)
	setProjectInstanceFullGetPlanNameAttributeType(&this.PlanName, planName)
	setProjectInstanceFullGetServiceNameAttributeType(&this.ServiceName, serviceName)
	setProjectInstanceFullGetStatusAttributeType(&this.Status, status)
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
func (o *ProjectInstanceFull) GetError() (res ProjectInstanceFullGetErrorRetType) {
	res, _ = o.GetErrorOk()
	return
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectInstanceFull) GetErrorOk() (ret ProjectInstanceFullGetErrorRetType, ok bool) {
	return getProjectInstanceFullGetErrorAttributeTypeOk(o.Error)
}

// HasError returns a boolean if a field has been set.
func (o *ProjectInstanceFull) HasError() bool {
	_, ok := o.GetErrorOk()
	return ok
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ProjectInstanceFull) SetError(v ProjectInstanceFullGetErrorRetType) {
	setProjectInstanceFullGetErrorAttributeType(&o.Error, v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *ProjectInstanceFull) SetErrorNil() {
	o.Error = nil
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ProjectInstanceFull) UnsetError() {
	o.Error = nil
}

// GetId returns the Id field value
func (o *ProjectInstanceFull) GetId() (ret ProjectInstanceFullGetIdRetType) {
	ret, _ = o.GetIdOk()
	return ret
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectInstanceFull) GetIdOk() (ret ProjectInstanceFullGetIdRetType, ok bool) {
	return getProjectInstanceFullGetIdAttributeTypeOk(o.Id)
}

// SetId sets field value
func (o *ProjectInstanceFull) SetId(v ProjectInstanceFullGetIdRetType) {
	setProjectInstanceFullGetIdAttributeType(&o.Id, v)
}

// GetInstance returns the Instance field value
func (o *ProjectInstanceFull) GetInstance() (ret ProjectInstanceFullGetInstanceRetType) {
	ret, _ = o.GetInstanceOk()
	return ret
}

// GetInstanceOk returns a tuple with the Instance field value
// and a boolean to check if the value has been set.
func (o *ProjectInstanceFull) GetInstanceOk() (ret ProjectInstanceFullGetInstanceRetType, ok bool) {
	return getProjectInstanceFullGetInstanceAttributeTypeOk(o.Instance)
}

// SetInstance sets field value
func (o *ProjectInstanceFull) SetInstance(v ProjectInstanceFullGetInstanceRetType) {
	setProjectInstanceFullGetInstanceAttributeType(&o.Instance, v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectInstanceFull) GetName() (res ProjectInstanceFullGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInstanceFull) GetNameOk() (ret ProjectInstanceFullGetNameRetType, ok bool) {
	return getProjectInstanceFullGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *ProjectInstanceFull) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectInstanceFull) SetName(v ProjectInstanceFullGetNameRetType) {
	setProjectInstanceFullGetNameAttributeType(&o.Name, v)
}

// GetPlanName returns the PlanName field value
func (o *ProjectInstanceFull) GetPlanName() (ret ProjectInstanceFullGetPlanNameRetType) {
	ret, _ = o.GetPlanNameOk()
	return ret
}

// GetPlanNameOk returns a tuple with the PlanName field value
// and a boolean to check if the value has been set.
func (o *ProjectInstanceFull) GetPlanNameOk() (ret ProjectInstanceFullGetPlanNameRetType, ok bool) {
	return getProjectInstanceFullGetPlanNameAttributeTypeOk(o.PlanName)
}

// SetPlanName sets field value
func (o *ProjectInstanceFull) SetPlanName(v ProjectInstanceFullGetPlanNameRetType) {
	setProjectInstanceFullGetPlanNameAttributeType(&o.PlanName, v)
}

// GetServiceName returns the ServiceName field value
func (o *ProjectInstanceFull) GetServiceName() (ret ProjectInstanceFullGetServiceNameRetType) {
	ret, _ = o.GetServiceNameOk()
	return ret
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *ProjectInstanceFull) GetServiceNameOk() (ret ProjectInstanceFullGetServiceNameRetType, ok bool) {
	return getProjectInstanceFullGetServiceNameAttributeTypeOk(o.ServiceName)
}

// SetServiceName sets field value
func (o *ProjectInstanceFull) SetServiceName(v ProjectInstanceFullGetServiceNameRetType) {
	setProjectInstanceFullGetServiceNameAttributeType(&o.ServiceName, v)
}

// GetStatus returns the Status field value
func (o *ProjectInstanceFull) GetStatus() (ret ProjectInstanceFullGetStatusRetType) {
	ret, _ = o.GetStatusOk()
	return ret
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProjectInstanceFull) GetStatusOk() (ret ProjectInstanceFullGetStatusRetType, ok bool) {
	return getProjectInstanceFullGetStatusAttributeTypeOk(o.Status)
}

// SetStatus sets field value
func (o *ProjectInstanceFull) SetStatus(v ProjectInstanceFullGetStatusRetType) {
	setProjectInstanceFullGetStatusAttributeType(&o.Status, v)
}

func (o ProjectInstanceFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getProjectInstanceFullGetErrorAttributeTypeOk(o.Error); ok {
		toSerialize["Error"] = val
	}
	if val, ok := getProjectInstanceFullGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getProjectInstanceFullGetInstanceAttributeTypeOk(o.Instance); ok {
		toSerialize["Instance"] = val
	}
	if val, ok := getProjectInstanceFullGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getProjectInstanceFullGetPlanNameAttributeTypeOk(o.PlanName); ok {
		toSerialize["PlanName"] = val
	}
	if val, ok := getProjectInstanceFullGetServiceNameAttributeTypeOk(o.ServiceName); ok {
		toSerialize["ServiceName"] = val
	}
	if val, ok := getProjectInstanceFullGetStatusAttributeTypeOk(o.Status); ok {
		toSerialize["Status"] = val
	}
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
