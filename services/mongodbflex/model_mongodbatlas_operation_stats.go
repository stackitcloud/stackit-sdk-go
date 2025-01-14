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

// checks if the MongodbatlasOperationStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MongodbatlasOperationStats{}

// MongodbatlasOperationStats Query statistics.
type MongodbatlasOperationStats struct {
	// Duration in milliseconds of the query.
	Ms *float64 `json:"ms,omitempty"`
	// Number of results returned by the query.
	NReturned *int64 `json:"nReturned,omitempty"`
	// Number of documents read by the query.
	NScanned *int64 `json:"nScanned,omitempty"`
	// Query timestamp, in seconds since epoch.
	Ts *int64 `json:"ts,omitempty"`
}

// NewMongodbatlasOperationStats instantiates a new MongodbatlasOperationStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongodbatlasOperationStats() *MongodbatlasOperationStats {
	this := MongodbatlasOperationStats{}
	return &this
}

// NewMongodbatlasOperationStatsWithDefaults instantiates a new MongodbatlasOperationStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongodbatlasOperationStatsWithDefaults() *MongodbatlasOperationStats {
	this := MongodbatlasOperationStats{}
	return &this
}

// GetMs returns the Ms field value if set, zero value otherwise.
func (o *MongodbatlasOperationStats) GetMs() *float64 {
	if o == nil || IsNil(o.Ms) {
		var ret *float64
		return ret
	}
	return o.Ms
}

// GetMsOk returns a tuple with the Ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongodbatlasOperationStats) GetMsOk() (*float64, bool) {
	if o == nil || IsNil(o.Ms) {
		return nil, false
	}
	return o.Ms, true
}

// HasMs returns a boolean if a field has been set.
func (o *MongodbatlasOperationStats) HasMs() bool {
	if o != nil && !IsNil(o.Ms) {
		return true
	}

	return false
}

// SetMs gets a reference to the given float64 and assigns it to the Ms field.
func (o *MongodbatlasOperationStats) SetMs(v *float64) {
	o.Ms = v
}

// GetNReturned returns the NReturned field value if set, zero value otherwise.
func (o *MongodbatlasOperationStats) GetNReturned() *int64 {
	if o == nil || IsNil(o.NReturned) {
		var ret *int64
		return ret
	}
	return o.NReturned
}

// GetNReturnedOk returns a tuple with the NReturned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongodbatlasOperationStats) GetNReturnedOk() (*int64, bool) {
	if o == nil || IsNil(o.NReturned) {
		return nil, false
	}
	return o.NReturned, true
}

// HasNReturned returns a boolean if a field has been set.
func (o *MongodbatlasOperationStats) HasNReturned() bool {
	if o != nil && !IsNil(o.NReturned) {
		return true
	}

	return false
}

// SetNReturned gets a reference to the given int64 and assigns it to the NReturned field.
func (o *MongodbatlasOperationStats) SetNReturned(v *int64) {
	o.NReturned = v
}

// GetNScanned returns the NScanned field value if set, zero value otherwise.
func (o *MongodbatlasOperationStats) GetNScanned() *int64 {
	if o == nil || IsNil(o.NScanned) {
		var ret *int64
		return ret
	}
	return o.NScanned
}

// GetNScannedOk returns a tuple with the NScanned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongodbatlasOperationStats) GetNScannedOk() (*int64, bool) {
	if o == nil || IsNil(o.NScanned) {
		return nil, false
	}
	return o.NScanned, true
}

// HasNScanned returns a boolean if a field has been set.
func (o *MongodbatlasOperationStats) HasNScanned() bool {
	if o != nil && !IsNil(o.NScanned) {
		return true
	}

	return false
}

// SetNScanned gets a reference to the given int64 and assigns it to the NScanned field.
func (o *MongodbatlasOperationStats) SetNScanned(v *int64) {
	o.NScanned = v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *MongodbatlasOperationStats) GetTs() *int64 {
	if o == nil || IsNil(o.Ts) {
		var ret *int64
		return ret
	}
	return o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongodbatlasOperationStats) GetTsOk() (*int64, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *MongodbatlasOperationStats) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given int64 and assigns it to the Ts field.
func (o *MongodbatlasOperationStats) SetTs(v *int64) {
	o.Ts = v
}

func (o MongodbatlasOperationStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ms) {
		toSerialize["ms"] = o.Ms
	}
	if !IsNil(o.NReturned) {
		toSerialize["nReturned"] = o.NReturned
	}
	if !IsNil(o.NScanned) {
		toSerialize["nScanned"] = o.NScanned
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableMongodbatlasOperationStats struct {
	value *MongodbatlasOperationStats
	isSet bool
}

func (v NullableMongodbatlasOperationStats) Get() *MongodbatlasOperationStats {
	return v.value
}

func (v *NullableMongodbatlasOperationStats) Set(val *MongodbatlasOperationStats) {
	v.value = val
	v.isSet = true
}

func (v NullableMongodbatlasOperationStats) IsSet() bool {
	return v.isSet
}

func (v *NullableMongodbatlasOperationStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMongodbatlasOperationStats(val *MongodbatlasOperationStats) *NullableMongodbatlasOperationStats {
	return &NullableMongodbatlasOperationStats{value: val, isSet: true}
}

func (v NullableMongodbatlasOperationStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMongodbatlasOperationStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
