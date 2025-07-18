/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"encoding/json"
	"fmt"
)

// checks if the GetInstanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInstanceResponse{}

/*
	types and functions for dashboardUrl
*/

// isNotNullableString
type GetInstanceResponseGetDashboardUrlAttributeType = *string

func getGetInstanceResponseGetDashboardUrlAttributeTypeOk(arg GetInstanceResponseGetDashboardUrlAttributeType) (ret GetInstanceResponseGetDashboardUrlRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetInstanceResponseGetDashboardUrlAttributeType(arg *GetInstanceResponseGetDashboardUrlAttributeType, val GetInstanceResponseGetDashboardUrlRetType) {
	*arg = &val
}

type GetInstanceResponseGetDashboardUrlArgType = string
type GetInstanceResponseGetDashboardUrlRetType = string

/*
	types and functions for error
*/

// isNullableString
type GetInstanceResponseGetErrorAttributeType = *NullableString

func getGetInstanceResponseGetErrorAttributeTypeOk(arg GetInstanceResponseGetErrorAttributeType) (ret GetInstanceResponseGetErrorRetType, ok bool) {
	if arg == nil {
		return nil, false
	}
	return arg.Get(), true
}

func setGetInstanceResponseGetErrorAttributeType(arg *GetInstanceResponseGetErrorAttributeType, val GetInstanceResponseGetErrorRetType) {
	if IsNil(*arg) {
		*arg = NewNullableString(val)
	} else {
		(*arg).Set(val)
	}
}

type GetInstanceResponseGetErrorArgType = *string
type GetInstanceResponseGetErrorRetType = *string

/*
	types and functions for id
*/

// isNotNullableString
type GetInstanceResponseGetIdAttributeType = *string

func getGetInstanceResponseGetIdAttributeTypeOk(arg GetInstanceResponseGetIdAttributeType) (ret GetInstanceResponseGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetInstanceResponseGetIdAttributeType(arg *GetInstanceResponseGetIdAttributeType, val GetInstanceResponseGetIdRetType) {
	*arg = &val
}

type GetInstanceResponseGetIdArgType = string
type GetInstanceResponseGetIdRetType = string

/*
	types and functions for instance
*/

// isModel
type GetInstanceResponseGetInstanceAttributeType = *InstanceSensitiveData
type GetInstanceResponseGetInstanceArgType = InstanceSensitiveData
type GetInstanceResponseGetInstanceRetType = InstanceSensitiveData

func getGetInstanceResponseGetInstanceAttributeTypeOk(arg GetInstanceResponseGetInstanceAttributeType) (ret GetInstanceResponseGetInstanceRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetInstanceResponseGetInstanceAttributeType(arg *GetInstanceResponseGetInstanceAttributeType, val GetInstanceResponseGetInstanceRetType) {
	*arg = &val
}

/*
	types and functions for isUpdatable
*/

// isBoolean
type GetInstanceResponsegetIsUpdatableAttributeType = *bool
type GetInstanceResponsegetIsUpdatableArgType = bool
type GetInstanceResponsegetIsUpdatableRetType = bool

func getGetInstanceResponsegetIsUpdatableAttributeTypeOk(arg GetInstanceResponsegetIsUpdatableAttributeType) (ret GetInstanceResponsegetIsUpdatableRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetInstanceResponsegetIsUpdatableAttributeType(arg *GetInstanceResponsegetIsUpdatableAttributeType, val GetInstanceResponsegetIsUpdatableRetType) {
	*arg = &val
}

/*
	types and functions for message
*/

// isNotNullableString
type GetInstanceResponseGetMessageAttributeType = *string

func getGetInstanceResponseGetMessageAttributeTypeOk(arg GetInstanceResponseGetMessageAttributeType) (ret GetInstanceResponseGetMessageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetInstanceResponseGetMessageAttributeType(arg *GetInstanceResponseGetMessageAttributeType, val GetInstanceResponseGetMessageRetType) {
	*arg = &val
}

type GetInstanceResponseGetMessageArgType = string
type GetInstanceResponseGetMessageRetType = string

/*
	types and functions for name
*/

// isNotNullableString
type GetInstanceResponseGetNameAttributeType = *string

func getGetInstanceResponseGetNameAttributeTypeOk(arg GetInstanceResponseGetNameAttributeType) (ret GetInstanceResponseGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetInstanceResponseGetNameAttributeType(arg *GetInstanceResponseGetNameAttributeType, val GetInstanceResponseGetNameRetType) {
	*arg = &val
}

type GetInstanceResponseGetNameArgType = string
type GetInstanceResponseGetNameRetType = string

/*
	types and functions for parameters
*/

// isContainer
type GetInstanceResponseGetParametersAttributeType = *map[string]string
type GetInstanceResponseGetParametersArgType = map[string]string
type GetInstanceResponseGetParametersRetType = map[string]string

func getGetInstanceResponseGetParametersAttributeTypeOk(arg GetInstanceResponseGetParametersAttributeType) (ret GetInstanceResponseGetParametersRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetInstanceResponseGetParametersAttributeType(arg *GetInstanceResponseGetParametersAttributeType, val GetInstanceResponseGetParametersRetType) {
	*arg = &val
}

/*
	types and functions for planId
*/

// isNotNullableString
type GetInstanceResponseGetPlanIdAttributeType = *string

func getGetInstanceResponseGetPlanIdAttributeTypeOk(arg GetInstanceResponseGetPlanIdAttributeType) (ret GetInstanceResponseGetPlanIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetInstanceResponseGetPlanIdAttributeType(arg *GetInstanceResponseGetPlanIdAttributeType, val GetInstanceResponseGetPlanIdRetType) {
	*arg = &val
}

type GetInstanceResponseGetPlanIdArgType = string
type GetInstanceResponseGetPlanIdRetType = string

/*
	types and functions for planName
*/

// isNotNullableString
type GetInstanceResponseGetPlanNameAttributeType = *string

func getGetInstanceResponseGetPlanNameAttributeTypeOk(arg GetInstanceResponseGetPlanNameAttributeType) (ret GetInstanceResponseGetPlanNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetInstanceResponseGetPlanNameAttributeType(arg *GetInstanceResponseGetPlanNameAttributeType, val GetInstanceResponseGetPlanNameRetType) {
	*arg = &val
}

type GetInstanceResponseGetPlanNameArgType = string
type GetInstanceResponseGetPlanNameRetType = string

/*
	types and functions for planSchema
*/

// isNotNullableString
type GetInstanceResponseGetPlanSchemaAttributeType = *string

func getGetInstanceResponseGetPlanSchemaAttributeTypeOk(arg GetInstanceResponseGetPlanSchemaAttributeType) (ret GetInstanceResponseGetPlanSchemaRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetInstanceResponseGetPlanSchemaAttributeType(arg *GetInstanceResponseGetPlanSchemaAttributeType, val GetInstanceResponseGetPlanSchemaRetType) {
	*arg = &val
}

type GetInstanceResponseGetPlanSchemaArgType = string
type GetInstanceResponseGetPlanSchemaRetType = string

/*
	types and functions for serviceName
*/

// isNotNullableString
type GetInstanceResponseGetServiceNameAttributeType = *string

func getGetInstanceResponseGetServiceNameAttributeTypeOk(arg GetInstanceResponseGetServiceNameAttributeType) (ret GetInstanceResponseGetServiceNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetInstanceResponseGetServiceNameAttributeType(arg *GetInstanceResponseGetServiceNameAttributeType, val GetInstanceResponseGetServiceNameRetType) {
	*arg = &val
}

type GetInstanceResponseGetServiceNameArgType = string
type GetInstanceResponseGetServiceNameRetType = string

/*
	types and functions for status
*/

// isEnum

// GetInstanceResponseStatus the model 'GetInstanceResponse'
// value type for enums
type GetInstanceResponseStatus string

// List of Status
const (
	GETINSTANCERESPONSESTATUS_CREATING         GetInstanceResponseStatus = "CREATING"
	GETINSTANCERESPONSESTATUS_CREATE_SUCCEEDED GetInstanceResponseStatus = "CREATE_SUCCEEDED"
	GETINSTANCERESPONSESTATUS_CREATE_FAILED    GetInstanceResponseStatus = "CREATE_FAILED"
	GETINSTANCERESPONSESTATUS_DELETING         GetInstanceResponseStatus = "DELETING"
	GETINSTANCERESPONSESTATUS_DELETE_SUCCEEDED GetInstanceResponseStatus = "DELETE_SUCCEEDED"
	GETINSTANCERESPONSESTATUS_DELETE_FAILED    GetInstanceResponseStatus = "DELETE_FAILED"
	GETINSTANCERESPONSESTATUS_UPDATING         GetInstanceResponseStatus = "UPDATING"
	GETINSTANCERESPONSESTATUS_UPDATE_SUCCEEDED GetInstanceResponseStatus = "UPDATE_SUCCEEDED"
	GETINSTANCERESPONSESTATUS_UPDATE_FAILED    GetInstanceResponseStatus = "UPDATE_FAILED"
)

// All allowed values of GetInstanceResponse enum
var AllowedGetInstanceResponseStatusEnumValues = []GetInstanceResponseStatus{
	"CREATING",
	"CREATE_SUCCEEDED",
	"CREATE_FAILED",
	"DELETING",
	"DELETE_SUCCEEDED",
	"DELETE_FAILED",
	"UPDATING",
	"UPDATE_SUCCEEDED",
	"UPDATE_FAILED",
}

func (v *GetInstanceResponseStatus) UnmarshalJSON(src []byte) error {
	// use a type alias to prevent infinite recursion during unmarshal,
	// see https://biscuit.ninja/posts/go-avoid-an-infitine-loop-with-custom-json-unmarshallers
	type TmpJson GetInstanceResponseStatus
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
	enumTypeValue := GetInstanceResponseStatus(value)
	for _, existing := range AllowedGetInstanceResponseStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetInstanceResponse", value)
}

// NewGetInstanceResponseStatusFromValue returns a pointer to a valid GetInstanceResponseStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetInstanceResponseStatusFromValue(v GetInstanceResponseStatus) (*GetInstanceResponseStatus, error) {
	ev := GetInstanceResponseStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetInstanceResponseStatus: valid values are %v", v, AllowedGetInstanceResponseStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetInstanceResponseStatus) IsValid() bool {
	for _, existing := range AllowedGetInstanceResponseStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusStatus value
func (v GetInstanceResponseStatus) Ptr() *GetInstanceResponseStatus {
	return &v
}

type NullableGetInstanceResponseStatus struct {
	value *GetInstanceResponseStatus
	isSet bool
}

func (v NullableGetInstanceResponseStatus) Get() *GetInstanceResponseStatus {
	return v.value
}

func (v *NullableGetInstanceResponseStatus) Set(val *GetInstanceResponseStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInstanceResponseStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInstanceResponseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInstanceResponseStatus(val *GetInstanceResponseStatus) *NullableGetInstanceResponseStatus {
	return &NullableGetInstanceResponseStatus{value: val, isSet: true}
}

func (v NullableGetInstanceResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInstanceResponseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type GetInstanceResponseGetStatusAttributeType = *GetInstanceResponseStatus
type GetInstanceResponseGetStatusArgType = GetInstanceResponseStatus
type GetInstanceResponseGetStatusRetType = GetInstanceResponseStatus

func getGetInstanceResponseGetStatusAttributeTypeOk(arg GetInstanceResponseGetStatusAttributeType) (ret GetInstanceResponseGetStatusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetInstanceResponseGetStatusAttributeType(arg *GetInstanceResponseGetStatusAttributeType, val GetInstanceResponseGetStatusRetType) {
	*arg = &val
}

// GetInstanceResponse struct for GetInstanceResponse
type GetInstanceResponse struct {
	// REQUIRED
	DashboardUrl GetInstanceResponseGetDashboardUrlAttributeType `json:"dashboardUrl" required:"true"`
	Error        GetInstanceResponseGetErrorAttributeType        `json:"error,omitempty"`
	// REQUIRED
	Id GetInstanceResponseGetIdAttributeType `json:"id" required:"true"`
	// REQUIRED
	Instance    GetInstanceResponseGetInstanceAttributeType    `json:"instance" required:"true"`
	IsUpdatable GetInstanceResponsegetIsUpdatableAttributeType `json:"isUpdatable,omitempty"`
	// REQUIRED
	Message    GetInstanceResponseGetMessageAttributeType    `json:"message" required:"true"`
	Name       GetInstanceResponseGetNameAttributeType       `json:"name,omitempty"`
	Parameters GetInstanceResponseGetParametersAttributeType `json:"parameters,omitempty"`
	// REQUIRED
	PlanId GetInstanceResponseGetPlanIdAttributeType `json:"planId" required:"true"`
	// REQUIRED
	PlanName   GetInstanceResponseGetPlanNameAttributeType   `json:"planName" required:"true"`
	PlanSchema GetInstanceResponseGetPlanSchemaAttributeType `json:"planSchema,omitempty"`
	// REQUIRED
	ServiceName GetInstanceResponseGetServiceNameAttributeType `json:"serviceName" required:"true"`
	// REQUIRED
	Status GetInstanceResponseGetStatusAttributeType `json:"status" required:"true"`
}

type _GetInstanceResponse GetInstanceResponse

// NewGetInstanceResponse instantiates a new GetInstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInstanceResponse(dashboardUrl GetInstanceResponseGetDashboardUrlArgType, id GetInstanceResponseGetIdArgType, instance GetInstanceResponseGetInstanceArgType, message GetInstanceResponseGetMessageArgType, planId GetInstanceResponseGetPlanIdArgType, planName GetInstanceResponseGetPlanNameArgType, serviceName GetInstanceResponseGetServiceNameArgType, status GetInstanceResponseGetStatusArgType) *GetInstanceResponse {
	this := GetInstanceResponse{}
	setGetInstanceResponseGetDashboardUrlAttributeType(&this.DashboardUrl, dashboardUrl)
	setGetInstanceResponseGetIdAttributeType(&this.Id, id)
	setGetInstanceResponseGetInstanceAttributeType(&this.Instance, instance)
	setGetInstanceResponseGetMessageAttributeType(&this.Message, message)
	setGetInstanceResponseGetPlanIdAttributeType(&this.PlanId, planId)
	setGetInstanceResponseGetPlanNameAttributeType(&this.PlanName, planName)
	setGetInstanceResponseGetServiceNameAttributeType(&this.ServiceName, serviceName)
	setGetInstanceResponseGetStatusAttributeType(&this.Status, status)
	return &this
}

// NewGetInstanceResponseWithDefaults instantiates a new GetInstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInstanceResponseWithDefaults() *GetInstanceResponse {
	this := GetInstanceResponse{}
	var isUpdatable bool = true
	this.IsUpdatable = &isUpdatable
	var planSchema string = ""
	this.PlanSchema = &planSchema
	return &this
}

// GetDashboardUrl returns the DashboardUrl field value
func (o *GetInstanceResponse) GetDashboardUrl() (ret GetInstanceResponseGetDashboardUrlRetType) {
	ret, _ = o.GetDashboardUrlOk()
	return ret
}

// GetDashboardUrlOk returns a tuple with the DashboardUrl field value
// and a boolean to check if the value has been set.
func (o *GetInstanceResponse) GetDashboardUrlOk() (ret GetInstanceResponseGetDashboardUrlRetType, ok bool) {
	return getGetInstanceResponseGetDashboardUrlAttributeTypeOk(o.DashboardUrl)
}

// SetDashboardUrl sets field value
func (o *GetInstanceResponse) SetDashboardUrl(v GetInstanceResponseGetDashboardUrlRetType) {
	setGetInstanceResponseGetDashboardUrlAttributeType(&o.DashboardUrl, v)
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetInstanceResponse) GetError() (res GetInstanceResponseGetErrorRetType) {
	res, _ = o.GetErrorOk()
	return
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetInstanceResponse) GetErrorOk() (ret GetInstanceResponseGetErrorRetType, ok bool) {
	return getGetInstanceResponseGetErrorAttributeTypeOk(o.Error)
}

// HasError returns a boolean if a field has been set.
func (o *GetInstanceResponse) HasError() bool {
	_, ok := o.GetErrorOk()
	return ok
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetInstanceResponse) SetError(v GetInstanceResponseGetErrorRetType) {
	setGetInstanceResponseGetErrorAttributeType(&o.Error, v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *GetInstanceResponse) SetErrorNil() {
	o.Error = nil
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *GetInstanceResponse) UnsetError() {
	o.Error = nil
}

// GetId returns the Id field value
func (o *GetInstanceResponse) GetId() (ret GetInstanceResponseGetIdRetType) {
	ret, _ = o.GetIdOk()
	return ret
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetInstanceResponse) GetIdOk() (ret GetInstanceResponseGetIdRetType, ok bool) {
	return getGetInstanceResponseGetIdAttributeTypeOk(o.Id)
}

// SetId sets field value
func (o *GetInstanceResponse) SetId(v GetInstanceResponseGetIdRetType) {
	setGetInstanceResponseGetIdAttributeType(&o.Id, v)
}

// GetInstance returns the Instance field value
func (o *GetInstanceResponse) GetInstance() (ret GetInstanceResponseGetInstanceRetType) {
	ret, _ = o.GetInstanceOk()
	return ret
}

// GetInstanceOk returns a tuple with the Instance field value
// and a boolean to check if the value has been set.
func (o *GetInstanceResponse) GetInstanceOk() (ret GetInstanceResponseGetInstanceRetType, ok bool) {
	return getGetInstanceResponseGetInstanceAttributeTypeOk(o.Instance)
}

// SetInstance sets field value
func (o *GetInstanceResponse) SetInstance(v GetInstanceResponseGetInstanceRetType) {
	setGetInstanceResponseGetInstanceAttributeType(&o.Instance, v)
}

// GetIsUpdatable returns the IsUpdatable field value if set, zero value otherwise.
func (o *GetInstanceResponse) GetIsUpdatable() (res GetInstanceResponsegetIsUpdatableRetType) {
	res, _ = o.GetIsUpdatableOk()
	return
}

// GetIsUpdatableOk returns a tuple with the IsUpdatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceResponse) GetIsUpdatableOk() (ret GetInstanceResponsegetIsUpdatableRetType, ok bool) {
	return getGetInstanceResponsegetIsUpdatableAttributeTypeOk(o.IsUpdatable)
}

// HasIsUpdatable returns a boolean if a field has been set.
func (o *GetInstanceResponse) HasIsUpdatable() bool {
	_, ok := o.GetIsUpdatableOk()
	return ok
}

// SetIsUpdatable gets a reference to the given bool and assigns it to the IsUpdatable field.
func (o *GetInstanceResponse) SetIsUpdatable(v GetInstanceResponsegetIsUpdatableRetType) {
	setGetInstanceResponsegetIsUpdatableAttributeType(&o.IsUpdatable, v)
}

// GetMessage returns the Message field value
func (o *GetInstanceResponse) GetMessage() (ret GetInstanceResponseGetMessageRetType) {
	ret, _ = o.GetMessageOk()
	return ret
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GetInstanceResponse) GetMessageOk() (ret GetInstanceResponseGetMessageRetType, ok bool) {
	return getGetInstanceResponseGetMessageAttributeTypeOk(o.Message)
}

// SetMessage sets field value
func (o *GetInstanceResponse) SetMessage(v GetInstanceResponseGetMessageRetType) {
	setGetInstanceResponseGetMessageAttributeType(&o.Message, v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetInstanceResponse) GetName() (res GetInstanceResponseGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceResponse) GetNameOk() (ret GetInstanceResponseGetNameRetType, ok bool) {
	return getGetInstanceResponseGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *GetInstanceResponse) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetInstanceResponse) SetName(v GetInstanceResponseGetNameRetType) {
	setGetInstanceResponseGetNameAttributeType(&o.Name, v)
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *GetInstanceResponse) GetParameters() (res GetInstanceResponseGetParametersRetType) {
	res, _ = o.GetParametersOk()
	return
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceResponse) GetParametersOk() (ret GetInstanceResponseGetParametersRetType, ok bool) {
	return getGetInstanceResponseGetParametersAttributeTypeOk(o.Parameters)
}

// HasParameters returns a boolean if a field has been set.
func (o *GetInstanceResponse) HasParameters() bool {
	_, ok := o.GetParametersOk()
	return ok
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *GetInstanceResponse) SetParameters(v GetInstanceResponseGetParametersRetType) {
	setGetInstanceResponseGetParametersAttributeType(&o.Parameters, v)
}

// GetPlanId returns the PlanId field value
func (o *GetInstanceResponse) GetPlanId() (ret GetInstanceResponseGetPlanIdRetType) {
	ret, _ = o.GetPlanIdOk()
	return ret
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *GetInstanceResponse) GetPlanIdOk() (ret GetInstanceResponseGetPlanIdRetType, ok bool) {
	return getGetInstanceResponseGetPlanIdAttributeTypeOk(o.PlanId)
}

// SetPlanId sets field value
func (o *GetInstanceResponse) SetPlanId(v GetInstanceResponseGetPlanIdRetType) {
	setGetInstanceResponseGetPlanIdAttributeType(&o.PlanId, v)
}

// GetPlanName returns the PlanName field value
func (o *GetInstanceResponse) GetPlanName() (ret GetInstanceResponseGetPlanNameRetType) {
	ret, _ = o.GetPlanNameOk()
	return ret
}

// GetPlanNameOk returns a tuple with the PlanName field value
// and a boolean to check if the value has been set.
func (o *GetInstanceResponse) GetPlanNameOk() (ret GetInstanceResponseGetPlanNameRetType, ok bool) {
	return getGetInstanceResponseGetPlanNameAttributeTypeOk(o.PlanName)
}

// SetPlanName sets field value
func (o *GetInstanceResponse) SetPlanName(v GetInstanceResponseGetPlanNameRetType) {
	setGetInstanceResponseGetPlanNameAttributeType(&o.PlanName, v)
}

// GetPlanSchema returns the PlanSchema field value if set, zero value otherwise.
func (o *GetInstanceResponse) GetPlanSchema() (res GetInstanceResponseGetPlanSchemaRetType) {
	res, _ = o.GetPlanSchemaOk()
	return
}

// GetPlanSchemaOk returns a tuple with the PlanSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceResponse) GetPlanSchemaOk() (ret GetInstanceResponseGetPlanSchemaRetType, ok bool) {
	return getGetInstanceResponseGetPlanSchemaAttributeTypeOk(o.PlanSchema)
}

// HasPlanSchema returns a boolean if a field has been set.
func (o *GetInstanceResponse) HasPlanSchema() bool {
	_, ok := o.GetPlanSchemaOk()
	return ok
}

// SetPlanSchema gets a reference to the given string and assigns it to the PlanSchema field.
func (o *GetInstanceResponse) SetPlanSchema(v GetInstanceResponseGetPlanSchemaRetType) {
	setGetInstanceResponseGetPlanSchemaAttributeType(&o.PlanSchema, v)
}

// GetServiceName returns the ServiceName field value
func (o *GetInstanceResponse) GetServiceName() (ret GetInstanceResponseGetServiceNameRetType) {
	ret, _ = o.GetServiceNameOk()
	return ret
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *GetInstanceResponse) GetServiceNameOk() (ret GetInstanceResponseGetServiceNameRetType, ok bool) {
	return getGetInstanceResponseGetServiceNameAttributeTypeOk(o.ServiceName)
}

// SetServiceName sets field value
func (o *GetInstanceResponse) SetServiceName(v GetInstanceResponseGetServiceNameRetType) {
	setGetInstanceResponseGetServiceNameAttributeType(&o.ServiceName, v)
}

// GetStatus returns the Status field value
func (o *GetInstanceResponse) GetStatus() (ret GetInstanceResponseGetStatusRetType) {
	ret, _ = o.GetStatusOk()
	return ret
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetInstanceResponse) GetStatusOk() (ret GetInstanceResponseGetStatusRetType, ok bool) {
	return getGetInstanceResponseGetStatusAttributeTypeOk(o.Status)
}

// SetStatus sets field value
func (o *GetInstanceResponse) SetStatus(v GetInstanceResponseGetStatusRetType) {
	setGetInstanceResponseGetStatusAttributeType(&o.Status, v)
}

func (o GetInstanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getGetInstanceResponseGetDashboardUrlAttributeTypeOk(o.DashboardUrl); ok {
		toSerialize["DashboardUrl"] = val
	}
	if val, ok := getGetInstanceResponseGetErrorAttributeTypeOk(o.Error); ok {
		toSerialize["Error"] = val
	}
	if val, ok := getGetInstanceResponseGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getGetInstanceResponseGetInstanceAttributeTypeOk(o.Instance); ok {
		toSerialize["Instance"] = val
	}
	if val, ok := getGetInstanceResponsegetIsUpdatableAttributeTypeOk(o.IsUpdatable); ok {
		toSerialize["IsUpdatable"] = val
	}
	if val, ok := getGetInstanceResponseGetMessageAttributeTypeOk(o.Message); ok {
		toSerialize["Message"] = val
	}
	if val, ok := getGetInstanceResponseGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getGetInstanceResponseGetParametersAttributeTypeOk(o.Parameters); ok {
		toSerialize["Parameters"] = val
	}
	if val, ok := getGetInstanceResponseGetPlanIdAttributeTypeOk(o.PlanId); ok {
		toSerialize["PlanId"] = val
	}
	if val, ok := getGetInstanceResponseGetPlanNameAttributeTypeOk(o.PlanName); ok {
		toSerialize["PlanName"] = val
	}
	if val, ok := getGetInstanceResponseGetPlanSchemaAttributeTypeOk(o.PlanSchema); ok {
		toSerialize["PlanSchema"] = val
	}
	if val, ok := getGetInstanceResponseGetServiceNameAttributeTypeOk(o.ServiceName); ok {
		toSerialize["ServiceName"] = val
	}
	if val, ok := getGetInstanceResponseGetStatusAttributeTypeOk(o.Status); ok {
		toSerialize["Status"] = val
	}
	return toSerialize, nil
}

type NullableGetInstanceResponse struct {
	value *GetInstanceResponse
	isSet bool
}

func (v NullableGetInstanceResponse) Get() *GetInstanceResponse {
	return v.value
}

func (v *NullableGetInstanceResponse) Set(val *GetInstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInstanceResponse(val *GetInstanceResponse) *NullableGetInstanceResponse {
	return &NullableGetInstanceResponse{value: val, isSet: true}
}

func (v NullableGetInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
