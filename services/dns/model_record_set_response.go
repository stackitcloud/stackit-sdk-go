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

// checks if the RecordSetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordSetResponse{}

/*
	types and functions for message
*/

// isNotNullableString
type RecordSetResponseGetMessageAttributeType = *string

func getRecordSetResponseGetMessageAttributeTypeOk(arg RecordSetResponseGetMessageAttributeType) (ret RecordSetResponseGetMessageRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetResponseGetMessageAttributeType(arg *RecordSetResponseGetMessageAttributeType, val RecordSetResponseGetMessageRetType) {
	*arg = &val
}

type RecordSetResponseGetMessageArgType = string
type RecordSetResponseGetMessageRetType = string

/*
	types and functions for rrset
*/

// isModel
type RecordSetResponseGetRrsetAttributeType = *RecordSet
type RecordSetResponseGetRrsetArgType = RecordSet
type RecordSetResponseGetRrsetRetType = RecordSet

func getRecordSetResponseGetRrsetAttributeTypeOk(arg RecordSetResponseGetRrsetAttributeType) (ret RecordSetResponseGetRrsetRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setRecordSetResponseGetRrsetAttributeType(arg *RecordSetResponseGetRrsetAttributeType, val RecordSetResponseGetRrsetRetType) {
	*arg = &val
}

// RecordSetResponse ResponseRRSet for rr set info.
type RecordSetResponse struct {
	Message RecordSetResponseGetMessageAttributeType `json:"message,omitempty"`
	// REQUIRED
	Rrset RecordSetResponseGetRrsetAttributeType `json:"rrset"`
}

type _RecordSetResponse RecordSetResponse

// NewRecordSetResponse instantiates a new RecordSetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordSetResponse(rrset RecordSetResponseGetRrsetArgType) *RecordSetResponse {
	this := RecordSetResponse{}
	setRecordSetResponseGetRrsetAttributeType(&this.Rrset, rrset)
	return &this
}

// NewRecordSetResponseWithDefaults instantiates a new RecordSetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordSetResponseWithDefaults() *RecordSetResponse {
	this := RecordSetResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RecordSetResponse) GetMessage() (res RecordSetResponseGetMessageRetType) {
	res, _ = o.GetMessageOk()
	return
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordSetResponse) GetMessageOk() (ret RecordSetResponseGetMessageRetType, ok bool) {
	return getRecordSetResponseGetMessageAttributeTypeOk(o.Message)
}

// HasMessage returns a boolean if a field has been set.
func (o *RecordSetResponse) HasMessage() bool {
	_, ok := o.GetMessageOk()
	return ok
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RecordSetResponse) SetMessage(v RecordSetResponseGetMessageRetType) {
	setRecordSetResponseGetMessageAttributeType(&o.Message, v)
}

// GetRrset returns the Rrset field value
func (o *RecordSetResponse) GetRrset() (ret RecordSetResponseGetRrsetRetType) {
	ret, _ = o.GetRrsetOk()
	return ret
}

// GetRrsetOk returns a tuple with the Rrset field value
// and a boolean to check if the value has been set.
func (o *RecordSetResponse) GetRrsetOk() (ret RecordSetResponseGetRrsetRetType, ok bool) {
	return getRecordSetResponseGetRrsetAttributeTypeOk(o.Rrset)
}

// SetRrset sets field value
func (o *RecordSetResponse) SetRrset(v RecordSetResponseGetRrsetRetType) {
	setRecordSetResponseGetRrsetAttributeType(&o.Rrset, v)
}

func (o RecordSetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getRecordSetResponseGetMessageAttributeTypeOk(o.Message); ok {
		toSerialize["Message"] = val
	}
	if val, ok := getRecordSetResponseGetRrsetAttributeTypeOk(o.Rrset); ok {
		toSerialize["Rrset"] = val
	}
	return toSerialize, nil
}

type NullableRecordSetResponse struct {
	value *RecordSetResponse
	isSet bool
}

func (v NullableRecordSetResponse) Get() *RecordSetResponse {
	return v.value
}

func (v *NullableRecordSetResponse) Set(val *RecordSetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordSetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordSetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordSetResponse(val *RecordSetResponse) *NullableRecordSetResponse {
	return &NullableRecordSetResponse{value: val, isSet: true}
}

func (v NullableRecordSetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordSetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
