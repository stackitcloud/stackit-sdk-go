/*
STACKIT Opensearch API

The STACKIT Opensearch API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opensearch

import (
	"encoding/json"
)

// checks if the ListOfferingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOfferingsResponse{}

/*
	types and functions for offerings
*/

// isArray
type ListOfferingsResponseGetOfferingsAttributeType = *[]Offering
type ListOfferingsResponseGetOfferingsArgType = []Offering
type ListOfferingsResponseGetOfferingsRetType = []Offering

func getListOfferingsResponseGetOfferingsAttributeTypeOk(arg ListOfferingsResponseGetOfferingsAttributeType) (ret ListOfferingsResponseGetOfferingsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListOfferingsResponseGetOfferingsAttributeType(arg *ListOfferingsResponseGetOfferingsAttributeType, val ListOfferingsResponseGetOfferingsRetType) {
	*arg = &val
}

// ListOfferingsResponse struct for ListOfferingsResponse
type ListOfferingsResponse struct {
	// REQUIRED
	Offerings ListOfferingsResponseGetOfferingsAttributeType `json:"offerings"`
}

type _ListOfferingsResponse ListOfferingsResponse

// NewListOfferingsResponse instantiates a new ListOfferingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOfferingsResponse(offerings ListOfferingsResponseGetOfferingsArgType) *ListOfferingsResponse {
	this := ListOfferingsResponse{}
	setListOfferingsResponseGetOfferingsAttributeType(&this.Offerings, offerings)
	return &this
}

// NewListOfferingsResponseWithDefaults instantiates a new ListOfferingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOfferingsResponseWithDefaults() *ListOfferingsResponse {
	this := ListOfferingsResponse{}
	return &this
}

// GetOfferings returns the Offerings field value
func (o *ListOfferingsResponse) GetOfferings() (ret ListOfferingsResponseGetOfferingsRetType) {
	ret, _ = o.GetOfferingsOk()
	return ret
}

// GetOfferingsOk returns a tuple with the Offerings field value
// and a boolean to check if the value has been set.
func (o *ListOfferingsResponse) GetOfferingsOk() (ret ListOfferingsResponseGetOfferingsRetType, ok bool) {
	return getListOfferingsResponseGetOfferingsAttributeTypeOk(o.Offerings)
}

// SetOfferings sets field value
func (o *ListOfferingsResponse) SetOfferings(v ListOfferingsResponseGetOfferingsRetType) {
	setListOfferingsResponseGetOfferingsAttributeType(&o.Offerings, v)
}

func (o ListOfferingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListOfferingsResponseGetOfferingsAttributeTypeOk(o.Offerings); ok {
		toSerialize["Offerings"] = val
	}
	return toSerialize, nil
}

type NullableListOfferingsResponse struct {
	value *ListOfferingsResponse
	isSet bool
}

func (v NullableListOfferingsResponse) Get() *ListOfferingsResponse {
	return v.value
}

func (v *NullableListOfferingsResponse) Set(val *ListOfferingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOfferingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOfferingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOfferingsResponse(val *ListOfferingsResponse) *NullableListOfferingsResponse {
	return &NullableListOfferingsResponse{value: val, isSet: true}
}

func (v NullableListOfferingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOfferingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
