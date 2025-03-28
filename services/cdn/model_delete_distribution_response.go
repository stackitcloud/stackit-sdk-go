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

// checks if the DeleteDistributionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteDistributionResponse{}

/*
	types and functions for distribution
*/

// isModel
type DeleteDistributionResponseGetDistributionAttributeType = *Distribution
type DeleteDistributionResponseGetDistributionArgType = Distribution
type DeleteDistributionResponseGetDistributionRetType = Distribution

func getDeleteDistributionResponseGetDistributionAttributeTypeOk(arg DeleteDistributionResponseGetDistributionAttributeType) (ret DeleteDistributionResponseGetDistributionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setDeleteDistributionResponseGetDistributionAttributeType(arg *DeleteDistributionResponseGetDistributionAttributeType, val DeleteDistributionResponseGetDistributionRetType) {
	*arg = &val
}

// DeleteDistributionResponse struct for DeleteDistributionResponse
type DeleteDistributionResponse struct {
	Distribution DeleteDistributionResponseGetDistributionAttributeType `json:"distribution,omitempty"`
}

// NewDeleteDistributionResponse instantiates a new DeleteDistributionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDistributionResponse() *DeleteDistributionResponse {
	this := DeleteDistributionResponse{}
	return &this
}

// NewDeleteDistributionResponseWithDefaults instantiates a new DeleteDistributionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDistributionResponseWithDefaults() *DeleteDistributionResponse {
	this := DeleteDistributionResponse{}
	return &this
}

// GetDistribution returns the Distribution field value if set, zero value otherwise.
func (o *DeleteDistributionResponse) GetDistribution() (res DeleteDistributionResponseGetDistributionRetType) {
	res, _ = o.GetDistributionOk()
	return
}

// GetDistributionOk returns a tuple with the Distribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDistributionResponse) GetDistributionOk() (ret DeleteDistributionResponseGetDistributionRetType, ok bool) {
	return getDeleteDistributionResponseGetDistributionAttributeTypeOk(o.Distribution)
}

// HasDistribution returns a boolean if a field has been set.
func (o *DeleteDistributionResponse) HasDistribution() bool {
	_, ok := o.GetDistributionOk()
	return ok
}

// SetDistribution gets a reference to the given Distribution and assigns it to the Distribution field.
func (o *DeleteDistributionResponse) SetDistribution(v DeleteDistributionResponseGetDistributionRetType) {
	setDeleteDistributionResponseGetDistributionAttributeType(&o.Distribution, v)
}

func (o DeleteDistributionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getDeleteDistributionResponseGetDistributionAttributeTypeOk(o.Distribution); ok {
		toSerialize["Distribution"] = val
	}
	return toSerialize, nil
}

type NullableDeleteDistributionResponse struct {
	value *DeleteDistributionResponse
	isSet bool
}

func (v NullableDeleteDistributionResponse) Get() *DeleteDistributionResponse {
	return v.value
}

func (v *NullableDeleteDistributionResponse) Set(val *DeleteDistributionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDistributionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDistributionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDistributionResponse(val *DeleteDistributionResponse) *NullableDeleteDistributionResponse {
	return &NullableDeleteDistributionResponse{value: val, isSet: true}
}

func (v NullableDeleteDistributionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDistributionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
