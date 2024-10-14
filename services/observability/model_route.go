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

// checks if the Route type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Route{}

// Route struct for Route
type Route struct {
	Continue      *bool              `json:"continue,omitempty"`
	GroupBy       *[]string          `json:"groupBy,omitempty"`
	GroupInterval *string            `json:"groupInterval,omitempty"`
	GroupWait     *string            `json:"groupWait,omitempty"`
	Match         *map[string]string `json:"match,omitempty"`
	MatchRe       *map[string]string `json:"matchRe,omitempty"`
	// REQUIRED
	Receiver       *string            `json:"receiver"`
	RepeatInterval *string            `json:"repeatInterval,omitempty"`
	Routes         *[]RouteSerializer `json:"routes,omitempty"`
}

type _Route Route

// NewRoute instantiates a new Route object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoute(receiver *string) *Route {
	this := Route{}
	var continue_ bool = false
	this.Continue = &continue_
	var groupInterval string = "5m"
	this.GroupInterval = &groupInterval
	var groupWait string = "30s"
	this.GroupWait = &groupWait
	this.Receiver = receiver
	var repeatInterval string = "4h"
	this.RepeatInterval = &repeatInterval
	return &this
}

// NewRouteWithDefaults instantiates a new Route object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteWithDefaults() *Route {
	this := Route{}
	var continue_ bool = false
	this.Continue = &continue_
	var groupInterval string = "5m"
	this.GroupInterval = &groupInterval
	var groupWait string = "30s"
	this.GroupWait = &groupWait
	var repeatInterval string = "4h"
	this.RepeatInterval = &repeatInterval
	return &this
}

// GetContinue returns the Continue field value if set, zero value otherwise.
func (o *Route) GetContinue() *bool {
	if o == nil || IsNil(o.Continue) {
		var ret *bool
		return ret
	}
	return o.Continue
}

// GetContinueOk returns a tuple with the Continue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetContinueOk() (*bool, bool) {
	if o == nil || IsNil(o.Continue) {
		return nil, false
	}
	return o.Continue, true
}

// HasContinue returns a boolean if a field has been set.
func (o *Route) HasContinue() bool {
	if o != nil && !IsNil(o.Continue) {
		return true
	}

	return false
}

// SetContinue gets a reference to the given bool and assigns it to the Continue field.
func (o *Route) SetContinue(v *bool) {
	o.Continue = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *Route) GetGroupBy() *[]string {
	if o == nil || IsNil(o.GroupBy) {
		var ret *[]string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetGroupByOk() (*[]string, bool) {
	if o == nil || IsNil(o.GroupBy) {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *Route) HasGroupBy() bool {
	if o != nil && !IsNil(o.GroupBy) {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *Route) SetGroupBy(v *[]string) {
	o.GroupBy = v
}

// GetGroupInterval returns the GroupInterval field value if set, zero value otherwise.
func (o *Route) GetGroupInterval() *string {
	if o == nil || IsNil(o.GroupInterval) {
		var ret *string
		return ret
	}
	return o.GroupInterval
}

// GetGroupIntervalOk returns a tuple with the GroupInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetGroupIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.GroupInterval) {
		return nil, false
	}
	return o.GroupInterval, true
}

// HasGroupInterval returns a boolean if a field has been set.
func (o *Route) HasGroupInterval() bool {
	if o != nil && !IsNil(o.GroupInterval) {
		return true
	}

	return false
}

// SetGroupInterval gets a reference to the given string and assigns it to the GroupInterval field.
func (o *Route) SetGroupInterval(v *string) {
	o.GroupInterval = v
}

// GetGroupWait returns the GroupWait field value if set, zero value otherwise.
func (o *Route) GetGroupWait() *string {
	if o == nil || IsNil(o.GroupWait) {
		var ret *string
		return ret
	}
	return o.GroupWait
}

// GetGroupWaitOk returns a tuple with the GroupWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetGroupWaitOk() (*string, bool) {
	if o == nil || IsNil(o.GroupWait) {
		return nil, false
	}
	return o.GroupWait, true
}

// HasGroupWait returns a boolean if a field has been set.
func (o *Route) HasGroupWait() bool {
	if o != nil && !IsNil(o.GroupWait) {
		return true
	}

	return false
}

// SetGroupWait gets a reference to the given string and assigns it to the GroupWait field.
func (o *Route) SetGroupWait(v *string) {
	o.GroupWait = v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *Route) GetMatch() *map[string]string {
	if o == nil || IsNil(o.Match) {
		var ret *map[string]string
		return ret
	}
	return o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetMatchOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *Route) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given map[string]string and assigns it to the Match field.
func (o *Route) SetMatch(v *map[string]string) {
	o.Match = v
}

// GetMatchRe returns the MatchRe field value if set, zero value otherwise.
func (o *Route) GetMatchRe() *map[string]string {
	if o == nil || IsNil(o.MatchRe) {
		var ret *map[string]string
		return ret
	}
	return o.MatchRe
}

// GetMatchReOk returns a tuple with the MatchRe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetMatchReOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.MatchRe) {
		return nil, false
	}
	return o.MatchRe, true
}

