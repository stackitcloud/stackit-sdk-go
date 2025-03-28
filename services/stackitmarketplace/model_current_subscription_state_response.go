/*
STACKIT Marketplace API

API to manage STACKIT Marketplace.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackitmarketplace

import (
	"encoding/json"
)

// checks if the CurrentSubscriptionStateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrentSubscriptionStateResponse{}

/*
	types and functions for lifecycleState
*/

// isEnumRef
type CurrentSubscriptionStateResponseGetLifecycleStateAttributeType = *string
type CurrentSubscriptionStateResponseGetLifecycleStateArgType = string
type CurrentSubscriptionStateResponseGetLifecycleStateRetType = string

func getCurrentSubscriptionStateResponseGetLifecycleStateAttributeTypeOk(arg CurrentSubscriptionStateResponseGetLifecycleStateAttributeType) (ret CurrentSubscriptionStateResponseGetLifecycleStateRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCurrentSubscriptionStateResponseGetLifecycleStateAttributeType(arg *CurrentSubscriptionStateResponseGetLifecycleStateAttributeType, val CurrentSubscriptionStateResponseGetLifecycleStateRetType) {
	*arg = &val
}

/*
	types and functions for subscriptionId
*/

// isNotNullableString
type CurrentSubscriptionStateResponseGetSubscriptionIdAttributeType = *string

func getCurrentSubscriptionStateResponseGetSubscriptionIdAttributeTypeOk(arg CurrentSubscriptionStateResponseGetSubscriptionIdAttributeType) (ret CurrentSubscriptionStateResponseGetSubscriptionIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCurrentSubscriptionStateResponseGetSubscriptionIdAttributeType(arg *CurrentSubscriptionStateResponseGetSubscriptionIdAttributeType, val CurrentSubscriptionStateResponseGetSubscriptionIdRetType) {
	*arg = &val
}

type CurrentSubscriptionStateResponseGetSubscriptionIdArgType = string
type CurrentSubscriptionStateResponseGetSubscriptionIdRetType = string

// CurrentSubscriptionStateResponse struct for CurrentSubscriptionStateResponse
type CurrentSubscriptionStateResponse struct {
	// Lifecycle state of the subscription.
	// REQUIRED
	LifecycleState CurrentSubscriptionStateResponseGetLifecycleStateAttributeType `json:"lifecycleState"`
	// The subscription ID.
	// REQUIRED
	SubscriptionId CurrentSubscriptionStateResponseGetSubscriptionIdAttributeType `json:"subscriptionId"`
}

type _CurrentSubscriptionStateResponse CurrentSubscriptionStateResponse

// NewCurrentSubscriptionStateResponse instantiates a new CurrentSubscriptionStateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentSubscriptionStateResponse(lifecycleState CurrentSubscriptionStateResponseGetLifecycleStateArgType, subscriptionId CurrentSubscriptionStateResponseGetSubscriptionIdArgType) *CurrentSubscriptionStateResponse {
	this := CurrentSubscriptionStateResponse{}
	setCurrentSubscriptionStateResponseGetLifecycleStateAttributeType(&this.LifecycleState, lifecycleState)
	setCurrentSubscriptionStateResponseGetSubscriptionIdAttributeType(&this.SubscriptionId, subscriptionId)
	return &this
}

// NewCurrentSubscriptionStateResponseWithDefaults instantiates a new CurrentSubscriptionStateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentSubscriptionStateResponseWithDefaults() *CurrentSubscriptionStateResponse {
	this := CurrentSubscriptionStateResponse{}
	return &this
}

// GetLifecycleState returns the LifecycleState field value
func (o *CurrentSubscriptionStateResponse) GetLifecycleState() (ret CurrentSubscriptionStateResponseGetLifecycleStateRetType) {
	ret, _ = o.GetLifecycleStateOk()
	return ret
}

// GetLifecycleStateOk returns a tuple with the LifecycleState field value
// and a boolean to check if the value has been set.
func (o *CurrentSubscriptionStateResponse) GetLifecycleStateOk() (ret CurrentSubscriptionStateResponseGetLifecycleStateRetType, ok bool) {
	return getCurrentSubscriptionStateResponseGetLifecycleStateAttributeTypeOk(o.LifecycleState)
}

// SetLifecycleState sets field value
func (o *CurrentSubscriptionStateResponse) SetLifecycleState(v CurrentSubscriptionStateResponseGetLifecycleStateRetType) {
	setCurrentSubscriptionStateResponseGetLifecycleStateAttributeType(&o.LifecycleState, v)
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *CurrentSubscriptionStateResponse) GetSubscriptionId() (ret CurrentSubscriptionStateResponseGetSubscriptionIdRetType) {
	ret, _ = o.GetSubscriptionIdOk()
	return ret
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *CurrentSubscriptionStateResponse) GetSubscriptionIdOk() (ret CurrentSubscriptionStateResponseGetSubscriptionIdRetType, ok bool) {
	return getCurrentSubscriptionStateResponseGetSubscriptionIdAttributeTypeOk(o.SubscriptionId)
}

// SetSubscriptionId sets field value
func (o *CurrentSubscriptionStateResponse) SetSubscriptionId(v CurrentSubscriptionStateResponseGetSubscriptionIdRetType) {
	setCurrentSubscriptionStateResponseGetSubscriptionIdAttributeType(&o.SubscriptionId, v)
}

func (o CurrentSubscriptionStateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCurrentSubscriptionStateResponseGetLifecycleStateAttributeTypeOk(o.LifecycleState); ok {
		toSerialize["LifecycleState"] = val
	}
	if val, ok := getCurrentSubscriptionStateResponseGetSubscriptionIdAttributeTypeOk(o.SubscriptionId); ok {
		toSerialize["SubscriptionId"] = val
	}
	return toSerialize, nil
}

type NullableCurrentSubscriptionStateResponse struct {
	value *CurrentSubscriptionStateResponse
	isSet bool
}

func (v NullableCurrentSubscriptionStateResponse) Get() *CurrentSubscriptionStateResponse {
	return v.value
}

func (v *NullableCurrentSubscriptionStateResponse) Set(val *CurrentSubscriptionStateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentSubscriptionStateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentSubscriptionStateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentSubscriptionStateResponse(val *CurrentSubscriptionStateResponse) *NullableCurrentSubscriptionStateResponse {
	return &NullableCurrentSubscriptionStateResponse{value: val, isSet: true}
}

func (v NullableCurrentSubscriptionStateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentSubscriptionStateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
