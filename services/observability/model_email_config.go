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

// checks if the EmailConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailConfig{}

// EmailConfig struct for EmailConfig
type EmailConfig struct {
	AuthIdentity *string `json:"authIdentity,omitempty"`
	AuthPassword *string `json:"authPassword,omitempty"`
	AuthUsername *string `json:"authUsername,omitempty"`
	From         *string `json:"from,omitempty"`
	SendResolved *bool   `json:"sendResolved,omitempty"`
	Smarthost    *string `json:"smarthost,omitempty"`
	// REQUIRED
	To *string `json:"to"`
}

type _EmailConfig EmailConfig

// NewEmailConfig instantiates a new EmailConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailConfig(to *string) *EmailConfig {
	this := EmailConfig{}
	var sendResolved bool = false
	this.SendResolved = &sendResolved
	this.To = to
	return &this
}

// NewEmailConfigWithDefaults instantiates a new EmailConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailConfigWithDefaults() *EmailConfig {
	this := EmailConfig{}
	var sendResolved bool = false
	this.SendResolved = &sendResolved
	return &this
}

// GetAuthIdentity returns the AuthIdentity field value if set, zero value otherwise.
func (o *EmailConfig) GetAuthIdentity() *string {
	if o == nil || IsNil(o.AuthIdentity) {
		var ret *string
		return ret
	}
	return o.AuthIdentity
}

// GetAuthIdentityOk returns a tuple with the AuthIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetAuthIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.AuthIdentity) {
		return nil, false
	}
	return o.AuthIdentity, true
}

// HasAuthIdentity returns a boolean if a field has been set.
func (o *EmailConfig) HasAuthIdentity() bool {
	if o != nil && !IsNil(o.AuthIdentity) {
		return true
	}

	return false
}

// SetAuthIdentity gets a reference to the given string and assigns it to the AuthIdentity field.
func (o *EmailConfig) SetAuthIdentity(v *string) {
	o.AuthIdentity = v
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *EmailConfig) GetAuthPassword() *string {
	if o == nil || IsNil(o.AuthPassword) {
		var ret *string
		return ret
	}
	return o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetAuthPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPassword) {
		return nil, false
	}
	return o.AuthPassword, true
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *EmailConfig) HasAuthPassword() bool {
	if o != nil && !IsNil(o.AuthPassword) {
		return true
	}

	return false
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *EmailConfig) SetAuthPassword(v *string) {
	o.AuthPassword = v
}

// GetAuthUsername returns the AuthUsername field value if set, zero value otherwise.
func (o *EmailConfig) GetAuthUsername() *string {
	if o == nil || IsNil(o.AuthUsername) {
		var ret *string
		return ret
	}
	return o.AuthUsername
}

// GetAuthUsernameOk returns a tuple with the AuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetAuthUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthUsername) {
		return nil, false
	}
	return o.AuthUsername, true
}

// HasAuthUsername returns a boolean if a field has been set.
func (o *EmailConfig) HasAuthUsername() bool {
	if o != nil && !IsNil(o.AuthUsername) {
		return true
	}

	return false
}

// SetAuthUsername gets a reference to the given string and assigns it to the AuthUsername field.
func (o *EmailConfig) SetAuthUsername(v *string) {
	o.AuthUsername = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *EmailConfig) GetFrom() *string {
	if o == nil || IsNil(o.From) {
		var ret *string
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *EmailConfig) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *EmailConfig) SetFrom(v *string) {
	o.From = v
}

// GetSendResolved returns the SendResolved field value if set, zero value otherwise.
func (o *EmailConfig) GetSendResolved() *bool {
	if o == nil || IsNil(o.SendResolved) {
		var ret *bool
		return ret
	}
	return o.SendResolved
}

// GetSendResolvedOk returns a tuple with the SendResolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetSendResolvedOk() (*bool, bool) {
	if o == nil || IsNil(o.SendResolved) {
		return nil, false
	}
	return o.SendResolved, true
}

// HasSendResolved returns a boolean if a field has been set.
func (o *EmailConfig) HasSendResolved() bool {
	if o != nil && !IsNil(o.SendResolved) {
		return true
	}

	return false
}

// SetSendResolved gets a reference to the given bool and assigns it to the SendResolved field.
func (o *EmailConfig) SetSendResolved(v *bool) {
	o.SendResolved = v
}

// GetSmarthost returns the Smarthost field value if set, zero value otherwise.
func (o *EmailConfig) GetSmarthost() *string {
	if o == nil || IsNil(o.Smarthost) {
		var ret *string
		return ret
	}
	return o.Smarthost
}

// GetSmarthostOk returns a tuple with the Smarthost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetSmarthostOk() (*string, bool) {
	if o == nil || IsNil(o.Smarthost) {
		return nil, false
	}
	return o.Smarthost, true
}

// HasSmarthost returns a boolean if a field has been set.
func (o *EmailConfig) HasSmarthost() bool {
	if o != nil && !IsNil(o.Smarthost) {
		return true
	}

	return false
}

// SetSmarthost gets a reference to the given string and assigns it to the Smarthost field.
func (o *EmailConfig) SetSmarthost(v *string) {
	o.Smarthost = v
}

// GetTo returns the To field value
func (o *EmailConfig) GetTo() *string {
	if o == nil || IsNil(o.To) {
		var ret *string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.To, true
}

// SetTo sets field value
func (o *EmailConfig) SetTo(v *string) {
	o.To = v
}

func (o EmailConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthIdentity) {
		toSerialize["authIdentity"] = o.AuthIdentity
	}
	if !IsNil(o.AuthPassword) {
		toSerialize["authPassword"] = o.AuthPassword
	}
	if !IsNil(o.AuthUsername) {
		toSerialize["authUsername"] = o.AuthUsername
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.SendResolved) {
		toSerialize["sendResolved"] = o.SendResolved
	}
	if !IsNil(o.Smarthost) {
		toSerialize["smarthost"] = o.Smarthost
	}
	toSerialize["to"] = o.To
	return toSerialize, nil
}

type NullableEmailConfig struct {
	value *EmailConfig
	isSet bool
}

func (v NullableEmailConfig) Get() *EmailConfig {
	return v.value
}

func (v *NullableEmailConfig) Set(val *EmailConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailConfig(val *EmailConfig) *NullableEmailConfig {
	return &NullableEmailConfig{value: val, isSet: true}
}

func (v NullableEmailConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
