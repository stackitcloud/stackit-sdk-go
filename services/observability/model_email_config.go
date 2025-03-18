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

/*
	types and functions for authIdentity
*/

// isNotNullableString
type EmailConfigGetAuthIdentityAttributeType = *string

func getEmailConfigGetAuthIdentityAttributeTypeOk(arg EmailConfigGetAuthIdentityAttributeType) (ret EmailConfigGetAuthIdentityRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setEmailConfigGetAuthIdentityAttributeType(arg *EmailConfigGetAuthIdentityAttributeType, val EmailConfigGetAuthIdentityRetType) {
	*arg = &val
}

type EmailConfigGetAuthIdentityArgType = string
type EmailConfigGetAuthIdentityRetType = string

/*
	types and functions for authPassword
*/

// isNotNullableString
type EmailConfigGetAuthPasswordAttributeType = *string

func getEmailConfigGetAuthPasswordAttributeTypeOk(arg EmailConfigGetAuthPasswordAttributeType) (ret EmailConfigGetAuthPasswordRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setEmailConfigGetAuthPasswordAttributeType(arg *EmailConfigGetAuthPasswordAttributeType, val EmailConfigGetAuthPasswordRetType) {
	*arg = &val
}

type EmailConfigGetAuthPasswordArgType = string
type EmailConfigGetAuthPasswordRetType = string

/*
	types and functions for authUsername
*/

// isNotNullableString
type EmailConfigGetAuthUsernameAttributeType = *string

func getEmailConfigGetAuthUsernameAttributeTypeOk(arg EmailConfigGetAuthUsernameAttributeType) (ret EmailConfigGetAuthUsernameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setEmailConfigGetAuthUsernameAttributeType(arg *EmailConfigGetAuthUsernameAttributeType, val EmailConfigGetAuthUsernameRetType) {
	*arg = &val
}

type EmailConfigGetAuthUsernameArgType = string
type EmailConfigGetAuthUsernameRetType = string

/*
	types and functions for from
*/

// isNotNullableString
type EmailConfigGetFromAttributeType = *string

func getEmailConfigGetFromAttributeTypeOk(arg EmailConfigGetFromAttributeType) (ret EmailConfigGetFromRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setEmailConfigGetFromAttributeType(arg *EmailConfigGetFromAttributeType, val EmailConfigGetFromRetType) {
	*arg = &val
}

type EmailConfigGetFromArgType = string
type EmailConfigGetFromRetType = string

/*
	types and functions for sendResolved
*/

// isBoolean
type EmailConfiggetSendResolvedAttributeType = *bool
type EmailConfiggetSendResolvedArgType = bool
type EmailConfiggetSendResolvedRetType = bool

func getEmailConfiggetSendResolvedAttributeTypeOk(arg EmailConfiggetSendResolvedAttributeType) (ret EmailConfiggetSendResolvedRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setEmailConfiggetSendResolvedAttributeType(arg *EmailConfiggetSendResolvedAttributeType, val EmailConfiggetSendResolvedRetType) {
	*arg = &val
}

/*
	types and functions for smarthost
*/

// isNotNullableString
type EmailConfigGetSmarthostAttributeType = *string

func getEmailConfigGetSmarthostAttributeTypeOk(arg EmailConfigGetSmarthostAttributeType) (ret EmailConfigGetSmarthostRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setEmailConfigGetSmarthostAttributeType(arg *EmailConfigGetSmarthostAttributeType, val EmailConfigGetSmarthostRetType) {
	*arg = &val
}

type EmailConfigGetSmarthostArgType = string
type EmailConfigGetSmarthostRetType = string

/*
	types and functions for to
*/

// isNotNullableString
type EmailConfigGetToAttributeType = *string

func getEmailConfigGetToAttributeTypeOk(arg EmailConfigGetToAttributeType) (ret EmailConfigGetToRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setEmailConfigGetToAttributeType(arg *EmailConfigGetToAttributeType, val EmailConfigGetToRetType) {
	*arg = &val
}

type EmailConfigGetToArgType = string
type EmailConfigGetToRetType = string

// EmailConfig struct for EmailConfig
type EmailConfig struct {
	AuthIdentity EmailConfigGetAuthIdentityAttributeType `json:"authIdentity,omitempty"`
	AuthPassword EmailConfigGetAuthPasswordAttributeType `json:"authPassword,omitempty"`
	AuthUsername EmailConfigGetAuthUsernameAttributeType `json:"authUsername,omitempty"`
	From         EmailConfigGetFromAttributeType         `json:"from,omitempty"`
	SendResolved EmailConfiggetSendResolvedAttributeType `json:"sendResolved,omitempty"`
	Smarthost    EmailConfigGetSmarthostAttributeType    `json:"smarthost,omitempty"`
	// REQUIRED
	To EmailConfigGetToAttributeType `json:"to"`
}

type _EmailConfig EmailConfig

// NewEmailConfig instantiates a new EmailConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailConfig(to EmailConfigGetToArgType) *EmailConfig {
	this := EmailConfig{}
	setEmailConfigGetToAttributeType(&this.To, to)
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
func (o *EmailConfig) GetAuthIdentity() (res EmailConfigGetAuthIdentityRetType) {
	res, _ = o.GetAuthIdentityOk()
	return
}

// GetAuthIdentityOk returns a tuple with the AuthIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetAuthIdentityOk() (ret EmailConfigGetAuthIdentityRetType, ok bool) {
	return getEmailConfigGetAuthIdentityAttributeTypeOk(o.AuthIdentity)
}

// HasAuthIdentity returns a boolean if a field has been set.
func (o *EmailConfig) HasAuthIdentity() bool {
	_, ok := o.GetAuthIdentityOk()
	return ok
}

// SetAuthIdentity gets a reference to the given string and assigns it to the AuthIdentity field.
func (o *EmailConfig) SetAuthIdentity(v EmailConfigGetAuthIdentityRetType) {
	setEmailConfigGetAuthIdentityAttributeType(&o.AuthIdentity, v)
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *EmailConfig) GetAuthPassword() (res EmailConfigGetAuthPasswordRetType) {
	res, _ = o.GetAuthPasswordOk()
	return
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetAuthPasswordOk() (ret EmailConfigGetAuthPasswordRetType, ok bool) {
	return getEmailConfigGetAuthPasswordAttributeTypeOk(o.AuthPassword)
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *EmailConfig) HasAuthPassword() bool {
	_, ok := o.GetAuthPasswordOk()
	return ok
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *EmailConfig) SetAuthPassword(v EmailConfigGetAuthPasswordRetType) {
	setEmailConfigGetAuthPasswordAttributeType(&o.AuthPassword, v)
}

// GetAuthUsername returns the AuthUsername field value if set, zero value otherwise.
func (o *EmailConfig) GetAuthUsername() (res EmailConfigGetAuthUsernameRetType) {
	res, _ = o.GetAuthUsernameOk()
	return
}

// GetAuthUsernameOk returns a tuple with the AuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetAuthUsernameOk() (ret EmailConfigGetAuthUsernameRetType, ok bool) {
	return getEmailConfigGetAuthUsernameAttributeTypeOk(o.AuthUsername)
}

// HasAuthUsername returns a boolean if a field has been set.
func (o *EmailConfig) HasAuthUsername() bool {
	_, ok := o.GetAuthUsernameOk()
	return ok
}

// SetAuthUsername gets a reference to the given string and assigns it to the AuthUsername field.
func (o *EmailConfig) SetAuthUsername(v EmailConfigGetAuthUsernameRetType) {
	setEmailConfigGetAuthUsernameAttributeType(&o.AuthUsername, v)
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *EmailConfig) GetFrom() (res EmailConfigGetFromRetType) {
	res, _ = o.GetFromOk()
	return
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetFromOk() (ret EmailConfigGetFromRetType, ok bool) {
	return getEmailConfigGetFromAttributeTypeOk(o.From)
}

// HasFrom returns a boolean if a field has been set.
func (o *EmailConfig) HasFrom() bool {
	_, ok := o.GetFromOk()
	return ok
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *EmailConfig) SetFrom(v EmailConfigGetFromRetType) {
	setEmailConfigGetFromAttributeType(&o.From, v)
}

// GetSendResolved returns the SendResolved field value if set, zero value otherwise.
func (o *EmailConfig) GetSendResolved() (res EmailConfiggetSendResolvedRetType) {
	res, _ = o.GetSendResolvedOk()
	return
}

// GetSendResolvedOk returns a tuple with the SendResolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetSendResolvedOk() (ret EmailConfiggetSendResolvedRetType, ok bool) {
	return getEmailConfiggetSendResolvedAttributeTypeOk(o.SendResolved)
}

// HasSendResolved returns a boolean if a field has been set.
func (o *EmailConfig) HasSendResolved() bool {
	_, ok := o.GetSendResolvedOk()
	return ok
}

// SetSendResolved gets a reference to the given bool and assigns it to the SendResolved field.
func (o *EmailConfig) SetSendResolved(v EmailConfiggetSendResolvedRetType) {
	setEmailConfiggetSendResolvedAttributeType(&o.SendResolved, v)
}

// GetSmarthost returns the Smarthost field value if set, zero value otherwise.
func (o *EmailConfig) GetSmarthost() (res EmailConfigGetSmarthostRetType) {
	res, _ = o.GetSmarthostOk()
	return
}

// GetSmarthostOk returns a tuple with the Smarthost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetSmarthostOk() (ret EmailConfigGetSmarthostRetType, ok bool) {
	return getEmailConfigGetSmarthostAttributeTypeOk(o.Smarthost)
}

// HasSmarthost returns a boolean if a field has been set.
func (o *EmailConfig) HasSmarthost() bool {
	_, ok := o.GetSmarthostOk()
	return ok
}

// SetSmarthost gets a reference to the given string and assigns it to the Smarthost field.
func (o *EmailConfig) SetSmarthost(v EmailConfigGetSmarthostRetType) {
	setEmailConfigGetSmarthostAttributeType(&o.Smarthost, v)
}

// GetTo returns the To field value
func (o *EmailConfig) GetTo() (ret EmailConfigGetToRetType) {
	ret, _ = o.GetToOk()
	return ret
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetToOk() (ret EmailConfigGetToRetType, ok bool) {
	return getEmailConfigGetToAttributeTypeOk(o.To)
}

// SetTo sets field value
func (o *EmailConfig) SetTo(v EmailConfigGetToRetType) {
	setEmailConfigGetToAttributeType(&o.To, v)
}

func (o EmailConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getEmailConfigGetAuthIdentityAttributeTypeOk(o.AuthIdentity); ok {
		toSerialize["AuthIdentity"] = val
	}
	if val, ok := getEmailConfigGetAuthPasswordAttributeTypeOk(o.AuthPassword); ok {
		toSerialize["AuthPassword"] = val
	}
	if val, ok := getEmailConfigGetAuthUsernameAttributeTypeOk(o.AuthUsername); ok {
		toSerialize["AuthUsername"] = val
	}
	if val, ok := getEmailConfigGetFromAttributeTypeOk(o.From); ok {
		toSerialize["From"] = val
	}
	if val, ok := getEmailConfiggetSendResolvedAttributeTypeOk(o.SendResolved); ok {
		toSerialize["SendResolved"] = val
	}
	if val, ok := getEmailConfigGetSmarthostAttributeTypeOk(o.Smarthost); ok {
		toSerialize["Smarthost"] = val
	}
	if val, ok := getEmailConfigGetToAttributeTypeOk(o.To); ok {
		toSerialize["To"] = val
	}
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
