/*
CDN API

API used to create and manage your CDN distributions.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
	"time"
)

// checks if the GetCacheInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCacheInfoResponse{}

/*
	types and functions for history
*/

// isArray
type GetCacheInfoResponseGetHistoryAttributeType = *[]GetCacheInfoResponseHistoryEntry
type GetCacheInfoResponseGetHistoryArgType = []GetCacheInfoResponseHistoryEntry
type GetCacheInfoResponseGetHistoryRetType = []GetCacheInfoResponseHistoryEntry

func getGetCacheInfoResponseGetHistoryAttributeTypeOk(arg GetCacheInfoResponseGetHistoryAttributeType) (ret GetCacheInfoResponseGetHistoryRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGetCacheInfoResponseGetHistoryAttributeType(arg *GetCacheInfoResponseGetHistoryAttributeType, val GetCacheInfoResponseGetHistoryRetType) {
	*arg = &val
}

/*
	types and functions for lastPurgeTime
*/

// isDateTime
type GetCacheInfoResponseGetLastPurgeTimeAttributeType = *time.Time
type GetCacheInfoResponseGetLastPurgeTimeArgType = *time.Time
type GetCacheInfoResponseGetLastPurgeTimeRetType = *time.Time

func getGetCacheInfoResponseGetLastPurgeTimeAttributeTypeOk(arg GetCacheInfoResponseGetLastPurgeTimeAttributeType) (ret GetCacheInfoResponseGetLastPurgeTimeRetType, ok bool) {
	if arg == nil {
		return nil, false
	}
	return arg, true
}

func setGetCacheInfoResponseGetLastPurgeTimeAttributeType(arg *GetCacheInfoResponseGetLastPurgeTimeAttributeType, val GetCacheInfoResponseGetLastPurgeTimeRetType) {
	*arg = val
}

// GetCacheInfoResponse struct for GetCacheInfoResponse
type GetCacheInfoResponse struct {
	// REQUIRED
	History GetCacheInfoResponseGetHistoryAttributeType `json:"history" required:"true"`
	// Returns the last time the cache was purged by calling the PurgeCache endpoint.  Time represented as RFC3339 compliant string. If the cache was never purged, this returns `null`
	// REQUIRED
	LastPurgeTime GetCacheInfoResponseGetLastPurgeTimeAttributeType `json:"lastPurgeTime" required:"true"`
}

type _GetCacheInfoResponse GetCacheInfoResponse

// NewGetCacheInfoResponse instantiates a new GetCacheInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCacheInfoResponse(history GetCacheInfoResponseGetHistoryArgType, lastPurgeTime GetCacheInfoResponseGetLastPurgeTimeArgType) *GetCacheInfoResponse {
	this := GetCacheInfoResponse{}
	setGetCacheInfoResponseGetHistoryAttributeType(&this.History, history)
	setGetCacheInfoResponseGetLastPurgeTimeAttributeType(&this.LastPurgeTime, lastPurgeTime)
	return &this
}

// NewGetCacheInfoResponseWithDefaults instantiates a new GetCacheInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCacheInfoResponseWithDefaults() *GetCacheInfoResponse {
	this := GetCacheInfoResponse{}
	return &this
}

// GetHistory returns the History field value
func (o *GetCacheInfoResponse) GetHistory() (ret GetCacheInfoResponseGetHistoryRetType) {
	ret, _ = o.GetHistoryOk()
	return ret
}

// GetHistoryOk returns a tuple with the History field value
// and a boolean to check if the value has been set.
func (o *GetCacheInfoResponse) GetHistoryOk() (ret GetCacheInfoResponseGetHistoryRetType, ok bool) {
	return getGetCacheInfoResponseGetHistoryAttributeTypeOk(o.History)
}

// SetHistory sets field value
func (o *GetCacheInfoResponse) SetHistory(v GetCacheInfoResponseGetHistoryRetType) {
	setGetCacheInfoResponseGetHistoryAttributeType(&o.History, v)
}

// GetLastPurgeTime returns the LastPurgeTime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *GetCacheInfoResponse) GetLastPurgeTime() (ret GetCacheInfoResponseGetLastPurgeTimeRetType) {
	ret, _ = o.GetLastPurgeTimeOk()
	return ret
}

// GetLastPurgeTimeOk returns a tuple with the LastPurgeTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetCacheInfoResponse) GetLastPurgeTimeOk() (ret GetCacheInfoResponseGetLastPurgeTimeRetType, ok bool) {
	return getGetCacheInfoResponseGetLastPurgeTimeAttributeTypeOk(o.LastPurgeTime)
}

// SetLastPurgeTime sets field value
func (o *GetCacheInfoResponse) SetLastPurgeTime(v GetCacheInfoResponseGetLastPurgeTimeRetType) {
	setGetCacheInfoResponseGetLastPurgeTimeAttributeType(&o.LastPurgeTime, v)
}

func (o GetCacheInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getGetCacheInfoResponseGetHistoryAttributeTypeOk(o.History); ok {
		toSerialize["History"] = val
	}
	if val, ok := getGetCacheInfoResponseGetLastPurgeTimeAttributeTypeOk(o.LastPurgeTime); ok {
		toSerialize["LastPurgeTime"] = val
	}
	return toSerialize, nil
}

type NullableGetCacheInfoResponse struct {
	value *GetCacheInfoResponse
	isSet bool
}

func (v NullableGetCacheInfoResponse) Get() *GetCacheInfoResponse {
	return v.value
}

func (v *NullableGetCacheInfoResponse) Set(val *GetCacheInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCacheInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCacheInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCacheInfoResponse(val *GetCacheInfoResponse) *NullableGetCacheInfoResponse {
	return &NullableGetCacheInfoResponse{value: val, isSet: true}
}

func (v NullableGetCacheInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCacheInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
