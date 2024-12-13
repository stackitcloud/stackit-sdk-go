/*
Resource Manager API

API v2 to manage resource containers - organizations, folders, projects incl. labels  ### Resource Management STACKIT resource management handles the terms _Organization_, _Folder_, _Project_, _Label_, and the hierarchical structure between them. Technically, organizations,  folders, and projects are _Resource Containers_ to which a _Label_ can be attached to. The STACKIT _Resource Manager_ provides CRUD endpoints to query and to modify the state.  ### Organizations STACKIT organizations are the base element to create and to use cloud-resources. An organization is bound to one customer account. Organizations have a lifecycle. - Organizations are always the root node in resource hierarchy and do not have a parent  ### Projects STACKIT projects are needed to use cloud-resources. Projects serve as wrapper for underlying technical structures and processes. Projects have a lifecycle. Projects compared to folders may have different policies. - Projects are optional, but mandatory for cloud-resource usage - A project can be created having either an organization, or a folder as parent - A project must not have a project as parent - Project names under the same parent must not be unique - Root organization cannot be changed  ### Label STACKIT labels are key-value pairs including a resource container reference. Labels can be defined and attached freely to resource containers by which resources can be organized and queried. - Policy-based, immutable labels may exists

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resourcemanager

import (
	"encoding/json"
	"time"
)

// checks if the ErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponse{}

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	// The reason phrase of the status code.
	// REQUIRED
	Error *string `json:"error"`
	// Description of the error.
	// REQUIRED
	Message *string `json:"message"`
	// Path which was called.
	// REQUIRED
	Path *string `json:"path"`
	// Http Status Code.
	// REQUIRED
	Status *float64 `json:"status"`
	// Timestamp at which the error occurred.
	// REQUIRED
	TimeStamp *time.Time `json:"timeStamp"`
}

type _ErrorResponse ErrorResponse

// NewErrorResponse instantiates a new ErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponse(error_ *string, message *string, path *string, status *float64, timeStamp *time.Time) *ErrorResponse {
	this := ErrorResponse{}
	this.Error = error_
	this.Message = message
	this.Path = path
	this.Status = status
	this.TimeStamp = timeStamp
	return &this
}

// NewErrorResponseWithDefaults instantiates a new ErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseWithDefaults() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// GetError returns the Error field value
func (o *ErrorResponse) GetError() *string {
	if o == nil || IsNil(o.Error) {
		var ret *string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error, true
}

// SetError sets field value
func (o *ErrorResponse) SetError(v *string) {
	o.Error = v
}

// GetMessage returns the Message field value
func (o *ErrorResponse) GetMessage() *string {
	if o == nil || IsNil(o.Message) {
		var ret *string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *ErrorResponse) SetMessage(v *string) {
	o.Message = v
}

// GetPath returns the Path field value
func (o *ErrorResponse) GetPath() *string {
	if o == nil || IsNil(o.Path) {
		var ret *string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path, true
}

// SetPath sets field value
func (o *ErrorResponse) SetPath(v *string) {
	o.Path = v
}

// GetStatus returns the Status field value
func (o *ErrorResponse) GetStatus() *float64 {
	if o == nil || IsNil(o.Status) {
		var ret *float64
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetStatusOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *ErrorResponse) SetStatus(v *float64) {
	o.Status = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *ErrorResponse) GetTimeStamp() *time.Time {
	if o == nil || IsNil(o.TimeStamp) {
		var ret *time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *ErrorResponse) SetTimeStamp(v *time.Time) {
	o.TimeStamp = v
}

func (o ErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["message"] = o.Message
	toSerialize["path"] = o.Path
	toSerialize["status"] = o.Status
	toSerialize["timeStamp"] = o.TimeStamp
	return toSerialize, nil
}

type NullableErrorResponse struct {
	value *ErrorResponse
	isSet bool
}

func (v NullableErrorResponse) Get() *ErrorResponse {
	return v.value
}

func (v *NullableErrorResponse) Set(val *ErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponse(val *ErrorResponse) *NullableErrorResponse {
	return &NullableErrorResponse{value: val, isSet: true}
}

func (v NullableErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
