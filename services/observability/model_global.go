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

// checks if the Global type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Global{}

/*
	types and functions for opsgenieApiKey
*/

// isNotNullableString
type GlobalGetOpsgenieApiKeyAttributeType = *string

func getGlobalGetOpsgenieApiKeyAttributeTypeOk(arg GlobalGetOpsgenieApiKeyAttributeType) (ret GlobalGetOpsgenieApiKeyRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGlobalGetOpsgenieApiKeyAttributeType(arg *GlobalGetOpsgenieApiKeyAttributeType, val GlobalGetOpsgenieApiKeyRetType) {
	*arg = &val
}

type GlobalGetOpsgenieApiKeyArgType = string
type GlobalGetOpsgenieApiKeyRetType = string

/*
	types and functions for opsgenieApiUrl
*/

// isNotNullableString
type GlobalGetOpsgenieApiUrlAttributeType = *string

func getGlobalGetOpsgenieApiUrlAttributeTypeOk(arg GlobalGetOpsgenieApiUrlAttributeType) (ret GlobalGetOpsgenieApiUrlRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGlobalGetOpsgenieApiUrlAttributeType(arg *GlobalGetOpsgenieApiUrlAttributeType, val GlobalGetOpsgenieApiUrlRetType) {
	*arg = &val
}

type GlobalGetOpsgenieApiUrlArgType = string
type GlobalGetOpsgenieApiUrlRetType = string

/*
	types and functions for resolveTimeout
*/

// isNotNullableString
type GlobalGetResolveTimeoutAttributeType = *string

func getGlobalGetResolveTimeoutAttributeTypeOk(arg GlobalGetResolveTimeoutAttributeType) (ret GlobalGetResolveTimeoutRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGlobalGetResolveTimeoutAttributeType(arg *GlobalGetResolveTimeoutAttributeType, val GlobalGetResolveTimeoutRetType) {
	*arg = &val
}

type GlobalGetResolveTimeoutArgType = string
type GlobalGetResolveTimeoutRetType = string

/*
	types and functions for smtpAuthIdentity
*/

// isNotNullableString
type GlobalGetSmtpAuthIdentityAttributeType = *string

func getGlobalGetSmtpAuthIdentityAttributeTypeOk(arg GlobalGetSmtpAuthIdentityAttributeType) (ret GlobalGetSmtpAuthIdentityRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGlobalGetSmtpAuthIdentityAttributeType(arg *GlobalGetSmtpAuthIdentityAttributeType, val GlobalGetSmtpAuthIdentityRetType) {
	*arg = &val
}

type GlobalGetSmtpAuthIdentityArgType = string
type GlobalGetSmtpAuthIdentityRetType = string

/*
	types and functions for smtpAuthPassword
*/

// isNotNullableString
type GlobalGetSmtpAuthPasswordAttributeType = *string

func getGlobalGetSmtpAuthPasswordAttributeTypeOk(arg GlobalGetSmtpAuthPasswordAttributeType) (ret GlobalGetSmtpAuthPasswordRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGlobalGetSmtpAuthPasswordAttributeType(arg *GlobalGetSmtpAuthPasswordAttributeType, val GlobalGetSmtpAuthPasswordRetType) {
	*arg = &val
}

type GlobalGetSmtpAuthPasswordArgType = string
type GlobalGetSmtpAuthPasswordRetType = string

/*
	types and functions for smtpAuthUsername
*/

// isNotNullableString
type GlobalGetSmtpAuthUsernameAttributeType = *string

func getGlobalGetSmtpAuthUsernameAttributeTypeOk(arg GlobalGetSmtpAuthUsernameAttributeType) (ret GlobalGetSmtpAuthUsernameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGlobalGetSmtpAuthUsernameAttributeType(arg *GlobalGetSmtpAuthUsernameAttributeType, val GlobalGetSmtpAuthUsernameRetType) {
	*arg = &val
}

type GlobalGetSmtpAuthUsernameArgType = string
type GlobalGetSmtpAuthUsernameRetType = string

/*
	types and functions for smtpFrom
*/

// isNotNullableString
type GlobalGetSmtpFromAttributeType = *string

func getGlobalGetSmtpFromAttributeTypeOk(arg GlobalGetSmtpFromAttributeType) (ret GlobalGetSmtpFromRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGlobalGetSmtpFromAttributeType(arg *GlobalGetSmtpFromAttributeType, val GlobalGetSmtpFromRetType) {
	*arg = &val
}

type GlobalGetSmtpFromArgType = string
type GlobalGetSmtpFromRetType = string

/*
	types and functions for smtpSmarthost
*/

// isNotNullableString
type GlobalGetSmtpSmarthostAttributeType = *string

func getGlobalGetSmtpSmarthostAttributeTypeOk(arg GlobalGetSmtpSmarthostAttributeType) (ret GlobalGetSmtpSmarthostRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setGlobalGetSmtpSmarthostAttributeType(arg *GlobalGetSmtpSmarthostAttributeType, val GlobalGetSmtpSmarthostRetType) {
	*arg = &val
}

type GlobalGetSmtpSmarthostArgType = string
type GlobalGetSmtpSmarthostRetType = string

// Global struct for Global
type Global struct {
	OpsgenieApiKey   GlobalGetOpsgenieApiKeyAttributeType   `json:"opsgenieApiKey,omitempty"`
	OpsgenieApiUrl   GlobalGetOpsgenieApiUrlAttributeType   `json:"opsgenieApiUrl,omitempty"`
	ResolveTimeout   GlobalGetResolveTimeoutAttributeType   `json:"resolveTimeout,omitempty"`
	SmtpAuthIdentity GlobalGetSmtpAuthIdentityAttributeType `json:"smtpAuthIdentity,omitempty"`
	SmtpAuthPassword GlobalGetSmtpAuthPasswordAttributeType `json:"smtpAuthPassword,omitempty"`
	SmtpAuthUsername GlobalGetSmtpAuthUsernameAttributeType `json:"smtpAuthUsername,omitempty"`
	SmtpFrom         GlobalGetSmtpFromAttributeType         `json:"smtpFrom,omitempty"`
	SmtpSmarthost    GlobalGetSmtpSmarthostAttributeType    `json:"smtpSmarthost,omitempty"`
}

// NewGlobal instantiates a new Global object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobal() *Global {
	this := Global{}
	return &this
}

// NewGlobalWithDefaults instantiates a new Global object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalWithDefaults() *Global {
	this := Global{}
	var resolveTimeout string = "5m"
	this.ResolveTimeout = &resolveTimeout
	return &this
}

// GetOpsgenieApiKey returns the OpsgenieApiKey field value if set, zero value otherwise.
func (o *Global) GetOpsgenieApiKey() (res GlobalGetOpsgenieApiKeyRetType) {
	res, _ = o.GetOpsgenieApiKeyOk()
	return
}

// GetOpsgenieApiKeyOk returns a tuple with the OpsgenieApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetOpsgenieApiKeyOk() (ret GlobalGetOpsgenieApiKeyRetType, ok bool) {
	return getGlobalGetOpsgenieApiKeyAttributeTypeOk(o.OpsgenieApiKey)
}

// HasOpsgenieApiKey returns a boolean if a field has been set.
func (o *Global) HasOpsgenieApiKey() bool {
	_, ok := o.GetOpsgenieApiKeyOk()
	return ok
}

// SetOpsgenieApiKey gets a reference to the given string and assigns it to the OpsgenieApiKey field.
func (o *Global) SetOpsgenieApiKey(v GlobalGetOpsgenieApiKeyRetType) {
	setGlobalGetOpsgenieApiKeyAttributeType(&o.OpsgenieApiKey, v)
}

// GetOpsgenieApiUrl returns the OpsgenieApiUrl field value if set, zero value otherwise.
func (o *Global) GetOpsgenieApiUrl() (res GlobalGetOpsgenieApiUrlRetType) {
	res, _ = o.GetOpsgenieApiUrlOk()
	return
}

// GetOpsgenieApiUrlOk returns a tuple with the OpsgenieApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetOpsgenieApiUrlOk() (ret GlobalGetOpsgenieApiUrlRetType, ok bool) {
	return getGlobalGetOpsgenieApiUrlAttributeTypeOk(o.OpsgenieApiUrl)
}

// HasOpsgenieApiUrl returns a boolean if a field has been set.
func (o *Global) HasOpsgenieApiUrl() bool {
	_, ok := o.GetOpsgenieApiUrlOk()
	return ok
}

// SetOpsgenieApiUrl gets a reference to the given string and assigns it to the OpsgenieApiUrl field.
func (o *Global) SetOpsgenieApiUrl(v GlobalGetOpsgenieApiUrlRetType) {
	setGlobalGetOpsgenieApiUrlAttributeType(&o.OpsgenieApiUrl, v)
}

// GetResolveTimeout returns the ResolveTimeout field value if set, zero value otherwise.
func (o *Global) GetResolveTimeout() (res GlobalGetResolveTimeoutRetType) {
	res, _ = o.GetResolveTimeoutOk()
	return
}

// GetResolveTimeoutOk returns a tuple with the ResolveTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetResolveTimeoutOk() (ret GlobalGetResolveTimeoutRetType, ok bool) {
	return getGlobalGetResolveTimeoutAttributeTypeOk(o.ResolveTimeout)
}

// HasResolveTimeout returns a boolean if a field has been set.
func (o *Global) HasResolveTimeout() bool {
	_, ok := o.GetResolveTimeoutOk()
	return ok
}

// SetResolveTimeout gets a reference to the given string and assigns it to the ResolveTimeout field.
func (o *Global) SetResolveTimeout(v GlobalGetResolveTimeoutRetType) {
	setGlobalGetResolveTimeoutAttributeType(&o.ResolveTimeout, v)
}

// GetSmtpAuthIdentity returns the SmtpAuthIdentity field value if set, zero value otherwise.
func (o *Global) GetSmtpAuthIdentity() (res GlobalGetSmtpAuthIdentityRetType) {
	res, _ = o.GetSmtpAuthIdentityOk()
	return
}

// GetSmtpAuthIdentityOk returns a tuple with the SmtpAuthIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetSmtpAuthIdentityOk() (ret GlobalGetSmtpAuthIdentityRetType, ok bool) {
	return getGlobalGetSmtpAuthIdentityAttributeTypeOk(o.SmtpAuthIdentity)
}

// HasSmtpAuthIdentity returns a boolean if a field has been set.
func (o *Global) HasSmtpAuthIdentity() bool {
	_, ok := o.GetSmtpAuthIdentityOk()
	return ok
}

// SetSmtpAuthIdentity gets a reference to the given string and assigns it to the SmtpAuthIdentity field.
func (o *Global) SetSmtpAuthIdentity(v GlobalGetSmtpAuthIdentityRetType) {
	setGlobalGetSmtpAuthIdentityAttributeType(&o.SmtpAuthIdentity, v)
}

// GetSmtpAuthPassword returns the SmtpAuthPassword field value if set, zero value otherwise.
func (o *Global) GetSmtpAuthPassword() (res GlobalGetSmtpAuthPasswordRetType) {
	res, _ = o.GetSmtpAuthPasswordOk()
	return
}

// GetSmtpAuthPasswordOk returns a tuple with the SmtpAuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetSmtpAuthPasswordOk() (ret GlobalGetSmtpAuthPasswordRetType, ok bool) {
	return getGlobalGetSmtpAuthPasswordAttributeTypeOk(o.SmtpAuthPassword)
}

// HasSmtpAuthPassword returns a boolean if a field has been set.
func (o *Global) HasSmtpAuthPassword() bool {
	_, ok := o.GetSmtpAuthPasswordOk()
	return ok
}

// SetSmtpAuthPassword gets a reference to the given string and assigns it to the SmtpAuthPassword field.
func (o *Global) SetSmtpAuthPassword(v GlobalGetSmtpAuthPasswordRetType) {
	setGlobalGetSmtpAuthPasswordAttributeType(&o.SmtpAuthPassword, v)
}

// GetSmtpAuthUsername returns the SmtpAuthUsername field value if set, zero value otherwise.
func (o *Global) GetSmtpAuthUsername() (res GlobalGetSmtpAuthUsernameRetType) {
	res, _ = o.GetSmtpAuthUsernameOk()
	return
}

// GetSmtpAuthUsernameOk returns a tuple with the SmtpAuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetSmtpAuthUsernameOk() (ret GlobalGetSmtpAuthUsernameRetType, ok bool) {
	return getGlobalGetSmtpAuthUsernameAttributeTypeOk(o.SmtpAuthUsername)
}

