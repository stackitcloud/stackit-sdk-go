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

// checks if the UpdateAlertConfigsPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAlertConfigsPayload{}

// UpdateAlertConfigsPayload Alert config
type UpdateAlertConfigsPayload struct {
	Global       *UpdateAlertConfigsPayloadGlobal       `json:"global,omitempty"`
	InhibitRules *UpdateAlertConfigsPayloadInhibitRules `json:"inhibitRules,omitempty"`
	// A list of notification receivers.
	// REQUIRED
	Receivers *[]UpdateAlertConfigsPayloadReceiversInner `json:"receivers"`
	// REQUIRED
	Route *UpdateAlertConfigsPayloadRoute `json:"route"`
}

type _UpdateAlertConfigsPayload UpdateAlertConfigsPayload

// NewUpdateAlertConfigsPayload instantiates a new UpdateAlertConfigsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAlertConfigsPayload(receivers *[]UpdateAlertConfigsPayloadReceiversInner, route *UpdateAlertConfigsPayloadRoute) *UpdateAlertConfigsPayload {
	this := UpdateAlertConfigsPayload{}
	this.Receivers = receivers
	this.Route = route
	return &this
}

// NewUpdateAlertConfigsPayloadWithDefaults instantiates a new UpdateAlertConfigsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAlertConfigsPayloadWithDefaults() *UpdateAlertConfigsPayload {
	this := UpdateAlertConfigsPayload{}
	return &this
}

// GetGlobal returns the Global field value if set, zero value otherwise.
func (o *UpdateAlertConfigsPayload) GetGlobal() *UpdateAlertConfigsPayloadGlobal {
	if o == nil || IsNil(o.Global) {
		var ret *UpdateAlertConfigsPayloadGlobal
		return ret
	}
	return o.Global
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAlertConfigsPayload) GetGlobalOk() (*UpdateAlertConfigsPayloadGlobal, bool) {
	if o == nil || IsNil(o.Global) {
		return nil, false
	}
	return o.Global, true
}

// HasGlobal returns a boolean if a field has been set.
func (o *UpdateAlertConfigsPayload) HasGlobal() bool {
	if o != nil && !IsNil(o.Global) {
		return true
	}

	return false
}

// SetGlobal gets a reference to the given UpdateAlertConfigsPayloadGlobal and assigns it to the Global field.
func (o *UpdateAlertConfigsPayload) SetGlobal(v *UpdateAlertConfigsPayloadGlobal) {
	o.Global = v
}

// GetInhibitRules returns the InhibitRules field value if set, zero value otherwise.
func (o *UpdateAlertConfigsPayload) GetInhibitRules() *UpdateAlertConfigsPayloadInhibitRules {
	if o == nil || IsNil(o.InhibitRules) {
		var ret *UpdateAlertConfigsPayloadInhibitRules
		return ret
	}
	return o.InhibitRules
}

// GetInhibitRulesOk returns a tuple with the InhibitRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAlertConfigsPayload) GetInhibitRulesOk() (*UpdateAlertConfigsPayloadInhibitRules, bool) {
	if o == nil || IsNil(o.InhibitRules) {
		return nil, false
	}
	return o.InhibitRules, true
}

// HasInhibitRules returns a boolean if a field has been set.
func (o *UpdateAlertConfigsPayload) HasInhibitRules() bool {
	if o != nil && !IsNil(o.InhibitRules) {
		return true
	}

	return false
}

// SetInhibitRules gets a reference to the given UpdateAlertConfigsPayloadInhibitRules and assigns it to the InhibitRules field.
func (o *UpdateAlertConfigsPayload) SetInhibitRules(v *UpdateAlertConfigsPayloadInhibitRules) {
	o.InhibitRules = v
}

// GetReceivers returns the Receivers field value
func (o *UpdateAlertConfigsPayload) GetReceivers() *[]UpdateAlertConfigsPayloadReceiversInner {
	if o == nil || IsNil(o.Receivers) {
		var ret *[]UpdateAlertConfigsPayloadReceiversInner
		return ret
	}

	return o.Receivers
}

// GetReceiversOk returns a tuple with the Receivers field value
// and a boolean to check if the value has been set.
func (o *UpdateAlertConfigsPayload) GetReceiversOk() (*[]UpdateAlertConfigsPayloadReceiversInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Receivers, true
}

// SetReceivers sets field value
func (o *UpdateAlertConfigsPayload) SetReceivers(v *[]UpdateAlertConfigsPayloadReceiversInner) {
	o.Receivers = v
}

// GetRoute returns the Route field value
func (o *UpdateAlertConfigsPayload) GetRoute() *UpdateAlertConfigsPayloadRoute {
	if o == nil || IsNil(o.Route) {
		var ret *UpdateAlertConfigsPayloadRoute
		return ret
	}

	return o.Route
}

// GetRouteOk returns a tuple with the Route field value
// and a boolean to check if the value has been set.
func (o *UpdateAlertConfigsPayload) GetRouteOk() (*UpdateAlertConfigsPayloadRoute, bool) {
	if o == nil {
		return nil, false
	}
	return o.Route, true
}

// SetRoute sets field value
func (o *UpdateAlertConfigsPayload) SetRoute(v *UpdateAlertConfigsPayloadRoute) {
	o.Route = v
}

func (o UpdateAlertConfigsPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Global) {
		toSerialize["global"] = o.Global
	}
	if !IsNil(o.InhibitRules) {
		toSerialize["inhibitRules"] = o.InhibitRules
	}
	toSerialize["receivers"] = o.Receivers
	toSerialize["route"] = o.Route
	return toSerialize, nil
}

type NullableUpdateAlertConfigsPayload struct {
	value *UpdateAlertConfigsPayload
	isSet bool
}

func (v NullableUpdateAlertConfigsPayload) Get() *UpdateAlertConfigsPayload {
	return v.value
}

func (v *NullableUpdateAlertConfigsPayload) Set(val *UpdateAlertConfigsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAlertConfigsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAlertConfigsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAlertConfigsPayload(val *UpdateAlertConfigsPayload) *NullableUpdateAlertConfigsPayload {
	return &NullableUpdateAlertConfigsPayload{value: val, isSet: true}
}

func (v NullableUpdateAlertConfigsPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAlertConfigsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
