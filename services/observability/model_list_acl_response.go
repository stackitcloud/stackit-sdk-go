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

// checks if the ListACLResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListACLResponse{}

// ListACLResponse struct for ListACLResponse
type ListACLResponse struct {
	// REQUIRED
	Acl *[]string `json:"acl"`
	// REQUIRED
	Message *string `json:"message"`
}

type _ListACLResponse ListACLResponse

// NewListACLResponse instantiates a new ListACLResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListACLResponse(acl *[]string, message *string) *ListACLResponse {
	this := ListACLResponse{}
	this.Acl = acl
	this.Message = message
	return &this
}

// NewListACLResponseWithDefaults instantiates a new ListACLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListACLResponseWithDefaults() *ListACLResponse {
	this := ListACLResponse{}
	return &this
}

// GetAcl returns the Acl field value
func (o *ListACLResponse) GetAcl() *[]string {
	if o == nil || IsNil(o.Acl) {
		var ret *[]string
		return ret
	}

	return o.Acl
}

// GetAclOk returns a tuple with the Acl field value
// and a boolean to check if the value has been set.
func (o *ListACLResponse) GetAclOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Acl, true
}

// SetAcl sets field value
func (o *ListACLResponse) SetAcl(v *[]string) {
	o.Acl = v
}

// GetMessage returns the Message field value
func (o *ListACLResponse) GetMessage() *string {
	if o == nil || IsNil(o.Message) {
		var ret *string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ListACLResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *ListACLResponse) SetMessage(v *string) {
	o.Message = v
}

func (o ListACLResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acl"] = o.Acl
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableListACLResponse struct {
	value *ListACLResponse
	isSet bool
}

func (v NullableListACLResponse) Get() *ListACLResponse {
	return v.value
}

func (v *NullableListACLResponse) Set(val *ListACLResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListACLResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListACLResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListACLResponse(val *ListACLResponse) *NullableListACLResponse {
	return &NullableListACLResponse{value: val, isSet: true}
}

func (v NullableListACLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListACLResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
