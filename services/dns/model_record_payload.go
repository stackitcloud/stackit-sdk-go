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

// checks if the RecordPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordPayload{}

/*
	types and functions for content
*/

// isNotNullableString
type RecordPayloadGetContentAttributeType = *string

func getRecordPayloadGetContentAttributeTypeOk(arg RecordPayloadGetContentAttributeType) (ret RecordPayloadGetContentRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordPayloadGetContentAttributeType(arg *RecordPayloadGetContentAttributeType, val RecordPayloadGetContentRetType) {
	*arg = &val
}

type RecordPayloadGetContentArgType = string
type RecordPayloadGetContentRetType = string

// RecordPayload RecordPost for rr set info.
type RecordPayload struct {
	// content of the record
	// REQUIRED
	Content RecordPayloadGetContentAttributeType `json:"content" required:"true"`
}

type _RecordPayload RecordPayload

// NewRecordPayload instantiates a new RecordPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordPayload(content RecordPayloadGetContentArgType) *RecordPayload {
	this := RecordPayload{}
	setRecordPayloadGetContentAttributeType(&this.Content, content)
	return &this
}

// NewRecordPayloadWithDefaults instantiates a new RecordPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordPayloadWithDefaults() *RecordPayload {
	this := RecordPayload{}
	return &this
}

// GetContent returns the Content field value
func (o *RecordPayload) GetContent() (ret RecordPayloadGetContentRetType) {
	ret, _ = o.GetContentOk()
	return ret
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *RecordPayload) GetContentOk() (ret RecordPayloadGetContentRetType, ok bool) {
	return getRecordPayloadGetContentAttributeTypeOk(o.Content)
}

// SetContent sets field value
func (o *RecordPayload) SetContent(v RecordPayloadGetContentRetType) {
	setRecordPayloadGetContentAttributeType(&o.Content, v)
}

func (o RecordPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getRecordPayloadGetContentAttributeTypeOk(o.Content); ok {
		toSerialize["Content"] = val
	}
	return toSerialize, nil
}

type NullableRecordPayload struct {
	value *RecordPayload
	isSet bool
}

func (v NullableRecordPayload) Get() *RecordPayload {
	return v.value
}

func (v *NullableRecordPayload) Set(val *RecordPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordPayload(val *RecordPayload) *NullableRecordPayload {
	return &NullableRecordPayload{value: val, isSet: true}
}

func (v NullableRecordPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