// HasMatchRe returns a boolean if a field has been set.
func (o *Route) HasMatchRe() bool {
	if o != nil && !IsNil(o.MatchRe) {
		return true
	}

	return false
}

// SetMatchRe gets a reference to the given map[string]string and assigns it to the MatchRe field.
func (o *Route) SetMatchRe(v *map[string]string) {
	o.MatchRe = v
}

// GetReceiver returns the Receiver field value
func (o *Route) GetReceiver() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value
// and a boolean to check if the value has been set.
func (o *Route) GetReceiverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Receiver, true
}

// SetReceiver sets field value
func (o *Route) SetReceiver(v *string) {
	o.Receiver = v
}

// GetRepeatInterval returns the RepeatInterval field value if set, zero value otherwise.
func (o *Route) GetRepeatInterval() *string {
	if o == nil || IsNil(o.RepeatInterval) {
		var ret *string
		return ret
	}
	return o.RepeatInterval
}

// GetRepeatIntervalOk returns a tuple with the RepeatInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetRepeatIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.RepeatInterval) {
		return nil, false
	}
	return o.RepeatInterval, true
}

// HasRepeatInterval returns a boolean if a field has been set.
func (o *Route) HasRepeatInterval() bool {
	if o != nil && !IsNil(o.RepeatInterval) {
		return true
	}

	return false
}

// SetRepeatInterval gets a reference to the given string and assigns it to the RepeatInterval field.
func (o *Route) SetRepeatInterval(v *string) {
	o.RepeatInterval = v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *Route) GetRoutes() *[]RouteSerializer {
	if o == nil || IsNil(o.Routes) {
		var ret *[]RouteSerializer
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetRoutesOk() (*[]RouteSerializer, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *Route) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []RouteSerializer and assigns it to the Routes field.
func (o *Route) SetRoutes(v *[]RouteSerializer) {
	o.Routes = v
}

func (o Route) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Continue) {
		toSerialize["continue"] = o.Continue
	}
	if !IsNil(o.GroupBy) {
		toSerialize["groupBy"] = o.GroupBy
	}
	if !IsNil(o.GroupInterval) {
		toSerialize["groupInterval"] = o.GroupInterval
	}
	if !IsNil(o.GroupWait) {
		toSerialize["groupWait"] = o.GroupWait
	}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	if !IsNil(o.MatchRe) {
		toSerialize["matchRe"] = o.MatchRe
	}
	toSerialize["receiver"] = o.Receiver
	if !IsNil(o.RepeatInterval) {
		toSerialize["repeatInterval"] = o.RepeatInterval
	}
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	return toSerialize, nil
}

type NullableRoute struct {
	value *Route
	isSet bool
}

func (v NullableRoute) Get() *Route {
	return v.value
}

func (v *NullableRoute) Set(val *Route) {
	v.value = val
	v.isSet = true
}

func (v NullableRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoute(val *Route) *NullableRoute {
	return &NullableRoute{value: val, isSet: true}
}

func (v NullableRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
