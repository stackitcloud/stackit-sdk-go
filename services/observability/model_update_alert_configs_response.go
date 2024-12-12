/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"encoding/json"
)

// checks if the UpdateAlertConfigsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAlertConfigsResponse{}

// UpdateAlertConfigsResponse struct for UpdateAlertConfigsResponse
type UpdateAlertConfigsResponse struct {
	// REQUIRED
	Data *Alert `json:"data"`
	// REQUIRED
	Message *string `json:"message"`
}

type _UpdateAlertConfigsResponse UpdateAlertConfigsResponse

// NewUpdateAlertConfigsResponse instantiates a new UpdateAlertConfigsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAlertConfigsResponse(data *Alert, message *string) *UpdateAlertConfigsResponse {
	this := UpdateAlertConfigsResponse{}
	this.Data = data
	this.Message = message
	return &this
}

// NewUpdateAlertConfigsResponseWithDefaults instantiates a new UpdateAlertConfigsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAlertConfigsResponseWithDefaults() *UpdateAlertConfigsResponse {
	this := UpdateAlertConfigsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateAlertConfigsResponse) GetData() *Alert {
	if o == nil || IsNil(o.Data) {
		var ret *Alert
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateAlertConfigsResponse) GetDataOk() (*Alert, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *UpdateAlertConfigsResponse) SetData(v *Alert) {
	o.Data = v
}

// GetMessage returns the Message field value
func (o *UpdateAlertConfigsResponse) GetMessage() *string {
	if o == nil || IsNil(o.Message) {
		var ret *string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *UpdateAlertConfigsResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *UpdateAlertConfigsResponse) SetMessage(v *string) {
	o.Message = v
}

func (o UpdateAlertConfigsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableUpdateAlertConfigsResponse struct {
	value *UpdateAlertConfigsResponse
	isSet bool
}

func (v NullableUpdateAlertConfigsResponse) Get() *UpdateAlertConfigsResponse {
	return v.value
}

func (v *NullableUpdateAlertConfigsResponse) Set(val *UpdateAlertConfigsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAlertConfigsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAlertConfigsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAlertConfigsResponse(val *UpdateAlertConfigsResponse) *NullableUpdateAlertConfigsResponse {
	return &NullableUpdateAlertConfigsResponse{value: val, isSet: true}
}

func (v NullableUpdateAlertConfigsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAlertConfigsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
