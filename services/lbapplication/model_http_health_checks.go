/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lbapplication

import (
	"encoding/json"
)

// checks if the HttpHealthChecks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpHealthChecks{}

/*
	types and functions for okStatuses
*/

// isArray
type HttpHealthChecksGetOkStatusesAttributeType = *[]string
type HttpHealthChecksGetOkStatusesArgType = []string
type HttpHealthChecksGetOkStatusesRetType = []string

func getHttpHealthChecksGetOkStatusesAttributeTypeOk(arg HttpHealthChecksGetOkStatusesAttributeType) (ret HttpHealthChecksGetOkStatusesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setHttpHealthChecksGetOkStatusesAttributeType(arg *HttpHealthChecksGetOkStatusesAttributeType, val HttpHealthChecksGetOkStatusesRetType) {
	*arg = &val
}

/*
	types and functions for path
*/

// isNotNullableString
type HttpHealthChecksGetPathAttributeType = *string

func getHttpHealthChecksGetPathAttributeTypeOk(arg HttpHealthChecksGetPathAttributeType) (ret HttpHealthChecksGetPathRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setHttpHealthChecksGetPathAttributeType(arg *HttpHealthChecksGetPathAttributeType, val HttpHealthChecksGetPathRetType) {
	*arg = &val
}

type HttpHealthChecksGetPathArgType = string
type HttpHealthChecksGetPathRetType = string

// HttpHealthChecks struct for HttpHealthChecks
type HttpHealthChecks struct {
	// List of HTTP status codes that indicate a healthy response
	OkStatuses HttpHealthChecksGetOkStatusesAttributeType `json:"okStatuses,omitempty"`
	// Path to send the health check request to
	Path HttpHealthChecksGetPathAttributeType `json:"path,omitempty"`
}

// NewHttpHealthChecks instantiates a new HttpHealthChecks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpHealthChecks() *HttpHealthChecks {
	this := HttpHealthChecks{}
	return &this
}

// NewHttpHealthChecksWithDefaults instantiates a new HttpHealthChecks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpHealthChecksWithDefaults() *HttpHealthChecks {
	this := HttpHealthChecks{}
	return &this
}

// GetOkStatuses returns the OkStatuses field value if set, zero value otherwise.
func (o *HttpHealthChecks) GetOkStatuses() (res HttpHealthChecksGetOkStatusesRetType) {
	res, _ = o.GetOkStatusesOk()
	return
}

// GetOkStatusesOk returns a tuple with the OkStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpHealthChecks) GetOkStatusesOk() (ret HttpHealthChecksGetOkStatusesRetType, ok bool) {
	return getHttpHealthChecksGetOkStatusesAttributeTypeOk(o.OkStatuses)
}

// HasOkStatuses returns a boolean if a field has been set.
func (o *HttpHealthChecks) HasOkStatuses() bool {
	_, ok := o.GetOkStatusesOk()
	return ok
}

// SetOkStatuses gets a reference to the given []string and assigns it to the OkStatuses field.
func (o *HttpHealthChecks) SetOkStatuses(v HttpHealthChecksGetOkStatusesRetType) {
	setHttpHealthChecksGetOkStatusesAttributeType(&o.OkStatuses, v)
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HttpHealthChecks) GetPath() (res HttpHealthChecksGetPathRetType) {
	res, _ = o.GetPathOk()
	return
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpHealthChecks) GetPathOk() (ret HttpHealthChecksGetPathRetType, ok bool) {
	return getHttpHealthChecksGetPathAttributeTypeOk(o.Path)
}

// HasPath returns a boolean if a field has been set.
func (o *HttpHealthChecks) HasPath() bool {
	_, ok := o.GetPathOk()
	return ok
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *HttpHealthChecks) SetPath(v HttpHealthChecksGetPathRetType) {
	setHttpHealthChecksGetPathAttributeType(&o.Path, v)
}

func (o HttpHealthChecks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getHttpHealthChecksGetOkStatusesAttributeTypeOk(o.OkStatuses); ok {
		toSerialize["OkStatuses"] = val
	}
	if val, ok := getHttpHealthChecksGetPathAttributeTypeOk(o.Path); ok {
		toSerialize["Path"] = val
	}
	return toSerialize, nil
}

type NullableHttpHealthChecks struct {
	value *HttpHealthChecks
	isSet bool
}

func (v NullableHttpHealthChecks) Get() *HttpHealthChecks {
	return v.value
}

func (v *NullableHttpHealthChecks) Set(val *HttpHealthChecks) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpHealthChecks) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpHealthChecks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpHealthChecks(val *HttpHealthChecks) *NullableHttpHealthChecks {
	return &NullableHttpHealthChecks{value: val, isSet: true}
}

func (v NullableHttpHealthChecks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpHealthChecks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
