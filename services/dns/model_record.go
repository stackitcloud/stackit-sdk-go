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

// checks if the Record type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Record{}

/*
	types and functions for content
*/

// isNotNullableString
type RecordGetContentAttributeType = *string

func getRecordGetContentAttributeTypeOk(arg RecordGetContentAttributeType) (ret RecordGetContentRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordGetContentAttributeType(arg *RecordGetContentAttributeType, val RecordGetContentRetType) {
	*arg = &val
}

type RecordGetContentArgType = string
type RecordGetContentRetType = string

/*
	types and functions for id
*/

// isNotNullableString
type RecordGetIdAttributeType = *string

func getRecordGetIdAttributeTypeOk(arg RecordGetIdAttributeType) (ret RecordGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordGetIdAttributeType(arg *RecordGetIdAttributeType, val RecordGetIdRetType) {
	*arg = &val
}

type RecordGetIdArgType = string
type RecordGetIdRetType = string

// Record Record.
type Record struct {
	// content of the record
	// REQUIRED
	Content RecordGetContentAttributeType `json:"content" required:"true"`
	// rr set id
	// REQUIRED
	Id RecordGetIdAttributeType `json:"id" required:"true"`
}

type _Record Record

// NewRecord instantiates a new Record object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecord(content RecordGetContentArgType, id RecordGetIdArgType) *Record {
	this := Record{}
	setRecordGetContentAttributeType(&this.Content, content)
	setRecordGetIdAttributeType(&this.Id, id)
	return &this
}

// NewRecordWithDefaults instantiates a new Record object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordWithDefaults() *Record {
	this := Record{}
	return &this
}

// GetContent returns the Content field value
func (o *Record) GetContent() (ret RecordGetContentRetType) {
	ret, _ = o.GetContentOk()
	return ret
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *Record) GetContentOk() (ret RecordGetContentRetType, ok bool) {
	return getRecordGetContentAttributeTypeOk(o.Content)
}

// SetContent sets field value
func (o *Record) SetContent(v RecordGetContentRetType) {
	setRecordGetContentAttributeType(&o.Content, v)
}

// GetId returns the Id field value
func (o *Record) GetId() (ret RecordGetIdRetType) {
	ret, _ = o.GetIdOk()
	return ret
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Record) GetIdOk() (ret RecordGetIdRetType, ok bool) {
	return getRecordGetIdAttributeTypeOk(o.Id)
}

// SetId sets field value
func (o *Record) SetId(v RecordGetIdRetType) {
	setRecordGetIdAttributeType(&o.Id, v)
}

func (o Record) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getRecordGetContentAttributeTypeOk(o.Content); ok {
		toSerialize["Content"] = val
	}
	if val, ok := getRecordGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	return toSerialize, nil
}

type NullableRecord struct {
	value *Record
	isSet bool
}

func (v NullableRecord) Get() *Record {
	return v.value
}

func (v *NullableRecord) Set(val *Record) {
	v.value = val
	v.isSet = true
}

func (v NullableRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecord(val *Record) *NullableRecord {
	return &NullableRecord{value: val, isSet: true}
}

func (v NullableRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