// HasSmtpAuthUsername returns a boolean if a field has been set.
func (o *Global) HasSmtpAuthUsername() bool {
	_, ok := o.GetSmtpAuthUsernameOk()
	return ok
}

// SetSmtpAuthUsername gets a reference to the given string and assigns it to the SmtpAuthUsername field.
func (o *Global) SetSmtpAuthUsername(v GlobalGetSmtpAuthUsernameRetType) {
	setGlobalGetSmtpAuthUsernameAttributeType(&o.SmtpAuthUsername, v)
}

// GetSmtpFrom returns the SmtpFrom field value if set, zero value otherwise.
func (o *Global) GetSmtpFrom() (res GlobalGetSmtpFromRetType) {
	res, _ = o.GetSmtpFromOk()
	return
}

// GetSmtpFromOk returns a tuple with the SmtpFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetSmtpFromOk() (ret GlobalGetSmtpFromRetType, ok bool) {
	return getGlobalGetSmtpFromAttributeTypeOk(o.SmtpFrom)
}

// HasSmtpFrom returns a boolean if a field has been set.
func (o *Global) HasSmtpFrom() bool {
	_, ok := o.GetSmtpFromOk()
	return ok
}

// SetSmtpFrom gets a reference to the given string and assigns it to the SmtpFrom field.
func (o *Global) SetSmtpFrom(v GlobalGetSmtpFromRetType) {
	setGlobalGetSmtpFromAttributeType(&o.SmtpFrom, v)
}

