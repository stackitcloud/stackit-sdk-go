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

// checks if the MongodbatlasSuggestedIndex type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MongodbatlasSuggestedIndex{}

// MongodbatlasSuggestedIndex struct for MongodbatlasSuggestedIndex
type MongodbatlasSuggestedIndex struct {
	// Unique id for this suggested index.
	Id *string `json:"id,omitempty"`
	// List of unique identifiers which correspond the query shapes in this response which pertain to this suggested index.
	Impact *[]string `json:"impact,omitempty"`
	// Array of documents that specifies a key in the index and its sort order, ascending or descending.
	Index *[]map[string]int32 `json:"index,omitempty"`
	// Namespace of the suggested index.
	Namespace *string `json:"namespace,omitempty"`
	// Estimated percentage performance improvement that the suggested index would provide.
	Weight *float64 `json:"weight,omitempty"`
}

// NewMongodbatlasSuggestedIndex instantiates a new MongodbatlasSuggestedIndex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongodbatlasSuggestedIndex() *MongodbatlasSuggestedIndex {
	this := MongodbatlasSuggestedIndex{}
	return &this
}

// NewMongodbatlasSuggestedIndexWithDefaults instantiates a new MongodbatlasSuggestedIndex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongodbatlasSuggestedIndexWithDefaults() *MongodbatlasSuggestedIndex {
	this := MongodbatlasSuggestedIndex{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MongodbatlasSuggestedIndex) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongodbatlasSuggestedIndex) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MongodbatlasSuggestedIndex) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MongodbatlasSuggestedIndex) SetId(v *string) {
	o.Id = v
}

// GetImpact returns the Impact field value if set, zero value otherwise.
func (o *MongodbatlasSuggestedIndex) GetImpact() *[]string {
	if o == nil || IsNil(o.Impact) {
		var ret *[]string
		return ret
	}
	return o.Impact
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongodbatlasSuggestedIndex) GetImpactOk() (*[]string, bool) {
	if o == nil || IsNil(o.Impact) {
		return nil, false
	}
	return o.Impact, true
}

// HasImpact returns a boolean if a field has been set.
func (o *MongodbatlasSuggestedIndex) HasImpact() bool {
	if o != nil && !IsNil(o.Impact) {
		return true
	}

	return false
}

// SetImpact gets a reference to the given []string and assigns it to the Impact field.
func (o *MongodbatlasSuggestedIndex) SetImpact(v *[]string) {
	o.Impact = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *MongodbatlasSuggestedIndex) GetIndex() *[]map[string]int32 {
	if o == nil || IsNil(o.Index) {
		var ret *[]map[string]int32
		return ret
	}
	return o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongodbatlasSuggestedIndex) GetIndexOk() (*[]map[string]int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *MongodbatlasSuggestedIndex) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given []map[string]int32 and assigns it to the Index field.
func (o *MongodbatlasSuggestedIndex) SetIndex(v *[]map[string]int32) {
	o.Index = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *MongodbatlasSuggestedIndex) GetNamespace() *string {
	if o == nil || IsNil(o.Namespace) {
		var ret *string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongodbatlasSuggestedIndex) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *MongodbatlasSuggestedIndex) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *MongodbatlasSuggestedIndex) SetNamespace(v *string) {
	o.Namespace = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *MongodbatlasSuggestedIndex) GetWeight() *float64 {
	if o == nil || IsNil(o.Weight) {
		var ret *float64
		return ret
	}
	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongodbatlasSuggestedIndex) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *MongodbatlasSuggestedIndex) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *MongodbatlasSuggestedIndex) SetWeight(v *float64) {
	o.Weight = v
}

func (o MongodbatlasSuggestedIndex) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Impact) {
		toSerialize["impact"] = o.Impact
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableMongodbatlasSuggestedIndex struct {
	value *MongodbatlasSuggestedIndex
	isSet bool
}

func (v NullableMongodbatlasSuggestedIndex) Get() *MongodbatlasSuggestedIndex {
	return v.value
}

func (v *NullableMongodbatlasSuggestedIndex) Set(val *MongodbatlasSuggestedIndex) {
	v.value = val
	v.isSet = true
}

func (v NullableMongodbatlasSuggestedIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableMongodbatlasSuggestedIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMongodbatlasSuggestedIndex(val *MongodbatlasSuggestedIndex) *NullableMongodbatlasSuggestedIndex {
	return &NullableMongodbatlasSuggestedIndex{value: val, isSet: true}
}

func (v NullableMongodbatlasSuggestedIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMongodbatlasSuggestedIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
