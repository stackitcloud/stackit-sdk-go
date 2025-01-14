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

// checks if the CreateAlertConfigReceiverPayloadWebHookConfigsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAlertConfigReceiverPayloadWebHookConfigsInner{}

// CreateAlertConfigReceiverPayloadWebHookConfigsInner struct for CreateAlertConfigReceiverPayloadWebHookConfigsInner
type CreateAlertConfigReceiverPayloadWebHookConfigsInner struct {
	// Microsoft Teams webhooks require special handling. If you set this property to true, it is treated as such
	MsTeams *bool `json:"msTeams,omitempty"`
	// The endpoint to send HTTP POST requests to. `Additional Validators:` * must be a syntactically valid url address
	Url *string `json:"url,omitempty"`
}

// NewCreateAlertConfigReceiverPayloadWebHookConfigsInner instantiates a new CreateAlertConfigReceiverPayloadWebHookConfigsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAlertConfigReceiverPayloadWebHookConfigsInner() *CreateAlertConfigReceiverPayloadWebHookConfigsInner {
	this := CreateAlertConfigReceiverPayloadWebHookConfigsInner{}
	var msTeams bool = false
	this.MsTeams = &msTeams
	return &this
}

// NewCreateAlertConfigReceiverPayloadWebHookConfigsInnerWithDefaults instantiates a new CreateAlertConfigReceiverPayloadWebHookConfigsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAlertConfigReceiverPayloadWebHookConfigsInnerWithDefaults() *CreateAlertConfigReceiverPayloadWebHookConfigsInner {
	this := CreateAlertConfigReceiverPayloadWebHookConfigsInner{}
	var msTeams bool = false
	this.MsTeams = &msTeams
	return &this
}

// GetMsTeams returns the MsTeams field value if set, zero value otherwise.
func (o *CreateAlertConfigReceiverPayloadWebHookConfigsInner) GetMsTeams() *bool {
	if o == nil || IsNil(o.MsTeams) {
		var ret *bool
		return ret
	}
	return o.MsTeams
}

// GetMsTeamsOk returns a tuple with the MsTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadWebHookConfigsInner) GetMsTeamsOk() (*bool, bool) {
	if o == nil || IsNil(o.MsTeams) {
		return nil, false
	}
	return o.MsTeams, true
}

// HasMsTeams returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadWebHookConfigsInner) HasMsTeams() bool {
	if o != nil && !IsNil(o.MsTeams) {
		return true
	}

	return false
}

// SetMsTeams gets a reference to the given bool and assigns it to the MsTeams field.
func (o *CreateAlertConfigReceiverPayloadWebHookConfigsInner) SetMsTeams(v *bool) {
	o.MsTeams = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateAlertConfigReceiverPayloadWebHookConfigsInner) GetUrl() *string {
	if o == nil || IsNil(o.Url) {
		var ret *string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadWebHookConfigsInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadWebHookConfigsInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateAlertConfigReceiverPayloadWebHookConfigsInner) SetUrl(v *string) {
	o.Url = v
}

func (o CreateAlertConfigReceiverPayloadWebHookConfigsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MsTeams) {
		toSerialize["msTeams"] = o.MsTeams
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCreateAlertConfigReceiverPayloadWebHookConfigsInner struct {
	value *CreateAlertConfigReceiverPayloadWebHookConfigsInner
	isSet bool
}

func (v NullableCreateAlertConfigReceiverPayloadWebHookConfigsInner) Get() *CreateAlertConfigReceiverPayloadWebHookConfigsInner {
	return v.value
}

func (v *NullableCreateAlertConfigReceiverPayloadWebHookConfigsInner) Set(val *CreateAlertConfigReceiverPayloadWebHookConfigsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAlertConfigReceiverPayloadWebHookConfigsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAlertConfigReceiverPayloadWebHookConfigsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAlertConfigReceiverPayloadWebHookConfigsInner(val *CreateAlertConfigReceiverPayloadWebHookConfigsInner) *NullableCreateAlertConfigReceiverPayloadWebHookConfigsInner {
	return &NullableCreateAlertConfigReceiverPayloadWebHookConfigsInner{value: val, isSet: true}
}

func (v NullableCreateAlertConfigReceiverPayloadWebHookConfigsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAlertConfigReceiverPayloadWebHookConfigsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
