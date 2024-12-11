/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

import (
	"encoding/json"
)

// checks if the SlowQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlowQuery{}

// SlowQuery struct for SlowQuery
type SlowQuery struct {
	// The raw log line pertaining to the slow query.
	Line *string `json:"line,omitempty"`
	// The namespace in which the slow query ran.
	Namespace *string `json:"namespace,omitempty"`
}

// NewSlowQuery instantiates a new SlowQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlowQuery() *SlowQuery {
	this := SlowQuery{}
	return &this
}

// NewSlowQueryWithDefaults instantiates a new SlowQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlowQueryWithDefaults() *SlowQuery {
	this := SlowQuery{}
	return &this
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *SlowQuery) GetLine() *string {
	if o == nil || IsNil(o.Line) {
		var ret *string
		return ret
	}
	return o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlowQuery) GetLineOk() (*string, bool) {
	if o == nil || IsNil(o.Line) {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *SlowQuery) HasLine() bool {
	if o != nil && !IsNil(o.Line) && !IsNil(o.Line) {
		return true
	}

	return false
}

// SetLine gets a reference to the given string and assigns it to the Line field.
func (o *SlowQuery) SetLine(v *string) {
	o.Line = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *SlowQuery) GetNamespace() *string {
	if o == nil || IsNil(o.Namespace) {
		var ret *string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlowQuery) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *SlowQuery) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *SlowQuery) SetNamespace(v *string) {
	o.Namespace = v
}

func (o SlowQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Line) {
		toSerialize["line"] = o.Line
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

type NullableSlowQuery struct {
	value *SlowQuery
	isSet bool
}

func (v NullableSlowQuery) Get() *SlowQuery {
	return v.value
}

func (v *NullableSlowQuery) Set(val *SlowQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSlowQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSlowQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlowQuery(val *SlowQuery) *NullableSlowQuery {
	return &NullableSlowQuery{value: val, isSet: true}
}

func (v NullableSlowQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlowQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
