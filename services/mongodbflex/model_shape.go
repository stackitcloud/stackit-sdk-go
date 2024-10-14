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

// checks if the Shape type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Shape{}

// Shape struct for Shape
type Shape struct {
	// Average duration in milliseconds for the queries examined that match this shape.
	AvgMs *float64 `json:"avgMs,omitempty"`
	// Number of queries examined that match this shape.
	Count *int64 `json:"count,omitempty"`
	// Unique id for this shape. Exists only for the duration of the API request.
	Id *string `json:"id,omitempty"`
	// Average number of documents read for every document returned by the query.
	InefficiencyScore *int64 `json:"inefficiencyScore,omitempty"`
	// The namespace in which the slow query ran.
	Namespace *string `json:"namespace,omitempty"`
	// It represents documents with specific information and log lines for individual queries.
	Operations *[]MongodbatlasOperation `json:"operations,omitempty"`
}

// NewShape instantiates a new Shape object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShape() *Shape {
	this := Shape{}
	return &this
}

// NewShapeWithDefaults instantiates a new Shape object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShapeWithDefaults() *Shape {
	this := Shape{}
	return &this
}

// GetAvgMs returns the AvgMs field value if set, zero value otherwise.
func (o *Shape) GetAvgMs() *float64 {
	if o == nil || IsNil(o.AvgMs) {
		var ret *float64
		return ret
	}
	return o.AvgMs
}

// GetAvgMsOk returns a tuple with the AvgMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shape) GetAvgMsOk() (*float64, bool) {
	if o == nil || IsNil(o.AvgMs) {
		return nil, false
	}
	return o.AvgMs, true
}

// HasAvgMs returns a boolean if a field has been set.
func (o *Shape) HasAvgMs() bool {
	if o != nil && !IsNil(o.AvgMs) {
		return true
	}

	return false
}

// SetAvgMs gets a reference to the given float64 and assigns it to the AvgMs field.
func (o *Shape) SetAvgMs(v *float64) {
	o.AvgMs = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Shape) GetCount() *int64 {
	if o == nil || IsNil(o.Count) {
		var ret *int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shape) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Shape) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *Shape) SetCount(v *int64) {
	o.Count = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Shape) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shape) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Shape) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Shape) SetId(v *string) {
	o.Id = v
}

// GetInefficiencyScore returns the InefficiencyScore field value if set, zero value otherwise.
func (o *Shape) GetInefficiencyScore() *int64 {
	if o == nil || IsNil(o.InefficiencyScore) {
		var ret *int64
		return ret
	}
	return o.InefficiencyScore
}

// GetInefficiencyScoreOk returns a tuple with the InefficiencyScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shape) GetInefficiencyScoreOk() (*int64, bool) {
	if o == nil || IsNil(o.InefficiencyScore) {
		return nil, false
	}
	return o.InefficiencyScore, true
}

// HasInefficiencyScore returns a boolean if a field has been set.
func (o *Shape) HasInefficiencyScore() bool {
	if o != nil && !IsNil(o.InefficiencyScore) {
		return true
	}

	return false
}

// SetInefficiencyScore gets a reference to the given int64 and assigns it to the InefficiencyScore field.
func (o *Shape) SetInefficiencyScore(v *int64) {
	o.InefficiencyScore = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Shape) GetNamespace() *string {
	if o == nil || IsNil(o.Namespace) {
		var ret *string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shape) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Shape) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *Shape) SetNamespace(v *string) {
	o.Namespace = v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *Shape) GetOperations() *[]MongodbatlasOperation {
	if o == nil || IsNil(o.Operations) {
		var ret *[]MongodbatlasOperation
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shape) GetOperationsOk() (*[]MongodbatlasOperation, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *Shape) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []MongodbatlasOperation and assigns it to the Operations field.
func (o *Shape) SetOperations(v *[]MongodbatlasOperation) {
	o.Operations = v
}

func (o Shape) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvgMs) {
		toSerialize["avgMs"] = o.AvgMs
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InefficiencyScore) {
		toSerialize["inefficiencyScore"] = o.InefficiencyScore
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	return toSerialize, nil
}

type NullableShape struct {
	value *Shape
	isSet bool
}

func (v NullableShape) Get() *Shape {
	return v.value
}

func (v *NullableShape) Set(val *Shape) {
	v.value = val
	v.isSet = true
}

func (v NullableShape) IsSet() bool {
	return v.isSet
}

func (v *NullableShape) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShape(val *Shape) *NullableShape {
	return &NullableShape{value: val, isSet: true}
}

func (v NullableShape) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShape) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
