/*
STACKIT LogMe API

The STACKIT LogMe API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logme

import (
	"encoding/json"
)

// checks if the CreateBackupResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBackupResponseItem{}

// CreateBackupResponseItem struct for CreateBackupResponseItem
type CreateBackupResponseItem struct {
	// REQUIRED
	Id *int64 `json:"id"`
	// REQUIRED
	Message *string `json:"message"`
}

type _CreateBackupResponseItem CreateBackupResponseItem

// NewCreateBackupResponseItem instantiates a new CreateBackupResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBackupResponseItem(id *int64, message *string) *CreateBackupResponseItem {
	this := CreateBackupResponseItem{}
	this.Id = id
	this.Message = message
	return &this
}

// NewCreateBackupResponseItemWithDefaults instantiates a new CreateBackupResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBackupResponseItemWithDefaults() *CreateBackupResponseItem {
	this := CreateBackupResponseItem{}
	return &this
}

// GetId returns the Id field value
func (o *CreateBackupResponseItem) GetId() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateBackupResponseItem) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *CreateBackupResponseItem) SetId(v *int64) {
	o.Id = v
}

// GetMessage returns the Message field value
func (o *CreateBackupResponseItem) GetMessage() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CreateBackupResponseItem) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *CreateBackupResponseItem) SetMessage(v *string) {
	o.Message = v
}

func (o CreateBackupResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableCreateBackupResponseItem struct {
	value *CreateBackupResponseItem
	isSet bool
}

func (v NullableCreateBackupResponseItem) Get() *CreateBackupResponseItem {
	return v.value
}

func (v *NullableCreateBackupResponseItem) Set(val *CreateBackupResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBackupResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBackupResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBackupResponseItem(val *CreateBackupResponseItem) *NullableCreateBackupResponseItem {
	return &NullableCreateBackupResponseItem{value: val, isSet: true}
}

func (v NullableCreateBackupResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBackupResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
