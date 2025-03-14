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

// checks if the CreateAlertConfigReceiverPayloadEmailConfigsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAlertConfigReceiverPayloadEmailConfigsInner{}

/*
	types and functions for authIdentity
*/

// isNotNullableString
type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityAttributeType = *string

func getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityAttributeTypeOk(arg CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityAttributeType) (ret CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityAttributeType(arg *CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityAttributeType, val CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityRetType) {
	*arg = &val
}

type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityArgType = string
type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityRetType = string

/*
	types and functions for authPassword
*/

// isNotNullableString
type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordAttributeType = *string

func getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordAttributeTypeOk(arg CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordAttributeType) (ret CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordAttributeType(arg *CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordAttributeType, val CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordRetType) {
	*arg = &val
}

type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordArgType = string
type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordRetType = string

/*
	types and functions for authUsername
*/

// isNotNullableString
type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameAttributeType = *string

func getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameAttributeTypeOk(arg CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameAttributeType) (ret CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameAttributeType(arg *CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameAttributeType, val CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameRetType) {
	*arg = &val
}

type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameArgType = string
type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameRetType = string

/*
	types and functions for from
*/

// isNotNullableString
type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromAttributeType = *string

func getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromAttributeTypeOk(arg CreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromAttributeType) (ret CreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromAttributeType(arg *CreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromAttributeType, val CreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromRetType) {
	*arg = &val
}

type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromArgType = string
type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromRetType = string

/*
	types and functions for smarthost
*/

// isNotNullableString
type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostAttributeType = *string

func getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostAttributeTypeOk(arg CreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostAttributeType) (ret CreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostAttributeType(arg *CreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostAttributeType, val CreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostRetType) {
	*arg = &val
}

type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostArgType = string
type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostRetType = string

/*
	types and functions for to
*/

// isNotNullableString
type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetToAttributeType = *string

func getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetToAttributeTypeOk(arg CreateAlertConfigReceiverPayloadEmailConfigsInnerGetToAttributeType) (ret CreateAlertConfigReceiverPayloadEmailConfigsInnerGetToRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCreateAlertConfigReceiverPayloadEmailConfigsInnerGetToAttributeType(arg *CreateAlertConfigReceiverPayloadEmailConfigsInnerGetToAttributeType, val CreateAlertConfigReceiverPayloadEmailConfigsInnerGetToRetType) {
	*arg = &val
}

type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetToArgType = string
type CreateAlertConfigReceiverPayloadEmailConfigsInnerGetToRetType = string

// CreateAlertConfigReceiverPayloadEmailConfigsInner struct for CreateAlertConfigReceiverPayloadEmailConfigsInner
type CreateAlertConfigReceiverPayloadEmailConfigsInner struct {
	// SMTP authentication information. `Additional Validators:` * must be a syntactically valid email address
	AuthIdentity CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityAttributeType `json:"authIdentity,omitempty"`
	// SMTP authentication information.
	AuthPassword CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordAttributeType `json:"authPassword,omitempty"`
	// SMTP authentication information.
	AuthUsername CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameAttributeType `json:"authUsername,omitempty"`
	// The sender address. `Additional Validators:` * must be a syntactically valid email address
	From CreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromAttributeType `json:"from,omitempty"`
	// The SMTP host through which emails are sent. `Additional Validators:` * should only include the characters: a-zA-Z0-9_./@&?:-
	Smarthost CreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostAttributeType `json:"smarthost,omitempty"`
	// The email address to send notifications to. `Additional Validators:` * must be a syntactically valid email address
	To CreateAlertConfigReceiverPayloadEmailConfigsInnerGetToAttributeType `json:"to,omitempty"`
}

// NewCreateAlertConfigReceiverPayloadEmailConfigsInner instantiates a new CreateAlertConfigReceiverPayloadEmailConfigsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAlertConfigReceiverPayloadEmailConfigsInner() *CreateAlertConfigReceiverPayloadEmailConfigsInner {
	this := CreateAlertConfigReceiverPayloadEmailConfigsInner{}
	return &this
}

// NewCreateAlertConfigReceiverPayloadEmailConfigsInnerWithDefaults instantiates a new CreateAlertConfigReceiverPayloadEmailConfigsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAlertConfigReceiverPayloadEmailConfigsInnerWithDefaults() *CreateAlertConfigReceiverPayloadEmailConfigsInner {
	this := CreateAlertConfigReceiverPayloadEmailConfigsInner{}
	return &this
}

// GetAuthIdentity returns the AuthIdentity field value if set, zero value otherwise.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetAuthIdentity() (res CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityRetType) {
	res, _ = o.GetAuthIdentityOk()
	return
}

// GetAuthIdentityOk returns a tuple with the AuthIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetAuthIdentityOk() (ret CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityRetType, ok bool) {
	return getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityAttributeTypeOk(o.AuthIdentity)
}

// HasAuthIdentity returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) HasAuthIdentity() bool {
	_, ok := o.GetAuthIdentityOk()
	return ok
}

// SetAuthIdentity gets a reference to the given string and assigns it to the AuthIdentity field.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) SetAuthIdentity(v CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityRetType) {
	setCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityAttributeType(&o.AuthIdentity, v)
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetAuthPassword() (res CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordRetType) {
	res, _ = o.GetAuthPasswordOk()
	return
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetAuthPasswordOk() (ret CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordRetType, ok bool) {
	return getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordAttributeTypeOk(o.AuthPassword)
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) HasAuthPassword() bool {
	_, ok := o.GetAuthPasswordOk()
	return ok
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) SetAuthPassword(v CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordRetType) {
	setCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordAttributeType(&o.AuthPassword, v)
}

// GetAuthUsername returns the AuthUsername field value if set, zero value otherwise.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetAuthUsername() (res CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameRetType) {
	res, _ = o.GetAuthUsernameOk()
	return
}

// GetAuthUsernameOk returns a tuple with the AuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetAuthUsernameOk() (ret CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameRetType, ok bool) {
	return getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameAttributeTypeOk(o.AuthUsername)
}

