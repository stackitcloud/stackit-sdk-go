/*
CDN API

API used to create and manage your CDN distributions.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the GetStatisticsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatisticsResponse{}

// GetStatisticsResponse struct for GetStatisticsResponse
type GetStatisticsResponse struct {
	// REQUIRED
	Records *[]DistributionStatisticsRecord `json:"records"`
}

type _GetStatisticsResponse GetStatisticsResponse

// NewGetStatisticsResponse instantiates a new GetStatisticsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatisticsResponse(records *[]DistributionStatisticsRecord) *GetStatisticsResponse {
	this := GetStatisticsResponse{}
	this.Records = records
	return &this
}

// NewGetStatisticsResponseWithDefaults instantiates a new GetStatisticsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatisticsResponseWithDefaults() *GetStatisticsResponse {
	this := GetStatisticsResponse{}
	return &this
}

// GetRecords returns the Records field value
func (o *GetStatisticsResponse) GetRecords() *[]DistributionStatisticsRecord {
	if o == nil || IsNil(o.Records) {
		var ret *[]DistributionStatisticsRecord
		return ret
	}

	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *GetStatisticsResponse) GetRecordsOk() (*[]DistributionStatisticsRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.Records, true
}

// SetRecords sets field value
func (o *GetStatisticsResponse) SetRecords(v *[]DistributionStatisticsRecord) {
	o.Records = v
}

func (o GetStatisticsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["records"] = o.Records
	return toSerialize, nil
}

type NullableGetStatisticsResponse struct {
	value *GetStatisticsResponse
	isSet bool
}

func (v NullableGetStatisticsResponse) Get() *GetStatisticsResponse {
	return v.value
}

func (v *NullableGetStatisticsResponse) Set(val *GetStatisticsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatisticsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatisticsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatisticsResponse(val *GetStatisticsResponse) *NullableGetStatisticsResponse {
	return &NullableGetStatisticsResponse{value: val, isSet: true}
}

func (v NullableGetStatisticsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatisticsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
