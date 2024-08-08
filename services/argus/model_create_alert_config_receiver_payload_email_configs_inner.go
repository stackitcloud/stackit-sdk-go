/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

import (
	"encoding/json"
)

// checks if the CreateAlertConfigReceiverPayloadEmailConfigsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAlertConfigReceiverPayloadEmailConfigsInner{}

// CreateAlertConfigReceiverPayloadEmailConfigsInner struct for CreateAlertConfigReceiverPayloadEmailConfigsInner
type CreateAlertConfigReceiverPayloadEmailConfigsInner struct {
	// SMTP authentication information. `Additional Validators:` * must be a syntactically valid email address
	AuthIdentity *string `json:"authIdentity,omitempty"`
	// SMTP authentication information.
	AuthPassword *string `json:"authPassword,omitempty"`
	// SMTP authentication information.
	AuthUsername *string `json:"authUsername,omitempty"`
	// The sender address. `Additional Validators:` * must be a syntactically valid email address
	From *string `json:"from,omitempty"`
	// The SMTP host through which emails are sent. `Additional Validators:` * should only include the characters: a-zA-Z0-9_./@&?:-
	Smarthost *string `json:"smarthost,omitempty"`
	// The email address to send notifications to. `Additional Validators:` * must be a syntactically valid email address
	To *string `json:"to,omitempty"`
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
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetAuthIdentity() *string {
	if o == nil || IsNil(o.AuthIdentity) {
		var ret *string
		return ret
	}
	return o.AuthIdentity
}

// GetAuthIdentityOk returns a tuple with the AuthIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetAuthIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.AuthIdentity) {
		return nil, false
	}
	return o.AuthIdentity, true
}

// HasAuthIdentity returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) HasAuthIdentity() bool {
	if o != nil && !IsNil(o.AuthIdentity) {
		return true
	}

	return false
}

// SetAuthIdentity gets a reference to the given string and assigns it to the AuthIdentity field.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) SetAuthIdentity(v *string) {
	o.AuthIdentity = v
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetAuthPassword() *string {
	if o == nil || IsNil(o.AuthPassword) {
		var ret *string
		return ret
	}
	return o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetAuthPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPassword) {
		return nil, false
	}
	return o.AuthPassword, true
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) HasAuthPassword() bool {
	if o != nil && !IsNil(o.AuthPassword) {
		return true
	}

	return false
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) SetAuthPassword(v *string) {
	o.AuthPassword = v
}

// GetAuthUsername returns the AuthUsername field value if set, zero value otherwise.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetAuthUsername() *string {
	if o == nil || IsNil(o.AuthUsername) {
		var ret *string
		return ret
	}
	return o.AuthUsername
}

// GetAuthUsernameOk returns a tuple with the AuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetAuthUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthUsername) {
		return nil, false
	}
	return o.AuthUsername, true
}

// HasAuthUsername returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) HasAuthUsername() bool {
	if o != nil && !IsNil(o.AuthUsername) {
		return true
	}

	return false
}

// SetAuthUsername gets a reference to the given string and assigns it to the AuthUsername field.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) SetAuthUsername(v *string) {
	o.AuthUsername = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetFrom() *string {
	if o == nil || IsNil(o.From) {
		var ret *string
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) SetFrom(v *string) {
	o.From = v
}

// GetSmarthost returns the Smarthost field value if set, zero value otherwise.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetSmarthost() *string {
	if o == nil || IsNil(o.Smarthost) {
		var ret *string
		return ret
	}
	return o.Smarthost
}

// GetSmarthostOk returns a tuple with the Smarthost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetSmarthostOk() (*string, bool) {
	if o == nil || IsNil(o.Smarthost) {
		return nil, false
	}
	return o.Smarthost, true
}

// HasSmarthost returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) HasSmarthost() bool {
	if o != nil && !IsNil(o.Smarthost) {
		return true
	}

	return false
}

// SetSmarthost gets a reference to the given string and assigns it to the Smarthost field.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) SetSmarthost(v *string) {
	o.Smarthost = v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetTo() *string {
	if o == nil || IsNil(o.To) {
		var ret *string
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *CreateAlertConfigReceiverPayloadEmailConfigsInner) SetTo(v *string) {
	o.To = v
}

func (o CreateAlertConfigReceiverPayloadEmailConfigsInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Smarthost) {
		toSerialize["smarthost"] = o.Smarthost
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
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