// HasAuthUsername returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) HasAuthUsername() bool {
	_, ok := o.GetAuthUsernameOk()
	return ok
}

// SetAuthUsername gets a reference to the given string and assigns it to the AuthUsername field.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) SetAuthUsername(v CreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameRetType) {
	setCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameAttributeType(&o.AuthUsername, v)
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetFrom() (res CreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromRetType) {
	res, _ = o.GetFromOk()
	return
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetFromOk() (ret CreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromRetType, ok bool) {
	return getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromAttributeTypeOk(o.From)
}

// HasFrom returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) HasFrom() bool {
	_, ok := o.GetFromOk()
	return ok
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) SetFrom(v CreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromRetType) {
	setCreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromAttributeType(&o.From, v)
}

// GetSmarthost returns the Smarthost field value if set, zero value otherwise.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetSmarthost() (res CreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostRetType) {
	res, _ = o.GetSmarthostOk()
	return
}

// GetSmarthostOk returns a tuple with the Smarthost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetSmarthostOk() (ret CreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostRetType, ok bool) {
	return getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostAttributeTypeOk(o.Smarthost)
}

// HasSmarthost returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) HasSmarthost() bool {
	_, ok := o.GetSmarthostOk()
	return ok
}

// SetSmarthost gets a reference to the given string and assigns it to the Smarthost field.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) SetSmarthost(v CreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostRetType) {
	setCreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostAttributeType(&o.Smarthost, v)
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetTo() (res CreateAlertConfigReceiverPayloadEmailConfigsInnerGetToRetType) {
	res, _ = o.GetToOk()
	return
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetToOk() (ret CreateAlertConfigReceiverPayloadEmailConfigsInnerGetToRetType, ok bool) {
	return getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetToAttributeTypeOk(o.To)
}

// HasTo returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) HasTo() bool {
	_, ok := o.GetToOk()
	return ok
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) SetTo(v CreateAlertConfigReceiverPayloadEmailConfigsInnerGetToRetType) {
	setCreateAlertConfigReceiverPayloadEmailConfigsInnerGetToAttributeType(&o.To, v)
}

func (o CreateAlertConfigReceiverPayloadEmailConfigsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthIdentityAttributeTypeOk(o.AuthIdentity); ok {
		toSerialize["AuthIdentity"] = val
	}
	if val, ok := getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthPasswordAttributeTypeOk(o.AuthPassword); ok {
		toSerialize["AuthPassword"] = val
	}
	if val, ok := getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetAuthUsernameAttributeTypeOk(o.AuthUsername); ok {
		toSerialize["AuthUsername"] = val
	}
	if val, ok := getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetFromAttributeTypeOk(o.From); ok {
		toSerialize["From"] = val
	}
	if val, ok := getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetSmarthostAttributeTypeOk(o.Smarthost); ok {
		toSerialize["Smarthost"] = val
	}
	if val, ok := getCreateAlertConfigReceiverPayloadEmailConfigsInnerGetToAttributeTypeOk(o.To); ok {
		toSerialize["To"] = val
	}
	return toSerialize, nil
}

type NullableCreateAlertConfigReceiverPayloadEmailConfigsInner struct {
	value *CreateAlertConfigReceiverPayloadEmailConfigsInner
	isSet bool
}

func (v NullableCreateAlertConfigReceiverPayloadEmailConfigsInner) Get() *CreateAlertConfigReceiverPayloadEmailConfigsInner {
	return v.value
}

func (v *NullableCreateAlertConfigReceiverPayloadEmailConfigsInner) Set(val *CreateAlertConfigReceiverPayloadEmailConfigsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAlertConfigReceiverPayloadEmailConfigsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAlertConfigReceiverPayloadEmailConfigsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAlertConfigReceiverPayloadEmailConfigsInner(val *CreateAlertConfigReceiverPayloadEmailConfigsInner) *NullableCreateAlertConfigReceiverPayloadEmailConfigsInner {
	return &NullableCreateAlertConfigReceiverPayloadEmailConfigsInner{value: val, isSet: true}
}

func (v NullableCreateAlertConfigReceiverPayloadEmailConfigsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAlertConfigReceiverPayloadEmailConfigsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
