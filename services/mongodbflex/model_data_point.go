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

// checks if the DataPoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataPoint{}

// DataPoint struct for DataPoint
type DataPoint struct {
	Timestamp *string  `json:"timestamp,omitempty"`
	Value     *float64 `json:"value,omitempty"`
}

// NewDataPoint instantiates a new DataPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPoint() *DataPoint {
	this := DataPoint{}
	return &this
}

// NewDataPointWithDefaults instantiates a new DataPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPointWithDefaults() *DataPoint {
	this := DataPoint{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *DataPoint) GetTimestamp() *string {
	if o == nil || IsNil(o.Timestamp) {
		var ret *string
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPoint) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DataPoint) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *DataPoint) SetTimestamp(v *string) {
	o.Timestamp = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DataPoint) GetValue() *float64 {
	if o == nil || IsNil(o.Value) {
		var ret *float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPoint) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DataPoint) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *DataPoint) SetValue(v *float64) {
	o.Value = v
}

func (o DataPoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableDataPoint struct {
	value *DataPoint
	isSet bool
}

func (v NullableDataPoint) Get() *DataPoint {
	return v.value
}

func (v *NullableDataPoint) Set(val *DataPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPoint(val *DataPoint) *NullableDataPoint {
	return &NullableDataPoint{value: val, isSet: true}
}

func (v NullableDataPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
