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

// checks if the ImportRecordSetsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportRecordSetsResponse{}

// ImportRecordSetsResponse ImportSummaryResponse is the response of the import.
type ImportRecordSetsResponse struct {
	Message *string `json:"message,omitempty"`
	// REQUIRED
	Summary *ImportSummary `json:"summary"`
}

type _ImportRecordSetsResponse ImportRecordSetsResponse

// NewImportRecordSetsResponse instantiates a new ImportRecordSetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportRecordSetsResponse(summary *ImportSummary) *ImportRecordSetsResponse {
	this := ImportRecordSetsResponse{}
	this.Summary = summary
	return &this
}

// NewImportRecordSetsResponseWithDefaults instantiates a new ImportRecordSetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportRecordSetsResponseWithDefaults() *ImportRecordSetsResponse {
	this := ImportRecordSetsResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ImportRecordSetsResponse) GetMessage() *string {
	if o == nil || IsNil(o.Message) {
		var ret *string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRecordSetsResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ImportRecordSetsResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ImportRecordSetsResponse) SetMessage(v *string) {
	o.Message = v
}

// GetSummary returns the Summary field value
func (o *ImportRecordSetsResponse) GetSummary() *ImportSummary {
	if o == nil {
		var ret *ImportSummary
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ImportRecordSetsResponse) GetSummaryOk() (*ImportSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Summary, true
}

// SetSummary sets field value
func (o *ImportRecordSetsResponse) SetSummary(v *ImportSummary) {
	o.Summary = v
}

func (o ImportRecordSetsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["summary"] = o.Summary
	return toSerialize, nil
}

type NullableImportRecordSetsResponse struct {
	value *ImportRecordSetsResponse
	isSet bool
}

func (v NullableImportRecordSetsResponse) Get() *ImportRecordSetsResponse {
	return v.value
}

func (v *NullableImportRecordSetsResponse) Set(val *ImportRecordSetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableImportRecordSetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableImportRecordSetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportRecordSetsResponse(val *ImportRecordSetsResponse) *NullableImportRecordSetsResponse {
	return &NullableImportRecordSetsResponse{value: val, isSet: true}
}

func (v NullableImportRecordSetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportRecordSetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
