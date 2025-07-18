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

// checks if the WebHook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebHook{}

/*
	types and functions for msTeams
*/

// isBoolean
type WebHookgetMsTeamsAttributeType = *bool
type WebHookgetMsTeamsArgType = bool
type WebHookgetMsTeamsRetType = bool

func getWebHookgetMsTeamsAttributeTypeOk(arg WebHookgetMsTeamsAttributeType) (ret WebHookgetMsTeamsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setWebHookgetMsTeamsAttributeType(arg *WebHookgetMsTeamsAttributeType, val WebHookgetMsTeamsRetType) {
	*arg = &val
}

/*
	types and functions for sendResolved
*/

// isBoolean
type WebHookgetSendResolvedAttributeType = *bool
type WebHookgetSendResolvedArgType = bool
type WebHookgetSendResolvedRetType = bool

func getWebHookgetSendResolvedAttributeTypeOk(arg WebHookgetSendResolvedAttributeType) (ret WebHookgetSendResolvedRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setWebHookgetSendResolvedAttributeType(arg *WebHookgetSendResolvedAttributeType, val WebHookgetSendResolvedRetType) {
	*arg = &val
}

/*
	types and functions for url
*/

// isNotNullableString
type WebHookGetUrlAttributeType = *string

func getWebHookGetUrlAttributeTypeOk(arg WebHookGetUrlAttributeType) (ret WebHookGetUrlRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setWebHookGetUrlAttributeType(arg *WebHookGetUrlAttributeType, val WebHookGetUrlRetType) {
	*arg = &val
}

type WebHookGetUrlArgType = string
type WebHookGetUrlRetType = string

// WebHook struct for WebHook
type WebHook struct {
	MsTeams      WebHookgetMsTeamsAttributeType      `json:"msTeams,omitempty"`
	SendResolved WebHookgetSendResolvedAttributeType `json:"sendResolved,omitempty"`
	// REQUIRED
	Url WebHookGetUrlAttributeType `json:"url" required:"true"`
}

type _WebHook WebHook

// NewWebHook instantiates a new WebHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebHook(url WebHookGetUrlArgType) *WebHook {
	this := WebHook{}
	setWebHookGetUrlAttributeType(&this.Url, url)
	return &this
}

// NewWebHookWithDefaults instantiates a new WebHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebHookWithDefaults() *WebHook {
	this := WebHook{}
	var msTeams bool = false
	this.MsTeams = &msTeams
	var sendResolved bool = true
	this.SendResolved = &sendResolved
	return &this
}

// GetMsTeams returns the MsTeams field value if set, zero value otherwise.
func (o *WebHook) GetMsTeams() (res WebHookgetMsTeamsRetType) {
	res, _ = o.GetMsTeamsOk()
	return
}

// GetMsTeamsOk returns a tuple with the MsTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHook) GetMsTeamsOk() (ret WebHookgetMsTeamsRetType, ok bool) {
	return getWebHookgetMsTeamsAttributeTypeOk(o.MsTeams)
}

// HasMsTeams returns a boolean if a field has been set.
func (o *WebHook) HasMsTeams() bool {
	_, ok := o.GetMsTeamsOk()
	return ok
}

// SetMsTeams gets a reference to the given bool and assigns it to the MsTeams field.
func (o *WebHook) SetMsTeams(v WebHookgetMsTeamsRetType) {
	setWebHookgetMsTeamsAttributeType(&o.MsTeams, v)
}

// GetSendResolved returns the SendResolved field value if set, zero value otherwise.
func (o *WebHook) GetSendResolved() (res WebHookgetSendResolvedRetType) {
	res, _ = o.GetSendResolvedOk()
	return
}

// GetSendResolvedOk returns a tuple with the SendResolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHook) GetSendResolvedOk() (ret WebHookgetSendResolvedRetType, ok bool) {
	return getWebHookgetSendResolvedAttributeTypeOk(o.SendResolved)
}

// HasSendResolved returns a boolean if a field has been set.
func (o *WebHook) HasSendResolved() bool {
	_, ok := o.GetSendResolvedOk()
	return ok
}

// SetSendResolved gets a reference to the given bool and assigns it to the SendResolved field.
func (o *WebHook) SetSendResolved(v WebHookgetSendResolvedRetType) {
	setWebHookgetSendResolvedAttributeType(&o.SendResolved, v)
}

// GetUrl returns the Url field value
func (o *WebHook) GetUrl() (ret WebHookGetUrlRetType) {
	ret, _ = o.GetUrlOk()
	return ret
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebHook) GetUrlOk() (ret WebHookGetUrlRetType, ok bool) {
	return getWebHookGetUrlAttributeTypeOk(o.Url)
}

// SetUrl sets field value
func (o *WebHook) SetUrl(v WebHookGetUrlRetType) {
	setWebHookGetUrlAttributeType(&o.Url, v)
}

func (o WebHook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getWebHookgetMsTeamsAttributeTypeOk(o.MsTeams); ok {
		toSerialize["MsTeams"] = val
	}
	if val, ok := getWebHookgetSendResolvedAttributeTypeOk(o.SendResolved); ok {
		toSerialize["SendResolved"] = val
	}
	if val, ok := getWebHookGetUrlAttributeTypeOk(o.Url); ok {
		toSerialize["Url"] = val
	}
	return toSerialize, nil
}

type NullableWebHook struct {
	value *WebHook
	isSet bool
}

func (v NullableWebHook) Get() *WebHook {
	return v.value
}

func (v *NullableWebHook) Set(val *WebHook) {
	v.value = val
	v.isSet = true
}

func (v NullableWebHook) IsSet() bool {
	return v.isSet
}

func (v *NullableWebHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebHook(val *WebHook) *NullableWebHook {
	return &NullableWebHook{value: val, isSet: true}
}

func (v NullableWebHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
