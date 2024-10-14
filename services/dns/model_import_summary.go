/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the ImportSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportSummary{}

// ImportSummary ImportSummary is the summary of the import.
type ImportSummary struct {
	CreatedRRSets  *int64 `json:"createdRRSets,omitempty"`
	CreatedRecords *int64 `json:"createdRecords,omitempty"`
	DeletedRRSets  *int64 `json:"deletedRRSets,omitempty"`
	DeletedRecords *int64 `json:"deletedRecords,omitempty"`
	UpdatedRRSets  *int64 `json:"updatedRRSets,omitempty"`
	UpdatedRecords *int64 `json:"updatedRecords,omitempty"`
}

// NewImportSummary instantiates a new ImportSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportSummary() *ImportSummary {
	this := ImportSummary{}
	return &this
}

// NewImportSummaryWithDefaults instantiates a new ImportSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportSummaryWithDefaults() *ImportSummary {
	this := ImportSummary{}
	return &this
}

// GetCreatedRRSets returns the CreatedRRSets field value if set, zero value otherwise.
func (o *ImportSummary) GetCreatedRRSets() *int64 {
	if o == nil || IsNil(o.CreatedRRSets) {
		var ret *int64
		return ret
	}
	return o.CreatedRRSets
}

// GetCreatedRRSetsOk returns a tuple with the CreatedRRSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetCreatedRRSetsOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedRRSets) {
		return nil, false
	}
	return o.CreatedRRSets, true
}

// HasCreatedRRSets returns a boolean if a field has been set.
func (o *ImportSummary) HasCreatedRRSets() bool {
	if o != nil && !IsNil(o.CreatedRRSets) {
		return true
	}

	return false
}

// SetCreatedRRSets gets a reference to the given int64 and assigns it to the CreatedRRSets field.
func (o *ImportSummary) SetCreatedRRSets(v *int64) {
	o.CreatedRRSets = v
}

// GetCreatedRecords returns the CreatedRecords field value if set, zero value otherwise.
func (o *ImportSummary) GetCreatedRecords() *int64 {
	if o == nil || IsNil(o.CreatedRecords) {
		var ret *int64
		return ret
	}
	return o.CreatedRecords
}

// GetCreatedRecordsOk returns a tuple with the CreatedRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetCreatedRecordsOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedRecords) {
		return nil, false
	}
	return o.CreatedRecords, true
}

// HasCreatedRecords returns a boolean if a field has been set.
func (o *ImportSummary) HasCreatedRecords() bool {
	if o != nil && !IsNil(o.CreatedRecords) {
		return true
	}

	return false
}

// SetCreatedRecords gets a reference to the given int64 and assigns it to the CreatedRecords field.
func (o *ImportSummary) SetCreatedRecords(v *int64) {
	o.CreatedRecords = v
}

// GetDeletedRRSets returns the DeletedRRSets field value if set, zero value otherwise.
func (o *ImportSummary) GetDeletedRRSets() *int64 {
	if o == nil || IsNil(o.DeletedRRSets) {
		var ret *int64
		return ret
	}
	return o.DeletedRRSets
}

// GetDeletedRRSetsOk returns a tuple with the DeletedRRSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetDeletedRRSetsOk() (*int64, bool) {
	if o == nil || IsNil(o.DeletedRRSets) {
		return nil, false
	}
	return o.DeletedRRSets, true
}

// HasDeletedRRSets returns a boolean if a field has been set.
func (o *ImportSummary) HasDeletedRRSets() bool {
	if o != nil && !IsNil(o.DeletedRRSets) {
		return true
	}

	return false
}

// SetDeletedRRSets gets a reference to the given int64 and assigns it to the DeletedRRSets field.
func (o *ImportSummary) SetDeletedRRSets(v *int64) {
	o.DeletedRRSets = v
}