// GetSmtpSmarthost returns the SmtpSmarthost field value if set, zero value otherwise.
func (o *Global) GetSmtpSmarthost() (res GlobalGetSmtpSmarthostRetType) {
	res, _ = o.GetSmtpSmarthostOk()
	return
}

// GetSmtpSmarthostOk returns a tuple with the SmtpSmarthost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Global) GetSmtpSmarthostOk() (ret GlobalGetSmtpSmarthostRetType, ok bool) {
	return getGlobalGetSmtpSmarthostAttributeTypeOk(o.SmtpSmarthost)
}

// HasSmtpSmarthost returns a boolean if a field has been set.
func (o *Global) HasSmtpSmarthost() bool {
	_, ok := o.GetSmtpSmarthostOk()
	return ok
}

// SetSmtpSmarthost gets a reference to the given string and assigns it to the SmtpSmarthost field.
func (o *Global) SetSmtpSmarthost(v GlobalGetSmtpSmarthostRetType) {
	setGlobalGetSmtpSmarthostAttributeType(&o.SmtpSmarthost, v)
}

func (o Global) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getGlobalGetOpsgenieApiKeyAttributeTypeOk(o.OpsgenieApiKey); ok {
		toSerialize["OpsgenieApiKey"] = val
	}
	if val, ok := getGlobalGetOpsgenieApiUrlAttributeTypeOk(o.OpsgenieApiUrl); ok {
		toSerialize["OpsgenieApiUrl"] = val
	}
	if val, ok := getGlobalGetResolveTimeoutAttributeTypeOk(o.ResolveTimeout); ok {
		toSerialize["ResolveTimeout"] = val
	}
	if val, ok := getGlobalGetSmtpAuthIdentityAttributeTypeOk(o.SmtpAuthIdentity); ok {
		toSerialize["SmtpAuthIdentity"] = val
	}
	if val, ok := getGlobalGetSmtpAuthPasswordAttributeTypeOk(o.SmtpAuthPassword); ok {
		toSerialize["SmtpAuthPassword"] = val
	}
	if val, ok := getGlobalGetSmtpAuthUsernameAttributeTypeOk(o.SmtpAuthUsername); ok {
		toSerialize["SmtpAuthUsername"] = val
	}
	if val, ok := getGlobalGetSmtpFromAttributeTypeOk(o.SmtpFrom); ok {
		toSerialize["SmtpFrom"] = val
	}
	if val, ok := getGlobalGetSmtpSmarthostAttributeTypeOk(o.SmtpSmarthost); ok {
		toSerialize["SmtpSmarthost"] = val
	}
	return toSerialize, nil
}

type NullableGlobal struct {
	value *Global
	isSet bool
}

func (v NullableGlobal) Get() *Global {
	return v.value
}

func (v *NullableGlobal) Set(val *Global) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobal) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobal(val *Global) *NullableGlobal {
	return &NullableGlobal{value: val, isSet: true}
}

func (v NullableGlobal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
