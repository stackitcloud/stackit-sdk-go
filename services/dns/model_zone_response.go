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

// checks if the ZoneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneResponse{}

// ZoneResponse ResponseZone for user info.
type ZoneResponse struct {
	Message *string `json:"message,omitempty"`
	// REQUIRED
	Zone *Zone `json:"zone"`
}

type _ZoneResponse ZoneResponse

// NewZoneResponse instantiates a new ZoneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneResponse(zone *Zone) *ZoneResponse {
	this := ZoneResponse{}
	this.Zone = zone
	return &this
}

// NewZoneResponseWithDefaults instantiates a new ZoneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneResponseWithDefaults() *ZoneResponse {
	this := ZoneResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ZoneResponse) GetMessage() *string {
	if o == nil || IsNil(o.Message) {
		var ret *string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ZoneResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ZoneResponse) SetMessage(v *string) {
	o.Message = v
}

// GetZone returns the Zone field value
func (o *ZoneResponse) GetZone() *Zone {
	if o == nil {
		var ret *Zone
		return ret
	}

	return o.Zone
}

// GetZoneOk returns a tuple with the Zone field value
// and a boolean to check if the value has been set.
func (o *ZoneResponse) GetZoneOk() (*Zone, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zone, true
}

// SetZone sets field value
func (o *ZoneResponse) SetZone(v *Zone) {
	o.Zone = v
}

func (o ZoneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["zone"] = o.Zone
	return toSerialize, nil
}

type NullableZoneResponse struct {
	value *ZoneResponse
	isSet bool
}

func (v NullableZoneResponse) Get() *ZoneResponse {
	return v.value
}

func (v *NullableZoneResponse) Set(val *ZoneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneResponse(val *ZoneResponse) *NullableZoneResponse {
	return &NullableZoneResponse{value: val, isSet: true}
}

func (v NullableZoneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