// GetDeletedRecords returns the DeletedRecords field value if set, zero value otherwise.
func (o *ImportSummary) GetDeletedRecords() *int64 {
	if o == nil || IsNil(o.DeletedRecords) {
		var ret *int64
		return ret
	}
	return o.DeletedRecords
}

// GetDeletedRecordsOk returns a tuple with the DeletedRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetDeletedRecordsOk() (*int64, bool) {
	if o == nil || IsNil(o.DeletedRecords) {
		return nil, false
	}
	return o.DeletedRecords, true
}

// HasDeletedRecords returns a boolean if a field has been set.
func (o *ImportSummary) HasDeletedRecords() bool {
	if o != nil && !IsNil(o.DeletedRecords) {
		return true
	}

	return false
}

// SetDeletedRecords gets a reference to the given int64 and assigns it to the DeletedRecords field.
func (o *ImportSummary) SetDeletedRecords(v *int64) {
	o.DeletedRecords = v
}

// GetUpdatedRRSets returns the UpdatedRRSets field value if set, zero value otherwise.
func (o *ImportSummary) GetUpdatedRRSets() *int64 {
	if o == nil || IsNil(o.UpdatedRRSets) {
		var ret *int64
		return ret
	}
	return o.UpdatedRRSets
}

// GetUpdatedRRSetsOk returns a tuple with the UpdatedRRSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetUpdatedRRSetsOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedRRSets) {
		return nil, false
	}
	return o.UpdatedRRSets, true
}

// HasUpdatedRRSets returns a boolean if a field has been set.
func (o *ImportSummary) HasUpdatedRRSets() bool {
	if o != nil && !IsNil(o.UpdatedRRSets) {
		return true
	}

	return false
}

// SetUpdatedRRSets gets a reference to the given int64 and assigns it to the UpdatedRRSets field.
func (o *ImportSummary) SetUpdatedRRSets(v *int64) {
	o.UpdatedRRSets = v
}

// GetUpdatedRecords returns the UpdatedRecords field value if set, zero value otherwise.
func (o *ImportSummary) GetUpdatedRecords() *int64 {
	if o == nil || IsNil(o.UpdatedRecords) {
		var ret *int64
		return ret
	}
	return o.UpdatedRecords
}

// GetUpdatedRecordsOk returns a tuple with the UpdatedRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummary) GetUpdatedRecordsOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedRecords) {
		return nil, false
	}
	return o.UpdatedRecords, true
}

// HasUpdatedRecords returns a boolean if a field has been set.
func (o *ImportSummary) HasUpdatedRecords() bool {
	if o != nil && !IsNil(o.UpdatedRecords) {
		return true
	}

	return false
}

// SetUpdatedRecords gets a reference to the given int64 and assigns it to the UpdatedRecords field.
func (o *ImportSummary) SetUpdatedRecords(v *int64) {
	o.UpdatedRecords = v
}

func (o ImportSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedRRSets) {
		toSerialize["createdRRSets"] = o.CreatedRRSets
	}
	if !IsNil(o.CreatedRecords) {
		toSerialize["createdRecords"] = o.CreatedRecords
	}
	if !IsNil(o.DeletedRRSets) {
		toSerialize["deletedRRSets"] = o.DeletedRRSets
	}
	if !IsNil(o.DeletedRecords) {
		toSerialize["deletedRecords"] = o.DeletedRecords
	}
	if !IsNil(o.UpdatedRRSets) {
		toSerialize["updatedRRSets"] = o.UpdatedRRSets
	}
	if !IsNil(o.UpdatedRecords) {
		toSerialize["updatedRecords"] = o.UpdatedRecords
	}
	return toSerialize, nil
}

type NullableImportSummary struct {
	value *ImportSummary
	isSet bool
}

func (v NullableImportSummary) Get() *ImportSummary {
	return v.value
}

func (v *NullableImportSummary) Set(val *ImportSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableImportSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableImportSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportSummary(val *ImportSummary) *NullableImportSummary {
	return &NullableImportSummary{value: val, isSet: true}
}

func (v NullableImportSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
